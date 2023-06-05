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

package google

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

func ResourceNetworkManagementConnectivityTest() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkManagementConnectivityTestCreate,
		Read:   resourceNetworkManagementConnectivityTestRead,
		Update: resourceNetworkManagementConnectivityTestUpdate,
		Delete: resourceNetworkManagementConnectivityTestDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkManagementConnectivityTestImport,
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
			"destination": {
				Type:     schema.TypeList,
				Required: true,
				Description: `Required. Destination specification of the Connectivity Test.

You can use a combination of destination IP address, Compute
Engine VM instance, or VPC network to uniquely identify the
destination location.

Even if the destination IP address is not unique, the source IP
location is unique. Usually, the analysis can infer the destination
endpoint from route information.

If the destination you specify is a VM instance and the instance has
multiple network interfaces, then you must also specify either a
destination IP address or VPC network to identify the destination
interface.

A reachability analysis proceeds even if the destination location
is ambiguous. However, the result can include endpoints that you
don't intend to test.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `A Compute Engine instance URI.`,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The IP address of the endpoint, which can be an external or
internal IP. An IPv6 address is only allowed when the test's
destination is a global load balancer VIP.`,
						},
						"network": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `A Compute Engine network URI.`,
						},
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Description: `The IP protocol port of the endpoint. Only applicable when
protocol is TCP or UDP.`,
						},
						"project_id": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Project ID where the endpoint is located. The Project ID can be
derived from the URI if you provide a VM instance or network URI.
The following are two cases where you must provide the project ID:
1. Only the IP address is specified, and the IP address is within
a GCP project. 2. When you are using Shared VPC and the IP address
that you provide is from the service project. In this case, the
network that the IP address resides in is defined in the host
project.`,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Unique name for the connectivity test.`,
			},
			"source": {
				Type:     schema.TypeList,
				Required: true,
				Description: `Required. Source specification of the Connectivity Test.

You can use a combination of source IP address, virtual machine
(VM) instance, or Compute Engine network to uniquely identify the
source location.

Examples: If the source IP address is an internal IP address within
a Google Cloud Virtual Private Cloud (VPC) network, then you must
also specify the VPC network. Otherwise, specify the VM instance,
which already contains its internal IP address and VPC network
information.

If the source of the test is within an on-premises network, then
you must provide the destination VPC network.

If the source endpoint is a Compute Engine VM instance with multiple
network interfaces, the instance itself is not sufficient to
identify the endpoint. So, you must also specify the source IP
address or VPC network.

A reachability analysis proceeds even if the source location is
ambiguous. However, the test result may include endpoints that
you don't intend to test.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `A Compute Engine instance URI.`,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The IP address of the endpoint, which can be an external or
internal IP. An IPv6 address is only allowed when the test's
destination is a global load balancer VIP.`,
						},
						"network": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `A Compute Engine network URI.`,
						},
						"network_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"GCP_NETWORK", "NON_GCP_NETWORK", ""}),
							Description:  `Type of the network where the endpoint is located. Possible values: ["GCP_NETWORK", "NON_GCP_NETWORK"]`,
						},
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Description: `The IP protocol port of the endpoint. Only applicable when
protocol is TCP or UDP.`,
						},
						"project_id": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Project ID where the endpoint is located. The Project ID can be
derived from the URI if you provide a VM instance or network URI.
The following are two cases where you must provide the project ID:

1. Only the IP address is specified, and the IP address is
   within a GCP project.
