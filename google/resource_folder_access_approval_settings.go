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
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var accessApprovalCloudProductMapping = map[string]string{
	"appengine.googleapis.com": "App Engine",
	"bigquery.googleapis.com":  "BigQuery",
	"bigtable.googleapis.com":  "Cloud Bigtable",
	"cloudkms.googleapis.com":  "Cloud Key Management Service",
	"compute.googleapis.com":   "Compute Engine",
	"dataflow.googleapis.com":  "Cloud Dataflow",
	"iam.googleapis.com":       "Cloud Identity and Access Management",
	"pubsub.googleapis.com":    "Cloud Pub/Sub",
	"storage.googleapis.com":   "Cloud Storage",
}

func accessApprovalEnrolledServicesHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	cp := m["cloud_product"].(string)
	if n, ok := accessApprovalCloudProductMapping[cp]; ok {
		cp = n
	}
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(cp))) // ToLower just in case
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["enrollment_level"].(string))))
	return hashcode(buf.String())
}

func ResourceAccessApprovalFolderSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessApprovalFolderSettingsCreate,
		Read:   resourceAccessApprovalFolderSettingsRead,
		Update: resourceAccessApprovalFolderSettingsUpdate,
		Delete: resourceAccessApprovalFolderSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessApprovalFolderSettingsImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"enrolled_services": {
				Type:     schema.TypeSet,
				Required: true,
				Description: `A list of Google Cloud Services for which the given resource has Access Approval enrolled.
Access requests for the resource given by name against any of these services contained here will be required
to have explicit approval. Enrollment can only be done on an all or nothing basis.

A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.`,
				Elem: accessapprovalFolderSettingsEnrolledServicesSchema(),
				Set:  accessApprovalEnrolledServicesHash,
			},
			"folder_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `ID of the folder of the access approval settings.`,
			},
			"active_key_version": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The asymmetric crypto key version to use for signing approval requests.
Empty active_key_version indicates that a Google-managed key should be used for signing.
This property will be ignored if set by an ancestor of the resource, and new non-empty values may not be set.`,
			},
			"notification_emails": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Description: `A list of email addresses to which notifications relating to approval requests should be sent.
Notifications relating to a resource will be sent to all emails in the settings of ancestor
resources of that resource. A maximum of 50 email addresses are allowed.`,
				MaxItems: 50,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"ancestor_has_active_key_version": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `If the field is true, that indicates that an ancestor of this Folder has set active_key_version.`,
			},
			"enrolled_ancestor": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: `If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors of the Folder.`,
			},
			"invalid_key_version": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: `If the field is true, that indicates that there is some configuration issue with the active_key_version
configured on this Folder (e.g. it doesn't exist or the Access Approval service account doesn't have the
correct permissions on it, etc.) This key version is not necessarily the effective key version at this level,
as key versions are inherited top-down.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the settings. Format is "folders/{folder_id}/accessApprovalSettings"`,
			},
		},
		UseJSONNumber: true,
	}
}

func accessapprovalFolderSettingsEnrolledServicesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_product": {
				Type:     schema.TypeString,
				Required: true,
				Description: `The product for which Access Approval will be enrolled. Allowed values are listed (case-sensitive):
  * all
  * App Engine
  * BigQuery
  * Cloud Bigtable
  * Cloud Key Management Service
  * Compute Engine
  * Cloud Dataflow
  * Cloud Identity and Access Management
  * Cloud Pub/Sub
  * Cloud Storage
  * Persistent Disk

