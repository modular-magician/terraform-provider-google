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
	"bytes"
	"fmt"
	"log"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"google.golang.org/api/compute/v1"
)

// validatePeerAddr returns false if a tunnel's peer_ip property
// is invalid. Currently, only addresses that collide with RFC
// 5735 (https://tools.ietf.org/html/rfc5735) fail validation.
func validatePeerAddr(i interface{}, val string) ([]string, []error) {
	ip := net.ParseIP(i.(string))
	if ip == nil {
		return nil, []error{fmt.Errorf("could not parse %q to IP address", val)}
	}
	for _, test := range invalidPeerAddrs {
		if bytes.Compare(ip, test.from) >= 0 && bytes.Compare(ip, test.to) <= 0 {
			return nil, []error{fmt.Errorf("address is invalid (is between %q and %q, conflicting with RFC5735)", test.from, test.to)}
		}
	}
	return nil, nil
}

// invalidPeerAddrs is a collection of IP address ranges that represent
// a conflict with RFC 5735 (https://tools.ietf.org/html/rfc5735#page-3).
// CIDR range notations in the RFC were converted to a (from, to) pair
// for easy checking with bytes.Compare.
var invalidPeerAddrs = []struct {
	from net.IP
	to   net.IP
}{
	{
		from: net.ParseIP("0.0.0.0"),
		to:   net.ParseIP("0.255.255.255"),
	},
	{
		from: net.ParseIP("10.0.0.0"),
		to:   net.ParseIP("10.255.255.255"),
	},
	{
		from: net.ParseIP("127.0.0.0"),
		to:   net.ParseIP("127.255.255.255"),
	},
	{
		from: net.ParseIP("169.254.0.0"),
		to:   net.ParseIP("169.254.255.255"),
	},
	{
		from: net.ParseIP("172.16.0.0"),
		to:   net.ParseIP("172.31.255.255"),
	},
	{
		from: net.ParseIP("192.0.0.0"),
		to:   net.ParseIP("192.0.0.255"),
	},
	{
		from: net.ParseIP("192.0.2.0"),
		to:   net.ParseIP("192.0.2.255"),
	},
	{
		from: net.ParseIP("192.88.99.0"),
		to:   net.ParseIP("192.88.99.255"),
	},
	{
		from: net.ParseIP("192.168.0.0"),
		to:   net.ParseIP("192.168.255.255"),
	},
	{
		from: net.ParseIP("198.18.0.0"),
		to:   net.ParseIP("198.19.255.255"),
	},
	{
		from: net.ParseIP("198.51.100.0"),
		to:   net.ParseIP("198.51.100.255"),
	},
	{
		from: net.ParseIP("203.0.113.0"),
		to:   net.ParseIP("203.0.113.255"),
	},
	{
		from: net.ParseIP("224.0.0.0"),
		to:   net.ParseIP("239.255.255.255"),
	},
	{
		from: net.ParseIP("240.0.0.0"),
		to:   net.ParseIP("255.255.255.255"),
	},
	{
		from: net.ParseIP("255.255.255.255"),
		to:   net.ParseIP("255.255.255.255"),
	},
}

func getVpnTunnelLink(config *Config, project string, region string, tunnel string) (string, error) {
	if !strings.Contains(tunnel, "/") {
		// Tunnel value provided is just the name, lookup the tunnel SelfLink
		tunnelData, err := config.clientCompute.VpnTunnels.Get(
			project, region, tunnel).Do()
		if err != nil {
			return "", fmt.Errorf("Error reading tunnel: %s", err)
		}
		tunnel = tunnelData.SelfLink
	}

	return tunnel, nil

}

func resourceComputeVpnTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeVpnTunnelCreate,
		Read:   resourceComputeVpnTunnelRead,
		Delete: resourceComputeVpnTunnelDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeVpnTunnelImport,
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
			"shared_secret": {
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ike_version": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  2,
			},
			"local_traffic_selector": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"peer_ip": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validatePeerAddr,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"remote_traffic_selector": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"router": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"target_vpn_gateway": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"detailed_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shared_secret_hash": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel_id": {
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

