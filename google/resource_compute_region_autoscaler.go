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
)

func resourceComputeRegionAutoscaler() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionAutoscalerCreate,
		Read:   resourceComputeRegionAutoscalerRead,
		Update: resourceComputeRegionAutoscalerUpdate,
		Delete: resourceComputeRegionAutoscalerDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionAutoscalerImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"autoscaling_policy": {
				Type:     schema.TypeList,
				Required: true,
				Description: `The configuration parameters for the autoscaling algorithm. You can
define one or more of the policies for an autoscaler: cpuUtilization,
customMetricUtilizations, and loadBalancingUtilization.

If none of these are specified, the default will be to autoscale based
on cpuUtilization to 0.6 or 60%.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_replicas": {
							Type:     schema.TypeInt,
							Required: true,
							Description: `The maximum number of instances that the autoscaler can scale up
to. This is required when creating or updating an autoscaler. The
maximum number of replicas should not be lower than minimal number
of replicas.`,
						},
						"min_replicas": {
							Type:     schema.TypeInt,
							Required: true,
							Description: `The minimum number of replicas that the autoscaler can scale down
to. This cannot be less than 0. If not provided, autoscaler will
choose a default value depending on maximum number of instances
allowed.`,
						},
						"cooldown_period": {
							Type:     schema.TypeInt,
							Optional: true,
							Description: `The number of seconds that the autoscaler should wait before it
starts collecting information from a new instance. This prevents
the autoscaler from collecting information when the instance is
initializing, during which the collected usage would not be
reliable. The default time autoscaler waits is 60 seconds.

Virtual machine initialization times might vary because of
numerous factors. We recommend that you test how long an
instance may take to initialize. To do this, create an instance
and time the startup process.`,
							Default: 60,
						},
						"cpu_utilization": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							Description: `Defines the CPU utilization policy that allows the autoscaler to
scale based on the average CPU utilization of a managed instance
group.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target": {
										Type:     schema.TypeFloat,
										Required: true,
										Description: `The target CPU utilization that the autoscaler should maintain.
Must be a float value in the range (0, 1]. If not specified, the
default is 0.6.

If the CPU level is below the target utilization, the autoscaler
scales down the number of instances until it reaches the minimum
number of instances you specified or until the average CPU of
your instances reaches the target utilization.

If the average CPU is above the target utilization, the autoscaler
scales up until it reaches the maximum number of instances you
specified or until the average utilization reaches the target
utilization.`,
									},
								},
							},
						},
						"load_balancing_utilization": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Configuration parameters of autoscaling based on a load balancer.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target": {
										Type:     schema.TypeFloat,
										Required: true,
										Description: `Fraction of backend capacity utilization (set in HTTP(s) load
balancing configuration) that autoscaler should maintain. Must
be a positive float value. If not defined, the default is 0.8.`,
									},
								},
							},
						},
						"metric": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `Defines the CPU utilization policy that allows the autoscaler to
scale based on the average CPU utilization of a managed instance
group.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
										Description: `The identifier (type) of the Stackdriver Monitoring metric.
The metric cannot have negative values.

The metric must have a value type of INT64 or DOUBLE.`,
									},
									"target": {
										Type:     schema.TypeFloat,
										Optional: true,
										Description: `The target value of the metric that autoscaler should
maintain. This must be a positive value. A utilization
metric scales number of virtual machines handling requests
to increase or decrease proportionally to the metric.

For example, a good metric to use as a utilizationTarget is
www.googleapis.com/compute/instance/network/received_bytes_count.
The autoscaler will work to keep this value constant for each
of the instances.`,
									},
									"type": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"GAUGE", "DELTA_PER_SECOND", "DELTA_PER_MINUTE", ""}, false),
										Description: `Defines how target utilization value is expressed for a
Stackdriver Monitoring metric. Either GAUGE, DELTA_PER_SECOND,
or DELTA_PER_MINUTE.`,
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateGCPName,
				Description: `Name of the resource. The name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"target": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `URL of the managed instance group that this autoscaler will scale.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of this resource.`,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `URL of the region where the instance group resides.`,
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
	}
}

func resourceComputeRegionAutoscalerCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionAutoscalerName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeRegionAutoscalerDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	autoscalingPolicyProp, err := expandComputeRegionAutoscalerAutoscalingPolicy(d.Get("autoscaling_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("autoscaling_policy"); !isEmptyValue(reflect.ValueOf(autoscalingPolicyProp)) && (ok || !reflect.DeepEqual(v, autoscalingPolicyProp)) {
		obj["autoscalingPolicy"] = autoscalingPolicyProp
	}
	targetProp, err := expandComputeRegionAutoscalerTarget(d.Get("target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	regionProp, err := expandComputeRegionAutoscalerRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/autoscalers")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionAutoscaler: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating RegionAutoscaler: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/autoscalers/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	waitErr := computeOperationWaitTime(
		config, res, project, "Creating RegionAutoscaler",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionAutoscaler: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating RegionAutoscaler %q: %#v", d.Id(), res)

	return resourceComputeRegionAutoscalerRead(d, meta)
}

func resourceComputeRegionAutoscalerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/autoscalers/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRegionAutoscaler %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeRegionAutoscalerCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("name", flattenComputeRegionAutoscalerName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("description", flattenComputeRegionAutoscalerDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("autoscaling_policy", flattenComputeRegionAutoscalerAutoscalingPolicy(res["autoscalingPolicy"], d)); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("target", flattenComputeRegionAutoscalerTarget(res["target"], d)); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("region", flattenComputeRegionAutoscalerRegion(res["region"], d)); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}

	return nil
}

func resourceComputeRegionAutoscalerUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionAutoscalerName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeRegionAutoscalerDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	autoscalingPolicyProp, err := expandComputeRegionAutoscalerAutoscalingPolicy(d.Get("autoscaling_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("autoscaling_policy"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, autoscalingPolicyProp)) {
		obj["autoscalingPolicy"] = autoscalingPolicyProp
	}
	targetProp, err := expandComputeRegionAutoscalerTarget(d.Get("target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	regionProp, err := expandComputeRegionAutoscalerRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/autoscalers?autoscaler={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RegionAutoscaler %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating RegionAutoscaler %q: %s", d.Id(), err)
	}

	err = computeOperationWaitTime(
		config, res, project, "Updating RegionAutoscaler",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))
	if err != nil {
		return err
	}

	return resourceComputeRegionAutoscalerRead(d, meta)
}

func resourceComputeRegionAutoscalerDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/autoscalers/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting RegionAutoscaler %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "RegionAutoscaler")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting RegionAutoscaler",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RegionAutoscaler %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionAutoscalerImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/autoscalers/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/autoscalers/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRegionAutoscalerCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionAutoscalerName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionAutoscalerDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicy(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["min_replicas"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyMinReplicas(original["minNumReplicas"], d)
	transformed["max_replicas"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyMaxReplicas(original["maxNumReplicas"], d)
	transformed["cooldown_period"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyCooldownPeriod(original["coolDownPeriodSec"], d)
	transformed["cpu_utilization"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyCpuUtilization(original["cpuUtilization"], d)
	transformed["metric"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyMetric(original["customMetricUtilizations"], d)
	transformed["load_balancing_utilization"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(original["loadBalancingUtilization"], d)
	return []interface{}{transformed}
}
func flattenComputeRegionAutoscalerAutoscalingPolicyMinReplicas(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyMaxReplicas(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyCooldownPeriod(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyCpuUtilization(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["target"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationTarget(original["utilizationTarget"], d)
	return []interface{}{transformed}
}
func flattenComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationTarget(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyMetric(v interface{}, d *schema.ResourceData) interface{} {
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
			"name":   flattenComputeRegionAutoscalerAutoscalingPolicyMetricName(original["metric"], d),
			"target": flattenComputeRegionAutoscalerAutoscalingPolicyMetricTarget(original["utilizationTarget"], d),
			"type":   flattenComputeRegionAutoscalerAutoscalingPolicyMetricType(original["utilizationTargetType"], d),
		})
	}
	return transformed
}
func flattenComputeRegionAutoscalerAutoscalingPolicyMetricName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyMetricTarget(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyMetricType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["target"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(original["utilizationTarget"], d)
	return []interface{}{transformed}
}
func flattenComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionAutoscalerTarget(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionAutoscalerRegion(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeRegionAutoscalerName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinReplicas, err := expandComputeRegionAutoscalerAutoscalingPolicyMinReplicas(original["min_replicas"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["minNumReplicas"] = transformedMinReplicas
	}

	transformedMaxReplicas, err := expandComputeRegionAutoscalerAutoscalingPolicyMaxReplicas(original["max_replicas"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxReplicas); val.IsValid() && !isEmptyValue(val) {
		transformed["maxNumReplicas"] = transformedMaxReplicas
	}

	transformedCooldownPeriod, err := expandComputeRegionAutoscalerAutoscalingPolicyCooldownPeriod(original["cooldown_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCooldownPeriod); val.IsValid() && !isEmptyValue(val) {
		transformed["coolDownPeriodSec"] = transformedCooldownPeriod
	}

	transformedCpuUtilization, err := expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilization(original["cpu_utilization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpuUtilization); val.IsValid() && !isEmptyValue(val) {
		transformed["cpuUtilization"] = transformedCpuUtilization
	}

	transformedMetric, err := expandComputeRegionAutoscalerAutoscalingPolicyMetric(original["metric"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMetric); val.IsValid() && !isEmptyValue(val) {
		transformed["customMetricUtilizations"] = transformedMetric
	}

	transformedLoadBalancingUtilization, err := expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(original["load_balancing_utilization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLoadBalancingUtilization); val.IsValid() && !isEmptyValue(val) {
		transformed["loadBalancingUtilization"] = transformedLoadBalancingUtilization
	}

	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMinReplicas(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMaxReplicas(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyCooldownPeriod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilization(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTarget, err := expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationTarget(original["target"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTarget); val.IsValid() && !isEmptyValue(val) {
		transformed["utilizationTarget"] = transformedTarget
	}

	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetric(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandComputeRegionAutoscalerAutoscalingPolicyMetricName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["metric"] = transformedName
		}

		transformedTarget, err := expandComputeRegionAutoscalerAutoscalingPolicyMetricTarget(original["target"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTarget); val.IsValid() && !isEmptyValue(val) {
			transformed["utilizationTarget"] = transformedTarget
		}

		transformedType, err := expandComputeRegionAutoscalerAutoscalingPolicyMetricType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !isEmptyValue(val) {
			transformed["utilizationTargetType"] = transformedType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetricName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetricTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetricType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTarget, err := expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(original["target"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTarget); val.IsValid() && !isEmptyValue(val) {
		transformed["utilizationTarget"] = transformedTarget
	}

	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
