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

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"google.golang.org/api/compute/v1"
)

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
			Create: schema.DefaultTimeout(240 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		SchemaVersion: 1,

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
				ValidateFunc: validation.StringInSlice([]string{"TCP", "UDP", ""}, false),
			},
			"region": {
				Type:             schema.TypeString,
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
	connectionDrainingProp, err := expandComputeRegionBackendServiceConnectionDraining(d, config)
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

	obj, err = resourceComputeRegionBackendServiceEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/backendServices")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionBackendService: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating RegionBackendService: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
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

	// security_policy isn't set by Create / Update
	if v, ok := d.GetOk("security_policy"); ok {
		pol, err := ParseSecurityPolicyFieldValue(v.(string), d, config)
		if err != nil {
			return errwrap.Wrapf("Error parsing Backend Service security policy: {{err}}", err)
		}
		op, err := config.clientCompute.BackendServices.SetSecurityPolicy(
			project, obj["name"].(string), &compute.SecurityPolicyReference{
				SecurityPolicy: pol.RelativeLink(),
			}).Do()
		if err != nil {
			return errwrap.Wrapf("Error setting Backend Service security policy: {{err}}", err)
		}
		waitErr := computeSharedOperationWait(config.clientCompute, op, project, "Setting Backend Service Security Policy")
		if waitErr != nil {
			return waitErr
		}
	}

	return resourceComputeRegionBackendServiceRead(d, meta)
}

func resourceComputeRegionBackendServiceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRegionBackendService %q", d.Id()))
	}

	res, err = resourceComputeRegionBackendServiceDecoder(d, meta, res)
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
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
	if v, ok := res["connectionDraining"].(map[string]interface{}); res["connectionDraining"] != nil && ok {
		if err := d.Set("connection_draining_timeout_sec", flattenComputeRegionBackendServiceConnectionDrainingConnection_draining_timeout_sec(v["drainingTimeoutSec"], d)); err != nil {
			return fmt.Errorf("Error reading RegionBackendService: %s", err)
		}
	} else {
		d.Set("connection_draining_timeout_sec", nil)
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
	connectionDrainingProp, err := expandComputeRegionBackendServiceConnectionDraining(d, config)
	if err != nil {
		return err
	} else if !isEmptyValue(reflect.ValueOf(connectionDrainingProp)) {
		obj["connectionDraining"] = connectionDrainingProp
	}
	loadBalancingSchemeProp, err := expandComputeRegionBackendServiceLoadBalancingScheme(d.Get("load_balancing_scheme"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("load_balancing_scheme"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, loadBalancingSchemeProp)) {
		obj["loadBalancingScheme"] = loadBalancingSchemeProp
	}

	obj, err = resourceComputeRegionBackendServiceEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RegionBackendService %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PUT", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating RegionBackendService %q: %s", d.Id(), err)
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
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

	// security_policy isn't set by Create / Update
	if v, ok := d.GetOk("security_policy"); ok {
		pol, err := ParseSecurityPolicyFieldValue(v.(string), d, config)
		if err != nil {
			return errwrap.Wrapf("Error parsing Backend Service security policy: {{err}}", err)
		}
		op, err := config.clientCompute.BackendServices.SetSecurityPolicy(
			project, obj["name"].(string), &compute.SecurityPolicyReference{
				SecurityPolicy: pol.RelativeLink(),
			}).Do()
		if err != nil {
			return errwrap.Wrapf("Error setting Backend Service security policy: {{err}}", err)
		}
		waitErr := computeSharedOperationWait(config.clientCompute, op, project, "Setting Backend Service Security Policy")
		if waitErr != nil {
			return waitErr
		}
	}
	return resourceComputeRegionBackendServiceRead(d, meta)
}

func resourceComputeRegionBackendServiceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting RegionBackendService %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "RegionBackendService")
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
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
	if err := parseImportId([]string{"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/backendServices/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
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

func flattenComputeRegionBackendServiceDescription(v interface{}, d *schema.ResourceData) interface{} {
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
	return ConvertSelfLinkToV1(v.(string))
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

func flattenComputeRegionBackendServiceConnectionDrainingConnection_draining_timeout_sec(v interface{}, d *schema.ResourceData) interface{} {
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

func expandComputeRegionBackendServiceDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
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

func expandComputeRegionBackendServiceConnectionDraining(d TerraformResourceData, config *Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	// Note that nesting flattened objects won't work because we don't handle them properly here.
	transformedConnection_draining_timeout_sec, err := expandComputeRegionBackendServiceConnectionDrainingConnection_draining_timeout_sec(d.Get("connection_draining_timeout_sec"), d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["drainingTimeoutSec"] = transformedConnection_draining_timeout_sec
	}

	return transformed, nil
}

func expandComputeRegionBackendServiceConnectionDrainingConnection_draining_timeout_sec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceLoadBalancingScheme(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceComputeRegionBackendServiceEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// The BackendService API's Update / PUT API is badly formed and behaves like
	// a PATCH field for at least IAP. When sent a `null` `iap` field, the API
	// doesn't disable an existing field. To work around this, we need to emulate
	// the old Terraform behaviour of always sending the block (at both update and
	// create), and force sending each subfield as empty when the block isn't
	// present in config.

	iapVal := obj["iap"]
	if iapVal == nil {
		data := map[string]interface{}{}
		data["enabled"] = false
		data["oauth2ClientId"] = ""
		data["oauth2ClientSecret"] = ""
		obj["iap"] = data
	} else {
		iap := iapVal.(map[string]interface{})
		iap["enabled"] = true
		obj["iap"] = iap
	}

	return obj, nil
}

func resourceComputeRegionBackendServiceDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	// We need to pretend IAP isn't there if it's disabled for Terraform to maintain
	// BC behaviour with the handwritten resource.
	v, ok := res["iap"]
	if !ok || v == nil {
		delete(res, "iap")
		return res, nil
	}
	m := v.(map[string]interface{})
	if ok && m["enabled"] == false {
		delete(res, "iap")
	}

	return res, nil
}
