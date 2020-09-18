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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComputeTargetGrpcProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeTargetGrpcProxyCreate,
		Read:   resourceComputeTargetGrpcProxyRead,
		Update: resourceComputeTargetGrpcProxyUpdate,
		Delete: resourceComputeTargetGrpcProxyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeTargetGrpcProxyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. Provided by the client when the resource
is created. The name must be 1-63 characters long, and comply
with RFC1035. Specifically, the name must be 1-63 characters long
and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which
means the first character must be a lowercase letter, and all
following characters must be a dash, lowercase letter, or digit,
except the last character, which cannot be a dash.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of this resource.`,
			},
			"url_map": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `URL to the UrlMap resource that defines the mapping from URL to
the BackendService. The protocol field in the BackendService
must be set to GRPC.`,
			},
			"validate_for_proxyless": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `If true, indicates that the BackendServices referenced by
the urlMap may be accessed by gRPC applications without using
a sidecar proxy. This will enable configuration checks on urlMap
and its referenced BackendServices to not allow unsupported features.
A gRPC application must use "xds:///" scheme in the target URI
of the service it is connecting to. If false, indicates that the
BackendServices referenced by the urlMap will be accessed by gRPC
applications via a sidecar proxy. In this case, a gRPC application
must not use "xds:///" scheme in the target URI of the service
it is connecting to`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Fingerprint of this resource. A hash of the contents stored in
this object. This field is used in optimistic locking. This field
will be ignored when inserting a TargetGrpcProxy. An up-to-date
fingerprint must be provided in order to patch/update the
TargetGrpcProxy; otherwise, the request will fail with error
412 conditionNotMet. To see the latest fingerprint, make a get()
request to retrieve the TargetGrpcProxy. A base64-encoded string.`,
			},
			"self_link_with_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Server-defined URL with id for the resource.`,
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

func resourceComputeTargetGrpcProxyCreate(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	obj := make(map[string]interface{})
	nameProp, err := expandComputeTargetGrpcProxyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeTargetGrpcProxyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	urlMapProp, err := expandComputeTargetGrpcProxyUrlMap(d.Get("url_map"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("url_map"); !isEmptyValue(reflect.ValueOf(urlMapProp)) && (ok || !reflect.DeepEqual(v, urlMapProp)) {
		obj["urlMap"] = urlMapProp
	}
	validateForProxylessProp, err := expandComputeTargetGrpcProxyValidateForProxyless(d.Get("validate_for_proxyless"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("validate_for_proxyless"); !isEmptyValue(reflect.ValueOf(validateForProxylessProp)) && (ok || !reflect.DeepEqual(v, validateForProxylessProp)) {
		obj["validateForProxyless"] = validateForProxylessProp
	}
	fingerprintProp, err := expandComputeTargetGrpcProxyFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetGrpcProxies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TargetGrpcProxy: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating TargetGrpcProxy: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/global/targetGrpcProxies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating TargetGrpcProxy",
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create TargetGrpcProxy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating TargetGrpcProxy %q: %#v", d.Id(), res)

	return resourceComputeTargetGrpcProxyRead(d, meta)
}

func resourceComputeTargetGrpcProxyRead(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetGrpcProxies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeTargetGrpcProxy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading TargetGrpcProxy: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeTargetGrpcProxyCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetGrpcProxy: %s", err)
	}
	if err := d.Set("name", flattenComputeTargetGrpcProxyName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetGrpcProxy: %s", err)
	}
	if err := d.Set("description", flattenComputeTargetGrpcProxyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetGrpcProxy: %s", err)
	}
	if err := d.Set("self_link_with_id", flattenComputeTargetGrpcProxySelfLinkWithId(res["selfLinkWithId"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetGrpcProxy: %s", err)
	}
	if err := d.Set("url_map", flattenComputeTargetGrpcProxyUrlMap(res["urlMap"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetGrpcProxy: %s", err)
	}
	if err := d.Set("validate_for_proxyless", flattenComputeTargetGrpcProxyValidateForProxyless(res["validateForProxyless"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetGrpcProxy: %s", err)
	}
	if err := d.Set("fingerprint", flattenComputeTargetGrpcProxyFingerprint(res["fingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading TargetGrpcProxy: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading TargetGrpcProxy: %s", err)
	}

	return nil
}

func resourceComputeTargetGrpcProxyUpdate(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeTargetGrpcProxyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeTargetGrpcProxyFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetGrpcProxies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating TargetGrpcProxy %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating TargetGrpcProxy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating TargetGrpcProxy %q: %#v", d.Id(), res)
	}

	err = computeOperationWaitTime(
		config, res, project, "Updating TargetGrpcProxy",
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputeTargetGrpcProxyRead(d, meta)
}

func resourceComputeTargetGrpcProxyDelete(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleKey)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/targetGrpcProxies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting TargetGrpcProxy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "TargetGrpcProxy")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting TargetGrpcProxy",
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting TargetGrpcProxy %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeTargetGrpcProxyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/targetGrpcProxies/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/global/targetGrpcProxies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeTargetGrpcProxyCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeTargetGrpcProxyName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeTargetGrpcProxyDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeTargetGrpcProxySelfLinkWithId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeTargetGrpcProxyUrlMap(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeTargetGrpcProxyValidateForProxyless(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeTargetGrpcProxyFingerprint(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandComputeTargetGrpcProxyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetGrpcProxyDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetGrpcProxyUrlMap(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetGrpcProxyValidateForProxyless(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeTargetGrpcProxyFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
