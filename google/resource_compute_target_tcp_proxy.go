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

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"google.golang.org/api/compute/v1"
)

func resourceComputeTargetTcpProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeTargetTcpProxyCreate,
		Read:   resourceComputeTargetTcpProxyRead,
		Update: resourceComputeTargetTcpProxyUpdate,
		Delete: resourceComputeTargetTcpProxyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeTargetTcpProxyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"backend_service": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"proxy_header": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "PROXY_V1", ""}, false),
				Default:      "NONE",
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

func resourceComputeTargetTcpProxyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeTargetTcpProxyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeTargetTcpProxyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	proxyHeaderProp, err := expandComputeTargetTcpProxyProxyHeader(d.Get("proxy_header"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("proxy_header"); !isEmptyValue(reflect.ValueOf(proxyHeaderProp)) && (ok || !reflect.DeepEqual(v, proxyHeaderProp)) {
		obj["proxyHeader"] = proxyHeaderProp
	}
	serviceProp, err := expandComputeTargetTcpProxyBackendService(d.Get("backend_service"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backend_service"); !isEmptyValue(reflect.ValueOf(serviceProp)) && (ok || !reflect.DeepEqual(v, serviceProp)) {
		obj["service"] = serviceProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetTcpProxies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TargetTcpProxy: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating TargetTcpProxy: %s", err)
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
		config.clientCompute, op, project, "Creating TargetTcpProxy",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create TargetTcpProxy: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating TargetTcpProxy %q: %#v", d.Id(), res)

	return resourceComputeTargetTcpProxyRead(d, meta)
}

func resourceComputeTargetTcpProxyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetTcpProxies/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeTargetTcpProxy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading TargetTcpProxy: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeTargetTcpProxyCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading TargetTcpProxy: %s", err)
	}
	if err := d.Set("description", flattenComputeTargetTcpProxyDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading TargetTcpProxy: %s", err)
	}
	if err := d.Set("proxy_id", flattenComputeTargetTcpProxyProxyId(res["id"], d)); err != nil {
		return fmt.Errorf("Error reading TargetTcpProxy: %s", err)
	}
	if err := d.Set("name", flattenComputeTargetTcpProxyName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading TargetTcpProxy: %s", err)
	}
	if err := d.Set("proxy_header", flattenComputeTargetTcpProxyProxyHeader(res["proxyHeader"], d)); err != nil {
		return fmt.Errorf("Error reading TargetTcpProxy: %s", err)
	}
	if err := d.Set("backend_service", flattenComputeTargetTcpProxyBackendService(res["service"], d)); err != nil {
		return fmt.Errorf("Error reading TargetTcpProxy: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading TargetTcpProxy: %s", err)
	}

	return nil
}

func resourceComputeTargetTcpProxyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	d.Partial(true)

	if d.HasChange("proxy_header") {
		obj := make(map[string]interface{})
		proxyHeaderProp, err := expandComputeTargetTcpProxyProxyHeader(d.Get("proxy_header"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("proxy_header"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, proxyHeaderProp)) {
			obj["proxyHeader"] = proxyHeaderProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetTcpProxies/{{name}}/setProxyHeader")
		if err != nil {
			return err
		}
		res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating TargetTcpProxy %q: %s", d.Id(), err)
		}

		op := &compute.Operation{}
		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating TargetTcpProxy",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("proxy_header")
	}
	if d.HasChange("backend_service") {
		obj := make(map[string]interface{})
		serviceProp, err := expandComputeTargetTcpProxyBackendService(d.Get("backend_service"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("backend_service"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, serviceProp)) {
			obj["service"] = serviceProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetTcpProxies/{{name}}/setBackendService")
		if err != nil {
			return err
		}
		res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating TargetTcpProxy %q: %s", d.Id(), err)
		}

		op := &compute.Operation{}
		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating TargetTcpProxy",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("backend_service")
	}

	d.Partial(false)

	return resourceComputeTargetTcpProxyRead(d, meta)
}

func resourceComputeTargetTcpProxyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetTcpProxies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting TargetTcpProxy %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "TargetTcpProxy")
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting TargetTcpProxy",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting TargetTcpProxy %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeTargetTcpProxyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/targetTcpProxies/(?P<name>[^/]+)",
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

func flattenComputeTargetTcpProxyCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeTargetTcpProxyDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeTargetTcpProxyProxyId(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeTargetTcpProxyName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeTargetTcpProxyProxyHeader(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeTargetTcpProxyBackendService(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeTargetTcpProxyDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetTcpProxyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetTcpProxyProxyHeader(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetTcpProxyBackendService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("backendServices", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for backend_service: %s", err)
	}
	return f.RelativeLink(), nil
}
