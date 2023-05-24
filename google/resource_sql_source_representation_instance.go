// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"

	"google.golang.org/api/googleapi"
)

func ResourceSQLSourceRepresentationInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceSQLSourceRepresentationInstanceCreate,
		Read:   resourceSQLSourceRepresentationInstanceRead,
		Delete: resourceSQLSourceRepresentationInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSQLSourceRepresentationInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"database_version": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"MYSQL_5_6", "MYSQL_5_7", "MYSQL_8_0", "POSTGRES_9_6", "POSTGRES_10", "POSTGRES_11", "POSTGRES_12", "POSTGRES_13", "POSTGRES_14"}),
				Description:  `The MySQL version running on your source database server. Possible values: ["MYSQL_5_6", "MYSQL_5_7", "MYSQL_8_0", "POSTGRES_9_6", "POSTGRES_10", "POSTGRES_11", "POSTGRES_12", "POSTGRES_13", "POSTGRES_14"]`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the source representation instance. Use any valid Cloud SQL instance name.`,
			},
			"host": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateIpAddress,
				Description:  `The externally accessible IPv4 address for the source database server.`,
			},
			"ca_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The CA certificate on the external server. Include only if SSL/TLS is used on the external server.`,
			},
			"client_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The client certificate on the external server. Required only for server-client authentication. Include only if SSL/TLS is used on the external server.`,
			},
			"client_key": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The private key file for the client certificate on the external server. Required only for server-client authentication. Include only if SSL/TLS is used on the external server.`,
			},
			"dump_file_path": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `A file in the bucket that contains the data from the external server.`,
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The password for the replication user account.`,
				Sensitive:   true,
			},
			"port": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(0, 65535),
				Description: `The externally accessible port for the source database server.
Defaults to 3306.`,
				Default: 3306,
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The replication user account on the external server.`,
			},

			"region": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `The Region in which the created instance should reside.
If it is not provided, the provider region is used.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceSQLSourceRepresentationInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandSQLSourceRepresentationInstanceName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	regionProp, err := expandSQLSourceRepresentationInstanceRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}
	databaseVersionProp, err := expandSQLSourceRepresentationInstanceDatabaseVersion(d.Get("database_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("database_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(databaseVersionProp)) && (ok || !reflect.DeepEqual(v, databaseVersionProp)) {
		obj["databaseVersion"] = databaseVersionProp
	}
	onPremisesConfigurationProp, err := expandSQLSourceRepresentationInstanceOnPremisesConfiguration(nil, d, config)
	if err != nil {
		return err
	} else if !tpgresource.IsEmptyValue(reflect.ValueOf(onPremisesConfigurationProp)) {
		obj["onPremisesConfiguration"] = onPremisesConfigurationProp
	}

	obj, err = resourceSQLSourceRepresentationInstanceEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SQLBasePath}}projects/{{project}}/instances")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SourceRepresentationInstance: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SourceRepresentationInstance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating SourceRepresentationInstance: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/instances/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = SqlAdminOperationWaitTime(
		config, res, project, "Creating SourceRepresentationInstance", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create SourceRepresentationInstance: %s", err)
	}

	log.Printf("[DEBUG] Finished creating SourceRepresentationInstance %q: %#v", d.Id(), res)

	return resourceSQLSourceRepresentationInstanceRead(d, meta)
}

func resourceSQLSourceRepresentationInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SQLBasePath}}projects/{{project}}/instances/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SourceRepresentationInstance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SQLSourceRepresentationInstance %q", d.Id()))
	}

	res, err = resourceSQLSourceRepresentationInstanceDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing SQLSourceRepresentationInstance because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading SourceRepresentationInstance: %s", err)
	}

	if err := d.Set("name", flattenSQLSourceRepresentationInstanceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading SourceRepresentationInstance: %s", err)
	}
	if err := d.Set("region", flattenSQLSourceRepresentationInstanceRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading SourceRepresentationInstance: %s", err)
	}
	if err := d.Set("database_version", flattenSQLSourceRepresentationInstanceDatabaseVersion(res["databaseVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading SourceRepresentationInstance: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if flattenedProp := flattenSQLSourceRepresentationInstanceOnPremisesConfiguration(res["onPremisesConfiguration"], d, config); flattenedProp != nil {
		if gerr, ok := flattenedProp.(*googleapi.Error); ok {
			return fmt.Errorf("Error reading SourceRepresentationInstance: %s", gerr)
		}
		casted := flattenedProp.([]interface{})[0]
		if casted != nil {
			for k, v := range casted.(map[string]interface{}) {
				if err := d.Set(k, v); err != nil {
					return fmt.Errorf("Error setting %s: %s", k, err)
				}
			}
		}
	}

	return nil
}

func resourceSQLSourceRepresentationInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SourceRepresentationInstance: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{SQLBasePath}}projects/{{project}}/instances/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting SourceRepresentationInstance %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "SourceRepresentationInstance")
	}

	err = SqlAdminOperationWaitTime(
		config, res, project, "Deleting SourceRepresentationInstance", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting SourceRepresentationInstance %q: %#v", d.Id(), res)
	return nil
}

func resourceSQLSourceRepresentationInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/instances/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/instances/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenSQLSourceRepresentationInstanceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSQLSourceRepresentationInstanceRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSQLSourceRepresentationInstanceDatabaseVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSQLSourceRepresentationInstanceOnPremisesConfiguration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["host"] =
		flattenSQLSourceRepresentationInstanceOnPremisesConfigurationHost(original["host"], d, config)
	transformed["port"] =
		flattenSQLSourceRepresentationInstanceOnPremisesConfigurationPort(original["port"], d, config)
	transformed["username"] =
		flattenSQLSourceRepresentationInstanceOnPremisesConfigurationUsername(original["username"], d, config)
	transformed["password"] =
		flattenSQLSourceRepresentationInstanceOnPremisesConfigurationPassword(original["password"], d, config)
	transformed["dump_file_path"] =
		flattenSQLSourceRepresentationInstanceOnPremisesConfigurationDumpFilePath(original["dumpFilePath"], d, config)
	transformed["ca_certificate"] =
		flattenSQLSourceRepresentationInstanceOnPremisesConfigurationCaCertificate(original["caCertificate"], d, config)
	transformed["client_certificate"] =
		flattenSQLSourceRepresentationInstanceOnPremisesConfigurationClientCertificate(original["clientCertificate"], d, config)
	transformed["client_key"] =
		flattenSQLSourceRepresentationInstanceOnPremisesConfigurationClientKey(original["clientKey"], d, config)
	return []interface{}{transformed}
}
func flattenSQLSourceRepresentationInstanceOnPremisesConfigurationHost(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSQLSourceRepresentationInstanceOnPremisesConfigurationPort(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenSQLSourceRepresentationInstanceOnPremisesConfigurationUsername(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSQLSourceRepresentationInstanceOnPremisesConfigurationPassword(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return d.Get("password")
}

func flattenSQLSourceRepresentationInstanceOnPremisesConfigurationDumpFilePath(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSQLSourceRepresentationInstanceOnPremisesConfigurationCaCertificate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSQLSourceRepresentationInstanceOnPremisesConfigurationClientCertificate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSQLSourceRepresentationInstanceOnPremisesConfigurationClientKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSQLSourceRepresentationInstanceName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSQLSourceRepresentationInstanceRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSQLSourceRepresentationInstanceDatabaseVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSQLSourceRepresentationInstanceOnPremisesConfiguration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedHost, err := expandSQLSourceRepresentationInstanceOnPremisesConfigurationHost(d.Get("host"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	transformedPort, err := expandSQLSourceRepresentationInstanceOnPremisesConfigurationPort(d.Get("port"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedUsername, err := expandSQLSourceRepresentationInstanceOnPremisesConfigurationUsername(d.Get("username"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUsername); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["username"] = transformedUsername
	}

	transformedPassword, err := expandSQLSourceRepresentationInstanceOnPremisesConfigurationPassword(d.Get("password"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPassword); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["password"] = transformedPassword
	}

	transformedDumpFilePath, err := expandSQLSourceRepresentationInstanceOnPremisesConfigurationDumpFilePath(d.Get("dump_file_path"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDumpFilePath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dumpFilePath"] = transformedDumpFilePath
	}

	transformedCaCertificate, err := expandSQLSourceRepresentationInstanceOnPremisesConfigurationCaCertificate(d.Get("ca_certificate"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCaCertificate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["caCertificate"] = transformedCaCertificate
	}

	transformedClientCertificate, err := expandSQLSourceRepresentationInstanceOnPremisesConfigurationClientCertificate(d.Get("client_certificate"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientCertificate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientCertificate"] = transformedClientCertificate
	}

	transformedClientKey, err := expandSQLSourceRepresentationInstanceOnPremisesConfigurationClientKey(d.Get("client_key"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clientKey"] = transformedClientKey
	}

	return transformed, nil
}

func expandSQLSourceRepresentationInstanceOnPremisesConfigurationHost(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSQLSourceRepresentationInstanceOnPremisesConfigurationPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSQLSourceRepresentationInstanceOnPremisesConfigurationUsername(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSQLSourceRepresentationInstanceOnPremisesConfigurationPassword(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSQLSourceRepresentationInstanceOnPremisesConfigurationDumpFilePath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSQLSourceRepresentationInstanceOnPremisesConfigurationCaCertificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSQLSourceRepresentationInstanceOnPremisesConfigurationClientCertificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSQLSourceRepresentationInstanceOnPremisesConfigurationClientKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceSQLSourceRepresentationInstanceEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	opc := obj["onPremisesConfiguration"].(map[string]interface{})
	opc["hostPort"] = fmt.Sprintf("%v:%v", opc["host"], opc["port"])
	delete(opc, "host")
	delete(opc, "port")
	return obj, nil
}

func resourceSQLSourceRepresentationInstanceDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if v, ok := res["onPremisesConfiguration"]; ok {
		opc := v.(map[string]interface{})
		hostPort := opc["hostPort"]
		spl := strings.Split(hostPort.(string), ":")
		if len(spl) != 2 {
			return nil, fmt.Errorf("unexpected value for hostPort, expected [host]:[port], got %q", hostPort)
		}
		opc["host"] = spl[0]
		p, err := strconv.Atoi(spl[1])
		if err != nil {
			return nil, fmt.Errorf("error converting port %q to int: %v", spl[1], err)
		}
		opc["port"] = p
		delete(opc, "hostPort")
	}
	return res, nil
}
