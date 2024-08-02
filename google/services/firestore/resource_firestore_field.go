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

package firestore

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceFirestoreField() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirestoreFieldCreate,
		Read:   resourceFirestoreFieldRead,
		Update: resourceFirestoreFieldUpdate,
		Delete: resourceFirestoreFieldDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirestoreFieldImport,
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
			"collection": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The id of the collection group to configure.`,
			},
			"field": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The id of the field to configure.`,
			},
			"database": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The Firestore database id. Defaults to '"(default)"'.`,
				Default:     "(default)",
			},
			"index_config": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `The single field index configuration for this field.
Creating an index configuration for this field will override any inherited configuration with the
indexes specified. Configuring the index configuration with an empty block disables all indexes on
the field.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"indexes": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: `The indexes to configure on the field. Order or array contains must be specified.`,
							Elem:        firestoreFieldIndexConfigIndexesSchema(),
							// Default schema.HashSchema is used.
						},
					},
				},
			},
			"ttl_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The TTL configuration for this Field. If set to an empty block (i.e. 'ttl_config {}'), a TTL policy is configured based on the field. If unset, a TTL policy is not configured (or will be disabled upon updating the resource).`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The state of TTL (time-to-live) configuration for documents that have this Field set.`,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The name of this field. Format:
'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/fields/{{field}}'`,
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

func firestoreFieldIndexConfigIndexesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"array_config": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"CONTAINS", ""}),
				Description: `Indicates that this field supports operations on arrayValues. Only one of 'order' and 'arrayConfig' can
be specified. Possible values: ["CONTAINS"]`,
				ExactlyOneOf: []string{},
			},
			"order": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ASCENDING", "DESCENDING", ""}),
				Description: `Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=, !=.
Only one of 'order' and 'arrayConfig' can be specified. Possible values: ["ASCENDING", "DESCENDING"]`,
				ExactlyOneOf: []string{},
			},
			"query_scope": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"COLLECTION", "COLLECTION_GROUP", ""}),
				Description: `The scope at which a query is run. Collection scoped queries require you specify
