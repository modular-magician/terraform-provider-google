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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComputeHttpsHealthCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeHttpsHealthCheckCreate,
		Read:   resourceComputeHttpsHealthCheckRead,
		Update: resourceComputeHttpsHealthCheckUpdate,
		Delete: resourceComputeHttpsHealthCheckDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeHttpsHealthCheckImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035.  Specifically, the name must be 1-63 characters long and
match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
the first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the
last character, which cannot be a dash.`,
			},
			"check_interval_sec": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `How often (in seconds) to send a health check. The default value is 5
seconds.`,
				Default: 5,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `An optional description of this resource. Provide this property when
you create the resource.`,
			},
			"healthy_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `A so-far unhealthy instance will be marked healthy after this many
consecutive successes. The default value is 2.`,
				Default: 2,
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The value of the host header in the HTTPS health check request. If
left empty (default value), the public IP on behalf of which this
health check is performed will be used.`,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `The TCP port number for the HTTPS health check request.
The default value is 443.`,
				Default: 443,
			},
			"request_path": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The request path of the HTTPS health check request.
The default value is /.`,
				Default: "/",
			},
			"timeout_sec": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `How long (in seconds) to wait before claiming failure.
The default value is 5 seconds.  It is invalid for timeoutSec to have
greater value than checkIntervalSec.`,
				Default: 5,
			},
			"unhealthy_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `A so-far healthy instance will be marked unhealthy after this many
