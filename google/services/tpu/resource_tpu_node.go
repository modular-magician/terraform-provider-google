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

package tpu

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

// compareTpuNodeSchedulingConfig diff suppresses for the default
// scheduling, i.e. if preemptible is false, the API may either return no
// schedulingConfig or an empty schedulingConfig.
func compareTpuNodeSchedulingConfig(k, old, new string, d *schema.ResourceData) bool {
	if k == "scheduling_config.0.preemptible" {
		return old == "" && new == "false"
	}
	if k == "scheduling_config.#" {
		o, n := d.GetChange("scheduling_config.0.preemptible")
		return o.(bool) == n.(bool)
	}
	return false
}

func tpuNodeCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, meta interface{}) error {
	old, new := diff.GetChange("network")
	config := meta.(*transport_tpg.Config)

	networkLinkRegex := regexp.MustCompile("projects/(.+)/global/networks/(.+)")

	var pid string

	if networkLinkRegex.MatchString(new.(string)) {
		parts := networkLinkRegex.FindStringSubmatch(new.(string))
		pid = parts[1]
	}

	if pid == "" {
		return nil
	}

	project, err := config.NewResourceManagerClient(config.UserAgent).Projects.Get(pid).Do()
	if err != nil {
		return fmt.Errorf("Failed to retrieve project, pid: %s, err: %s", pid, err)
	}

	if networkLinkRegex.MatchString(old.(string)) {
		parts := networkLinkRegex.FindStringSubmatch(old.(string))
		i, err := tpgresource.StringToFixed64(parts[1])
		if err == nil {
			if project.ProjectNumber == i {
				if err := diff.SetNew("network", old); err != nil {
					return err
				}
				return nil
			}
		}
	}
	return nil
}

func ResourceTPUNode() *schema.Resource {
	return &schema.Resource{
		Create: resourceTPUNodeCreate,
		Read:   resourceTPUNodeRead,
		Update: resourceTPUNodeUpdate,
		Delete: resourceTPUNodeDelete,

		Importer: &schema.ResourceImporter{
			State: resourceTPUNodeImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpuNodeCustomizeDiff,
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"accelerator_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The type of hardware accelerators associated with this node.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The immutable name of the TPU.`,
			},
			"tensorflow_version": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The version of Tensorflow running in the Node.`,
			},
			"cidr_block": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `The CIDR block that the TPU node will use when selecting an IP
address. This CIDR block must be a /29 block; the Compute Engine
networks API forbids a smaller block, and using a larger block would
be wasteful (a node can only consume one IP address).

Errors will occur if the CIDR block has already been used for a
currently existing TPU node, the CIDR block conflicts with any
subnetworks in the user's provided network, or the provided network
is peered with another network that is using that CIDR block.`,
				ConflictsWith: []string{"use_service_networking"},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The user-supplied description of the TPU. Maximum of 512 characters.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: `Resource labels to represent user provided metadata.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"network": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `The name of a network to peer the TPU node to. It must be a
preexisting Compute Engine network inside of the project on which
this API has been activated. If none is provided, "default" will be
used.`,
			},
			"scheduling_config": {
				Type:             schema.TypeList,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareTpuNodeSchedulingConfig,
				Description:      `Sets the scheduling options for this TPU instance.`,
				MaxItems:         1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"preemptible": {
							Type:             schema.TypeBool,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: compareTpuNodeSchedulingConfig,
							Description:      `Defines whether the TPU instance is preemptible.`,
						},
					},
				},
			},
			"use_service_networking": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `Whether the VPC peering for the node is set up through Service Networking API.
The VPC Peering should be set up before provisioning the node. If this field is set,
cidr_block field should not be specified. If the network that you want to peer the
TPU Node to is a Shared VPC network, the node must be created with this this field enabled.`,
				Default:       false,
				ConflictsWith: []string{"cidr_block"},
			},
			"zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The GCP location for the TPU. If it is not provided, the provider zone is used.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"network_endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `The network endpoints where TPU workers can be accessed and sent work.
It is recommended that Tensorflow clients of the node first reach out
to the first (index 0) entry.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The IP address of this network endpoint.`,
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `The port of this network endpoint.`,
						},
					},
				},
			},
			"service_account": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The service account used to run the tensor flow services within the
