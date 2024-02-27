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

package logging

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceLoggingOrganizationSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceLoggingOrganizationSettingsCreate,
		Read:   resourceLoggingOrganizationSettingsRead,
		Update: resourceLoggingOrganizationSettingsUpdate,
		Delete: resourceLoggingOrganizationSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: resourceLoggingOrganizationSettingsImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"organization": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The organization for which to retrieve or configure settings.`,
			},
			"disable_default_sink": {
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
				Description: `If set to true, the _Default sink in newly created projects and folders will created in a disabled state. This can be used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The _Default sink can be re-enabled manually if needed.`,
			},
			"kms_key_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: `The resource name for the configured Cloud KMS key.`,
			},
			"kms_service_account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: `The service account that will be used by the Log Router to access your Cloud KMS key. This can be modified only once to migrate from the legacy CMEK service account to the logging service account as described in [Migrate CMEK SA](https://cloud.google.com/logging/docs/routing/troubleshoot-cmek-orgs#migrate-cmek-sa).`,
			},
			"storage_location": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: `The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly provided.`,
			},
			"logging_service_account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The service account for the given container. Sinks use this service account as their writerIdentity if no custom service account is provided.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the settings.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceLoggingOrganizationSettingsCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	kmsKeyNameProp, err := expandLoggingOrganizationSettingsKmsKeyName(d.Get("kms_key_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kms_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(kmsKeyNameProp)) && (ok || !reflect.DeepEqual(v, kmsKeyNameProp)) {
		obj["kmsKeyName"] = kmsKeyNameProp
	}
	kmsServiceAccountIdProp, err := expandLoggingOrganizationSettingsKmsServiceAccountId(d.Get("kms_service_account_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kms_service_account_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(kmsServiceAccountIdProp)) && (ok || !reflect.DeepEqual(v, kmsServiceAccountIdProp)) {
		obj["kmsServiceAccountId"] = kmsServiceAccountIdProp
	}
	storageLocationProp, err := expandLoggingOrganizationSettingsStorageLocation(d.Get("storage_location"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("storage_location"); !tpgresource.IsEmptyValue(reflect.ValueOf(storageLocationProp)) && (ok || !reflect.DeepEqual(v, storageLocationProp)) {
		obj["storageLocation"] = storageLocationProp
	}
	disableDefaultSinkProp, err := expandLoggingOrganizationSettingsDisableDefaultSink(d.Get("disable_default_sink"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disable_default_sink"); !tpgresource.IsEmptyValue(reflect.ValueOf(disableDefaultSinkProp)) && (ok || !reflect.DeepEqual(v, disableDefaultSinkProp)) {
		obj["disableDefaultSink"] = disableDefaultSinkProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{LoggingBasePath}}organizations/{{organization}}/settings?updateMask=disableDefaultSink,storageLocation,kmsKeyName,kmsServiceAccountId")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new OrganizationSettings: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating OrganizationSettings: %s", err)
	}
	if err := d.Set("name", flattenLoggingOrganizationSettingsName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{organization}}/settings")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating OrganizationSettings %q: %#v", d.Id(), res)

	return resourceLoggingOrganizationSettingsRead(d, meta)
}

func resourceLoggingOrganizationSettingsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{LoggingBasePath}}organizations/{{organization}}/settings")
	if err != nil {
		return err
	}

	billingProject := ""

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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("LoggingOrganizationSettings %q", d.Id()))
	}

	if err := d.Set("name", flattenLoggingOrganizationSettingsName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSettings: %s", err)
	}
	if err := d.Set("kms_key_name", flattenLoggingOrganizationSettingsKmsKeyName(res["kmsKeyName"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSettings: %s", err)
	}
	if err := d.Set("kms_service_account_id", flattenLoggingOrganizationSettingsKmsServiceAccountId(res["kmsServiceAccountId"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSettings: %s", err)
	}
	if err := d.Set("storage_location", flattenLoggingOrganizationSettingsStorageLocation(res["storageLocation"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSettings: %s", err)
	}
	if err := d.Set("disable_default_sink", flattenLoggingOrganizationSettingsDisableDefaultSink(res["disableDefaultSink"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSettings: %s", err)
	}
	if err := d.Set("logging_service_account_id", flattenLoggingOrganizationSettingsLoggingServiceAccountId(res["loggingServiceAccountId"], d, config)); err != nil {
		return fmt.Errorf("Error reading OrganizationSettings: %s", err)
	}

	return nil
}

func resourceLoggingOrganizationSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	kmsKeyNameProp, err := expandLoggingOrganizationSettingsKmsKeyName(d.Get("kms_key_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kms_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, kmsKeyNameProp)) {
		obj["kmsKeyName"] = kmsKeyNameProp
	}
	kmsServiceAccountIdProp, err := expandLoggingOrganizationSettingsKmsServiceAccountId(d.Get("kms_service_account_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kms_service_account_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, kmsServiceAccountIdProp)) {
		obj["kmsServiceAccountId"] = kmsServiceAccountIdProp
	}
	storageLocationProp, err := expandLoggingOrganizationSettingsStorageLocation(d.Get("storage_location"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("storage_location"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, storageLocationProp)) {
		obj["storageLocation"] = storageLocationProp
	}
	disableDefaultSinkProp, err := expandLoggingOrganizationSettingsDisableDefaultSink(d.Get("disable_default_sink"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disable_default_sink"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, disableDefaultSinkProp)) {
		obj["disableDefaultSink"] = disableDefaultSinkProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{LoggingBasePath}}organizations/{{organization}}/settings?updateMask=disableDefaultSink,storageLocation,kmsKeyName,kmsServiceAccountId")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating OrganizationSettings %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating OrganizationSettings %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating OrganizationSettings %q: %#v", d.Id(), res)
	}

	return resourceLoggingOrganizationSettingsRead(d, meta)
}

func resourceLoggingOrganizationSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] Logging OrganizationSettings resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceLoggingOrganizationSettingsImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^organizations/(?P<organization>[^/]+)/settings$",
		"^(?P<organization>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{organization}}/settings")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenLoggingOrganizationSettingsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenLoggingOrganizationSettingsKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenLoggingOrganizationSettingsKmsServiceAccountId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenLoggingOrganizationSettingsStorageLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenLoggingOrganizationSettingsDisableDefaultSink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenLoggingOrganizationSettingsLoggingServiceAccountId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandLoggingOrganizationSettingsKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingOrganizationSettingsKmsServiceAccountId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingOrganizationSettingsStorageLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingOrganizationSettingsDisableDefaultSink(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
