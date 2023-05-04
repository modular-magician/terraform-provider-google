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
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceDialogflowFulfillment() *schema.Resource {
	return &schema.Resource{
		Create: resourceDialogflowFulfillmentCreate,
		Read:   resourceDialogflowFulfillmentRead,
		Update: resourceDialogflowFulfillmentUpdate,
		Delete: resourceDialogflowFulfillmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDialogflowFulfillmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The human-readable name of the fulfillment, unique within the agent.`,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Whether fulfillment is enabled.`,
			},
			"features": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The field defines whether the fulfillment is enabled for certain features.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: verify.ValidateEnum([]string{"SMALLTALK"}),
							Description: `The type of the feature that enabled for fulfillment.
* SMALLTALK: Fulfillment is enabled for SmallTalk. Possible values: ["SMALLTALK"]`,
						},
					},
				},
			},
			"generic_web_service": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uri": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The fulfillment URI for receiving POST requests. It must use https protocol.`,
						},
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The password for HTTP Basic authentication.`,
						},
						"request_headers": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: `The HTTP request headers to send together with fulfillment requests.`,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The user name for HTTP Basic authentication.`,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The unique identifier of the fulfillment.
Format: projects/<Project ID>/agent/fulfillment - projects/<Project ID>/locations/<Location ID>/agent/fulfillment`,
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

func resourceDialogflowFulfillmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowFulfillmentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enabledProp, err := expandDialogflowFulfillmentEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !isEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}
	featuresProp, err := expandDialogflowFulfillmentFeatures(d.Get("features"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("features"); !isEmptyValue(reflect.ValueOf(featuresProp)) && (ok || !reflect.DeepEqual(v, featuresProp)) {
		obj["features"] = featuresProp
	}
	genericWebServiceProp, err := expandDialogflowFulfillmentGenericWebService(d.Get("generic_web_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("generic_web_service"); !isEmptyValue(reflect.ValueOf(genericWebServiceProp)) && (ok || !reflect.DeepEqual(v, genericWebServiceProp)) {
		obj["genericWebService"] = genericWebServiceProp
	}

	url, err := ReplaceVars(d, config, "{{DialogflowBasePath}}projects/{{project}}/agent/fulfillment/?updateMask=name,displayName,enabled,genericWebService,features")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Fulfillment: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Fulfillment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Fulfillment: %s", err)
	}
	if err := d.Set("name", flattenDialogflowFulfillmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		respBody, ok := res["response"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}

		name, ok = respBody.(map[string]interface{})["name"]
		if !ok {
			return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
		}
	}
	if err := d.Set("name", name.(string)); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name.(string))

	log.Printf("[DEBUG] Finished creating Fulfillment %q: %#v", d.Id(), res)

	return resourceDialogflowFulfillmentRead(d, meta)
}

func resourceDialogflowFulfillmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{DialogflowBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Fulfillment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DialogflowFulfillment %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Fulfillment: %s", err)
	}

	if err := d.Set("name", flattenDialogflowFulfillmentName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Fulfillment: %s", err)
	}
	if err := d.Set("display_name", flattenDialogflowFulfillmentDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Fulfillment: %s", err)
	}
	if err := d.Set("enabled", flattenDialogflowFulfillmentEnabled(res["enabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading Fulfillment: %s", err)
	}
	if err := d.Set("features", flattenDialogflowFulfillmentFeatures(res["features"], d, config)); err != nil {
		return fmt.Errorf("Error reading Fulfillment: %s", err)
	}
	if err := d.Set("generic_web_service", flattenDialogflowFulfillmentGenericWebService(res["genericWebService"], d, config)); err != nil {
		return fmt.Errorf("Error reading Fulfillment: %s", err)
	}

	return nil
}

func resourceDialogflowFulfillmentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Fulfillment: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowFulfillmentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enabledProp, err := expandDialogflowFulfillmentEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}
	featuresProp, err := expandDialogflowFulfillmentFeatures(d.Get("features"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("features"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, featuresProp)) {
		obj["features"] = featuresProp
	}
	genericWebServiceProp, err := expandDialogflowFulfillmentGenericWebService(d.Get("generic_web_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("generic_web_service"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, genericWebServiceProp)) {
		obj["genericWebService"] = genericWebServiceProp
	}

	url, err := ReplaceVars(d, config, "{{DialogflowBasePath}}projects/{{project}}/agent/fulfillment/")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Fulfillment %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("enabled") {
		updateMask = append(updateMask, "enabled")
	}

	if d.HasChange("features") {
		updateMask = append(updateMask, "features")
	}

	if d.HasChange("generic_web_service") {
		updateMask = append(updateMask, "genericWebService")
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
		return fmt.Errorf("Error updating Fulfillment %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Fulfillment %q: %#v", d.Id(), res)
	}

	return resourceDialogflowFulfillmentRead(d, meta)
}

func resourceDialogflowFulfillmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Fulfillment: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{DialogflowBasePath}}projects/{{project}}/agent/fulfillment/?updateMask=name,displayName,enabled,genericWebService,features")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Fulfillment %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Fulfillment")
	}

	log.Printf("[DEBUG] Finished deleting Fulfillment %q: %#v", d.Id(), res)
	return nil
}

func resourceDialogflowFulfillmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	stringParts := strings.Split(d.Get("name").(string), "/")
	if len(stringParts) < 2 {
		return nil, fmt.Errorf(
			"Could not split project from name: %s",
			d.Get("name"),
		)
	}

	if err := d.Set("project", stringParts[1]); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

func flattenDialogflowFulfillmentName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowFulfillmentDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowFulfillmentEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowFulfillmentFeatures(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"type": flattenDialogflowFulfillmentFeaturesType(original["type"], d, config),
		})
	}
	return transformed
}
func flattenDialogflowFulfillmentFeaturesType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowFulfillmentGenericWebService(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["uri"] =
		flattenDialogflowFulfillmentGenericWebServiceUri(original["uri"], d, config)
	transformed["username"] =
		flattenDialogflowFulfillmentGenericWebServiceUsername(original["username"], d, config)
	transformed["password"] =
		flattenDialogflowFulfillmentGenericWebServicePassword(original["password"], d, config)
	transformed["request_headers"] =
		flattenDialogflowFulfillmentGenericWebServiceRequestHeaders(original["requestHeaders"], d, config)
	return []interface{}{transformed}
}
func flattenDialogflowFulfillmentGenericWebServiceUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowFulfillmentGenericWebServiceUsername(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowFulfillmentGenericWebServicePassword(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowFulfillmentGenericWebServiceRequestHeaders(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDialogflowFulfillmentDisplayName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowFulfillmentEnabled(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowFulfillmentFeatures(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedType, err := expandDialogflowFulfillmentFeaturesType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !isEmptyValue(val) {
			transformed["type"] = transformedType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDialogflowFulfillmentFeaturesType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowFulfillmentGenericWebService(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUri, err := expandDialogflowFulfillmentGenericWebServiceUri(original["uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !isEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	transformedUsername, err := expandDialogflowFulfillmentGenericWebServiceUsername(original["username"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !isEmptyValue(val) {
		transformed["username"] = transformedUsername
	}

	transformedPassword, err := expandDialogflowFulfillmentGenericWebServicePassword(original["password"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPassword); val.IsValid() && !isEmptyValue(val) {
		transformed["password"] = transformedPassword
	}

	transformedRequestHeaders, err := expandDialogflowFulfillmentGenericWebServiceRequestHeaders(original["request_headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestHeaders); val.IsValid() && !isEmptyValue(val) {
		transformed["requestHeaders"] = transformedRequestHeaders
	}

	return transformed, nil
}

func expandDialogflowFulfillmentGenericWebServiceUri(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowFulfillmentGenericWebServiceUsername(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowFulfillmentGenericWebServicePassword(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowFulfillmentGenericWebServiceRequestHeaders(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
