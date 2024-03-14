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

package edgecontainer

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceEdgecontainerVpnConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceEdgecontainerVpnConnectionCreate,
		Read:   resourceEdgecontainerVpnConnectionRead,
		Update: resourceEdgecontainerVpnConnectionUpdate,
		Delete: resourceEdgecontainerVpnConnectionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceEdgecontainerVpnConnectionImport,
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
			"cluster": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The canonical Cluster name to connect to. It is in the form of projects/{project}/locations/{location}/clusters/{cluster}.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Google Cloud Platform location.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource name of VPN connection`,
			},
			"enable_high_availability": {
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Whether this VPN connection has HA enabled on cluster side. If enabled, when creating VPN connection we will attempt to use 2 ANG floating IPs.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels associated with this resource.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"nat_gateway_ip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `NAT gateway IP, or WAN IP address. If a customer has multiple NAT IPs, the customer needs to configure NAT such that only one external IP maps to the GMEC Anthos cluster.
This is empty if NAT is not used.`,
			},
			"router": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The VPN connection Cloud Router name.`,
			},
			"vpc": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The network ID of VPC to connect to.`,
			},
			"vpc_project": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Project detail of the VPC network. Required if VPC is in a different project than the cluster project.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_id": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `The project of the VPC to connect to. If not specified, it is the same as the cluster project.`,
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time when the VPN connection was created.`,
			},
			"details": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `A nested object resource`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloud_router": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `The Cloud Router info.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `The associated Cloud Router name.`,
									},
								},
							},
						},
						"cloud_vpns": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `Each connection has multiple Cloud VPN gateways.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gateway": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `The created Cloud VPN gateway name.`,
									},
								},
							},
						},
						"error": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The error message. This is only populated when state=ERROR.`,
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The current connection state.`,
						},
					},
				},
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				ForceNew:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
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
				Description: `The time when the VPN connection was last updated.`,
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

func resourceEdgecontainerVpnConnectionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	natGatewayIpProp, err := expandEdgecontainerVpnConnectionNatGatewayIp(d.Get("nat_gateway_ip"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("nat_gateway_ip"); !tpgresource.IsEmptyValue(reflect.ValueOf(natGatewayIpProp)) && (ok || !reflect.DeepEqual(v, natGatewayIpProp)) {
		obj["natGatewayIp"] = natGatewayIpProp
	}
	clusterProp, err := expandEdgecontainerVpnConnectionCluster(d.Get("cluster"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cluster"); !tpgresource.IsEmptyValue(reflect.ValueOf(clusterProp)) && (ok || !reflect.DeepEqual(v, clusterProp)) {
		obj["cluster"] = clusterProp
	}
	vpcProp, err := expandEdgecontainerVpnConnectionVpc(d.Get("vpc"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vpc"); !tpgresource.IsEmptyValue(reflect.ValueOf(vpcProp)) && (ok || !reflect.DeepEqual(v, vpcProp)) {
		obj["vpc"] = vpcProp
	}
	vpcProjectProp, err := expandEdgecontainerVpnConnectionVpcProject(d.Get("vpc_project"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vpc_project"); !tpgresource.IsEmptyValue(reflect.ValueOf(vpcProjectProp)) && (ok || !reflect.DeepEqual(v, vpcProjectProp)) {
		obj["vpcProject"] = vpcProjectProp
	}
	enableHighAvailabilityProp, err := expandEdgecontainerVpnConnectionEnableHighAvailability(d.Get("enable_high_availability"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_high_availability"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableHighAvailabilityProp)) && (ok || !reflect.DeepEqual(v, enableHighAvailabilityProp)) {
		obj["enableHighAvailability"] = enableHighAvailabilityProp
	}
	routerProp, err := expandEdgecontainerVpnConnectionRouter(d.Get("router"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("router"); !tpgresource.IsEmptyValue(reflect.ValueOf(routerProp)) && (ok || !reflect.DeepEqual(v, routerProp)) {
		obj["router"] = routerProp
	}
	labelsProp, err := expandEdgecontainerVpnConnectionEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{EdgecontainerBasePath}}projects/{{project}}/locations/{{location}}/vpnConnections?vpnConnectionId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new VpnConnection: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for VpnConnection: %s", err)
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
		return fmt.Errorf("Error creating VpnConnection: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/vpnConnections/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = EdgecontainerOperationWaitTime(
		config, res, project, "Creating VpnConnection", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create VpnConnection: %s", err)
	}

	log.Printf("[DEBUG] Finished creating VpnConnection %q: %#v", d.Id(), res)

	return resourceEdgecontainerVpnConnectionRead(d, meta)
}

func resourceEdgecontainerVpnConnectionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{EdgecontainerBasePath}}projects/{{project}}/locations/{{location}}/vpnConnections/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for VpnConnection: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("EdgecontainerVpnConnection %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading VpnConnection: %s", err)
	}

	if err := d.Set("create_time", flattenEdgecontainerVpnConnectionCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnConnection: %s", err)
	}
	if err := d.Set("update_time", flattenEdgecontainerVpnConnectionUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnConnection: %s", err)
	}
	if err := d.Set("labels", flattenEdgecontainerVpnConnectionLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnConnection: %s", err)
	}
	if err := d.Set("nat_gateway_ip", flattenEdgecontainerVpnConnectionNatGatewayIp(res["natGatewayIp"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnConnection: %s", err)
	}
	if err := d.Set("cluster", flattenEdgecontainerVpnConnectionCluster(res["cluster"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnConnection: %s", err)
	}
	if err := d.Set("vpc", flattenEdgecontainerVpnConnectionVpc(res["vpc"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnConnection: %s", err)
	}
	if err := d.Set("vpc_project", flattenEdgecontainerVpnConnectionVpcProject(res["vpcProject"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnConnection: %s", err)
	}
	if err := d.Set("enable_high_availability", flattenEdgecontainerVpnConnectionEnableHighAvailability(res["enableHighAvailability"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnConnection: %s", err)
	}
	if err := d.Set("router", flattenEdgecontainerVpnConnectionRouter(res["router"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnConnection: %s", err)
	}
	if err := d.Set("details", flattenEdgecontainerVpnConnectionDetails(res["details"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnConnection: %s", err)
	}
	if err := d.Set("terraform_labels", flattenEdgecontainerVpnConnectionTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnConnection: %s", err)
	}
	if err := d.Set("effective_labels", flattenEdgecontainerVpnConnectionEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading VpnConnection: %s", err)
	}

	return nil
}

func resourceEdgecontainerVpnConnectionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for VpnConnection: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	natGatewayIpProp, err := expandEdgecontainerVpnConnectionNatGatewayIp(d.Get("nat_gateway_ip"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("nat_gateway_ip"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, natGatewayIpProp)) {
		obj["natGatewayIp"] = natGatewayIpProp
	}
	clusterProp, err := expandEdgecontainerVpnConnectionCluster(d.Get("cluster"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cluster"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clusterProp)) {
		obj["cluster"] = clusterProp
	}
	vpcProp, err := expandEdgecontainerVpnConnectionVpc(d.Get("vpc"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vpc"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, vpcProp)) {
		obj["vpc"] = vpcProp
	}
	vpcProjectProp, err := expandEdgecontainerVpnConnectionVpcProject(d.Get("vpc_project"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vpc_project"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, vpcProjectProp)) {
		obj["vpcProject"] = vpcProjectProp
	}
	enableHighAvailabilityProp, err := expandEdgecontainerVpnConnectionEnableHighAvailability(d.Get("enable_high_availability"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_high_availability"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enableHighAvailabilityProp)) {
		obj["enableHighAvailability"] = enableHighAvailabilityProp
	}
	routerProp, err := expandEdgecontainerVpnConnectionRouter(d.Get("router"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("router"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, routerProp)) {
		obj["router"] = routerProp
	}
	labelsProp, err := expandEdgecontainerVpnConnectionEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{EdgecontainerBasePath}}projects/{{project}}/locations/{{location}}/vpnConnections/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating VpnConnection %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating VpnConnection %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating VpnConnection %q: %#v", d.Id(), res)
	}

	err = EdgecontainerOperationWaitTime(
		config, res, project, "Updating VpnConnection", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceEdgecontainerVpnConnectionRead(d, meta)
}

func resourceEdgecontainerVpnConnectionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for VpnConnection: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{EdgecontainerBasePath}}projects/{{project}}/locations/{{location}}/vpnConnections/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	log.Printf("[DEBUG] Deleting VpnConnection %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "VpnConnection")
	}

	err = EdgecontainerOperationWaitTime(
		config, res, project, "Deleting VpnConnection", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting VpnConnection %q: %#v", d.Id(), res)
	return nil
}

func resourceEdgecontainerVpnConnectionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/vpnConnections/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/vpnConnections/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenEdgecontainerVpnConnectionCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerVpnConnectionUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerVpnConnectionLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenEdgecontainerVpnConnectionNatGatewayIp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerVpnConnectionCluster(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerVpnConnectionVpc(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerVpnConnectionVpcProject(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["project_id"] =
		flattenEdgecontainerVpnConnectionVpcProjectProjectId(original["projectId"], d, config)
	return []interface{}{transformed}
}
func flattenEdgecontainerVpnConnectionVpcProjectProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerVpnConnectionEnableHighAvailability(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerVpnConnectionRouter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerVpnConnectionDetails(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["state"] =
		flattenEdgecontainerVpnConnectionDetailsState(original["state"], d, config)
	transformed["error"] =
		flattenEdgecontainerVpnConnectionDetailsError(original["error"], d, config)
	transformed["cloud_router"] =
		flattenEdgecontainerVpnConnectionDetailsCloudRouter(original["cloudRouter"], d, config)
	transformed["cloud_vpns"] =
		flattenEdgecontainerVpnConnectionDetailsCloudVpns(original["cloudVpns"], d, config)
	return []interface{}{transformed}
}
func flattenEdgecontainerVpnConnectionDetailsState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerVpnConnectionDetailsError(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerVpnConnectionDetailsCloudRouter(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["name"] =
		flattenEdgecontainerVpnConnectionDetailsCloudRouterName(original["name"], d, config)
	return []interface{}{transformed}
}
func flattenEdgecontainerVpnConnectionDetailsCloudRouterName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerVpnConnectionDetailsCloudVpns(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["gateway"] =
		flattenEdgecontainerVpnConnectionDetailsCloudVpnsGateway(original["gateway"], d, config)
	return []interface{}{transformed}
}
func flattenEdgecontainerVpnConnectionDetailsCloudVpnsGateway(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenEdgecontainerVpnConnectionTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenEdgecontainerVpnConnectionEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandEdgecontainerVpnConnectionNatGatewayIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerVpnConnectionCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerVpnConnectionVpc(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerVpnConnectionVpcProject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProjectId, err := expandEdgecontainerVpnConnectionVpcProjectProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	return transformed, nil
}

func expandEdgecontainerVpnConnectionVpcProjectProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerVpnConnectionEnableHighAvailability(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerVpnConnectionRouter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerVpnConnectionEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
