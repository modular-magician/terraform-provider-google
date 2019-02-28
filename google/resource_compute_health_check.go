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
)

func resourceComputeHealthCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeHealthCheckCreate,
		Read:   resourceComputeHealthCheckRead,
		Update: resourceComputeHealthCheckUpdate,
		Delete: resourceComputeHealthCheckDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeHealthCheckImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"check_interval_sec": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"healthy_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"http_health_check": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  80,
						},
						"proxy_header": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"NONE", "PROXY_V1", ""}, false),
							Default:      "NONE",
						},
						"request_path": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "/",
						},
						"response": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				ConflictsWith: []string{"https_health_check", "tcp_health_check", "ssl_health_check"},
			},
			"https_health_check": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  443,
						},
						"proxy_header": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"NONE", "PROXY_V1", ""}, false),
							Default:      "NONE",
						},
						"request_path": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "/",
						},
						"response": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				ConflictsWith: []string{"http_health_check", "tcp_health_check", "ssl_health_check"},
			},
			"ssl_health_check": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  443,
						},
						"proxy_header": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"NONE", "PROXY_V1", ""}, false),
							Default:      "NONE",
						},
						"request": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"response": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				ConflictsWith: []string{"http_health_check", "https_health_check", "tcp_health_check"},
			},
			"tcp_health_check": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  80,
						},
						"proxy_header": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"NONE", "PROXY_V1", ""}, false),
							Default:      "NONE",
						},
						"request": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"response": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				ConflictsWith: []string{"http_health_check", "https_health_check", "ssl_health_check"},
			},
			"timeout_sec": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"unhealthy_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
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

func resourceComputeHealthCheckCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	checkIntervalSecProp, err := expandComputeHealthCheckCheckIntervalSec(d.Get("check_interval_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("check_interval_sec"); !isEmptyValue(reflect.ValueOf(checkIntervalSecProp)) && (ok || !reflect.DeepEqual(v, checkIntervalSecProp)) {
		obj["checkIntervalSec"] = checkIntervalSecProp
	}
	descriptionProp, err := expandComputeHealthCheckDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	healthyThresholdProp, err := expandComputeHealthCheckHealthyThreshold(d.Get("healthy_threshold"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("healthy_threshold"); !isEmptyValue(reflect.ValueOf(healthyThresholdProp)) && (ok || !reflect.DeepEqual(v, healthyThresholdProp)) {
		obj["healthyThreshold"] = healthyThresholdProp
	}
	nameProp, err := expandComputeHealthCheckName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	timeoutSecProp, err := expandComputeHealthCheckTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout_sec"); !isEmptyValue(reflect.ValueOf(timeoutSecProp)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}
	unhealthyThresholdProp, err := expandComputeHealthCheckUnhealthyThreshold(d.Get("unhealthy_threshold"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("unhealthy_threshold"); !isEmptyValue(reflect.ValueOf(unhealthyThresholdProp)) && (ok || !reflect.DeepEqual(v, unhealthyThresholdProp)) {
		obj["unhealthyThreshold"] = unhealthyThresholdProp
	}
	httpHealthCheckProp, err := expandComputeHealthCheckHttpHealthCheck(d.Get("http_health_check"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("http_health_check"); !isEmptyValue(reflect.ValueOf(httpHealthCheckProp)) && (ok || !reflect.DeepEqual(v, httpHealthCheckProp)) {
		obj["httpHealthCheck"] = httpHealthCheckProp
	}
	httpsHealthCheckProp, err := expandComputeHealthCheckHttpsHealthCheck(d.Get("https_health_check"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("https_health_check"); !isEmptyValue(reflect.ValueOf(httpsHealthCheckProp)) && (ok || !reflect.DeepEqual(v, httpsHealthCheckProp)) {
		obj["httpsHealthCheck"] = httpsHealthCheckProp
	}
	tcpHealthCheckProp, err := expandComputeHealthCheckTcpHealthCheck(d.Get("tcp_health_check"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tcp_health_check"); !isEmptyValue(reflect.ValueOf(tcpHealthCheckProp)) && (ok || !reflect.DeepEqual(v, tcpHealthCheckProp)) {
		obj["tcpHealthCheck"] = tcpHealthCheckProp
	}
	sslHealthCheckProp, err := expandComputeHealthCheckSslHealthCheck(d.Get("ssl_health_check"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ssl_health_check"); !isEmptyValue(reflect.ValueOf(sslHealthCheckProp)) && (ok || !reflect.DeepEqual(v, sslHealthCheckProp)) {
		obj["sslHealthCheck"] = sslHealthCheckProp
	}

	obj, err = resourceComputeHealthCheckEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/healthChecks")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new HealthCheck: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating HealthCheck: %s", err)
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
		config.clientCompute, op, project, "Creating HealthCheck",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create HealthCheck: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating HealthCheck %q: %#v", d.Id(), res)

	return resourceComputeHealthCheckRead(d, meta)
}

func resourceComputeHealthCheckRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/healthChecks/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeHealthCheck %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading HealthCheck: %s", err)
	}

	if err := d.Set("check_interval_sec", flattenComputeHealthCheckCheckIntervalSec(res["checkIntervalSec"], d)); err != nil {
		return fmt.Errorf("Error reading HealthCheck: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeHealthCheckCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading HealthCheck: %s", err)
	}
	if err := d.Set("description", flattenComputeHealthCheckDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading HealthCheck: %s", err)
	}
	if err := d.Set("healthy_threshold", flattenComputeHealthCheckHealthyThreshold(res["healthyThreshold"], d)); err != nil {
		return fmt.Errorf("Error reading HealthCheck: %s", err)
	}
	if err := d.Set("name", flattenComputeHealthCheckName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading HealthCheck: %s", err)
	}
	if err := d.Set("timeout_sec", flattenComputeHealthCheckTimeoutSec(res["timeoutSec"], d)); err != nil {
		return fmt.Errorf("Error reading HealthCheck: %s", err)
	}
	if err := d.Set("unhealthy_threshold", flattenComputeHealthCheckUnhealthyThreshold(res["unhealthyThreshold"], d)); err != nil {
		return fmt.Errorf("Error reading HealthCheck: %s", err)
	}
	if err := d.Set("type", flattenComputeHealthCheckType(res["type"], d)); err != nil {
		return fmt.Errorf("Error reading HealthCheck: %s", err)
	}
	if err := d.Set("http_health_check", flattenComputeHealthCheckHttpHealthCheck(res["httpHealthCheck"], d)); err != nil {
		return fmt.Errorf("Error reading HealthCheck: %s", err)
	}
	if err := d.Set("https_health_check", flattenComputeHealthCheckHttpsHealthCheck(res["httpsHealthCheck"], d)); err != nil {
		return fmt.Errorf("Error reading HealthCheck: %s", err)
	}
	if err := d.Set("tcp_health_check", flattenComputeHealthCheckTcpHealthCheck(res["tcpHealthCheck"], d)); err != nil {
		return fmt.Errorf("Error reading HealthCheck: %s", err)
	}
	if err := d.Set("ssl_health_check", flattenComputeHealthCheckSslHealthCheck(res["sslHealthCheck"], d)); err != nil {
		return fmt.Errorf("Error reading HealthCheck: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading HealthCheck: %s", err)
	}

	return nil
}

func resourceComputeHealthCheckUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	checkIntervalSecProp, err := expandComputeHealthCheckCheckIntervalSec(d.Get("check_interval_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("check_interval_sec"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, checkIntervalSecProp)) {
		obj["checkIntervalSec"] = checkIntervalSecProp
	}
	descriptionProp, err := expandComputeHealthCheckDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	healthyThresholdProp, err := expandComputeHealthCheckHealthyThreshold(d.Get("healthy_threshold"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("healthy_threshold"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, healthyThresholdProp)) {
		obj["healthyThreshold"] = healthyThresholdProp
	}
	nameProp, err := expandComputeHealthCheckName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	timeoutSecProp, err := expandComputeHealthCheckTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout_sec"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}
	unhealthyThresholdProp, err := expandComputeHealthCheckUnhealthyThreshold(d.Get("unhealthy_threshold"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("unhealthy_threshold"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, unhealthyThresholdProp)) {
		obj["unhealthyThreshold"] = unhealthyThresholdProp
	}
	httpHealthCheckProp, err := expandComputeHealthCheckHttpHealthCheck(d.Get("http_health_check"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("http_health_check"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, httpHealthCheckProp)) {
		obj["httpHealthCheck"] = httpHealthCheckProp
	}
	httpsHealthCheckProp, err := expandComputeHealthCheckHttpsHealthCheck(d.Get("https_health_check"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("https_health_check"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, httpsHealthCheckProp)) {
		obj["httpsHealthCheck"] = httpsHealthCheckProp
	}
	tcpHealthCheckProp, err := expandComputeHealthCheckTcpHealthCheck(d.Get("tcp_health_check"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tcp_health_check"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, tcpHealthCheckProp)) {
		obj["tcpHealthCheck"] = tcpHealthCheckProp
	}
	sslHealthCheckProp, err := expandComputeHealthCheckSslHealthCheck(d.Get("ssl_health_check"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ssl_health_check"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sslHealthCheckProp)) {
		obj["sslHealthCheck"] = sslHealthCheckProp
	}

	obj, err = resourceComputeHealthCheckEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/healthChecks/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating HealthCheck %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PUT", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating HealthCheck %q: %s", d.Id(), err)
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
		config.clientCompute, op, project, "Updating HealthCheck",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceComputeHealthCheckRead(d, meta)
}

func resourceComputeHealthCheckDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/healthChecks/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting HealthCheck %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "HealthCheck")
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
		config.clientCompute, op, project, "Deleting HealthCheck",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting HealthCheck %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeHealthCheckImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{"projects/(?P<project>[^/]+)/global/healthChecks/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config); err != nil {
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

func flattenComputeHealthCheckCheckIntervalSec(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeHealthCheckCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckHealthyThreshold(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeHealthCheckName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckTimeoutSec(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeHealthCheckUnhealthyThreshold(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeHealthCheckType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckHttpHealthCheck(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["host"] =
		flattenComputeHealthCheckHttpHealthCheckHost(original["host"], d)
	transformed["request_path"] =
		flattenComputeHealthCheckHttpHealthCheckRequestPath(original["requestPath"], d)
	transformed["response"] =
		flattenComputeHealthCheckHttpHealthCheckResponse(original["response"], d)
	transformed["port"] =
		flattenComputeHealthCheckHttpHealthCheckPort(original["port"], d)
	transformed["proxy_header"] =
		flattenComputeHealthCheckHttpHealthCheckProxyHeader(original["proxyHeader"], d)
	return []interface{}{transformed}
}
func flattenComputeHealthCheckHttpHealthCheckHost(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckHttpHealthCheckRequestPath(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckHttpHealthCheckResponse(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckHttpHealthCheckPort(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeHealthCheckHttpHealthCheckProxyHeader(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckHttpsHealthCheck(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["host"] =
		flattenComputeHealthCheckHttpsHealthCheckHost(original["host"], d)
	transformed["request_path"] =
		flattenComputeHealthCheckHttpsHealthCheckRequestPath(original["requestPath"], d)
	transformed["response"] =
		flattenComputeHealthCheckHttpsHealthCheckResponse(original["response"], d)
	transformed["port"] =
		flattenComputeHealthCheckHttpsHealthCheckPort(original["port"], d)
	transformed["proxy_header"] =
		flattenComputeHealthCheckHttpsHealthCheckProxyHeader(original["proxyHeader"], d)
	return []interface{}{transformed}
}
func flattenComputeHealthCheckHttpsHealthCheckHost(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckHttpsHealthCheckRequestPath(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckHttpsHealthCheckResponse(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckHttpsHealthCheckPort(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeHealthCheckHttpsHealthCheckProxyHeader(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckTcpHealthCheck(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["request"] =
		flattenComputeHealthCheckTcpHealthCheckRequest(original["request"], d)
	transformed["response"] =
		flattenComputeHealthCheckTcpHealthCheckResponse(original["response"], d)
	transformed["port"] =
		flattenComputeHealthCheckTcpHealthCheckPort(original["port"], d)
	transformed["proxy_header"] =
		flattenComputeHealthCheckTcpHealthCheckProxyHeader(original["proxyHeader"], d)
	return []interface{}{transformed}
}
func flattenComputeHealthCheckTcpHealthCheckRequest(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckTcpHealthCheckResponse(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckTcpHealthCheckPort(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeHealthCheckTcpHealthCheckProxyHeader(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckSslHealthCheck(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["request"] =
		flattenComputeHealthCheckSslHealthCheckRequest(original["request"], d)
	transformed["response"] =
		flattenComputeHealthCheckSslHealthCheckResponse(original["response"], d)
	transformed["port"] =
		flattenComputeHealthCheckSslHealthCheckPort(original["port"], d)
	transformed["proxy_header"] =
		flattenComputeHealthCheckSslHealthCheckProxyHeader(original["proxyHeader"], d)
	return []interface{}{transformed}
}
func flattenComputeHealthCheckSslHealthCheckRequest(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckSslHealthCheckResponse(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeHealthCheckSslHealthCheckPort(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeHealthCheckSslHealthCheckProxyHeader(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandComputeHealthCheckCheckIntervalSec(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHealthyThreshold(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTimeoutSec(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckUnhealthyThreshold(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheck(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHost, err := expandComputeHealthCheckHttpHealthCheckHost(original["host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !isEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	transformedRequestPath, err := expandComputeHealthCheckHttpHealthCheckRequestPath(original["request_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestPath); val.IsValid() && !isEmptyValue(val) {
		transformed["requestPath"] = transformedRequestPath
	}

	transformedResponse, err := expandComputeHealthCheckHttpHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !isEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeHealthCheckHttpHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedProxyHeader, err := expandComputeHealthCheckHttpHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !isEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	return transformed, nil
}

func expandComputeHealthCheckHttpHealthCheckHost(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheckRequestPath(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheckResponse(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheckPort(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheckProxyHeader(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheck(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHost, err := expandComputeHealthCheckHttpsHealthCheckHost(original["host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !isEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	transformedRequestPath, err := expandComputeHealthCheckHttpsHealthCheckRequestPath(original["request_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestPath); val.IsValid() && !isEmptyValue(val) {
		transformed["requestPath"] = transformedRequestPath
	}

	transformedResponse, err := expandComputeHealthCheckHttpsHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !isEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeHealthCheckHttpsHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedProxyHeader, err := expandComputeHealthCheckHttpsHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !isEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	return transformed, nil
}

func expandComputeHealthCheckHttpsHealthCheckHost(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheckRequestPath(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheckResponse(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheckPort(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheckProxyHeader(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTcpHealthCheck(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRequest, err := expandComputeHealthCheckTcpHealthCheckRequest(original["request"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequest); val.IsValid() && !isEmptyValue(val) {
		transformed["request"] = transformedRequest
	}

	transformedResponse, err := expandComputeHealthCheckTcpHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !isEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeHealthCheckTcpHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedProxyHeader, err := expandComputeHealthCheckTcpHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !isEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	return transformed, nil
}

func expandComputeHealthCheckTcpHealthCheckRequest(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTcpHealthCheckResponse(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTcpHealthCheckPort(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTcpHealthCheckProxyHeader(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckSslHealthCheck(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRequest, err := expandComputeHealthCheckSslHealthCheckRequest(original["request"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequest); val.IsValid() && !isEmptyValue(val) {
		transformed["request"] = transformedRequest
	}

	transformedResponse, err := expandComputeHealthCheckSslHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !isEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeHealthCheckSslHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedProxyHeader, err := expandComputeHealthCheckSslHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !isEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	return transformed, nil
}

func expandComputeHealthCheckSslHealthCheckRequest(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckSslHealthCheckResponse(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckSslHealthCheckPort(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckSslHealthCheckProxyHeader(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceComputeHealthCheckEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	if _, ok := d.GetOk("http_health_check"); ok {
		obj["type"] = "HTTP"
		return obj, nil
	}
	if _, ok := d.GetOk("https_health_check"); ok {
		obj["type"] = "HTTPS"
		return obj, nil
	}
	if _, ok := d.GetOk("tcp_health_check"); ok {
		obj["type"] = "TCP"
		return obj, nil
	}
	if _, ok := d.GetOk("ssl_health_check"); ok {
		obj["type"] = "SSL"
		return obj, nil
	}

	return nil, fmt.Errorf("error in HealthCheck %s: No health check block specified.", d.Get("name").(string))
}
