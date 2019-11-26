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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAccessContextManagerAccessLevel() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessContextManagerAccessLevelCreate,
		Read:   resourceAccessContextManagerAccessLevelRead,
		Update: resourceAccessContextManagerAccessLevelUpdate,
		Delete: resourceAccessContextManagerAccessLevelDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessContextManagerAccessLevelImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(6 * time.Minute),
			Update: schema.DefaultTimeout(6 * time.Minute),
			Delete: schema.DefaultTimeout(6 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Resource name for the Access Level. The short_name component must begin
with a letter and only include alphanumeric and '_'.
Format: accessPolicies/{policy_id}/accessLevels/{short_name}`,
			},
			"parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The AccessPolicy this AccessLevel lives in.
Format: accessPolicies/{policy_id}`,
			},
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Human readable title. Must be unique within the Policy.`,
			},
			"basic": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `A set of predefined conditions for the access level and a combining function.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"conditions": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `A set of requirements for the AccessLevel to be granted.`,
							MinItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_policy": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `Device specific restrictions, all restrictions must hold for
the Condition to be true. If not specified, all devices are
allowed.`,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"allowed_device_management_levels": {
													Type:     schema.TypeList,
													Optional: true,
													Description: `A list of allowed device management levels.
An empty list allows all management levels.`,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"allowed_encryption_statuses": {
													Type:     schema.TypeList,
													Optional: true,
													Description: `A list of allowed encryptions statuses.
An empty list allows all statuses.`,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"os_constraints": {
													Type:     schema.TypeList,
													Optional: true,
													Description: `A list of allowed OS versions.
An empty list allows all types and all versions.`,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"os_type": {
																Type:         schema.TypeString,
																Required:     true,
																ValidateFunc: validation.StringInSlice([]string{"OS_UNSPECIFIED", "DESKTOP_MAC", "DESKTOP_WINDOWS", "DESKTOP_LINUX", "DESKTOP_CHROME_OS"}, false),
																Description:  `The operating system type of the device.`,
															},
															"minimum_version": {
																Type:     schema.TypeString,
																Optional: true,
																Description: `The minimum allowed OS version. If not set, any version
of this OS satisfies the constraint.
Format: "major.minor.patch" such as "10.5.301", "9.2.1".`,
															},
														},
													},
												},
												"require_admin_approval": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: `Whether the device needs to be approved by the customer admin.`,
												},
												"require_corp_owned": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: `Whether the device needs to be corp owned.`,
												},
												"require_screen_lock": {
													Type:     schema.TypeBool,
													Optional: true,
													Description: `Whether or not screenlock is required for the DevicePolicy
to be true. Defaults to false.`,
												},
											},
										},
									},
									"ip_subnetworks": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `A list of CIDR block IP subnetwork specification. May be IPv4
