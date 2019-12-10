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
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceComputeNetworkEndpointGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNetworkEndpointGroupCreate,
		Read:   resourceComputeNetworkEndpointGroupRead,
		Delete: resourceComputeNetworkEndpointGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeNetworkEndpointGroupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateGCPName,
				Description: `Name of the resource; provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"network": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `The network to which all network endpoints in the NEG belong.
Uses "default" project network if unspecified.`,
			},
			"default_port": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Description: `The default port used if the port number is not specified in the
network endpoint.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `An optional description of this resource. Provide this property when
you create the resource.`,
			},
			"network_endpoint_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"GCE_VM_IP_PORT", ""}, false),
				Description: `Type of network endpoints in this network endpoint group. Currently
the only supported value is GCE_VM_IP_PORT.`,
				Default: "GCE_VM_IP_PORT",
			},
			"subnetwork": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `Optional subnetwork to which all network endpoints in the NEG belong.`,
			},
			"zone": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `Zone where the network endpoint group is located.`,
			},
			"size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Number of network endpoints in the network endpoint group.`,
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

func resourceComputeNetworkEndpointGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandComputeNetworkEndpointGroupName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeNetworkEndpointGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networkEndpointTypeProp, err := expandComputeNetworkEndpointGroupNetworkEndpointType(d.Get("network_endpoint_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network_endpoint_type"); !isEmptyValue(reflect.ValueOf(networkEndpointTypeProp)) && (ok || !reflect.DeepEqual(v, networkEndpointTypeProp)) {
		obj["networkEndpointType"] = networkEndpointTypeProp
	}
	networkProp, err := expandComputeNetworkEndpointGroupNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	subnetworkProp, err := expandComputeNetworkEndpointGroupSubnetwork(d.Get("subnetwork"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("subnetwork"); !isEmptyValue(reflect.ValueOf(subnetworkProp)) && (ok || !reflect.DeepEqual(v, subnetworkProp)) {
		obj["subnetwork"] = subnetworkProp
	}
	defaultPortProp, err := expandComputeNetworkEndpointGroupDefaultPort(d.Get("default_port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("default_port"); !isEmptyValue(reflect.ValueOf(defaultPortProp)) && (ok || !reflect.DeepEqual(v, defaultPortProp)) {
		obj["defaultPort"] = defaultPortProp
	}
	zoneProp, err := expandComputeNetworkEndpointGroupZone(d.Get("zone"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("zone"); !isEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NetworkEndpointGroup: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating NetworkEndpointGroup: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating NetworkEndpointGroup",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if err != nil {
		// Remove ID to show resource wasn't created.
		d.SetId("")
		return fmt.Errorf("Error waiting to create NetworkEndpointGroup: %s", err)
	}

	log.Printf("[DEBUG] Finished creating NetworkEndpointGroup %q: %#v", d.Id(), res)

	return resourceComputeNetworkEndpointGroupRead(d, meta)
}

func resourceComputeNetworkEndpointGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeNetworkEndpointGroup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}

	if err := d.Set("name", flattenComputeNetworkEndpointGroupName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("description", flattenComputeNetworkEndpointGroupDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("network_endpoint_type", flattenComputeNetworkEndpointGroupNetworkEndpointType(res["networkEndpointType"], d)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("size", flattenComputeNetworkEndpointGroupSize(res["size"], d)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("network", flattenComputeNetworkEndpointGroupNetwork(res["network"], d)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("subnetwork", flattenComputeNetworkEndpointGroupSubnetwork(res["subnetwork"], d)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("default_port", flattenComputeNetworkEndpointGroupDefaultPort(res["defaultPort"], d)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("zone", flattenComputeNetworkEndpointGroupZone(res["zone"], d)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading NetworkEndpointGroup: %s", err)
	}

	return nil
}

func resourceComputeNetworkEndpointGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting NetworkEndpointGroup %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "NetworkEndpointGroup")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting NetworkEndpointGroup",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting NetworkEndpointGroup %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeNetworkEndpointGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/networkEndpointGroups/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<zone>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeNetworkEndpointGroupName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeNetworkEndpointGroupDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeNetworkEndpointGroupNetworkEndpointType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeNetworkEndpointGroupSize(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeNetworkEndpointGroupNetwork(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeNetworkEndpointGroupSubnetwork(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeNetworkEndpointGroupDefaultPort(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeNetworkEndpointGroupZone(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeNetworkEndpointGroupName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkEndpointGroupDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkEndpointGroupNetworkEndpointType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkEndpointGroupNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeNetworkEndpointGroupSubnetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("subnetworks", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for subnetwork: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeNetworkEndpointGroupDefaultPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkEndpointGroupZone(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}
