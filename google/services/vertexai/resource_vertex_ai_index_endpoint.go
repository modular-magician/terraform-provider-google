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

package vertexai

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

func ResourceVertexAIIndexEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceVertexAIIndexEndpointCreate,
		Read:   resourceVertexAIIndexEndpointRead,
		Update: resourceVertexAIIndexEndpointUpdate,
		Delete: resourceVertexAIIndexEndpointDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVertexAIIndexEndpointImport,
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
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The description of the Index.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `The labels with user-defined metadata to organize your Indexes.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"network": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the index endpoint should be peered.
Private services access must already be configured for the network. If left unspecified, the index endpoint is not peered with any network.
[Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): 'projects/{project}/global/networks/{network}'.
Where '{project}' is a project number, as in '12345', and '{network}' is network name.`,
			},
			"public_endpoint_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `If true, the deployed index will be accessible through public endpoint.`,
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The region of the index endpoint. eg us-central1`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp of when the Index was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Used to perform consistent read-modify-write updates.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the Index.`,
			},
			"public_endpoint_domain_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `If publicEndpointEnabled is true, this field will be populated with the domain name to use for this index endpoint.`,
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
				Description: `The timestamp of when the Index was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.`,
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

func resourceVertexAIIndexEndpointCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandVertexAIIndexEndpointDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandVertexAIIndexEndpointDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networkProp, err := expandVertexAIIndexEndpointNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	publicEndpointEnabledProp, err := expandVertexAIIndexEndpointPublicEndpointEnabled(d.Get("public_endpoint_enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("public_endpoint_enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(publicEndpointEnabledProp)) && (ok || !reflect.DeepEqual(v, publicEndpointEnabledProp)) {
		obj["publicEndpointEnabled"] = publicEndpointEnabledProp
	}
	labelsProp, err := expandVertexAIIndexEndpointEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/indexEndpoints")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new IndexEndpoint: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for IndexEndpoint: %s", err)
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
		return fmt.Errorf("Error creating IndexEndpoint: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/indexEndpoints/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = VertexAIOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating IndexEndpoint", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create IndexEndpoint: %s", err)
	}

	if err := d.Set("name", flattenVertexAIIndexEndpointName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/indexEndpoints/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating IndexEndpoint %q: %#v", d.Id(), res)

	return resourceVertexAIIndexEndpointRead(d, meta)
}

func resourceVertexAIIndexEndpointRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/indexEndpoints/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for IndexEndpoint: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("VertexAIIndexEndpoint %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading IndexEndpoint: %s", err)
	}

	if err := d.Set("name", flattenVertexAIIndexEndpointName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading IndexEndpoint: %s", err)
	}
	if err := d.Set("display_name", flattenVertexAIIndexEndpointDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading IndexEndpoint: %s", err)
	}
	if err := d.Set("description", flattenVertexAIIndexEndpointDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading IndexEndpoint: %s", err)
	}
	if err := d.Set("labels", flattenVertexAIIndexEndpointLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading IndexEndpoint: %s", err)
	}
	if err := d.Set("create_time", flattenVertexAIIndexEndpointCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading IndexEndpoint: %s", err)
	}
	if err := d.Set("update_time", flattenVertexAIIndexEndpointUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading IndexEndpoint: %s", err)
	}
	if err := d.Set("network", flattenVertexAIIndexEndpointNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading IndexEndpoint: %s", err)
	}
	if err := d.Set("public_endpoint_domain_name", flattenVertexAIIndexEndpointPublicEndpointDomainName(res["publicEndpointDomainName"], d, config)); err != nil {
		return fmt.Errorf("Error reading IndexEndpoint: %s", err)
	}
	if err := d.Set("terraform_labels", flattenVertexAIIndexEndpointTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading IndexEndpoint: %s", err)
	}
	if err := d.Set("effective_labels", flattenVertexAIIndexEndpointEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading IndexEndpoint: %s", err)
	}

	return nil
}

func resourceVertexAIIndexEndpointUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for IndexEndpoint: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandVertexAIIndexEndpointDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandVertexAIIndexEndpointDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandVertexAIIndexEndpointEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/indexEndpoints/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating IndexEndpoint %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
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
		return fmt.Errorf("Error updating IndexEndpoint %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating IndexEndpoint %q: %#v", d.Id(), res)
	}

	return resourceVertexAIIndexEndpointRead(d, meta)
}

func resourceVertexAIIndexEndpointDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for IndexEndpoint: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/indexEndpoints/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting IndexEndpoint %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "IndexEndpoint")
	}

	err = VertexAIOperationWaitTime(
		config, res, project, "Deleting IndexEndpoint", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting IndexEndpoint %q: %#v", d.Id(), res)
	return nil
}

func resourceVertexAIIndexEndpointImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/indexEndpoints/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{region}}/indexEndpoints/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenVertexAIIndexEndpointName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenVertexAIIndexEndpointDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIIndexEndpointDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIIndexEndpointLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenVertexAIIndexEndpointCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIIndexEndpointUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIIndexEndpointNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIIndexEndpointPublicEndpointDomainName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVertexAIIndexEndpointTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenVertexAIIndexEndpointEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandVertexAIIndexEndpointDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointPublicEndpointEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
