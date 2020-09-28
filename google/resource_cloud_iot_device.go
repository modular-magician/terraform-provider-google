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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudIotDevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudIotDeviceCreate,
		Read:   resourceCloudIotDeviceRead,
		Update: resourceCloudIotDeviceUpdate,
		Delete: resourceCloudIotDeviceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCloudIotDeviceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `A unique name for the resource.`,
			},
			"registry": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the device registry where this device should be created.`,
			},
			"blocked": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `If a device is blocked, connections or requests from this device will fail.`,
			},
			"credentials": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The credentials used to authenticate this device.`,
				MaxItems:    3,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"public_key": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `A public key used to verify the signature of JSON Web Tokens (JWTs).`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"format": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"RSA_PEM", "RSA_X509_PEM", "ES256_PEM", "ES256_X509_PEM"}, false),
										Description:  `The format of the key. Possible values: ["RSA_PEM", "RSA_X509_PEM", "ES256_PEM", "ES256_X509_PEM"]`,
									},
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The key data.`,
									},
								},
							},
						},
						"expiration_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: `The time at which this credential becomes invalid.`,
						},
					},
				},
			},
			"gateway_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Gateway-related configuration and state.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gateway_auth_method": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"ASSOCIATION_ONLY", "DEVICE_AUTH_TOKEN_ONLY", "ASSOCIATION_AND_DEVICE_AUTH_TOKEN", ""}, false),
							Description:  `Indicates whether the device is a gateway. Possible values: ["ASSOCIATION_ONLY", "DEVICE_AUTH_TOKEN_ONLY", "ASSOCIATION_AND_DEVICE_AUTH_TOKEN"]`,
						},
						"gateway_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"GATEWAY", "NON_GATEWAY", ""}, false),
							Description:  `Indicates whether the device is a gateway. Default value: "NON_GATEWAY" Possible values: ["GATEWAY", "NON_GATEWAY"]`,
							Default:      "NON_GATEWAY",
						},
						"last_accessed_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The ID of the gateway the device accessed most recently.`,
						},
						"last_accessed_gateway_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The most recent time at which the device accessed the gateway specified in last_accessed_gateway.`,
						},
					},
				},
			},
			"log_level": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"NONE", "ERROR", "INFO", "DEBUG", ""}, false),
				Description:  `The logging verbosity for device activity. Possible values: ["NONE", "ERROR", "INFO", "DEBUG"]`,
			},
			"metadata": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `The metadata key-value pairs assigned to the device.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"config": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"binary_data": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The device configuration data.`,
						},
						"cloud_update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The time at which this configuration version was updated in Cloud IoT Core.`,
						},
						"device_ack_time": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `The time at which Cloud IoT Core received the acknowledgment from the device,
indicating that the device has received this configuration version.`,
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The version of this update.`,
						},
					},
				},
			},
			"last_config_ack_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The last time a cloud-to-device config version acknowledgment was received from the device.`,
			},
			"last_config_send_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The last time a cloud-to-device config version was sent to the device.`,
			},
			"last_error_status": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"details": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `A list of messages that carry the error details.`,
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
						"message": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `A developer-facing error message, which should be in English.`,
						},
						"number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: `The status code, which should be an enum value of google.rpc.Code.`,
						},
					},
				},
			},
			"last_error_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.`,
			},
			"last_event_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The last time a telemetry event was received.`,
			},
			"last_heartbeat_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The last time an MQTT PINGREQ was received.`,
			},
			"last_state_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The last time a state event was received.`,
			},
			"num_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `A server-defined unique numeric ID for the device.
This is a more compact way to identify devices, and it is globally unique.`,
			},
			"state": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The state most recently received from the device.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"binary_data": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The device state data.`,
						},
						"update_time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The time at which this state version was updated in Cloud IoT Core.`,
						},
					},
				},
			},
		},
	}
}

func resourceCloudIotDeviceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	idProp, err := expandCloudIotDeviceName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	credentialsProp, err := expandCloudIotDeviceCredentials(d.Get("credentials"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("credentials"); !isEmptyValue(reflect.ValueOf(credentialsProp)) && (ok || !reflect.DeepEqual(v, credentialsProp)) {
		obj["credentials"] = credentialsProp
	}
	blockedProp, err := expandCloudIotDeviceBlocked(d.Get("blocked"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("blocked"); !isEmptyValue(reflect.ValueOf(blockedProp)) && (ok || !reflect.DeepEqual(v, blockedProp)) {
		obj["blocked"] = blockedProp
	}
	logLevelProp, err := expandCloudIotDeviceLogLevel(d.Get("log_level"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("log_level"); !isEmptyValue(reflect.ValueOf(logLevelProp)) && (ok || !reflect.DeepEqual(v, logLevelProp)) {
		obj["logLevel"] = logLevelProp
	}
	metadataProp, err := expandCloudIotDeviceMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}
	gatewayConfigProp, err := expandCloudIotDeviceGatewayConfig(d.Get("gateway_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gateway_config"); !isEmptyValue(reflect.ValueOf(gatewayConfigProp)) && (ok || !reflect.DeepEqual(v, gatewayConfigProp)) {
		obj["gatewayConfig"] = gatewayConfigProp
	}

	url, err := replaceVars(d, config, "{{CloudIotBasePath}}{{registry}}/devices")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Device: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Device: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{registry}}/devices/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Device %q: %#v", d.Id(), res)

	return resourceCloudIotDeviceRead(d, meta)
}

func resourceCloudIotDeviceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{CloudIotBasePath}}{{registry}}/devices/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("CloudIotDevice %q", d.Id()))
	}

	if err := d.Set("name", flattenCloudIotDeviceName(res["id"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}
	if err := d.Set("num_id", flattenCloudIotDeviceNumId(res["numId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}
	if err := d.Set("credentials", flattenCloudIotDeviceCredentials(res["credentials"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}
	if err := d.Set("last_heartbeat_time", flattenCloudIotDeviceLastHeartbeatTime(res["lastHeartbeatTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}
	if err := d.Set("last_event_time", flattenCloudIotDeviceLastEventTime(res["lastEventTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}
	if err := d.Set("last_state_time", flattenCloudIotDeviceLastStateTime(res["lastStateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}
	if err := d.Set("last_config_ack_time", flattenCloudIotDeviceLastConfigAckTime(res["lastConfigAckTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}
	if err := d.Set("last_config_send_time", flattenCloudIotDeviceLastConfigSendTime(res["lastConfigSendTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}
	if err := d.Set("blocked", flattenCloudIotDeviceBlocked(res["blocked"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}
	if err := d.Set("last_error_time", flattenCloudIotDeviceLastErrorTime(res["lastErrorTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}
	if err := d.Set("last_error_status", flattenCloudIotDeviceLastErrorStatus(res["lastErrorStatus"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}
	if err := d.Set("config", flattenCloudIotDeviceConfig(res["config"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}
	if err := d.Set("state", flattenCloudIotDeviceState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}
	if err := d.Set("log_level", flattenCloudIotDeviceLogLevel(res["logLevel"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}
	if err := d.Set("metadata", flattenCloudIotDeviceMetadata(res["metadata"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}
	if err := d.Set("gateway_config", flattenCloudIotDeviceGatewayConfig(res["gatewayConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Device: %s", err)
	}

	return nil
}

func resourceCloudIotDeviceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	config.userAgent = userAgent

	billingProject := ""

	obj := make(map[string]interface{})
	credentialsProp, err := expandCloudIotDeviceCredentials(d.Get("credentials"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("credentials"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, credentialsProp)) {
		obj["credentials"] = credentialsProp
	}
	blockedProp, err := expandCloudIotDeviceBlocked(d.Get("blocked"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("blocked"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, blockedProp)) {
		obj["blocked"] = blockedProp
	}
	logLevelProp, err := expandCloudIotDeviceLogLevel(d.Get("log_level"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("log_level"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, logLevelProp)) {
		obj["logLevel"] = logLevelProp
	}
	metadataProp, err := expandCloudIotDeviceMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}
	gatewayConfigProp, err := expandCloudIotDeviceGatewayConfig(d.Get("gateway_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gateway_config"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, gatewayConfigProp)) {
		obj["gatewayConfig"] = gatewayConfigProp
	}

	url, err := replaceVars(d, config, "{{CloudIotBasePath}}{{registry}}/devices/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Device %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("credentials") {
		updateMask = append(updateMask, "credentials")
	}

	if d.HasChange("blocked") {
		updateMask = append(updateMask, "blocked")
	}

	if d.HasChange("log_level") {
		updateMask = append(updateMask, "logLevel")
	}

	if d.HasChange("metadata") {
		updateMask = append(updateMask, "metadata")
	}

	if d.HasChange("gateway_config") {
		updateMask = append(updateMask, "gateway_config.gateway_auth_method")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Device %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Device %q: %#v", d.Id(), res)
	}

	return resourceCloudIotDeviceRead(d, meta)
}

func resourceCloudIotDeviceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	config.userAgent = userAgent

	billingProject := ""

	url, err := replaceVars(d, config, "{{CloudIotBasePath}}{{registry}}/devices/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Device %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Device")
	}

	log.Printf("[DEBUG] Finished deleting Device %q: %#v", d.Id(), res)
	return nil
}

func resourceCloudIotDeviceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"(?P<registry>.+)/devices/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{registry}}/devices/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenCloudIotDeviceName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceNumId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceCredentials(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"expiration_time": flattenCloudIotDeviceCredentialsExpirationTime(original["expirationTime"], d, config),
			"public_key":      flattenCloudIotDeviceCredentialsPublicKey(original["publicKey"], d, config),
		})
	}
	return transformed
}
func flattenCloudIotDeviceCredentialsExpirationTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceCredentialsPublicKey(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["format"] =
		flattenCloudIotDeviceCredentialsPublicKeyFormat(original["format"], d, config)
	transformed["key"] =
		flattenCloudIotDeviceCredentialsPublicKeyKey(original["key"], d, config)
	return []interface{}{transformed}
}
func flattenCloudIotDeviceCredentialsPublicKeyFormat(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceCredentialsPublicKeyKey(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceLastHeartbeatTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceLastEventTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceLastStateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceLastConfigAckTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceLastConfigSendTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceBlocked(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceLastErrorTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceLastErrorStatus(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["number"] =
		flattenCloudIotDeviceLastErrorStatusNumber(original["number"], d, config)
	transformed["message"] =
		flattenCloudIotDeviceLastErrorStatusMessage(original["message"], d, config)
	transformed["details"] =
		flattenCloudIotDeviceLastErrorStatusDetails(original["details"], d, config)
	return []interface{}{transformed}
}
func flattenCloudIotDeviceLastErrorStatusNumber(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
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

func flattenCloudIotDeviceLastErrorStatusMessage(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceLastErrorStatusDetails(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["version"] =
		flattenCloudIotDeviceConfigVersion(original["version"], d, config)
	transformed["cloud_update_time"] =
		flattenCloudIotDeviceConfigCloudUpdateTime(original["cloudUpdateTime"], d, config)
	transformed["device_ack_time"] =
		flattenCloudIotDeviceConfigDeviceAckTime(original["deviceAckTime"], d, config)
	transformed["binary_data"] =
		flattenCloudIotDeviceConfigBinaryData(original["binaryData"], d, config)
	return []interface{}{transformed}
}
func flattenCloudIotDeviceConfigVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceConfigCloudUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceConfigDeviceAckTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceConfigBinaryData(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["update_time"] =
		flattenCloudIotDeviceStateUpdateTime(original["updateTime"], d, config)
	transformed["binary_data"] =
		flattenCloudIotDeviceStateBinaryData(original["binaryData"], d, config)
	return []interface{}{transformed}
}
func flattenCloudIotDeviceStateUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceStateBinaryData(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceLogLevel(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceMetadata(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceGatewayConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["gateway_type"] =
		flattenCloudIotDeviceGatewayConfigGatewayType(original["gatewayType"], d, config)
	transformed["gateway_auth_method"] =
		flattenCloudIotDeviceGatewayConfigGatewayAuthMethod(original["gatewayAuthMethod"], d, config)
	transformed["last_accessed_gateway_id"] =
		flattenCloudIotDeviceGatewayConfigLastAccessedGatewayId(original["lastAccessedGatewayId"], d, config)
	transformed["last_accessed_gateway_time"] =
		flattenCloudIotDeviceGatewayConfigLastAccessedGatewayTime(original["lastAccessedGatewayTime"], d, config)
	return []interface{}{transformed}
}
func flattenCloudIotDeviceGatewayConfigGatewayType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceGatewayConfigGatewayAuthMethod(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceGatewayConfigLastAccessedGatewayId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCloudIotDeviceGatewayConfigLastAccessedGatewayTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandCloudIotDeviceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceCredentials(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedExpirationTime, err := expandCloudIotDeviceCredentialsExpirationTime(original["expiration_time"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedExpirationTime); val.IsValid() && !isEmptyValue(val) {
			transformed["expirationTime"] = transformedExpirationTime
		}

		transformedPublicKey, err := expandCloudIotDeviceCredentialsPublicKey(original["public_key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPublicKey); val.IsValid() && !isEmptyValue(val) {
			transformed["publicKey"] = transformedPublicKey
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCloudIotDeviceCredentialsExpirationTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceCredentialsPublicKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFormat, err := expandCloudIotDeviceCredentialsPublicKeyFormat(original["format"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFormat); val.IsValid() && !isEmptyValue(val) {
		transformed["format"] = transformedFormat
	}

	transformedKey, err := expandCloudIotDeviceCredentialsPublicKeyKey(original["key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !isEmptyValue(val) {
		transformed["key"] = transformedKey
	}

	return transformed, nil
}

func expandCloudIotDeviceCredentialsPublicKeyFormat(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceCredentialsPublicKeyKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceBlocked(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceLogLevel(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceMetadata(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCloudIotDeviceGatewayConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGatewayType, err := expandCloudIotDeviceGatewayConfigGatewayType(original["gateway_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGatewayType); val.IsValid() && !isEmptyValue(val) {
		transformed["gatewayType"] = transformedGatewayType
	}

	transformedGatewayAuthMethod, err := expandCloudIotDeviceGatewayConfigGatewayAuthMethod(original["gateway_auth_method"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGatewayAuthMethod); val.IsValid() && !isEmptyValue(val) {
		transformed["gatewayAuthMethod"] = transformedGatewayAuthMethod
	}

	transformedLastAccessedGatewayId, err := expandCloudIotDeviceGatewayConfigLastAccessedGatewayId(original["last_accessed_gateway_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLastAccessedGatewayId); val.IsValid() && !isEmptyValue(val) {
		transformed["lastAccessedGatewayId"] = transformedLastAccessedGatewayId
	}

	transformedLastAccessedGatewayTime, err := expandCloudIotDeviceGatewayConfigLastAccessedGatewayTime(original["last_accessed_gateway_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLastAccessedGatewayTime); val.IsValid() && !isEmptyValue(val) {
		transformed["lastAccessedGatewayTime"] = transformedLastAccessedGatewayTime
	}

	return transformed, nil
}

func expandCloudIotDeviceGatewayConfigGatewayType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceGatewayConfigGatewayAuthMethod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceGatewayConfigLastAccessedGatewayId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIotDeviceGatewayConfigLastAccessedGatewayTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
