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
	compute "google.golang.org/api/compute/v1"
)

func resourceComputeTargetSslProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeTargetSslProxyCreate,
		Read:   resourceComputeTargetSslProxyRead,
		Update: resourceComputeTargetSslProxyUpdate,
		Delete: resourceComputeTargetSslProxyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeTargetSslProxyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Update: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
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
			"ssl_certificates": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: compareSelfLinkOrResourceName,
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"proxy_header": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "PROXY_V1"}, false),
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
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeTargetSslProxyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	descriptionProp, err := expandComputeTargetSslProxyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	}
	nameProp, err := expandComputeTargetSslProxyName(d.Get("name"), d, config)
	if err != nil {
		return err
	}
	proxyHeaderProp, err := expandComputeTargetSslProxyProxyHeader(d.Get("proxy_header"), d, config)
	if err != nil {
		return err
	}
	serviceProp, err := expandComputeTargetSslProxyBackendService(d.Get("backend_service"), d, config)
	if err != nil {
		return err
	}
	sslCertificatesProp, err := expandComputeTargetSslProxySslCertificates(d.Get("ssl_certificates"), d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"description":     descriptionProp,
		"name":            nameProp,
		"proxyHeader":     proxyHeaderProp,
		"service":         serviceProp,
		"sslCertificates": sslCertificatesProp,
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/targetSslProxies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TargetSslProxy: %#v", obj)
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating TargetSslProxy: %s", err)
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
		config.clientCompute, op, project, "Creating TargetSslProxy",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return waitErr
	}

	return resourceComputeTargetSslProxyRead(d, meta)
}

func resourceComputeTargetSslProxyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/targetSslProxies/{{name}}")
	if err != nil {
		return err
	}

	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeTargetSslProxy %q", d.Id()))
	}

	d.Set("creation_timestamp", flattenComputeTargetSslProxyCreationTimestamp(res["creationTimestamp"]))
	d.Set("description", flattenComputeTargetSslProxyDescription(res["description"]))
	d.Set("proxy_id", flattenComputeTargetSslProxyProxyId(res["id"]))
	d.Set("name", flattenComputeTargetSslProxyName(res["name"]))
	d.Set("proxy_header", flattenComputeTargetSslProxyProxyHeader(res["proxyHeader"]))
	d.Set("backend_service", flattenComputeTargetSslProxyBackendService(res["service"]))
	d.Set("ssl_certificates", flattenComputeTargetSslProxySslCertificates(res["sslCertificates"]))
	d.Set("self_link", res["selfLink"])
	d.Set("project", project)

	return nil
}

func resourceComputeTargetSslProxyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	var url string
	var res map[string]interface{}
	op := &compute.Operation{}

	if d.HasChange("proxy_header") {
		descriptionProp, err := expandComputeTargetSslProxyDescription(d.Get("description"), d, config)
		if err != nil {
			return err
		}
		nameProp, err := expandComputeTargetSslProxyName(d.Get("name"), d, config)
		if err != nil {
			return err
		}
		proxyHeaderProp, err := expandComputeTargetSslProxyProxyHeader(d.Get("proxy_header"), d, config)
		if err != nil {
			return err
		}
		serviceProp, err := expandComputeTargetSslProxyBackendService(d.Get("backend_service"), d, config)
		if err != nil {
			return err
		}
		sslCertificatesProp, err := expandComputeTargetSslProxySslCertificates(d.Get("ssl_certificates"), d, config)
		if err != nil {
			return err
		}

		obj := map[string]interface{}{
			"description":     descriptionProp,
			"name":            nameProp,
			"proxyHeader":     proxyHeaderProp,
			"service":         serviceProp,
			"sslCertificates": sslCertificatesProp,
		}
		url, err = replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/targetSslProxies/{{name}}/setProxyHeader")
		if err != nil {
			return err
		}
		res, err = sendRequest(config, "POST", url, obj)
		if err != nil {
			return fmt.Errorf("Error updating TargetSslProxy %q: %s", d.Id(), err)
		}

		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating TargetSslProxy",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}
	}
	if d.HasChange("backend_service") {
		descriptionProp, err := expandComputeTargetSslProxyDescription(d.Get("description"), d, config)
		if err != nil {
			return err
		}
		nameProp, err := expandComputeTargetSslProxyName(d.Get("name"), d, config)
		if err != nil {
			return err
		}
		proxyHeaderProp, err := expandComputeTargetSslProxyProxyHeader(d.Get("proxy_header"), d, config)
		if err != nil {
			return err
		}
		serviceProp, err := expandComputeTargetSslProxyBackendService(d.Get("backend_service"), d, config)
		if err != nil {
			return err
		}
		sslCertificatesProp, err := expandComputeTargetSslProxySslCertificates(d.Get("ssl_certificates"), d, config)
		if err != nil {
			return err
		}

		obj := map[string]interface{}{
			"description":     descriptionProp,
			"name":            nameProp,
			"proxyHeader":     proxyHeaderProp,
			"service":         serviceProp,
			"sslCertificates": sslCertificatesProp,
		}
		url, err = replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/targetSslProxies/{{name}}/setBackendService")
		if err != nil {
			return err
		}
		res, err = sendRequest(config, "POST", url, obj)
		if err != nil {
			return fmt.Errorf("Error updating TargetSslProxy %q: %s", d.Id(), err)
		}

		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating TargetSslProxy",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}
	}
	if d.HasChange("ssl_certificates") {
		descriptionProp, err := expandComputeTargetSslProxyDescription(d.Get("description"), d, config)
		if err != nil {
			return err
		}
		nameProp, err := expandComputeTargetSslProxyName(d.Get("name"), d, config)
		if err != nil {
			return err
		}
		proxyHeaderProp, err := expandComputeTargetSslProxyProxyHeader(d.Get("proxy_header"), d, config)
		if err != nil {
			return err
		}
		serviceProp, err := expandComputeTargetSslProxyBackendService(d.Get("backend_service"), d, config)
		if err != nil {
			return err
		}
		sslCertificatesProp, err := expandComputeTargetSslProxySslCertificates(d.Get("ssl_certificates"), d, config)
		if err != nil {
			return err
		}

		obj := map[string]interface{}{
			"description":     descriptionProp,
			"name":            nameProp,
			"proxyHeader":     proxyHeaderProp,
			"service":         serviceProp,
			"sslCertificates": sslCertificatesProp,
		}
		url, err = replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/targetSslProxies/{{name}}/setSslCertificates")
		if err != nil {
			return err
		}
		res, err = sendRequest(config, "POST", url, obj)
		if err != nil {
			return fmt.Errorf("Error updating TargetSslProxy %q: %s", d.Id(), err)
		}

		err = Convert(res, op)
		if err != nil {
			return err
		}

		err = computeOperationWaitTime(
			config.clientCompute, op, project, "Updating TargetSslProxy",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}
	}

	return resourceComputeTargetSslProxyRead(d, meta)
}

func resourceComputeTargetSslProxyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/targetSslProxies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting TargetSslProxy %q", d.Id())
	res, err := Delete(config, url)
	if err != nil {
		return fmt.Errorf("Error deleting TargetSslProxy %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting TargetSslProxy",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	return nil
}

func resourceComputeTargetSslProxyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"projects/(?P<project>[^/]+)/global/targetSslProxies/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeTargetSslProxyCreationTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeTargetSslProxyDescription(v interface{}) interface{} {
	return v
}

func flattenComputeTargetSslProxyProxyId(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.Atoi(strVal); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeTargetSslProxyName(v interface{}) interface{} {
	return v
}

func flattenComputeTargetSslProxyProxyHeader(v interface{}) interface{} {
	return v
}

func flattenComputeTargetSslProxyBackendService(v interface{}) interface{} {
	return v
}

func flattenComputeTargetSslProxySslCertificates(v interface{}) interface{} {
	return v
}

func expandComputeTargetSslProxyDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetSslProxyName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetSslProxyProxyHeader(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetSslProxyBackendService(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("backendServices", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for backend_service: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeTargetSslProxySslCertificates(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		f, err := parseGlobalFieldValue("sslCertificates", raw.(string), "project", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for ssl_certificates: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}
