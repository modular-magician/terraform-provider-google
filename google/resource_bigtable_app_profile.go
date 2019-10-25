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
	"google.golang.org/api/bigtableadmin/v2"
)

func resourceBigtableAppProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigtableAppProfileCreate,
		Read:   resourceBigtableAppProfileRead,
		Update: resourceBigtableAppProfileUpdate,
		Delete: resourceBigtableAppProfileDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBigtableAppProfileImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"app_profile_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ignore_warnings": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"instance": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"multi_cluster_routing_use_any": {
				Type:          schema.TypeBool,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"single_cluster_routing"},
			},
			"single_cluster_routing": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_transactional_writes": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"cluster_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				ConflictsWith: []string{"multi_cluster_routing_use_any"},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
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

func resourceBigtableAppProfileCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandBigtableAppProfileDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	multiClusterRoutingUseAnyProp, err := expandBigtableAppProfileMultiClusterRoutingUseAny(d.Get("multi_cluster_routing_use_any"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("multi_cluster_routing_use_any"); !isEmptyValue(reflect.ValueOf(multiClusterRoutingUseAnyProp)) && (ok || !reflect.DeepEqual(v, multiClusterRoutingUseAnyProp)) {
		obj["multiClusterRoutingUseAny"] = multiClusterRoutingUseAnyProp
	}
	singleClusterRoutingProp, err := expandBigtableAppProfileSingleClusterRouting(d.Get("single_cluster_routing"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("single_cluster_routing"); !isEmptyValue(reflect.ValueOf(singleClusterRoutingProp)) && (ok || !reflect.DeepEqual(v, singleClusterRoutingProp)) {
		obj["singleClusterRouting"] = singleClusterRoutingProp
	}

	url, err := replaceVars(d, config, "{{BigtableBasePath}}projects/{{project}}/instances/{{instance}}/appProfiles?appProfileId={{app_profile_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AppProfile: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating AppProfile: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/instances/{{instance}}/appProfiles/{{app_profile_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating AppProfile %q: %#v", d.Id(), res)

	return resourceBigtableAppProfileRead(d, meta)
}

func resourceBigtableAppProfileRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{BigtableBasePath}}projects/{{project}}/instances/{{instance}}/appProfiles/{{app_profile_id}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("BigtableAppProfile %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading AppProfile: %s", err)
	}

	if err := d.Set("name", flattenBigtableAppProfileName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading AppProfile: %s", err)
	}
	if err := d.Set("description", flattenBigtableAppProfileDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading AppProfile: %s", err)
	}
	if err := d.Set("multi_cluster_routing_use_any", flattenBigtableAppProfileMultiClusterRoutingUseAny(res["multiClusterRoutingUseAny"], d)); err != nil {
		return fmt.Errorf("Error reading AppProfile: %s", err)
	}
	if err := d.Set("single_cluster_routing", flattenBigtableAppProfileSingleClusterRouting(res["singleClusterRouting"], d)); err != nil {
		return fmt.Errorf("Error reading AppProfile: %s", err)
	}

	return nil
}

func resourceBigtableAppProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandBigtableAppProfileDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	url, err := replaceVars(d, config, "{{BigtableBasePath}}projects/{{project}}/instances/{{instance}}/appProfiles/{{app_profile_id}}?ignoreWarnings={{ignore_warnings}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AppProfile %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	_, err = sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating AppProfile %q: %s", d.Id(), err)
	}

	return resourceBigtableAppProfileRead(d, meta)
}

func resourceBigtableAppProfileDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{BigtableBasePath}}projects/{{project}}/instances/{{instance}}/appProfiles/{{app_profile_id}}?ignoreWarnings={{ignore_warnings}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting AppProfile %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "AppProfile")
	}

	log.Printf("[DEBUG] Finished deleting AppProfile %q: %#v", d.Id(), res)
	return nil
}

func resourceBigtableAppProfileImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/instances/(?P<instance>[^/]+)/appProfiles/(?P<app_profile_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<instance>[^/]+)/(?P<app_profile_id>[^/]+)",
		"(?P<instance>[^/]+)/(?P<app_profile_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/instances/{{instance}}/appProfiles/{{app_profile_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBigtableAppProfileName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigtableAppProfileDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigtableAppProfileMultiClusterRoutingUseAny(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return false
	}

	return true
}

func flattenBigtableAppProfileSingleClusterRouting(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["cluster_id"] =
		flattenBigtableAppProfileSingleClusterRoutingClusterId(original["clusterId"], d)
	transformed["allow_transactional_writes"] =
		flattenBigtableAppProfileSingleClusterRoutingAllowTransactionalWrites(original["allowTransactionalWrites"], d)
	return []interface{}{transformed}
}
func flattenBigtableAppProfileSingleClusterRoutingClusterId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenBigtableAppProfileSingleClusterRoutingAllowTransactionalWrites(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandBigtableAppProfileDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigtableAppProfileMultiClusterRoutingUseAny(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == nil || !v.(bool) {
		return nil, nil
	}

	return bigtableadmin.MultiClusterRoutingUseAny{}, nil
}

func expandBigtableAppProfileSingleClusterRouting(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedClusterId, err := expandBigtableAppProfileSingleClusterRoutingClusterId(original["cluster_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClusterId); val.IsValid() && !isEmptyValue(val) {
		transformed["clusterId"] = transformedClusterId
	}

	transformedAllowTransactionalWrites, err := expandBigtableAppProfileSingleClusterRoutingAllowTransactionalWrites(original["allow_transactional_writes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowTransactionalWrites); val.IsValid() && !isEmptyValue(val) {
		transformed["allowTransactionalWrites"] = transformedAllowTransactionalWrites
	}

	return transformed, nil
}

func expandBigtableAppProfileSingleClusterRoutingClusterId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigtableAppProfileSingleClusterRoutingAllowTransactionalWrites(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
