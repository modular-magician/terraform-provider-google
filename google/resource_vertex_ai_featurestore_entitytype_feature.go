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
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceVertexAIFeaturestoreEntitytypeFeature() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAIFeaturestoreEntitytypeFeatureCreate,
		Read:   resourceVertexAIFeaturestoreEntitytypeFeatureRead,
		Update: resourceVertexAIFeaturestoreEntitytypeFeatureUpdate,
		Delete: resourceVertexAIFeaturestoreEntitytypeFeatureDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAIFeaturestoreEntitytypeFeatureImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"entitytype": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.`,
			},
			"value_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Description of the feature.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `A set of key/value label pairs to assign to the feature.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the entity type was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Used to perform consistent read-modify-write updates.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp when the entity type was most recently updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceVertexAIFeaturestoreEntitytypeFeatureCreate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandVertexAIFeaturestoreEntitytypeFeatureLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandVertexAIFeaturestoreEntitytypeFeatureDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	valueTypeProp, err := expandVertexAIFeaturestoreEntitytypeFeatureValueType(d.Get("value_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("value_type"); !isEmptyValue(reflect.ValueOf(valueTypeProp)) && (ok || !reflect.DeepEqual(v, valueTypeProp)) {
		obj["valueType"] = valueTypeProp
	}

	url, err := ReplaceVars(d, config, "{{VertexAIBasePath}}{{entitytype}}/features?featureId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FeaturestoreEntitytypeFeature: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	if v, ok := d.GetOk("entitytype"); ok {
		re := regexp.MustCompile("projects/([a-zA-Z0-9-]*)/(?:locations|regions)/([a-zA-Z0-9-]*)")
		switch {
		case re.MatchString(v.(string)):
			if res := re.FindStringSubmatch(v.(string)); len(res) == 3 && res[1] != "" {
				project = res[1]
			}
		}
	}
	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating FeaturestoreEntitytypeFeature: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{entitytype}}/features/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = VertexAIOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating FeaturestoreEntitytypeFeature", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create FeaturestoreEntitytypeFeature: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = ReplaceVars(d, config, "{{entitytype}}/features/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating FeaturestoreEntitytypeFeature %q: %#v", d.Id(), res)

	return resourceVertexAIFeaturestoreEntitytypeFeatureRead(d, meta)
}

func resourceVertexAIFeaturestoreEntitytypeFeatureRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{VertexAIBasePath}}{{entitytype}}/features/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("VertexAIFeaturestoreEntitytypeFeature %q", d.Id()))
	}

	if err := d.Set("create_time", flattenVertexAIFeaturestoreEntitytypeFeatureCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeaturestoreEntitytypeFeature: %s", err)
	}
	if err := d.Set("update_time", flattenVertexAIFeaturestoreEntitytypeFeatureUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeaturestoreEntitytypeFeature: %s", err)
	}
	if err := d.Set("labels", flattenVertexAIFeaturestoreEntitytypeFeatureLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeaturestoreEntitytypeFeature: %s", err)
	}
	if err := d.Set("description", flattenVertexAIFeaturestoreEntitytypeFeatureDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeaturestoreEntitytypeFeature: %s", err)
	}
	if err := d.Set("value_type", flattenVertexAIFeaturestoreEntitytypeFeatureValueType(res["valueType"], d, config)); err != nil {
		return fmt.Errorf("Error reading FeaturestoreEntitytypeFeature: %s", err)
	}

	return nil
}

func resourceVertexAIFeaturestoreEntitytypeFeatureUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	labelsProp, err := expandVertexAIFeaturestoreEntitytypeFeatureLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandVertexAIFeaturestoreEntitytypeFeatureDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	url, err := ReplaceVars(d, config, "{{VertexAIBasePath}}{{entitytype}}/features/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FeaturestoreEntitytypeFeature %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating FeaturestoreEntitytypeFeature %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating FeaturestoreEntitytypeFeature %q: %#v", d.Id(), res)
	}

	return resourceVertexAIFeaturestoreEntitytypeFeatureRead(d, meta)
}

func resourceVertexAIFeaturestoreEntitytypeFeatureDelete(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := ReplaceVars(d, config, "{{VertexAIBasePath}}{{entitytype}}/features/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	if v, ok := d.GetOk("entitytype"); ok {
		re := regexp.MustCompile("projects/([a-zA-Z0-9-]*)/(?:locations|regions)/([a-zA-Z0-9-]*)")
		switch {
		case re.MatchString(v.(string)):
			if res := re.FindStringSubmatch(v.(string)); len(res) == 3 && res[1] != "" {
				project = res[1]
			}
		}
	}
	log.Printf("[DEBUG] Deleting FeaturestoreEntitytypeFeature %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "FeaturestoreEntitytypeFeature")
	}

	err = VertexAIOperationWaitTime(
		config, res, project, "Deleting FeaturestoreEntitytypeFeature", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting FeaturestoreEntitytypeFeature %q: %#v", d.Id(), res)
	return nil
}

func resourceVertexAIFeaturestoreEntitytypeFeatureImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"(?P<entitytype>.+)/features/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "{{entitytype}}/features/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenVertexAIFeaturestoreEntitytypeFeatureCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeFeatureUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeFeatureLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeFeatureDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIFeaturestoreEntitytypeFeatureValueType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandVertexAIFeaturestoreEntitytypeFeatureLabels(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandVertexAIFeaturestoreEntitytypeFeatureDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIFeaturestoreEntitytypeFeatureValueType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
