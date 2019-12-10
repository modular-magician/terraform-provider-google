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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func migrateStateNoop(v int, is *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	return is, nil
}

func resourceComputeRegionBackendService() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionBackendServiceCreate,
		Read:   resourceComputeRegionBackendServiceRead,
		Update: resourceComputeRegionBackendServiceUpdate,
		Delete: resourceComputeRegionBackendServiceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionBackendServiceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		SchemaVersion: 1,
		MigrateState:  migrateStateNoop,

		Schema: map[string]*schema.Schema{
			"health_checks": {
				Type:     schema.TypeSet,
				Required: true,
				Description: `The set of URLs to HealthCheck resources for health checking
this RegionBackendService. Currently at most one health
check can be specified, and a health check is required.`,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: selfLinkRelativePathHash,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"backend": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: `The set of backends that serve this RegionBackendService.`,
				Elem:        computeRegionBackendServiceBackendSchema(),
				Set:         resourceGoogleComputeBackendServiceBackendHash,
			},
			"connection_draining_timeout_sec": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `Time for which instance will be drained (not accept new
connections, but still work to finish started).`,
				Default: 0,
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An optional description of this resource.`,
			},
			"load_balancing_scheme": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"INTERNAL", "INTERNAL_MANAGED", ""}, false),
				Description: `Indicates what kind of load balancing this regional backend service
will be used for. A backend service created for one type of load
balancing cannot be used with the other(s). Must be 'INTERNAL' or
'INTERNAL_MANAGED'. Defaults to 'INTERNAL'.`,
				Default: "INTERNAL",
			},
			"protocol": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"HTTP", "HTTPS", "HTTP2", "SSL", "TCP", "UDP", ""}, false),
				Description: `The protocol this RegionBackendService uses to communicate with backends.
Possible values are HTTP, HTTPS, HTTP2, SSL, TCP, and UDP. The default is
HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
types and may result in errors if used with the GA API.`,
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `The Region in which the created backend service should reside.
If it is not provided, the provider region is used.`,
			},
			"session_affinity": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "CLIENT_IP", "CLIENT_IP_PORT_PROTO", "CLIENT_IP_PROTO", "GENERATED_COOKIE", "HEADER_FIELD", "HTTP_COOKIE", ""}, false),
				Description: `Type of session affinity to use. The default is NONE. Session affinity is
not applicable if the protocol is UDP.`,
			},
			"timeout_sec": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
				Description: `How many seconds to wait for the backend before considering it a
failed request. Default is 30 seconds. Valid range is [1, 86400].`,
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Fingerprint of this resource. A hash of the contents stored in this
object. This field is used in optimistic locking.`,
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

func computeRegionBackendServiceBackendSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"group": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkRelativePaths,
				Description: `The fully-qualified URL of an Instance Group or Network Endpoint
Group resource. In case of instance group this defines the list
of instances that serve traffic. Member virtual machine
instances from each instance group must live in the same zone as
the instance group itself. No two backends in a backend service
are allowed to use same Instance Group resource.

For Network Endpoint Groups this defines list of endpoints. All
endpoints of Network Endpoint Group must be hosted on instances
located in the same zone as the Network Endpoint Group.

Backend services cannot mix Instance Group and
Network Endpoint Group backends.

When the 'load_balancing_scheme' is INTERNAL, only instance groups
are supported.

Note that you must specify an Instance Group or Network Endpoint
Group resource using the fully-qualified URL, rather than a
partial URL.`,
			},
			"balancing_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"UTILIZATION", "RATE", "CONNECTION", ""}, false),
				Description:  `Specifies the balancing mode for this backend. Defaults to CONNECTION.`,
				Default:      "CONNECTION",
			},
			"capacity_scaler": {
				Type:     schema.TypeFloat,
				Computed: true,
				Optional: true,
				Description: `A multiplier applied to the group's maximum servicing capacity
(based on UTILIZATION, RATE or CONNECTION).

A setting of 0 means the group is completely drained, offering
0% of its available Capacity. Valid range is [0.0,1.0].`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `An optional description of this resource.
Provide this property when you create the resource.`,
			},
			"max_connections": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `The max number of simultaneous connections for the group. Can
be used with either CONNECTION or UTILIZATION balancing modes.

For CONNECTION mode, either maxConnections or one
of maxConnectionsPerInstance or maxConnectionsPerEndpoint,
as appropriate for group type, must be set.`,
			},
			"max_connections_per_endpoint": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `The max number of simultaneous connections that a single backend
network endpoint can handle. This is used to calculate the
capacity of the group. Can be used in either CONNECTION or
UTILIZATION balancing modes.

For CONNECTION mode, either
maxConnections or maxConnectionsPerEndpoint must be set.`,
			},
			"max_connections_per_instance": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `The max number of simultaneous connections that a single
backend instance can handle. This is used to calculate the
capacity of the group. Can be used in either CONNECTION or
UTILIZATION balancing modes.

For CONNECTION mode, either maxConnections or
maxConnectionsPerInstance must be set.`,
			},
			"max_rate": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `The max requests per second (RPS) of the group.

Can be used with either RATE or UTILIZATION balancing modes,
but required if RATE mode. Either maxRate or one
of maxRatePerInstance or maxRatePerEndpoint, as appropriate for
group type, must be set.`,
			},
			"max_rate_per_endpoint": {
				Type:     schema.TypeFloat,
				Optional: true,
				Description: `The max requests per second (RPS) that a single backend network
endpoint can handle. This is used to calculate the capacity of
the group. Can be used in either balancing mode. For RATE mode,
either maxRate or maxRatePerEndpoint must be set.`,
			},
			"max_rate_per_instance": {
				Type:     schema.TypeFloat,
				Optional: true,
				Description: `The max requests per second (RPS) that a single backend
instance can handle. This is used to calculate the capacity of
the group. Can be used in either balancing mode. For RATE mode,
either maxRate or maxRatePerInstance must be set.`,
			},
			"max_utilization": {
				Type:     schema.TypeFloat,
				Optional: true,
				Description: `Used when balancingMode is UTILIZATION. This ratio defines the
CPU utilization target for the group. Valid range is [0.0, 1.0].`,
			},
		},
	}
}

func resourceComputeRegionBackendServiceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	backendsProp, err := expandComputeRegionBackendServiceBackend(d.Get("backend"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backend"); !isEmptyValue(reflect.ValueOf(backendsProp)) && (ok || !reflect.DeepEqual(v, backendsProp)) {
		obj["backends"] = backendsProp
	}
	connectionDrainingProp, err := expandComputeRegionBackendServiceConnectionDraining(nil, d, config)
	if err != nil {
		return err
	} else if !isEmptyValue(reflect.ValueOf(connectionDrainingProp)) {
		obj["connectionDraining"] = connectionDrainingProp
	}
	descriptionProp, err := expandComputeRegionBackendServiceDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeRegionBackendServiceFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	healthChecksProp, err := expandComputeRegionBackendServiceHealthChecks(d.Get("health_checks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("health_checks"); !isEmptyValue(reflect.ValueOf(healthChecksProp)) && (ok || !reflect.DeepEqual(v, healthChecksProp)) {
		obj["healthChecks"] = healthChecksProp
	}
	loadBalancingSchemeProp, err := expandComputeRegionBackendServiceLoadBalancingScheme(d.Get("load_balancing_scheme"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("load_balancing_scheme"); !isEmptyValue(reflect.ValueOf(loadBalancingSchemeProp)) && (ok || !reflect.DeepEqual(v, loadBalancingSchemeProp)) {
		obj["loadBalancingScheme"] = loadBalancingSchemeProp
	}
	nameProp, err := expandComputeRegionBackendServiceName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	protocolProp, err := expandComputeRegionBackendServiceProtocol(d.Get("protocol"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("protocol"); !isEmptyValue(reflect.ValueOf(protocolProp)) && (ok || !reflect.DeepEqual(v, protocolProp)) {
		obj["protocol"] = protocolProp
	}
	sessionAffinityProp, err := expandComputeRegionBackendServiceSessionAffinity(d.Get("session_affinity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("session_affinity"); !isEmptyValue(reflect.ValueOf(sessionAffinityProp)) && (ok || !reflect.DeepEqual(v, sessionAffinityProp)) {
		obj["sessionAffinity"] = sessionAffinityProp
	}
	timeoutSecProp, err := expandComputeRegionBackendServiceTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout_sec"); !isEmptyValue(reflect.ValueOf(timeoutSecProp)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}
	regionProp, err := expandComputeRegionBackendServiceRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/backendServices")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionBackendService: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating RegionBackendService: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating RegionBackendService",
		int(d.Timeout(schema.TimeoutCreate).Minutes()))

	if err != nil {
		// Remove ID to show resource wasn't created.
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionBackendService: %s", err)
	}

	log.Printf("[DEBUG] Finished creating RegionBackendService %q: %#v", d.Id(), res)

	return resourceComputeRegionBackendServiceRead(d, meta)
}

func resourceComputeRegionBackendServiceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRegionBackendService %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}

	if err := d.Set("backend", flattenComputeRegionBackendServiceBackend(res["backends"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if flattenedProp := flattenComputeRegionBackendServiceConnectionDraining(res["connectionDraining"], d); flattenedProp != nil {
		casted := flattenedProp.([]interface{})[0]
		if casted != nil {
			for k, v := range casted.(map[string]interface{}) {
				d.Set(k, v)
			}
		}
	}
	if err := d.Set("creation_timestamp", flattenComputeRegionBackendServiceCreationTimestamp(res["creationTimestamp"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("description", flattenComputeRegionBackendServiceDescription(res["description"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("fingerprint", flattenComputeRegionBackendServiceFingerprint(res["fingerprint"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("health_checks", flattenComputeRegionBackendServiceHealthChecks(res["healthChecks"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("load_balancing_scheme", flattenComputeRegionBackendServiceLoadBalancingScheme(res["loadBalancingScheme"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("name", flattenComputeRegionBackendServiceName(res["name"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("protocol", flattenComputeRegionBackendServiceProtocol(res["protocol"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("session_affinity", flattenComputeRegionBackendServiceSessionAffinity(res["sessionAffinity"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("timeout_sec", flattenComputeRegionBackendServiceTimeoutSec(res["timeoutSec"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("region", flattenComputeRegionBackendServiceRegion(res["region"], d)); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading RegionBackendService: %s", err)
	}

	return nil
}

func resourceComputeRegionBackendServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	backendsProp, err := expandComputeRegionBackendServiceBackend(d.Get("backend"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("backend"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, backendsProp)) {
		obj["backends"] = backendsProp
	}
	connectionDrainingProp, err := expandComputeRegionBackendServiceConnectionDraining(nil, d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("connection_draining"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, connectionDrainingProp)) {
		obj["connectionDraining"] = connectionDrainingProp
	}
	descriptionProp, err := expandComputeRegionBackendServiceDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeRegionBackendServiceFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	healthChecksProp, err := expandComputeRegionBackendServiceHealthChecks(d.Get("health_checks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("health_checks"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, healthChecksProp)) {
		obj["healthChecks"] = healthChecksProp
	}
	loadBalancingSchemeProp, err := expandComputeRegionBackendServiceLoadBalancingScheme(d.Get("load_balancing_scheme"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("load_balancing_scheme"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, loadBalancingSchemeProp)) {
		obj["loadBalancingScheme"] = loadBalancingSchemeProp
	}
	nameProp, err := expandComputeRegionBackendServiceName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	protocolProp, err := expandComputeRegionBackendServiceProtocol(d.Get("protocol"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("protocol"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, protocolProp)) {
		obj["protocol"] = protocolProp
	}
	sessionAffinityProp, err := expandComputeRegionBackendServiceSessionAffinity(d.Get("session_affinity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("session_affinity"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sessionAffinityProp)) {
		obj["sessionAffinity"] = sessionAffinityProp
	}
	timeoutSecProp, err := expandComputeRegionBackendServiceTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("timeout_sec"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}
	regionProp, err := expandComputeRegionBackendServiceRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RegionBackendService %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating RegionBackendService %q: %s", d.Id(), err)
	}

	err = computeOperationWaitTime(
		config, res, project, "Updating RegionBackendService",
		int(d.Timeout(schema.TimeoutUpdate).Minutes()))

	if err != nil {
		return err
	}

	return resourceComputeRegionBackendServiceRead(d, meta)
}

func resourceComputeRegionBackendServiceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting RegionBackendService %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "RegionBackendService")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting RegionBackendService",
		int(d.Timeout(schema.TimeoutDelete).Minutes()))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RegionBackendService %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionBackendServiceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/backendServices/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRegionBackendServiceBackend(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(resourceGoogleComputeBackendServiceBackendHash, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"balancing_mode":               flattenComputeRegionBackendServiceBackendBalancingMode(original["balancingMode"], d),
			"capacity_scaler":              flattenComputeRegionBackendServiceBackendCapacityScaler(original["capacityScaler"], d),
			"description":                  flattenComputeRegionBackendServiceBackendDescription(original["description"], d),
			"group":                        flattenComputeRegionBackendServiceBackendGroup(original["group"], d),
			"max_connections":              flattenComputeRegionBackendServiceBackendMaxConnections(original["maxConnections"], d),
			"max_connections_per_instance": flattenComputeRegionBackendServiceBackendMaxConnectionsPerInstance(original["maxConnectionsPerInstance"], d),
			"max_connections_per_endpoint": flattenComputeRegionBackendServiceBackendMaxConnectionsPerEndpoint(original["maxConnectionsPerEndpoint"], d),
			"max_rate":                     flattenComputeRegionBackendServiceBackendMaxRate(original["maxRate"], d),
			"max_rate_per_instance":        flattenComputeRegionBackendServiceBackendMaxRatePerInstance(original["maxRatePerInstance"], d),
			"max_rate_per_endpoint":        flattenComputeRegionBackendServiceBackendMaxRatePerEndpoint(original["maxRatePerEndpoint"], d),
			"max_utilization":              flattenComputeRegionBackendServiceBackendMaxUtilization(original["maxUtilization"], d),
		})
	}
	return transformed
}
func flattenComputeRegionBackendServiceBackendBalancingMode(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceBackendCapacityScaler(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceBackendDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceBackendGroup(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeRegionBackendServiceBackendMaxConnections(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionBackendServiceBackendMaxConnectionsPerInstance(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionBackendServiceBackendMaxConnectionsPerEndpoint(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionBackendServiceBackendMaxRate(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionBackendServiceBackendMaxRatePerInstance(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceBackendMaxRatePerEndpoint(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceBackendMaxUtilization(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceConnectionDraining(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["connection_draining_timeout_sec"] =
		flattenComputeRegionBackendServiceConnectionDrainingConnectionDrainingTimeoutSec(original["drainingTimeoutSec"], d)
	return []interface{}{transformed}
}
func flattenComputeRegionBackendServiceConnectionDrainingConnectionDrainingTimeoutSec(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionBackendServiceCreationTimestamp(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceFingerprint(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceHealthChecks(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return convertAndMapStringArr(v.([]interface{}), ConvertSelfLinkToV1)
}

func flattenComputeRegionBackendServiceLoadBalancingScheme(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceProtocol(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceSessionAffinity(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenComputeRegionBackendServiceTimeoutSec(v interface{}, d *schema.ResourceData) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenComputeRegionBackendServiceRegion(v interface{}, d *schema.ResourceData) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandComputeRegionBackendServiceBackend(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedBalancingMode, err := expandComputeRegionBackendServiceBackendBalancingMode(original["balancing_mode"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedBalancingMode); val.IsValid() && !isEmptyValue(val) {
			transformed["balancingMode"] = transformedBalancingMode
		}

		transformedCapacityScaler, err := expandComputeRegionBackendServiceBackendCapacityScaler(original["capacity_scaler"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCapacityScaler); val.IsValid() && !isEmptyValue(val) {
			transformed["capacityScaler"] = transformedCapacityScaler
		}

		transformedDescription, err := expandComputeRegionBackendServiceBackendDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedGroup, err := expandComputeRegionBackendServiceBackendGroup(original["group"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGroup); val.IsValid() && !isEmptyValue(val) {
			transformed["group"] = transformedGroup
		}

		transformedMaxConnections, err := expandComputeRegionBackendServiceBackendMaxConnections(original["max_connections"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMaxConnections); val.IsValid() && !isEmptyValue(val) {
			transformed["maxConnections"] = transformedMaxConnections
		}

		transformedMaxConnectionsPerInstance, err := expandComputeRegionBackendServiceBackendMaxConnectionsPerInstance(original["max_connections_per_instance"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMaxConnectionsPerInstance); val.IsValid() && !isEmptyValue(val) {
			transformed["maxConnectionsPerInstance"] = transformedMaxConnectionsPerInstance
		}

		transformedMaxConnectionsPerEndpoint, err := expandComputeRegionBackendServiceBackendMaxConnectionsPerEndpoint(original["max_connections_per_endpoint"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMaxConnectionsPerEndpoint); val.IsValid() && !isEmptyValue(val) {
			transformed["maxConnectionsPerEndpoint"] = transformedMaxConnectionsPerEndpoint
		}

		transformedMaxRate, err := expandComputeRegionBackendServiceBackendMaxRate(original["max_rate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMaxRate); val.IsValid() && !isEmptyValue(val) {
			transformed["maxRate"] = transformedMaxRate
		}

		transformedMaxRatePerInstance, err := expandComputeRegionBackendServiceBackendMaxRatePerInstance(original["max_rate_per_instance"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMaxRatePerInstance); val.IsValid() && !isEmptyValue(val) {
			transformed["maxRatePerInstance"] = transformedMaxRatePerInstance
		}

		transformedMaxRatePerEndpoint, err := expandComputeRegionBackendServiceBackendMaxRatePerEndpoint(original["max_rate_per_endpoint"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMaxRatePerEndpoint); val.IsValid() && !isEmptyValue(val) {
			transformed["maxRatePerEndpoint"] = transformedMaxRatePerEndpoint
		}

		transformedMaxUtilization, err := expandComputeRegionBackendServiceBackendMaxUtilization(original["max_utilization"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMaxUtilization); val.IsValid() && !isEmptyValue(val) {
			transformed["maxUtilization"] = transformedMaxUtilization
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeRegionBackendServiceBackendBalancingMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceBackendCapacityScaler(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceBackendDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceBackendGroup(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceBackendMaxConnections(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceBackendMaxConnectionsPerInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceBackendMaxConnectionsPerEndpoint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceBackendMaxRate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceBackendMaxRatePerInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceBackendMaxRatePerEndpoint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceBackendMaxUtilization(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceConnectionDraining(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedConnectionDrainingTimeoutSec, err := expandComputeRegionBackendServiceConnectionDrainingConnectionDrainingTimeoutSec(d.Get("connection_draining_timeout_sec"), d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["drainingTimeoutSec"] = transformedConnectionDrainingTimeoutSec
	}

	return transformed, nil
}

func expandComputeRegionBackendServiceConnectionDrainingConnectionDrainingTimeoutSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceHealthChecks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeRegionBackendServiceLoadBalancingScheme(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceSessionAffinity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceTimeoutSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionBackendServiceRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
