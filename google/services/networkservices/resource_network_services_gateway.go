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

package networkservices

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
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

// Checks if there is another gateway under the same location.
func gatewaysSameLocation(d *schema.ResourceData, config *transport_tpg.Config, billingProject, userAgent string) ([]interface{}, error) {
	log.Print("[DEBUG] Looking for gateways under the same location.")
	var gateways []interface{}

	gatewaysUrl, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/{{location}}/gateways")
	if err != nil {
		return gateways, err
	}

	resp, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    gatewaysUrl,
		UserAgent: userAgent,
	})
	if err != nil {
		return gateways, err
	}

	data, ok := resp["gateways"]
	if !ok || data == nil {
		log.Print("[DEBUG] No gateways under the same location found.")
		return gateways, nil
	}

	gateways = data.([]interface{})

	log.Printf("[DEBUG] There are still gateways under the same location: %#v", gateways)

	return gateways, nil
}

// Checks if the given list of gateways contains a gateway of type SECURE_WEB_GATEWAY.
func isLastSWGGateway(gateways []interface{}, network string) bool {
	log.Print("[DEBUG] Checking if this is the last gateway of type SECURE_WEB_GATEWAY.")
	for _, itemRaw := range gateways {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		gType, ok := item["type"]
		if !ok || gType == nil {
			continue
		}

		gNetwork, ok := item["network"]
		if !ok || gNetwork == nil {
			continue
		}

		if gType.(string) == "SECURE_WEB_GATEWAY" && gNetwork.(string) == network {
			return false
		}
	}

	log.Print("[DEBUG] There is no other gateway of type SECURE_WEB_GATEWAY.")
	// no gateways of type SWG found.
	return true
}

// Deletes the swg-autogen-router if the current gateway being deleted is the type of swg so there is no other gateway using it.
func deleteSWGAutoGenRouter(d *schema.ResourceData, config *transport_tpg.Config, billingProject, userAgent string) error {
	log.Printf("[DEBUG] Searching the network id by name %q.", d.Get("network"))

	networkPath := fmt.Sprintf("{{ComputeBasePath}}%s", d.Get("network"))
	networkUrl, err := tpgresource.ReplaceVars(d, config, networkPath)
	if err != nil {
		return err
	}

	resp, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    networkUrl,
		UserAgent: userAgent,
	})
	if err != nil {
		return err
	}

	// The name of swg auto generated router is in the following format: swg-autogen-router-{NETWORK-ID}
	routerId := fmt.Sprintf("swg-autogen-router-%s", resp["id"])
	log.Printf("[DEBUG] Deleting the auto generated router %q.", routerId)

	routerPath := fmt.Sprintf("{{ComputeBasePath}}projects/{{project}}/regions/{{location}}/routers/%s", routerId)
	routerUrl, err := tpgresource.ReplaceVars(d, config, routerPath)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "DELETE",
		Project:              billingProject,
		RawURL:               routerUrl,
		UserAgent:            userAgent,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsSwgAutogenRouterRetryable},
	})
	if err != nil {
		if transport_tpg.IsGoogleApiErrorWithCode(err, 404) {
			// The swg auto gen router may have already been deleted.
			// No further action needed.
			return nil
		}

		return err
	}

	return nil
}

func ResourceNetworkServicesGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkServicesGatewayCreate,
		Read:   resourceNetworkServicesGatewayRead,
		Update: resourceNetworkServicesGatewayUpdate,
		Delete: resourceNetworkServicesGatewayDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkServicesGatewayImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Short name of the Gateway resource to be created.`,
			},
			"ports": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Description: `One or more port numbers (1-65535), on which the Gateway will receive traffic.
The proxy binds to the specified ports. Gateways of type 'SECURE_WEB_GATEWAY' are
limited to 1 port. Gateways of type 'OPEN_MESH' listen on 0.0.0.0 and support multiple ports.`,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"TYPE_UNSPECIFIED", "OPEN_MESH", "SECURE_WEB_GATEWAY"}),
				Description:  `Immutable. The type of the customer-managed gateway. Possible values are: * OPEN_MESH * SECURE_WEB_GATEWAY. Possible values: ["TYPE_UNSPECIFIED", "OPEN_MESH", "SECURE_WEB_GATEWAY"]`,
			},
			"addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `Zero or one IPv4-address on which the Gateway will receive the traffic. When no address is provided,
an IP from the subnetwork is allocated This field only applies to gateways of type 'SECURE_WEB_GATEWAY'.
Gateways of type 'OPEN_MESH' listen on 0.0.0.0.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"certificate_urls": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A fully-qualified Certificates URL reference. The proxy presents a Certificate (selected based on SNI) when establishing a TLS connection.
This feature only applies to gateways of type 'SECURE_WEB_GATEWAY'.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A free-text description of the resource. Max length 1024 characters.`,
			},
			"gateway_security_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `A fully-qualified GatewaySecurityPolicy URL reference. Defines how a server should apply security policy to inbound (VM to Proxy) initiated connections.
