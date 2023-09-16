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
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceNetworkConnectivityServiceConnectionPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkConnectivityServiceConnectionPolicyCreate,
		Read:   resourceNetworkConnectivityServiceConnectionPolicyRead,
		Update: resourceNetworkConnectivityServiceConnectionPolicyUpdate,
		Delete: resourceNetworkConnectivityServiceConnectionPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetworkConnectivityServiceConnectionPolicyImport,
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
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location of the ServiceConnectionPolicy.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of a ServiceConnectionPolicy. Format: projects/{project}/locations/{location}/serviceConnectionPolicies/{service_connection_policy} See: https://google.aip.dev/122#fields-representing-resource-names`,
			},
			"network": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource path of the consumer network. Example: - projects/{projectNumOrId}/global/networks/{resourceId}.`,
			},
			"service_class": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The service class identifier for which this ServiceConnectionPolicy is for. The service class identifier is a unique, symbolic representation of a ServiceClass.
It is provided by the Service Producer. Google services have a prefix of gcp. For example, gcp-cloud-sql. 3rd party services do not. For example, test-service-a3dfcx.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Free-text description of the resource.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `User-defined labels.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"psc_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Configuration used for Private Service Connect connections. Used when Infrastructure is PSC.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnetworks": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `IDs of the subnetworks or fully qualified identifiers for the subnetworks`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"limit": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Max number of PSC connections for this policy.`,
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp when the resource was created.`,
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The etag is computed by the server, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.`,
			},
			"infrastructure": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The type of underlying resources used to create the connection.`,
			},
			"psc_connections": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Information about each Private Service Connect connection.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"consumer_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The resource reference of the consumer address.`,
						},
						"consumer_forwarding_rule": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The resource reference of the PSC Forwarding Rule within the consumer VPC.`,
						},
						"consumer_target_project": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The project where the PSC connection is created.`,
						},
						"error": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `The most recent error during operating this connection.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: `The status code, which should be an enum value of [google.rpc.Code][].`,
									},
									"message": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `A developer-facing error message.`,
									},
									"details": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: `A list of messages that carry the error details.`,
										Elem: &schema.Schema{
											Type: schema.TypeMap,
										},
									},
								},
							},
						},
						"error_info": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `The error info for the latest error during operating this connection.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The logical grouping to which the "reason" belongs.`,
									},
									"metadata": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: `Additional structured details about this error.`,
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"reason": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The reason of the error.`,
									},
								},
							},
						},
						"error_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"CONNECTION_ERROR_TYPE_UNSPECIFIED", "ERROR_INTERNAL", "ERROR_CONSUMER_SIDE", "ERROR_PRODUCER_SIDE", ""}),
							Description: `The error type indicates whether the error is consumer facing, producer
facing or system internal. Possible values: ["CONNECTION_ERROR_TYPE_UNSPECIFIED", "ERROR_INTERNAL", "ERROR_CONSUMER_SIDE", "ERROR_PRODUCER_SIDE"]`,
						},
						"gce_operation": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The last Compute Engine operation to setup PSC connection.`,
						},
						"psc_connection_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The PSC connection id of the PSC forwarding rule.`,
						},
						"state": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"STATE_UNSPECIFIED", "ACTIVE", "CREATING", "DELETING", "FAILED", ""}),
							Description:  `The state of the PSC connection. Possible values: ["STATE_UNSPECIFIED", "ACTIVE", "CREATING", "DELETING", "FAILED"]`,
						},
					},
				},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The timestamp when the resource was updated.`,
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

func resourceNetworkConnectivityServiceConnectionPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	serviceClassProp, err := expandNetworkConnectivityServiceConnectionPolicyServiceClass(d.Get("service_class"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service_class"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceClassProp)) && (ok || !reflect.DeepEqual(v, serviceClassProp)) {
		obj["serviceClass"] = serviceClassProp
	}
	descriptionProp, err := expandNetworkConnectivityServiceConnectionPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networkProp, err := expandNetworkConnectivityServiceConnectionPolicyNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	pscConfigProp, err := expandNetworkConnectivityServiceConnectionPolicyPscConfig(d.Get("psc_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("psc_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(pscConfigProp)) && (ok || !reflect.DeepEqual(v, pscConfigProp)) {
		obj["pscConfig"] = pscConfigProp
	}
	etagProp, err := expandNetworkConnectivityServiceConnectionPolicyEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	labelsProp, err := expandNetworkConnectivityServiceConnectionPolicyLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/{{location}}/serviceConnectionPolicies?serviceConnectionPolicyId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ServiceConnectionPolicy: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceConnectionPolicy: %s", err)
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
		return fmt.Errorf("Error creating ServiceConnectionPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/serviceConnectionPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetworkConnectivityOperationWaitTime(
		config, res, project, "Creating ServiceConnectionPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ServiceConnectionPolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ServiceConnectionPolicy %q: %#v", d.Id(), res)

	return resourceNetworkConnectivityServiceConnectionPolicyRead(d, meta)
}

func resourceNetworkConnectivityServiceConnectionPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/{{location}}/serviceConnectionPolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceConnectionPolicy: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetworkConnectivityServiceConnectionPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ServiceConnectionPolicy: %s", err)
	}

	if err := d.Set("create_time", flattenNetworkConnectivityServiceConnectionPolicyCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceConnectionPolicy: %s", err)
	}
	if err := d.Set("update_time", flattenNetworkConnectivityServiceConnectionPolicyUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceConnectionPolicy: %s", err)
	}
	if err := d.Set("service_class", flattenNetworkConnectivityServiceConnectionPolicyServiceClass(res["serviceClass"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceConnectionPolicy: %s", err)
	}
	if err := d.Set("description", flattenNetworkConnectivityServiceConnectionPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceConnectionPolicy: %s", err)
	}
	if err := d.Set("network", flattenNetworkConnectivityServiceConnectionPolicyNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceConnectionPolicy: %s", err)
	}
	if err := d.Set("psc_config", flattenNetworkConnectivityServiceConnectionPolicyPscConfig(res["pscConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceConnectionPolicy: %s", err)
	}
	if err := d.Set("etag", flattenNetworkConnectivityServiceConnectionPolicyEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceConnectionPolicy: %s", err)
	}
	if err := d.Set("psc_connections", flattenNetworkConnectivityServiceConnectionPolicyPscConnections(res["pscConnections"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceConnectionPolicy: %s", err)
	}
	if err := d.Set("infrastructure", flattenNetworkConnectivityServiceConnectionPolicyInfrastructure(res["infrastructure"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceConnectionPolicy: %s", err)
	}
	if err := d.Set("labels", flattenNetworkConnectivityServiceConnectionPolicyLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceConnectionPolicy: %s", err)
	}

	return nil
}

func resourceNetworkConnectivityServiceConnectionPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceConnectionPolicy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkConnectivityServiceConnectionPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	pscConfigProp, err := expandNetworkConnectivityServiceConnectionPolicyPscConfig(d.Get("psc_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("psc_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, pscConfigProp)) {
		obj["pscConfig"] = pscConfigProp
	}
	etagProp, err := expandNetworkConnectivityServiceConnectionPolicyEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	labelsProp, err := expandNetworkConnectivityServiceConnectionPolicyLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	obj, err = resourceNetworkConnectivityServiceConnectionPolicyUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/{{location}}/serviceConnectionPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ServiceConnectionPolicy %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("psc_config") {
		updateMask = append(updateMask, "pscConfig")
	}

	if d.HasChange("etag") {
		updateMask = append(updateMask, "etag")
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
		return fmt.Errorf("Error updating ServiceConnectionPolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ServiceConnectionPolicy %q: %#v", d.Id(), res)
	}

	err = NetworkConnectivityOperationWaitTime(
		config, res, project, "Updating ServiceConnectionPolicy", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceNetworkConnectivityServiceConnectionPolicyRead(d, meta)
}

func resourceNetworkConnectivityServiceConnectionPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceConnectionPolicy: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/{{location}}/serviceConnectionPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ServiceConnectionPolicy %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "ServiceConnectionPolicy")
	}

	err = NetworkConnectivityOperationWaitTime(
		config, res, project, "Deleting ServiceConnectionPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ServiceConnectionPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceNetworkConnectivityServiceConnectionPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/serviceConnectionPolicies/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/serviceConnectionPolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetworkConnectivityServiceConnectionPolicyCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyServiceClass(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyPscConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["subnetworks"] =
		flattenNetworkConnectivityServiceConnectionPolicyPscConfigSubnetworks(original["subnetworks"], d, config)
	transformed["limit"] =
		flattenNetworkConnectivityServiceConnectionPolicyPscConfigLimit(original["limit"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkConnectivityServiceConnectionPolicyPscConfigSubnetworks(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyPscConfigLimit(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyPscConnections(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"state":                    flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsState(original["state"], d, config),
			"consumer_forwarding_rule": flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsConsumerForwardingRule(original["consumerForwardingRule"], d, config),
			"consumer_address":         flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsConsumerAddress(original["consumerAddress"], d, config),
			"error_type":               flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorType(original["errorType"], d, config),
			"error":                    flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsError(original["error"], d, config),
			"gce_operation":            flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsGceOperation(original["gceOperation"], d, config),
			"consumer_target_project":  flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsConsumerTargetProject(original["consumerTargetProject"], d, config),
			"psc_connection_id":        flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsPscConnectionId(original["pscConnectionId"], d, config),
			"error_info":               flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorInfo(original["errorInfo"], d, config),
		})
	}
	return transformed
}
func flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsConsumerForwardingRule(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsConsumerAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsError(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["message"] =
		flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorMessage(original["message"], d, config)
	transformed["code"] =
		flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorCode(original["code"], d, config)
	transformed["details"] =
		flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorDetails(original["details"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorCode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorDetails(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsGceOperation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsConsumerTargetProject(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsPscConnectionId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorInfo(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["reason"] =
		flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorInfoReason(original["reason"], d, config)
	transformed["domain"] =
		flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorInfoDomain(original["domain"], d, config)
	transformed["metadata"] =
		flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorInfoMetadata(original["metadata"], d, config)
	return []interface{}{transformed}
}
func flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorInfoReason(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorInfoDomain(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyPscConnectionsErrorInfoMetadata(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyInfrastructure(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetworkConnectivityServiceConnectionPolicyLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetworkConnectivityServiceConnectionPolicyServiceClass(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityServiceConnectionPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityServiceConnectionPolicyNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityServiceConnectionPolicyPscConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSubnetworks, err := expandNetworkConnectivityServiceConnectionPolicyPscConfigSubnetworks(original["subnetworks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubnetworks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subnetworks"] = transformedSubnetworks
	}

	transformedLimit, err := expandNetworkConnectivityServiceConnectionPolicyPscConfigLimit(original["limit"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLimit); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["limit"] = transformedLimit
	}

	return transformed, nil
}

func expandNetworkConnectivityServiceConnectionPolicyPscConfigSubnetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityServiceConnectionPolicyPscConfigLimit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityServiceConnectionPolicyEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityServiceConnectionPolicyLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func resourceNetworkConnectivityServiceConnectionPolicyUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	obj["network"] = d.Get("network").(string)
	return obj, nil
}
