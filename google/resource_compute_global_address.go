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

func resourceComputeGlobalAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeGlobalAddressCreate,
		Read:   resourceComputeGlobalAddressRead,
		Delete: resourceComputeGlobalAddressDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeGlobalAddressImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"address_type": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateFunc:     validation.StringInSlice([]string{"EXTERNAL", "INTERNAL", ""}, false),
				DiffSuppressFunc: emptyOrDefaultStringSuppress("EXTERNAL"),
				Default:          "EXTERNAL",
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ip_version": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateFunc:     validation.StringInSlice([]string{"IPV4", "IPV6", ""}, false),
				DiffSuppressFunc: emptyOrDefaultStringSuppress("IPV4"),
			},
			"network": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"prefix_length": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"purpose": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"VPC_PEERING", ""}, false),
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

func resourceComputeGlobalAddressCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	addressProp, err := expandComputeGlobalAddressAddress(d.Get("address"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("address"); !isEmptyValue(reflect.ValueOf(addressProp)) && (ok || !reflect.DeepEqual(v, addressProp)) {
		obj["address"] = addressProp
	}
	descriptionProp, err := expandComputeGlobalAddressDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeGlobalAddressName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	ipVersionProp, err := expandComputeGlobalAddressIpVersion(d.Get("ip_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_version"); !isEmptyValue(reflect.ValueOf(ipVersionProp)) && (ok || !reflect.DeepEqual(v, ipVersionProp)) {
		obj["ipVersion"] = ipVersionProp
	}
	prefixLengthProp, err := expandComputeGlobalAddressPrefixLength(d.Get("prefix_length"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("prefix_length"); !isEmptyValue(reflect.ValueOf(prefixLengthProp)) && (ok || !reflect.DeepEqual(v, prefixLengthProp)) {
		obj["prefixLength"] = prefixLengthProp
	}
	addressTypeProp, err := expandComputeGlobalAddressAddressType(d.Get("address_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("address_type"); !isEmptyValue(reflect.ValueOf(addressTypeProp)) && (ok || !reflect.DeepEqual(v, addressTypeProp)) {
		obj["addressType"] = addressTypeProp
	}
	purposeProp, err := expandComputeGlobalAddressPurpose(d.Get("purpose"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("purpose"); !isEmptyValue(reflect.ValueOf(purposeProp)) && (ok || !reflect.DeepEqual(v, purposeProp)) {
		obj["purpose"] = purposeProp
	}
	networkProp, err := expandComputeGlobalAddressNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/addresses")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GlobalAddress: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating GlobalAddress: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	waitErr := computeOperationWaitTime(
		config, res, project, "Creating GlobalAddress",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create GlobalAddress: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating GlobalAddress %q: %#v", d.Id(), res)

	return resourceComputeGlobalAddressRead(d, meta)
}

func resourceComputeGlobalAddressRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/addresses/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeGlobalAddress %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}

	if err := d.Set("address", flattenComputeGlobalAddressAddress(res["address"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeGlobalAddressCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("description", flattenComputeGlobalAddressDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("name", flattenComputeGlobalAddressName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("ip_version", flattenComputeGlobalAddressIpVersion(res["ipVersion"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("prefix_length", flattenComputeGlobalAddressPrefixLength(res["prefixLength"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("address_type", flattenComputeGlobalAddressAddressType(res["addressType"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("purpose", flattenComputeGlobalAddressPurpose(res["purpose"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("network", flattenComputeGlobalAddressNetwork(res["network"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}

	return nil
}

func resourceComputeGlobalAddressDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/addresses/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting GlobalAddress %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "GlobalAddress")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting GlobalAddress",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting GlobalAddress %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeGlobalAddressImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/addresses/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeGlobalAddressAddress(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressIpVersion(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressPrefixLength(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeGlobalAddressAddressType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressPurpose(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeGlobalAddressNetwork(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeGlobalAddressAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressIpVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressPrefixLength(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressAddressType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressPurpose(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}
