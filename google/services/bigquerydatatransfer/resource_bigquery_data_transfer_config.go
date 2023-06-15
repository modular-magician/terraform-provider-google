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

package bigquerydatatransfer

import (
	"context"
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

var sensitiveParams = []string{"secret_access_key"}

func sensitiveParamCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	for _, sp := range sensitiveParams {
		mapLabel := diff.Get("params." + sp).(string)
		authLabel := diff.Get("sensitive_params.0." + sp).(string)
		if mapLabel != "" && authLabel != "" {
			return fmt.Errorf("Sensitive param [%s] cannot be set in both `params` and the `sensitive_params` block.", sp)
		}
	}
	return nil
}

// This customizeDiff is to use ForceNew for params fields data_path_template and
// destination_table_name_template only if the value of "data_source_id" is "google_cloud_storage".
func ParamsCustomizeDiffFunc(diff tpgresource.TerraformResourceDiff) error {
	old, new := diff.GetChange("params")
	dsId := diff.Get("data_source_id").(string)
	oldParams := old.(map[string]interface{})
	newParams := new.(map[string]interface{})
	var err error

	if dsId == "google_cloud_storage" {
		if oldParams["data_path_template"] != nil && newParams["data_path_template"] != nil && oldParams["data_path_template"].(string) != newParams["data_path_template"].(string) {
			err = diff.ForceNew("params")
			if err != nil {
				return fmt.Errorf("ForceNew failed for params, old - %v and new - %v", oldParams, newParams)
			}
			return nil
		}

		if oldParams["destination_table_name_template"] != nil && newParams["destination_table_name_template"] != nil && oldParams["destination_table_name_template"].(string) != newParams["destination_table_name_template"].(string) {
			err = diff.ForceNew("params")
			if err != nil {
				return fmt.Errorf("ForceNew failed for params, old - %v and new - %v", oldParams, newParams)
			}
			return nil
		}
	}

	return nil
}

func paramsCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	return ParamsCustomizeDiffFunc(diff)
}

func ResourceBigqueryDataTransferConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigqueryDataTransferConfigCreate,
		Read:   resourceBigqueryDataTransferConfigRead,
		Update: resourceBigqueryDataTransferConfigUpdate,
		Delete: resourceBigqueryDataTransferConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBigqueryDataTransferConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(sensitiveParamCustomizeDiff, paramsCustomizeDiff),

		Schema: map[string]*schema.Schema{
			"data_source_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The data source id. Cannot be changed once the transfer config is created.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The user specified display name for the transfer config.`,
			},
			"params": {
				Type:     schema.TypeMap,
				Required: true,
				Description: `Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer'
section for each data source. For example the parameters for Cloud Storage transfers are listed here:
https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq

**NOTE** : If you are attempting to update a parameter that cannot be updated (due to api limitations) [please force recreation of the resource](https://www.terraform.io/cli/state/taint#forcing-re-creation-of-resources).`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"data_refresh_window_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `The number of days to look back to automatically refresh the data.
For example, if dataRefreshWindowDays = 10, then every day BigQuery
reingests data for [today-10, today-1], rather than ingesting data for
just [today-1]. Only valid if the data source supports the feature.
Set the value to 0 to use the default value.`,
			},
			"destination_dataset_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The BigQuery target dataset id.`,
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `When set to true, no runs are scheduled for a given transfer.`,
			},
			"email_preferences": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Email notifications will be sent according to these preferences to the
email address of the user who owns this transfer config.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_failure_email": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: `If true, email notifications will be sent on transfer run failures.`,
						},
					},
				},
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The geographic location where the transfer config should reside.
Examples: US, EU, asia-northeast1. The default value is US.`,
				Default: "US",
			},
			"notification_pubsub_topic": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Pub/Sub topic where notifications will be sent after transfer runs
associated with this transfer config finish.`,
			},
			"schedule": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Data transfer schedule. If the data source does not support a custom
schedule, this should be empty. If it is empty, the default value for
the data source will be used. The specified times are in UTC. Examples
of valid format: 1st,3rd monday of month 15:30, every wed,fri of jan,
jun 13:15, and first sunday of quarter 00:00. See more explanation
about the format here:
https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
NOTE: the granularity should be at least 8 hours, or less frequent.`,
			},
			"schedule_options": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Options customizing the data transfer schedule.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable_auto_scheduling": {
							Type:     schema.TypeBool,
							Optional: true,
							Description: `If true, automatic scheduling of data transfer runs for this
