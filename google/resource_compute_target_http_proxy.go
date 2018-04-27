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
	compute "google.golang.org/api/compute/v1"
)

func resourceComputeTargetHttpProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeTargetHttpProxyCreate,
		Read:   resourceComputeTargetHttpProxyRead,
		Update: resourceComputeTargetHttpProxyUpdate,
		Delete: resourceComputeTargetHttpProxyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeTargetHttpProxyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"url_map": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_id": {
				Type:     schema.TypeInt,
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

func resourceComputeTargetHttpProxyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	descriptionProp, err := expandComputeTargetHttpProxyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	}
	nameProp, err := expandComputeTargetHttpProxyName(d.Get("name"), d, config)
	if err != nil {
		return err
	}
	urlMapProp, err := expandComputeTargetHttpProxyUrlMap(d.Get("url_map"), d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"description": descriptionProp,
		"name":        nameProp,
		"urlMap":      urlMapProp,
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/targetHttpProxies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TargetHttpProxy: %#v", obj)
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating TargetHttpProxy: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWaitTime(
		config.clientCompute, op, project, "Creating TargetHttpProxy",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return waitErr
	}

	return resourceComputeTargetHttpProxyRead(d, meta)
}

func resourceComputeTargetHttpProxyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/targetHttpProxies/{{name}}")
	if err != nil {
		return err
	}

	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeTargetHttpProxy %q", d.Id()))
	}

	d.Set("creation_timestamp", flattenComputeTargetHttpProxyCreationTimestamp(res["creationTimestamp"]))
	d.Set("description", flattenComputeTargetHttpProxyDescription(res["description"]))
	d.Set("proxy_id", flattenComputeTargetHttpProxyProxyId(res["id"]))
	d.Set("name", flattenComputeTargetHttpProxyName(res["name"]))
	d.Set("url_map", flattenComputeTargetHttpProxyUrlMap(res["urlMap"]))
	d.Set("self_link", res["selfLink"])
	d.Set("project", project)

	return nil
}

func resourceComputeTargetHttpProxyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	var url string
	var res map[string]interface{}
	op := &compute.Operation{}

	d.Partial(true)

	if d.HasChange("url_map") {
		descriptionProp, err := expandComputeTargetHttpProxyDescription(d.Get("description"), d, config)
		if err != nil {
			return err
		}
		nameProp, err := expandComputeTargetHttpProxyName(d.Get("name"), d, config)
		if err != nil {
			return err
		}
		urlMapProp, err := expandComputeTargetHttpProxyUrlMap(d.Get("url_map"), d, config)
		if err != nil {
			return err
		}

		obj := map[string]interface{}{
			"description": descriptionProp,
			"name":        nameProp,
			"urlMap":      urlMapProp,
		}
		url, err = replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/targetHttpProxies/{{name}}/setUrlMap")
		if err != nil {
			return err
		}
		res, err = sendRequest(config, "POST", url, obj)
		if err != nil {
			return fmt.Errorf("Error updating TargetHttpProxy %q: %s", d.Id(), err)
		}

		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating TargetHttpProxy",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("url_map")
	}

	d.Partial(false)

	return resourceComputeTargetHttpProxyRead(d, meta)
}

func resourceComputeTargetHttpProxyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/targetHttpProxies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting TargetHttpProxy %q", d.Id())
	res, err := Delete(config, url)
	if err != nil {
		return fmt.Errorf("Error deleting TargetHttpProxy %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting TargetHttpProxy",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	return nil
}

func resourceComputeTargetHttpProxyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"projects/(?P<project>[^/]+)/global/targetHttpProxies/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeTargetHttpProxyCreationTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeTargetHttpProxyDescription(v interface{}) interface{} {
	return v
}

func flattenComputeTargetHttpProxyProxyId(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.Atoi(strVal); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeTargetHttpProxyName(v interface{}) interface{} {
	return v
}

func flattenComputeTargetHttpProxyUrlMap(v interface{}) interface{} {
	return v
}

func expandComputeTargetHttpProxyDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetHttpProxyName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetHttpProxyUrlMap(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("urlMaps", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for url_map: %s", err)
	}
	return f.RelativeLink(), nil
}
