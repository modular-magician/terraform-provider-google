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

package netapp

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceNetappbackupPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetappbackupPolicyCreate,
		Read:   resourceNetappbackupPolicyRead,
		Update: resourceNetappbackupPolicyUpdate,
		Delete: resourceNetappbackupPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetappbackupPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"daily_backup_limit": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Number of daily backups to keep. Note that the minimum daily backup limit is 2.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the region for the policy to apply to.`,
			},
			"monthly_backup_limit": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Number of monthly backups to keep. Note that the sum of daily, weekly and monthly backups should be greater than 1.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the backup policy. Needs to be unique per location.`,
			},
			"weekly_backup_limit": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: `Number of weekly backups to keep. Note that the sum of daily, weekly and monthly backups should be greater than 1.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of this resource.`,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `If enabled, make backups automatically according to the schedules.
This will be applied to all volumes that have this policy attached and enforced on volume level.`,
				Default: true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"assigned_volume_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The total number of volumes assigned by this backup policy.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Create time of the backup policy. A timestamp in RFC3339 UTC "Zulu" format. Examples: "2023-06-22T09:13:01.617Z".`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The state of the backup policy.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceNetappbackupPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	dailyBackupLimitProp, err := expandNetappbackupPolicyDailyBackupLimit(d.Get("daily_backup_limit"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("daily_backup_limit"); !tpgresource.IsEmptyValue(reflect.ValueOf(dailyBackupLimitProp)) && (ok || !reflect.DeepEqual(v, dailyBackupLimitProp)) {
		obj["dailyBackupLimit"] = dailyBackupLimitProp
	}
	weeklyBackupLimitProp, err := expandNetappbackupPolicyWeeklyBackupLimit(d.Get("weekly_backup_limit"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("weekly_backup_limit"); !tpgresource.IsEmptyValue(reflect.ValueOf(weeklyBackupLimitProp)) && (ok || !reflect.DeepEqual(v, weeklyBackupLimitProp)) {
		obj["weeklyBackupLimit"] = weeklyBackupLimitProp
	}
	monthlyBackupLimitProp, err := expandNetappbackupPolicyMonthlyBackupLimit(d.Get("monthly_backup_limit"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("monthly_backup_limit"); !tpgresource.IsEmptyValue(reflect.ValueOf(monthlyBackupLimitProp)) && (ok || !reflect.DeepEqual(v, monthlyBackupLimitProp)) {
		obj["monthlyBackupLimit"] = monthlyBackupLimitProp
	}
	descriptionProp, err := expandNetappbackupPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	enabledProp, err := expandNetappbackupPolicyEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); ok || !reflect.DeepEqual(v, enabledProp) {
		obj["enabled"] = enabledProp
	}
	labelsProp, err := expandNetappbackupPolicyEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetappBasePath}}projects/{{project}}/locations/{{location}}/backupPolicies?backupPolicyId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new backupPolicy: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for backupPolicy: %s", err)
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
		return fmt.Errorf("Error creating backupPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backupPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetappOperationWaitTime(
		config, res, project, "Creating backupPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create backupPolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating backupPolicy %q: %#v", d.Id(), res)

	return resourceNetappbackupPolicyRead(d, meta)
}

func resourceNetappbackupPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetappBasePath}}projects/{{project}}/locations/{{location}}/backupPolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for backupPolicy: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetappbackupPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading backupPolicy: %s", err)
	}

	if err := d.Set("create_time", flattenNetappbackupPolicyCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading backupPolicy: %s", err)
	}
	if err := d.Set("labels", flattenNetappbackupPolicyLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading backupPolicy: %s", err)
	}
	if err := d.Set("state", flattenNetappbackupPolicyState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading backupPolicy: %s", err)
	}
	if err := d.Set("daily_backup_limit", flattenNetappbackupPolicyDailyBackupLimit(res["dailyBackupLimit"], d, config)); err != nil {
		return fmt.Errorf("Error reading backupPolicy: %s", err)
	}
	if err := d.Set("weekly_backup_limit", flattenNetappbackupPolicyWeeklyBackupLimit(res["weeklyBackupLimit"], d, config)); err != nil {
		return fmt.Errorf("Error reading backupPolicy: %s", err)
	}
	if err := d.Set("monthly_backup_limit", flattenNetappbackupPolicyMonthlyBackupLimit(res["monthlyBackupLimit"], d, config)); err != nil {
		return fmt.Errorf("Error reading backupPolicy: %s", err)
	}
	if err := d.Set("description", flattenNetappbackupPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading backupPolicy: %s", err)
	}
	if err := d.Set("enabled", flattenNetappbackupPolicyEnabled(res["enabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading backupPolicy: %s", err)
	}
	if err := d.Set("assigned_volume_count", flattenNetappbackupPolicyAssignedVolumeCount(res["assignedVolumeCount"], d, config)); err != nil {
		return fmt.Errorf("Error reading backupPolicy: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetappbackupPolicyTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading backupPolicy: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetappbackupPolicyEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading backupPolicy: %s", err)
	}

	return nil
}

func resourceNetappbackupPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for backupPolicy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	dailyBackupLimitProp, err := expandNetappbackupPolicyDailyBackupLimit(d.Get("daily_backup_limit"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("daily_backup_limit"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, dailyBackupLimitProp)) {
		obj["dailyBackupLimit"] = dailyBackupLimitProp
	}
	weeklyBackupLimitProp, err := expandNetappbackupPolicyWeeklyBackupLimit(d.Get("weekly_backup_limit"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("weekly_backup_limit"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, weeklyBackupLimitProp)) {
		obj["weeklyBackupLimit"] = weeklyBackupLimitProp
	}
	monthlyBackupLimitProp, err := expandNetappbackupPolicyMonthlyBackupLimit(d.Get("monthly_backup_limit"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("monthly_backup_limit"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, monthlyBackupLimitProp)) {
		obj["monthlyBackupLimit"] = monthlyBackupLimitProp
	}
	descriptionProp, err := expandNetappbackupPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	enabledProp, err := expandNetappbackupPolicyEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); ok || !reflect.DeepEqual(v, enabledProp) {
		obj["enabled"] = enabledProp
	}
	labelsProp, err := expandNetappbackupPolicyEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetappBasePath}}projects/{{project}}/locations/{{location}}/backupPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating backupPolicy %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("daily_backup_limit") {
		updateMask = append(updateMask, "dailyBackupLimit")
	}

	if d.HasChange("weekly_backup_limit") {
		updateMask = append(updateMask, "weeklyBackupLimit")
	}

	if d.HasChange("monthly_backup_limit") {
		updateMask = append(updateMask, "monthlyBackupLimit")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("enabled") {
		updateMask = append(updateMask, "enabled")
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
		})

		if err != nil {
			return fmt.Errorf("Error updating backupPolicy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating backupPolicy %q: %#v", d.Id(), res)
		}

		err = NetappOperationWaitTime(
			config, res, project, "Updating backupPolicy", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetappbackupPolicyRead(d, meta)
}

func resourceNetappbackupPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for backupPolicy: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetappBasePath}}projects/{{project}}/locations/{{location}}/backupPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	log.Printf("[DEBUG] Deleting backupPolicy %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "backupPolicy")
	}

	err = NetappOperationWaitTime(
		config, res, project, "Deleting backupPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting backupPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceNetappbackupPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/backupPolicies/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backupPolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetappbackupPolicyCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetappbackupPolicyLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetappbackupPolicyState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetappbackupPolicyDailyBackupLimit(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
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

func flattenNetappbackupPolicyWeeklyBackupLimit(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
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

func flattenNetappbackupPolicyMonthlyBackupLimit(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
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

func flattenNetappbackupPolicyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetappbackupPolicyEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetappbackupPolicyAssignedVolumeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
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

func flattenNetappbackupPolicyTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetappbackupPolicyEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetappbackupPolicyDailyBackupLimit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappbackupPolicyWeeklyBackupLimit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappbackupPolicyMonthlyBackupLimit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappbackupPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappbackupPolicyEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappbackupPolicyEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
