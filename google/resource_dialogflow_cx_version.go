// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceDialogflowCXVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceDialogflowCXVersionCreate,
		Read:   resourceDialogflowCXVersionRead,
		Update: resourceDialogflowCXVersionUpdate,
		Delete: resourceDialogflowCXVersionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDialogflowCXVersionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(40 * time.Minute),
			Update: schema.DefaultTimeout(40 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Description:  `The human-readable name of the version. Limit of 64 characters.`,
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 500),
				Description:  `The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.`,
			},
			"parent": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The Flow to create an Version for.
Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow upon version creation.`,
			},
			"nlu_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The NLU settings of the flow at version creation.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"classification_threshold": {
							Type:     schema.TypeFloat,
							Optional: true,
							Description: `To filter out false positive results and still get variety in matched natural language inputs for your agent, you can tune the machine learning classification threshold. If the returned score value is less than the threshold value, then a no-match event will be triggered.
The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.`,
						},
						"model_training_mode": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"MODEL_TRAINING_MODE_AUTOMATIC", "MODEL_TRAINING_MODE_MANUAL", ""}),
							Description: `Indicates NLU model training mode.
* MODEL_TRAINING_MODE_AUTOMATIC: NLU model training is automatically triggered when a flow gets modified. User can also manually trigger model training in this mode.
* MODEL_TRAINING_MODE_MANUAL: User needs to manually trigger NLU model training. Best for large flows whose models take long time to train. Possible values: ["MODEL_TRAINING_MODE_AUTOMATIC", "MODEL_TRAINING_MODE_MANUAL"]`,
						},
						"model_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"MODEL_TYPE_STANDARD", "MODEL_TYPE_ADVANCED", ""}),
							Description: `Indicates the type of NLU model.
* MODEL_TYPE_STANDARD: Use standard NLU model.
* MODEL_TYPE_ADVANCED: Use advanced NLU model. Possible values: ["MODEL_TYPE_STANDARD", "MODEL_TYPE_ADVANCED"]`,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The state of this version.
* RUNNING: Version is not ready to serve (e.g. training is running).
* SUCCEEDED: Training has succeeded and this version is ready to serve.
* FAILED: Version training failed.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDialogflowCXVersionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXVersionDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDialogflowCXVersionDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/versions")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Version: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// extract location from the parent
	location := ""

	if parts := regexp.MustCompile(`locations\/([^\/]*)\/`).FindStringSubmatch(d.Get("parent").(string)); parts != nil {
		location = parts[1]
	} else {
		return fmt.Errorf(
			"Saw %s when the parent is expected to contains location %s",
			d.Get("parent"),
			"projects/{{project}}/locations/{{location}}/...",
		)
	}

	url = strings.Replace(url, "-dialogflow", fmt.Sprintf("%s-dialogflow", location), 1)
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
		return fmt.Errorf("Error creating Version: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/versions/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = DialogflowCXOperationWaitTimeWithResponse(
		config, res, &opRes, "Creating Version", userAgent, location,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Version: %s", err)
	}

	if err := d.Set("name", flattenDialogflowCXVersionName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "{{parent}}/versions/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Version %q: %#v", d.Id(), res)

	return resourceDialogflowCXVersionRead(d, meta)
}

func resourceDialogflowCXVersionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/versions/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// extract location from the parent
	location := ""

	if parts := regexp.MustCompile(`locations\/([^\/]*)\/`).FindStringSubmatch(d.Get("parent").(string)); parts != nil {
		location = parts[1]
	} else {
		return fmt.Errorf(
			"Saw %s when the parent is expected to contains location %s",
			d.Get("parent"),
			"projects/{{project}}/locations/{{location}}/...",
		)
	}

	url = strings.Replace(url, "-dialogflow", fmt.Sprintf("%s-dialogflow", location), 1)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DialogflowCXVersion %q", d.Id()))
	}

	if err := d.Set("name", flattenDialogflowCXVersionName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Version: %s", err)
	}
	if err := d.Set("display_name", flattenDialogflowCXVersionDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Version: %s", err)
	}
	if err := d.Set("description", flattenDialogflowCXVersionDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Version: %s", err)
	}
	if err := d.Set("nlu_settings", flattenDialogflowCXVersionNluSettings(res["nluSettings"], d, config)); err != nil {
		return fmt.Errorf("Error reading Version: %s", err)
	}
	if err := d.Set("create_time", flattenDialogflowCXVersionCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Version: %s", err)
	}
	if err := d.Set("state", flattenDialogflowCXVersionState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Version: %s", err)
	}

	return nil
}

func resourceDialogflowCXVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXVersionDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDialogflowCXVersionDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/versions/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Version %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
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

	// extract location from the parent
	location := ""

	if parts := regexp.MustCompile(`locations\/([^\/]*)\/`).FindStringSubmatch(d.Get("parent").(string)); parts != nil {
		location = parts[1]
	} else {
		return fmt.Errorf(
			"Saw %s when the parent is expected to contains location %s",
			d.Get("parent"),
			"projects/{{project}}/locations/{{location}}/...",
		)
	}

	url = strings.Replace(url, "-dialogflow", fmt.Sprintf("%s-dialogflow", location), 1)

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
		return fmt.Errorf("Error updating Version %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Version %q: %#v", d.Id(), res)
	}

	err = DialogflowCXOperationWaitTime(
		config, res, "Updating Version", userAgent, location,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceDialogflowCXVersionRead(d, meta)
}

func resourceDialogflowCXVersionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}{{parent}}/versions/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// extract location from the parent
	location := ""

	if parts := regexp.MustCompile(`locations\/([^\/]*)\/`).FindStringSubmatch(d.Get("parent").(string)); parts != nil {
		location = parts[1]
	} else {
		return fmt.Errorf(
			"Saw %s when the parent is expected to contains location %s",
			d.Get("parent"),
			"projects/{{project}}/locations/{{location}}/...",
		)
	}

	url = strings.Replace(url, "-dialogflow", fmt.Sprintf("%s-dialogflow", location), 1)
	log.Printf("[DEBUG] Deleting Version %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Version")
	}

	err = DialogflowCXOperationWaitTime(
		config, res, "Deleting Version", userAgent, location,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Version %q: %#v", d.Id(), res)
	return nil
}

func resourceDialogflowCXVersionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value and parent contains slashes
	if err := tpgresource.ParseImportId([]string{
		"(?P<parent>.+)/versions/(?P<name>[^/]+)",
		"(?P<parent>.+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/versions/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDialogflowCXVersionName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenDialogflowCXVersionDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXVersionDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXVersionNluSettings(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["model_type"] =
		flattenDialogflowCXVersionNluSettingsModelType(original["modelType"], d, config)
	transformed["classification_threshold"] =
		flattenDialogflowCXVersionNluSettingsClassificationThreshold(original["classificationThreshold"], d, config)
	transformed["model_training_mode"] =
		flattenDialogflowCXVersionNluSettingsModelTrainingMode(original["modelTrainingMode"], d, config)
	return []interface{}{transformed}
}
func flattenDialogflowCXVersionNluSettingsModelType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXVersionNluSettingsClassificationThreshold(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXVersionNluSettingsModelTrainingMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXVersionCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXVersionState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDialogflowCXVersionDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXVersionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