consecutive failures. The default value is 2.`,
				Default: 2,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
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
		UseJSONNumber: true,
	}
}

func resourceComputeHttpsHealthCheckCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	checkIntervalSecProp, err := expandComputeHttpsHealthCheckCheckIntervalSec(d.Get("check_interval_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("check_interval_sec"); !isEmptyValue(reflect.ValueOf(checkIntervalSecProp)) && (ok || !reflect.DeepEqual(v, checkIntervalSecProp)) {
		obj["checkIntervalSec"] = checkIntervalSecProp
	}
	descriptionProp, err := expandComputeHttpsHealthCheckDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	healthyThresholdProp, err := expandComputeHttpsHealthCheckHealthyThreshold(d.Get("healthy_threshold"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("healthy_threshold"); !isEmptyValue(reflect.ValueOf(healthyThresholdProp)) && (ok || !reflect.DeepEqual(v, healthyThresholdProp)) {
		obj["healthyThreshold"] = healthyThresholdProp
	}
	hostProp, err := expandComputeHttpsHealthCheckHost(d.Get("host"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("host"); !isEmptyValue(reflect.ValueOf(hostProp)) && (ok || !reflect.DeepEqual(v, hostProp)) {
		obj["host"] = hostProp
	}
	nameProp, err := expandComputeHttpsHealthCheckName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	portProp, err := expandComputeHttpsHealthCheckPort(d.Get("port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port"); !isEmptyValue(reflect.ValueOf(portProp)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}
	requestPathProp, err := expandComputeHttpsHealthCheckRequestPath(d.Get("request_path"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("request_path"); !isEmptyValue(reflect.ValueOf(requestPathProp)) && (ok || !reflect.DeepEqual(v, requestPathProp)) {
		obj["requestPath"] = requestPathProp
	}
	timeoutSecProp, err := expandComputeHttpsHealthCheckTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout_sec"); !isEmptyValue(reflect.ValueOf(timeoutSecProp)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}
	unhealthyThresholdProp, err := expandComputeHttpsHealthCheckUnhealthyThreshold(d.Get("unhealthy_threshold"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("unhealthy_threshold"); !isEmptyValue(reflect.ValueOf(unhealthyThresholdProp)) && (ok || !reflect.DeepEqual(v, unhealthyThresholdProp)) {
		obj["unhealthyThreshold"] = unhealthyThresholdProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/httpsHealthChecks")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new HttpsHealthCheck: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for HttpsHealthCheck: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating HttpsHealthCheck: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/global/httpsHealthChecks/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating HttpsHealthCheck", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create HttpsHealthCheck: %s", err)
	}

	log.Printf("[DEBUG] Finished creating HttpsHealthCheck %q: %#v", d.Id(), res)

	return resourceComputeHttpsHealthCheckRead(d, meta)
}

func resourceComputeHttpsHealthCheckRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/httpsHealthChecks/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for HttpsHealthCheck: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeHttpsHealthCheck %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}

	if err := d.Set("check_interval_sec", flattenComputeHttpsHealthCheckCheckIntervalSec(res["checkIntervalSec"], d, config)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeHttpsHealthCheckCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("description", flattenComputeHttpsHealthCheckDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("healthy_threshold", flattenComputeHttpsHealthCheckHealthyThreshold(res["healthyThreshold"], d, config)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("host", flattenComputeHttpsHealthCheckHost(res["host"], d, config)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("name", flattenComputeHttpsHealthCheckName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("port", flattenComputeHttpsHealthCheckPort(res["port"], d, config)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("request_path", flattenComputeHttpsHealthCheckRequestPath(res["requestPath"], d, config)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("timeout_sec", flattenComputeHttpsHealthCheckTimeoutSec(res["timeoutSec"], d, config)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("unhealthy_threshold", flattenComputeHttpsHealthCheckUnhealthyThreshold(res["unhealthyThreshold"], d, config)); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading HttpsHealthCheck: %s", err)
	}

	return nil
}

func resourceComputeHttpsHealthCheckUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for HttpsHealthCheck: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	checkIntervalSecProp, err := expandComputeHttpsHealthCheckCheckIntervalSec(d.Get("check_interval_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("check_interval_sec"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, checkIntervalSecProp)) {
		obj["checkIntervalSec"] = checkIntervalSecProp
	}
	descriptionProp, err := expandComputeHttpsHealthCheckDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	healthyThresholdProp, err := expandComputeHttpsHealthCheckHealthyThreshold(d.Get("healthy_threshold"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("healthy_threshold"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, healthyThresholdProp)) {
		obj["healthyThreshold"] = healthyThresholdProp
	}
	hostProp, err := expandComputeHttpsHealthCheckHost(d.Get("host"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("host"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, hostProp)) {
		obj["host"] = hostProp
	}
	nameProp, err := expandComputeHttpsHealthCheckName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	portProp, err := expandComputeHttpsHealthCheckPort(d.Get("port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}
	requestPathProp, err := expandComputeHttpsHealthCheckRequestPath(d.Get("request_path"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("request_path"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, requestPathProp)) {
		obj["requestPath"] = requestPathProp
	}
	timeoutSecProp, err := expandComputeHttpsHealthCheckTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout_sec"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}
	unhealthyThresholdProp, err := expandComputeHttpsHealthCheckUnhealthyThreshold(d.Get("unhealthy_threshold"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("unhealthy_threshold"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, unhealthyThresholdProp)) {
		obj["unhealthyThreshold"] = unhealthyThresholdProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/httpsHealthChecks/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating HttpsHealthCheck %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PUT", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating HttpsHealthCheck %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating HttpsHealthCheck %q: %#v", d.Id(), res)
	}

	err = computeOperationWaitTime(
		config, res, project, "Updating HttpsHealthCheck", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputeHttpsHealthCheckRead(d, meta)
}

func resourceComputeHttpsHealthCheckDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for HttpsHealthCheck: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/httpsHealthChecks/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting HttpsHealthCheck %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "HttpsHealthCheck")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting HttpsHealthCheck", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting HttpsHealthCheck %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeHttpsHealthCheckImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/httpsHealthChecks/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/global/httpsHealthChecks/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeHttpsHealthCheckCheckIntervalSec(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
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

func flattenComputeHttpsHealthCheckCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeHttpsHealthCheckDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeHttpsHealthCheckHealthyThreshold(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
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

func flattenComputeHttpsHealthCheckHost(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeHttpsHealthCheckName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeHttpsHealthCheckPort(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
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

func flattenComputeHttpsHealthCheckRequestPath(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeHttpsHealthCheckTimeoutSec(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
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

func flattenComputeHttpsHealthCheckUnhealthyThreshold(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
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

func expandComputeHttpsHealthCheckCheckIntervalSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckHealthyThreshold(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckHost(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckRequestPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckTimeoutSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHttpsHealthCheckUnhealthyThreshold(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
