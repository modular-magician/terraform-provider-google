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

package identityplatform

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
)

func ResourceIdentityPlatformTenantOauthIdpConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityPlatformTenantOauthIdpConfigCreate,
		Read:   resourceIdentityPlatformTenantOauthIdpConfigRead,
		Update: resourceIdentityPlatformTenantOauthIdpConfigUpdate,
		Delete: resourceIdentityPlatformTenantOauthIdpConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIdentityPlatformTenantOauthIdpConfigImport,
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
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The client id of an OAuth client.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Human friendly display name.`,
			},
			"issuer": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `For OIDC Idps, the issuer identifier.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the OauthIdpConfig. Must start with 'oidc.'.`,
			},
			"tenant": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the tenant where this OIDC IDP configuration resource exists`,
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The client secret of the OAuth client, to enable OIDC code flow.`,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `If this config allows users to sign in with the provider.`,
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

func resourceIdentityPlatformTenantOauthIdpConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandIdentityPlatformTenantOauthIdpConfigName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	displayNameProp, err := expandIdentityPlatformTenantOauthIdpConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enabledProp, err := expandIdentityPlatformTenantOauthIdpConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}
	issuerProp, err := expandIdentityPlatformTenantOauthIdpConfigIssuer(d.Get("issuer"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("issuer"); !tpgresource.IsEmptyValue(reflect.ValueOf(issuerProp)) && (ok || !reflect.DeepEqual(v, issuerProp)) {
		obj["issuer"] = issuerProp
	}
	clientIdProp, err := expandIdentityPlatformTenantOauthIdpConfigClientId(d.Get("client_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(clientIdProp)) && (ok || !reflect.DeepEqual(v, clientIdProp)) {
		obj["clientId"] = clientIdProp
	}
	clientSecretProp, err := expandIdentityPlatformTenantOauthIdpConfigClientSecret(d.Get("client_secret"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_secret"); !tpgresource.IsEmptyValue(reflect.ValueOf(clientSecretProp)) && (ok || !reflect.DeepEqual(v, clientSecretProp)) {
		obj["clientSecret"] = clientSecretProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs?oauthIdpConfigId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TenantOauthIdpConfig: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TenantOauthIdpConfig: %s", err)
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
		return fmt.Errorf("Error creating TenantOauthIdpConfig: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating TenantOauthIdpConfig %q: %#v", d.Id(), res)

	return resourceIdentityPlatformTenantOauthIdpConfigRead(d, meta)
}

func resourceIdentityPlatformTenantOauthIdpConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TenantOauthIdpConfig: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IdentityPlatformTenantOauthIdpConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading TenantOauthIdpConfig: %s", err)
	}

	if err := d.Set("name", flattenIdentityPlatformTenantOauthIdpConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading TenantOauthIdpConfig: %s", err)
	}
	if err := d.Set("display_name", flattenIdentityPlatformTenantOauthIdpConfigDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading TenantOauthIdpConfig: %s", err)
	}
	if err := d.Set("enabled", flattenIdentityPlatformTenantOauthIdpConfigEnabled(res["enabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading TenantOauthIdpConfig: %s", err)
	}
	if err := d.Set("issuer", flattenIdentityPlatformTenantOauthIdpConfigIssuer(res["issuer"], d, config)); err != nil {
		return fmt.Errorf("Error reading TenantOauthIdpConfig: %s", err)
	}
	if err := d.Set("client_id", flattenIdentityPlatformTenantOauthIdpConfigClientId(res["clientId"], d, config)); err != nil {
		return fmt.Errorf("Error reading TenantOauthIdpConfig: %s", err)
	}
	if err := d.Set("client_secret", flattenIdentityPlatformTenantOauthIdpConfigClientSecret(res["clientSecret"], d, config)); err != nil {
		return fmt.Errorf("Error reading TenantOauthIdpConfig: %s", err)
	}

	return nil
}

func resourceIdentityPlatformTenantOauthIdpConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TenantOauthIdpConfig: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandIdentityPlatformTenantOauthIdpConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enabledProp, err := expandIdentityPlatformTenantOauthIdpConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}
	issuerProp, err := expandIdentityPlatformTenantOauthIdpConfigIssuer(d.Get("issuer"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("issuer"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, issuerProp)) {
		obj["issuer"] = issuerProp
	}
	clientIdProp, err := expandIdentityPlatformTenantOauthIdpConfigClientId(d.Get("client_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clientIdProp)) {
		obj["clientId"] = clientIdProp
	}
	clientSecretProp, err := expandIdentityPlatformTenantOauthIdpConfigClientSecret(d.Get("client_secret"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("client_secret"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, clientSecretProp)) {
		obj["clientSecret"] = clientSecretProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating TenantOauthIdpConfig %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("enabled") {
		updateMask = append(updateMask, "enabled")
	}

	if d.HasChange("issuer") {
		updateMask = append(updateMask, "issuer")
	}

	if d.HasChange("client_id") {
		updateMask = append(updateMask, "clientId")
	}

	if d.HasChange("client_secret") {
		updateMask = append(updateMask, "clientSecret")
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
		return fmt.Errorf("Error updating TenantOauthIdpConfig %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating TenantOauthIdpConfig %q: %#v", d.Id(), res)
	}

	return resourceIdentityPlatformTenantOauthIdpConfigRead(d, meta)
}

func resourceIdentityPlatformTenantOauthIdpConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TenantOauthIdpConfig: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting TenantOauthIdpConfig %q", d.Id())

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
		return transport_tpg.HandleNotFoundError(err, d, "TenantOauthIdpConfig")
	}

	log.Printf("[DEBUG] Finished deleting TenantOauthIdpConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceIdentityPlatformTenantOauthIdpConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/tenants/(?P<tenant>[^/]+)/oauthIdpConfigs/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<tenant>[^/]+)/(?P<name>[^/]+)",
		"(?P<tenant>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/tenants/{{tenant}}/oauthIdpConfigs/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIdentityPlatformTenantOauthIdpConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenIdentityPlatformTenantOauthIdpConfigDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantOauthIdpConfigEnabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantOauthIdpConfigIssuer(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantOauthIdpConfigClientId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformTenantOauthIdpConfigClientSecret(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIdentityPlatformTenantOauthIdpConfigName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantOauthIdpConfigDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantOauthIdpConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantOauthIdpConfigIssuer(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantOauthIdpConfigClientId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantOauthIdpConfigClientSecret(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
