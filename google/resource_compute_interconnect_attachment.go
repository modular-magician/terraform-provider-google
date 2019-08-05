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
)

func resourceComputeInterconnectAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeInterconnectAttachmentCreate,
		Read:   resourceComputeInterconnectAttachmentRead,
		Delete: resourceComputeInterconnectAttachmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeInterconnectAttachmentImport,
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
				ValidateFunc: validateRegexp(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
			},
			"router": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"candidate_subnets": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"edge_availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"interconnect": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"type": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"DEDICATED", "PARTNER", "PARTNER_PROVIDER", ""}, false),
			},
			"vlan_tag8021q": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"cloud_router_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"customer_router_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"google_reference_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pairing_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"partner_asn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_interconnect_info": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tag8021q": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"state": {
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

func resourceComputeInterconnectAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	interconnectProp, err := expandComputeInterconnectAttachmentInterconnect(d.Get("interconnect"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("interconnect"); !isEmptyValue(reflect.ValueOf(interconnectProp)) && (ok || !reflect.DeepEqual(v, interconnectProp)) {
		obj["interconnect"] = interconnectProp
	}
	descriptionProp, err := expandComputeInterconnectAttachmentDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	edgeAvailabilityDomainProp, err := expandComputeInterconnectAttachmentEdgeAvailabilityDomain(d.Get("edge_availability_domain"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("edge_availability_domain"); !isEmptyValue(reflect.ValueOf(edgeAvailabilityDomainProp)) && (ok || !reflect.DeepEqual(v, edgeAvailabilityDomainProp)) {
		obj["edgeAvailabilityDomain"] = edgeAvailabilityDomainProp
	}
	typeProp, err := expandComputeInterconnectAttachmentType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	routerProp, err := expandComputeInterconnectAttachmentRouter(d.Get("router"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("router"); !isEmptyValue(reflect.ValueOf(routerProp)) && (ok || !reflect.DeepEqual(v, routerProp)) {
		obj["router"] = routerProp
	}
	nameProp, err := expandComputeInterconnectAttachmentName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	candidateSubnetsProp, err := expandComputeInterconnectAttachmentCandidateSubnets(d.Get("candidate_subnets"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("candidate_subnets"); !isEmptyValue(reflect.ValueOf(candidateSubnetsProp)) && (ok || !reflect.DeepEqual(v, candidateSubnetsProp)) {
		obj["candidateSubnets"] = candidateSubnetsProp
	}
	vlanTag8021qProp, err := expandComputeInterconnectAttachmentVlanTag8021q(d.Get("vlan_tag8021q"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vlan_tag8021q"); !isEmptyValue(reflect.ValueOf(vlanTag8021qProp)) && (ok || !reflect.DeepEqual(v, vlanTag8021qProp)) {
		obj["vlanTag8021q"] = vlanTag8021qProp
	}
	regionProp, err := expandComputeInterconnectAttachmentRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/interconnectAttachments")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new InterconnectAttachment: %#v", obj)
	res, err := sendRequestWithTimeout(config, "POST", url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating InterconnectAttachment: %s", err)
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
		config.clientCompute, op, project, "Creating InterconnectAttachment",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create InterconnectAttachment: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating InterconnectAttachment %q: %#v", d.Id(), res)

	return resourceComputeInterconnectAttachmentRead(d, meta)
}

func resourceComputeInterconnectAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeInterconnectAttachment %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}

	if err := d.Set("cloud_router_ip_address", flattenComputeInterconnectAttachmentCloudRouterIpAddress(res["cloudRouterIpAddress"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("customer_router_ip_address", flattenComputeInterconnectAttachmentCustomerRouterIpAddress(res["customerRouterIpAddress"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("interconnect", flattenComputeInterconnectAttachmentInterconnect(res["interconnect"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("description", flattenComputeInterconnectAttachmentDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("edge_availability_domain", flattenComputeInterconnectAttachmentEdgeAvailabilityDomain(res["edgeAvailabilityDomain"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("pairing_key", flattenComputeInterconnectAttachmentPairingKey(res["pairingKey"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("partner_asn", flattenComputeInterconnectAttachmentPartnerAsn(res["partnerAsn"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("private_interconnect_info", flattenComputeInterconnectAttachmentPrivateInterconnectInfo(res["privateInterconnectInfo"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("type", flattenComputeInterconnectAttachmentType(res["type"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("state", flattenComputeInterconnectAttachmentState(res["state"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("google_reference_id", flattenComputeInterconnectAttachmentGoogleReferenceId(res["googleReferenceId"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("router", flattenComputeInterconnectAttachmentRouter(res["router"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeInterconnectAttachmentCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("name", flattenComputeInterconnectAttachmentName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("candidate_subnets", flattenComputeInterconnectAttachmentCandidateSubnets(res["candidateSubnets"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("vlan_tag8021q", flattenComputeInterconnectAttachmentVlanTag8021q(res["vlanTag8021q"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("region", flattenComputeInterconnectAttachmentRegion(res["region"], d)); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading InterconnectAttachment: %s", err)
	}

	return nil
}

func resourceComputeInterconnectAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting InterconnectAttachment %q", d.Id())
	res, err := sendRequestWithTimeout(config, "DELETE", url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "InterconnectAttachment")
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
		config.clientCompute, op, project, "Deleting InterconnectAttachment",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting InterconnectAttachment %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeInterconnectAttachmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/interconnectAttachments/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
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

func flattenComputeInterconnectAttachmentCloudRouterIpAddress(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentCustomerRouterIpAddress(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentInterconnect(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentEdgeAvailabilityDomain(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentPairingKey(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentPartnerAsn(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentPrivateInterconnectInfo(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["tag8021q"] =
		flattenComputeInterconnectAttachmentPrivateInterconnectInfoTag8021q(original["tag8021q"], d)
	return []interface{}{transformed}
}
func flattenComputeInterconnectAttachmentPrivateInterconnectInfoTag8021q(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeInterconnectAttachmentType(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentState(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentGoogleReferenceId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentRouter(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeInterconnectAttachmentCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentCandidateSubnets(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeInterconnectAttachmentVlanTag8021q(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeInterconnectAttachmentRegion(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeInterconnectAttachmentInterconnect(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentEdgeAvailabilityDomain(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentRouter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("routers", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for router: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeInterconnectAttachmentName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentCandidateSubnets(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentVlanTag8021q(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeInterconnectAttachmentRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
