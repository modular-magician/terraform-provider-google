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

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"google.golang.org/api/compute/v1"
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
			Create: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
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
	} else if v := d.Get("address"); !isEmptyValue(reflect.ValueOf(addressProp)) {
		obj["address"] = addressProp
	}
	descriptionProp, err := expandComputeGlobalAddressDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v := d.Get("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeGlobalAddressName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v := d.Get("name"); !isEmptyValue(reflect.ValueOf(nameProp)) {
		obj["name"] = nameProp
	}
	ipVersionProp, err := expandComputeGlobalAddressIpVersion(d.Get("ip_version"), d, config)
	if err != nil {
		return err
	} else if v := d.Get("ip_version"); !isEmptyValue(reflect.ValueOf(ipVersionProp)) {
		obj["ipVersion"] = ipVersionProp
	}
	addressTypeProp, err := expandComputeGlobalAddressAddressType(d.Get("address_type"), d, config)
	if err != nil {
		return err
	} else if v := d.Get("address_type"); !isEmptyValue(reflect.ValueOf(addressTypeProp)) {
		obj["addressType"] = addressTypeProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/addresses")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GlobalAddress: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating GlobalAddress: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating GlobalAddress",
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

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/addresses/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeGlobalAddress %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
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
	if err := d.Set("address_type", flattenComputeGlobalAddressAddressType(res["addressType"], d)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}

	return nil
}

func resourceComputeGlobalAddressDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/addresses/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting GlobalAddress %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "GlobalAddress")
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting GlobalAddress",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting GlobalAddress %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeGlobalAddressImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{"projects/(?P<project>[^/]+)/global/addresses/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config); err != nil {
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

func flattenComputeGlobalAddressAddressType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandComputeGlobalAddressAddress(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressIpVersion(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressAddressType(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}
