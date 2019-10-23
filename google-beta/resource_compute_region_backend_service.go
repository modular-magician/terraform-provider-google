// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"google.golang.org/api/compute/v1"
)

func migrateStateNoop(v int, is *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	return is, nil
}

func resourceComputeRegionBackendService() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionBackendServiceCreate,
		Read:   resourceComputeRegionBackendServiceRead,
		Update: resourceComputeRegionBackendServiceUpdate,
		Delete: resourceComputeRegionBackendServiceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionBackendServiceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		SchemaVersion: 1,
		MigrateState:  migrateStateNoop,

		Schema: map[string]*schema.Schema{
			"health_checks": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: selfLinkRelativePathHash,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"backend": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     computeRegionBackendServiceBackendSchema(),
				Set:      resourceGoogleComputeBackendServiceBackendHash,
			},
			"connection_draining_timeout_sec": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"failover_policy": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable_connection_drain_on_failover": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"drop_traffic_if_unhealthy": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"failover_ratio": {
							Type:     schema.TypeFloat,
							Optional: true,
						},
					},
				},
			},
			"load_balancing_scheme": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"INTERNAL", ""}, false),
				Default:      "INTERNAL",
			},
			"protocol": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"HTTP", "HTTPS", "HTTP2", "SSL", "TCP", "UDP", ""}, false),
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"session_affinity": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "CLIENT_IP", "CLIENT_IP_PROTO", "CLIENT_IP_PORT_PROTO", ""}, false),
			},
			"timeout_sec": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func computeRegionBackendServiceBackendSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"failover": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"group": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkRelativePaths,
			},
		},
	}
}

func resourceComputeRegionBackendServiceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionBackendServiceName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	healthChecksProp, err := expandComputeRegionBackendServiceHealthChecks(d.Get("health_checks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("health_checks"); !isEmptyValue(reflect.ValueOf(healthChecksProp)) && (ok || !reflect.DeepEqual(v, healthChecksProp)) {
		obj["healthChecks"] = healthChecksProp
	}
	backendsProp, err := expandComputeRegionBackendServiceBackend(d.Get("backend"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backend"); !isEmptyValue(reflect.ValueOf(backendsProp)) && (ok || !reflect.DeepEqual(v, backendsProp)) {
		obj["backends"] = backendsProp
	}
	descriptionProp, err := expandComputeRegionBackendServiceDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	failoverPolicyProp, err := expandComputeRegionBackendServiceFailoverPolicy(d.Get("failover_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("failover_policy"); !isEmptyValue(reflect.ValueOf(failoverPolicyProp)) && (ok || !reflect.DeepEqual(v, failoverPolicyProp)) {
		obj["failoverPolicy"] = failoverPolicyProp
	}
	fingerprintProp, err := expandComputeRegionBackendServiceFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	protocolProp, err := expandComputeRegionBackendServiceProtocol(d.Get("protocol"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("protocol"); !isEmptyValue(reflect.ValueOf(protocolProp)) && (ok || !reflect.DeepEqual(v, protocolProp)) {
		obj["protocol"] = protocolProp
	}
	sessionAffinityProp, err := expandComputeRegionBackendServiceSessionAffinity(d.Get("session_affinity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("session_affinity"); !isEmptyValue(reflect.ValueOf(sessionAffinityProp)) && (ok || !reflect.DeepEqual(v, sessionAffinityProp)) {
		obj["sessionAffinity"] = sessionAffinityProp
	}
	regionProp, err := expandComputeRegionBackendServiceRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}
	timeoutSecProp, err := expandComputeRegionBackendServiceTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout_sec"); !isEmptyValue(reflect.ValueOf(timeoutSecProp)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}
	connectionDrainingProp, err := expandComputeRegionBackendServiceConnectionDraining(nil, d, config)
	if err != nil {
		return err
	} else if !isEmptyValue(reflect.ValueOf(connectionDrainingProp)) {
		obj["connectionDraining"] = connectionDrainingProp
	}
	loadBalancingSchemeProp, err := expandComputeRegionBackendServiceLoadBalancingScheme(d.Get("load_balancing_scheme"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("load_balancing_scheme"); !isEmptyValue(reflect.ValueOf(loadBalancingSchemeProp)) && (ok || !reflect.DeepEqual(v, loadBalancingSchemeProp)) {
		obj["loadBalancingScheme"] = loadBalancingSchemeProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/backendServices")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionBackendService: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating RegionBackendService: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating RegionBackendService",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionBackendService: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating RegionBackendService %q: %#v", d.Id(), res)

	return resourceComputeRegionBackendServiceRead(d, meta)
}

func resourceComputeRegionBackendServiceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRegionBackendService %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}

	if err := d.Set("name", flattenComputeRegionBackendServiceName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("health_checks", flattenComputeRegionBackendServiceHealthChecks(res["healthChecks"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("backend", flattenComputeRegionBackendServiceBackend(res["backends"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("description", flattenComputeRegionBackendServiceDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("failover_policy", flattenComputeRegionBackendServiceFailoverPolicy(res["failoverPolicy"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("fingerprint", flattenComputeRegionBackendServiceFingerprint(res["fingerprint"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("protocol", flattenComputeRegionBackendServiceProtocol(res["protocol"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("session_affinity", flattenComputeRegionBackendServiceSessionAffinity(res["sessionAffinity"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("region", flattenComputeRegionBackendServiceRegion(res["region"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("timeout_sec", flattenComputeRegionBackendServiceTimeoutSec(res["timeoutSec"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if flattenedProp := flattenComputeRegionBackendServiceConnectionDraining(res["connectionDraining"], d); flattenedProp != nil {
		casted := flattenedProp.([]interface{})[0]
		if casted != nil {
			for k, v := range casted.(map[string]interface{}) {
				d.Set(k, v)
			}
		}
	}
	if err := d.Set("load_balancing_scheme", flattenComputeRegionBackendServiceLoadBalancingScheme(res["loadBalancingScheme"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}

	return nil
}

func resourceComputeRegionBackendServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionBackendServiceName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	healthChecksProp, err := expandComputeRegionBackendServiceHealthChecks(d.Get("health_checks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("health_checks"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, healthChecksProp)) {
		obj["healthChecks"] = healthChecksProp
	}
	backendsProp, err := expandComputeRegionBackendServiceBackend(d.Get("backend"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backend"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, backendsProp)) {
		obj["backends"] = backendsProp
	}
	descriptionProp, err := expandComputeRegionBackendServiceDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	failoverPolicyProp, err := expandComputeRegionBackendServiceFailoverPolicy(d.Get("failover_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("failover_policy"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, failoverPolicyProp)) {
		obj["failoverPolicy"] = failoverPolicyProp
	}
	fingerprintProp, err := expandComputeRegionBackendServiceFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	protocolProp, err := expandComputeRegionBackendServiceProtocol(d.Get("protocol"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("protocol"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, protocolProp)) {
		obj["protocol"] = protocolProp
	}
	sessionAffinityProp, err := expandComputeRegionBackendServiceSessionAffinity(d.Get("session_affinity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("session_affinity"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sessionAffinityProp)) {
		obj["sessionAffinity"] = sessionAffinityProp
	}
	regionProp, err := expandComputeRegionBackendServiceRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}
	timeoutSecProp, err := expandComputeRegionBackendServiceTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout_sec"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}
	connectionDrainingProp, err := expandComputeRegionBackendServiceConnectionDraining(nil, d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("connection_draining"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, connectionDrainingProp)) {
		obj["connectionDraining"] = connectionDrainingProp
	}
	loadBalancingSchemeProp, err := expandComputeRegionBackendServiceLoadBalancingScheme(d.Get("load_balancing_scheme"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("load_balancing_scheme"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, loadBalancingSchemeProp)) {
		obj["loadBalancingScheme"] = loadBalancingSchemeProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RegionBackendService %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating RegionBackendService %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Updating RegionBackendService",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceComputeRegionBackendServiceRead(d, meta)
}

func resourceComputeRegionBackendServiceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting RegionBackendService %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "RegionBackendService")
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting RegionBackendService",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RegionBackendService %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionBackendServiceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/backendServices/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRegionBackendServiceName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceHealthChecks(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return convertAndMapStringArr(v.([]interface{}), ConvertSelfLinkToV1)
}

func flattenComputeRegionBackendServiceBackend(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(resourceGoogleComputeBackendServiceBackendHash, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"description": flattenComputeRegionBackendServiceBackendDescription(original["description"], d),
			"group":       flattenComputeRegionBackendServiceBackendGroup(original["group"], d),
			"failover":    flattenComputeRegionBackendServiceBackendFailover(original["failover"], d),
		})
	}
	return transformed
}
func flattenComputeRegionBackendServiceBackendDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceBackendGroup(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRegionBackendServiceBackendFailover(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceFailoverPolicy(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["disable_connection_drain_on_failover"] =
		flattenComputeRegionBackendServiceFailoverPolicyDisableConnectionDrainOnFailover(original["disableConnectionDrainOnFailover"], d)
	transformed["drop_traffic_if_unhealthy"] =
		flattenComputeRegionBackendServiceFailoverPolicyDropTrafficIfUnhealthy(original["dropTrafficIfUnhealthy"], d)
	transformed["failover_ratio"] =
		flattenComputeRegionBackendServiceFailoverPolicyFailoverRatio(original["failoverRatio"], d)
	return []interface{}{transformed}
}
func flattenComputeRegionBackendServiceFailoverPolicyDisableConnectionDrainOnFailover(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceFailoverPolicyDropTrafficIfUnhealthy(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceFailoverPolicyFailoverRatio(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceFingerprint(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceProtocol(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceSessionAffinity(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceRegion(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenComputeRegionBackendServiceTimeoutSec(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionBackendServiceConnectionDraining(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["connection_draining_timeout_sec"] =
		flattenComputeRegionBackendServiceConnectionDrainingConnectionDrainingTimeoutSec(original["drainingTimeoutSec"], d)
	return []interface{}{transformed}
}
func flattenComputeRegionBackendServiceConnectionDrainingConnectionDrainingTimeoutSec(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionBackendServiceLoadBalancingScheme(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandComputeRegionBackendServiceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceHealthChecks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeRegionBackendServiceBackend(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDescription, err := expandComputeRegionBackendServiceBackendDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedGroup, err := expandComputeRegionBackendServiceBackendGroup(original["group"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGroup); val.IsValid() && !isEmptyValue(val) {
			transformed["group"] = transformedGroup
		}

		transformedFailover, err := expandComputeRegionBackendServiceBackendFailover(original["failover"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFailover); val.IsValid() && !isEmptyValue(val) {
			transformed["failover"] = transformedFailover
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRegionBackendServiceBackendDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceBackendGroup(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceBackendFailover(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceFailoverPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDisableConnectionDrainOnFailover, err := expandComputeRegionBackendServiceFailoverPolicyDisableConnectionDrainOnFailover(original["disable_connection_drain_on_failover"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisableConnectionDrainOnFailover); val.IsValid() && !isEmptyValue(val) {
		transformed["disableConnectionDrainOnFailover"] = transformedDisableConnectionDrainOnFailover
	}

	transformedDropTrafficIfUnhealthy, err := expandComputeRegionBackendServiceFailoverPolicyDropTrafficIfUnhealthy(original["drop_traffic_if_unhealthy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDropTrafficIfUnhealthy); val.IsValid() && !isEmptyValue(val) {
		transformed["dropTrafficIfUnhealthy"] = transformedDropTrafficIfUnhealthy
	}

	transformedFailoverRatio, err := expandComputeRegionBackendServiceFailoverPolicyFailoverRatio(original["failover_ratio"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFailoverRatio); val.IsValid() && !isEmptyValue(val) {
		transformed["failoverRatio"] = transformedFailoverRatio
	}

	return transformed, nil
}

func expandComputeRegionBackendServiceFailoverPolicyDisableConnectionDrainOnFailover(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceFailoverPolicyDropTrafficIfUnhealthy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceFailoverPolicyFailoverRatio(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceSessionAffinity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRegionBackendServiceTimeoutSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceConnectionDraining(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedConnectionDrainingTimeoutSec, err := expandComputeRegionBackendServiceConnectionDrainingConnectionDrainingTimeoutSec(d.Get("connection_draining_timeout_sec"), d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["drainingTimeoutSec"] = transformedConnectionDrainingTimeoutSec
	}

	return transformed, nil
}

func expandComputeRegionBackendServiceConnectionDrainingConnectionDrainingTimeoutSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceLoadBalancingScheme(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
