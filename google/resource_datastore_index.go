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
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceDatastoreIndex() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatastoreIndexCreate,
		Read:   resourceDatastoreIndexRead,
		Delete: resourceDatastoreIndexDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDatastoreIndexImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"kind": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The entity kind which the index applies to.`,
			},
			"ancestor": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"NONE", "ALL_ANCESTORS", ""}),
				Description:  `Policy for including ancestors in the index. Default value: "NONE" Possible values: ["NONE", "ALL_ANCESTORS"]`,
				Default:      "NONE",
			},
			"properties": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `An ordered list of properties to index on.`,
				MinItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"direction": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: verify.ValidateEnum([]string{"ASCENDING", "DESCENDING"}),
							Description:  `The direction the index should optimize for sorting. Possible values: ["ASCENDING", "DESCENDING"]`,
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `The property name to index.`,
						},
					},
				},
			},
			"index_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The index id.`,
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

func resourceDatastoreIndexCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	kindProp, err := expandDatastoreIndexKind(d.Get("kind"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kind"); !tpgresource.IsEmptyValue(reflect.ValueOf(kindProp)) && (ok || !reflect.DeepEqual(v, kindProp)) {
		obj["kind"] = kindProp
	}
	ancestorProp, err := expandDatastoreIndexAncestor(d.Get("ancestor"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ancestor"); !tpgresource.IsEmptyValue(reflect.ValueOf(ancestorProp)) && (ok || !reflect.DeepEqual(v, ancestorProp)) {
		obj["ancestor"] = ancestorProp
	}
	propertiesProp, err := expandDatastoreIndexProperties(d.Get("properties"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("properties"); !tpgresource.IsEmptyValue(reflect.ValueOf(propertiesProp)) && (ok || !reflect.DeepEqual(v, propertiesProp)) {
		obj["properties"] = propertiesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DatastoreBasePath}}projects/{{project}}/indexes")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Index: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Index: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate), transport_tpg.DatastoreIndex409Contention)
	if err != nil {
		return fmt.Errorf("Error creating Index: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/indexes/{{index_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = DatastoreOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Index", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Index: %s", err)
	}

	if err := d.Set("index_id", flattenDatastoreIndexIndexId(opRes["indexId"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/indexes/{{index_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Index %q: %#v", d.Id(), res)

	return resourceDatastoreIndexRead(d, meta)
}

func resourceDatastoreIndexRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DatastoreBasePath}}projects/{{project}}/indexes/{{index_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Index: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil, transport_tpg.DatastoreIndex409Contention)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DatastoreIndex %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}

	if err := d.Set("index_id", flattenDatastoreIndexIndexId(res["indexId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("kind", flattenDatastoreIndexKind(res["kind"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("ancestor", flattenDatastoreIndexAncestor(res["ancestor"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}
	if err := d.Set("properties", flattenDatastoreIndexProperties(res["properties"], d, config)); err != nil {
		return fmt.Errorf("Error reading Index: %s", err)
	}

	return nil
}

func resourceDatastoreIndexDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Index: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DatastoreBasePath}}projects/{{project}}/indexes/{{index_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Index %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete), transport_tpg.DatastoreIndex409Contention)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Index")
	}

	err = DatastoreOperationWaitTime(
		config, res, project, "Deleting Index", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Index %q: %#v", d.Id(), res)
	return nil
}

func resourceDatastoreIndexImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/indexes/(?P<index_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<index_id>[^/]+)",
		"(?P<index_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/indexes/{{index_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDatastoreIndexIndexId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatastoreIndexKind(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatastoreIndexAncestor(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatastoreIndexProperties(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"name":      flattenDatastoreIndexPropertiesName(original["name"], d, config),
			"direction": flattenDatastoreIndexPropertiesDirection(original["direction"], d, config),
		})
	}
	return transformed
}
func flattenDatastoreIndexPropertiesName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDatastoreIndexPropertiesDirection(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDatastoreIndexKind(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatastoreIndexAncestor(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatastoreIndexProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDatastoreIndexPropertiesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedDirection, err := expandDatastoreIndexPropertiesDirection(original["direction"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDirection); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["direction"] = transformedDirection
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDatastoreIndexPropertiesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatastoreIndexPropertiesDirection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
