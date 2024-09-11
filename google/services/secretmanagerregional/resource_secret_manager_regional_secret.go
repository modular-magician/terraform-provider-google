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

package secretmanagerregional

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceSecretManagerRegionalRegionalSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretManagerRegionalRegionalSecretCreate,
		Read:   resourceSecretManagerRegionalRegionalSecretRead,
		Update: resourceSecretManagerRegionalRegionalSecretUpdate,
		Delete: resourceSecretManagerRegionalRegionalSecretDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecretManagerRegionalRegionalSecretImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.SetAnnotationsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location of the regional secret. eg us-central1`,
			},
			"secret_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `This must be unique within the project.`,
			},
			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Custom metadata about the regional secret.

Annotations are distinct from various forms of labels. Annotations exist to allow
client tools to store their own state information without requiring a database.

Annotation keys must be between 1 and 63 characters long, have a UTF-8 encoding of
maximum 128 bytes, begin and end with an alphanumeric character ([a-z0-9A-Z]), and
may have dashes (-), underscores (_), dots (.), and alphanumerics in between these
symbols.

The total size of annotation keys and values must be less than 16KiB.

An object containing a list of "key": value pairs. Example:
{ "name": "wrench", "mass": "1.3kg", "count": "3" }.


**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
Please refer to the field 'effective_annotations' for all of the annotations present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"customer_managed_encryption": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The customer-managed encryption configuration of the regional secret.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The resource name of the Cloud KMS CryptoKey used to encrypt secret payloads.`,
						},
					},
				},
			},
			"expire_time": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				Description: `Timestamp in UTC when the regional secret is scheduled to expire. This is always provided on
output, regardless of what was sent on input. A timestamp in RFC3339 UTC "Zulu" format, with
nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and
"2014-10-02T15:01:23.045123456Z". Only one of 'expire_time' or 'ttl' can be provided.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `The labels assigned to this regional secret.

Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}

Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}

No more than 64 labels can be assigned to a given resource.

An object containing a list of "key": value pairs. Example:
{ "name": "wrench", "mass": "1.3kg", "count": "3" }.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"rotation": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `The rotation time and period for a regional secret. At 'next_rotation_time', Secret Manager
will send a Pub/Sub notification to the topics configured on the Secret. 'topics' must be
set to configure rotation.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"next_rotation_time": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Timestamp in UTC at which the Secret is scheduled to rotate.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
						},
						"rotation_period": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The Duration between rotation notifications. Must be in seconds and at least 3600s (1h)
and at most 3153600000s (100 years). If rotationPeriod is set, 'next_rotation_time' must
be set. 'next_rotation_time' will be advanced by this period when the service
automatically sends rotation notifications.`,
							RequiredWith: []string{"rotation.0.next_rotation_time"},
						},
					},
				},
				RequiredWith: []string{"topics"},
			},
			"topics": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A list of up to 10 Pub/Sub topics to which messages are published when control plane
