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
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceRedisInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceRedisInstanceCreate,
		Read:   resourceRedisInstanceRead,
		Delete: resourceRedisInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceRedisInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(360 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"memory_size_gb": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"alternative_location_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"authorized_network": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"location_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"redis_version": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"reserved_ip_range": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"tier": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"BASIC", "STANDARD_HA", ""}, false),
				Default:      "BASIC",
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_location_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
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

func resourceRedisInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	alternativeLocationIdProp, err := expandRedisInstanceAlternativeLocationId(d.Get("alternative_location_id"), d, config)
	if err != nil {
		return err
	}
	authorizedNetworkProp, err := expandRedisInstanceAuthorizedNetwork(d.Get("authorized_network"), d, config)
	if err != nil {
		return err
	}
	displayNameProp, err := expandRedisInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	}
	labelsProp, err := expandRedisInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	}
	locationIdProp, err := expandRedisInstanceLocationId(d.Get("location_id"), d, config)
	if err != nil {
		return err
	}
	nameProp, err := expandRedisInstanceName(d.Get("name"), d, config)
	if err != nil {
		return err
	}
	memorySizeGbProp, err := expandRedisInstanceMemorySizeGb(d.Get("memory_size_gb"), d, config)
	if err != nil {
		return err
	}
	redisVersionProp, err := expandRedisInstanceRedisVersion(d.Get("redis_version"), d, config)
	if err != nil {
		return err
	}
	reservedIpRangeProp, err := expandRedisInstanceReservedIpRange(d.Get("reserved_ip_range"), d, config)
	if err != nil {
		return err
	}
	tierProp, err := expandRedisInstanceTier(d.Get("tier"), d, config)
	if err != nil {
		return err
	}
	regionProp, err := expandRedisInstanceRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"alternativeLocationId": alternativeLocationIdProp,
		"authorizedNetwork":     authorizedNetworkProp,
		"displayName":           displayNameProp,
		"labels":                labelsProp,
		"locationId":            locationIdProp,
		"name":                  nameProp,
		"memorySizeGb":          memorySizeGbProp,
		"redisVersion":          redisVersionProp,
		"reservedIpRange":       reservedIpRangeProp,
		"tier":                  tierProp,
		"region":                regionProp,
	}
	obj, err = resourceRedisInstanceEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://redis.googleapis.com/v1beta1/projects/{{project}}/locations/{{region}}/instances?instanceId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Instance: %#v", obj)
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating Instance: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{project}}/{{region}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	op := &redis.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := redisOperationWaitTime(
		config.clientRedis, op, project, "Creating Instance",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Instance: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceRedisInstanceRead(d, meta)
}

func resourceRedisInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://redis.googleapis.com/v1beta1/projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return err
	}

	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("RedisInstance %q", d.Id()))
	}

	res, err = resourceRedisInstanceDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if err := d.Set("alternative_location_id", flattenRedisInstanceAlternativeLocationId(res["alternativeLocationId"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("authorized_network", flattenRedisInstanceAuthorizedNetwork(res["authorizedNetwork"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("create_time", flattenRedisInstanceCreateTime(res["createTime"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("current_location_id", flattenRedisInstanceCurrentLocationId(res["currentLocationId"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("display_name", flattenRedisInstanceDisplayName(res["displayName"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("host", flattenRedisInstanceHost(res["host"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("labels", flattenRedisInstanceLabels(res["labels"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("location_id", flattenRedisInstanceLocationId(res["locationId"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("name", flattenRedisInstanceName(res["name"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("memory_size_gb", flattenRedisInstanceMemorySizeGb(res["memorySizeGb"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("port", flattenRedisInstancePort(res["port"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("redis_version", flattenRedisInstanceRedisVersion(res["redisVersion"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("reserved_ip_range", flattenRedisInstanceReservedIpRange(res["reservedIpRange"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("tier", flattenRedisInstanceTier(res["tier"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("region", flattenRedisInstanceRegion(res["region"])); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	return nil
}

func resourceRedisInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://redis.googleapis.com/v1beta1/projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting Instance %q", d.Id())
	res, err := Delete(config, url)
	if err != nil {
		return fmt.Errorf("Error deleting Instance %q: %s", d.Id(), err)
	}

	op := &redis.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = redisOperationWaitTime(
		config.clientRedis, op, project, "Deleting Instance",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id(), res)
	return nil
}

func resourceRedisInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/instances/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{project}}/{{region}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenRedisInstanceAlternativeLocationId(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceAuthorizedNetwork(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceCreateTime(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceCurrentLocationId(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceDisplayName(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceHost(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceLabels(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceLocationId(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceName(v interface{}) interface{} {
	return NameFromSelfLinkStateFunc(v)
}

func flattenRedisInstanceMemorySizeGb(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenRedisInstancePort(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenRedisInstanceRedisVersion(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceReservedIpRange(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceTier(v interface{}) interface{} {
	return v
}

func flattenRedisInstanceRegion(v interface{}) interface{} {
	return v
}

func expandRedisInstanceAlternativeLocationId(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceAuthorizedNetwork(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceDisplayName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceLabels(v interface{}, d *schema.ResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandRedisInstanceLocationId(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	project, err := getProject(d, config)
	if err != nil {
		return nil, err
	}

	region, err := getRegion(d, config)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("projects/%s/locations/%s/instances/%s", project, region, v.(string)), nil
}

func expandRedisInstanceMemorySizeGb(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceRedisVersion(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceReservedIpRange(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceTier(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandRedisInstanceRegion(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceRedisInstanceEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "region")
	return obj, nil
}

func resourceRedisInstanceDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)
	region, err := getRegion(d, config)
	if err != nil {
		return nil, err
	}
	res["region"] = region
	return res, nil
}