For example: 'projects/*/locations/*/gatewaySecurityPolicies/swg-policy'.
This policy is specific to gateways of type 'SECURE_WEB_GATEWAY'.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Set of label tags associated with the Gateway resource.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The location of the gateway.
The default value is 'global'.`,
				Default: "global",
			},
			"network": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The relative resource name identifying the VPC network that is using this configuration.
For example: 'projects/*/global/networks/network-1'.
Currently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY'.`,
			},
			"scope": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Immutable. Scope determines how configuration across multiple Gateway instances are merged.
The configuration for multiple Gateway instances with the same scope will be merged as presented as
a single coniguration to the proxy/load balancer.
Max length 64 characters. Scope should start with a letter and can only have letters, numbers, hyphens.`,
			},
			"server_tls_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `A fully-qualified ServerTLSPolicy URL reference. Specifies how TLS traffic is terminated.
If empty, TLS termination is disabled.`,
			},
			"subnetwork": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The relative resource name identifying the subnetwork in which this SWG is allocated.
For example: 'projects/*/regions/us-central1/subnetworks/network-1'.
Currently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the AccessPolicy was created in UTC.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"self_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL of this resource.`,
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
				Description: `Time the AccessPolicy was updated in UTC.`,
			},
			"delete_swg_autogen_router_on_destroy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				Description: `UPDATEDWhen deleting a gateway of type 'SECURE_WEB_GATEWAY', this boolean option will also delete auto generated router by the gateway creation.
