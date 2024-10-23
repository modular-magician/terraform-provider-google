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

func ResourceNetworkSecurityGatewaySecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkSecurityGatewaySecurityPolicyCreate,
		Read:   resourceNetworkSecurityGatewaySecurityPolicyRead,
		Update: resourceNetworkSecurityGatewaySecurityPolicyUpdate,
		Delete: resourceNetworkSecurityGatewaySecurityPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkSecurityGatewaySecurityPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Description: `Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gatewaySecurityPolicy}
gatewaySecurityPolicy should match the pattern:(^a-z?$).`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A free-text description of the resource. Max length 1024 characters.`,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The location of the gateway security policy.
The default value is 'global'.`,
				Default: "global",
			},
			"tls_inspection_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Name of a TlsInspectionPolicy resource that defines how TLS inspection is performed for any rule that enables it.
Note: google_network_security_tls_inspection_policy resource is still in [Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html) therefore it will need to import the provider.`,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The timestamp when the resource was created.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"`,
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL of this resource.`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The timestamp when the resource was updated.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
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

func resourceNetworkSecurityGatewaySecurityPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityGatewaySecurityPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	tlsInspectionPolicyProp, err := expandNetworkSecurityGatewaySecurityPolicyTlsInspectionPolicy(d.Get("tls_inspection_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tls_inspection_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(tlsInspectionPolicyProp)) && (ok || !reflect.DeepEqual(v, tlsInspectionPolicyProp)) {
		obj["tlsInspectionPolicy"] = tlsInspectionPolicyProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies?gatewaySecurityPolicyId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GatewaySecurityPolicy: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GatewaySecurityPolicy: %s", err)
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
		return fmt.Errorf("Error creating GatewaySecurityPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Creating GatewaySecurityPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create GatewaySecurityPolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating GatewaySecurityPolicy %q: %#v", d.Id(), res)

	return resourceNetworkSecurityGatewaySecurityPolicyRead(d, meta)
}

func resourceNetworkSecurityGatewaySecurityPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GatewaySecurityPolicy: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkSecurityGatewaySecurityPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicy: %s", err)
	}

	if err := d.Set("self_link", flattenNetworkSecurityGatewaySecurityPolicySelfLink(res["selfLink"], d, config)); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicy: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkSecurityGatewaySecurityPolicyCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicy: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkSecurityGatewaySecurityPolicyUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicy: %s", err)
	}
	if err := d.Set("description", flattenNetworkSecurityGatewaySecurityPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading GatewaySecurityPolicy: %s", err)
	}

	return nil
}

func resourceNetworkSecurityGatewaySecurityPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GatewaySecurityPolicy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityGatewaySecurityPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	tlsInspectionPolicyProp, err := expandNetworkSecurityGatewaySecurityPolicyTlsInspectionPolicy(d.Get("tls_inspection_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tls_inspection_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, tlsInspectionPolicyProp)) {
		obj["tlsInspectionPolicy"] = tlsInspectionPolicyProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating GatewaySecurityPolicy %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("tls_inspection_policy") {
		updateMask = append(updateMask, "tlsInspectionPolicy")
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
			return fmt.Errorf("Error updating GatewaySecurityPolicy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating GatewaySecurityPolicy %q: %#v", d.Id(), res)
		}

		err = NetworkSecurityOperationWaitTime(
			config, res, project, "Updating GatewaySecurityPolicy", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkSecurityGatewaySecurityPolicyRead(d, meta)
}

func resourceNetworkSecurityGatewaySecurityPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GatewaySecurityPolicy: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkSecurityBasePath}}projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting GatewaySecurityPolicy %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "GatewaySecurityPolicy")
	}

	err = NetworkSecurityOperationWaitTime(
		config, res, project, "Deleting GatewaySecurityPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting GatewaySecurityPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkSecurityGatewaySecurityPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/gatewaySecurityPolicies/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkSecurityGatewaySecurityPolicySelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityGatewaySecurityPolicyCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityGatewaySecurityPolicyUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkSecurityGatewaySecurityPolicyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkSecurityGatewaySecurityPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityGatewaySecurityPolicyTlsInspectionPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
