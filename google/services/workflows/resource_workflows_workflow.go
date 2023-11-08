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

package workflows

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceWorkflowsWorkflow() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkflowsWorkflowCreate,
		Read:   resourceWorkflowsWorkflowRead,
		Update: resourceWorkflowsWorkflowUpdate,
		Delete: resourceWorkflowsWorkflowDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		SchemaVersion: 1,

		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceWorkflowsWorkflowResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: ResourceWorkflowsWorkflowUpgradeV0,
				Version: 0,
			},
		},
		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"crypto_key_name": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The KMS key used to encrypt workflow and execution data.

Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}`,
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: `Description of the workflow provided by the user. Must be at most 1000 unicode characters long.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `A set of key/value label pairs to assign to this Workflow.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Name of the Workflow.`,
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of the workflow.`,
			},
			"service_account": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `Name of the service account associated with the latest workflow version. This service
account represents the identity of the workflow and determines what permissions the workflow has.
Format: projects/{project}/serviceAccounts/{account} or {account}.
Using - as a wildcard for the {project} or not providing one at all will infer the project from the account.
The {account} value can be the email address or the unique_id of the service account.
If not provided, workflow will use the project's default service account.
Modifying this field for an existing workflow results in a new workflow revision.`,
			},
			"source_contents": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Workflow code to be executed. The size limit is 32KB.`,
			},
			"user_env_vars": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `User-defined environment variables associated with this workflow revision. This map has a maximum length of 20. Each string can take up to 40KiB. Keys cannot be empty strings and cannot start with “GOOGLE” or “WORKFLOWS".`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the workflow was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"revision_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The revision of the workflow. A new one is generated if the service account or source contents is changed.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `State of the workflow deployment.`,
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
				Description: `The timestamp of when the workflow was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"name_prefix": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"name"},
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

func resourceWorkflowsWorkflowCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandWorkflowsWorkflowName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandWorkflowsWorkflowDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	serviceAccountProp, err := expandWorkflowsWorkflowServiceAccount(d.Get("service_account"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service_account"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceAccountProp)) && (ok || !reflect.DeepEqual(v, serviceAccountProp)) {
		obj["serviceAccount"] = serviceAccountProp
	}
	sourceContentsProp, err := expandWorkflowsWorkflowSourceContents(d.Get("source_contents"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_contents"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceContentsProp)) && (ok || !reflect.DeepEqual(v, sourceContentsProp)) {
		obj["sourceContents"] = sourceContentsProp
	}
	cryptoKeyNameProp, err := expandWorkflowsWorkflowCryptoKeyName(d.Get("crypto_key_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("crypto_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(cryptoKeyNameProp)) && (ok || !reflect.DeepEqual(v, cryptoKeyNameProp)) {
		obj["cryptoKeyName"] = cryptoKeyNameProp
	}
	userEnvVarsProp, err := expandWorkflowsWorkflowUserEnvVars(d.Get("user_env_vars"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_env_vars"); !tpgresource.IsEmptyValue(reflect.ValueOf(userEnvVarsProp)) && (ok || !reflect.DeepEqual(v, userEnvVarsProp)) {
		obj["userEnvVars"] = userEnvVarsProp
	}
	labelsProp, err := expandWorkflowsWorkflowEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	obj, err = resourceWorkflowsWorkflowEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{WorkflowsBasePath}}projects/{{project}}/locations/{{region}}/workflows?workflowId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Workflow: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Workflow: %s", err)
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
		return fmt.Errorf("Error creating Workflow: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/workflows/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = WorkflowsOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Workflow", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Workflow: %s", err)
	}

	if err := d.Set("name", flattenWorkflowsWorkflowName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/workflows/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Workflow %q: %#v", d.Id(), res)

	return resourceWorkflowsWorkflowRead(d, meta)
}

func resourceWorkflowsWorkflowRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{WorkflowsBasePath}}projects/{{project}}/locations/{{region}}/workflows/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Workflow: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("WorkflowsWorkflow %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Workflow: %s", err)
	}

	if err := d.Set("name", flattenWorkflowsWorkflowName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workflow: %s", err)
	}
	if err := d.Set("description", flattenWorkflowsWorkflowDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workflow: %s", err)
	}
	if err := d.Set("create_time", flattenWorkflowsWorkflowCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workflow: %s", err)
	}
	if err := d.Set("update_time", flattenWorkflowsWorkflowUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workflow: %s", err)
	}
	if err := d.Set("state", flattenWorkflowsWorkflowState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workflow: %s", err)
	}
	if err := d.Set("labels", flattenWorkflowsWorkflowLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workflow: %s", err)
	}
	if err := d.Set("service_account", flattenWorkflowsWorkflowServiceAccount(res["serviceAccount"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workflow: %s", err)
	}
	if err := d.Set("source_contents", flattenWorkflowsWorkflowSourceContents(res["sourceContents"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workflow: %s", err)
	}
	if err := d.Set("revision_id", flattenWorkflowsWorkflowRevisionId(res["revisionId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workflow: %s", err)
	}
	if err := d.Set("crypto_key_name", flattenWorkflowsWorkflowCryptoKeyName(res["cryptoKeyName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workflow: %s", err)
	}
	if err := d.Set("user_env_vars", flattenWorkflowsWorkflowUserEnvVars(res["userEnvVars"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workflow: %s", err)
	}
	if err := d.Set("terraform_labels", flattenWorkflowsWorkflowTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workflow: %s", err)
	}
	if err := d.Set("effective_labels", flattenWorkflowsWorkflowEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Workflow: %s", err)
	}

	return nil
}

func resourceWorkflowsWorkflowUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Workflow: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandWorkflowsWorkflowDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	serviceAccountProp, err := expandWorkflowsWorkflowServiceAccount(d.Get("service_account"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service_account"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, serviceAccountProp)) {
		obj["serviceAccount"] = serviceAccountProp
	}
	sourceContentsProp, err := expandWorkflowsWorkflowSourceContents(d.Get("source_contents"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_contents"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sourceContentsProp)) {
		obj["sourceContents"] = sourceContentsProp
	}
	cryptoKeyNameProp, err := expandWorkflowsWorkflowCryptoKeyName(d.Get("crypto_key_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("crypto_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, cryptoKeyNameProp)) {
		obj["cryptoKeyName"] = cryptoKeyNameProp
	}
	userEnvVarsProp, err := expandWorkflowsWorkflowUserEnvVars(d.Get("user_env_vars"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("user_env_vars"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, userEnvVarsProp)) {
		obj["userEnvVars"] = userEnvVarsProp
	}
	labelsProp, err := expandWorkflowsWorkflowEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	obj, err = resourceWorkflowsWorkflowEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{WorkflowsBasePath}}projects/{{project}}/locations/{{region}}/workflows/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Workflow %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("service_account") {
		updateMask = append(updateMask, "serviceAccount")
	}

	if d.HasChange("source_contents") {
		updateMask = append(updateMask, "sourceContents")
	}

	if d.HasChange("crypto_key_name") {
		updateMask = append(updateMask, "cryptoKeyName")
	}

	if d.HasChange("user_env_vars") {
		updateMask = append(updateMask, "userEnvVars")
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
			return fmt.Errorf("Error updating Workflow %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Workflow %q: %#v", d.Id(), res)
		}

		err = WorkflowsOperationWaitTime(
			config, res, project, "Updating Workflow", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceWorkflowsWorkflowRead(d, meta)
}

func resourceWorkflowsWorkflowDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Workflow: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{WorkflowsBasePath}}projects/{{project}}/locations/{{region}}/workflows/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Workflow %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Workflow")
	}

	err = WorkflowsOperationWaitTime(
		config, res, project, "Deleting Workflow", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Workflow %q: %#v", d.Id(), res)
	return nil
}

func flattenWorkflowsWorkflowName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenWorkflowsWorkflowDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkflowsWorkflowCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkflowsWorkflowUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkflowsWorkflowState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkflowsWorkflowLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenWorkflowsWorkflowServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkflowsWorkflowSourceContents(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkflowsWorkflowRevisionId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkflowsWorkflowCryptoKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkflowsWorkflowUserEnvVars(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenWorkflowsWorkflowTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenWorkflowsWorkflowEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandWorkflowsWorkflowName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkflowsWorkflowDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkflowsWorkflowServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkflowsWorkflowSourceContents(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkflowsWorkflowCryptoKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandWorkflowsWorkflowUserEnvVars(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandWorkflowsWorkflowEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func resourceWorkflowsWorkflowEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	var ResName string
	if v, ok := d.GetOk("name"); ok {
		ResName = v.(string)
	} else if v, ok := d.GetOk("name_prefix"); ok {
		ResName = resource.PrefixedUniqueId(v.(string))
	} else {
		ResName = resource.UniqueId()
	}

	if err := d.Set("name", ResName); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}

	return obj, nil
}

func resourceWorkflowsWorkflowResourceV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: `Description of the workflow provided by the user. Must be at most 1000 unicode characters long.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `A set of key/value label pairs to assign to this Workflow.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Name of the Workflow.`,
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of the workflow.`,
			},
			"service_account": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `Name of the service account associated with the latest workflow version. This service
account represents the identity of the workflow and determines what permissions the workflow has.

Format: projects/{project}/serviceAccounts/{account}.`,
			},
			"source_contents": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Workflow code to be executed. The size limit is 32KB.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the workflow was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"revision_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The revision of the workflow. A new one is generated if the service account or source contents is changed.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `State of the workflow deployment.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the workflow was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"name_prefix": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"name"},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func ResourceWorkflowsWorkflowUpgradeV0(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	log.Printf("[DEBUG] Attributes before migration: %#v", rawState)

	rawState["name"] = tpgresource.GetResourceNameFromSelfLink(rawState["name"].(string))

	log.Printf("[DEBUG] Attributes after migration: %#v", rawState)
	return rawState, nil
}
