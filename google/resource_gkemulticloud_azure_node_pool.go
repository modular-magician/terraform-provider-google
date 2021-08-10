// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	gkemulticloud "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud"
)

func resourceGkemulticloudAzureNodePool() *schema.Resource {
	return &schema.Resource{
		Create: resourceGkemulticloudAzureNodePoolCreate,
		Read:   resourceGkemulticloudAzureNodePoolRead,
		Delete: resourceGkemulticloudAzureNodePoolDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGkemulticloudAzureNodePoolImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"autoscaling": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAzureNodePoolAutoscalingSchema(),
			},

			"azure_cluster": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"config": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAzureNodePoolConfigSchema(),
			},

			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"max_pods_constraint": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAzureNodePoolMaxPodsConstraintSchema(),
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"annotations": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"azure_availability_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"reconciling": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: ``,
			},

			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},
		},
	}
}

func GkemulticloudAzureNodePoolAutoscalingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_node_count": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"min_node_count": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func GkemulticloudAzureNodePoolConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ssh_config": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAzureNodePoolConfigSshConfigSchema(),
			},

			"root_volume": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAzureNodePoolConfigRootVolumeSchema(),
			},

			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"vm_size": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func GkemulticloudAzureNodePoolConfigSshConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authorized_key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func GkemulticloudAzureNodePoolConfigRootVolumeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"size_gib": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func GkemulticloudAzureNodePoolMaxPodsConstraintSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_pods_per_node": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func resourceGkemulticloudAzureNodePoolCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkemulticloud.AzureNodePool{
		Autoscaling:           expandGkemulticloudAzureNodePoolAutoscaling(d.Get("autoscaling")),
		AzureCluster:          dcl.String(d.Get("azure_cluster").(string)),
		Config:                expandGkemulticloudAzureNodePoolConfig(d.Get("config")),
		Location:              dcl.String(d.Get("location").(string)),
		MaxPodsConstraint:     expandGkemulticloudAzureNodePoolMaxPodsConstraint(d.Get("max_pods_constraint")),
		Name:                  dcl.String(d.Get("name").(string)),
		SubnetId:              dcl.String(d.Get("subnet_id").(string)),
		Version:               dcl.String(d.Get("version").(string)),
		Annotations:           checkStringMap(d.Get("annotations")),
		AzureAvailabilityZone: dcl.String(d.Get("azure_availability_zone").(string)),
		Project:               dcl.String(project),
	}

	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/azureClusters/{{azure_cluster}}/azureNodePools/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	createDirective := CreateDirective
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLGkemulticloudClient(config, userAgent, billingProject)
	res, err := client.ApplyAzureNodePool(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating AzureNodePool: %s", err)
	}

	log.Printf("[DEBUG] Finished creating AzureNodePool %q: %#v", d.Id(), res)

	return resourceGkemulticloudAzureNodePoolRead(d, meta)
}

func resourceGkemulticloudAzureNodePoolRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkemulticloud.AzureNodePool{
		Autoscaling:           expandGkemulticloudAzureNodePoolAutoscaling(d.Get("autoscaling")),
		AzureCluster:          dcl.String(d.Get("azure_cluster").(string)),
		Config:                expandGkemulticloudAzureNodePoolConfig(d.Get("config")),
		Location:              dcl.String(d.Get("location").(string)),
		MaxPodsConstraint:     expandGkemulticloudAzureNodePoolMaxPodsConstraint(d.Get("max_pods_constraint")),
		Name:                  dcl.String(d.Get("name").(string)),
		SubnetId:              dcl.String(d.Get("subnet_id").(string)),
		Version:               dcl.String(d.Get("version").(string)),
		Annotations:           checkStringMap(d.Get("annotations")),
		AzureAvailabilityZone: dcl.String(d.Get("azure_availability_zone").(string)),
		Project:               dcl.String(project),
	}

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLGkemulticloudClient(config, userAgent, billingProject)
	res, err := client.GetAzureNodePool(context.Background(), obj)
	if err != nil {
		// Resource not found
		d.SetId("")
		return err
	}

	if err = d.Set("autoscaling", flattenGkemulticloudAzureNodePoolAutoscaling(res.Autoscaling)); err != nil {
		return fmt.Errorf("error setting autoscaling in state: %s", err)
	}
	if err = d.Set("azure_cluster", res.AzureCluster); err != nil {
		return fmt.Errorf("error setting azure_cluster in state: %s", err)
	}
	if err = d.Set("config", flattenGkemulticloudAzureNodePoolConfig(res.Config)); err != nil {
		return fmt.Errorf("error setting config in state: %s", err)
	}
	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("max_pods_constraint", flattenGkemulticloudAzureNodePoolMaxPodsConstraint(res.MaxPodsConstraint)); err != nil {
		return fmt.Errorf("error setting max_pods_constraint in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("subnet_id", res.SubnetId); err != nil {
		return fmt.Errorf("error setting subnet_id in state: %s", err)
	}
	if err = d.Set("version", res.Version); err != nil {
		return fmt.Errorf("error setting version in state: %s", err)
	}
	if err = d.Set("annotations", res.Annotations); err != nil {
		return fmt.Errorf("error setting annotations in state: %s", err)
	}
	if err = d.Set("azure_availability_zone", res.AzureAvailabilityZone); err != nil {
		return fmt.Errorf("error setting azure_availability_zone in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("etag", res.Etag); err != nil {
		return fmt.Errorf("error setting etag in state: %s", err)
	}
	if err = d.Set("reconciling", res.Reconciling); err != nil {
		return fmt.Errorf("error setting reconciling in state: %s", err)
	}
	if err = d.Set("state", res.State); err != nil {
		return fmt.Errorf("error setting state in state: %s", err)
	}
	if err = d.Set("uid", res.Uid); err != nil {
		return fmt.Errorf("error setting uid in state: %s", err)
	}
	if err = d.Set("update_time", res.UpdateTime); err != nil {
		return fmt.Errorf("error setting update_time in state: %s", err)
	}

	return nil
}

func resourceGkemulticloudAzureNodePoolDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkemulticloud.AzureNodePool{
		Autoscaling:           expandGkemulticloudAzureNodePoolAutoscaling(d.Get("autoscaling")),
		AzureCluster:          dcl.String(d.Get("azure_cluster").(string)),
		Config:                expandGkemulticloudAzureNodePoolConfig(d.Get("config")),
		Location:              dcl.String(d.Get("location").(string)),
		MaxPodsConstraint:     expandGkemulticloudAzureNodePoolMaxPodsConstraint(d.Get("max_pods_constraint")),
		Name:                  dcl.String(d.Get("name").(string)),
		SubnetId:              dcl.String(d.Get("subnet_id").(string)),
		Version:               dcl.String(d.Get("version").(string)),
		Annotations:           checkStringMap(d.Get("annotations")),
		AzureAvailabilityZone: dcl.String(d.Get("azure_availability_zone").(string)),
		Project:               dcl.String(project),
	}

	log.Printf("[DEBUG] Deleting AzureNodePool %q", d.Id())
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLGkemulticloudClient(config, userAgent, billingProject)
	if err := client.DeleteAzureNodePool(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting AzureNodePool: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting AzureNodePool %q", d.Id())
	return nil
}

func resourceGkemulticloudAzureNodePoolImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/azureClusters/(?P<azure_cluster>[^/]+)/azureNodePools/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<azure_cluster>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<azure_cluster>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/azureClusters/{{azure_cluster}}/azureNodePools/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandGkemulticloudAzureNodePoolAutoscaling(o interface{}) *gkemulticloud.AzureNodePoolAutoscaling {
	if o == nil {
		return gkemulticloud.EmptyAzureNodePoolAutoscaling
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAzureNodePoolAutoscaling
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AzureNodePoolAutoscaling{
		MaxNodeCount: dcl.Int64(int64(obj["max_node_count"].(int))),
		MinNodeCount: dcl.Int64(int64(obj["min_node_count"].(int))),
	}
}

func flattenGkemulticloudAzureNodePoolAutoscaling(obj *gkemulticloud.AzureNodePoolAutoscaling) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"max_node_count": obj.MaxNodeCount,
		"min_node_count": obj.MinNodeCount,
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAzureNodePoolConfig(o interface{}) *gkemulticloud.AzureNodePoolConfig {
	if o == nil {
		return gkemulticloud.EmptyAzureNodePoolConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAzureNodePoolConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AzureNodePoolConfig{
		SshConfig:  expandGkemulticloudAzureNodePoolConfigSshConfig(obj["ssh_config"]),
		RootVolume: expandGkemulticloudAzureNodePoolConfigRootVolume(obj["root_volume"]),
		Tags:       checkStringMap(obj["tags"]),
		VmSize:     dcl.String(obj["vm_size"].(string)),
	}
}

func flattenGkemulticloudAzureNodePoolConfig(obj *gkemulticloud.AzureNodePoolConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"ssh_config":  flattenGkemulticloudAzureNodePoolConfigSshConfig(obj.SshConfig),
		"root_volume": flattenGkemulticloudAzureNodePoolConfigRootVolume(obj.RootVolume),
		"tags":        obj.Tags,
		"vm_size":     obj.VmSize,
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAzureNodePoolConfigSshConfig(o interface{}) *gkemulticloud.AzureNodePoolConfigSshConfig {
	if o == nil {
		return gkemulticloud.EmptyAzureNodePoolConfigSshConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAzureNodePoolConfigSshConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AzureNodePoolConfigSshConfig{
		AuthorizedKey: dcl.String(obj["authorized_key"].(string)),
	}
}

func flattenGkemulticloudAzureNodePoolConfigSshConfig(obj *gkemulticloud.AzureNodePoolConfigSshConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"authorized_key": obj.AuthorizedKey,
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAzureNodePoolConfigRootVolume(o interface{}) *gkemulticloud.AzureNodePoolConfigRootVolume {
	if o == nil {
		return gkemulticloud.EmptyAzureNodePoolConfigRootVolume
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAzureNodePoolConfigRootVolume
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AzureNodePoolConfigRootVolume{
		SizeGib: dcl.Int64(int64(obj["size_gib"].(int))),
	}
}

func flattenGkemulticloudAzureNodePoolConfigRootVolume(obj *gkemulticloud.AzureNodePoolConfigRootVolume) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"size_gib": obj.SizeGib,
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAzureNodePoolMaxPodsConstraint(o interface{}) *gkemulticloud.AzureNodePoolMaxPodsConstraint {
	if o == nil {
		return gkemulticloud.EmptyAzureNodePoolMaxPodsConstraint
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAzureNodePoolMaxPodsConstraint
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AzureNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64(int64(obj["max_pods_per_node"].(int))),
	}
}

func flattenGkemulticloudAzureNodePoolMaxPodsConstraint(obj *gkemulticloud.AzureNodePoolMaxPodsConstraint) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"max_pods_per_node": obj.MaxPodsPerNode,
	}

	return []interface{}{transformed}

}
