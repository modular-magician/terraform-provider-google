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

package compute

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceComputeInstanceGroupMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeInstanceGroupMembershipCreate,
		Read:   resourceComputeInstanceGroupMembershipRead,
		Delete: resourceComputeInstanceGroupMembershipDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeInstanceGroupMembershipImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"instance": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `An instance being added to the InstanceGroup`,
			},
			"instance_group": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareResourceNames,
				Description:      `Represents an Instance Group resource name that the instance belongs to.`,
			},
			"zone": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `A reference to the zone where the instance group resides.`,
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

func resourceComputeInstanceGroupMembershipCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	instanceProp, err := expandNestedComputeInstanceGroupMembershipInstance(d.Get("instance"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("instance"); !tpgresource.IsEmptyValue(reflect.ValueOf(instanceProp)) && (ok || !reflect.DeepEqual(v, instanceProp)) {
		obj["instance"] = instanceProp
	}

	obj, err = resourceComputeInstanceGroupMembershipEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "instanceGroups/{{project}}/zones/{{zone}}/{{instance_group}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/instanceGroups/{{instance_group}}/addInstances")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new InstanceGroupMembership: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InstanceGroupMembership: %s", err)
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
		return fmt.Errorf("Error creating InstanceGroupMembership: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{zone}}/{{instance_group}}/{{instance}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating InstanceGroupMembership", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create InstanceGroupMembership: %s", err)
	}

	log.Printf("[DEBUG] Finished creating InstanceGroupMembership %q: %#v", d.Id(), res)

	return resourceComputeInstanceGroupMembershipRead(d, meta)
}

func resourceComputeInstanceGroupMembershipRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/instanceGroups/{{instance_group}}/listInstances")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InstanceGroupMembership: %s", err)
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
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeInstanceGroupMembership %q", d.Id()))
	}

	res, err = flattenNestedComputeInstanceGroupMembership(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeInstanceGroupMembership because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading InstanceGroupMembership: %s", err)
	}

	if err := d.Set("instance", flattenNestedComputeInstanceGroupMembershipInstance(res["instance"], d, config)); err != nil {
		return fmt.Errorf("Error reading InstanceGroupMembership: %s", err)
	}

	return nil
}

func resourceComputeInstanceGroupMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for InstanceGroupMembership: %s", err)
	}
	billingProject = project

	lockName, err := tpgresource.ReplaceVars(d, config, "instanceGroups/{{project}}/zones/{{zone}}/{{instance_group}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/instanceGroups/{{instance_group}}/removeInstances")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	toDelete := make(map[string]interface{})

	// Instance
	instanceProp := flattenNestedComputeInstanceGroupMembershipInstance(d.Get("instance"), d, config)

	if instanceProp != "" {
		toDelete["instance"] = instanceProp
	}

	obj = map[string]interface{}{
		"instances": []map[string]interface{}{toDelete},
	}
	log.Printf("[DEBUG] Deleting InstanceGroupMembership %q", d.Id())

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
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "InstanceGroupMembership")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting InstanceGroupMembership", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting InstanceGroupMembership %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeInstanceGroupMembershipImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/instanceGroups/(?P<instance_group>[^/]+)/(?P<instance>.+)",
		"(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<instance_group>[^/]+)/(?P<instance>.+)",
		"(?P<zone>[^/]+)/(?P<instance_group>[^/]+)/(?P<instance>.+)",
		"(?P<instance_group>[^/]+)/(?P<instance>.+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{project}}/{{zone}}/{{instance_group}}/{{instance}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNestedComputeInstanceGroupMembershipInstance(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	relative, err := tpgresource.GetRelativePath(v.(string))
	if err != nil {
		return v
	}
	return relative
}

func expandNestedComputeInstanceGroupMembershipInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseZonalFieldValue("instances", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for instance: %s", err)
	}
	return f.RelativeLink(), nil
}

func resourceComputeInstanceGroupMembershipEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Instance Group is a URL parameter only, so replace self-link/path with resource name only.
	if err := d.Set("instance_group", tpgresource.GetResourceNameFromSelfLink(d.Get("instance_group").(string))); err != nil {
		return nil, fmt.Errorf("Error setting instance_group: %s", err)
	}

	wrappedReq := map[string]interface{}{
		"instances": []interface{}{obj},
	}
	return wrappedReq, nil
}

func flattenNestedComputeInstanceGroupMembership(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["items"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value items. Actual value: %v", v)
	}

	_, item, err := resourceComputeInstanceGroupMembershipFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceComputeInstanceGroupMembershipFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedInstance, err := expandNestedComputeInstanceGroupMembershipInstance(d.Get("instance"), d, meta.(*transport_tpg.Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedInstance := flattenNestedComputeInstanceGroupMembershipInstance(expectedInstance, d, meta.(*transport_tpg.Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemInstance := flattenNestedComputeInstanceGroupMembershipInstance(item["instance"], d, meta.(*transport_tpg.Config))
		// IsEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(tpgresource.IsEmptyValue(reflect.ValueOf(itemInstance)) && tpgresource.IsEmptyValue(reflect.ValueOf(expectedFlattenedInstance))) && !reflect.DeepEqual(itemInstance, expectedFlattenedInstance) {
			log.Printf("[DEBUG] Skipping item with instance= %#v, looking for %#v)", itemInstance, expectedFlattenedInstance)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
