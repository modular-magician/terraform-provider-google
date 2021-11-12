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
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComputeAcceleratorType() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeAcceleratorTypeCreate,
		Read:   resourceComputeAcceleratorTypeRead,
		Update: resourceComputeAcceleratorTypeUpdate,
		Delete: resourceComputeAcceleratorTypeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeAcceleratorTypeImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Name of the resource.`,
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The name of the zone where the accelerator type resides.`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"deprecated": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The deprecation status associated with this accelerator type.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deprecated": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `An optional RFC3339 timestamp on or after which the state of this
resource is intended to change to DEPRECATED. This is only
informational and the status will not change unless the client
explicitly changes it.`,
						},
						"obsolete": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `An optional RFC3339 timestamp on or after which the state of this
resource is intended to change to OBSOLETE. This is only
informational and the status will not change unless the client
explicitly changes it.`,
						},
						"replacement": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The URL of the suggested replacement for a deprecated resource.
The suggested replacement resource must be the same kind of
resource as the deprecated resource.`,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `An optional RFC3339 timestamp on or after which the state of this
resource is intended to change to DELETED. This is only
informational and the status will not change unless the client
explicitly changes it.`,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The deprecation state of this resource. This can be DEPRECATED,
OBSOLETE, or DELETED. Operations which create a new resource
using a DEPRECATED resource will return successfully, but with a
warning indicating the deprecated resource and recommending its
replacement. Operations which use OBSOLETE or DELETED resources
will be rejected and result in an error.`,
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `An optional textual description of the resource.`,
			},
			"id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The unique identifier for the resource.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeAcceleratorTypeCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeAcceleratorTypeName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	zoneProp, err := expandComputeAcceleratorTypeZone(d.Get("zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zone"); !isEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/acceleratorTypes")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AcceleratorType: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AcceleratorType: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating AcceleratorType: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/zones/{{zone}}/acceleratorTypes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating AcceleratorType %q: %#v", d.Id(), res)

	return resourceComputeAcceleratorTypeRead(d, meta)
}

func resourceComputeAcceleratorTypeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/acceleratorTypes/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AcceleratorType: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeAcceleratorType %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading AcceleratorType: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeAcceleratorTypeCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading AcceleratorType: %s", err)
	}
	if err := d.Set("deprecated", flattenComputeAcceleratorTypeDeprecated(res["deprecated"], d, config)); err != nil {
		return fmt.Errorf("Error reading AcceleratorType: %s", err)
	}
	if err := d.Set("description", flattenComputeAcceleratorTypeDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading AcceleratorType: %s", err)
	}
	if err := d.Set("id", flattenComputeAcceleratorTypeId(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading AcceleratorType: %s", err)
	}
	if err := d.Set("name", flattenComputeAcceleratorTypeName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading AcceleratorType: %s", err)
	}
	if err := d.Set("zone", flattenComputeAcceleratorTypeZone(res["zone"], d, config)); err != nil {
		return fmt.Errorf("Error reading AcceleratorType: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading AcceleratorType: %s", err)
	}

	return nil
}

func resourceComputeAcceleratorTypeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AcceleratorType: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	nameProp, err := expandComputeAcceleratorTypeName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	zoneProp, err := expandComputeAcceleratorTypeZone(d.Get("zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zone"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/acceleratorTypes/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AcceleratorType %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PUT", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating AcceleratorType %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating AcceleratorType %q: %#v", d.Id(), res)
	}

	return resourceComputeAcceleratorTypeRead(d, meta)
}

func resourceComputeAcceleratorTypeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for AcceleratorType: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/acceleratorTypes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting AcceleratorType %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "AcceleratorType")
	}

	log.Printf("[DEBUG] Finished deleting AcceleratorType %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeAcceleratorTypeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/acceleratorTypes/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/zones/{{zone}}/acceleratorTypes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeAcceleratorTypeCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeAcceleratorTypeDeprecated(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["state"] =
		flattenComputeAcceleratorTypeDeprecatedState(original["state"], d, config)
	transformed["deprecated"] =
		flattenComputeAcceleratorTypeDeprecatedDeprecated(original["deprecated"], d, config)
	transformed["obsolete"] =
		flattenComputeAcceleratorTypeDeprecatedObsolete(original["obsolete"], d, config)
	transformed["replacement"] =
		flattenComputeAcceleratorTypeDeprecatedReplacement(original["replacement"], d, config)
	transformed["state"] =
		flattenComputeAcceleratorTypeDeprecatedState(original["state"], d, config)
	return []interface{}{transformed}
}
func flattenComputeAcceleratorTypeDeprecatedState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeAcceleratorTypeDeprecatedDeprecated(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeAcceleratorTypeDeprecatedObsolete(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeAcceleratorTypeDeprecatedReplacement(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeAcceleratorTypeDeprecatedState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeAcceleratorTypeDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeAcceleratorTypeId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
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

func flattenComputeAcceleratorTypeName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeAcceleratorTypeZone(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandComputeAcceleratorTypeName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAcceleratorTypeZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
