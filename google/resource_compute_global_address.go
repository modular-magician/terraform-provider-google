// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceComputeGlobalAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeGlobalAddressCreate,
		Read:   resourceComputeGlobalAddressRead,
		Delete: resourceComputeGlobalAddressDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeGlobalAddressImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035.  Specifically, the name must be 1-63 characters long and
match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
the first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"address": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `The IP address or beginning of the address range represented by this
resource. This can be supplied as an input to reserve a specific
address or omitted to allow GCP to choose a valid one for you.`,
			},
			"address_type": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateFunc:     validateEnum([]string{"EXTERNAL", "INTERNAL", ""}),
				DiffSuppressFunc: EmptyOrDefaultStringSuppress("EXTERNAL"),
				Description: `The type of the address to reserve.

* EXTERNAL indicates public/external single IP address.
* INTERNAL indicates internal IP ranges belonging to some network. Default value: "EXTERNAL" Possible values: ["EXTERNAL", "INTERNAL"]`,
				Default: "EXTERNAL",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"ip_version": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateFunc:     validateEnum([]string{"IPV4", "IPV6", ""}),
				DiffSuppressFunc: EmptyOrDefaultStringSuppress("IPV4"),
				Description:      `The IP Version that will be used by this address. The default value is 'IPV4'. Possible values: ["IPV4", "IPV6"]`,
			},
			"network": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `The URL of the network in which to reserve the IP range. The IP range
must be in RFC1918 space. The network cannot be deleted if there are
any reserved IP ranges referring to it.

This should only be set when using an Internal address.`,
			},
			"prefix_length": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Description: `The prefix length of the IP range. If not present, it means the
address field is a single IP address.

This field is not applicable to addresses with addressType=EXTERNAL,
or addressType=INTERNAL when purpose=PRIVATE_SERVICE_CONNECT`,
			},
			"purpose": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The purpose of the resource. Possible values include:

* VPC_PEERING - for peer networks

* PRIVATE_SERVICE_CONNECT - for ([Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html) only) Private Service Connect networks`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
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
		UseJSONNumber: true,
	}
}

func resourceComputeGlobalAddressCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

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

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/addresses")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GlobalAddress: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GlobalAddress: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating GlobalAddress: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/global/addresses/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating GlobalAddress", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create GlobalAddress: %s", err)
	}

	log.Printf("[DEBUG] Finished creating GlobalAddress %q: %#v", d.Id(), res)

	return resourceComputeGlobalAddressRead(d, meta)
}

func resourceComputeGlobalAddressRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/addresses/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GlobalAddress: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeGlobalAddress %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}

	if err := d.Set("address", flattenComputeGlobalAddressAddress(res["address"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeGlobalAddressCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("description", flattenComputeGlobalAddressDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("name", flattenComputeGlobalAddressName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("ip_version", flattenComputeGlobalAddressIpVersion(res["ipVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("prefix_length", flattenComputeGlobalAddressPrefixLength(res["prefixLength"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("address_type", flattenComputeGlobalAddressAddressType(res["addressType"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("purpose", flattenComputeGlobalAddressPurpose(res["purpose"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("network", flattenComputeGlobalAddressNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading GlobalAddress: %s", err)
	}

	return nil
}

func resourceComputeGlobalAddressDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GlobalAddress: %s", err)
	}
	billingProject = project

	url, err := ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/addresses/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting GlobalAddress %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "GlobalAddress")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting GlobalAddress", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting GlobalAddress %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeGlobalAddressImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/global/addresses/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/global/addresses/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeGlobalAddressAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeGlobalAddressCreationTimestamp(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeGlobalAddressDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeGlobalAddressName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeGlobalAddressIpVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeGlobalAddressPrefixLength(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeGlobalAddressAddressType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeGlobalAddressPurpose(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeGlobalAddressNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeGlobalAddressAddress(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressIpVersion(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressPrefixLength(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressAddressType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressPurpose(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressNetwork(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}
