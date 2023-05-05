// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceComputeInstanceGroupNamedPort() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeInstanceGroupNamedPortCreate,
		Read:   resourceComputeInstanceGroupNamedPortRead,
		Delete: resourceComputeInstanceGroupNamedPortDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeInstanceGroupNamedPortImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"group": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description:      `The name of the instance group.`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The name for this named port. The name must be 1-63 characters
long, and comply with RFC1035.`,
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: `The port number, which can be a value between 1 and 65535.`,
			},
			"zone": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The zone of the instance group.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeInstanceGroupNamedPortCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandNestedComputeInstanceGroupNamedPortName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	portProp, err := expandNestedComputeInstanceGroupNamedPortPort(d.Get("port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port"); !isEmptyValue(reflect.ValueOf(portProp)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}

	obj, err = resourceComputeInstanceGroupNamedPortEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := ReplaceVars(d, config, "projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}/setNamedPorts")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new InstanceGroupNamedPort: %#v", obj)

	obj, err = resourceComputeInstanceGroupNamedPortPatchCreateEncoder(d, meta, obj)
	if err != nil {
		return err
	}
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InstanceGroupNamedPort: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating InstanceGroupNamedPort: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}/{{port}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating InstanceGroupNamedPort", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create InstanceGroupNamedPort: %s", err)
	}

	log.Printf("[DEBUG] Finished creating InstanceGroupNamedPort %q: %#v", d.Id(), res)

	return resourceComputeInstanceGroupNamedPortRead(d, meta)
}

func resourceComputeInstanceGroupNamedPortRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InstanceGroupNamedPort: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeInstanceGroupNamedPort %q", d.Id()))
	}

	res, err = flattenNestedComputeInstanceGroupNamedPort(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeInstanceGroupNamedPort because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading InstanceGroupNamedPort: %s", err)
	}

	if err := d.Set("name", flattenNestedComputeInstanceGroupNamedPortName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading InstanceGroupNamedPort: %s", err)
	}
	if err := d.Set("port", flattenNestedComputeInstanceGroupNamedPortPort(res["port"], d, config)); err != nil {
		return fmt.Errorf("Error reading InstanceGroupNamedPort: %s", err)
	}

	return nil
}

func resourceComputeInstanceGroupNamedPortDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InstanceGroupNamedPort: %s", err)
	}
	billingProject = project

	lockName, err := ReplaceVars(d, config, "projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}/setNamedPorts")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	obj, err = resourceComputeInstanceGroupNamedPortPatchDeleteEncoder(d, meta, obj)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "InstanceGroupNamedPort")
	}
	log.Printf("[DEBUG] Deleting InstanceGroupNamedPort %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "InstanceGroupNamedPort")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting InstanceGroupNamedPort", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting InstanceGroupNamedPort %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeInstanceGroupNamedPortImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/instanceGroups/(?P<group>[^/]+)/(?P<port>[^/]+)/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<group>[^/]+)/(?P<port>[^/]+)/(?P<name>[^/]+)",
		"(?P<zone>[^/]+)/(?P<group>[^/]+)/(?P<port>[^/]+)/(?P<name>[^/]+)",
		"(?P<group>[^/]+)/(?P<port>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}/{{port}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNestedComputeInstanceGroupNamedPortName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedComputeInstanceGroupNamedPortPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func expandNestedComputeInstanceGroupNamedPortName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeInstanceGroupNamedPortPort(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceComputeInstanceGroupNamedPortEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	ig, err := ParseInstanceGroupFieldValue(d.Get("group").(string), d, config)
	if err != nil {
		return nil, err
	}

	if err := d.Set("group", ig.Name); err != nil {
		return nil, fmt.Errorf("Error setting group: %s", err)
	}
	if err := d.Set("zone", ig.Zone); err != nil {
		return nil, fmt.Errorf("Error setting zone: %s", err)
	}
	if err := d.Set("project", ig.Project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}

	return obj, nil
}

func flattenNestedComputeInstanceGroupNamedPort(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["namedPorts"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value namedPorts. Actual value: %v", v)
	}

	_, item, err := resourceComputeInstanceGroupNamedPortFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceComputeInstanceGroupNamedPortFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedPort, err := expandNestedComputeInstanceGroupNamedPortPort(d.Get("port"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedPort := flattenNestedComputeInstanceGroupNamedPortPort(expectedPort, d, meta.(*transport_tpg.Config))
	expectedName, err := expandNestedComputeInstanceGroupNamedPortName(d.Get("name"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedName := flattenNestedComputeInstanceGroupNamedPortName(expectedName, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemPort := flattenNestedComputeInstanceGroupNamedPortPort(item["port"], d, meta.(*transport_tpg.Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemPort)) && isEmptyValue(reflect.ValueOf(expectedFlattenedPort))) && !reflect.DeepEqual(itemPort, expectedFlattenedPort) {
			log.Printf("[DEBUG] Skipping item with port= %#v, looking for %#v)", itemPort, expectedFlattenedPort)
			continue
		}
		itemName := flattenNestedComputeInstanceGroupNamedPortName(item["name"], d, meta.(*transport_tpg.Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemName)) && isEmptyValue(reflect.ValueOf(expectedFlattenedName))) && !reflect.DeepEqual(itemName, expectedFlattenedName) {
			log.Printf("[DEBUG] Skipping item with name= %#v, looking for %#v)", itemName, expectedFlattenedName)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}

// PatchCreateEncoder handles creating request data to PATCH parent resource
// with list including new object.
func resourceComputeInstanceGroupNamedPortPatchCreateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceComputeInstanceGroupNamedPortListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	_, found, err := resourceComputeInstanceGroupNamedPortFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}

	// Return error if item already created.
	if found != nil {
		return nil, fmt.Errorf("Unable to create InstanceGroupNamedPort, existing object already found: %+v", found)
	}

	// Return list with the resource to create appended
	res := map[string]interface{}{
		"namedPorts": append(currItems, obj),
	}

	return res, nil
}

// PatchDeleteEncoder handles creating request data to PATCH parent resource
// with list excluding object to delete.
func resourceComputeInstanceGroupNamedPortPatchDeleteEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceComputeInstanceGroupNamedPortListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	idx, item, err := resourceComputeInstanceGroupNamedPortFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}
	if item == nil {
		// Spoof 404 error for proper handling by Delete (i.e. no-op)
		return nil, fake404("nested", "ComputeInstanceGroupNamedPort")
	}

	updatedItems := append(currItems[:idx], currItems[idx+1:]...)
	res := map[string]interface{}{
		"namedPorts": updatedItems,
	}

	return res, nil
}

// ListForPatch handles making API request to get parent resource and
// extracting list of objects.
func resourceComputeInstanceGroupNamedPortListForPatch(d *schema.ResourceData, meta interface{}) ([]interface{}, error) {
	config := meta.(*transport_tpg.Config)
	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}")
	if err != nil {
		return nil, err
	}
	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}

	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return nil, err
	}

	res, err := transport_tpg.SendRequest(config, "GET", project, url, userAgent, nil)
	if err != nil {
		return nil, err
	}

	var v interface{}
	var ok bool

	v, ok = res["namedPorts"]
	if ok && v != nil {
		ls, lsOk := v.([]interface{})
		if !lsOk {
			return nil, fmt.Errorf(`expected list for nested field "namedPorts"`)
		}
		return ls, nil
	}
	return nil, nil
}