or IPv6.
Note that for a CIDR IP address block, the specified IP address
portion must be properly truncated (i.e. all the host bits must
be zero) or the input is considered malformed. For example,
"192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
is not. The originating IP of a request must be in one of the
listed subnets in order for this Condition to be true.
If empty, all IP addresses are allowed.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"members": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `An allowed list of members (users, groups, service accounts).
The signed-in user originating the request must be a part of one
of the provided members. If not specified, a request may come
from any user (logged in/not logged in, not present in any
groups, etc.).
Formats: 'user:{emailid}', 'group:{emailid}', 'serviceAccount:{emailid}'`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"negate": {
										Type:     schema.TypeBool,
										Optional: true,
										Description: `Whether to negate the Condition. If true, the Condition becomes
a NAND over its non-empty fields, each field must be false for
the Condition overall to be satisfied. Defaults to false.`,
									},
									"required_access_levels": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `A list of other access levels defined in the same Policy,
referenced by resource name. Referencing an AccessLevel which
does not exist is an error. All access levels listed must be
granted for the Condition to be true.
Format: accessPolicies/{policy_id}/accessLevels/{short_name}`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"combining_function": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"AND", "OR", ""}, false),
							Description: `How the conditions list should be combined to determine if a request
is granted this AccessLevel. If AND is used, each Condition in
conditions must be satisfied for the AccessLevel to be applied. If
OR is used, at least one Condition in conditions must be satisfied
for the AccessLevel to be applied. Defaults to AND if unspecified.`,
							Default: "AND",
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Description of the AccessLevel and its use. Does not affect behavior.`,
			},
		},
	}
}

func resourceAccessContextManagerAccessLevelCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	titleProp, err := expandAccessContextManagerAccessLevelTitle(d.Get("title"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("title"); !isEmptyValue(reflect.ValueOf(titleProp)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}
	descriptionProp, err := expandAccessContextManagerAccessLevelDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	basicProp, err := expandAccessContextManagerAccessLevelBasic(d.Get("basic"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("basic"); !isEmptyValue(reflect.ValueOf(basicProp)) && (ok || !reflect.DeepEqual(v, basicProp)) {
		obj["basic"] = basicProp
	}
	parentProp, err := expandAccessContextManagerAccessLevelParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	nameProp, err := expandAccessContextManagerAccessLevelName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	obj, err = resourceAccessContextManagerAccessLevelEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{parent}}/accessLevels")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AccessLevel: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", "", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating AccessLevel: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	waitErr := accessContextManagerOperationWaitTime(
		config, res, "Creating AccessLevel",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create AccessLevel: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating AccessLevel %q: %#v", d.Id(), res)

	return resourceAccessContextManagerAccessLevelRead(d, meta)
}

func resourceAccessContextManagerAccessLevelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", "", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AccessContextManagerAccessLevel %q", d.Id()))
	}

	if err := d.Set("title", flattenAccessContextManagerAccessLevelTitle(res["title"], d)); err != nil {
		return fmt.Errorf("Error reading AccessLevel: %s", err)
	}
	if err := d.Set("description", flattenAccessContextManagerAccessLevelDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading AccessLevel: %s", err)
	}
	if err := d.Set("basic", flattenAccessContextManagerAccessLevelBasic(res["basic"], d)); err != nil {
		return fmt.Errorf("Error reading AccessLevel: %s", err)
	}
	if err := d.Set("name", flattenAccessContextManagerAccessLevelName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading AccessLevel: %s", err)
	}

	return nil
}

func resourceAccessContextManagerAccessLevelUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	titleProp, err := expandAccessContextManagerAccessLevelTitle(d.Get("title"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("title"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}
	descriptionProp, err := expandAccessContextManagerAccessLevelDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	basicProp, err := expandAccessContextManagerAccessLevelBasic(d.Get("basic"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("basic"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, basicProp)) {
		obj["basic"] = basicProp
	}

	obj, err = resourceAccessContextManagerAccessLevelEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AccessLevel %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("title") {
		updateMask = append(updateMask, "title")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("basic") {
		updateMask = append(updateMask, "basic")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", "", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating AccessLevel %q: %s", d.Id(), err)
	}

	err = accessContextManagerOperationWaitTime(
		config, res, "Updating AccessLevel",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))
	if err != nil {
		return err
	}

	return resourceAccessContextManagerAccessLevelRead(d, meta)
}

func resourceAccessContextManagerAccessLevelDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting AccessLevel %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", "", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "AccessLevel")
	}

	err = accessContextManagerOperationWaitTime(
		config, res, "Deleting AccessLevel",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting AccessLevel %q: %#v", d.Id(), res)
	return nil
}

func resourceAccessContextManagerAccessLevelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}
	stringParts := strings.Split(d.Get("name").(string), "/")
	if len(stringParts) < 2 {
		return nil, fmt.Errorf("Error parsing parent name. Should be in form accessPolicies/{{policy_id}}/accessLevels/{{short_name}}")
	}
	d.Set("parent", fmt.Sprintf("%s/%s", stringParts[0], stringParts[1]))
	return []*schema.ResourceData{d}, nil
}

func flattenAccessContextManagerAccessLevelTitle(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasic(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["combining_function"] =
		flattenAccessContextManagerAccessLevelBasicCombiningFunction(original["combiningFunction"], d)
	transformed["conditions"] =
		flattenAccessContextManagerAccessLevelBasicConditions(original["conditions"], d)
	return []interface{}{transformed}
}
func flattenAccessContextManagerAccessLevelBasicCombiningFunction(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil || v.(string) == "" {
		return "AND"
	}
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditions(v interface{}, d *schema.ResourceData) interface{} {
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
			"ip_subnetworks":         flattenAccessContextManagerAccessLevelBasicConditionsIpSubnetworks(original["ipSubnetworks"], d),
			"required_access_levels": flattenAccessContextManagerAccessLevelBasicConditionsRequiredAccessLevels(original["requiredAccessLevels"], d),
			"members":                flattenAccessContextManagerAccessLevelBasicConditionsMembers(original["members"], d),
			"negate":                 flattenAccessContextManagerAccessLevelBasicConditionsNegate(original["negate"], d),
			"device_policy":          flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicy(original["devicePolicy"], d),
		})
	}
	return transformed
}
func flattenAccessContextManagerAccessLevelBasicConditionsIpSubnetworks(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsRequiredAccessLevels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsMembers(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsNegate(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicy(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["require_screen_lock"] =
		flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireScreenLock(original["requireScreenLock"], d)
	transformed["allowed_encryption_statuses"] =
		flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatuses(original["allowedEncryptionStatuses"], d)
	transformed["allowed_device_management_levels"] =
		flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevels(original["allowedDeviceManagementLevels"], d)
	transformed["os_constraints"] =
		flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraints(original["osConstraints"], d)
	transformed["require_admin_approval"] =
		flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireAdminApproval(original["requireAdminApproval"], d)
	transformed["require_corp_owned"] =
		flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireCorpOwned(original["requireCorpOwned"], d)
	return []interface{}{transformed}
}
func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireScreenLock(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatuses(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraints(v interface{}, d *schema.ResourceData) interface{} {
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
			"minimum_version": flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsMinimumVersion(original["minimumVersion"], d),
			"os_type":         flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOsType(original["osType"], d),
		})
	}
	return transformed
}
func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsMinimumVersion(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOsType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireAdminApproval(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireCorpOwned(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenAccessContextManagerAccessLevelName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandAccessContextManagerAccessLevelTitle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCombiningFunction, err := expandAccessContextManagerAccessLevelBasicCombiningFunction(original["combining_function"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCombiningFunction); val.IsValid() && !isEmptyValue(val) {
		transformed["combiningFunction"] = transformedCombiningFunction
	}

	transformedConditions, err := expandAccessContextManagerAccessLevelBasicConditions(original["conditions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConditions); val.IsValid() && !isEmptyValue(val) {
		transformed["conditions"] = transformedConditions
	}

	return transformed, nil
}

func expandAccessContextManagerAccessLevelBasicCombiningFunction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpSubnetworks, err := expandAccessContextManagerAccessLevelBasicConditionsIpSubnetworks(original["ip_subnetworks"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpSubnetworks); val.IsValid() && !isEmptyValue(val) {
			transformed["ipSubnetworks"] = transformedIpSubnetworks
		}

		transformedRequiredAccessLevels, err := expandAccessContextManagerAccessLevelBasicConditionsRequiredAccessLevels(original["required_access_levels"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRequiredAccessLevels); val.IsValid() && !isEmptyValue(val) {
			transformed["requiredAccessLevels"] = transformedRequiredAccessLevels
		}

		transformedMembers, err := expandAccessContextManagerAccessLevelBasicConditionsMembers(original["members"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMembers); val.IsValid() && !isEmptyValue(val) {
			transformed["members"] = transformedMembers
		}

		transformedNegate, err := expandAccessContextManagerAccessLevelBasicConditionsNegate(original["negate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNegate); val.IsValid() && !isEmptyValue(val) {
			transformed["negate"] = transformedNegate
		}

		transformedDevicePolicy, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicy(original["device_policy"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDevicePolicy); val.IsValid() && !isEmptyValue(val) {
			transformed["devicePolicy"] = transformedDevicePolicy
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsIpSubnetworks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsRequiredAccessLevels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsMembers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsNegate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRequireScreenLock, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireScreenLock(original["require_screen_lock"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireScreenLock); val.IsValid() && !isEmptyValue(val) {
		transformed["requireScreenLock"] = transformedRequireScreenLock
	}

	transformedAllowedEncryptionStatuses, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatuses(original["allowed_encryption_statuses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedEncryptionStatuses); val.IsValid() && !isEmptyValue(val) {
		transformed["allowedEncryptionStatuses"] = transformedAllowedEncryptionStatuses
	}

	transformedAllowedDeviceManagementLevels, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevels(original["allowed_device_management_levels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedDeviceManagementLevels); val.IsValid() && !isEmptyValue(val) {
		transformed["allowedDeviceManagementLevels"] = transformedAllowedDeviceManagementLevels
	}

	transformedOsConstraints, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraints(original["os_constraints"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOsConstraints); val.IsValid() && !isEmptyValue(val) {
		transformed["osConstraints"] = transformedOsConstraints
	}

	transformedRequireAdminApproval, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireAdminApproval(original["require_admin_approval"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireAdminApproval); val.IsValid() && !isEmptyValue(val) {
		transformed["requireAdminApproval"] = transformedRequireAdminApproval
	}

	transformedRequireCorpOwned, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireCorpOwned(original["require_corp_owned"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireCorpOwned); val.IsValid() && !isEmptyValue(val) {
		transformed["requireCorpOwned"] = transformedRequireCorpOwned
	}

	return transformed, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireScreenLock(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedEncryptionStatuses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyAllowedDeviceManagementLevels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraints(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMinimumVersion, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsMinimumVersion(original["minimum_version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMinimumVersion); val.IsValid() && !isEmptyValue(val) {
			transformed["minimumVersion"] = transformedMinimumVersion
		}

		transformedOsType, err := expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOsType(original["os_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOsType); val.IsValid() && !isEmptyValue(val) {
			transformed["osType"] = transformedOsType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsMinimumVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOsType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireAdminApproval(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelBasicConditionsDevicePolicyRequireCorpOwned(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelParent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessLevelName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceAccessContextManagerAccessLevelEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "parent")
	return obj, nil
}
