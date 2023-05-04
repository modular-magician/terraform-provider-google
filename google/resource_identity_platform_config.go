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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceIdentityPlatformConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityPlatformConfigCreate,
		Read:   resourceIdentityPlatformConfigRead,
		Update: resourceIdentityPlatformConfigUpdate,
		Delete: resourceIdentityPlatformConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIdentityPlatformConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"autodelete_anonymous_users": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Whether anonymous users will be auto-deleted after a period of 30 days`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the Config resource`,
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

func resourceIdentityPlatformConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/identityPlatform:initializeAuth")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Config: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, nil, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Config: %s", err)
	}
	if err := d.Set("name", flattenIdentityPlatformConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "projects/{{project}}/config")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Update the resource after initializing auth to set fields.
	if err := resourceIdentityPlatformConfigUpdate(d, meta); err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished creating Config %q: %#v", d.Id(), res)

	return resourceIdentityPlatformConfigRead(d, meta)
}

func resourceIdentityPlatformConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/config")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Config: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IdentityPlatformConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}

	if err := d.Set("name", flattenIdentityPlatformConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}
	if err := d.Set("autodelete_anonymous_users", flattenIdentityPlatformConfigAutodeleteAnonymousUsers(res["autodeleteAnonymousUsers"], d, config)); err != nil {
		return fmt.Errorf("Error reading Config: %s", err)
	}

	return nil
}

func resourceIdentityPlatformConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Config: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	autodeleteAnonymousUsersProp, err := expandIdentityPlatformConfigAutodeleteAnonymousUsers(d.Get("autodelete_anonymous_users"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("autodelete_anonymous_users"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, autodeleteAnonymousUsersProp)) {
		obj["autodeleteAnonymousUsers"] = autodeleteAnonymousUsersProp
	}

	url, err := ReplaceVars(d, config, "{{IdentityPlatformBasePath}}projects/{{project}}/config")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Config %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("autodelete_anonymous_users") {
		updateMask = append(updateMask, "autodeleteAnonymousUsers")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Config %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Config %q: %#v", d.Id(), res)
	}

	return resourceIdentityPlatformConfigRead(d, meta)
}

func resourceIdentityPlatformConfigDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] IdentityPlatform Config resources"+
		" cannot be deleted from Google Cloud. The resource %s will be removed from Terraform"+
		" state, but will still be present on Google Cloud.", d.Id())
	d.SetId("")

	return nil
}

func resourceIdentityPlatformConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/config",
		"projects/(?P<project>[^/]+)",
		"(?P<project>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "projects/{{project}}/config")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIdentityPlatformConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIdentityPlatformConfigAutodeleteAnonymousUsers(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIdentityPlatformConfigAutodeleteAnonymousUsers(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
