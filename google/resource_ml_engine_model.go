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

func ResourceMLEngineModel() *schema.Resource {
	return &schema.Resource{
		Create: resourceMLEngineModelCreate,
		Read:   resourceMLEngineModelRead,
		Delete: resourceMLEngineModelDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMLEngineModelImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name specified for the model.`,
			},
			"default_version": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `The default version of the model. This version will be used to handle
prediction requests that do not specify a version.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `The name specified for the version when it was created.`,
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The description specified for the model when it was created.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: `One or more labels that you can add, to organize your models.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"online_prediction_console_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging`,
			},
			"online_prediction_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `If true, online prediction access logs are sent to StackDriver Logging.`,
			},
			"regions": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `The list of regions where the model is going to be deployed.
Currently only one region per model is supported`,
				MaxItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func resourceMLEngineModelCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandMLEngineModelName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandMLEngineModelDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	defaultVersionProp, err := expandMLEngineModelDefaultVersion(d.Get("default_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultVersionProp)) && (ok || !reflect.DeepEqual(v, defaultVersionProp)) {
		obj["defaultVersion"] = defaultVersionProp
	}
	regionsProp, err := expandMLEngineModelRegions(d.Get("regions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("regions"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionsProp)) && (ok || !reflect.DeepEqual(v, regionsProp)) {
		obj["regions"] = regionsProp
	}
	onlinePredictionLoggingProp, err := expandMLEngineModelOnlinePredictionLogging(d.Get("online_prediction_logging"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("online_prediction_logging"); !tpgresource.IsEmptyValue(reflect.ValueOf(onlinePredictionLoggingProp)) && (ok || !reflect.DeepEqual(v, onlinePredictionLoggingProp)) {
		obj["onlinePredictionLogging"] = onlinePredictionLoggingProp
	}
	onlinePredictionConsoleLoggingProp, err := expandMLEngineModelOnlinePredictionConsoleLogging(d.Get("online_prediction_console_logging"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("online_prediction_console_logging"); !tpgresource.IsEmptyValue(reflect.ValueOf(onlinePredictionConsoleLoggingProp)) && (ok || !reflect.DeepEqual(v, onlinePredictionConsoleLoggingProp)) {
		obj["onlinePredictionConsoleLogging"] = onlinePredictionConsoleLoggingProp
	}
	labelsProp, err := expandMLEngineModelLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{MLEngineBasePath}}projects/{{project}}/models")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Model: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Model: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating Model: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/models/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Model %q: %#v", d.Id(), res)

	return resourceMLEngineModelRead(d, meta)
}

func resourceMLEngineModelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{MLEngineBasePath}}projects/{{project}}/models/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Model: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("MLEngineModel %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}

	if err := d.Set("name", flattenMLEngineModelName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}
	if err := d.Set("description", flattenMLEngineModelDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}
	if err := d.Set("default_version", flattenMLEngineModelDefaultVersion(res["defaultVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}
	if err := d.Set("regions", flattenMLEngineModelRegions(res["regions"], d, config)); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}
	if err := d.Set("online_prediction_logging", flattenMLEngineModelOnlinePredictionLogging(res["onlinePredictionLogging"], d, config)); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}
	if err := d.Set("online_prediction_console_logging", flattenMLEngineModelOnlinePredictionConsoleLogging(res["onlinePredictionConsoleLogging"], d, config)); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}
	if err := d.Set("labels", flattenMLEngineModelLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Model: %s", err)
	}

	return nil
}

func resourceMLEngineModelDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Model: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{MLEngineBasePath}}projects/{{project}}/models/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Model %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Model")
	}

	err = MLEngineOperationWaitTime(
		config, res, project, "Deleting Model", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Model %q: %#v", d.Id(), res)
	return nil
}

func resourceMLEngineModelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/models/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/models/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenMLEngineModelName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenMLEngineModelDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMLEngineModelDefaultVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["name"] =
		flattenMLEngineModelDefaultVersionName(original["name"], d, config)
	return []interface{}{transformed}
}
func flattenMLEngineModelDefaultVersionName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMLEngineModelRegions(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMLEngineModelOnlinePredictionLogging(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMLEngineModelOnlinePredictionConsoleLogging(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMLEngineModelLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandMLEngineModelName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelDefaultVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandMLEngineModelDefaultVersionName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandMLEngineModelDefaultVersionName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelRegions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelOnlinePredictionLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelOnlinePredictionConsoleLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMLEngineModelLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
