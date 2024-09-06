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

package monitoring

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
)

func ResourceMonitoringSnooze() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitoringSnoozeCreate,
		Read:   resourceMonitoringSnoozeRead,
		Update: resourceMonitoringSnoozeUpdate,
		Delete: resourceMonitoringSnoozeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMonitoringSnoozeImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"criteria": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `This defines the criteria for applying the Snooze`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policies": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `The specific AlertPolicy names for the alert that should be snoozed.
The format is: 'projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[POLICY_ID]'
There is a limit of 16 policies per snooze.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				Description: `A short name or phrase used to identify the Snooze.
The display name is limited to 512 Unicode characters.`,
			},
			"interval": {
				Type:     schema.TypeList,
				Required: true,
				Description: `The Snooze will be active from interval.start_time through interval.end_time.
The interval.start_time cannot be in the past.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_time": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The end of the time interval.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
						},
						"start_time": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The beginning of the time interval. The default value for the start time is the end time.
The start time must not be later than the end time.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The unique resource name for this Snooze.
Its syntax is: projects/[PROJECT_ID_OR_NUMBER]/snoozes/[SNOOZE_ID]`,
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

func resourceMonitoringSnoozeCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandMonitoringSnoozeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	criteriaProp, err := expandMonitoringSnoozeCriteria(d.Get("criteria"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("criteria"); !tpgresource.IsEmptyValue(reflect.ValueOf(criteriaProp)) && (ok || !reflect.DeepEqual(v, criteriaProp)) {
		obj["criteria"] = criteriaProp
	}
	intervalProp, err := expandMonitoringSnoozeInterval(d.Get("interval"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("interval"); !tpgresource.IsEmptyValue(reflect.ValueOf(intervalProp)) && (ok || !reflect.DeepEqual(v, intervalProp)) {
		obj["interval"] = intervalProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{MonitoringBasePath}}v3/projects/{{project}}/snoozes")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Snooze: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Snooze: %s", err)
	}
	billingProject = project

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
		return fmt.Errorf("Error creating Snooze: %s", err)
	}
	if err := d.Set("name", flattenMonitoringSnoozeName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
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

	log.Printf("[DEBUG] Finished creating Snooze %q: %#v", d.Id(), res)

	return resourceMonitoringSnoozeRead(d, meta)
}

func resourceMonitoringSnoozeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Snooze: %s", err)
	}
	billingProject = project

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("MonitoringSnooze %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Snooze: %s", err)
	}

	if err := d.Set("name", flattenMonitoringSnoozeName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snooze: %s", err)
	}
	if err := d.Set("display_name", flattenMonitoringSnoozeDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snooze: %s", err)
	}
	if err := d.Set("criteria", flattenMonitoringSnoozeCriteria(res["criteria"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snooze: %s", err)
	}
	if err := d.Set("interval", flattenMonitoringSnoozeInterval(res["interval"], d, config)); err != nil {
		return fmt.Errorf("Error reading Snooze: %s", err)
	}

	return nil
}

func resourceMonitoringSnoozeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Snooze: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandMonitoringSnoozeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	criteriaProp, err := expandMonitoringSnoozeCriteria(d.Get("criteria"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("criteria"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, criteriaProp)) {
		obj["criteria"] = criteriaProp
	}
	intervalProp, err := expandMonitoringSnoozeInterval(d.Get("interval"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("interval"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, intervalProp)) {
		obj["interval"] = intervalProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Snooze %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("criteria") {
		updateMask = append(updateMask, "criteria")
	}

	if d.HasChange("interval") {
		updateMask = append(updateMask, "interval")
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
			return fmt.Errorf("Error updating Snooze %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Snooze %q: %#v", d.Id(), res)
		}

	}

	return resourceMonitoringSnoozeRead(d, meta)
}

func resourceMonitoringSnoozeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Snooze: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{MonitoringBasePath}}v3/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Snooze %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "Snooze")
	}

	log.Printf("[DEBUG] Finished deleting Snooze %q: %#v", d.Id(), res)
	return nil
}

func resourceMonitoringSnoozeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<project>[^ ]+) (?P<name>[^ ]+)", "(?P<name>[^ ]+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenMonitoringSnoozeName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringSnoozeDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringSnoozeCriteria(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["policies"] =
		flattenMonitoringSnoozeCriteriaPolicies(original["policies"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringSnoozeCriteriaPolicies(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringSnoozeInterval(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["start_time"] =
		flattenMonitoringSnoozeIntervalStartTime(original["startTime"], d, config)
	transformed["end_time"] =
		flattenMonitoringSnoozeIntervalEndTime(original["endTime"], d, config)
	return []interface{}{transformed}
}
func flattenMonitoringSnoozeIntervalStartTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenMonitoringSnoozeIntervalEndTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandMonitoringSnoozeDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringSnoozeCriteria(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPolicies, err := expandMonitoringSnoozeCriteriaPolicies(original["policies"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPolicies); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["policies"] = transformedPolicies
	}

	return transformed, nil
}

func expandMonitoringSnoozeCriteriaPolicies(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringSnoozeInterval(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStartTime, err := expandMonitoringSnoozeIntervalStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	transformedEndTime, err := expandMonitoringSnoozeIntervalEndTime(original["end_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEndTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["endTime"] = transformedEndTime
	}

	return transformed, nil
}

func expandMonitoringSnoozeIntervalStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringSnoozeIntervalEndTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
