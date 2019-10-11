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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"google.golang.org/api/compute/v1"
)

func resourceComputeGlobalForwardingRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeGlobalForwardingRuleCreate,
		Read:   resourceComputeGlobalForwardingRuleRead,
		Update: resourceComputeGlobalForwardingRuleUpdate,
		Delete: resourceComputeGlobalForwardingRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeGlobalForwardingRuleImport,
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
			},
			"target": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkRelativePaths,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"ip_protocol": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				ValidateFunc:     validation.StringInSlice([]string{"TCP", "UDP", "ESP", "AH", "SCTP", "ICMP", ""}, false),
				DiffSuppressFunc: caseDiffSuppress,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ip_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"IPV4", "IPV6", ""}, false),
			},
			"load_balancing_scheme": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"INTERNAL_SELF_MANAGED", "EXTERNAL", ""}, false),
				Default:      "EXTERNAL",
			},
			"metadata_filters": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_labels": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MinItems: 1,
							MaxItems: 64,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
								},
							},
						},
						"filter_match_criteria": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"MATCH_ANY", "MATCH_ALL"}, false),
						},
					},
				},
			},
			"port_range": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: portRangeDiffSuppress,
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

func resourceComputeGlobalForwardingRuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeGlobalForwardingRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	IPAddressProp, err := expandComputeGlobalForwardingRuleIPAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_address"); !isEmptyValue(reflect.ValueOf(IPAddressProp)) && (ok || !reflect.DeepEqual(v, IPAddressProp)) {
		obj["IPAddress"] = IPAddressProp
	}
	IPProtocolProp, err := expandComputeGlobalForwardingRuleIPProtocol(d.Get("ip_protocol"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_protocol"); !isEmptyValue(reflect.ValueOf(IPProtocolProp)) && (ok || !reflect.DeepEqual(v, IPProtocolProp)) {
		obj["IPProtocol"] = IPProtocolProp
	}
	ipVersionProp, err := expandComputeGlobalForwardingRuleIpVersion(d.Get("ip_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_version"); !isEmptyValue(reflect.ValueOf(ipVersionProp)) && (ok || !reflect.DeepEqual(v, ipVersionProp)) {
		obj["ipVersion"] = ipVersionProp
	}
	loadBalancingSchemeProp, err := expandComputeGlobalForwardingRuleLoadBalancingScheme(d.Get("load_balancing_scheme"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("load_balancing_scheme"); !isEmptyValue(reflect.ValueOf(loadBalancingSchemeProp)) && (ok || !reflect.DeepEqual(v, loadBalancingSchemeProp)) {
		obj["loadBalancingScheme"] = loadBalancingSchemeProp
	}
	metadataFiltersProp, err := expandComputeGlobalForwardingRuleMetadataFilters(d.Get("metadata_filters"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata_filters"); !isEmptyValue(reflect.ValueOf(metadataFiltersProp)) && (ok || !reflect.DeepEqual(v, metadataFiltersProp)) {
		obj["metadataFilters"] = metadataFiltersProp
	}
	nameProp, err := expandComputeGlobalForwardingRuleName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	portRangeProp, err := expandComputeGlobalForwardingRulePortRange(d.Get("port_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port_range"); !isEmptyValue(reflect.ValueOf(portRangeProp)) && (ok || !reflect.DeepEqual(v, portRangeProp)) {
		obj["portRange"] = portRangeProp
	}
	targetProp, err := expandComputeGlobalForwardingRuleTarget(d.Get("target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/forwardingRules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GlobalForwardingRule: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating GlobalForwardingRule: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/global/forwardingRules/{{name}}")
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
		config.clientCompute, op, project, "Creating GlobalForwardingRule",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create GlobalForwardingRule: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating GlobalForwardingRule %q: %#v", d.Id(), res)

	return resourceComputeGlobalForwardingRuleRead(d, meta)
}

func resourceComputeGlobalForwardingRuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/forwardingRules/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeGlobalForwardingRule %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}

	if err := d.Set("description", flattenComputeGlobalForwardingRuleDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("ip_address", flattenComputeGlobalForwardingRuleIPAddress(res["IPAddress"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("ip_protocol", flattenComputeGlobalForwardingRuleIPProtocol(res["IPProtocol"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("ip_version", flattenComputeGlobalForwardingRuleIpVersion(res["ipVersion"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("load_balancing_scheme", flattenComputeGlobalForwardingRuleLoadBalancingScheme(res["loadBalancingScheme"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("metadata_filters", flattenComputeGlobalForwardingRuleMetadataFilters(res["metadataFilters"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("name", flattenComputeGlobalForwardingRuleName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("port_range", flattenComputeGlobalForwardingRulePortRange(res["portRange"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("target", flattenComputeGlobalForwardingRuleTarget(res["target"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}

	return nil
}

func resourceComputeGlobalForwardingRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	d.Partial(true)

	if d.HasChange("target") {
		obj := make(map[string]interface{})
		targetProp, err := expandComputeGlobalForwardingRuleTarget(d.Get("target"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetProp)) {
			obj["target"] = targetProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/forwardingRules/{{name}}/setTarget")
		if err != nil {
			return err
		}
		res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating GlobalForwardingRule %q: %s", d.Id(), err)
		}

		op := &compute.Operation{}
		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating GlobalForwardingRule",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("target")
	}

	d.Partial(false)

	return resourceComputeGlobalForwardingRuleRead(d, meta)
}

func resourceComputeGlobalForwardingRuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/forwardingRules/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting GlobalForwardingRule %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "GlobalForwardingRule")
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting GlobalForwardingRule",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting GlobalForwardingRule %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeGlobalForwardingRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/forwardingRules/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/global/forwardingRules/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeGlobalForwardingRuleDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleIPAddress(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleIPProtocol(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleIpVersion(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleLoadBalancingScheme(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleMetadataFilters(v interface{}, d *schema.ResourceData) interface{} {
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
			"filter_match_criteria": flattenComputeGlobalForwardingRuleMetadataFiltersFilterMatchCriteria(original["filterMatchCriteria"], d),
			"filter_labels":         flattenComputeGlobalForwardingRuleMetadataFiltersFilterLabels(original["filterLabels"], d),
		})
	}
	return transformed
}
func flattenComputeGlobalForwardingRuleMetadataFiltersFilterMatchCriteria(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleMetadataFiltersFilterLabels(v interface{}, d *schema.ResourceData) interface{} {
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
			"name":  flattenComputeGlobalForwardingRuleMetadataFiltersFilterLabelsName(original["name"], d),
			"value": flattenComputeGlobalForwardingRuleMetadataFiltersFilterLabelsValue(original["value"], d),
		})
	}
	return transformed
}
func flattenComputeGlobalForwardingRuleMetadataFiltersFilterLabelsName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleMetadataFiltersFilterLabelsValue(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalForwardingRulePortRange(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleTarget(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandComputeGlobalForwardingRuleDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleIPAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleIPProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleIpVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleLoadBalancingScheme(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleMetadataFilters(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedFilterMatchCriteria, err := expandComputeGlobalForwardingRuleMetadataFiltersFilterMatchCriteria(original["filter_match_criteria"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFilterMatchCriteria); val.IsValid() && !isEmptyValue(val) {
			transformed["filterMatchCriteria"] = transformedFilterMatchCriteria
		}

		transformedFilterLabels, err := expandComputeGlobalForwardingRuleMetadataFiltersFilterLabels(original["filter_labels"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFilterLabels); val.IsValid() && !isEmptyValue(val) {
			transformed["filterLabels"] = transformedFilterLabels
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeGlobalForwardingRuleMetadataFiltersFilterMatchCriteria(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleMetadataFiltersFilterLabels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandComputeGlobalForwardingRuleMetadataFiltersFilterLabelsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValue, err := expandComputeGlobalForwardingRuleMetadataFiltersFilterLabelsValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !isEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeGlobalForwardingRuleMetadataFiltersFilterLabelsName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleMetadataFiltersFilterLabelsValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRulePortRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
