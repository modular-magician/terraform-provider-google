package google

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var iamBindingSchema = map[string]*schema.Schema{
	"role": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"members": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Schema{
			Type:             schema.TypeString,
			DiffSuppressFunc: caseDiffSuppress,
		},
		Set: func(v interface{}) int {
			return schema.HashString(strings.ToLower(v.(string)))
		},
	},
	"etag": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

func ResourceIamBinding(parentSpecificSchema map[string]*schema.Schema, newUpdaterFunc newResourceIamUpdaterFunc, resourceIdParser resourceIdParserFunc) *schema.Resource {
	return ResourceIamBindingWithBatching(parentSpecificSchema, newUpdaterFunc, resourceIdParser, IamBatchingDisabled)
}

// Resource that batches requests to the same IAM policy across multiple IAM fine-grained resources
func ResourceIamBindingWithBatching(parentSpecificSchema map[string]*schema.Schema, newUpdaterFunc newResourceIamUpdaterFunc, resourceIdParser resourceIdParserFunc, enableBatching bool) *schema.Resource {
	return &schema.Resource{
		Create: resourceIamBindingCreate(newUpdaterFunc, enableBatching),
		Read:   resourceIamBindingRead(newUpdaterFunc),
		Update: resourceIamBindingUpdate(newUpdaterFunc, enableBatching),
		Delete: resourceIamBindingDelete(newUpdaterFunc, enableBatching),
		Schema: mergeSchemas(iamBindingSchema, parentSpecificSchema),
		Importer: &schema.ResourceImporter{
			State: iamBindingImport(resourceIdParser),
		},
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type: resourceIAMBindingResourceV0(parentSpecificSchema).CoreConfigSchema().ImpliedType(),
				Upgrade: func(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
					// IAM conditions require parsing information from the ID, which means the ID needs
					// a separator that won't occur naturally in the project or role.
					// Previous format: projects/{project}/{resourceType}/{resourceName}/roles/{role}
					// New format: projects/{project}/{resourceType}/{resourceName} roles/{role}
					oldId := rawState["id"].(string)
					role := rawState["role"].(string)
					rawState["id"] = strings.TrimSuffix(oldId, "/"+role) + " " + role
					return rawState, nil
				},
				Version: 0,
			},
		},
	}
}

func resourceIAMBindingResourceV0(parentSpecificSchema map[string]*schema.Schema) *schema.Resource {
	return &schema.Resource{Schema: mergeSchemas(iamBindingSchema, parentSpecificSchema)}
}

func resourceIamBindingCreate(newUpdaterFunc newResourceIamUpdaterFunc, enableBatching bool) func(*schema.ResourceData, interface{}) error {
	return func(d *schema.ResourceData, meta interface{}) error {
		config := meta.(*Config)
		updater, err := newUpdaterFunc(d, config)
		if err != nil {
			return err
		}

		binding := getResourceIamBinding(d)
		modifyF := func(ep *cloudresourcemanager.Policy) error {
			cleaned := removeAllBindingsWithRoleAndCondition(ep.Bindings, binding.Role, binding.Condition)
			ep.Bindings = append(cleaned, binding)
			return nil
		}

		if enableBatching {
			err = BatchRequestModifyIamPolicy(updater, modifyF, config, fmt.Sprintf(
				"Set IAM Binding for role %q on %q", binding.Role, updater.DescribeResource()))
		} else {
			err = iamPolicyReadModifyWrite(updater, modifyF)
		}
		if err != nil {
			return err
		}
		conditionTitle := ""
		if binding.Condition != nil {
			conditionTitle = binding.Condition.Title
		}
		d.SetId(updater.GetResourceId() + " " + binding.Role + " " + conditionTitle)
		return resourceIamBindingRead(newUpdaterFunc)(d, meta)
	}
}

func resourceIamBindingRead(newUpdaterFunc newResourceIamUpdaterFunc) schema.ReadFunc {
	return func(d *schema.ResourceData, meta interface{}) error {
		config := meta.(*Config)
		updater, err := newUpdaterFunc(d, config)
		if err != nil {
			return err
		}

		var role, conditionTitle string
		s := strings.SplitN(d.Id(), " ", 3)
		if len(s) == 2 {
			role = s[1]
		} else if len(s) == 3 {
			role, conditionTitle = s[1], s[2]
		} else {
			return fmt.Errorf("Unexpected resource ID %s", d.Id())
		}
		p, err := iamPolicyReadWithRetry(updater)
		if err != nil {
			return handleNotFoundError(err, d, fmt.Sprintf("Resource %q with IAM Binding (Role %q)", updater.DescribeResource(), role))
		}
		log.Printf("[DEBUG]: Retrieved policy for %s: %+v", updater.DescribeResource(), p)

		log.Printf("Looking for binding with role %s and condition %s", role, conditionTitle)
		var binding *cloudresourcemanager.Binding
		for _, b := range p.Bindings {
			if b.Role == role && ((b.Condition == nil && conditionTitle == "") || (b.Condition != nil && b.Condition.Title == conditionTitle)) {
				c := ""
				if b.Condition != nil {
					c = b.Condition.Title
				}
				log.Printf("Found binding with role %s and condition %s", b.Role, c)
				binding = b
				break
			}
		}
		if binding == nil {
			log.Printf("[DEBUG]: Binding for role %q and condition %q not found in policy for %s, assuming it has no members.", role, conditionTitle, updater.DescribeResource())
			d.Set("role", role)
			d.Set("members", nil)
			return nil
		} else {
			d.Set("role", binding.Role)
			d.Set("members", binding.Members)
		}
		d.Set("etag", p.Etag)
		return nil
	}
}