If there is no other gateway of type 'SECURE_WEB_GATEWAY' remaining for that region and network it will be deleted.`,
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

func resourceNetworkServicesGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesGatewayDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	typeProp, err := expandNetworkServicesGatewayType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	portsProp, err := expandNetworkServicesGatewayPorts(d.Get("ports"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ports"); !tpgresource.IsEmptyValue(reflect.ValueOf(portsProp)) && (ok || !reflect.DeepEqual(v, portsProp)) {
		obj["ports"] = portsProp
	}
	scopeProp, err := expandNetworkServicesGatewayScope(d.Get("scope"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("scope"); !tpgresource.IsEmptyValue(reflect.ValueOf(scopeProp)) && (ok || !reflect.DeepEqual(v, scopeProp)) {
		obj["scope"] = scopeProp
	}
	serverTlsPolicyProp, err := expandNetworkServicesGatewayServerTlsPolicy(d.Get("server_tls_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("server_tls_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(serverTlsPolicyProp)) && (ok || !reflect.DeepEqual(v, serverTlsPolicyProp)) {
		obj["serverTlsPolicy"] = serverTlsPolicyProp
	}
	addressesProp, err := expandNetworkServicesGatewayAddresses(d.Get("addresses"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("addresses"); !tpgresource.IsEmptyValue(reflect.ValueOf(addressesProp)) && (ok || !reflect.DeepEqual(v, addressesProp)) {
		obj["addresses"] = addressesProp
	}
	subnetworkProp, err := expandNetworkServicesGatewaySubnetwork(d.Get("subnetwork"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("subnetwork"); !tpgresource.IsEmptyValue(reflect.ValueOf(subnetworkProp)) && (ok || !reflect.DeepEqual(v, subnetworkProp)) {
		obj["subnetwork"] = subnetworkProp
	}
	networkProp, err := expandNetworkServicesGatewayNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	gatewaySecurityPolicyProp, err := expandNetworkServicesGatewayGatewaySecurityPolicy(d.Get("gateway_security_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gateway_security_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(gatewaySecurityPolicyProp)) && (ok || !reflect.DeepEqual(v, gatewaySecurityPolicyProp)) {
		obj["gatewaySecurityPolicy"] = gatewaySecurityPolicyProp
	}
	certificateUrlsProp, err := expandNetworkServicesGatewayCertificateUrls(d.Get("certificate_urls"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("certificate_urls"); !tpgresource.IsEmptyValue(reflect.ValueOf(certificateUrlsProp)) && (ok || !reflect.DeepEqual(v, certificateUrlsProp)) {
		obj["certificateUrls"] = certificateUrlsProp
	}
	labelsProp, err := expandNetworkServicesGatewayEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/{{location}}/gateways?gatewayId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Gateway: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Gateway: %s", err)
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
		return fmt.Errorf("Error creating Gateway: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/gateways/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkServicesOperationWaitTime(
		config, res, project, "Creating Gateway", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Gateway: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Gateway %q: %#v", d.Id(), res)

	return resourceNetworkServicesGatewayRead(d, meta)
}

func resourceNetworkServicesGatewayRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/{{location}}/gateways/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Gateway: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkServicesGateway %q", d.Id()))
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("delete_swg_autogen_router_on_destroy"); !ok {
		if err := d.Set("delete_swg_autogen_router_on_destroy", false); err != nil {
			return fmt.Errorf("Error setting delete_swg_autogen_router_on_destroy: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}

	if err := d.Set("self_link", flattenNetworkServicesGatewaySelfLink(res["selfLink"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkServicesGatewayCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkServicesGatewayUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("labels", flattenNetworkServicesGatewayLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("description", flattenNetworkServicesGatewayDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("type", flattenNetworkServicesGatewayType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("ports", flattenNetworkServicesGatewayPorts(res["ports"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("scope", flattenNetworkServicesGatewayScope(res["scope"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("server_tls_policy", flattenNetworkServicesGatewayServerTlsPolicy(res["serverTlsPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("addresses", flattenNetworkServicesGatewayAddresses(res["addresses"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("subnetwork", flattenNetworkServicesGatewaySubnetwork(res["subnetwork"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("network", flattenNetworkServicesGatewayNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("gateway_security_policy", flattenNetworkServicesGatewayGatewaySecurityPolicy(res["gatewaySecurityPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("certificate_urls", flattenNetworkServicesGatewayCertificateUrls(res["certificateUrls"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkServicesGatewayTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkServicesGatewayEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Gateway: %s", err)
	}

	return nil
}

func resourceNetworkServicesGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Gateway: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesGatewayDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	serverTlsPolicyProp, err := expandNetworkServicesGatewayServerTlsPolicy(d.Get("server_tls_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("server_tls_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, serverTlsPolicyProp)) {
		obj["serverTlsPolicy"] = serverTlsPolicyProp
	}
	gatewaySecurityPolicyProp, err := expandNetworkServicesGatewayGatewaySecurityPolicy(d.Get("gateway_security_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gateway_security_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, gatewaySecurityPolicyProp)) {
		obj["gatewaySecurityPolicy"] = gatewaySecurityPolicyProp
	}
	certificateUrlsProp, err := expandNetworkServicesGatewayCertificateUrls(d.Get("certificate_urls"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("certificate_urls"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, certificateUrlsProp)) {
		obj["certificateUrls"] = certificateUrlsProp
	}
	labelsProp, err := expandNetworkServicesGatewayEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/{{location}}/gateways/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Gateway %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("server_tls_policy") {
		updateMask = append(updateMask, "serverTlsPolicy")
	}

	if d.HasChange("gateway_security_policy") {
		updateMask = append(updateMask, "gatewaySecurityPolicy")
	}

	if d.HasChange("certificate_urls") {
		updateMask = append(updateMask, "certificateUrls")
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
	if d.Get("type") == "SECURE_WEB_GATEWAY" {
		obj["name"] = d.Get("name")
		obj["type"] = d.Get("type")
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
			return fmt.Errorf("Error updating Gateway %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Gateway %q: %#v", d.Id(), res)
		}

		err = NetworkServicesOperationWaitTime(
			config, res, project, "Updating Gateway", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkServicesGatewayRead(d, meta)
}

func resourceNetworkServicesGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Gateway: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkServicesBasePath}}projects/{{project}}/locations/{{location}}/gateways/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Gateway %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "Gateway")
	}

	err = NetworkServicesOperationWaitTime(
		config, res, project, "Deleting Gateway", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}
	if d.Get("delete_swg_autogen_router_on_destroy").(bool) {
		log.Print("[DEBUG] The field delete_swg_autogen_router_on_destroy is true. Deleting swg_autogen_router.")
		gateways, err := gatewaysSameLocation(d, config, billingProject, userAgent)
		if err != nil {
			return err
		}

		network := d.Get("network").(string)
		if isLastSWGGateway(gateways, network) {
			err := deleteSWGAutoGenRouter(d, config, billingProject, userAgent)
			if err != nil {
				return err
			}
		}
	}

	log.Printf("[DEBUG] Finished deleting Gateway %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkServicesGatewayImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/gateways/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/gateways/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	if err := d.Set("delete_swg_autogen_router_on_destroy", false); err != nil {
		return nil, fmt.Errorf("Error setting delete_swg_autogen_router_on_destroy: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkServicesGatewaySelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesGatewayCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesGatewayUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesGatewayLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkServicesGatewayDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesGatewayType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesGatewayPorts(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesGatewayScope(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesGatewayServerTlsPolicy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesGatewayAddresses(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesGatewaySubnetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesGatewayNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesGatewayGatewaySecurityPolicy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesGatewayCertificateUrls(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkServicesGatewayTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkServicesGatewayEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkServicesGatewayDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayPorts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayServerTlsPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayAddresses(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewaySubnetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayGatewaySecurityPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayCertificateUrls(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesGatewayEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
