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
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceMonitoringUptimeCheckConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitoringUptimeCheckConfigCreate,
		Read:   resourceMonitoringUptimeCheckConfigRead,
		Update: resourceMonitoringUptimeCheckConfigUpdate,
		Delete: resourceMonitoringUptimeCheckConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMonitoringUptimeCheckConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"timeout": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content_matchers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"http_check": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_info": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"password": {
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
									"username": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"headers": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"mask_headers": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"path": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "/",
						},
						"port": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
						},
						"use_ssl": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
				ConflictsWith: []string{"tcp_check"},
			},
			"monitored_resource": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"labels": {
							Type:     schema.TypeMap,
							Required: true,
							ForceNew: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
				ConflictsWith: []string{"resource_group"},
			},
			"period": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "300s",
			},
			"resource_group": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:             schema.TypeString,
							Optional:         true,
							ForceNew:         true,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
						},
						"resource_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"RESOURCE_TYPE_UNSPECIFIED", "INSTANCE", "AWS_ELB_LOAD_BALANCER", ""}, false),
						},
					},
				},
				ConflictsWith: []string{"monitored_resource"},
			},
			"selected_regions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tcp_check": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
				ConflictsWith: []string{"http_check"},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uptime_check_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_internal": {
				Type:       schema.TypeBool,
				Optional:   true,
				Computed:   true,
				Deprecated: "This field never worked, and will be removed in 3.0.0.",
			},
			"internal_checkers": {
				Type:       schema.TypeList,
				Optional:   true,
				Computed:   true,
				Deprecated: "This field never worked, and will be removed in 3.0.0.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_name": {
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "This field never worked, and will be removed in 3.0.0.",
						},
						"gcp_zone": {
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "This field never worked, and will be removed in 3.0.0.",
						},
						"name": {
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "This field never worked, and will be removed in 3.0.0.",
						},
						"network": {
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "This field never worked, and will be removed in 3.0.0.",
						},
						"peer_project_id": {
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "This field never worked, and will be removed in 3.0.0.",
						},
					},
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceMonitoringUptimeCheckConfigCreate(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)

	obj := make(map[string]interface{})
	displayNameProp, err := expandMonitoringUptimeCheckConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	periodProp, err := expandMonitoringUptimeCheckConfigPeriod(d.Get("period"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("period"); !isEmptyValue(reflect.ValueOf(periodProp)) && (ok || !reflect.DeepEqual(v, periodProp)) {
		obj["period"] = periodProp
	}
	timeoutProp, err := expandMonitoringUptimeCheckConfigTimeout(d.Get("timeout"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout"); !isEmptyValue(reflect.ValueOf(timeoutProp)) && (ok || !reflect.DeepEqual(v, timeoutProp)) {
		obj["timeout"] = timeoutProp
	}
	contentMatchersProp, err := expandMonitoringUptimeCheckConfigContentMatchers(d.Get("content_matchers"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("content_matchers"); !isEmptyValue(reflect.ValueOf(contentMatchersProp)) && (ok || !reflect.DeepEqual(v, contentMatchersProp)) {
		obj["contentMatchers"] = contentMatchersProp
	}
	selectedRegionsProp, err := expandMonitoringUptimeCheckConfigSelectedRegions(d.Get("selected_regions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("selected_regions"); !isEmptyValue(reflect.ValueOf(selectedRegionsProp)) && (ok || !reflect.DeepEqual(v, selectedRegionsProp)) {
		obj["selectedRegions"] = selectedRegionsProp
	}
	httpCheckProp, err := expandMonitoringUptimeCheckConfigHttpCheck(d.Get("http_check"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("http_check"); !isEmptyValue(reflect.ValueOf(httpCheckProp)) && (ok || !reflect.DeepEqual(v, httpCheckProp)) {
		obj["httpCheck"] = httpCheckProp
	}
	tcpCheckProp, err := expandMonitoringUptimeCheckConfigTcpCheck(d.Get("tcp_check"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tcp_check"); !isEmptyValue(reflect.ValueOf(tcpCheckProp)) && (ok || !reflect.DeepEqual(v, tcpCheckProp)) {
		obj["tcpCheck"] = tcpCheckProp
	}
	resourceGroupProp, err := expandMonitoringUptimeCheckConfigResourceGroup(d.Get("resource_group"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("resource_group"); !isEmptyValue(reflect.ValueOf(resourceGroupProp)) && (ok || !reflect.DeepEqual(v, resourceGroupProp)) {
		obj["resourceGroup"] = resourceGroupProp
	}
	monitoredResourceProp, err := expandMonitoringUptimeCheckConfigMonitoredResource(d.Get("monitored_resource"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("monitored_resource"); !isEmptyValue(reflect.ValueOf(monitoredResourceProp)) && (ok || !reflect.DeepEqual(v, monitoredResourceProp)) {
		obj["monitoredResource"] = monitoredResourceProp
	}

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}projects/{{project}}/uptimeCheckConfigs")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new UptimeCheckConfig: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating UptimeCheckConfig: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating UptimeCheckConfig %q: %#v", d.Id(), res)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
	}
	d.Set("name", name.(string))
	d.SetId(name.(string))

	return resourceMonitoringUptimeCheckConfigRead(d, meta)
}

func resourceMonitoringUptimeCheckConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("MonitoringUptimeCheckConfig %q", d.Id()))
	}

	res, err = resourceMonitoringUptimeCheckConfigDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing MonitoringUptimeCheckConfig because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading UptimeCheckConfig: %s", err)
	}

	if err := d.Set("name", flattenMonitoringUptimeCheckConfigName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading UptimeCheckConfig: %s", err)
	}
	if err := d.Set("uptime_check_id", flattenMonitoringUptimeCheckConfigUptimeCheckId(res["id"], d)); err != nil {
		return fmt.Errorf("Error reading UptimeCheckConfig: %s", err)
	}
	if err := d.Set("display_name", flattenMonitoringUptimeCheckConfigDisplayName(res["displayName"], d)); err != nil {
		return fmt.Errorf("Error reading UptimeCheckConfig: %s", err)
	}
	if err := d.Set("period", flattenMonitoringUptimeCheckConfigPeriod(res["period"], d)); err != nil {
		return fmt.Errorf("Error reading UptimeCheckConfig: %s", err)
	}
	if err := d.Set("timeout", flattenMonitoringUptimeCheckConfigTimeout(res["timeout"], d)); err != nil {
		return fmt.Errorf("Error reading UptimeCheckConfig: %s", err)
	}
	if err := d.Set("content_matchers", flattenMonitoringUptimeCheckConfigContentMatchers(res["contentMatchers"], d)); err != nil {
		return fmt.Errorf("Error reading UptimeCheckConfig: %s", err)
	}
	if err := d.Set("selected_regions", flattenMonitoringUptimeCheckConfigSelectedRegions(res["selectedRegions"], d)); err != nil {
		return fmt.Errorf("Error reading UptimeCheckConfig: %s", err)
	}
	if err := d.Set("http_check", flattenMonitoringUptimeCheckConfigHttpCheck(res["httpCheck"], d)); err != nil {
		return fmt.Errorf("Error reading UptimeCheckConfig: %s", err)
	}
	if err := d.Set("tcp_check", flattenMonitoringUptimeCheckConfigTcpCheck(res["tcpCheck"], d)); err != nil {
		return fmt.Errorf("Error reading UptimeCheckConfig: %s", err)
	}
	if err := d.Set("resource_group", flattenMonitoringUptimeCheckConfigResourceGroup(res["resourceGroup"], d)); err != nil {
		return fmt.Errorf("Error reading UptimeCheckConfig: %s", err)
	}
	if err := d.Set("monitored_resource", flattenMonitoringUptimeCheckConfigMonitoredResource(res["monitoredResource"], d)); err != nil {
		return fmt.Errorf("Error reading UptimeCheckConfig: %s", err)
	}

	return nil
}

func resourceMonitoringUptimeCheckConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandMonitoringUptimeCheckConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	timeoutProp, err := expandMonitoringUptimeCheckConfigTimeout(d.Get("timeout"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, timeoutProp)) {
		obj["timeout"] = timeoutProp
	}
	contentMatchersProp, err := expandMonitoringUptimeCheckConfigContentMatchers(d.Get("content_matchers"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("content_matchers"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, contentMatchersProp)) {
		obj["contentMatchers"] = contentMatchersProp
	}
	selectedRegionsProp, err := expandMonitoringUptimeCheckConfigSelectedRegions(d.Get("selected_regions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("selected_regions"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, selectedRegionsProp)) {
		obj["selectedRegions"] = selectedRegionsProp
	}
	httpCheckProp, err := expandMonitoringUptimeCheckConfigHttpCheck(d.Get("http_check"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("http_check"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, httpCheckProp)) {
		obj["httpCheck"] = httpCheckProp
	}
	tcpCheckProp, err := expandMonitoringUptimeCheckConfigTcpCheck(d.Get("tcp_check"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("tcp_check"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, tcpCheckProp)) {
		obj["tcpCheck"] = tcpCheckProp
	}

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating UptimeCheckConfig %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("timeout") {
		updateMask = append(updateMask, "timeout")
	}

	if d.HasChange("content_matchers") {
		updateMask = append(updateMask, "contentMatchers")
	}

	if d.HasChange("selected_regions") {
		updateMask = append(updateMask, "selectedRegions")
	}

	if d.HasChange("http_check") {
		updateMask = append(updateMask, "httpCheck")
	}

	if d.HasChange("tcp_check") {
		updateMask = append(updateMask, "tcpCheck")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating UptimeCheckConfig %q: %s", d.Id(), err)
	}

	return resourceMonitoringUptimeCheckConfigRead(d, meta)
}

func resourceMonitoringUptimeCheckConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{MonitoringBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting UptimeCheckConfig %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "UptimeCheckConfig")
	}

	log.Printf("[DEBUG] Finished deleting UptimeCheckConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceMonitoringUptimeCheckConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func flattenMonitoringUptimeCheckConfigName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringUptimeCheckConfigUptimeCheckId(v interface{}, d *schema.ResourceData) interface{} {
	parts := strings.Split(d.Get("name").(string), "/")
	return parts[len(parts)-1]
}

func flattenMonitoringUptimeCheckConfigDisplayName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringUptimeCheckConfigPeriod(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringUptimeCheckConfigTimeout(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringUptimeCheckConfigContentMatchers(v interface{}, d *schema.ResourceData) interface{} {
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
			"content": flattenMonitoringUptimeCheckConfigContentMatchersContent(original["content"], d),
		})
	}
	return transformed
}
func flattenMonitoringUptimeCheckConfigContentMatchersContent(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringUptimeCheckConfigSelectedRegions(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringUptimeCheckConfigHttpCheck(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["auth_info"] =
		flattenMonitoringUptimeCheckConfigHttpCheckAuthInfo(original["authInfo"], d)
	transformed["port"] =
		flattenMonitoringUptimeCheckConfigHttpCheckPort(original["port"], d)
	transformed["headers"] =
		flattenMonitoringUptimeCheckConfigHttpCheckHeaders(original["headers"], d)
	transformed["path"] =
		flattenMonitoringUptimeCheckConfigHttpCheckPath(original["path"], d)
	transformed["use_ssl"] =
		flattenMonitoringUptimeCheckConfigHttpCheckUseSsl(original["useSsl"], d)
	transformed["mask_headers"] =
		flattenMonitoringUptimeCheckConfigHttpCheckMaskHeaders(original["maskHeaders"], d)
	return []interface{}{transformed}
}
func flattenMonitoringUptimeCheckConfigHttpCheckAuthInfo(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["password"] =
		flattenMonitoringUptimeCheckConfigHttpCheckAuthInfoPassword(original["password"], d)
	transformed["username"] =
		flattenMonitoringUptimeCheckConfigHttpCheckAuthInfoUsername(original["username"], d)
	return []interface{}{transformed}
}
func flattenMonitoringUptimeCheckConfigHttpCheckAuthInfoPassword(v interface{}, d *schema.ResourceData) interface{} {
	return d.Get("http_check.0.auth_info.0.password")
}

func flattenMonitoringUptimeCheckConfigHttpCheckAuthInfoUsername(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringUptimeCheckConfigHttpCheckPort(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenMonitoringUptimeCheckConfigHttpCheckHeaders(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringUptimeCheckConfigHttpCheckPath(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringUptimeCheckConfigHttpCheckUseSsl(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringUptimeCheckConfigHttpCheckMaskHeaders(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringUptimeCheckConfigTcpCheck(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["port"] =
		flattenMonitoringUptimeCheckConfigTcpCheckPort(original["port"], d)
	return []interface{}{transformed}
}
func flattenMonitoringUptimeCheckConfigTcpCheckPort(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenMonitoringUptimeCheckConfigResourceGroup(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["resource_type"] =
		flattenMonitoringUptimeCheckConfigResourceGroupResourceType(original["resourceType"], d)
	transformed["group_id"] =
		flattenMonitoringUptimeCheckConfigResourceGroupGroupId(original["groupId"], d)
	return []interface{}{transformed}
}
func flattenMonitoringUptimeCheckConfigResourceGroupResourceType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringUptimeCheckConfigResourceGroupGroupId(v interface{}, d *schema.ResourceData) interface{} {
	project := d.Get("project").(string)
	return fmt.Sprintf("projects/%s/groups/%s", project, v)
}

func flattenMonitoringUptimeCheckConfigMonitoredResource(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["type"] =
		flattenMonitoringUptimeCheckConfigMonitoredResourceType(original["type"], d)
	transformed["labels"] =
		flattenMonitoringUptimeCheckConfigMonitoredResourceLabels(original["labels"], d)
	return []interface{}{transformed}
}
func flattenMonitoringUptimeCheckConfigMonitoredResourceType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringUptimeCheckConfigMonitoredResourceLabels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandMonitoringUptimeCheckConfigDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigPeriod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigTimeout(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigContentMatchers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedContent, err := expandMonitoringUptimeCheckConfigContentMatchersContent(original["content"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedContent); val.IsValid() && !isEmptyValue(val) {
			transformed["content"] = transformedContent
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMonitoringUptimeCheckConfigContentMatchersContent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigSelectedRegions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheck(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAuthInfo, err := expandMonitoringUptimeCheckConfigHttpCheckAuthInfo(original["auth_info"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAuthInfo); val.IsValid() && !isEmptyValue(val) {
		transformed["authInfo"] = transformedAuthInfo
	}

	transformedPort, err := expandMonitoringUptimeCheckConfigHttpCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedHeaders, err := expandMonitoringUptimeCheckConfigHttpCheckHeaders(original["headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHeaders); val.IsValid() && !isEmptyValue(val) {
		transformed["headers"] = transformedHeaders
	}

	transformedPath, err := expandMonitoringUptimeCheckConfigHttpCheckPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	transformedUseSsl, err := expandMonitoringUptimeCheckConfigHttpCheckUseSsl(original["use_ssl"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUseSsl); val.IsValid() && !isEmptyValue(val) {
		transformed["useSsl"] = transformedUseSsl
	}

	transformedMaskHeaders, err := expandMonitoringUptimeCheckConfigHttpCheckMaskHeaders(original["mask_headers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaskHeaders); val.IsValid() && !isEmptyValue(val) {
		transformed["maskHeaders"] = transformedMaskHeaders
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckAuthInfo(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPassword, err := expandMonitoringUptimeCheckConfigHttpCheckAuthInfoPassword(original["password"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPassword); val.IsValid() && !isEmptyValue(val) {
		transformed["password"] = transformedPassword
	}

	transformedUsername, err := expandMonitoringUptimeCheckConfigHttpCheckAuthInfoUsername(original["username"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !isEmptyValue(val) {
		transformed["username"] = transformedUsername
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckAuthInfoPassword(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckAuthInfoUsername(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckHeaders(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckUseSsl(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigHttpCheckMaskHeaders(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigTcpCheck(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPort, err := expandMonitoringUptimeCheckConfigTcpCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigTcpCheckPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigResourceGroup(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResourceType, err := expandMonitoringUptimeCheckConfigResourceGroupResourceType(original["resource_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceType); val.IsValid() && !isEmptyValue(val) {
		transformed["resourceType"] = transformedResourceType
	}

	transformedGroupId, err := expandMonitoringUptimeCheckConfigResourceGroupGroupId(original["group_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGroupId); val.IsValid() && !isEmptyValue(val) {
		transformed["groupId"] = transformedGroupId
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigResourceGroupResourceType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigResourceGroupGroupId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return GetResourceNameFromSelfLink(v.(string)), nil
}

func expandMonitoringUptimeCheckConfigMonitoredResource(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedType, err := expandMonitoringUptimeCheckConfigMonitoredResourceType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !isEmptyValue(val) {
		transformed["type"] = transformedType
	}

	transformedLabels, err := expandMonitoringUptimeCheckConfigMonitoredResourceLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !isEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	return transformed, nil
}

func expandMonitoringUptimeCheckConfigMonitoredResourceType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringUptimeCheckConfigMonitoredResourceLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func resourceMonitoringUptimeCheckConfigDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	d.Set("internal_checkers", nil)
	return res, nil
}