func resourceIamBindingUpdate(newUpdaterFunc newResourceIamUpdaterFunc, enableBatching bool) func(*schema.ResourceData, interface{}) error {
	return func(d *schema.ResourceData, meta interface{}) error {
		config := meta.(*Config)
		updater, err := newUpdaterFunc(d, config)
		if err != nil {
			return err
		}

		oBinding, nBinding := getResourceIamBindingChange(d)
		modifyF := func(ep *cloudresourcemanager.Policy) error {
			cleaned := removeAllBindingsWithRoleAndCondition(ep.Bindings, oBinding.Role, oBinding.Condition)
			cleaned = removeAllBindingsWithRoleAndCondition(cleaned, nBinding.Role, nBinding.Condition)
			ep.Bindings = append(cleaned, nBinding)
			return nil
		}

		if enableBatching {
			err = BatchRequestModifyIamPolicy(updater, modifyF, config, fmt.Sprintf(
				"Set IAM Binding for role %q on %q", nBinding.Role, updater.DescribeResource()))
		} else {
			err = iamPolicyReadModifyWrite(updater, modifyF)
		}
		if err != nil {
			return err
		}

		conditionTitle := ""
		if nBinding.Condition != nil {
			conditionTitle = nBinding.Condition.Title
		}
		d.SetId(updater.GetResourceId() + " " + nBinding.Role + " " + conditionTitle)
		return resourceIamBindingRead(newUpdaterFunc)(d, meta)
	}
}

func iamBindingImport(resourceIdParser resourceIdParserFunc) schema.StateFunc {
	return func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
		if resourceIdParser == nil {
			return nil, errors.New("Import not supported for this IAM resource.")
		}
		config := m.(*Config)
		s := strings.Fields(d.Id())
		var id, role, conditionTitle string
		if len(s) == 2 {
			id, role = s[0], s[1]
		} else if len(s) >= 3 {
			// condition titles can have any characters in them, so re-join the split string
			id, role, conditionTitle = s[0], s[1], strings.Join(s[2:], " ")
		} else {
			d.SetId("")
			return nil, fmt.Errorf("Wrong number of parts to Binding id %s; expected 'resource_name role [condition_title]'.", s)
		}

		// Set the ID only to the first part so all IAM types can share the same resourceIdParserFunc.
		d.SetId(id)
		d.Set("role", role)
		err := resourceIdParser(d, config)
		if err != nil {
			return nil, err
		}

		// Set the ID again so that the ID matches the ID it would have if it had been created via TF.
		// Use the current ID in case it changed in the resourceIdParserFunc.
		// Use " " as the delimiter because projects can have ':'s and roles can have '/'s.
		d.SetId(d.Id() + " " + role + " " + conditionTitle)
		// It is possible to return multiple bindings, since we can learn about all the bindings
		// for this resource here.  Unfortunately, `terraform import` has some messy behavior here -
		// there's no way to know at this point which resource is being imported, so it's not possible
		// to order this list in a useful way.  In the event of a complex set of bindings, the user
		// will have a terribly confusing set of imported resources and no way to know what matches
		// up to what.  And since the only users who will do a terraform import on their IAM bindings
		// are users who aren't too familiar with Google Cloud IAM (because a "create" for bindings or
		// members is idempotent), it's reasonable to expect that the user will be very alarmed by the
		// plan that terraform will output which mentions destroying a dozen-plus IAM bindings.  With
		// that in mind, we return only the binding that matters.
		return []*schema.ResourceData{d}, nil
	}
}

func resourceIamBindingDelete(newUpdaterFunc newResourceIamUpdaterFunc, enableBatching bool) schema.DeleteFunc {
	return func(d *schema.ResourceData, meta interface{}) error {
		config := meta.(*Config)
		updater, err := newUpdaterFunc(d, config)
		if err != nil {
			return err
		}

		binding := getResourceIamBinding(d)
		modifyF := func(p *cloudresourcemanager.Policy) error {
			p.Bindings = removeAllBindingsWithRoleAndCondition(p.Bindings, binding.Role, binding.Condition)
			return nil
		}

		if enableBatching {
			err = BatchRequestModifyIamPolicy(updater, modifyF, config, fmt.Sprintf(
				"Delete IAM Binding for role %q on %q", binding.Role, updater.DescribeResource()))
		} else {
			err = iamPolicyReadModifyWrite(updater, modifyF)
		}
		if err != nil {
			return handleNotFoundError(err, d, fmt.Sprintf("Resource %q for IAM binding with role %q", updater.DescribeResource(), binding.Role))
		}

		return resourceIamBindingRead(newUpdaterFunc)(d, meta)
	}
}

func getResourceIamBinding(d *schema.ResourceData) *cloudresourcemanager.Binding {
	members := d.Get("members").(*schema.Set).List()
	return &cloudresourcemanager.Binding{
		Members: convertStringArr(members),
		Role:    d.Get("role").(string),
	}
}

func getResourceIamBindingChange(d *schema.ResourceData) (*cloudresourcemanager.Binding, *cloudresourcemanager.Binding) {
	oMembers, nMembers := d.GetChange("members")
	oRole, nRole := d.GetChange("role")

	oBinding := &cloudresourcemanager.Binding{
		Members: convertStringArr(oMembers.(*schema.Set).List()),
		Role:    oRole.(string),
	}
	nBinding := &cloudresourcemanager.Binding{
		Members: convertStringArr(nMembers.(*schema.Set).List()),
		Role:    nRole.(string),
	}
	return oBinding, nBinding
}
