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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceTagsTagValue() *schema.Resource {
	return &schema.Resource{
		Create: resourceTagsTagValueCreate,
		Read:   resourceTagsTagValueRead,
		Update: resourceTagsTagValueUpdate,
		Delete: resourceTagsTagValueDelete,

		Importer: &schema.ResourceImporter{
			State: resourceTagsTagValueImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"parent": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Input only. The resource name of the new TagValue's parent. Must be of the form tagKeys/{tag_key_id}.`,
			},
			"short_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 63),
				Description: `Input only. User-assigned short name for TagValue. The short name should be unique for TagValues within the same parent TagKey.

The short name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.`,
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 256),
				Description:  `User-assigned description of the TagValue. Must not exceed 256 characters.`,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. Creation time.

A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The generated numeric id for the TagValue.`,
			},
			"namespaced_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Namespaced name of the TagValue. Will be in the format {organizationId}/{tag_key_short_name}/{shortName}.`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. Update time.

A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceTagsTagValueCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	parentProp, err := expandTagsTagValueParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	shortNameProp, err := expandTagsTagValueShortName(d.Get("short_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("short_name"); !isEmptyValue(reflect.ValueOf(shortNameProp)) && (ok || !reflect.DeepEqual(v, shortNameProp)) {
		obj["shortName"] = shortNameProp
	}
	descriptionProp, err := expandTagsTagValueDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	lockName, err := replaceVars(d, config, "tagValues/{{parent}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{TagsBasePath}}tagValues")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TagValue: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating TagValue: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "tagValues/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = tagsOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating TagValue", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create TagValue: %s", err)
	}

	if err := d.Set("name", flattenTagsTagValueName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "tagValues/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating TagValue %q: %#v", d.Id(), res)

	return resourceTagsTagValueRead(d, meta)
}

func resourceTagsTagValueRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{TagsBasePath}}tagValues/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("TagsTagValue %q", d.Id()))
	}

	if err := d.Set("name", flattenTagsTagValueName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagValue: %s", err)
	}
	if err := d.Set("parent", flattenTagsTagValueParent(res["parent"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagValue: %s", err)
	}
	if err := d.Set("short_name", flattenTagsTagValueShortName(res["shortName"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagValue: %s", err)
	}
	if err := d.Set("namespaced_name", flattenTagsTagValueNamespacedName(res["namespacedName"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagValue: %s", err)
	}
	if err := d.Set("description", flattenTagsTagValueDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagValue: %s", err)
	}
	if err := d.Set("create_time", flattenTagsTagValueCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagValue: %s", err)
	}
	if err := d.Set("update_time", flattenTagsTagValueUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagValue: %s", err)
	}

	return nil
}

func resourceTagsTagValueUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	descriptionProp, err := expandTagsTagValueDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	lockName, err := replaceVars(d, config, "tagValues/{{parent}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{TagsBasePath}}tagValues/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating TagValue %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating TagValue %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating TagValue %q: %#v", d.Id(), res)
	}

	err = tagsOperationWaitTime(
		config, res, "Updating TagValue", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceTagsTagValueRead(d, meta)
}

func resourceTagsTagValueDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	lockName, err := replaceVars(d, config, "tagValues/{{parent}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{TagsBasePath}}tagValues/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting TagValue %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "TagValue")
	}

	err = tagsOperationWaitTime(
		config, res, "Deleting TagValue", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting TagValue %q: %#v", d.Id(), res)
	return nil
}

func resourceTagsTagValueImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"tagValues/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "tagValues/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenTagsTagValueName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenTagsTagValueParent(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenTagsTagValueShortName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenTagsTagValueNamespacedName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenTagsTagValueDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenTagsTagValueCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenTagsTagValueUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandTagsTagValueParent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTagsTagValueShortName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandTagsTagValueDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
