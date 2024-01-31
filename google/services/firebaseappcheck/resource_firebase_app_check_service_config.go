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

package firebaseappcheck

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceFirebaseAppCheckServiceConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseAppCheckServiceConfigCreate,
		Read:   resourceFirebaseAppCheckServiceConfigRead,
		Update: resourceFirebaseAppCheckServiceConfigUpdate,
		Delete: resourceFirebaseAppCheckServiceConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseAppCheckServiceConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"service_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The identifier of the service to configure enforcement. Currently, the following service IDs are supported:
  firebasestorage.googleapis.com (Cloud Storage for Firebase)
  firebasedatabase.googleapis.com (Firebase Realtime Database)
  firestore.googleapis.com (Cloud Firestore)
  identitytoolkit.googleapis.com (Authentication)`,
			},
			"enforcement_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"UNENFORCED", "ENFORCED", ""}),
				Description: `The App Check enforcement mode for a service supported by App Check. Valid values are

(Unset)
Firebase App Check is not enforced for the service, nor are App Check metrics collected.
Though the service is not protected by App Check in this mode, other applicable protections,
such as user authorization, are still enforced. An unconfigured service is in this mode by default.
This is equivalent to OFF in the REST API. Deleting the Terraform resource will also switch the
enforcement to OFF for this service.

UNENFORCED
Firebase App Check is not enforced for the service. App Check metrics are collected to help you
decide when to turn on enforcement for the service. Though the service is not protected by App Check
in this mode, other applicable protections, such as user authorization, are still enforced.

ENFORCED
Firebase App Check is enforced for the service. The service will reject any request that attempts to
access your project's resources if it does not have valid App Check token attached, with some exceptions
depending on the service; for example, some services will still allow requests bearing the developer's
privileged service account credentials without an App Check token. App Check metrics continue to be
collected to help you detect issues with your App Check integration and monitor the composition of your
callers. While the service is protected by App Check, other applicable protections, such as user
authorization, continue to be enforced at the same time.

Use caution when choosing to enforce App Check on a Firebase service. If your users have not updated
to an App Check capable version of your app, their apps will no longer be able to use your Firebase
services that are enforcing App Check. App Check metrics can help you decide whether to enforce App
Check on your Firebase services.

If your app has not launched yet, you should enable enforcement immediately, since there are no outdated
clients in use. Possible values: ["UNENFORCED", "ENFORCED"]`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The fully-qualified resource name of the service enforcement configuration.`,
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

func resourceFirebaseAppCheckServiceConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	enforcementModeProp, err := expandFirebaseAppCheckServiceConfigEnforcementMode(d.Get("enforcement_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enforcement_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(enforcementModeProp)) && (ok || !reflect.DeepEqual(v, enforcementModeProp)) {
		obj["enforcementMode"] = enforcementModeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseAppCheckBasePath}}projects/{{project}}/services/{{service_id}}?updateMask=enforcementMode")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ServiceConfig: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceConfig: %s", err)
	}
	billingProject = project

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
		return fmt.Errorf("Error creating ServiceConfig: %s", err)
	}
	if err := d.Set("name", flattenFirebaseAppCheckServiceConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/services/{{service_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ServiceConfig %q: %#v", d.Id(), res)

	return resourceFirebaseAppCheckServiceConfigRead(d, meta)
}

func resourceFirebaseAppCheckServiceConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseAppCheckBasePath}}projects/{{project}}/services/{{service_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceConfig: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("FirebaseAppCheckServiceConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ServiceConfig: %s", err)
	}

	if err := d.Set("name", flattenFirebaseAppCheckServiceConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceConfig: %s", err)
	}
	if err := d.Set("enforcement_mode", flattenFirebaseAppCheckServiceConfigEnforcementMode(res["enforcementMode"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServiceConfig: %s", err)
	}

	return nil
}

func resourceFirebaseAppCheckServiceConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	enforcementModeProp, err := expandFirebaseAppCheckServiceConfigEnforcementMode(d.Get("enforcement_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enforcement_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enforcementModeProp)) {
		obj["enforcementMode"] = enforcementModeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseAppCheckBasePath}}projects/{{project}}/services/{{service_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ServiceConfig %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("enforcement_mode") {
		updateMask = append(updateMask, "enforcementMode")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
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
			return fmt.Errorf("Error updating ServiceConfig %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating ServiceConfig %q: %#v", d.Id(), res)
		}

	}

	return resourceFirebaseAppCheckServiceConfigRead(d, meta)
}

func resourceFirebaseAppCheckServiceConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ServiceConfig: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseAppCheckBasePath}}projects/{{project}}/services/{{service_id}}?updateMask=enforcementMode")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ServiceConfig %q", d.Id())

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
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ServiceConfig")
	}

	log.Printf("[DEBUG] Finished deleting ServiceConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceFirebaseAppCheckServiceConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/services/(?P<service_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<service_id>[^/]+)$",
		"^(?P<service_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/services/{{service_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenFirebaseAppCheckServiceConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseAppCheckServiceConfigEnforcementMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandFirebaseAppCheckServiceConfigEnforcementMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
