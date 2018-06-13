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

func resourceComputeRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRouteCreate,
		Read:   resourceComputeRouteRead,
		Delete: resourceComputeRouteDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRouteImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(240 * time.Second),
			Delete: schema.DefaultTimeout(240 * time.Second),
		},

		Schema: map[string]*schema.Schema{
			"dest_range": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  1000,
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"next_hop_gateway": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"next_hop_instance": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"next_hop_ip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"next_hop_vpn_tunnel": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"next_hop_instance_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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

func resourceComputeRouteCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	destRangeProp, err := expandComputeRouteDestRange(d.Get("dest_range"), d, config)
	if err != nil {
		return err
	}
	descriptionProp, err := expandComputeRouteDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	}
	nameProp, err := expandComputeRouteName(d.Get("name"), d, config)
	if err != nil {
		return err
	}
	networkProp, err := expandComputeRouteNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	}
	priorityProp, err := expandComputeRoutePriority(d.Get("priority"), d, config)
	if err != nil {
		return err
	}
	tagsProp, err := expandComputeRouteTags(d.Get("tags"), d, config)
	if err != nil {
		return err
	}
	nextHopGatewayProp, err := expandComputeRouteNextHopGateway(d.Get("next_hop_gateway"), d, config)
	if err != nil {
		return err
	}
	nextHopInstanceProp, err := expandComputeRouteNextHopInstance(d.Get("next_hop_instance"), d, config)
	if err != nil {
		return err
	}
	nextHopIpProp, err := expandComputeRouteNextHopIp(d.Get("next_hop_ip"), d, config)
	if err != nil {
		return err
	}
	nextHopVpnTunnelProp, err := expandComputeRouteNextHopVpnTunnel(d.Get("next_hop_vpn_tunnel"), d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"destRange":        destRangeProp,
		"description":      descriptionProp,
		"name":             nameProp,
		"network":          networkProp,
		"priority":         priorityProp,
		"tags":             tagsProp,
		"nextHopGateway":   nextHopGatewayProp,
		"nextHopInstance":  nextHopInstanceProp,
		"nextHopIp":        nextHopIpProp,
		"nextHopVpnTunnel": nextHopVpnTunnelProp,
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/routes")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Route: %#v", obj)
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating Route: %s", err)
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
		config.clientCompute, op, project, "Creating Route",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Route: %s", waitErr)
	}

	log.Printf("[DEBUG] Finished creating Route %q: %#v", d.Id(), res)

	return resourceComputeRouteRead(d, meta)
}

func resourceComputeRouteRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/routes/{{name}}")
	if err != nil {
		return err
	}

	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRoute %q", d.Id()))
	}

	res, err = resourceComputeRouteDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if err := d.Set("dest_range", flattenComputeRouteDestRange(res["destRange"])); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("description", flattenComputeRouteDescription(res["description"])); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("name", flattenComputeRouteName(res["name"])); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("network", flattenComputeRouteNetwork(res["network"])); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("priority", flattenComputeRoutePriority(res["priority"])); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("tags", flattenComputeRouteTags(res["tags"])); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("next_hop_gateway", flattenComputeRouteNextHopGateway(res["nextHopGateway"])); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("next_hop_instance", flattenComputeRouteNextHopInstance(res["nextHopInstance"])); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("next_hop_ip", flattenComputeRouteNextHopIp(res["nextHopIp"])); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("next_hop_vpn_tunnel", flattenComputeRouteNextHopVpnTunnel(res["nextHopVpnTunnel"])); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Route: %s", err)
	}

	return nil
}

func resourceComputeRouteDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/routes/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting Route %q", d.Id())
	res, err := Delete(config, url)
	if err != nil {
		return handleNotFoundError(err, d, "Route")
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWaitTime(
		config.clientCompute, op, project, "Deleting Route",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Route %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRouteImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"projects/(?P<project>[^/]+)/global/routes/(?P<name>[^/]+)", "(?P<project>[^/]+)/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRouteDestRange(v interface{}) interface{} {
	return v
}

func flattenComputeRouteDescription(v interface{}) interface{} {
	return v
}

func flattenComputeRouteName(v interface{}) interface{} {
	return v
}

func flattenComputeRouteNetwork(v interface{}) interface{} {
	return v
}

func flattenComputeRoutePriority(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRouteTags(v interface{}) interface{} {
	return v
}

func flattenComputeRouteNextHopGateway(v interface{}) interface{} {
	return v
}

func flattenComputeRouteNextHopInstance(v interface{}) interface{} {
	return v
}

func flattenComputeRouteNextHopIp(v interface{}) interface{} {
	return v
}

func flattenComputeRouteNextHopVpnTunnel(v interface{}) interface{} {
	return v
}

func expandComputeRouteDestRange(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteDescription(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteNetwork(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeRoutePriority(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteTags(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v.(*schema.Set).List(), nil
}

func expandComputeRouteNextHopGateway(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	if v == "default-internet-gateway" {
		project, err := getProject(d, config)
		if err != nil {
			return nil, err
		}
		return fmt.Sprintf("projects/%s/global/gateways/default-internet-gateway", project), nil
	} else {
		return v, nil
	}
}

func expandComputeRouteNextHopInstance(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	if v == "" {
		return v, nil
	}
	val, err := parseZonalFieldValue("instances", v.(string), "project", "next_hop_instance_zone", d, config, true)
	if err != nil {
		return nil, err
	}
	project, err := getProject(d, config)
	if err != nil {
		// Can't happen - we wouldn't make it this far if there wasn't a project.
		return nil, err
	}
	nextInstance, err := config.clientCompute.Instances.Get(project, val.Zone, val.Name).Do()
	if err != nil {
		return nil, err
	}
	return nextInstance.SelfLink, nil
}

func expandComputeRouteNextHopIp(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRouteNextHopVpnTunnel(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceComputeRouteDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if v, ok := res["nextHopInstance"]; ok {
		val, err := parseZonalFieldValue("instances", v.(string), "project", "next_hop_instance_zone", d, meta.(*Config), true)
		if err != nil {
			return nil, err
		}
		d.Set("next_hop_instance_zone", val.Zone)
		res["nextHopInstance"] = val.RelativeLink()
	}

	return res, nil
}
