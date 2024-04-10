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

package netapp

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

func ResourceNetappVolumeSnapshot() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetappVolumeSnapshotCreate,
		Read:   resourceNetappVolumeSnapshotRead,
		Update: resourceNetappVolumeSnapshotUpdate,
		Delete: resourceNetappVolumeSnapshotDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNetappVolumeSnapshotImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Name of the snapshot location. Snapshots are child resources of volumes and live in the same location.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the snapshot.`,
			},
			"volume_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the volume to create the snapshot in.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Description for the snapshot.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"used_bytes": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `Storage used to store blocks unique to this snapshot.`,
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

func resourceNetappVolumeSnapshotCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetappVolumeSnapshotDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetappVolumeSnapshotEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetappBasePath}}projects/{{project}}/locations/{{location}}/volumes/{{volume_name}}/snapshots?snapshotId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new VolumeSnapshot: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for VolumeSnapshot: %s", err)
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
		return fmt.Errorf("Error creating VolumeSnapshot: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/volumes/{{volume_name}}/snapshots/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = NetappOperationWaitTime(
		config, res, project, "Creating VolumeSnapshot", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create VolumeSnapshot: %s", err)
	}

	log.Printf("[DEBUG] Finished creating VolumeSnapshot %q: %#v", d.Id(), res)

	return resourceNetappVolumeSnapshotRead(d, meta)
}

func resourceNetappVolumeSnapshotRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetappBasePath}}projects/{{project}}/locations/{{location}}/volumes/{{volume_name}}/snapshots/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for VolumeSnapshot: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("NetappVolumeSnapshot %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading VolumeSnapshot: %s", err)
	}

	if err := d.Set("description", flattenNetappVolumeSnapshotDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading VolumeSnapshot: %s", err)
	}
	if err := d.Set("labels", flattenNetappVolumeSnapshotLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading VolumeSnapshot: %s", err)
	}
	if err := d.Set("used_bytes", flattenNetappVolumeSnapshotUsedBytes(res["usedBytes"], d, config)); err != nil {
		return fmt.Errorf("Error reading VolumeSnapshot: %s", err)
	}
	if err := d.Set("terraform_labels", flattenNetappVolumeSnapshotTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading VolumeSnapshot: %s", err)
	}
	if err := d.Set("effective_labels", flattenNetappVolumeSnapshotEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading VolumeSnapshot: %s", err)
	}

	return nil
}

func resourceNetappVolumeSnapshotUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for VolumeSnapshot: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandNetappVolumeSnapshotDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandNetappVolumeSnapshotEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{NetappBasePath}}projects/{{project}}/locations/{{location}}/volumes/{{volume_name}}/snapshots/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating VolumeSnapshot %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
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
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating VolumeSnapshot %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating VolumeSnapshot %q: %#v", d.Id(), res)
		}

		err = NetappOperationWaitTime(
			config, res, project, "Updating VolumeSnapshot", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceNetappVolumeSnapshotRead(d, meta)
}

func resourceNetappVolumeSnapshotDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for VolumeSnapshot: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{NetappBasePath}}projects/{{project}}/locations/{{location}}/volumes/{{volume_name}}/snapshots/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting VolumeSnapshot %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "VolumeSnapshot")
	}

	err = NetappOperationWaitTime(
		config, res, project, "Deleting VolumeSnapshot", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting VolumeSnapshot %q: %#v", d.Id(), res)
	return nil
}

func resourceNetappVolumeSnapshotImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/volumes/(?P<volume_name>[^/]+)/snapshots/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<volume_name>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<volume_name>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/volumes/{{volume_name}}/snapshots/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNetappVolumeSnapshotDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenNetappVolumeSnapshotLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetappVolumeSnapshotUsedBytes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
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

func flattenNetappVolumeSnapshotTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
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

func flattenNetappVolumeSnapshotEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandNetappVolumeSnapshotDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappVolumeSnapshotEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
