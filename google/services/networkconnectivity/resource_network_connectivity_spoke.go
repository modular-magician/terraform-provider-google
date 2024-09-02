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

package networkconnectivity

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

func ResourceNetworkConnectivitySpoke() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkConnectivitySpokeCreate,
		Read:   resourceNetworkConnectivitySpokeRead,
		Update: resourceNetworkConnectivitySpokeUpdate,
		Delete: resourceNetworkConnectivitySpokeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkConnectivitySpokeImport,
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
			"hub": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `Immutable. The URI of the hub that this spoke is attached to.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location for the resource`,
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `Immutable. The name of the spoke. Spoke names must be unique.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of the spoke.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"linked_interconnect_attachments": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"site_to_site_data_transfer": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: `A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.`,
						},
						"uris": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Description: `The URIs of linked interconnect attachment resources`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
				ConflictsWith: []string{"linked_vpn_tunnels", "linked_router_appliance_instances", "linked_vpc_network"},
			},
			"linked_router_appliance_instances": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `The URIs of linked Router appliance resources`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instances": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Description: `The list of router appliance instances`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: `The IP address on the VM to use for peering.`,
									},
									"virtual_machine": {
										Type:             schema.TypeString,
										Optional:         true,
										ForceNew:         true,
										DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
										Description:      `The URI of the virtual machine resource`,
									},
								},
							},
						},
						"site_to_site_data_transfer": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: `A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.`,
						},
					},
				},
				ConflictsWith: []string{"linked_interconnect_attachments", "linked_vpn_tunnels", "linked_vpc_network"},
			},
			"linked_vpc_network": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `VPC network that is associated with the spoke.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uri": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
							Description:      `The URI of the VPC network resource.`,
						},
						"exclude_export_ranges": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `IP ranges encompassing the subnets to be excluded from peering.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"include_export_ranges": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `IP ranges allowed to be included from peering.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
				ConflictsWith: []string{"linked_interconnect_attachments", "linked_router_appliance_instances", "linked_vpn_tunnels"},
			},
			"linked_vpn_tunnels": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `The URIs of linked VPN tunnel resources`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"site_to_site_data_transfer": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: `A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations.`,
						},
						"uris": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Description: `The URIs of linked VPN tunnel resources.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
				ConflictsWith: []string{"linked_interconnect_attachments", "linked_router_appliance_instances", "linked_vpc_network"},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The time the spoke was created.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"state": {
				Type:         schema.TypeString,
				Computed:     true,
				Description:  `Output only. The current lifecycle state of this spoke.`,
				ExactlyOneOf: []string{},
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"unique_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is deleted and another with the same name is created, the new spoke is assigned a different unique_id.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The time the spoke was last updated.`,
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

func resourceNetworkConnectivitySpokeCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandNetworkConnectivitySpokeName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandNetworkConnectivitySpokeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	hubProp, err := expandNetworkConnectivitySpokeHub(d.Get("hub"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("hub"); !tpgresource.IsEmptyValue(reflect.ValueOf(hubProp)) && (ok || !reflect.DeepEqual(v, hubProp)) {
		obj["hub"] = hubProp
	}
	linkedVpnTunnelsProp, err := expandNetworkConnectivitySpokeLinkedVpnTunnels(d.Get("linked_vpn_tunnels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("linked_vpn_tunnels"); !tpgresource.IsEmptyValue(reflect.ValueOf(linkedVpnTunnelsProp)) && (ok || !reflect.DeepEqual(v, linkedVpnTunnelsProp)) {
		obj["linkedVpnTunnels"] = linkedVpnTunnelsProp
	}
	linkedInterconnectAttachmentsProp, err := expandNetworkConnectivitySpokeLinkedInterconnectAttachments(d.Get("linked_interconnect_attachments"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("linked_interconnect_attachments"); !tpgresource.IsEmptyValue(reflect.ValueOf(linkedInterconnectAttachmentsProp)) && (ok || !reflect.DeepEqual(v, linkedInterconnectAttachmentsProp)) {
		obj["linkedInterconnectAttachments"] = linkedInterconnectAttachmentsProp
	}
	linkedRouterApplianceInstancesProp, err := expandNetworkConnectivitySpokeLinkedRouterApplianceInstances(d.Get("linked_router_appliance_instances"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("linked_router_appliance_instances"); !tpgresource.IsEmptyValue(reflect.ValueOf(linkedRouterApplianceInstancesProp)) && (ok || !reflect.DeepEqual(v, linkedRouterApplianceInstancesProp)) {
		obj["linkedRouterApplianceInstances"] = linkedRouterApplianceInstancesProp
	}
	linkedVpcNetworkProp, err := expandNetworkConnectivitySpokeLinkedVpcNetwork(d.Get("linked_vpc_network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("linked_vpc_network"); !tpgresource.IsEmptyValue(reflect.ValueOf(linkedVpcNetworkProp)) && (ok || !reflect.DeepEqual(v, linkedVpcNetworkProp)) {
		obj["linkedVpcNetwork"] = linkedVpcNetworkProp
	}
	labelsProp, err := expandNetworkConnectivitySpokeEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/{{location}}/spokes?spokeId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Spoke: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Spoke: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

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
		return fmt.Errorf("Error creating Spoke: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/spokes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkConnectivityOperationWaitTime(
		config, res, tpgresource.GetResourceNameFromSelfLink(project), "Creating Spoke", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Spoke: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Spoke %q: %#v", d.Id(), res)

	return resourceNetworkConnectivitySpokeRead(d, meta)
}

func resourceNetworkConnectivitySpokeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/{{location}}/spokes/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Spoke: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkConnectivitySpoke %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Spoke: %s", err)
	}

	if err := d.Set("name", flattenNetworkConnectivitySpokeName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Spoke: %s", err)
	}
	if err := d.Set("create_time", flattenNetworkConnectivitySpokeCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Spoke: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkConnectivitySpokeUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Spoke: %s", err)
	}
	if err := d.Set("labels", flattenNetworkConnectivitySpokeLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Spoke: %s", err)
	}
	if err := d.Set("description", flattenNetworkConnectivitySpokeDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Spoke: %s", err)
	}
	if err := d.Set("hub", flattenNetworkConnectivitySpokeHub(res["hub"], d, config)); err != nil {
		return fmt.Errorf("Error reading Spoke: %s", err)
	}
	if err := d.Set("linked_vpn_tunnels", flattenNetworkConnectivitySpokeLinkedVpnTunnels(res["linkedVpnTunnels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Spoke: %s", err)
	}
	if err := d.Set("linked_interconnect_attachments", flattenNetworkConnectivitySpokeLinkedInterconnectAttachments(res["linkedInterconnectAttachments"], d, config)); err != nil {
		return fmt.Errorf("Error reading Spoke: %s", err)
	}
	if err := d.Set("linked_router_appliance_instances", flattenNetworkConnectivitySpokeLinkedRouterApplianceInstances(res["linkedRouterApplianceInstances"], d, config)); err != nil {
		return fmt.Errorf("Error reading Spoke: %s", err)
	}
	if err := d.Set("linked_vpc_network", flattenNetworkConnectivitySpokeLinkedVpcNetwork(res["linkedVpcNetwork"], d, config)); err != nil {
		return fmt.Errorf("Error reading Spoke: %s", err)
	}
	if err := d.Set("unique_id", flattenNetworkConnectivitySpokeUniqueId(res["uniqueId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Spoke: %s", err)
	}
	if err := d.Set("state", flattenNetworkConnectivitySpokeState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Spoke: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetworkConnectivitySpokeTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Spoke: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetworkConnectivitySpokeEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Spoke: %s", err)
	}

	return nil
}

func resourceNetworkConnectivitySpokeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Spoke: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkConnectivitySpokeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetworkConnectivitySpokeEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/{{location}}/spokes/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Spoke %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

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
			return fmt.Errorf("Error updating Spoke %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Spoke %q: %#v", d.Id(), res)
		}

		err = NetworkConnectivityOperationWaitTime(
			config, res, tpgresource.GetResourceNameFromSelfLink(project), "Updating Spoke", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetworkConnectivitySpokeRead(d, meta)
}

func resourceNetworkConnectivitySpokeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Spoke: %s", err)
	}
	billingProject = strings.TrimPrefix(project, "projects/")

	url, err := tpgresource.ReplaceVarsForId(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/{{location}}/spokes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Spoke %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "Spoke")
	}

	err = NetworkConnectivityOperationWaitTime(
		config, res, tpgresource.GetResourceNameFromSelfLink(project), "Deleting Spoke", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Spoke %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkConnectivitySpokeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/spokes/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/spokes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkConnectivitySpokeName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkConnectivitySpokeDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeHub(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenNetworkConnectivitySpokeLinkedVpnTunnels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["uris"] =
		flattenNetworkConnectivitySpokeLinkedVpnTunnelsUris(original["uris"], d, config)
	transformed["site_to_site_data_transfer"] =
		flattenNetworkConnectivitySpokeLinkedVpnTunnelsSiteToSiteDataTransfer(original["siteToSiteDataTransfer"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkConnectivitySpokeLinkedVpnTunnelsUris(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeLinkedVpnTunnelsSiteToSiteDataTransfer(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeLinkedInterconnectAttachments(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["uris"] =
		flattenNetworkConnectivitySpokeLinkedInterconnectAttachmentsUris(original["uris"], d, config)
	transformed["site_to_site_data_transfer"] =
		flattenNetworkConnectivitySpokeLinkedInterconnectAttachmentsSiteToSiteDataTransfer(original["siteToSiteDataTransfer"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkConnectivitySpokeLinkedInterconnectAttachmentsUris(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeLinkedInterconnectAttachmentsSiteToSiteDataTransfer(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeLinkedRouterApplianceInstances(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["instances"] =
		flattenNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstances(original["instances"], d, config)
	transformed["site_to_site_data_transfer"] =
		flattenNetworkConnectivitySpokeLinkedRouterApplianceInstancesSiteToSiteDataTransfer(original["siteToSiteDataTransfer"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstances(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"virtual_machine": flattenNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstancesVirtualMachine(original["virtualMachine"], d, config),
			"ip_address":      flattenNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstancesIpAddress(original["ipAddress"], d, config),
		})
	}
	return transformed
}
func flattenNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstancesVirtualMachine(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstancesIpAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeLinkedRouterApplianceInstancesSiteToSiteDataTransfer(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeLinkedVpcNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["uri"] =
		flattenNetworkConnectivitySpokeLinkedVpcNetworkUri(original["uri"], d, config)
	transformed["exclude_export_ranges"] =
		flattenNetworkConnectivitySpokeLinkedVpcNetworkExcludeExportRanges(original["excludeExportRanges"], d, config)
	transformed["include_export_ranges"] =
		flattenNetworkConnectivitySpokeLinkedVpcNetworkIncludeExportRanges(original["includeExportRanges"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkConnectivitySpokeLinkedVpcNetworkUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeLinkedVpcNetworkExcludeExportRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeLinkedVpcNetworkIncludeExportRanges(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeUniqueId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivitySpokeTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkConnectivitySpokeEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkConnectivitySpokeName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeHub(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedVpnTunnels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUris, err := expandNetworkConnectivitySpokeLinkedVpnTunnelsUris(original["uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uris"] = transformedUris
	}

	transformedSiteToSiteDataTransfer, err := expandNetworkConnectivitySpokeLinkedVpnTunnelsSiteToSiteDataTransfer(original["site_to_site_data_transfer"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSiteToSiteDataTransfer); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["siteToSiteDataTransfer"] = transformedSiteToSiteDataTransfer
	}

	return transformed, nil
}

func expandNetworkConnectivitySpokeLinkedVpnTunnelsUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedVpnTunnelsSiteToSiteDataTransfer(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedInterconnectAttachments(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUris, err := expandNetworkConnectivitySpokeLinkedInterconnectAttachmentsUris(original["uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uris"] = transformedUris
	}

	transformedSiteToSiteDataTransfer, err := expandNetworkConnectivitySpokeLinkedInterconnectAttachmentsSiteToSiteDataTransfer(original["site_to_site_data_transfer"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSiteToSiteDataTransfer); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["siteToSiteDataTransfer"] = transformedSiteToSiteDataTransfer
	}

	return transformed, nil
}

func expandNetworkConnectivitySpokeLinkedInterconnectAttachmentsUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedInterconnectAttachmentsSiteToSiteDataTransfer(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedRouterApplianceInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedInstances, err := expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstances(original["instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstances); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["instances"] = transformedInstances
	}

	transformedSiteToSiteDataTransfer, err := expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesSiteToSiteDataTransfer(original["site_to_site_data_transfer"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSiteToSiteDataTransfer); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["siteToSiteDataTransfer"] = transformedSiteToSiteDataTransfer
	}

	return transformed, nil
}

func expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedVirtualMachine, err := expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstancesVirtualMachine(original["virtual_machine"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVirtualMachine); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["virtualMachine"] = transformedVirtualMachine
		}

		transformedIpAddress, err := expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstancesIpAddress(original["ip_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipAddress"] = transformedIpAddress
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstancesVirtualMachine(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstancesIpAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesSiteToSiteDataTransfer(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedVpcNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUri, err := expandNetworkConnectivitySpokeLinkedVpcNetworkUri(original["uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	transformedExcludeExportRanges, err := expandNetworkConnectivitySpokeLinkedVpcNetworkExcludeExportRanges(original["exclude_export_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExcludeExportRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["excludeExportRanges"] = transformedExcludeExportRanges
	}

	transformedIncludeExportRanges, err := expandNetworkConnectivitySpokeLinkedVpcNetworkIncludeExportRanges(original["include_export_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeExportRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["includeExportRanges"] = transformedIncludeExportRanges
	}

	return transformed, nil
}

func expandNetworkConnectivitySpokeLinkedVpcNetworkUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedVpcNetworkExcludeExportRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedVpcNetworkIncludeExportRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
