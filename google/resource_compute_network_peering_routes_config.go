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
)

func resourceComputeNetworkPeeringRoutesConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNetworkPeeringRoutesConfigCreate,
		Read:   resourceComputeNetworkPeeringRoutesConfigRead,
		Update: resourceComputeNetworkPeeringRoutesConfigUpdate,
		Delete: resourceComputeNetworkPeeringRoutesConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeNetworkPeeringRoutesConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(6 * time.Minute),
			Update: schema.DefaultTimeout(6 * time.Minute),
			Delete: schema.DefaultTimeout(6 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"export_custom_routes": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: `Whether to export the custom routes to the peer network.`,
			},
			"import_custom_routes": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: `Whether to import the custom routes to the peer network.`,
			},
			"network": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The name of the primary network for the peering.`,
			},
			"peering": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Name of the peering.`,
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

func resourceComputeNetworkPeeringRoutesConfigCreate(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandComputeNetworkPeeringRoutesConfigPeering(d.Get("peering"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peering"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	exportCustomRoutesProp, err := expandComputeNetworkPeeringRoutesConfigExportCustomRoutes(d.Get("export_custom_routes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("export_custom_routes"); !isEmptyValue(reflect.ValueOf(exportCustomRoutesProp)) && (ok || !reflect.DeepEqual(v, exportCustomRoutesProp)) {
		obj["exportCustomRoutes"] = exportCustomRoutesProp
	}
	importCustomRoutesProp, err := expandComputeNetworkPeeringRoutesConfigImportCustomRoutes(d.Get("import_custom_routes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("import_custom_routes"); !isEmptyValue(reflect.ValueOf(importCustomRoutesProp)) && (ok || !reflect.DeepEqual(v, importCustomRoutesProp)) {
		obj["importCustomRoutes"] = importCustomRoutesProp
	}

	obj, err = resourceComputeNetworkPeeringRoutesConfigEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "projects/{{project}}/global/networks/{{network}}/peerings")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networks/{{network}}/updatePeering")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NetworkPeeringRoutesConfig: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating NetworkPeeringRoutesConfig: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/global/networks/{{network}}/networkPeerings/{{peering}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating NetworkPeeringRoutesConfig",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create NetworkPeeringRoutesConfig: %s", err)
	}

	log.Printf("[DEBUG] Finished creating NetworkPeeringRoutesConfig %q: %#v", d.Id(), res)

	return resourceComputeNetworkPeeringRoutesConfigRead(d, meta)
}

func resourceComputeNetworkPeeringRoutesConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networks/{{network}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeNetworkPeeringRoutesConfig %q", d.Id()))
	}

	res, err = flattenNestedComputeNetworkPeeringRoutesConfig(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeNetworkPeeringRoutesConfig because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading NetworkPeeringRoutesConfig: %s", err)
	}

	if err := d.Set("peering", flattenComputeNetworkPeeringRoutesConfigPeering(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPeeringRoutesConfig: %s", err)
	}
	if err := d.Set("export_custom_routes", flattenComputeNetworkPeeringRoutesConfigExportCustomRoutes(res["exportCustomRoutes"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPeeringRoutesConfig: %s", err)
	}
	if err := d.Set("import_custom_routes", flattenComputeNetworkPeeringRoutesConfigImportCustomRoutes(res["importCustomRoutes"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkPeeringRoutesConfig: %s", err)
	}

	return nil
}

func resourceComputeNetworkPeeringRoutesConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeNetworkPeeringRoutesConfigPeering(d.Get("peering"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peering"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	exportCustomRoutesProp, err := expandComputeNetworkPeeringRoutesConfigExportCustomRoutes(d.Get("export_custom_routes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("export_custom_routes"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, exportCustomRoutesProp)) {
		obj["exportCustomRoutes"] = exportCustomRoutesProp
	}
	importCustomRoutesProp, err := expandComputeNetworkPeeringRoutesConfigImportCustomRoutes(d.Get("import_custom_routes"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("import_custom_routes"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, importCustomRoutesProp)) {
		obj["importCustomRoutes"] = importCustomRoutesProp
	}

	obj, err = resourceComputeNetworkPeeringRoutesConfigEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "projects/{{project}}/global/networks/{{network}}/peerings")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/networks/{{network}}/updatePeering")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating NetworkPeeringRoutesConfig %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PATCH", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating NetworkPeeringRoutesConfig %q: %s", d.Id(), err)
	}

	err = computeOperationWaitTime(
		config, res, project, "Updating NetworkPeeringRoutesConfig",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceComputeNetworkPeeringRoutesConfigRead(d, meta)
}

func resourceComputeNetworkPeeringRoutesConfigDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Compute NetworkPeeringRoutesConfig resources"+
		" cannot be deleted from GCP. The resource %s will be removed from Terraform"+
		" state, but will still be present on the server.", d.Id())
	d.SetId("")

	return nil
}

func resourceComputeNetworkPeeringRoutesConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/networks/(?P<network>[^/]+)/networkPeerings/(?P<peering>[^/]+)",
		"(?P<project>[^/]+)/(?P<network>[^/]+)/(?P<peering>[^/]+)",
		"(?P<network>[^/]+)/(?P<peering>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/global/networks/{{network}}/networkPeerings/{{peering}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeNetworkPeeringRoutesConfigPeering(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeNetworkPeeringRoutesConfigExportCustomRoutes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeNetworkPeeringRoutesConfigImportCustomRoutes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandComputeNetworkPeeringRoutesConfigPeering(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkPeeringRoutesConfigExportCustomRoutes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkPeeringRoutesConfigImportCustomRoutes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceComputeNetworkPeeringRoutesConfigEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Stick request in a networkPeering block as in
	// https://cloud.google.com/compute/docs/reference/rest/v1/networks/updatePeering
	newObj := make(map[string]interface{})
	newObj["networkPeering"] = obj
	return newObj, nil
}

func flattenNestedComputeNetworkPeeringRoutesConfig(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["peerings"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value peerings. Actual value: %v", v)
	}

	_, item, err := resourceComputeNetworkPeeringRoutesConfigFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceComputeNetworkPeeringRoutesConfigFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedPeering, err := expandComputeNetworkPeeringRoutesConfigPeering(d.Get("peering"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemPeering := flattenComputeNetworkPeeringRoutesConfigPeering(item["name"], d, meta.(*Config))
		if !reflect.DeepEqual(itemPeering, expectedPeering) {
			log.Printf("[DEBUG] Skipping item with name= %#v, looking for %#v)", itemPeering, expectedPeering)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
