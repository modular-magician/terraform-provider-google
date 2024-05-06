// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package networksecurity

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceNetworkSecuritySecurityProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecuritySecurityProfileCreate,
		Read:   resourceNetworkSecuritySecurityProfileRead,
		Update: resourceNetworkSecuritySecurityProfileUpdate,
		Delete: resourceNetworkSecuritySecurityProfileDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkSecuritySecurityProfileImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the security profile resource.`,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"THREAT_PREVENTION"}),
				Description:  `The type of security profile. Possible values: ["THREAT_PREVENTION"]`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of the security profile. The Max length is 512 characters.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `A map of key/value label pairs to assign to the resource.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The location of the security profile.
The default value is 'global'.`,
				Default: "global",
			},
			"parent": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The name of the parent this security profile belongs to.
Format: organizations/{organization_id}.`,
			},
			"threat_prevention_profile": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The threat prevention configuration for the security profile.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"severity_overrides": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `The configuration for overriding threats actions by severity match.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: verify.ValidateEnum([]string{"ALERT", "ALLOW", "DEFAULT_ACTION", "DENY"}),
										Description:  `Threat action override. Possible values: ["ALERT", "ALLOW", "DEFAULT_ACTION", "DENY"]`,
									},
									"severity": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: verify.ValidateEnum([]string{"CRITICAL", "HIGH", "INFORMATIONAL", "LOW", "MEDIUM"}),
										Description:  `Severity level to match. Possible values: ["CRITICAL", "HIGH", "INFORMATIONAL", "LOW", "MEDIUM"]`,
									},
								},
							},
						},
						"threat_overrides": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `The configuration for overriding threats actions by threat id match.
If a threat is matched both by configuration provided in severity overrides
and threat overrides, the threat overrides action is applied.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: verify.ValidateEnum([]string{"ALERT", "ALLOW", "DEFAULT_ACTION", "DENY"}),
										Description:  `Threat action. Possible values: ["ALERT", "ALLOW", "DEFAULT_ACTION", "DENY"]`,
									},
									"threat_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Vendor-specific ID of a threat to override.`,
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `Type of threat.`,
									},
								},
							},
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the security profile was created in UTC.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `This checksum is computed by the server based on the value of other fields,
and may be sent on update and delete requests to ensure the client has an up-to-date
value before proceeding.`,
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL of this resource.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the security profile was updated in UTC.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceNetworkSecuritySecurityProfileCreate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecuritySecurityProfileDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	threatPreventionProfileProp, err := expandNetworkSecuritySecurityProfileThreatPreventionProfile(d.Get("threat_prevention_profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("threat_prevention_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(threatPreventionProfileProp)) && (ok || !reflect.DeepEqual(v, threatPreventionProfileProp)) {
		obj["threatPreventionProfile"] = threatPreventionProfileProp
	}
	typeProp, err := expandNetworkSecuritySecurityProfileType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	labelsProp, err := expandNetworkSecuritySecurityProfileEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/securityProfiles?securityProfileId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SecurityProfile: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating SecurityProfile: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/locations/{{location}}/securityProfiles/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Creating SecurityProfile", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create SecurityProfile: %s", err)
	}

	log.Printf("[DEBUG] Finished creating SecurityProfile %q: %#v", d.Id(), res)

	return resourceNetworkSecuritySecurityProfileRead(d, meta)
}

func resourceNetworkSecuritySecurityProfileRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/securityProfiles/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkSecuritySecurityProfile %q", d.Id()))
	}

	if err := d.Set("self_link", flattenNetworkSecuritySecurityProfileSelfLink(res["selfLink"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfile: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkSecuritySecurityProfileCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfile: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkSecuritySecurityProfileUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfile: %s", err)
	}
	if err := d.Set("etag", flattenNetworkSecuritySecurityProfileEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfile: %s", err)
	}
	if err := d.Set("description", flattenNetworkSecuritySecurityProfileDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfile: %s", err)
	}
	if err := d.Set("labels", flattenNetworkSecuritySecurityProfileLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfile: %s", err)
	}
	if err := d.Set("threat_prevention_profile", flattenNetworkSecuritySecurityProfileThreatPreventionProfile(res["threatPreventionProfile"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfile: %s", err)
	}
	if err := d.Set("type", flattenNetworkSecuritySecurityProfileType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfile: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkSecuritySecurityProfileTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfile: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkSecuritySecurityProfileEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecurityProfile: %s", err)
	}

	return nil
}

func resourceNetworkSecuritySecurityProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecuritySecurityProfileDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	threatPreventionProfileProp, err := expandNetworkSecuritySecurityProfileThreatPreventionProfile(d.Get("threat_prevention_profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("threat_prevention_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, threatPreventionProfileProp)) {
		obj["threatPreventionProfile"] = threatPreventionProfileProp
	}
	labelsProp, err := expandNetworkSecuritySecurityProfileEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/securityProfiles/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating SecurityProfile %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("threat_prevention_profile") {
		updateMask = append(updateMask, "threatPreventionProfile")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating SecurityProfile %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating SecurityProfile %q: %#v", d.Id(), res)
		}

		err = NetworkSecurityOperationWaitTime(
			config, res, project, "Updating SecurityProfile", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkSecuritySecurityProfileRead(d, meta)
}

func resourceNetworkSecuritySecurityProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/securityProfiles/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting SecurityProfile %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "SecurityProfile")
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Deleting SecurityProfile", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting SecurityProfile %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkSecuritySecurityProfileImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<parent>.+)/locations/(?P<location>[^/]+)/securityProfiles/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/locations/{{location}}/securityProfiles/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkSecuritySecurityProfileSelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenNetworkSecuritySecurityProfileThreatPreventionProfile(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["severity_overrides"] =
		flattenNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverrides(original["severityOverrides"], d, config)
	transformed["threat_overrides"] =
		flattenNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverrides(original["threatOverrides"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverrides(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"action":   flattenNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverridesAction(original["action"], d, config),
			"severity": flattenNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverridesSeverity(original["severity"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverridesAction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverridesSeverity(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverrides(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"action":    flattenNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesAction(original["action"], d, config),
			"threat_id": flattenNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesThreatId(original["threatId"], d, config),
			"type":      flattenNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesType(original["type"], d, config),
		})
	}
	return transformed
}
func flattenNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesAction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesThreatId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecuritySecurityProfileTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenNetworkSecuritySecurityProfileEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkSecuritySecurityProfileDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSeverityOverrides, err := expandNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverrides(original["severity_overrides"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSeverityOverrides); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["severityOverrides"] = transformedSeverityOverrides
	}

	transformedThreatOverrides, err := expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverrides(original["threat_overrides"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedThreatOverrides); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["threatOverrides"] = transformedThreatOverrides
	}

	return transformed, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverrides(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAction, err := expandNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverridesAction(original["action"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["action"] = transformedAction
		}

		transformedSeverity, err := expandNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverridesSeverity(original["severity"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSeverity); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["severity"] = transformedSeverity
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverridesAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverridesSeverity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverrides(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAction, err := expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesAction(original["action"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["action"] = transformedAction
		}

		transformedThreatId, err := expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesThreatId(original["threat_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedThreatId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["threatId"] = transformedThreatId
		}

		transformedType, err := expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["type"] = transformedType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesThreatId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
