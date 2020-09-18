// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
	"google.golang.org/api/googleapi"
)

func resourceAccessContextManagerServicePerimeterResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessContextManagerServicePerimeterResourceCreate,
		Read:   resourceAccessContextManagerServicePerimeterResourceRead,
		Delete: resourceAccessContextManagerServicePerimeterResourceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessContextManagerServicePerimeterResourceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"perimeter_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The name of the Service Perimeter to add this resource to.`,
			},
			"resource": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `A GCP resource that is inside of the service perimeter.
Currently only projects are allowed.
Format: projects/{project_number}`,
			},
		},
	}
}

func resourceAccessContextManagerServicePerimeterResourceCreate(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	obj := make(map[string]interface{})
	resourceProp, err := expandNestedAccessContextManagerServicePerimeterResourceResource(d.Get("resource"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("resource"); !isEmptyValue(reflect.ValueOf(resourceProp)) && (ok || !reflect.DeepEqual(v, resourceProp)) {
		obj["resource"] = resourceProp
	}

	lockName, err := replaceVars(d, config, "{{perimeter_name}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{perimeter_name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ServicePerimeterResource: %#v", obj)

	obj, err = resourceAccessContextManagerServicePerimeterResourcePatchCreateEncoder(d, meta, obj)
	if err != nil {
		return err
	}
	url, err = addQueryParams(url, map[string]string{"updateMask": "status.resources"})
	if err != nil {
		return err
	}
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ServicePerimeterResource: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{perimeter_name}}/{{resource}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = accessContextManagerOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating ServicePerimeterResource",
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ServicePerimeterResource: %s", err)
	}

	if _, ok := opRes["status"]; ok {
		opRes, err = flattenNestedAccessContextManagerServicePerimeterResource(d, meta, opRes)
		if err != nil {
			return fmt.Errorf("Error getting nested object from operation response: %s", err)
		}
		if opRes == nil {
			// Object isn't there any more - remove it from the state.
			return fmt.Errorf("Error decoding response from operation, could not find nested object")
		}
	}
	if err := d.Set("resource", flattenNestedAccessContextManagerServicePerimeterResourceResource(opRes["resource"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "{{perimeter_name}}/{{resource}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ServicePerimeterResource %q: %#v", d.Id(), res)

	return resourceAccessContextManagerServicePerimeterResourceRead(d, meta)
}

func resourceAccessContextManagerServicePerimeterResourceRead(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{perimeter_name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AccessContextManagerServicePerimeterResource %q", d.Id()))
	}

	res, err = flattenNestedAccessContextManagerServicePerimeterResource(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing AccessContextManagerServicePerimeterResource because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("resource", flattenNestedAccessContextManagerServicePerimeterResourceResource(res["resource"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeterResource: %s", err)
	}

	return nil
}

func resourceAccessContextManagerServicePerimeterResourceDelete(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	billingProject := ""

	lockName, err := replaceVars(d, config, "{{perimeter_name}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{perimeter_name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	obj, err = resourceAccessContextManagerServicePerimeterResourcePatchDeleteEncoder(d, meta, obj)
	if err != nil {
		return handleNotFoundError(err, d, "ServicePerimeterResource")
	}
	url, err = addQueryParams(url, map[string]string{"updateMask": "status.resources"})
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Deleting ServicePerimeterResource %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ServicePerimeterResource")
	}

	err = accessContextManagerOperationWaitTime(
		config, res, "Deleting ServicePerimeterResource",
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ServicePerimeterResource %q: %#v", d.Id(), res)
	return nil
}

func resourceAccessContextManagerServicePerimeterResourceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	parts, err := getImportIdQualifiers([]string{"accessPolicies/(?P<accessPolicy>[^/]+)/servicePerimeters/(?P<perimeter>[^/]+)/(?P<resource>.+)"}, d, config, d.Id())
	if err != nil {
		return nil, err
	}

	d.Set("perimeter_name", fmt.Sprintf("accessPolicies/%s/servicePerimeters/%s", parts["accessPolicy"], parts["perimeter"]))
	d.Set("resource", parts["resource"])
	return []*schema.ResourceData{d}, nil
}

func flattenNestedAccessContextManagerServicePerimeterResourceResource(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandNestedAccessContextManagerServicePerimeterResourceResource(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func flattenNestedAccessContextManagerServicePerimeterResource(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["status"]
	if !ok || v == nil {
		return nil, nil
	}
	res = v.(map[string]interface{})

	v, ok = res["resources"]
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
		return nil, fmt.Errorf("expected list or map for value status.resources. Actual value: %v", v)
	}

	_, item, err := resourceAccessContextManagerServicePerimeterResourceFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceAccessContextManagerServicePerimeterResourceFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedResource, err := expandNestedAccessContextManagerServicePerimeterResourceResource(d.Get("resource"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedResource := flattenNestedAccessContextManagerServicePerimeterResourceResource(expectedResource, d, meta.(*Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		// List response only contains the ID - construct a response object.
		item := map[string]interface{}{
			"resource": itemRaw,
		}

		itemResource := flattenNestedAccessContextManagerServicePerimeterResourceResource(item["resource"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemResource)) && isEmptyValue(reflect.ValueOf(expectedFlattenedResource))) && !reflect.DeepEqual(itemResource, expectedFlattenedResource) {
			log.Printf("[DEBUG] Skipping item with resource= %#v, looking for %#v)", itemResource, expectedFlattenedResource)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}

// PatchCreateEncoder handles creating request data to PATCH parent resource
// with list including new object.
func resourceAccessContextManagerServicePerimeterResourcePatchCreateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceAccessContextManagerServicePerimeterResourceListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	_, found, err := resourceAccessContextManagerServicePerimeterResourceFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}

	// Return error if item already created.
	if found != nil {
		return nil, fmt.Errorf("Unable to create ServicePerimeterResource, existing object already found: %+v", found)
	}

	// Return list with the resource to create appended
	res := map[string]interface{}{
		"resources": append(currItems, obj["resource"]),
	}
	wrapped := map[string]interface{}{
		"status": res,
	}
	res = wrapped

	return res, nil
}

// PatchDeleteEncoder handles creating request data to PATCH parent resource
// with list excluding object to delete.
func resourceAccessContextManagerServicePerimeterResourcePatchDeleteEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceAccessContextManagerServicePerimeterResourceListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	idx, item, err := resourceAccessContextManagerServicePerimeterResourceFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}
	if item == nil {
		// Spoof 404 error for proper handling by Delete (i.e. no-op)
		return nil, &googleapi.Error{
			Code:    404,
			Message: "ServicePerimeterResource not found in list",
		}
	}

	updatedItems := append(currItems[:idx], currItems[idx+1:]...)
	res := map[string]interface{}{
		"resources": updatedItems,
	}
	wrapped := map[string]interface{}{
		"status": res,
	}
	res = wrapped

	return res, nil
}

// ListForPatch handles making API request to get parent resource and
// extracting list of objects.
func resourceAccessContextManagerServicePerimeterResourceListForPatch(d *schema.ResourceData, meta interface{}) ([]interface{}, error) {
	config := meta.(*Config)
	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{perimeter_name}}")
	if err != nil {
		return nil, err
	}
	res, err := sendRequest(config, "GET", "", url, nil)
	if err != nil {
		return nil, err
	}

	var v interface{}
	var ok bool
	if v, ok = res["status"]; ok && v != nil {
		res = v.(map[string]interface{})
	} else {
		return nil, nil
	}

	v, ok = res["resources"]
	if ok && v != nil {
		ls, lsOk := v.([]interface{})
		if !lsOk {
			return nil, fmt.Errorf(`expected list for nested field "resources"`)
		}
		return ls, nil
	}
	return nil, nil
}
