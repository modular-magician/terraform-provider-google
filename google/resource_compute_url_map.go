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
)

func resourceComputeUrlMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeUrlMapCreate,
		Read:   resourceComputeUrlMapRead,
		Update: resourceComputeUrlMapUpdate,
		Delete: resourceComputeUrlMapDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeUrlMapImport,
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
				Description: `Name of the resource. Provided by the client when the resource is created. The
name must be 1-63 characters long, and comply with RFC1035. Specifically, the
name must be 1-63 characters long and match the regular expression
'[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase
letter, and all following characters must be a dash, lowercase letter, or digit,
except the last character, which cannot be a dash.`,
			},
			"default_service": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The backend service or backend bucket to use when none of the given rules match.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `An optional description of this resource. Provide this property when you create
the resource.`,
			},
			"host_rule": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: `The list of HostRules to use against the URL.`,
				Elem:        computeUrlMapHostRuleSchema(),
				// Default schema.HashSchema is used.
			},
			"path_matcher": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The list of named PathMatchers to use against the URL.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The name to which this PathMatcher is referred by the HostRule.`,
						},
						"default_service": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
							Description:      `The backend service or backend bucket to use when none of the given paths match.`,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `An optional description of this resource. Provide this property when you create
the resource.`,
						},
						"path_rule": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `The list of path rules.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"paths": {
										Type:     schema.TypeSet,
										Required: true,
										Description: `The list of path patterns to match. Each must start with /
and the only place a * is allowed is at the end following
a /. The string fed to the path matcher does not include
any text after the first ? or #, and those chars are not
allowed here.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set: schema.HashString,
									},
									"service": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: compareSelfLinkOrResourceName,
										Description:      `The backend service or backend bucket to use if any of the given paths match.`,
									},
								},
							},
						},
					},
				},
			},
			"test": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `The list of expected URL mapping tests. Request to update this UrlMap will
succeed only if all of the test cases pass. You can specify a maximum of 100
tests per UrlMap.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Host portion of the URL.`,
						},
						"path": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Path portion of the URL.`,
						},
						"service": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
							Description:      `The backend service or backend bucket link that should be matched by this test.`,
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Description of this test case.`,
						},
					},
				},
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Fingerprint of this resource. A hash of the contents stored in this object. This
field is used in optimistic locking.`,
			},
			"map_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The unique identifier for the resource.`,
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

func computeUrlMapHostRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hosts": {
				Type:     schema.TypeSet,
				Required: true,
				Description: `The list of host patterns to match. They must be valid hostnames, except * will
match any string of ([a-z0-9-.]*). In that case, * must be the first character
and must be followed in the pattern by either - or ..`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"path_matcher": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The name of the PathMatcher to use to match the path portion of the URL if the
hostRule matches the URL's host portion.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `An optional description of this resource. Provide this property when you create
the resource.`,
			},
		},
	}
}

func resourceComputeUrlMapCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	defaultServiceProp, err := expandComputeUrlMapDefaultService(d.Get("default_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_service"); !isEmptyValue(reflect.ValueOf(defaultServiceProp)) && (ok || !reflect.DeepEqual(v, defaultServiceProp)) {
		obj["defaultService"] = defaultServiceProp
	}
	descriptionProp, err := expandComputeUrlMapDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeUrlMapFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	hostRulesProp, err := expandComputeUrlMapHostRule(d.Get("host_rule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("host_rule"); !isEmptyValue(reflect.ValueOf(hostRulesProp)) && (ok || !reflect.DeepEqual(v, hostRulesProp)) {
		obj["hostRules"] = hostRulesProp
	}
	nameProp, err := expandComputeUrlMapName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	pathMatchersProp, err := expandComputeUrlMapPathMatcher(d.Get("path_matcher"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("path_matcher"); !isEmptyValue(reflect.ValueOf(pathMatchersProp)) && (ok || !reflect.DeepEqual(v, pathMatchersProp)) {
		obj["pathMatchers"] = pathMatchersProp
	}
	testsProp, err := expandComputeUrlMapTest(d.Get("test"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("test"); !isEmptyValue(reflect.ValueOf(testsProp)) && (ok || !reflect.DeepEqual(v, testsProp)) {
		obj["tests"] = testsProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/urlMaps")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new UrlMap: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating UrlMap: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/global/urlMaps/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	waitErr := computeOperationWaitTime(
		config, res, project, "Creating UrlMap",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create UrlMap: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating UrlMap %q: %#v", d.Id(), res)

	return resourceComputeUrlMapRead(d, meta)
}

func resourceComputeUrlMapRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/urlMaps/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeUrlMap %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}

	if err := d.Set("map_id", flattenComputeUrlMapMapId(res["id"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeUrlMapCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("default_service", flattenComputeUrlMapDefaultService(res["defaultService"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("description", flattenComputeUrlMapDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("fingerprint", flattenComputeUrlMapFingerprint(res["fingerprint"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("host_rule", flattenComputeUrlMapHostRule(res["hostRules"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("name", flattenComputeUrlMapName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("path_matcher", flattenComputeUrlMapPathMatcher(res["pathMatchers"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("test", flattenComputeUrlMapTest(res["tests"], d)); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading UrlMap: %s", err)
	}

	return nil
}

func resourceComputeUrlMapUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	defaultServiceProp, err := expandComputeUrlMapDefaultService(d.Get("default_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_service"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, defaultServiceProp)) {
		obj["defaultService"] = defaultServiceProp
	}
	descriptionProp, err := expandComputeUrlMapDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeUrlMapFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	hostRulesProp, err := expandComputeUrlMapHostRule(d.Get("host_rule"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("host_rule"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, hostRulesProp)) {
		obj["hostRules"] = hostRulesProp
	}
	nameProp, err := expandComputeUrlMapName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	pathMatchersProp, err := expandComputeUrlMapPathMatcher(d.Get("path_matcher"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("path_matcher"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, pathMatchersProp)) {
		obj["pathMatchers"] = pathMatchersProp
	}
	testsProp, err := expandComputeUrlMapTest(d.Get("test"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("test"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, testsProp)) {
		obj["tests"] = testsProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/urlMaps/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating UrlMap %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating UrlMap %q: %s", d.Id(), err)
	}

	err = computeOperationWaitTime(
		config, res, project, "Updating UrlMap",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceComputeUrlMapRead(d, meta)
}

func resourceComputeUrlMapDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/urlMaps/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting UrlMap %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "UrlMap")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting UrlMap",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting UrlMap %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeUrlMapImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/urlMaps/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/global/urlMaps/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeUrlMapMapId(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeUrlMapCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapDefaultService(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeUrlMapDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapFingerprint(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapHostRule(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(computeUrlMapHostRuleSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"description":  flattenComputeUrlMapHostRuleDescription(original["description"], d),
			"hosts":        flattenComputeUrlMapHostRuleHosts(original["hosts"], d),
			"path_matcher": flattenComputeUrlMapHostRulePathMatcher(original["pathMatcher"], d),
		})
	}
	return transformed
}
func flattenComputeUrlMapHostRuleDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapHostRuleHosts(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeUrlMapHostRulePathMatcher(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapPathMatcher(v interface{}, d *schema.ResourceData) interface{} {
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
			"default_service": flattenComputeUrlMapPathMatcherDefaultService(original["defaultService"], d),
			"description":     flattenComputeUrlMapPathMatcherDescription(original["description"], d),
			"path_rule":       flattenComputeUrlMapPathMatcherPathRule(original["pathRules"], d),
			"name":            flattenComputeUrlMapPathMatcherName(original["name"], d),
		})
	}
	return transformed
}
func flattenComputeUrlMapPathMatcherDefaultService(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeUrlMapPathMatcherDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapPathMatcherPathRule(v interface{}, d *schema.ResourceData) interface{} {
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
			"paths":   flattenComputeUrlMapPathMatcherPathRulePaths(original["paths"], d),
			"service": flattenComputeUrlMapPathMatcherPathRuleService(original["service"], d),
		})
	}
	return transformed
}
func flattenComputeUrlMapPathMatcherPathRulePaths(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeUrlMapPathMatcherPathRuleService(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeUrlMapPathMatcherName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapTest(v interface{}, d *schema.ResourceData) interface{} {
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
			"description": flattenComputeUrlMapTestDescription(original["description"], d),
			"host":        flattenComputeUrlMapTestHost(original["host"], d),
			"path":        flattenComputeUrlMapTestPath(original["path"], d),
			"service":     flattenComputeUrlMapTestService(original["service"], d),
		})
	}
	return transformed
}
func flattenComputeUrlMapTestDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapTestHost(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapTestPath(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeUrlMapTestService(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeUrlMapDefaultService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapHostRule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDescription, err := expandComputeUrlMapHostRuleDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedHosts, err := expandComputeUrlMapHostRuleHosts(original["hosts"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedHosts); val.IsValid() && !isEmptyValue(val) {
			transformed["hosts"] = transformedHosts
		}

		transformedPathMatcher, err := expandComputeUrlMapHostRulePathMatcher(original["path_matcher"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPathMatcher); val.IsValid() && !isEmptyValue(val) {
			transformed["pathMatcher"] = transformedPathMatcher
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeUrlMapHostRuleDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapHostRuleHosts(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeUrlMapHostRulePathMatcher(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapPathMatcher(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDefaultService, err := expandComputeUrlMapPathMatcherDefaultService(original["default_service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDefaultService); val.IsValid() && !isEmptyValue(val) {
			transformed["defaultService"] = transformedDefaultService
		}

		transformedDescription, err := expandComputeUrlMapPathMatcherDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedPathRule, err := expandComputeUrlMapPathMatcherPathRule(original["path_rule"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPathRule); val.IsValid() && !isEmptyValue(val) {
			transformed["pathRules"] = transformedPathRule
		}

		transformedName, err := expandComputeUrlMapPathMatcherName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeUrlMapPathMatcherDefaultService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapPathMatcherDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapPathMatcherPathRule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPaths, err := expandComputeUrlMapPathMatcherPathRulePaths(original["paths"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPaths); val.IsValid() && !isEmptyValue(val) {
			transformed["paths"] = transformedPaths
		}

		transformedService, err := expandComputeUrlMapPathMatcherPathRuleService(original["service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
			transformed["service"] = transformedService
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeUrlMapPathMatcherPathRulePaths(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeUrlMapPathMatcherPathRuleService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapPathMatcherName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapTest(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDescription, err := expandComputeUrlMapTestDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedHost, err := expandComputeUrlMapTestHost(original["host"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !isEmptyValue(val) {
			transformed["host"] = transformedHost
		}

		transformedPath, err := expandComputeUrlMapTestPath(original["path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
			transformed["path"] = transformedPath
		}

		transformedService, err := expandComputeUrlMapTestService(original["service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
			transformed["service"] = transformedService
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeUrlMapTestDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapTestHost(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapTestPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeUrlMapTestService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
