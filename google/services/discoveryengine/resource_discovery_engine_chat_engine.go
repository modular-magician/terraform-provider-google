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

package discoveryengine

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
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceDiscoveryEngineChatEngine() *schema.Resource {
	return &schema.Resource{
		Create: resourceDiscoveryEngineChatEngineCreate,
		Read:   resourceDiscoveryEngineChatEngineRead,
		Update: resourceDiscoveryEngineChatEngineUpdate,
		Delete: resourceDiscoveryEngineChatEngineDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDiscoveryEngineChatEngineImport,
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
			"chat_engine_config": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: `Configurations for a chat Engine.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agent_creation_config": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `The configuration to generate the Dialogflow agent that is associated to this Engine.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default_language_code": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The default language of the agent as a language tag. See [Language Support](https://cloud.google.com/dialogflow/docs/reference/language) for a list of the currently supported language codes.`,
									},
									"time_zone": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The time zone of the agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris.`,
									},
									"business": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Name of the company, organization or other entity that the agent represents. Used for knowledge connector LLM prompt and for knowledge search.`,
									},
									"location": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Agent location for Agent creation, currently supported values: global/us/eu, it needs to be the same region as the Chat Engine.`,
									},
								},
							},
						},
					},
				},
			},
			"collection_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The collection ID.`,
			},
			"data_store_ids": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: `The data stores associated with this engine. Multiple DataStores in the same Collection can be associated here. All listed DataStores must be 'SOLUTION_TYPE_CHAT'. Adding or removing data stores will force recreation.`,
				MinItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The display name of the engine. Should be human readable. UTF-8 encoded string with limit of 1024 characters.`,
			},
			"engine_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID to use for chat engine.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Location.`,
			},
			"common_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Common config spec that specifies the metadata of the engine.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"company_name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `The name of the company, business or entity that is associated with the engine. Setting this may help improve LLM related features.`,
						},
					},
				},
			},
			"industry_vertical": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"GENERIC", ""}),
				Description:  `The industry vertical that the chat engine registers. Vertical on Engine has to match vertical of the DataStore linked to the engine. Default value: "GENERIC" Possible values: ["GENERIC"]`,
				Default:      "GENERIC",
			},
			"chat_engine_metadata": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Additional information of the Chat Engine.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dialogflow_agent": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The resource name of a Dialogflow agent, that this Chat Engine refers to.`,
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Timestamp the Engine was created at.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The unique full resource name of the chat engine. Values are of the format
'projects/{project}/locations/{location}/collections/{collection_id}/engines/{engine_id}'.
This field must be a UTF-8 encoded string with a length limit of 1024
characters.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Timestamp the Engine was last updated.`,
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

func resourceDiscoveryEngineChatEngineCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	industryVerticalProp, err := expandDiscoveryEngineChatEngineIndustryVertical(d.Get("industry_vertical"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("industry_vertical"); !tpgresource.IsEmptyValue(reflect.ValueOf(industryVerticalProp)) && (ok || !reflect.DeepEqual(v, industryVerticalProp)) {
		obj["industryVertical"] = industryVerticalProp
	}
	displayNameProp, err := expandDiscoveryEngineChatEngineDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	dataStoreIdsProp, err := expandDiscoveryEngineChatEngineDataStoreIds(d.Get("data_store_ids"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("data_store_ids"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataStoreIdsProp)) && (ok || !reflect.DeepEqual(v, dataStoreIdsProp)) {
		obj["dataStoreIds"] = dataStoreIdsProp
	}
	chatEngineConfigProp, err := expandDiscoveryEngineChatEngineChatEngineConfig(d.Get("chat_engine_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("chat_engine_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(chatEngineConfigProp)) && (ok || !reflect.DeepEqual(v, chatEngineConfigProp)) {
		obj["chatEngineConfig"] = chatEngineConfigProp
	}
	commonConfigProp, err := expandDiscoveryEngineChatEngineCommonConfig(d.Get("common_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("common_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(commonConfigProp)) && (ok || !reflect.DeepEqual(v, commonConfigProp)) {
		obj["commonConfig"] = commonConfigProp
	}

	obj, err = resourceDiscoveryEngineChatEngineEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines?engineId={{engine_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ChatEngine: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ChatEngine: %s", err)
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
		return fmt.Errorf("Error creating ChatEngine: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = DiscoveryEngineOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating ChatEngine", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create ChatEngine: %s", err)
	}

	if err := d.Set("name", flattenDiscoveryEngineChatEngineName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ChatEngine %q: %#v", d.Id(), res)

	return resourceDiscoveryEngineChatEngineRead(d, meta)
}

func resourceDiscoveryEngineChatEngineRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ChatEngine: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DiscoveryEngineChatEngine %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ChatEngine: %s", err)
	}

	if err := d.Set("name", flattenDiscoveryEngineChatEngineName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ChatEngine: %s", err)
	}
	if err := d.Set("industry_vertical", flattenDiscoveryEngineChatEngineIndustryVertical(res["industryVertical"], d, config)); err != nil {
		return fmt.Errorf("Error reading ChatEngine: %s", err)
	}
	if err := d.Set("display_name", flattenDiscoveryEngineChatEngineDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading ChatEngine: %s", err)
	}
	if err := d.Set("data_store_ids", flattenDiscoveryEngineChatEngineDataStoreIds(res["dataStoreIds"], d, config)); err != nil {
		return fmt.Errorf("Error reading ChatEngine: %s", err)
	}
	if err := d.Set("create_time", flattenDiscoveryEngineChatEngineCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ChatEngine: %s", err)
	}
	if err := d.Set("update_time", flattenDiscoveryEngineChatEngineUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ChatEngine: %s", err)
	}
	if err := d.Set("common_config", flattenDiscoveryEngineChatEngineCommonConfig(res["commonConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading ChatEngine: %s", err)
	}
	if err := d.Set("chat_engine_metadata", flattenDiscoveryEngineChatEngineChatEngineMetadata(res["chatEngineMetadata"], d, config)); err != nil {
		return fmt.Errorf("Error reading ChatEngine: %s", err)
	}

	return nil
}

func resourceDiscoveryEngineChatEngineUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ChatEngine: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandDiscoveryEngineChatEngineDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	obj, err = resourceDiscoveryEngineChatEngineEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ChatEngine %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
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
			return fmt.Errorf("Error updating ChatEngine %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ChatEngine %q: %#v", d.Id(), res)
		}

	}

	return resourceDiscoveryEngineChatEngineRead(d, meta)
}

func resourceDiscoveryEngineChatEngineDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ChatEngine: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DiscoveryEngineBasePath}}projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	log.Printf("[DEBUG] Deleting ChatEngine %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "ChatEngine")
	}

	err = DiscoveryEngineOperationWaitTime(
		config, res, project, "Deleting ChatEngine", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ChatEngine %q: %#v", d.Id(), res)
	return nil
}

func resourceDiscoveryEngineChatEngineImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/collections/(?P<collection_id>[^/]+)/engines/(?P<engine_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<collection_id>[^/]+)/(?P<engine_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<collection_id>[^/]+)/(?P<engine_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDiscoveryEngineChatEngineName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineChatEngineIndustryVertical(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineChatEngineDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineChatEngineDataStoreIds(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineChatEngineCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineChatEngineUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineChatEngineCommonConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["company_name"] =
		flattenDiscoveryEngineChatEngineCommonConfigCompanyName(original["companyName"], d, config)
	return []interface{}{transformed}
}
func flattenDiscoveryEngineChatEngineCommonConfigCompanyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDiscoveryEngineChatEngineChatEngineMetadata(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dialogflow_agent"] =
		flattenDiscoveryEngineChatEngineChatEngineMetadataDialogflowAgent(original["dialogflowAgent"], d, config)
	return []interface{}{transformed}
}
func flattenDiscoveryEngineChatEngineChatEngineMetadataDialogflowAgent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDiscoveryEngineChatEngineIndustryVertical(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineChatEngineDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineChatEngineDataStoreIds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineChatEngineChatEngineConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAgentCreationConfig, err := expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfig(original["agent_creation_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAgentCreationConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["agentCreationConfig"] = transformedAgentCreationConfig
	}

	return transformed, nil
}

func expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBusiness, err := expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigBusiness(original["business"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBusiness); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["business"] = transformedBusiness
	}

	transformedDefaultLanguageCode, err := expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigDefaultLanguageCode(original["default_language_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultLanguageCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultLanguageCode"] = transformedDefaultLanguageCode
	}

	transformedTimeZone, err := expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigTimeZone(original["time_zone"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeZone); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["timeZone"] = transformedTimeZone
	}

	transformedLocation, err := expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigBusiness(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigDefaultLanguageCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineChatEngineCommonConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCompanyName, err := expandDiscoveryEngineChatEngineCommonConfigCompanyName(original["company_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCompanyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["companyName"] = transformedCompanyName
	}

	return transformed, nil
}

func expandDiscoveryEngineChatEngineCommonConfigCompanyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceDiscoveryEngineChatEngineEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	//hard code solutionType to "SOLUTION_TYPE_CHAT" for chat engine resource
	obj["solutionType"] = "SOLUTION_TYPE_CHAT"
	return obj, nil
}