node. To share resources, including Google Cloud Storage data, with
the Tensorflow job running in the Node, this account must have
permissions to that data.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceTPUNodeCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandTPUNodeName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandTPUNodeDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	acceleratorTypeProp, err := expandTPUNodeAcceleratorType(d.Get("accelerator_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("accelerator_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(acceleratorTypeProp)) && (ok || !reflect.DeepEqual(v, acceleratorTypeProp)) {
		obj["acceleratorType"] = acceleratorTypeProp
	}
	tensorflowVersionProp, err := expandTPUNodeTensorflowVersion(d.Get("tensorflow_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tensorflow_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(tensorflowVersionProp)) && (ok || !reflect.DeepEqual(v, tensorflowVersionProp)) {
		obj["tensorflowVersion"] = tensorflowVersionProp
	}
	networkProp, err := expandTPUNodeNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	cidrBlockProp, err := expandTPUNodeCidrBlock(d.Get("cidr_block"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cidr_block"); !tpgresource.IsEmptyValue(reflect.ValueOf(cidrBlockProp)) && (ok || !reflect.DeepEqual(v, cidrBlockProp)) {
		obj["cidrBlock"] = cidrBlockProp
	}
	useServiceNetworkingProp, err := expandTPUNodeUseServiceNetworking(d.Get("use_service_networking"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("use_service_networking"); !tpgresource.IsEmptyValue(reflect.ValueOf(useServiceNetworkingProp)) && (ok || !reflect.DeepEqual(v, useServiceNetworkingProp)) {
		obj["useServiceNetworking"] = useServiceNetworkingProp
	}
	schedulingConfigProp, err := expandTPUNodeSchedulingConfig(d.Get("scheduling_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("scheduling_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(schedulingConfigProp)) && (ok || !reflect.DeepEqual(v, schedulingConfigProp)) {
		obj["schedulingConfig"] = schedulingConfigProp
	}
	labelsProp, err := expandTPUNodeEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{TPUBasePath}}projects/{{project}}/locations/{{zone}}/nodes?nodeId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Node: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Node: %s", err)
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
		return fmt.Errorf("Error creating Node: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{zone}}/nodes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = TPUOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Node", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Node: %s", err)
	}

	if err := d.Set("name", flattenTPUNodeName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{zone}}/nodes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Node %q: %#v", d.Id(), res)

	return resourceTPUNodeRead(d, meta)
}

func resourceTPUNodeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{TPUBasePath}}projects/{{project}}/locations/{{zone}}/nodes/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Node: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("TPUNode %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}

	if err := d.Set("name", flattenTPUNodeName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("description", flattenTPUNodeDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("accelerator_type", flattenTPUNodeAcceleratorType(res["acceleratorType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("tensorflow_version", flattenTPUNodeTensorflowVersion(res["tensorflowVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("network", flattenTPUNodeNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("cidr_block", flattenTPUNodeCidrBlock(res["cidrBlock"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("service_account", flattenTPUNodeServiceAccount(res["serviceAccount"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("use_service_networking", flattenTPUNodeUseServiceNetworking(res["useServiceNetworking"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("scheduling_config", flattenTPUNodeSchedulingConfig(res["schedulingConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("network_endpoints", flattenTPUNodeNetworkEndpoints(res["networkEndpoints"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("labels", flattenTPUNodeLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("terraform_labels", flattenTPUNodeTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}
	if err := d.Set("effective_labels", flattenTPUNodeEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Node: %s", err)
	}

	return nil
}

func resourceTPUNodeUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Node: %s", err)
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("tensorflow_version") {
		obj := make(map[string]interface{})

		tensorflowVersionProp, err := expandTPUNodeTensorflowVersion(d.Get("tensorflow_version"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("tensorflow_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, tensorflowVersionProp)) {
			obj["tensorflowVersion"] = tensorflowVersionProp
		}

		url, err := tpgresource.ReplaceVars(d, config, "{{TPUBasePath}}projects/{{project}}/locations/{{zone}}/nodes/{{name}}:reimage")
		if err != nil {
			return err
		}

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
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})
		if err != nil {
			return fmt.Errorf("Error updating Node %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Node %q: %#v", d.Id(), res)
		}

		err = TPUOperationWaitTime(
			config, res, project, "Updating Node", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceTPUNodeRead(d, meta)
}

func resourceTPUNodeDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Node: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{TPUBasePath}}projects/{{project}}/locations/{{zone}}/nodes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Node %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "Node")
	}

	err = TPUOperationWaitTime(
		config, res, project, "Deleting Node", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Node %q: %#v", d.Id(), res)
	return nil
}

func resourceTPUNodeImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<zone>[^/]+)/nodes/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{zone}}/nodes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenTPUNodeName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenTPUNodeDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeAcceleratorType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeTensorflowVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeCidrBlock(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeUseServiceNetworking(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeSchedulingConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["preemptible"] =
		flattenTPUNodeSchedulingConfigPreemptible(original["preemptible"], d, config)
	return []interface{}{transformed}
}
func flattenTPUNodeSchedulingConfigPreemptible(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeNetworkEndpoints(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"ip_address": flattenTPUNodeNetworkEndpointsIpAddress(original["ipAddress"], d, config),
			"port":       flattenTPUNodeNetworkEndpointsPort(original["port"], d, config),
		})
	}
	return transformed
}
func flattenTPUNodeNetworkEndpointsIpAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenTPUNodeNetworkEndpointsPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
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

func flattenTPUNodeLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenTPUNodeTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenTPUNodeEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandTPUNodeName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeAcceleratorType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeTensorflowVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeCidrBlock(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeUseServiceNetworking(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeSchedulingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPreemptible, err := expandTPUNodeSchedulingConfigPreemptible(original["preemptible"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPreemptible); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["preemptible"] = transformedPreemptible
	}

	return transformed, nil
}

func expandTPUNodeSchedulingConfigPreemptible(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandTPUNodeEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
