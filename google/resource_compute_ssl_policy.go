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
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	compute "google.golang.org/api/compute/v1"
)

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
		CustomizeDiff: func(diff *schema.ResourceDiff, v interface{}) error {
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
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"profile": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM", ""}, false),
				Default:      "COMPATIBLE",
			},
			"min_tls_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"TLS_1_0", "TLS_1_1", "TLS_1_2", ""}, false),
				Default:      "TLS_1_0",
			},
			"custom_features": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
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

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	descriptionProp, err := expandComputeSslPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	}
	nameProp, err := expandComputeSslPolicyName(d.Get("name"), d, config)
	if err != nil {
		return err
	}
	profileProp, err := expandComputeSslPolicyProfile(d.Get("profile"), d, config)
	if err != nil {
		return err
	}
	minTlsVersionProp, err := expandComputeSslPolicyMinTlsVersion(d.Get("min_tls_version"), d, config)
	if err != nil {
		return err
	}
	customFeaturesProp, err := expandComputeSslPolicyCustomFeatures(d.Get("custom_features"), d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"description":    descriptionProp,
		"name":           nameProp,
		"profile":        profileProp,
		"minTlsVersion":  minTlsVersionProp,
		"customFeatures": customFeaturesProp,
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/sslPolicies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SslPolicy: %#v", obj)
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating SslPolicy: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

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

	return resourceComputeSslPolicyRead(d, meta)
}

func resourceComputeSslPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/sslPolicies/{{name}}")
	if err != nil {
		return err
	}

	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeSslPolicy %q", d.Id()))
	}
	if err := d.Set("creation_timestamp", flattenComputeSslPolicyCreationTimestamp(res["creationTimestamp"])); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("description", flattenComputeSslPolicyDescription(res["description"])); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("name", flattenComputeSslPolicyName(res["name"])); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("profile", flattenComputeSslPolicyProfile(res["profile"])); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("min_tls_version", flattenComputeSslPolicyMinTlsVersion(res["minTlsVersion"])); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("enabled_features", flattenComputeSslPolicyEnabledFeatures(res["enabledFeatures"])); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("custom_features", flattenComputeSslPolicyCustomFeatures(res["customFeatures"])); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("fingerprint", flattenComputeSslPolicyFingerprint(res["fingerprint"])); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("self_link", res["selfLink"]); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading SslPolicy: %s", err)
	}

	return nil
}

func resourceComputeSslPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	descriptionProp, err := expandComputeSslPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	}
	nameProp, err := expandComputeSslPolicyName(d.Get("name"), d, config)
	if err != nil {
		return err
	}
	profileProp, err := expandComputeSslPolicyProfile(d.Get("profile"), d, config)
	if err != nil {
		return err
	}
	minTlsVersionProp, err := expandComputeSslPolicyMinTlsVersion(d.Get("min_tls_version"), d, config)
	if err != nil {
		return err
	}
	customFeaturesProp, err := expandComputeSslPolicyCustomFeatures(d.Get("custom_features"), d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"description":    descriptionProp,
		"name":           nameProp,
		"profile":        profileProp,
		"minTlsVersion":  minTlsVersionProp,
		"customFeatures": customFeaturesProp,
	}

	obj, err = resourceComputeSslPolicyUpdateEncoder(d, meta, obj)
	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/sslPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating SslPolicy %q: %#v", d.Id(), obj)
	res, err := sendRequest(config, "PATCH", url, obj)

	if err != nil {
		return fmt.Errorf("Error updating SslPolicy %q: %s", d.Id(), err)
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

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/sslPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting SslPolicy %q", d.Id())
	res, err := Delete(config, url)
	if err != nil {
		return fmt.Errorf("Error deleting SslPolicy %q: %s", d.Id(), err)
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

	return nil
}

func resourceComputeSslPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"projects/(?P<project>[^/]+)/global/sslPolicies/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeSslPolicyCreationTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeSslPolicyDescription(v interface{}) interface{} {
	return v
}

func flattenComputeSslPolicyName(v interface{}) interface{} {
	return v
}

func flattenComputeSslPolicyProfile(v interface{}) interface{} {
	return v
}

func flattenComputeSslPolicyMinTlsVersion(v interface{}) interface{} {
	return v
}

func flattenComputeSslPolicyEnabledFeatures(v interface{}) interface{} {
	return v
}

func flattenComputeSslPolicyCustomFeatures(v interface{}) interface{} {
	return v
}

func flattenComputeSslPolicyFingerprint(v interface{}) interface{} {
	return v
}

func expandComputeSslPolicyDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslPolicyName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslPolicyProfile(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslPolicyMinTlsVersion(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeSslPolicyCustomFeatures(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
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
