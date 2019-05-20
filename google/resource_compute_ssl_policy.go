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

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"google.golang.org/api/compute/v1"
)

func sslPolicyCustomizeDiff(diff *schema.ResourceDiff, v interface{}) error {
	profile := diff.Get("profile")
	customFeaturesCount := diff.Get("custom_features.#")

	// Validate that policy configs aren't incompatible during all phases
	// CUSTOM profile demands non-zero custom_features, and other profiles (i.e., not CUSTOM) demand zero custom_features
	if diff.HasChange("profile") || diff.HasChange("custom_features") {
		if profile.(string) == "CUSTOM" {
			if customFeaturesCount.(int) == 0 {
				return fmt.Errorf("Error in SSL Policy %s: the profile is set to %s but no custom_features are set.", diff.Get("name"), profile.(string))
			}
		} else {
			if customFeaturesCount != 0 {
				return fmt.Errorf("Error in SSL Policy %s: the profile is set to %s but using custom_features requires the profile to be CUSTOM.", diff.Get("name"), profile.(string))
			}
		}
		return nil
	}
	return nil
}

func resourceComputeSslPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeSslPolicyCreate,
		Read:   resourceComputeSslPolicyRead,
		Update: resourceComputeSslPolicyUpdate,
		Delete: resourceComputeSslPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeSslPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		CustomizeDiff: sslPolicyCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"custom_features": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"min_tls_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"TLS_1_0", "TLS_1_1", "TLS_1_2", ""}, false),
				Default:      "TLS_1_0",
			},
			"profile": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM", ""}, false),
				Default:      "COMPATIBLE",
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled_features": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
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
	}
}

func resourceComputeSslPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeSslPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeSslPolicyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	profileProp, err := expandComputeSslPolicyProfile(d.Get("profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("profile"); !isEmptyValue(reflect.ValueOf(profileProp)) && (ok || !reflect.DeepEqual(v, profileProp)) {
		obj["profile"] = profileProp
	}
	minTlsVersionProp, err := expandComputeSslPolicyMinTlsVersion(d.Get("min_tls_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("min_tls_version"); !isEmptyValue(reflect.ValueOf(minTlsVersionProp)) && (ok || !reflect.DeepEqual(v, minTlsVersionProp)) {
		obj["minTlsVersion"] = minTlsVersionProp
	}
	customFeaturesProp, err := expandComputeSslPolicyCustomFeatures(d.Get("custom_features"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("custom_features"); !isEmptyValue(reflect.ValueOf(customFeaturesProp)) && (ok || !reflect.DeepEqual(v, customFeaturesProp)) {
		obj["customFeatures"] = customFeaturesProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/sslPolicies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SslPolicy: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating SslPolicy: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating SslPolicy",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create SslPolicy: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating SslPolicy %q: %#v", d.Id(), res)

	return resourceComputeSslPolicyRead(d, meta)
}

func resourceComputeSslPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/sslPolicies/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeSslPolicy %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}

	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("creation_timestamp", flattenComputeSslPolicyCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("description", flattenComputeSslPolicyDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("name", flattenComputeSslPolicyName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("profile", flattenComputeSslPolicyProfile(res["profile"], d)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("min_tls_version", flattenComputeSslPolicyMinTlsVersion(res["minTlsVersion"], d)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("enabled_features", flattenComputeSslPolicyEnabledFeatures(res["enabledFeatures"], d)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("custom_features", flattenComputeSslPolicyCustomFeatures(res["customFeatures"], d)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if err := d.Set("fingerprint", flattenComputeSslPolicyFingerprint(res["fingerprint"], d)); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}

	return nil
}

func resourceComputeSslPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	profileProp, err := expandComputeSslPolicyProfile(d.Get("profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("profile"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, profileProp)) {
		obj["profile"] = profileProp
	}
	minTlsVersionProp, err := expandComputeSslPolicyMinTlsVersion(d.Get("min_tls_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("min_tls_version"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, minTlsVersionProp)) {
		obj["minTlsVersion"] = minTlsVersionProp
	}
	customFeaturesProp, err := expandComputeSslPolicyCustomFeatures(d.Get("custom_features"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("custom_features"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, customFeaturesProp)) {
		obj["customFeatures"] = customFeaturesProp
	}

	obj, err = resourceComputeSslPolicyUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/sslPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating SslPolicy %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PATCH", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating SslPolicy %q: %s", d.Id(), err)
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Updating SslPolicy",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceComputeSslPolicyRead(d, meta)
}

func resourceComputeSslPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/sslPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting SslPolicy %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "SslPolicy")
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting SslPolicy",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting SslPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeSslPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{"projects/(?P<project>[^/]+)/global/sslPolicies/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeSslPolicyCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeSslPolicyDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeSslPolicyName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeSslPolicyProfile(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeSslPolicyMinTlsVersion(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeSslPolicyEnabledFeatures(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeSslPolicyCustomFeatures(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeSslPolicyFingerprint(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandComputeSslPolicyDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslPolicyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslPolicyProfile(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslPolicyMinTlsVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslPolicyCustomFeatures(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func resourceComputeSslPolicyUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// TODO(https://github.com/GoogleCloudPlatform/magic-modules/issues/184): Handle fingerprint consistently
	obj["fingerprint"] = d.Get("fingerprint")

	// TODO(https://github.com/GoogleCloudPlatform/magic-modules/issues/183): Can we generalize this
	// Send a null fields if customFeatures is empty.
	if v, ok := obj["customFeatures"]; ok && len(v.([]interface{})) == 0 {
		obj["customFeatures"] = nil
	}

	return obj, nil
}