operations are called on the regional secret or its versions.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The resource name of the Pub/Sub topic that will be published to, in the following
format: projects/*/topics/*. For publication to succeed, the Secret Manager Service
Agent service account must have pubsub.publisher permissions on the topic.`,
						},
					},
				},
			},
			"ttl": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `The TTL for the regional secret. A duration in seconds with up to nine fractional digits,
terminated by 's'. Example: "3.5s". Only one of 'ttl' or 'expire_time' can be provided.`,
			},
			"version_destroy_ttl": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Secret Version TTL after destruction request.
This is a part of the delayed delete feature on Secret Version.
For secret with versionDestroyTtl>0, version destruction doesn't happen immediately
on calling destroy instead the version goes to a disabled state and
the actual destruction happens after this TTL expires. It must be atleast 24h.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time at which the regional secret was created.`,
			},
			"effective_annotations": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the regional secret. Format:
'projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}'`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceSecretManagerRegionalRegionalSecretCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	customerManagedEncryptionProp, err := expandSecretManagerRegionalRegionalSecretCustomerManagedEncryption(d.Get("customer_managed_encryption"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("customer_managed_encryption"); !tpgresource.IsEmptyValue(reflect.ValueOf(customerManagedEncryptionProp)) && (ok || !reflect.DeepEqual(v, customerManagedEncryptionProp)) {
		obj["customerManagedEncryption"] = customerManagedEncryptionProp
	}
	topicsProp, err := expandSecretManagerRegionalRegionalSecretTopics(d.Get("topics"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("topics"); !tpgresource.IsEmptyValue(reflect.ValueOf(topicsProp)) && (ok || !reflect.DeepEqual(v, topicsProp)) {
		obj["topics"] = topicsProp
	}
	rotationProp, err := expandSecretManagerRegionalRegionalSecretRotation(d.Get("rotation"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rotation"); !tpgresource.IsEmptyValue(reflect.ValueOf(rotationProp)) && (ok || !reflect.DeepEqual(v, rotationProp)) {
		obj["rotation"] = rotationProp
	}
	expireTimeProp, err := expandSecretManagerRegionalRegionalSecretExpireTime(d.Get("expire_time"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expire_time"); !tpgresource.IsEmptyValue(reflect.ValueOf(expireTimeProp)) && (ok || !reflect.DeepEqual(v, expireTimeProp)) {
		obj["expireTime"] = expireTimeProp
	}
	ttlProp, err := expandSecretManagerRegionalRegionalSecretTtl(d.Get("ttl"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(ttlProp)) && (ok || !reflect.DeepEqual(v, ttlProp)) {
		obj["ttl"] = ttlProp
	}
	versionDestroyTtlProp, err := expandSecretManagerRegionalRegionalSecretVersionDestroyTtl(d.Get("version_destroy_ttl"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("version_destroy_ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(versionDestroyTtlProp)) && (ok || !reflect.DeepEqual(v, versionDestroyTtlProp)) {
		obj["versionDestroyTtl"] = versionDestroyTtlProp
	}
	labelsProp, err := expandSecretManagerRegionalRegionalSecretEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandSecretManagerRegionalRegionalSecretEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerRegionalBasePath}}projects/{{project}}/locations/{{location}}/secrets?secretId={{secret_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionalSecret: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionalSecret: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating RegionalSecret: %s", err)
	}
	if err := d.Set("name", flattenSecretManagerRegionalRegionalSecretName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating RegionalSecret %q: %#v", d.Id(), res)

	return resourceSecretManagerRegionalRegionalSecretRead(d, meta)
}

func resourceSecretManagerRegionalRegionalSecretRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerRegionalBasePath}}projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionalSecret: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("SecretManagerRegionalRegionalSecret %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionalSecret: %s", err)
	}

	if err := d.Set("name", flattenSecretManagerRegionalRegionalSecretName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecret: %s", err)
	}
	if err := d.Set("create_time", flattenSecretManagerRegionalRegionalSecretCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecret: %s", err)
	}
	if err := d.Set("labels", flattenSecretManagerRegionalRegionalSecretLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecret: %s", err)
	}
	if err := d.Set("annotations", flattenSecretManagerRegionalRegionalSecretAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecret: %s", err)
	}
	if err := d.Set("customer_managed_encryption", flattenSecretManagerRegionalRegionalSecretCustomerManagedEncryption(res["customerManagedEncryption"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecret: %s", err)
	}
	if err := d.Set("topics", flattenSecretManagerRegionalRegionalSecretTopics(res["topics"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecret: %s", err)
	}
	if err := d.Set("rotation", flattenSecretManagerRegionalRegionalSecretRotation(res["rotation"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecret: %s", err)
	}
	if err := d.Set("expire_time", flattenSecretManagerRegionalRegionalSecretExpireTime(res["expireTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecret: %s", err)
	}
	if err := d.Set("version_destroy_ttl", flattenSecretManagerRegionalRegionalSecretVersionDestroyTtl(res["versionDestroyTtl"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecret: %s", err)
	}
	if err := d.Set("terraform_labels", flattenSecretManagerRegionalRegionalSecretTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecret: %s", err)
	}
	if err := d.Set("effective_labels", flattenSecretManagerRegionalRegionalSecretEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecret: %s", err)
	}
	if err := d.Set("effective_annotations", flattenSecretManagerRegionalRegionalSecretEffectiveAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionalSecret: %s", err)
	}

	return nil
}

func resourceSecretManagerRegionalRegionalSecretUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionalSecret: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	customerManagedEncryptionProp, err := expandSecretManagerRegionalRegionalSecretCustomerManagedEncryption(d.Get("customer_managed_encryption"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("customer_managed_encryption"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, customerManagedEncryptionProp)) {
		obj["customerManagedEncryption"] = customerManagedEncryptionProp
	}
	topicsProp, err := expandSecretManagerRegionalRegionalSecretTopics(d.Get("topics"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("topics"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, topicsProp)) {
		obj["topics"] = topicsProp
	}
	rotationProp, err := expandSecretManagerRegionalRegionalSecretRotation(d.Get("rotation"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rotation"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, rotationProp)) {
		obj["rotation"] = rotationProp
	}
	expireTimeProp, err := expandSecretManagerRegionalRegionalSecretExpireTime(d.Get("expire_time"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("expire_time"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, expireTimeProp)) {
		obj["expireTime"] = expireTimeProp
	}
	ttlProp, err := expandSecretManagerRegionalRegionalSecretTtl(d.Get("ttl"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, ttlProp)) {
		obj["ttl"] = ttlProp
	}
	versionDestroyTtlProp, err := expandSecretManagerRegionalRegionalSecretVersionDestroyTtl(d.Get("version_destroy_ttl"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("version_destroy_ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, versionDestroyTtlProp)) {
		obj["versionDestroyTtl"] = versionDestroyTtlProp
	}
	labelsProp, err := expandSecretManagerRegionalRegionalSecretEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandSecretManagerRegionalRegionalSecretEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerRegionalBasePath}}projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RegionalSecret %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("customer_managed_encryption") {
		updateMask = append(updateMask, "customerManagedEncryption")
	}

	if d.HasChange("topics") {
		updateMask = append(updateMask, "topics")
	}

	if d.HasChange("rotation") {
		updateMask = append(updateMask, "rotation")
	}

	if d.HasChange("expire_time") {
		updateMask = append(updateMask, "expireTime")
	}

	if d.HasChange("ttl") {
		updateMask = append(updateMask, "ttl")
	}

	if d.HasChange("version_destroy_ttl") {
		updateMask = append(updateMask, "versionDestroyTtl")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("effective_annotations") {
		updateMask = append(updateMask, "annotations")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}
	// As the API expects only one of ttl or expireTime
	if d.HasChange("ttl") && !d.HasChange("expire_time") {
		delete(obj, "expireTime")
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
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating RegionalSecret %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating RegionalSecret %q: %#v", d.Id(), res)
		}

	}

	return resourceSecretManagerRegionalRegionalSecretRead(d, meta)
}

func resourceSecretManagerRegionalRegionalSecretDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionalSecret: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{SecretManagerRegionalBasePath}}projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting RegionalSecret %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "RegionalSecret")
	}

	log.Printf("[DEBUG] Finished deleting RegionalSecret %q: %#v", d.Id(), res)
	return nil
}

func resourceSecretManagerRegionalRegionalSecretImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/secrets/(?P<secret_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<secret_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<secret_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenSecretManagerRegionalRegionalSecretName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerRegionalRegionalSecretCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerRegionalRegionalSecretLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenSecretManagerRegionalRegionalSecretAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("annotations"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenSecretManagerRegionalRegionalSecretCustomerManagedEncryption(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_name"] =
		flattenSecretManagerRegionalRegionalSecretCustomerManagedEncryptionKmsKeyName(original["kmsKeyName"], d, config)
	return []interface{}{transformed}
}
func flattenSecretManagerRegionalRegionalSecretCustomerManagedEncryptionKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerRegionalRegionalSecretTopics(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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
			"name": flattenSecretManagerRegionalRegionalSecretTopicsName(original["name"], d, config),
		})
	}
	return transformed
}
func flattenSecretManagerRegionalRegionalSecretTopicsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerRegionalRegionalSecretRotation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["next_rotation_time"] =
		flattenSecretManagerRegionalRegionalSecretRotationNextRotationTime(original["nextRotationTime"], d, config)
	transformed["rotation_period"] =
		flattenSecretManagerRegionalRegionalSecretRotationRotationPeriod(original["rotationPeriod"], d, config)
	return []interface{}{transformed}
}
func flattenSecretManagerRegionalRegionalSecretRotationNextRotationTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerRegionalRegionalSecretRotationRotationPeriod(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerRegionalRegionalSecretExpireTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerRegionalRegionalSecretVersionDestroyTtl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerRegionalRegionalSecretTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenSecretManagerRegionalRegionalSecretEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenSecretManagerRegionalRegionalSecretEffectiveAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandSecretManagerRegionalRegionalSecretCustomerManagedEncryption(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandSecretManagerRegionalRegionalSecretCustomerManagedEncryptionKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandSecretManagerRegionalRegionalSecretCustomerManagedEncryptionKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerRegionalRegionalSecretTopics(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandSecretManagerRegionalRegionalSecretTopicsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandSecretManagerRegionalRegionalSecretTopicsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerRegionalRegionalSecretRotation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNextRotationTime, err := expandSecretManagerRegionalRegionalSecretRotationNextRotationTime(original["next_rotation_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNextRotationTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nextRotationTime"] = transformedNextRotationTime
	}

	transformedRotationPeriod, err := expandSecretManagerRegionalRegionalSecretRotationRotationPeriod(original["rotation_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRotationPeriod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rotationPeriod"] = transformedRotationPeriod
	}

	return transformed, nil
}

func expandSecretManagerRegionalRegionalSecretRotationNextRotationTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerRegionalRegionalSecretRotationRotationPeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerRegionalRegionalSecretExpireTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerRegionalRegionalSecretTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerRegionalRegionalSecretVersionDestroyTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecretManagerRegionalRegionalSecretEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandSecretManagerRegionalRegionalSecretEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
