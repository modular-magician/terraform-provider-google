// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"context"
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

func ResourcePubsubSchemaV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The ID to use for the schema, which will become the final component of the schema's resource name.`,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateEnum([]string{"TYPE_UNSPECIFIED", "PROTOCOL_BUFFER", "AVRO"}),
				Description:  `The type of the schema definition Possible values: ["TYPE_UNSPECIFIED", "PROTOCOL_BUFFER", "AVRO"]`,
				Default:      "TYPE_UNSPECIFIED",
			},
			"definition": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The definition of the schema.
This should contain a string representing the full definition of the schema
that is a valid schema definition of the type specified in type.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func ResourcePubsubSchema() *schema.Resource {
	return &schema.Resource{
		Create: resourcePubsubSchemaCreate,
		Read:   resourcePubsubSchemaRead,
		Update: resourcePubsubSchemaUpdate,
		Delete: resourcePubsubSchemaDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			resourcePubsubSchemaCustomizeDiff,
		),

		Importer: &schema.ResourceImporter{
			State: resourcePubsubSchemaImport,
		},

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    ResourcePubsubSchemaV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourcePubsubSchemaUpgradeV0,
				Version: 0,
			},
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The ID to use for the schema, which will become the final component of the schema's resource name.`,
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validateEnum([]string{"TYPE_UNSPECIFIED", "PROTOCOL_BUFFER", "AVRO"}),
				Description:  `The type of the schema definition Possible values: ["TYPE_UNSPECIFIED", "PROTOCOL_BUFFER", "AVRO"]`,
				Default:      "TYPE_UNSPECIFIED",
			},
			"definition": {
				Type:       schema.TypeString,
				Optional:   true,
				ForceNew:   true,
				Deprecated: `Deprecated in favor of revision.`,
				Description: `The definition of the schema.
This should contain a string representing the full definition of the schema
that is a valid schema definition of the type specified in type.`,
			},
			"revision": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The schema's revisions.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"definition": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The definition of the schema.
This should contain a string representing the full definition of the
schema that is a valid schema definition of the type specified in
type.`,
						},
						"revision_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The revision ID of the schema.`,
						},
						"revision_create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The timestamp that the revision was created.`,
						},
					},
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"use_definition": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourcePubsubSchemaCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["type"] = d.Get("type")

	def := d.Get("definition").(string)
	l := d.Get("revision").([]interface{})
	if len(def) > 0 {
		obj["definition"] = def
	} else {
		if len(l) == 0 {
			return fmt.Errorf("%q must have at least one \"revision\".", d.Id())
		}
		o := l[0].(map[string]interface{})
		obj["definition"] = o["definition"]
		l = l[1:]
	}

	nameProp, err := expandPubsubSchemaName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else {
		obj["name"] = nameProp
	}

	rObj := make(map[string]interface{})
	rObj["name"] = obj["name"]
	rObj["schema"] = obj

	url, err := tpgresource.ReplaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/schemas?schemaId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Schema: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Schema: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	d.Partial(true)
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
		return fmt.Errorf("Error creating Schema: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/schemas/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Schema %q: %#v", d.Id(), res)

	rObj["name"] = d.Id()
	obj["name"] = d.Id()
	for _, raw := range l {
		rev := raw.(map[string]interface{})
		obj["definition"] = rev["definition"]
		rObj["schema"] = obj
		err = commitPubsubSchemaRevision(rObj, d, config, userAgent, billingProject)
		if err != nil {
			return err
		}
	}
	d.Partial(false)

	return resourcePubsubSchemaRead(d, meta)
}

func resourcePubsubSchemaPollRead(d *schema.ResourceData, meta interface{}) transport_tpg.PollReadFunc {
	return func() (map[string]interface{}, error) {
		config := meta.(*transport_tpg.Config)

		url, err := tpgresource.ReplaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/schemas/{{name}}")
		if err != nil {
			return nil, err
		}

		billingProject := ""

		project, err := tpgresource.GetProject(d, config)
		if err != nil {
			return nil, fmt.Errorf("Error fetching project for Schema: %s", err)
		}
		billingProject = project

		// err == nil indicates that the billing_project value was found
		if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
			billingProject = bp
		}

		userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
		if err != nil {
			return nil, err
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "GET",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
		})
		if err != nil {
			return res, err
		}
		return res, nil
	}
}

func resourcePubsubSchemaRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Schema: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/schemas/{{name}}:listRevisions?view=FULL")
	if err != nil {
		return err
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("PubsubSchema %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Schema: %s", err)
	}

	if err != nil {
		return err
	}
	ss := res["schemas"].([]interface{})
	var tss []interface{}
	df := false
	// listRevisions returns schemas in reverse chronological order.
	for i := len(ss) - 1; i >= 0; i-- {
		rs := ss[i].(map[string]interface{})
		if d.Get("use_definition").(bool) && !df {
			// The oldest revision should be the schema-wide "definition" if
			// "definition" is used.
			d.Set("definition", rs["definition"])
			df = true
		} else {
			ts := make(map[string]interface{})
			ts["definition"] = rs["definition"]
			ts["revision_id"] = rs["revisionId"]
			ts["revision_create_time"] = rs["revisionCreateTime"]
			tss = append(tss, ts)
		}
		d.Set("name", NameFromSelfLinkStateFunc(strings.Split(rs["name"].(string), "@")[0]))
		d.Set("type", rs["type"])
	}
	d.Set("revision", tss)
	return nil
}

func resourcePubsubSchemaUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Schema: %s", err)
	}
	billingProject = project

	rObj := make(map[string]interface{})
	rObj["name"] = d.Id()
	obj := make(map[string]interface{})
	typeProp := d.Get("type")
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	obj["name"] = d.Id()

	old, new := d.GetChange("revision")
	lold := old.([]interface{})
	lnew := new.([]interface{})

	adds, dels := computeRevisionDiffs(lold, lnew)

	oldr := make(map[string]bool)
	for _, r := range lold {
		rm := r.(map[string]interface{})
		oldr[rm["revision_id"].(string)] = true
	}

	revC := len(lold)
	if d.Get("definition") != "" {
		revC++
	}

	d.Partial(true)
	dObj := make(map[string]interface{})
	dObj["name"] = d.Id()

	dCount := 0
	aCount := 0

	// Alternate deletes and commmits of revisions in order to keep the number of
	// revisions within the allowed bounds. We initially do these in the order
	// delete/add, but if we are going to delete the last revision, we switch to
	// add/delete.
	for dCount < len(dels) || aCount < len(adds) {
		if dCount < len(dels) && revC > 1 {
			deletePubsubSchemaRevision(dels[dCount], dObj, d, config, userAgent, billingProject)
			dCount++
			revC--
		}
		if aCount < len(adds) {
			obj["definition"] = adds[aCount]
			rObj["schema"] = obj
			err = commitPubsubSchemaRevision(rObj, d, config, userAgent, billingProject)
			if err != nil {
				return err
			}
			aCount++
			revC++
		}
	}

	d.Partial(false)

	return resourcePubsubSchemaRead(d, meta)
}

func resourcePubsubSchemaCustomizeDiff(_ context.Context, d *schema.ResourceDiff, meta interface{}) error {
	rc := 0
	if len(d.Get("definition").(string)) > 0 {
		rc += 1
		d.SetNew("use_definition", true)
	}

	if d.HasChange("revision") {
		_, new := d.GetChange("revision")
		lnew := new.([]interface{})
		rc += len(lnew)
	} else {
		rs := d.Get("revision")
		rl := rs.([]interface{})
		rc += len(rl)
	}
	if rc > 20 {
		return fmt.Errorf("Cannot have more than 20 Schema revisions.")
	}

	if rc == 0 {
		return fmt.Errorf("Must have at least one Schema revision.")
	}

	return nil
}

func commitPubsubSchemaRevision(s map[string]interface{}, d *schema.ResourceData, config *transport_tpg.Config, userAgent string, billingProject string) error {
	url, err := tpgresource.ReplaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/schemas/{{name}}:commit")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Committing Schema revision %q: %#v", d.Id(), s)

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      s,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error committing Schema revision %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished committing Schema revision %q: %#v", d.Id(), res)
	}

	return nil
}

func deletePubsubSchemaRevision(rid string, s map[string]interface{}, d *schema.ResourceData, config *transport_tpg.Config, userAgent string, billingProject string) error {
	url, err := tpgresource.ReplaceVars(d, config, fmt.Sprintf("{{PubsubBasePath}}projects/{{project}}/schemas/{{name}}@%s:deleteRevision", rid))
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting Schema revision %q: %#v", d.Id(), rid)

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      s,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})

	if err != nil {
		return fmt.Errorf("Error deleting Schema revision %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished deleting Schema revision %q: %#v", d.Id(), res)
	}

	return nil
}

func computeRevisionDiffs(old []interface{}, new []interface{}) ([]string, []string) {
	del := []string{}
	add := []string{}
	oldi, newi := 0, 0
	for oldi < len(old) || newi < len(new) {
		if newi == len(new) {
			r := old[oldi].(map[string]interface{})
			del = append(del, r["revision_id"].(string))
			oldi++
		} else if oldi == len(old) {
			r := new[newi].(map[string]interface{})
			add = append(add, r["definition"].(string))
			newi++
		} else {
			rold := old[oldi].(map[string]interface{})
			rnew := new[newi].(map[string]interface{})
			if rold["definition"] == rnew["definition"] {
				oldi++
				newi++
			} else {
				del = append(del, rold["revision_id"].(string))
				oldi++
			}
		}
	}
	return add, del
}

func resourcePubsubSchemaDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Schema: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{PubsubBasePath}}projects/{{project}}/schemas/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting Schema %q", d.Id())

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
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Schema")
	}

	err = PollingWaitTime(resourcePubsubSchemaPollRead(d, meta), PollCheckForAbsence, "Deleting Schema", d.Timeout(schema.TimeoutCreate), 10)
	if err != nil {
		return fmt.Errorf("Error waiting to delete Schema: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Schema %q: %#v", d.Id(), res)
	return nil
}

func resourcePubsubSchemaImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/schemas/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/schemas/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.Set("use_definition", true)
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenPubsubSchemaType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenPubsubSchemaDefinition(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenPubsubSchemaName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func expandPubsubSchemaName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return GetResourceNameFromSelfLink(v.(string)), nil
}

func resourcePubsubSchemaUpgradeV0(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	rawState["use_definition"] = true
	return rawState, nil
}
