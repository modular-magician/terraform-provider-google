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
)

func resourceDnsManagedZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnsManagedZoneCreate,
		Read:   resourceDnsManagedZoneRead,
		Update: resourceDnsManagedZoneUpdate,
		Delete: resourceDnsManagedZoneDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDnsManagedZoneImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"dns_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Managed by Terraform",
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"name_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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

func resourceDnsManagedZoneCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandDnsManagedZoneDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	dnsNameProp, err := expandDnsManagedZoneDnsName(d.Get("dns_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dns_name"); !isEmptyValue(reflect.ValueOf(dnsNameProp)) && (ok || !reflect.DeepEqual(v, dnsNameProp)) {
		obj["dnsName"] = dnsNameProp
	}
	nameProp, err := expandDnsManagedZoneName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	labelsProp, err := expandDnsManagedZoneLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/dns/v1/projects/{{project}}/managedZones")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ManagedZone: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ManagedZone: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ManagedZone %q: %#v", d.Id(), res)

	return resourceDnsManagedZoneRead(d, meta)
}

func resourceDnsManagedZoneRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/dns/v1/projects/{{project}}/managedZones/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("DnsManagedZone %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}

	if err := d.Set("description", flattenDnsManagedZoneDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("dns_name", flattenDnsManagedZoneDnsName(res["dnsName"], d)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("name", flattenDnsManagedZoneName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("name_servers", flattenDnsManagedZoneNameServers(res["nameServers"], d)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}
	if err := d.Set("labels", flattenDnsManagedZoneLabels(res["labels"], d)); err != nil {
		return fmt.Errorf("Error reading ManagedZone: %s", err)
	}

	return nil
}

func resourceDnsManagedZoneUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	d.Partial(true)

	if d.HasChange("description") || d.HasChange("labels") {
		obj := make(map[string]interface{})
		descriptionProp, err := expandDnsManagedZoneDescription(d.Get("description"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
			obj["description"] = descriptionProp
		}
		labelsProp, err := expandDnsManagedZoneLabels(d.Get("labels"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
			obj["labels"] = labelsProp
		}

		url, err := replaceVars(d, config, "https://www.googleapis.com/dns/v1/projects/{{project}}/managedZones/{{name}}")
		if err != nil {
			return err
		}
		_, err = sendRequestWithTimeout(config, "PATCH", url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating ManagedZone %q: %s", d.Id(), err)
		}

		d.SetPartial("description")
		d.SetPartial("labels")
	}

	d.Partial(false)

	return resourceDnsManagedZoneRead(d, meta)
}

func resourceDnsManagedZoneDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/dns/v1/projects/{{project}}/managedZones/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ManagedZone %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ManagedZone")
	}

	log.Printf("[DEBUG] Finished deleting ManagedZone %q: %#v", d.Id(), res)
	return nil
}

func resourceDnsManagedZoneImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{"projects/(?P<project>[^/]+)/managedZones/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config); err != nil {
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

func flattenDnsManagedZoneDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDnsManagedZoneDnsName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDnsManagedZoneName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDnsManagedZoneNameServers(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenDnsManagedZoneLabels(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandDnsManagedZoneDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDnsManagedZoneDnsName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDnsManagedZoneName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDnsManagedZoneLabels(v interface{}, d *schema.ResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
