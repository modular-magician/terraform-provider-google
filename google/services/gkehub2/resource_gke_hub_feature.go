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

package gkehub2

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceGKEHub2Feature() *schema.Resource {
	return &schema.Resource{
		Create: resourceGKEHub2FeatureCreate,
		Read:   resourceGKEHub2FeatureRead,
		Update: resourceGKEHub2FeatureUpdate,
		Delete: resourceGKEHub2FeatureDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGKEHub2FeatureImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location for the resource`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `GCP labels for this Feature.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The full, unique name of this Feature resource`,
			},
			"spec": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"multiclusteringress": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Multicluster Ingress-specific spec.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"config_membership": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `Fully-qualified Membership name which hosts the MultiClusterIngress CRD. Example: 'projects/foo-proj/locations/global/memberships/bar'`,
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
				Description: `Output only. When the Feature resource was created.`,
			},
			"delete_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. When the Feature resource was deleted.`,
			},
			"resource_state": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `State of the Feature resource itself.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"has_resources": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: `Whether this Feature has outstanding resources that need to be cleaned up before it can be disabled.`,
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The current state of the Feature resource in the Hub API.`,
						},
					},
				},
			},
			"state": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Output only. The Hub-wide Feature state`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `Output only. The "running state" of the Feature in this Hub.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `The high-level, machine-readable status of this Feature.`,
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `A human-readable description of the current status.`,
									},
									"update_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `The time this status and any related Feature-specific details were updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"`,
									},
								},
							},
						},
					},
				},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. When the Feature resource was last updated.`,
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

func resourceGKEHub2FeatureCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandGKEHub2FeatureLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	specProp, err := expandGKEHub2FeatureSpec(d.Get("spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(specProp)) && (ok || !reflect.DeepEqual(v, specProp)) {
		obj["spec"] = specProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/{{location}}/features?featureId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Feature: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Feature: %s", err)
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
		return fmt.Errorf("Error creating Feature: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = GKEHub2OperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Feature", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Feature: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Feature %q: %#v", d.Id(), res)

	return resourceGKEHub2FeatureRead(d, meta)
}

func resourceGKEHub2FeatureRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Feature: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("GKEHub2Feature %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}

	if err := d.Set("labels", flattenGKEHub2FeatureLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}
	if err := d.Set("resource_state", flattenGKEHub2FeatureResourceState(res["resourceState"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}
	if err := d.Set("spec", flattenGKEHub2FeatureSpec(res["spec"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}
	if err := d.Set("state", flattenGKEHub2FeatureState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}
	if err := d.Set("create_time", flattenGKEHub2FeatureCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}
	if err := d.Set("update_time", flattenGKEHub2FeatureUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}
	if err := d.Set("delete_time", flattenGKEHub2FeatureDeleteTime(res["deleteTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Feature: %s", err)
	}

	return nil
}

func resourceGKEHub2FeatureUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Feature: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandGKEHub2FeatureLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	specProp, err := expandGKEHub2FeatureSpec(d.Get("spec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, specProp)) {
		obj["spec"] = specProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Feature %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("spec") {
		updateMask = append(updateMask, "spec")
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

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating Feature %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Feature %q: %#v", d.Id(), res)
	}

	err = GKEHub2OperationWaitTime(
		config, res, project, "Updating Feature", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceGKEHub2FeatureRead(d, meta)
}

func resourceGKEHub2FeatureDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Feature: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{GKEHub2BasePath}}projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Feature %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Feature")
	}

	err = GKEHub2OperationWaitTime(
		config, res, project, "Deleting Feature", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Feature %q: %#v", d.Id(), res)
	return nil
}

func resourceGKEHub2FeatureImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/features/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/features/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenGKEHub2FeatureLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureResourceState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["state"] =
		flattenGKEHub2FeatureResourceStateState(original["state"], d, config)
	transformed["has_resources"] =
		flattenGKEHub2FeatureResourceStateHasResources(original["hasResources"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FeatureResourceStateState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureResourceStateHasResources(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureSpec(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["multiclusteringress"] =
		flattenGKEHub2FeatureSpecMulticlusteringress(original["multiclusteringress"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FeatureSpecMulticlusteringress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["config_membership"] =
		flattenGKEHub2FeatureSpecMulticlusteringressConfigMembership(original["configMembership"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FeatureSpecMulticlusteringressConfigMembership(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["state"] =
		flattenGKEHub2FeatureStateState(original["state"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FeatureStateState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["code"] =
		flattenGKEHub2FeatureStateStateCode(original["code"], d, config)
	transformed["description"] =
		flattenGKEHub2FeatureStateStateDescription(original["description"], d, config)
	transformed["update_time"] =
		flattenGKEHub2FeatureStateStateUpdateTime(original["updateTime"], d, config)
	return []interface{}{transformed}
}
func flattenGKEHub2FeatureStateStateCode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureStateStateDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureStateStateUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGKEHub2FeatureDeleteTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandGKEHub2FeatureLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGKEHub2FeatureSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMulticlusteringress, err := expandGKEHub2FeatureSpecMulticlusteringress(original["multiclusteringress"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMulticlusteringress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["multiclusteringress"] = transformedMulticlusteringress
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecMulticlusteringress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedConfigMembership, err := expandGKEHub2FeatureSpecMulticlusteringressConfigMembership(original["config_membership"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConfigMembership); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["configMembership"] = transformedConfigMembership
	}

	return transformed, nil
}

func expandGKEHub2FeatureSpecMulticlusteringressConfigMembership(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
