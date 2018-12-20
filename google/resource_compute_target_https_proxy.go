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
	compute "google.golang.org/api/compute/v1"
)

func resourceComputeTargetHttpsProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeTargetHttpsProxyCreate,
		Read:   resourceComputeTargetHttpsProxyRead,
		Update: resourceComputeTargetHttpsProxyUpdate,
		Delete: resourceComputeTargetHttpsProxyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeTargetHttpsProxyImport,
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
			"ssl_certificates": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					DiffSuppressFunc: compareSelfLinkOrResourceName,
				},
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
			"quic_override": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "ENABLE", "DISABLE", ""}, false),
			},
			"ssl_policy": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
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

func resourceComputeTargetHttpsProxyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeTargetHttpsProxyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeTargetHttpsProxyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	quicOverrideProp, err := expandComputeTargetHttpsProxyQuicOverride(d.Get("quic_override"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("quic_override"); !isEmptyValue(reflect.ValueOf(quicOverrideProp)) && (ok || !reflect.DeepEqual(v, quicOverrideProp)) {
		obj["quicOverride"] = quicOverrideProp
	}
	sslCertificatesProp, err := expandComputeTargetHttpsProxySslCertificates(d.Get("ssl_certificates"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ssl_certificates"); !isEmptyValue(reflect.ValueOf(sslCertificatesProp)) && (ok || !reflect.DeepEqual(v, sslCertificatesProp)) {
		obj["sslCertificates"] = sslCertificatesProp
	}
	sslPolicyProp, err := expandComputeTargetHttpsProxySslPolicy(d.Get("ssl_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ssl_policy"); !isEmptyValue(reflect.ValueOf(sslPolicyProp)) && (ok || !reflect.DeepEqual(v, sslPolicyProp)) {
		obj["sslPolicy"] = sslPolicyProp
	}
	urlMapProp, err := expandComputeTargetHttpsProxyUrlMap(d.Get("url_map"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("url_map"); !isEmptyValue(reflect.ValueOf(urlMapProp)) && (ok || !reflect.DeepEqual(v, urlMapProp)) {
		obj["urlMap"] = urlMapProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/targetHttpsProxies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TargetHttpsProxy: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating TargetHttpsProxy: %s", err)
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
		config.clientCompute, op, project, "Creating TargetHttpsProxy",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create TargetHttpsProxy: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating TargetHttpsProxy %q: %#v", d.Id(), res)

	return resourceComputeTargetHttpsProxyRead(d, meta)
}

func resourceComputeTargetHttpsProxyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/targetHttpsProxies/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeTargetHttpsProxy %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeTargetHttpsProxyCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("description", flattenComputeTargetHttpsProxyDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("proxy_id", flattenComputeTargetHttpsProxyProxyId(res["id"], d)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("name", flattenComputeTargetHttpsProxyName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("quic_override", flattenComputeTargetHttpsProxyQuicOverride(res["quicOverride"], d)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("ssl_certificates", flattenComputeTargetHttpsProxySslCertificates(res["sslCertificates"], d)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("ssl_policy", flattenComputeTargetHttpsProxySslPolicy(res["sslPolicy"], d)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("url_map", flattenComputeTargetHttpsProxyUrlMap(res["urlMap"], d)); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading TargetHttpsProxy: %s", err)
	}

	return nil
}

func resourceComputeTargetHttpsProxyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	d.Partial(true)

	if d.HasChange("quic_override") {
		obj := make(map[string]interface{})
		quicOverrideProp, err := expandComputeTargetHttpsProxyQuicOverride(d.Get("quic_override"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("quic_override"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, quicOverrideProp)) {
			obj["quicOverride"] = quicOverrideProp
		}

		url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/targetHttpsProxies/{{name}}/setQuicOverride")
		if err != nil {
			return err
		}
		res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating TargetHttpsProxy %q: %s", d.Id(), err)
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
			config.clientCompute, op, project, "Updating TargetHttpsProxy",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("quic_override")
	}
	if d.HasChange("ssl_certificates") {
		obj := make(map[string]interface{})
		sslCertificatesProp, err := expandComputeTargetHttpsProxySslCertificates(d.Get("ssl_certificates"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("ssl_certificates"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sslCertificatesProp)) {
			obj["sslCertificates"] = sslCertificatesProp
		}

		url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/targetHttpsProxies/{{name}}/setSslCertificates")
		if err != nil {
			return err
		}
		res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating TargetHttpsProxy %q: %s", d.Id(), err)
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
			config.clientCompute, op, project, "Updating TargetHttpsProxy",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("ssl_certificates")
	}
	if d.HasChange("ssl_policy") {
		obj := make(map[string]interface{})
		sslPolicyProp, err := expandComputeTargetHttpsProxySslPolicy(d.Get("ssl_policy"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("ssl_policy"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sslPolicyProp)) {
			obj["sslPolicy"] = sslPolicyProp
		}

		url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/targetHttpsProxies/{{name}}/setSslPolicy")
		if err != nil {
			return err
		}
		res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating TargetHttpsProxy %q: %s", d.Id(), err)
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
			config.clientCompute, op, project, "Updating TargetHttpsProxy",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("ssl_policy")
	}
	if d.HasChange("url_map") {
		obj := make(map[string]interface{})
		urlMapProp, err := expandComputeTargetHttpsProxyUrlMap(d.Get("url_map"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("url_map"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, urlMapProp)) {
			obj["urlMap"] = urlMapProp
		}

		url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/targetHttpsProxies/{{name}}/setUrlMap")
		if err != nil {
			return err
		}
		res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating TargetHttpsProxy %q: %s", d.Id(), err)
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
			config.clientCompute, op, project, "Updating TargetHttpsProxy",
			int(d.Timeout(schema.TimeoutUpdate).Minutes()))

		if err != nil {
			return err
		}

		d.SetPartial("url_map")
	}

	d.Partial(false)

	return resourceComputeTargetHttpsProxyRead(d, meta)
}

func resourceComputeTargetHttpsProxyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/targetHttpsProxies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting TargetHttpsProxy %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "TargetHttpsProxy")
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
		config.clientCompute, op, project, "Deleting TargetHttpsProxy",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting TargetHttpsProxy %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeTargetHttpsProxyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{"projects/(?P<project>[^/]+)/global/targetHttpsProxies/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config); err != nil {
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

func flattenComputeTargetHttpsProxyCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeTargetHttpsProxyDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeTargetHttpsProxyProxyId(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeTargetHttpsProxyName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeTargetHttpsProxyQuicOverride(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeTargetHttpsProxySslCertificates(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return convertAndMapStringArr(v.([]interface{}), ConvertSelfLinkToV1)
}

func flattenComputeTargetHttpsProxySslPolicy(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeTargetHttpsProxyUrlMap(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeTargetHttpsProxyDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetHttpsProxyName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetHttpsProxyQuicOverride(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetHttpsProxySslCertificates(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
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

func expandComputeTargetHttpsProxySslPolicy(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("sslPolicies", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for ssl_policy: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeTargetHttpsProxyUrlMap(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("urlMaps", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for url_map: %s", err)
	}
	return f.RelativeLink(), nil
}