2. When you are using Shared VPC and the IP address
   that you provide is from the service project. In this case,
   the network that the IP address resides in is defined in the
   host project.`,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The user-supplied description of the Connectivity Test.
Maximum of 512 characters.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Resource labels to represent user-provided metadata.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `IP Protocol of the test. When not provided, "TCP" is assumed.`,
				Default:     "TCP",
			},
			"related_projects": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Other projects that may be relevant for reachability analysis.
This is applicable to scenarios where a test can cross project
boundaries.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func resourceNetworkManagementConnectivityTestCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandNetworkManagementConnectivityTestName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandNetworkManagementConnectivityTestDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sourceProp, err := expandNetworkManagementConnectivityTestSource(d.Get("source"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceProp)) && (ok || !reflect.DeepEqual(v, sourceProp)) {
		obj["source"] = sourceProp
	}
	destinationProp, err := expandNetworkManagementConnectivityTestDestination(d.Get("destination"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("destination"); !tpgresource.IsEmptyValue(reflect.ValueOf(destinationProp)) && (ok || !reflect.DeepEqual(v, destinationProp)) {
		obj["destination"] = destinationProp
	}
	protocolProp, err := expandNetworkManagementConnectivityTestProtocol(d.Get("protocol"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("protocol"); !tpgresource.IsEmptyValue(reflect.ValueOf(protocolProp)) && (ok || !reflect.DeepEqual(v, protocolProp)) {
		obj["protocol"] = protocolProp
	}
	relatedProjectsProp, err := expandNetworkManagementConnectivityTestRelatedProjects(d.Get("related_projects"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("related_projects"); !tpgresource.IsEmptyValue(reflect.ValueOf(relatedProjectsProp)) && (ok || !reflect.DeepEqual(v, relatedProjectsProp)) {
		obj["relatedProjects"] = relatedProjectsProp
	}
	labelsProp, err := expandNetworkManagementConnectivityTestLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkManagementBasePath}}projects/{{project}}/locations/global/connectivityTests?testId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ConnectivityTest: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ConnectivityTest: %s", err)
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
		return fmt.Errorf("Error creating ConnectivityTest: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/connectivityTests/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = NetworkManagementOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating ConnectivityTest", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create ConnectivityTest: %s", err)
	}

	if err := d.Set("name", flattenNetworkManagementConnectivityTestName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/connectivityTests/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ConnectivityTest %q: %#v", d.Id(), res)

	return resourceNetworkManagementConnectivityTestRead(d, meta)
}

func resourceNetworkManagementConnectivityTestRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkManagementBasePath}}projects/{{project}}/locations/global/connectivityTests/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ConnectivityTest: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkManagementConnectivityTest %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ConnectivityTest: %s", err)
	}

	if err := d.Set("name", flattenNetworkManagementConnectivityTestName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectivityTest: %s", err)
	}
	if err := d.Set("description", flattenNetworkManagementConnectivityTestDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectivityTest: %s", err)
	}
	if err := d.Set("source", flattenNetworkManagementConnectivityTestSource(res["source"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectivityTest: %s", err)
	}
	if err := d.Set("destination", flattenNetworkManagementConnectivityTestDestination(res["destination"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectivityTest: %s", err)
	}
	if err := d.Set("protocol", flattenNetworkManagementConnectivityTestProtocol(res["protocol"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectivityTest: %s", err)
	}
	if err := d.Set("related_projects", flattenNetworkManagementConnectivityTestRelatedProjects(res["relatedProjects"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectivityTest: %s", err)
	}
	if err := d.Set("labels", flattenNetworkManagementConnectivityTestLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ConnectivityTest: %s", err)
	}

	return nil
}

func resourceNetworkManagementConnectivityTestUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ConnectivityTest: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkManagementConnectivityTestDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sourceProp, err := expandNetworkManagementConnectivityTestSource(d.Get("source"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sourceProp)) {
		obj["source"] = sourceProp
	}
	destinationProp, err := expandNetworkManagementConnectivityTestDestination(d.Get("destination"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("destination"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, destinationProp)) {
		obj["destination"] = destinationProp
	}
	protocolProp, err := expandNetworkManagementConnectivityTestProtocol(d.Get("protocol"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("protocol"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, protocolProp)) {
		obj["protocol"] = protocolProp
	}
	relatedProjectsProp, err := expandNetworkManagementConnectivityTestRelatedProjects(d.Get("related_projects"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("related_projects"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, relatedProjectsProp)) {
		obj["relatedProjects"] = relatedProjectsProp
	}
	labelsProp, err := expandNetworkManagementConnectivityTestLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkManagementBasePath}}projects/{{project}}/locations/global/connectivityTests/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ConnectivityTest %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("source") {
		updateMask = append(updateMask, "source.ipAddress",
			"source.port",
			"source.instance",
			"source.network",
			"source.networkType",
			"source.projectId")
	}

	if d.HasChange("destination") {
		updateMask = append(updateMask, "destination.ipAddress",
			"destination.port",
			"destination.instance",
			"destination.network",
			"destination.projectId")
	}

	if d.HasChange("protocol") {
		updateMask = append(updateMask, "protocol")
	}

	if d.HasChange("related_projects") {
		updateMask = append(updateMask, "relatedProjects")
	}

	if d.HasChange("labels") {
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
		return fmt.Errorf("Error updating ConnectivityTest %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ConnectivityTest %q: %#v", d.Id(), res)
	}

	err = NetworkManagementOperationWaitTime(
		config, res, project, "Updating ConnectivityTest", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceNetworkManagementConnectivityTestRead(d, meta)
}

func resourceNetworkManagementConnectivityTestDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ConnectivityTest: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkManagementBasePath}}projects/{{project}}/locations/global/connectivityTests/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ConnectivityTest %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "ConnectivityTest")
	}

	err = NetworkManagementOperationWaitTime(
		config, res, project, "Deleting ConnectivityTest", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ConnectivityTest %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkManagementConnectivityTestImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/connectivityTests/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/global/connectivityTests/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkManagementConnectivityTestName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenNetworkManagementConnectivityTestDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkManagementConnectivityTestSource(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["ip_address"] =
		flattenNetworkManagementConnectivityTestSourceIpAddress(original["ipAddress"], d, config)
	transformed["port"] =
		flattenNetworkManagementConnectivityTestSourcePort(original["port"], d, config)
	transformed["instance"] =
		flattenNetworkManagementConnectivityTestSourceInstance(original["instance"], d, config)
	transformed["network"] =
		flattenNetworkManagementConnectivityTestSourceNetwork(original["network"], d, config)
	transformed["network_type"] =
		flattenNetworkManagementConnectivityTestSourceNetworkType(original["networkType"], d, config)
	transformed["project_id"] =
		flattenNetworkManagementConnectivityTestSourceProjectId(original["projectId"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkManagementConnectivityTestSourceIpAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkManagementConnectivityTestSourcePort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
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

func flattenNetworkManagementConnectivityTestSourceInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkManagementConnectivityTestSourceNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkManagementConnectivityTestSourceNetworkType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkManagementConnectivityTestSourceProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkManagementConnectivityTestDestination(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["ip_address"] =
		flattenNetworkManagementConnectivityTestDestinationIpAddress(original["ipAddress"], d, config)
	transformed["port"] =
		flattenNetworkManagementConnectivityTestDestinationPort(original["port"], d, config)
	transformed["instance"] =
		flattenNetworkManagementConnectivityTestDestinationInstance(original["instance"], d, config)
	transformed["network"] =
		flattenNetworkManagementConnectivityTestDestinationNetwork(original["network"], d, config)
	transformed["project_id"] =
		flattenNetworkManagementConnectivityTestDestinationProjectId(original["projectId"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkManagementConnectivityTestDestinationIpAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkManagementConnectivityTestDestinationPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
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

func flattenNetworkManagementConnectivityTestDestinationInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkManagementConnectivityTestDestinationNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkManagementConnectivityTestDestinationProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkManagementConnectivityTestProtocol(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkManagementConnectivityTestRelatedProjects(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkManagementConnectivityTestLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkManagementConnectivityTestName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	// projects/X/tests/Y - note not "connectivityTests"
	f, err := tpgresource.ParseGlobalFieldValue("tests", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandNetworkManagementConnectivityTestDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIpAddress, err := expandNetworkManagementConnectivityTestSourceIpAddress(original["ip_address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ipAddress"] = transformedIpAddress
	}

	transformedPort, err := expandNetworkManagementConnectivityTestSourcePort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedInstance, err := expandNetworkManagementConnectivityTestSourceInstance(original["instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["instance"] = transformedInstance
	}

	transformedNetwork, err := expandNetworkManagementConnectivityTestSourceNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	transformedNetworkType, err := expandNetworkManagementConnectivityTestSourceNetworkType(original["network_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworkType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["networkType"] = transformedNetworkType
	}

	transformedProjectId, err := expandNetworkManagementConnectivityTestSourceProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	return transformed, nil
}

func expandNetworkManagementConnectivityTestSourceIpAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSourcePort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSourceInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSourceNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSourceNetworkType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestSourceProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestDestination(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIpAddress, err := expandNetworkManagementConnectivityTestDestinationIpAddress(original["ip_address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ipAddress"] = transformedIpAddress
	}

	transformedPort, err := expandNetworkManagementConnectivityTestDestinationPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedInstance, err := expandNetworkManagementConnectivityTestDestinationInstance(original["instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["instance"] = transformedInstance
	}

	transformedNetwork, err := expandNetworkManagementConnectivityTestDestinationNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	transformedProjectId, err := expandNetworkManagementConnectivityTestDestinationProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	return transformed, nil
}

func expandNetworkManagementConnectivityTestDestinationIpAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestDestinationPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestDestinationInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestDestinationNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestDestinationProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestProtocol(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestRelatedProjects(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementConnectivityTestLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
