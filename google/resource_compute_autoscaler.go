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

func resourceComputeAutoscaler() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeAutoscalerCreate,
		Read:   resourceComputeAutoscalerRead,
		Update: resourceComputeAutoscalerUpdate,
		Delete: resourceComputeAutoscalerDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeAutoscalerImport,
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
										Computed: true,
										Optional: true,
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
										Optional: true,
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
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zone": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
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

func resourceComputeAutoscalerCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeAutoscalerName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeAutoscalerDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	autoscalingPolicyProp, err := expandComputeAutoscalerAutoscalingPolicy(d.Get("autoscaling_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("autoscaling_policy"); !isEmptyValue(reflect.ValueOf(autoscalingPolicyProp)) && (ok || !reflect.DeepEqual(v, autoscalingPolicyProp)) {
		obj["autoscalingPolicy"] = autoscalingPolicyProp
	}
	targetProp, err := expandComputeAutoscalerTarget(d.Get("target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	zoneProp, err := expandComputeAutoscalerZone(d.Get("zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zone"); !isEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/zones/{{zone}}/autoscalers")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Autoscaler: %#v", obj)
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating Autoscaler: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{zone}}/{{name}}")
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
		config.clientCompute, op, project, "Creating Autoscaler",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Autoscaler: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating Autoscaler %q: %#v", d.Id(), res)

	return resourceComputeAutoscalerRead(d, meta)
}

func resourceComputeAutoscalerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/zones/{{zone}}/autoscalers/{{name}}")
	if err != nil {
		return err
	}

	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeAutoscaler %q", d.Id()))
	}

	if err := d.Set("creation_timestamp", flattenComputeAutoscalerCreationTimestamp(res["creationTimestamp"])); err != nil {
		return fmt.Errorf("Error reading Autoscaler: %s", err)
	}
	if err := d.Set("name", flattenComputeAutoscalerName(res["name"])); err != nil {
		return fmt.Errorf("Error reading Autoscaler: %s", err)
	}
	if err := d.Set("description", flattenComputeAutoscalerDescription(res["description"])); err != nil {
		return fmt.Errorf("Error reading Autoscaler: %s", err)
	}
	if err := d.Set("autoscaling_policy", flattenComputeAutoscalerAutoscalingPolicy(res["autoscalingPolicy"])); err != nil {
		return fmt.Errorf("Error reading Autoscaler: %s", err)
	}
	if err := d.Set("target", flattenComputeAutoscalerTarget(res["target"])); err != nil {
		return fmt.Errorf("Error reading Autoscaler: %s", err)
	}
	if err := d.Set("zone", flattenComputeAutoscalerZone(res["zone"])); err != nil {
		return fmt.Errorf("Error reading Autoscaler: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading Autoscaler: %s", err)
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Autoscaler: %s", err)
	}

	return nil
}

func resourceComputeAutoscalerUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeAutoscalerName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeAutoscalerDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	autoscalingPolicyProp, err := expandComputeAutoscalerAutoscalingPolicy(d.Get("autoscaling_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("autoscaling_policy"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, autoscalingPolicyProp)) {
		obj["autoscalingPolicy"] = autoscalingPolicyProp
	}
	targetProp, err := expandComputeAutoscalerTarget(d.Get("target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	zoneProp, err := expandComputeAutoscalerZone(d.Get("zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zone"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/zones/{{zone}}/autoscalers?autoscaler={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Autoscaler %q: %#v", d.Id(), obj)
	res, err := sendRequest(config, "PUT", url, obj)

	if err != nil {
		return fmt.Errorf("Error updating Autoscaler %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Updating Autoscaler",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceComputeAutoscalerRead(d, meta)
}

func resourceComputeAutoscalerDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/zones/{{zone}}/autoscalers/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting Autoscaler %q", d.Id())
	res, err := Delete(config, url)
	if err != nil {
		return handleNotFoundError(err, d, "Autoscaler")
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting Autoscaler",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Autoscaler %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeAutoscalerImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/autoscalers/(?P<name>[^/]+)", "(?P<zone>[^/]+)/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{zone}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeAutoscalerCreationTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeAutoscalerName(v interface{}) interface{} {
	return v
}

func flattenComputeAutoscalerDescription(v interface{}) interface{} {
	return v
}

func flattenComputeAutoscalerAutoscalingPolicy(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["min_replicas"] =
		flattenComputeAutoscalerAutoscalingPolicyMinReplicas(original["minNumReplicas"])
	transformed["max_replicas"] =
		flattenComputeAutoscalerAutoscalingPolicyMaxReplicas(original["maxNumReplicas"])
	transformed["cooldown_period"] =
		flattenComputeAutoscalerAutoscalingPolicyCooldownPeriod(original["coolDownPeriodSec"])
	transformed["cpu_utilization"] =
		flattenComputeAutoscalerAutoscalingPolicyCpuUtilization(original["cpuUtilization"])
	transformed["metric"] =
		flattenComputeAutoscalerAutoscalingPolicyMetric(original["customMetricUtilizations"])
	transformed["load_balancing_utilization"] =
		flattenComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization(original["loadBalancingUtilization"])
	return []interface{}{transformed}
}
func flattenComputeAutoscalerAutoscalingPolicyMinReplicas(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeAutoscalerAutoscalingPolicyMaxReplicas(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeAutoscalerAutoscalingPolicyCooldownPeriod(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeAutoscalerAutoscalingPolicyCpuUtilization(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["target"] =
		flattenComputeAutoscalerAutoscalingPolicyCpuUtilizationTarget(original["utilizationTarget"])
	return []interface{}{transformed}
}
func flattenComputeAutoscalerAutoscalingPolicyCpuUtilizationTarget(v interface{}) interface{} {
	return v
}

func flattenComputeAutoscalerAutoscalingPolicyMetric(v interface{}) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"name":   flattenComputeAutoscalerAutoscalingPolicyMetricName(original["metric"]),
			"target": flattenComputeAutoscalerAutoscalingPolicyMetricTarget(original["utilizationTarget"]),
			"type":   flattenComputeAutoscalerAutoscalingPolicyMetricType(original["utilizationTargetType"]),
		})
	}
	return transformed
}
func flattenComputeAutoscalerAutoscalingPolicyMetricName(v interface{}) interface{} {
	return v
}

func flattenComputeAutoscalerAutoscalingPolicyMetricTarget(v interface{}) interface{} {
	return v
}

func flattenComputeAutoscalerAutoscalingPolicyMetricType(v interface{}) interface{} {
	return v
}

func flattenComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["target"] =
		flattenComputeAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(original["utilizationTarget"])
	return []interface{}{transformed}
}
func flattenComputeAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(v interface{}) interface{} {
	return v
}

func flattenComputeAutoscalerTarget(v interface{}) interface{} {
	return v
}

func flattenComputeAutoscalerZone(v interface{}) interface{} {
	return v
}

func expandComputeAutoscalerName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicy(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMinReplicas, err := expandComputeAutoscalerAutoscalingPolicyMinReplicas(original["min_replicas"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["minNumReplicas"] = transformedMinReplicas
		transformedMaxReplicas, err := expandComputeAutoscalerAutoscalingPolicyMaxReplicas(original["max_replicas"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["maxNumReplicas"] = transformedMaxReplicas
		transformedCooldownPeriod, err := expandComputeAutoscalerAutoscalingPolicyCooldownPeriod(original["cooldown_period"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["coolDownPeriodSec"] = transformedCooldownPeriod
		transformedCpuUtilization, err := expandComputeAutoscalerAutoscalingPolicyCpuUtilization(original["cpu_utilization"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["cpuUtilization"] = transformedCpuUtilization
		transformedMetric, err := expandComputeAutoscalerAutoscalingPolicyMetric(original["metric"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["customMetricUtilizations"] = transformedMetric
		transformedLoadBalancingUtilization, err := expandComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization(original["load_balancing_utilization"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["loadBalancingUtilization"] = transformedLoadBalancingUtilization

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeAutoscalerAutoscalingPolicyMinReplicas(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyMaxReplicas(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyCooldownPeriod(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyCpuUtilization(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedTarget, err := expandComputeAutoscalerAutoscalingPolicyCpuUtilizationTarget(original["target"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["utilizationTarget"] = transformedTarget

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeAutoscalerAutoscalingPolicyCpuUtilizationTarget(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyMetric(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandComputeAutoscalerAutoscalingPolicyMetricName(original["name"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["metric"] = transformedName
		transformedTarget, err := expandComputeAutoscalerAutoscalingPolicyMetricTarget(original["target"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["utilizationTarget"] = transformedTarget
		transformedType, err := expandComputeAutoscalerAutoscalingPolicyMetricType(original["type"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["utilizationTargetType"] = transformedType

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeAutoscalerAutoscalingPolicyMetricName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyMetricTarget(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyMetricType(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedTarget, err := expandComputeAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(original["target"], d, config)
		if err != nil {
			return nil, err
		}
		transformed["utilizationTarget"] = transformedTarget

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeAutoscalerAutoscalingPolicyLoadBalancingUtilizationTarget(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeAutoscalerTarget(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseZonalFieldValue("instanceGroupManagers", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for target: %s", err)
	}
	return "https://www.googleapis.com/compute/v1/" + f.RelativeLink(), nil
}

func expandComputeAutoscalerZone(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}