configuration will be disabled. The runs can be started on ad-hoc
basis using transferConfigs.startManualRuns API. When automatic
scheduling is disabled, the TransferConfig.schedule field will
be ignored.`,
							AtLeastOneOf: []string{"schedule_options.0.disable_auto_scheduling", "schedule_options.0.start_time", "schedule_options.0.end_time"},
						},
						"end_time": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Defines time to stop scheduling transfer runs. A transfer run cannot be
scheduled at or after the end time. The end time can be changed at any
moment. The time when a data transfer can be triggered manually is not
limited by this option.`,
							AtLeastOneOf: []string{"schedule_options.0.disable_auto_scheduling", "schedule_options.0.start_time", "schedule_options.0.end_time"},
						},
						"start_time": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Specifies time to start scheduling transfer runs. The first run will be
scheduled at or after the start time according to a recurrence pattern
defined in the schedule string. The start time can be changed at any
moment. The time when a data transfer can be triggered manually is not
limited by this option.`,
							AtLeastOneOf: []string{"schedule_options.0.disable_auto_scheduling", "schedule_options.0.start_time", "schedule_options.0.end_time"},
						},
					},
				},
			},
			"sensitive_params": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Different parameters are configured primarily using the the 'params' field on this
resource. This block contains the parameters which contain secrets or passwords so that they can be marked
sensitive and hidden from plan output. The name of the field, eg: secret_access_key, will be the key
in the 'params' map in the api request.

Credentials may not be specified in both locations and will cause an error. Changing from one location
to a different credential configuration in the config will require an apply to update state.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"secret_access_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The Secret Access Key of the AWS account transferring data from.`,
							Sensitive:   true,
						},
					},
				},
			},
			"service_account_name": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Service account email. If this field is set, transfer config will
be created with this service account credentials. It requires that
requesting user calling this API has permissions to act as this service account.`,
				Default: "",
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the transfer config. Transfer config names have the
form projects/{projectId}/locations/{location}/transferConfigs/{configId}.
Where configId is usually a uuid, but this is not required.
The name is ignored when creating a transfer config.`,
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

func resourceBigqueryDataTransferConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandBigqueryDataTransferConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	destinationDatasetIdProp, err := expandBigqueryDataTransferConfigDestinationDatasetId(d.Get("destination_dataset_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("destination_dataset_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(destinationDatasetIdProp)) && (ok || !reflect.DeepEqual(v, destinationDatasetIdProp)) {
		obj["destinationDatasetId"] = destinationDatasetIdProp
	}
	dataSourceIdProp, err := expandBigqueryDataTransferConfigDataSourceId(d.Get("data_source_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("data_source_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataSourceIdProp)) && (ok || !reflect.DeepEqual(v, dataSourceIdProp)) {
		obj["dataSourceId"] = dataSourceIdProp
	}
	scheduleProp, err := expandBigqueryDataTransferConfigSchedule(d.Get("schedule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("schedule"); !tpgresource.IsEmptyValue(reflect.ValueOf(scheduleProp)) && (ok || !reflect.DeepEqual(v, scheduleProp)) {
		obj["schedule"] = scheduleProp
	}
	scheduleOptionsProp, err := expandBigqueryDataTransferConfigScheduleOptions(d.Get("schedule_options"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("schedule_options"); !tpgresource.IsEmptyValue(reflect.ValueOf(scheduleOptionsProp)) && (ok || !reflect.DeepEqual(v, scheduleOptionsProp)) {
		obj["scheduleOptions"] = scheduleOptionsProp
	}
	emailPreferencesProp, err := expandBigqueryDataTransferConfigEmailPreferences(d.Get("email_preferences"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("email_preferences"); !tpgresource.IsEmptyValue(reflect.ValueOf(emailPreferencesProp)) && (ok || !reflect.DeepEqual(v, emailPreferencesProp)) {
		obj["emailPreferences"] = emailPreferencesProp
	}
	notificationPubsubTopicProp, err := expandBigqueryDataTransferConfigNotificationPubsubTopic(d.Get("notification_pubsub_topic"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_pubsub_topic"); !tpgresource.IsEmptyValue(reflect.ValueOf(notificationPubsubTopicProp)) && (ok || !reflect.DeepEqual(v, notificationPubsubTopicProp)) {
		obj["notificationPubsubTopic"] = notificationPubsubTopicProp
	}
	dataRefreshWindowDaysProp, err := expandBigqueryDataTransferConfigDataRefreshWindowDays(d.Get("data_refresh_window_days"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("data_refresh_window_days"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataRefreshWindowDaysProp)) && (ok || !reflect.DeepEqual(v, dataRefreshWindowDaysProp)) {
		obj["dataRefreshWindowDays"] = dataRefreshWindowDaysProp
	}
	disabledProp, err := expandBigqueryDataTransferConfigDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	paramsProp, err := expandBigqueryDataTransferConfigParams(d.Get("params"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("params"); !tpgresource.IsEmptyValue(reflect.ValueOf(paramsProp)) && (ok || !reflect.DeepEqual(v, paramsProp)) {
		obj["params"] = paramsProp
	}

	obj, err = resourceBigqueryDataTransferConfigEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryDataTransferBasePath}}projects/{{project}}/locations/{{location}}/transferConfigs?serviceAccountName={{service_account_name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Config: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Config: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "POST",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutCreate),
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IamMemberMissing},
	})
	if err != nil {
		return fmt.Errorf("Error creating Config: %s", err)
	}
	if err := d.Set("name", flattenBigqueryDataTransferConfigName(res["name"], d, config)); err != nil {
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

	log.Printf("[DEBUG] Finished creating Config %q: %#v", d.Id(), res)

	return resourceBigqueryDataTransferConfigRead(d, meta)
}

func resourceBigqueryDataTransferConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryDataTransferBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Config: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "GET",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IamMemberMissing},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("BigqueryDataTransferConfig %q", d.Id()))
	}

	res, err = resourceBigqueryDataTransferConfigDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing BigqueryDataTransferConfig because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}

	if err := d.Set("display_name", flattenBigqueryDataTransferConfigDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("name", flattenBigqueryDataTransferConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("destination_dataset_id", flattenBigqueryDataTransferConfigDestinationDatasetId(res["destinationDatasetId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("data_source_id", flattenBigqueryDataTransferConfigDataSourceId(res["dataSourceId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("schedule", flattenBigqueryDataTransferConfigSchedule(res["schedule"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("schedule_options", flattenBigqueryDataTransferConfigScheduleOptions(res["scheduleOptions"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("email_preferences", flattenBigqueryDataTransferConfigEmailPreferences(res["emailPreferences"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("notification_pubsub_topic", flattenBigqueryDataTransferConfigNotificationPubsubTopic(res["notificationPubsubTopic"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("data_refresh_window_days", flattenBigqueryDataTransferConfigDataRefreshWindowDays(res["dataRefreshWindowDays"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("disabled", flattenBigqueryDataTransferConfigDisabled(res["disabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("params", flattenBigqueryDataTransferConfigParams(res["params"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}

	return nil
}

func resourceBigqueryDataTransferConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Config: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandBigqueryDataTransferConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	destinationDatasetIdProp, err := expandBigqueryDataTransferConfigDestinationDatasetId(d.Get("destination_dataset_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("destination_dataset_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, destinationDatasetIdProp)) {
		obj["destinationDatasetId"] = destinationDatasetIdProp
	}
	scheduleProp, err := expandBigqueryDataTransferConfigSchedule(d.Get("schedule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("schedule"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, scheduleProp)) {
		obj["schedule"] = scheduleProp
	}
	scheduleOptionsProp, err := expandBigqueryDataTransferConfigScheduleOptions(d.Get("schedule_options"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("schedule_options"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, scheduleOptionsProp)) {
		obj["scheduleOptions"] = scheduleOptionsProp
	}
	emailPreferencesProp, err := expandBigqueryDataTransferConfigEmailPreferences(d.Get("email_preferences"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("email_preferences"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, emailPreferencesProp)) {
		obj["emailPreferences"] = emailPreferencesProp
	}
	notificationPubsubTopicProp, err := expandBigqueryDataTransferConfigNotificationPubsubTopic(d.Get("notification_pubsub_topic"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_pubsub_topic"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, notificationPubsubTopicProp)) {
		obj["notificationPubsubTopic"] = notificationPubsubTopicProp
	}
	dataRefreshWindowDaysProp, err := expandBigqueryDataTransferConfigDataRefreshWindowDays(d.Get("data_refresh_window_days"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("data_refresh_window_days"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, dataRefreshWindowDaysProp)) {
		obj["dataRefreshWindowDays"] = dataRefreshWindowDaysProp
	}
	disabledProp, err := expandBigqueryDataTransferConfigDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	paramsProp, err := expandBigqueryDataTransferConfigParams(d.Get("params"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("params"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, paramsProp)) {
		obj["params"] = paramsProp
	}

	obj, err = resourceBigqueryDataTransferConfigEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryDataTransferBasePath}}{{name}}?serviceAccountName={{service_account_name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Config %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("destination_dataset_id") {
		updateMask = append(updateMask, "destinationDatasetId")
	}

	if d.HasChange("schedule") {
		updateMask = append(updateMask, "schedule")
	}

	if d.HasChange("schedule_options") {
		updateMask = append(updateMask, "scheduleOptions")
	}

	if d.HasChange("email_preferences") {
		updateMask = append(updateMask, "emailPreferences")
	}

	if d.HasChange("notification_pubsub_topic") {
		updateMask = append(updateMask, "notificationPubsubTopic")
	}

	if d.HasChange("data_refresh_window_days") {
		updateMask = append(updateMask, "dataRefreshWindowDays")
	}

	if d.HasChange("disabled") {
		updateMask = append(updateMask, "disabled")
	}

	if d.HasChange("params") {
		updateMask = append(updateMask, "params")
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
		Config:               config,
		Method:               "PATCH",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutUpdate),
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IamMemberMissing},
	})

	if err != nil {
		return fmt.Errorf("Error updating Config %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Config %q: %#v", d.Id(), res)
	}

	return resourceBigqueryDataTransferConfigRead(d, meta)
}

func resourceBigqueryDataTransferConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Config: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{BigqueryDataTransferBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Config %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "DELETE",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutDelete),
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IamMemberMissing},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Config")
	}

	log.Printf("[DEBUG] Finished deleting Config %q: %#v", d.Id(), res)
	return nil
}

func resourceBigqueryDataTransferConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<project>[^ ]+) (?P<name>[^ ]+)", "(?P<name>[^ ]+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenBigqueryDataTransferConfigDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigDestinationDatasetId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigDataSourceId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigSchedule(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigScheduleOptions(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["disable_auto_scheduling"] =
		flattenBigqueryDataTransferConfigScheduleOptionsDisableAutoScheduling(original["disableAutoScheduling"], d, config)
	transformed["start_time"] =
		flattenBigqueryDataTransferConfigScheduleOptionsStartTime(original["startTime"], d, config)
	transformed["end_time"] =
		flattenBigqueryDataTransferConfigScheduleOptionsEndTime(original["endTime"], d, config)
	return []interface{}{transformed}
}
func flattenBigqueryDataTransferConfigScheduleOptionsDisableAutoScheduling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigScheduleOptionsStartTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigScheduleOptionsEndTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigEmailPreferences(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enable_failure_email"] =
		flattenBigqueryDataTransferConfigEmailPreferencesEnableFailureEmail(original["enableFailureEmail"], d, config)
	return []interface{}{transformed}
}
func flattenBigqueryDataTransferConfigEmailPreferencesEnableFailureEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigNotificationPubsubTopic(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigDataRefreshWindowDays(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenBigqueryDataTransferConfigDisabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigqueryDataTransferConfigParams(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	kv := v.(map[string]interface{})

	res := make(map[string]string)
	for key, value := range kv {
		res[key] = fmt.Sprintf("%v", value)
	}
	return res
}

func expandBigqueryDataTransferConfigDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDestinationDatasetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDataSourceId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigScheduleOptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDisableAutoScheduling, err := expandBigqueryDataTransferConfigScheduleOptionsDisableAutoScheduling(original["disable_auto_scheduling"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisableAutoScheduling); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["disableAutoScheduling"] = transformedDisableAutoScheduling
	}

	transformedStartTime, err := expandBigqueryDataTransferConfigScheduleOptionsStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	transformedEndTime, err := expandBigqueryDataTransferConfigScheduleOptionsEndTime(original["end_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEndTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["endTime"] = transformedEndTime
	}

	return transformed, nil
}

func expandBigqueryDataTransferConfigScheduleOptionsDisableAutoScheduling(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigScheduleOptionsStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigScheduleOptionsEndTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigEmailPreferences(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableFailureEmail, err := expandBigqueryDataTransferConfigEmailPreferencesEnableFailureEmail(original["enable_failure_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableFailureEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableFailureEmail"] = transformedEnableFailureEmail
	}

	return transformed, nil
}

func expandBigqueryDataTransferConfigEmailPreferencesEnableFailureEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigNotificationPubsubTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDataRefreshWindowDays(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigParams(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func resourceBigqueryDataTransferConfigEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	paramMap, ok := obj["params"]
	if !ok {
		paramMap = make(map[string]string)
	}

	var params map[string]string
	params = paramMap.(map[string]string)

	for _, sp := range sensitiveParams {
		if auth, _ := d.GetOkExists("sensitive_params.0." + sp); auth != "" {
			params[sp] = auth.(string)
		}
	}

	obj["params"] = params

	return obj, nil
}

func resourceBigqueryDataTransferConfigDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if paramMap, ok := res["params"]; ok {
		params := paramMap.(map[string]interface{})
		for _, sp := range sensitiveParams {
			if _, apiOk := params[sp]; apiOk {
				if _, exists := d.GetOkExists("sensitive_params.0." + sp); exists {
					delete(params, sp)
				} else {
					params[sp] = d.Get("params." + sp)
				}
			}
		}
	}

	return res, nil
}