func resourceComputeVpnTunnelCreate(d *schema.ResourceData, meta interface{}) error {

	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandComputeVpnTunnelName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeVpnTunnelDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	targetVpnGatewayProp, err := expandComputeVpnTunnelTargetVpnGateway(d.Get("target_vpn_gateway"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target_vpn_gateway"); !isEmptyValue(reflect.ValueOf(targetVpnGatewayProp)) && (ok || !reflect.DeepEqual(v, targetVpnGatewayProp)) {
		obj["targetVpnGateway"] = targetVpnGatewayProp
	}
	routerProp, err := expandComputeVpnTunnelRouter(d.Get("router"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("router"); !isEmptyValue(reflect.ValueOf(routerProp)) && (ok || !reflect.DeepEqual(v, routerProp)) {
		obj["router"] = routerProp
	}
	peerIpProp, err := expandComputeVpnTunnelPeerIp(d.Get("peer_ip"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("peer_ip"); !isEmptyValue(reflect.ValueOf(peerIpProp)) && (ok || !reflect.DeepEqual(v, peerIpProp)) {
		obj["peerIp"] = peerIpProp
	}
	sharedSecretProp, err := expandComputeVpnTunnelSharedSecret(d.Get("shared_secret"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("shared_secret"); !isEmptyValue(reflect.ValueOf(sharedSecretProp)) && (ok || !reflect.DeepEqual(v, sharedSecretProp)) {
		obj["sharedSecret"] = sharedSecretProp
	}
	ikeVersionProp, err := expandComputeVpnTunnelIkeVersion(d.Get("ike_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ike_version"); !isEmptyValue(reflect.ValueOf(ikeVersionProp)) && (ok || !reflect.DeepEqual(v, ikeVersionProp)) {
		obj["ikeVersion"] = ikeVersionProp
	}
	localTrafficSelectorProp, err := expandComputeVpnTunnelLocalTrafficSelector(d.Get("local_traffic_selector"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("local_traffic_selector"); !isEmptyValue(reflect.ValueOf(localTrafficSelectorProp)) && (ok || !reflect.DeepEqual(v, localTrafficSelectorProp)) {
		obj["localTrafficSelector"] = localTrafficSelectorProp
	}
	remoteTrafficSelectorProp, err := expandComputeVpnTunnelRemoteTrafficSelector(d.Get("remote_traffic_selector"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("remote_traffic_selector"); !isEmptyValue(reflect.ValueOf(remoteTrafficSelectorProp)) && (ok || !reflect.DeepEqual(v, remoteTrafficSelectorProp)) {
		obj["remoteTrafficSelector"] = remoteTrafficSelectorProp
	}
	regionProp, err := expandComputeVpnTunnelRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	obj, err = resourceComputeVpnTunnelEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/vpnTunnels")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new VpnTunnel: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating VpnTunnel: %s", err)
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
		config.clientCompute, op, project, "Creating VpnTunnel",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create VpnTunnel: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating VpnTunnel %q: %#v", d.Id(), res)

	return resourceComputeVpnTunnelRead(d, meta)
}

func resourceComputeVpnTunnelRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeVpnTunnel %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}

	if err := d.Set("tunnel_id", flattenComputeVpnTunnelTunnelId(res["id"], d)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("creation_timestamp", flattenComputeVpnTunnelCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("name", flattenComputeVpnTunnelName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("description", flattenComputeVpnTunnelDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("target_vpn_gateway", flattenComputeVpnTunnelTargetVpnGateway(res["targetVpnGateway"], d)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("router", flattenComputeVpnTunnelRouter(res["router"], d)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("peer_ip", flattenComputeVpnTunnelPeerIp(res["peerIp"], d)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("shared_secret_hash", flattenComputeVpnTunnelSharedSecretHash(res["sharedSecretHash"], d)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("ike_version", flattenComputeVpnTunnelIkeVersion(res["ikeVersion"], d)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("local_traffic_selector", flattenComputeVpnTunnelLocalTrafficSelector(res["localTrafficSelector"], d)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("remote_traffic_selector", flattenComputeVpnTunnelRemoteTrafficSelector(res["remoteTrafficSelector"], d)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("detailed_status", flattenComputeVpnTunnelDetailedStatus(res["detailedStatus"], d)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("region", flattenComputeVpnTunnelRegion(res["region"], d)); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading VpnTunnel: %s", err)
	}

	return nil
}

func resourceComputeVpnTunnelDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting VpnTunnel %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "VpnTunnel")
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting VpnTunnel",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting VpnTunnel %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeVpnTunnelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/vpnTunnels/(?P<name>[^/]+)",
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

func flattenComputeVpnTunnelTunnelId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeVpnTunnelCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeVpnTunnelName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeVpnTunnelDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeVpnTunnelTargetVpnGateway(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeVpnTunnelRouter(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeVpnTunnelPeerIp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeVpnTunnelSharedSecretHash(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeVpnTunnelIkeVersion(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeVpnTunnelLocalTrafficSelector(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeVpnTunnelRemoteTrafficSelector(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeVpnTunnelDetailedStatus(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeVpnTunnelRegion(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandComputeVpnTunnelName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelTargetVpnGateway(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("targetVpnGateways", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for target_vpn_gateway: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeVpnTunnelRouter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == nil || v.(string) == "" {
		return "", nil
	}
	f, err := parseRegionalFieldValue("routers", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for router: %s", err)
	}
	return "https://www.googleapis.com/compute/v1/" + f.RelativeLink(), nil
}

func expandComputeVpnTunnelPeerIp(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelSharedSecret(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelIkeVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelLocalTrafficSelector(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeVpnTunnelRemoteTrafficSelector(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeVpnTunnelRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}

func resourceComputeVpnTunnelEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)
	f, err := parseRegionalFieldValue("targetVpnGateways", d.Get("target_vpn_gateway").(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, err
	}
	if _, ok := d.GetOk("project"); !ok {
		d.Set("project", f.Project)
	}
	if _, ok := d.GetOk("region"); !ok {
		d.Set("region", f.Region)
	}
	return obj, nil
}
