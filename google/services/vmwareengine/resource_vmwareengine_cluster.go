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

package vmwareengine

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceVmwareengineCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceVmwareengineClusterCreate,
		Read:   resourceVmwareengineClusterRead,
		Update: resourceVmwareengineClusterUpdate,
		Delete: resourceVmwareengineClusterDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVmwareengineClusterImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(210 * time.Minute),
			Update: schema.DefaultTimeout(190 * time.Minute),
			Delete: schema.DefaultTimeout(150 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the Cluster.`,
			},
			"parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The resource name of the private cloud to create a new cluster in.
Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud`,
			},
			"node_type_configs": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: `The map of cluster node types in this cluster,
where the key is canonical identifier of the node type (corresponds to the NodeType).`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_type_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"node_count": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: `The number of nodes of this type in the cluster.`,
						},
						"custom_core_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Description: `Customized number of cores available to each node of the type.
This number must always be one of 'nodeType.availableCustomCoreCounts'.
If zero is provided max value from 'nodeType.availableCustomCoreCounts' will be used.
Once the customer is created then corecount cannot be changed.`,
							Default: 0,
						},
					},
				},
			},
			"management": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: `True if the cluster is a management cluster; false otherwise.
There can only be one management cluster in a private cloud and it has to be the first one.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `State of the Cluster.`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `System-generated unique identifier for the resource.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceVmwareengineClusterCreate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nodeTypeConfigsProp, err := expandVmwareengineClusterNodeTypeConfigs(d.Get("node_type_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_type_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeTypeConfigsProp)) && (ok || !reflect.DeepEqual(v, nodeTypeConfigsProp)) {
		obj["nodeTypeConfigs"] = nodeTypeConfigsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}{{parent}}/clusters?clusterId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Cluster: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "POST",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutCreate),
		Headers:              headers,
		ErrorAbortPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.Is429QuotaError},
	})
	if err != nil {
		return fmt.Errorf("Error creating Cluster: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/clusters/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = VmwareengineOperationWaitTime(
		config, res, project, "Creating Cluster", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Cluster: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Cluster %q: %#v", d.Id(), res)

	return resourceVmwareengineClusterRead(d, meta)
}

func resourceVmwareengineClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}{{parent}}/clusters/{{name}}")
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
		Config:               config,
		Method:               "GET",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Headers:              headers,
		ErrorAbortPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.Is429QuotaError},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("VmwareengineCluster %q", d.Id()))
	}

	if err := d.Set("management", flattenVmwareengineClusterManagement(res["management"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("uid", flattenVmwareengineClusterUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("state", flattenVmwareengineClusterState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("node_type_configs", flattenVmwareengineClusterNodeTypeConfigs(res["nodeTypeConfigs"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}

	return nil
}

func resourceVmwareengineClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	nodeTypeConfigsProp, err := expandVmwareengineClusterNodeTypeConfigs(d.Get("node_type_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("node_type_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nodeTypeConfigsProp)) {
		obj["nodeTypeConfigs"] = nodeTypeConfigsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}{{parent}}/clusters/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Cluster %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("node_type_configs") {
		updateMask = append(updateMask, "nodeTypeConfigs.*.nodeCount")
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
			Config:               config,
			Method:               "PATCH",
			Project:              billingProject,
			RawURL:               url,
			UserAgent:            userAgent,
			Body:                 obj,
			Timeout:              d.Timeout(schema.TimeoutUpdate),
			Headers:              headers,
			ErrorAbortPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.Is429QuotaError},
		})

		if err != nil {
			return fmt.Errorf("Error updating Cluster %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Cluster %q: %#v", d.Id(), res)
		}

		err = VmwareengineOperationWaitTime(
			config, res, project, "Updating Cluster", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceVmwareengineClusterRead(d, meta)
}

func resourceVmwareengineClusterDelete(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{VmwareengineBasePath}}{{parent}}/clusters/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Cluster %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "DELETE",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutDelete),
		Headers:              headers,
		ErrorAbortPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.Is429QuotaError},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Cluster")
	}

	err = VmwareengineOperationWaitTime(
		config, res, project, "Deleting Cluster", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Cluster %q: %#v", d.Id(), res)
	return nil
}

func resourceVmwareengineClusterImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<parent>.+)/clusters/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/clusters/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenVmwareengineClusterManagement(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineClusterUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineClusterState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenVmwareengineClusterNodeTypeConfigs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.(map[string]interface{})
	transformed := make([]interface{}, 0, len(l))
	for k, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"node_type_id":      k,
			"node_count":        flattenVmwareengineClusterNodeTypeConfigsNodeCount(original["nodeCount"], d, config),
			"custom_core_count": flattenVmwareengineClusterNodeTypeConfigsCustomCoreCount(original["customCoreCount"], d, config),
		})
	}
	return transformed
}
func flattenVmwareengineClusterNodeTypeConfigsNodeCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenVmwareengineClusterNodeTypeConfigsCustomCoreCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func expandVmwareengineClusterNodeTypeConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNodeCount, err := expandVmwareengineClusterNodeTypeConfigsNodeCount(original["node_count"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["nodeCount"] = transformedNodeCount
		}

		transformedCustomCoreCount, err := expandVmwareengineClusterNodeTypeConfigsCustomCoreCount(original["custom_core_count"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCustomCoreCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["customCoreCount"] = transformedCustomCoreCount
		}

		transformedNodeTypeId, err := tpgresource.ExpandString(original["node_type_id"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedNodeTypeId] = transformed
	}
	return m, nil
}

func expandVmwareengineClusterNodeTypeConfigsNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineClusterNodeTypeConfigsCustomCoreCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
