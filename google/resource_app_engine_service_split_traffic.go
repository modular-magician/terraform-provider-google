// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceAppEngineServiceSplitTraffic() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppEngineServiceSplitTrafficCreate,
		Read:   resourceAppEngineServiceSplitTrafficRead,
		Update: resourceAppEngineServiceSplitTrafficUpdate,
		Delete: resourceAppEngineServiceSplitTrafficDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAppEngineServiceSplitTrafficImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The name of the service these settings apply to.`,
			},
			"split": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Mapping that defines fractional HTTP traffic diversion to different versions within the service.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allocations": {
							Type:        schema.TypeMap,
							Required:    true,
							Description: `Mapping from version IDs within the service to fractional (0.000, 1] allocations of traffic for that version. Each version can be specified only once, but some versions in the service may not have any traffic allocation. Services that have traffic allocated cannot be deleted until either the service is deleted or their traffic allocation is removed. Allocations must sum to 1. Up to two decimal place precision is supported for IP-based splits and up to three decimal places is supported for cookie-based splits.`,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"shard_by": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"UNSPECIFIED", "COOKIE", "IP", "RANDOM", ""}),
							Description:  `Mechanism used to determine which version a request is sent to. The traffic selection algorithm will be stable for either type until allocations are changed. Possible values: ["UNSPECIFIED", "COOKIE", "IP", "RANDOM"]`,
						},
					},
				},
			},
			"migrate_traffic": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `If set to true traffic will be migrated to this version.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceAppEngineServiceSplitTrafficCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	idProp, err := expandAppEngineServiceSplitTrafficService(d.Get("service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service"); !isEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	splitProp, err := expandAppEngineServiceSplitTrafficSplit(d.Get("split"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("split"); !isEmptyValue(reflect.ValueOf(splitProp)) && (ok || !reflect.DeepEqual(v, splitProp)) {
		obj["split"] = splitProp
	}

	lockName, err := ReplaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/services/{{service}}?migrateTraffic={{migrate_traffic}}&updateMask=split")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ServiceSplitTraffic: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceSplitTraffic: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ServiceSplitTraffic: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "apps/{{project}}/services/{{service}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = AppEngineOperationWaitTime(
		config, res, project, "Creating ServiceSplitTraffic", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ServiceSplitTraffic: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ServiceSplitTraffic %q: %#v", d.Id(), res)

	return resourceAppEngineServiceSplitTrafficRead(d, meta)
}

func resourceAppEngineServiceSplitTrafficRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/services/{{service}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceSplitTraffic: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("AppEngineServiceSplitTraffic %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ServiceSplitTraffic: %s", err)
	}

	if err := d.Set("service", flattenAppEngineServiceSplitTrafficService(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceSplitTraffic: %s", err)
	}

	return nil
}

func resourceAppEngineServiceSplitTrafficUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceSplitTraffic: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	idProp, err := expandAppEngineServiceSplitTrafficService(d.Get("service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	splitProp, err := expandAppEngineServiceSplitTrafficSplit(d.Get("split"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("split"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, splitProp)) {
		obj["split"] = splitProp
	}

	lockName, err := ReplaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := ReplaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/services/{{service}}?migrateTraffic={{migrate_traffic}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ServiceSplitTraffic %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("service") {
		updateMask = append(updateMask, "id")
	}

	if d.HasChange("split") {
		updateMask = append(updateMask, "split")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating ServiceSplitTraffic %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ServiceSplitTraffic %q: %#v", d.Id(), res)
	}

	err = AppEngineOperationWaitTime(
		config, res, project, "Updating ServiceSplitTraffic", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceAppEngineServiceSplitTrafficRead(d, meta)
}

func resourceAppEngineServiceSplitTrafficDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] AppEngine ServiceSplitTraffic resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceAppEngineServiceSplitTrafficImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"apps/(?P<project>[^/]+)/services/(?P<service>[^/]+)",
		"(?P<project>[^/]+)/(?P<service>[^/]+)",
		"(?P<service>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "apps/{{project}}/services/{{service}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAppEngineServiceSplitTrafficService(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandAppEngineServiceSplitTrafficService(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineServiceSplitTrafficSplit(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedShardBy, err := expandAppEngineServiceSplitTrafficSplitShardBy(original["shard_by"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedShardBy); val.IsValid() && !isEmptyValue(val) {
		transformed["shardBy"] = transformedShardBy
	}

	transformedAllocations, err := expandAppEngineServiceSplitTrafficSplitAllocations(original["allocations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllocations); val.IsValid() && !isEmptyValue(val) {
		transformed["allocations"] = transformedAllocations
	}

	return transformed, nil
}

func expandAppEngineServiceSplitTrafficSplitShardBy(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineServiceSplitTrafficSplitAllocations(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
