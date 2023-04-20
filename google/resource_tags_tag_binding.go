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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceTagsTagBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceTagsTagBindingCreate,
		Read:   resourceTagsTagBindingRead,
		Delete: resourceTagsTagBindingDelete,

		Importer: &schema.ResourceImporter{
			State: resourceTagsTagBindingImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"parent": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The full resource name of the resource the TagValue is bound to. E.g. //cloudresourcemanager.googleapis.com/projects/123`,
			},
			"tag_value": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The TagValue of the TagBinding. Must be of the form tagValues/456.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The generated id for the TagBinding. This is a string of the form: 'tagBindings/{full-resource-name}/{tag-value-name}'`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceTagsTagBindingCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	parentProp, err := expandNestedTagsTagBindingParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	tagValueProp, err := expandNestedTagsTagBindingTagValue(d.Get("tag_value"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tag_value"); !isEmptyValue(reflect.ValueOf(tagValueProp)) && (ok || !reflect.DeepEqual(v, tagValueProp)) {
		obj["tagValue"] = tagValueProp
	}

	lockName, err := ReplaceVars(d, config, "tagBindings/{{parent}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := ReplaceVars(d, config, "{{TagsBasePath}}tagBindings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TagBinding: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating TagBinding: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "tagBindings/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = TagsOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating TagBinding", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create TagBinding: %s", err)
	}

	if _, ok := opRes["tagBindings"]; ok {
		opRes, err = flattenNestedTagsTagBinding(d, meta, opRes)
		if err != nil {
			return fmt.Errorf("Error getting nested object from operation response: %s", err)
		}
		if opRes == nil {
			// Object isn't there any more - remove it from the state.
			return fmt.Errorf("Error decoding response from operation, could not find nested object")
		}
	}
	if err := d.Set("name", flattenNestedTagsTagBindingName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "tagBindings/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating TagBinding %q: %#v", d.Id(), res)

	return resourceTagsTagBindingRead(d, meta)
}

func resourceTagsTagBindingRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{TagsBasePath}}tagBindings/?parent={{parent}}&pageSize=300")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("TagsTagBinding %q", d.Id()))
	}

	res, err = flattenNestedTagsTagBinding(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing TagsTagBinding because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("name", flattenNestedTagsTagBindingName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagBinding: %s", err)
	}
	if err := d.Set("parent", flattenNestedTagsTagBindingParent(res["parent"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagBinding: %s", err)
	}
	if err := d.Set("tag_value", flattenNestedTagsTagBindingTagValue(res["tagValue"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagBinding: %s", err)
	}

	return nil
}

func resourceTagsTagBindingDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	lockName, err := ReplaceVars(d, config, "tagBindings/{{parent}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := ReplaceVars(d, config, "{{TagsBasePath}}tagBindings/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting TagBinding %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "TagBinding")
	}

	err = TagsOperationWaitTime(
		config, res, "Deleting TagBinding", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting TagBinding %q: %#v", d.Id(), res)
	return nil
}

func resourceTagsTagBindingImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := ParseImportId([]string{
		"tagBindings/(?P<name>.+)",
		"(?P<name>.+)",
	}, d, config); err != nil {
		return nil, err
	}

	name := d.Get("name").(string)
	d.SetId(name)

	return []*schema.ResourceData{d}, nil
}

func flattenNestedTagsTagBindingName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	parts := strings.Split(v.(string), "/")
	return strings.Join(parts[len(parts)-3:], "/")
}

func flattenNestedTagsTagBindingParent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNestedTagsTagBindingTagValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNestedTagsTagBindingParent(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNestedTagsTagBindingTagValue(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func flattenNestedTagsTagBinding(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["tagBindings"]
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
		return nil, fmt.Errorf("expected list or map for value tagBindings. Actual value: %v", v)
	}

	_, item, err := resourceTagsTagBindingFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceTagsTagBindingFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedName := d.Get("name")
	expectedFlattenedName := flattenNestedTagsTagBindingName(expectedName, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemName := flattenNestedTagsTagBindingName(item["name"], d, meta.(*transport_tpg.Config))
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
