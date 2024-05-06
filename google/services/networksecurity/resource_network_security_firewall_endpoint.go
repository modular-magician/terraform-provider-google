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
)

func ResourceNetworkSecurityFirewallEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecurityFirewallEndpointCreate,
		Read:   resourceNetworkSecurityFirewallEndpointRead,
		Update: resourceNetworkSecurityFirewallEndpointUpdate,
		Delete: resourceNetworkSecurityFirewallEndpointDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkSecurityFirewallEndpointImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
		),

		Schema: map[string]*schema.Schema{
			"billing_project_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Project to bill on endpoint uptime usage.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location (zone) of the firewall endpoint.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the firewall endpoint resource.`,
			},
			"parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The name of the parent this firewall endpoint belongs to.
Format: organizations/{organization_id}.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `A map of key/value label pairs to assign to the resource.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"associated_networks": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `List of networks that are associated with this endpoint in the local zone.
This is a projection of the FirewallEndpointAssociations pointing at this
endpoint. A network will only appear in this list after traffic routing is
fully configured. Format: projects/{project}/global/networks/{name}.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the firewall endpoint was created in UTC.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"reconciling": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `Whether reconciling is in progress, recommended per https://google.aip.dev/128.`,
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL of this resource.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The current state of the endpoint.`,
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
				Description: `Time the firewall endpoint was updated in UTC.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceNetworkSecurityFirewallEndpointCreate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	billingProjectIdProp, err := expandNetworkSecurityFirewallEndpointBillingProjectId(d.Get("billing_project_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("billing_project_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(billingProjectIdProp)) && (ok || !reflect.DeepEqual(v, billingProjectIdProp)) {
		obj["billingProjectId"] = billingProjectIdProp
	}
	labelsProp, err := expandNetworkSecurityFirewallEndpointEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/firewallEndpoints?firewallEndpointId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FirewallEndpoint: %#v", obj)
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
		return fmt.Errorf("Error creating FirewallEndpoint: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/locations/{{location}}/firewallEndpoints/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Creating FirewallEndpoint", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create FirewallEndpoint: %s", err)
	}

	log.Printf("[DEBUG] Finished creating FirewallEndpoint %q: %#v", d.Id(), res)

	return resourceNetworkSecurityFirewallEndpointRead(d, meta)
}

func resourceNetworkSecurityFirewallEndpointRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/firewallEndpoints/{{name}}")
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkSecurityFirewallEndpoint %q", d.Id()))
	}

	if err := d.Set("labels", flattenNetworkSecurityFirewallEndpointLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallEndpoint: %s", err)
	}
	if err := d.Set("self_link", flattenNetworkSecurityFirewallEndpointSelfLink(res["selfLink"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallEndpoint: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkSecurityFirewallEndpointCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallEndpoint: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkSecurityFirewallEndpointUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallEndpoint: %s", err)
	}
	if err := d.Set("reconciling", flattenNetworkSecurityFirewallEndpointReconciling(res["reconciling"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallEndpoint: %s", err)
	}
	if err := d.Set("associated_networks", flattenNetworkSecurityFirewallEndpointAssociatedNetworks(res["associatedNetworks"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallEndpoint: %s", err)
	}
	if err := d.Set("state", flattenNetworkSecurityFirewallEndpointState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallEndpoint: %s", err)
	}
	if err := d.Set("billing_project_id", flattenNetworkSecurityFirewallEndpointBillingProjectId(res["billingProjectId"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallEndpoint: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkSecurityFirewallEndpointTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallEndpoint: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkSecurityFirewallEndpointEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading FirewallEndpoint: %s", err)
	}

	return nil
}

func resourceNetworkSecurityFirewallEndpointUpdate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	billingProjectIdProp, err := expandNetworkSecurityFirewallEndpointBillingProjectId(d.Get("billing_project_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("billing_project_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, billingProjectIdProp)) {
		obj["billingProjectId"] = billingProjectIdProp
	}
	labelsProp, err := expandNetworkSecurityFirewallEndpointEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/firewallEndpoints/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FirewallEndpoint %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("billing_project_id") {
		updateMask = append(updateMask, "billingProjectId")
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
			return fmt.Errorf("Error updating FirewallEndpoint %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating FirewallEndpoint %q: %#v", d.Id(), res)
		}

		err = NetworkSecurityOperationWaitTime(
			config, res, project, "Updating FirewallEndpoint", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkSecurityFirewallEndpointRead(d, meta)
}

func resourceNetworkSecurityFirewallEndpointDelete(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/firewallEndpoints/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting FirewallEndpoint %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "FirewallEndpoint")
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Deleting FirewallEndpoint", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting FirewallEndpoint %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkSecurityFirewallEndpointImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<parent>.+)/locations/(?P<location>[^/]+)/firewallEndpoints/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/locations/{{location}}/firewallEndpoints/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkSecurityFirewallEndpointLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkSecurityFirewallEndpointSelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityFirewallEndpointCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityFirewallEndpointUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityFirewallEndpointReconciling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityFirewallEndpointAssociatedNetworks(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityFirewallEndpointState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityFirewallEndpointBillingProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityFirewallEndpointTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkSecurityFirewallEndpointEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkSecurityFirewallEndpointBillingProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityFirewallEndpointEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
