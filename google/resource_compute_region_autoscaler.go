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

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	compute "google.golang.org/api/compute/v1"
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
			Create: schema.DefaultTimeout(240 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"autoscaling_policy": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_replicas": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"min_replicas": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"cooldown_period": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  60,
						},
						"cpu_utilization": {
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target": {
										Type:     schema.TypeFloat,
										Required: true,
									},
								},
							},
						},
						"metric": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"target": {
										Type:     schema.TypeFloat,
										Required: true,
									},
									"type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"GAUGE", "DELTA_PER_SECOND", "DELTA_PER_MINUTE"}, false),
									},
								},
							},
						},
						"load_balancing_utilization": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target": {
										Type:     schema.TypeFloat,
										Required: true,
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
			},
			"target": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"creation_timestamp": {
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

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/autoscalers")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionAutoscaler: %#v", obj)
	res, err := sendRequest(config, "POST", url, obj)
	if err != nil {
		return fmt.Errorf("Error creating RegionAutoscaler: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{region}}/{{name}}")
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
		config.clientCompute, op, project, "Creating RegionAutoscaler",
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

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/autoscalers/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRegionAutoscaler %q", d.Id()))
	}

	if err := d.Set("creation_timestamp", flattenComputeRegionAutoscalerCreationTimestamp(res["creationTimestamp"])); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("name", flattenComputeRegionAutoscalerName(res["name"])); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("description", flattenComputeRegionAutoscalerDescription(res["description"])); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("autoscaling_policy", flattenComputeRegionAutoscalerAutoscalingPolicy(res["autoscalingPolicy"])); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("target", flattenComputeRegionAutoscalerTarget(res["target"])); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("region", flattenComputeRegionAutoscalerRegion(res["region"])); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionAutoscaler: %s", err)
	}

	return nil
}

func resourceComputeRegionAutoscalerUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

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

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/autoscalers?autoscaler={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RegionAutoscaler %q: %#v", d.Id(), obj)
	res, err := sendRequest(config, "PUT", url, obj)

	if err != nil {
		return fmt.Errorf("Error updating RegionAutoscaler %q: %s", d.Id(), err)
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
		config.clientCompute, op, project, "Updating RegionAutoscaler",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceComputeRegionAutoscalerRead(d, meta)
}

func resourceComputeRegionAutoscalerDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/autoscalers/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting RegionAutoscaler %q", d.Id())
	res, err := sendRequest(config, "DELETE", url, obj)
	if err != nil {
		return handleNotFoundError(err, d, "RegionAutoscaler")
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
		config.clientCompute, op, project, "Deleting RegionAutoscaler",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RegionAutoscaler %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionAutoscalerImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/autoscalers/(?P<name>[^/]+)", "(?P<region>[^/]+)/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{region}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRegionAutoscalerCreationTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeRegionAutoscalerName(v interface{}) interface{} {
	return v
}

func flattenComputeRegionAutoscalerDescription(v interface{}) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicy(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["min_replicas"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyMinReplicas(original["minNumReplicas"])
	transformed["max_replicas"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyMaxReplicas(original["maxNumReplicas"])
	transformed["cooldown_period"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyCooldownPeriod(original["coolDownPeriodSec"])
	transformed["cpu_utilization"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyCpuUtilization(original["cpuUtilization"])
	transformed["metric"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyMetric(original["customMetricUtilizations"])
	transformed["load_balancing_utilization"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(original["loadBalancingUtilization"])
	return []interface{}{transformed}
}
func flattenComputeRegionAutoscalerAutoscalingPolicyMinReplicas(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyMaxReplicas(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyCooldownPeriod(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyCpuUtilization(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["target"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationTarget(original["utilizationTarget"])
	return []interface{}{transformed}
}
func flattenComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationTarget(v interface{}) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyMetric(v interface{}) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"name":   flattenComputeRegionAutoscalerAutoscalingPolicyMetricName(original["metric"]),
			"target": flattenComputeRegionAutoscalerAutoscalingPolicyMetricTarget(original["utilizationTarget"]),
			"type":   flattenComputeRegionAutoscalerAutoscalingPolicyMetricType(original["utilizationTargetType"]),
		})
	}
	return transformed
}
func flattenComputeRegionAutoscalerAutoscalingPolicyMetricName(v interface{}) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyMetricTarget(v interface{}) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyMetricType(v interface{}) interface{} {
	return v
}

func flattenComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["target"] =
		flattenComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(original["utilizationTarget"])
	return []interface{}{transformed}
}
func flattenComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(v interface{}) interface{} {
	return v
}

func flattenComputeRegionAutoscalerTarget(v interface{}) interface{} {
	return v
}

func flattenComputeRegionAutoscalerRegion(v interface{}) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeRegionAutoscalerName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicy(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinReplicas, err := expandComputeRegionAutoscalerAutoscalingPolicyMinReplicas(original["min_replicas"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["minNumReplicas"] = transformedMinReplicas
	transformedMaxReplicas, err := expandComputeRegionAutoscalerAutoscalingPolicyMaxReplicas(original["max_replicas"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["maxNumReplicas"] = transformedMaxReplicas
	transformedCooldownPeriod, err := expandComputeRegionAutoscalerAutoscalingPolicyCooldownPeriod(original["cooldown_period"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["coolDownPeriodSec"] = transformedCooldownPeriod
	transformedCpuUtilization, err := expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilization(original["cpu_utilization"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["cpuUtilization"] = transformedCpuUtilization
	transformedMetric, err := expandComputeRegionAutoscalerAutoscalingPolicyMetric(original["metric"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["customMetricUtilizations"] = transformedMetric
	transformedLoadBalancingUtilization, err := expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(original["load_balancing_utilization"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["loadBalancingUtilization"] = transformedLoadBalancingUtilization
	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMinReplicas(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMaxReplicas(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyCooldownPeriod(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilization(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTarget, err := expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationTarget(original["target"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["utilizationTarget"] = transformedTarget
	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationTarget(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetric(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandComputeRegionAutoscalerAutoscalingPolicyMetricName(original["name"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["metric"] = transformedName
		transformedTarget, err := expandComputeRegionAutoscalerAutoscalingPolicyMetricTarget(original["target"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["utilizationTarget"] = transformedTarget
		transformedType, err := expandComputeRegionAutoscalerAutoscalingPolicyMetricType(original["type"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["utilizationTargetType"] = transformedType
		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetricName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetricTarget(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyMetricType(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTarget, err := expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(original["target"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["utilizationTarget"] = transformedTarget
	return transformed, nil
}

func expandComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerTarget(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionAutoscalerRegion(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