the collection at query time. Collection group scope allows queries across all
collections with the same id. Default value: "COLLECTION" Possible values: ["COLLECTION", "COLLECTION_GROUP"]`,
				Default: "COLLECTION",
			},
		},
	}
}

func resourceFirestoreFieldCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	indexConfigProp, err := expandFirestoreFieldIndexConfig(d.Get("index_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("index_config"); ok || !reflect.DeepEqual(v, indexConfigProp) {
		obj["indexConfig"] = indexConfigProp
	}
	ttlConfigProp, err := expandFirestoreFieldTtlConfig(d.Get("ttl_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ttl_config"); ok || !reflect.DeepEqual(v, ttlConfigProp) {
		obj["ttlConfig"] = ttlConfigProp
	}

	obj, err = resourceFirestoreFieldEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirestoreBasePath}}projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/fields/{{field}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Field: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Field: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "PATCH",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutCreate),
		Headers:              headers,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.FirestoreField409RetryUnderlyingDataChanged},
	})
	if err != nil {
		return fmt.Errorf("Error creating Field: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = FirestoreOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Field", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Field: %s", err)
	}

	if err := d.Set("name", flattenFirestoreFieldName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Field %q: %#v", d.Id(), res)

	return resourceFirestoreFieldRead(d, meta)
}

func resourceFirestoreFieldRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirestoreBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Field: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "GET",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Headers:              headers,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.FirestoreField409RetryUnderlyingDataChanged},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("FirestoreField %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Field: %s", err)
	}

	if err := d.Set("name", flattenFirestoreFieldName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Field: %s", err)
	}
	if err := d.Set("index_config", flattenFirestoreFieldIndexConfig(res["indexConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Field: %s", err)
	}
	if err := d.Set("ttl_config", flattenFirestoreFieldTtlConfig(res["ttlConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Field: %s", err)
	}

	return nil
}

func resourceFirestoreFieldUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Field: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	indexConfigProp, err := expandFirestoreFieldIndexConfig(d.Get("index_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("index_config"); ok || !reflect.DeepEqual(v, indexConfigProp) {
		obj["indexConfig"] = indexConfigProp
	}
	ttlConfigProp, err := expandFirestoreFieldTtlConfig(d.Get("ttl_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ttl_config"); ok || !reflect.DeepEqual(v, ttlConfigProp) {
		obj["ttlConfig"] = ttlConfigProp
	}

	obj, err = resourceFirestoreFieldEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirestoreBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Field %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("index_config") {
		updateMask = append(updateMask, "indexConfig")
	}

	if d.HasChange("ttl_config") {
		updateMask = append(updateMask, "ttlConfig")
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
			Config:               config,
			Method:               "PATCH",
			Project:              billingProject,
			RawURL:               url,
			UserAgent:            userAgent,
			Body:                 obj,
			Timeout:              d.Timeout(schema.TimeoutUpdate),
			Headers:              headers,
			ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.FirestoreField409RetryUnderlyingDataChanged},
		})

		if err != nil {
			return fmt.Errorf("Error updating Field %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Field %q: %#v", d.Id(), res)
		}

		err = FirestoreOperationWaitTime(
			config, res, project, "Updating Field", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceFirestoreFieldRead(d, meta)
}

func resourceFirestoreFieldDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	// Firestore fields cannot be deleted, instead we clear the indexConfig and ttlConfig.

	log.Printf("[DEBUG] Deleting Field %q", d.Id())

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for App: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{FirestoreBasePath}}{{name}}")
	if err != nil {
		return err
	}

	updateMask := []string{"indexConfig", "ttlConfig"}

	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// Clear fields by sending an empty PATCH request with appropriate update mask.
	req := make(map[string]interface{})
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      req,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error deleting Field %q: %s", d.Id(), err)
	}

	err = FirestoreOperationWaitTime(
		config, res, project, "Deleting Field", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Field %q", d.Id())
	return nil
}

func resourceFirestoreFieldImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	// Re-populate split fields from the name.
	re := regexp.MustCompile("^projects/([^/]+)/databases/([^/]+)/collectionGroups/([^/]+)/fields/(.+)$")
	match := re.FindStringSubmatch(d.Get("name").(string))
	if len(match) > 0 {
		if err := d.Set("project", match[1]); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
		if err := d.Set("database", match[2]); err != nil {
			return nil, fmt.Errorf("Error setting database: %s", err)
		}
		if err := d.Set("collection", match[3]); err != nil {
			return nil, fmt.Errorf("Error setting collection: %s", err)
		}
		if err := d.Set("field", match[4]); err != nil {
			return nil, fmt.Errorf("Error setting field: %s", err)
		}
	} else {
		return nil, fmt.Errorf("import did not match the regex ^projects/([^/]+)/databases/([^/]+)/collectionGroups/([^/]+)/fields/(.+)$")
	}

	return []*schema.ResourceData{d}, nil
}

func flattenFirestoreFieldName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreFieldIndexConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	indexConfig := v.(map[string]interface{})

	usesAncestorConfig := false
	if indexConfig["usesAncestorConfig"] != nil {
		usesAncestorConfig = indexConfig["usesAncestorConfig"].(bool)
	}

	if usesAncestorConfig {
		// The intent when uses_ancestor_config is no config.
		return []interface{}{}
	}

	if indexConfig["indexes"] == nil {
		// No indexes, return an existing, but empty index config.
		return [1]interface{}{nil}
	}

	// For Single field indexes, we put the field configuration on the index to avoid forced nesting.
	l := indexConfig["indexes"].([]interface{})
	transformed := make(map[string]interface{})
	transformedIndexes := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		fields := original["fields"].([]interface{})
		sfi := fields[0].(map[string]interface{})
		transformedIndexes = append(transformedIndexes, map[string]interface{}{
			"query_scope":  original["queryScope"],
			"order":        sfi["order"],
			"array_config": sfi["arrayConfig"],
		})
	}
	transformed["indexes"] = transformedIndexes
	return []interface{}{transformed}
}

func flattenFirestoreFieldTtlConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["state"] =
		flattenFirestoreFieldTtlConfigState(original["state"], d, config)
	return []interface{}{transformed}
}
func flattenFirestoreFieldTtlConfigState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandFirestoreFieldIndexConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	// We drop all output only fields as they are unnecessary.
	if v == nil {
		return nil, nil
	}

	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	transformedIndexConfig := make(map[string]interface{})

	// A configured, but empty, index_config block should be sent. This is how a user would remove all indexes.
	if l[0] == nil {
		return transformedIndexConfig, nil
	}

	indexConfig := l[0].(map[string]interface{})

	// For Single field indexes, we put the field configuration on the index to avoid forced nesting.
	// Push all order/arrayConfig down into a single element fields list.
	l = indexConfig["indexes"].(*schema.Set).List()
	transformedIndexes := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})
		transformedField := make(map[string]interface{})

		if val := reflect.ValueOf(original["query_scope"]); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["queryScope"] = original["query_scope"]
		}

		if val := reflect.ValueOf(original["order"]); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformedField["order"] = original["order"]
		}

		if val := reflect.ValueOf(original["array_config"]); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformedField["arrayConfig"] = original["array_config"]
		}
		transformed["fields"] = [1]interface{}{
			transformedField,
		}

		transformedIndexes = append(transformedIndexes, transformed)
	}
	transformedIndexConfig["indexes"] = transformedIndexes
	return transformedIndexConfig, nil
}

/*
 * Expands an empty terraform config into an empty object.
 *
 * Used to differentate a user specifying an empty block versus a null/unset block.
 *
 * This is unique from send_empty_value, which will send an explicit null value
 * for empty configuration blocks.
 */
func expandFirestoreFieldTtlConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}
	// A set, but empty object.
	return struct{}{}, nil
}

func resourceFirestoreFieldEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {

	// We've added project / database / collection / field as split fields of the name, but
	// the API doesn't expect them.  Make sure we remove them from any requests.

	delete(obj, "project")
	delete(obj, "database")
	delete(obj, "collection")
	delete(obj, "field")
	return obj, nil
}