Note: These values are supported as input, but considered a legacy format:
  * all
  * appengine.googleapis.com
  * bigquery.googleapis.com
  * bigtable.googleapis.com
  * cloudkms.googleapis.com
  * compute.googleapis.com
  * dataflow.googleapis.com
  * iam.googleapis.com
  * pubsub.googleapis.com
  * storage.googleapis.com`,
			},
			"enrollment_level": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateEnum([]string{"BLOCK_ALL", ""}),
				Description:  `The enrollment level of the service. Default value: "BLOCK_ALL" Possible values: ["BLOCK_ALL"]`,
				Default:      "BLOCK_ALL",
			},
		},
	}
}

func resourceAccessApprovalFolderSettingsCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	notificationEmailsProp, err := expandAccessApprovalFolderSettingsNotificationEmails(d.Get("notification_emails"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_emails"); !isEmptyValue(reflect.ValueOf(notificationEmailsProp)) && (ok || !reflect.DeepEqual(v, notificationEmailsProp)) {
		obj["notificationEmails"] = notificationEmailsProp
	}
	enrolledServicesProp, err := expandAccessApprovalFolderSettingsEnrolledServices(d.Get("enrolled_services"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enrolled_services"); !isEmptyValue(reflect.ValueOf(enrolledServicesProp)) && (ok || !reflect.DeepEqual(v, enrolledServicesProp)) {
		obj["enrolledServices"] = enrolledServicesProp
	}
	activeKeyVersionProp, err := expandAccessApprovalFolderSettingsActiveKeyVersion(d.Get("active_key_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("active_key_version"); !isEmptyValue(reflect.ValueOf(activeKeyVersionProp)) && (ok || !reflect.DeepEqual(v, activeKeyVersionProp)) {
		obj["activeKeyVersion"] = activeKeyVersionProp
	}

	url, err := replaceVars(d, config, "{{AccessApprovalBasePath}}folders/{{folder_id}}/accessApprovalSettings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new FolderSettings: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	updateMask := []string{}

	if d.HasChange("notification_emails") {
		updateMask = append(updateMask, "notificationEmails")
	}

	if d.HasChange("enrolled_services") {
		updateMask = append(updateMask, "enrolledServices")
	}

	if d.HasChange("active_key_version") {
		updateMask = append(updateMask, "activeKeyVersion")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating FolderSettings: %s", err)
	}
	if err := d.Set("name", flattenAccessApprovalFolderSettingsName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "folders/{{folder_id}}/accessApprovalSettings")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating FolderSettings %q: %#v", d.Id(), res)

	return resourceAccessApprovalFolderSettingsRead(d, meta)
}

func resourceAccessApprovalFolderSettingsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{AccessApprovalBasePath}}folders/{{folder_id}}/accessApprovalSettings")
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
		return handleNotFoundError(err, d, fmt.Sprintf("AccessApprovalFolderSettings %q", d.Id()))
	}

	if err := d.Set("name", flattenAccessApprovalFolderSettingsName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderSettings: %s", err)
	}
	if err := d.Set("notification_emails", flattenAccessApprovalFolderSettingsNotificationEmails(res["notificationEmails"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderSettings: %s", err)
	}
	if err := d.Set("enrolled_services", flattenAccessApprovalFolderSettingsEnrolledServices(res["enrolledServices"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderSettings: %s", err)
	}
	if err := d.Set("enrolled_ancestor", flattenAccessApprovalFolderSettingsEnrolledAncestor(res["enrolledAncestor"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderSettings: %s", err)
	}
	if err := d.Set("active_key_version", flattenAccessApprovalFolderSettingsActiveKeyVersion(res["activeKeyVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderSettings: %s", err)
	}
	if err := d.Set("ancestor_has_active_key_version", flattenAccessApprovalFolderSettingsAncestorHasActiveKeyVersion(res["ancestorHasActiveKeyVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderSettings: %s", err)
	}
	if err := d.Set("invalid_key_version", flattenAccessApprovalFolderSettingsInvalidKeyVersion(res["invalidKeyVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading FolderSettings: %s", err)
	}

	return nil
}

func resourceAccessApprovalFolderSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	notificationEmailsProp, err := expandAccessApprovalFolderSettingsNotificationEmails(d.Get("notification_emails"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_emails"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, notificationEmailsProp)) {
		obj["notificationEmails"] = notificationEmailsProp
	}
	enrolledServicesProp, err := expandAccessApprovalFolderSettingsEnrolledServices(d.Get("enrolled_services"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enrolled_services"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, enrolledServicesProp)) {
		obj["enrolledServices"] = enrolledServicesProp
	}
	activeKeyVersionProp, err := expandAccessApprovalFolderSettingsActiveKeyVersion(d.Get("active_key_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("active_key_version"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, activeKeyVersionProp)) {
		obj["activeKeyVersion"] = activeKeyVersionProp
	}

	url, err := replaceVars(d, config, "{{AccessApprovalBasePath}}folders/{{folder_id}}/accessApprovalSettings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating FolderSettings %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("notification_emails") {
		updateMask = append(updateMask, "notificationEmails")
	}

	if d.HasChange("enrolled_services") {
		updateMask = append(updateMask, "enrolledServices")
	}

	if d.HasChange("active_key_version") {
		updateMask = append(updateMask, "activeKeyVersion")
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
		return fmt.Errorf("Error updating FolderSettings %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating FolderSettings %q: %#v", d.Id(), res)
	}

	return resourceAccessApprovalFolderSettingsRead(d, meta)
}

func resourceAccessApprovalFolderSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["notificationEmails"] = []string{}
	obj["enrolledServices"] = []string{}
	obj["activeKeyVersion"] = ""

	url, err := replaceVars(d, config, "{{AccessApprovalBasePath}}folders/{{folder_id}}/accessApprovalSettings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Emptying FolderSettings %q: %#v", d.Id(), obj)
	updateMask := []string{}

	updateMask = append(updateMask, "notificationEmails")
	updateMask = append(updateMask, "enrolledServices")
	updateMask = append(updateMask, "activeKeyVersion")

	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	res, err := sendRequestWithTimeout(config, "PATCH", "", url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error emptying FolderSettings %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished emptying FolderSettings %q: %#v", d.Id(), res)
	}

	return nil
}

func resourceAccessApprovalFolderSettingsImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"folders/(?P<folder_id>[^/]+)/accessApprovalSettings",
		"(?P<folder_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "folders/{{folder_id}}/accessApprovalSettings")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAccessApprovalFolderSettingsName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalFolderSettingsNotificationEmails(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenAccessApprovalFolderSettingsEnrolledServices(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(accessApprovalEnrolledServicesHash, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"cloud_product":    flattenAccessApprovalFolderSettingsEnrolledServicesCloudProduct(original["cloudProduct"], d, config),
			"enrollment_level": flattenAccessApprovalFolderSettingsEnrolledServicesEnrollmentLevel(original["enrollmentLevel"], d, config),
		})
	}
	return transformed
}
func flattenAccessApprovalFolderSettingsEnrolledServicesCloudProduct(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalFolderSettingsEnrolledServicesEnrollmentLevel(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalFolderSettingsEnrolledAncestor(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalFolderSettingsActiveKeyVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalFolderSettingsAncestorHasActiveKeyVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessApprovalFolderSettingsInvalidKeyVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandAccessApprovalFolderSettingsNotificationEmails(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandAccessApprovalFolderSettingsEnrolledServices(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCloudProduct, err := expandAccessApprovalFolderSettingsEnrolledServicesCloudProduct(original["cloud_product"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCloudProduct); val.IsValid() && !isEmptyValue(val) {
			transformed["cloudProduct"] = transformedCloudProduct
		}

		transformedEnrollmentLevel, err := expandAccessApprovalFolderSettingsEnrolledServicesEnrollmentLevel(original["enrollment_level"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEnrollmentLevel); val.IsValid() && !isEmptyValue(val) {
			transformed["enrollmentLevel"] = transformedEnrollmentLevel
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAccessApprovalFolderSettingsEnrolledServicesCloudProduct(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessApprovalFolderSettingsEnrolledServicesEnrollmentLevel(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessApprovalFolderSettingsActiveKeyVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
