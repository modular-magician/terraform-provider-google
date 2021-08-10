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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	gkemulticloud "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud"
)

func resourceGkemulticloudAwsNodePool() *schema.Resource {
	return &schema.Resource{
		Create: resourceGkemulticloudAwsNodePoolCreate,
		Read:   resourceGkemulticloudAwsNodePoolRead,
		Delete: resourceGkemulticloudAwsNodePoolDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGkemulticloudAwsNodePoolImport,
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
				Elem:        GkemulticloudAwsNodePoolAutoscalingSchema(),
			},

			"aws_cluster": {
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
				Elem:        GkemulticloudAwsNodePoolConfigSchema(),
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
				Elem:        GkemulticloudAwsNodePoolMaxPodsConstraintSchema(),
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

func GkemulticloudAwsNodePoolAutoscalingSchema() *schema.Resource {
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

func GkemulticloudAwsNodePoolConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"iam_instance_profile": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"instance_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"root_volume": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAwsNodePoolConfigRootVolumeSchema(),
			},

			"security_group_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"ssh_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAwsNodePoolConfigSshConfigSchema(),
			},

			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"taints": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        GkemulticloudAwsNodePoolConfigTaintsSchema(),
			},
		},
	}
}

func GkemulticloudAwsNodePoolConfigRootVolumeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"iops": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"kms_key_arn": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"size_gib": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"volume_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  ``,
				ValidateFunc: validation.StringInSlice([]string{"VOLUME_TYPE_UNSPECIFIED", "GP2", "GP3", ""}, false),
			},
		},
	}
}

func GkemulticloudAwsNodePoolConfigSshConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ec2_key_pair": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func GkemulticloudAwsNodePoolConfigTaintsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"effect": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  ``,
				ValidateFunc: validation.StringInSlice([]string{"EFFECT_UNSPECIFIED", "NO_SCHEDULE", "PREFER_NO_SCHEDULE", "NO_EXECUTE", ""}, false),
			},

			"key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"value": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func GkemulticloudAwsNodePoolMaxPodsConstraintSchema() *schema.Resource {
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

func resourceGkemulticloudAwsNodePoolCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkemulticloud.AwsNodePool{
		Autoscaling:       expandGkemulticloudAwsNodePoolAutoscaling(d.Get("autoscaling")),
		AwsCluster:        dcl.String(d.Get("aws_cluster").(string)),
		Config:            expandGkemulticloudAwsNodePoolConfig(d.Get("config")),
		Location:          dcl.String(d.Get("location").(string)),
		MaxPodsConstraint: expandGkemulticloudAwsNodePoolMaxPodsConstraint(d.Get("max_pods_constraint")),
		Name:              dcl.String(d.Get("name").(string)),
		SubnetId:          dcl.String(d.Get("subnet_id").(string)),
		Version:           dcl.String(d.Get("version").(string)),
		Annotations:       checkStringMap(d.Get("annotations")),
		Project:           dcl.String(project),
	}

	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/awsClusters/{{aws_cluster}}/awsNodePools/{{name}}")
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
	res, err := client.ApplyAwsNodePool(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating AwsNodePool: %s", err)
	}

	log.Printf("[DEBUG] Finished creating AwsNodePool %q: %#v", d.Id(), res)

	return resourceGkemulticloudAwsNodePoolRead(d, meta)
}

func resourceGkemulticloudAwsNodePoolRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkemulticloud.AwsNodePool{
		Autoscaling:       expandGkemulticloudAwsNodePoolAutoscaling(d.Get("autoscaling")),
		AwsCluster:        dcl.String(d.Get("aws_cluster").(string)),
		Config:            expandGkemulticloudAwsNodePoolConfig(d.Get("config")),
		Location:          dcl.String(d.Get("location").(string)),
		MaxPodsConstraint: expandGkemulticloudAwsNodePoolMaxPodsConstraint(d.Get("max_pods_constraint")),
		Name:              dcl.String(d.Get("name").(string)),
		SubnetId:          dcl.String(d.Get("subnet_id").(string)),
		Version:           dcl.String(d.Get("version").(string)),
		Annotations:       checkStringMap(d.Get("annotations")),
		Project:           dcl.String(project),
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
	res, err := client.GetAwsNodePool(context.Background(), obj)
	if err != nil {
		// Resource not found
		d.SetId("")
		return err
	}

	if err = d.Set("autoscaling", flattenGkemulticloudAwsNodePoolAutoscaling(res.Autoscaling)); err != nil {
		return fmt.Errorf("error setting autoscaling in state: %s", err)
	}
	if err = d.Set("aws_cluster", res.AwsCluster); err != nil {
		return fmt.Errorf("error setting aws_cluster in state: %s", err)
	}
	if err = d.Set("config", flattenGkemulticloudAwsNodePoolConfig(res.Config)); err != nil {
		return fmt.Errorf("error setting config in state: %s", err)
	}
	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("max_pods_constraint", flattenGkemulticloudAwsNodePoolMaxPodsConstraint(res.MaxPodsConstraint)); err != nil {
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

func resourceGkemulticloudAwsNodePoolDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkemulticloud.AwsNodePool{
		Autoscaling:       expandGkemulticloudAwsNodePoolAutoscaling(d.Get("autoscaling")),
		AwsCluster:        dcl.String(d.Get("aws_cluster").(string)),
		Config:            expandGkemulticloudAwsNodePoolConfig(d.Get("config")),
		Location:          dcl.String(d.Get("location").(string)),
		MaxPodsConstraint: expandGkemulticloudAwsNodePoolMaxPodsConstraint(d.Get("max_pods_constraint")),
		Name:              dcl.String(d.Get("name").(string)),
		SubnetId:          dcl.String(d.Get("subnet_id").(string)),
		Version:           dcl.String(d.Get("version").(string)),
		Annotations:       checkStringMap(d.Get("annotations")),
		Project:           dcl.String(project),
	}

	log.Printf("[DEBUG] Deleting AwsNodePool %q", d.Id())
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
	if err := client.DeleteAwsNodePool(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting AwsNodePool: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting AwsNodePool %q", d.Id())
	return nil
}

func resourceGkemulticloudAwsNodePoolImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/awsClusters/(?P<aws_cluster>[^/]+)/awsNodePools/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<aws_cluster>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<aws_cluster>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/awsClusters/{{aws_cluster}}/awsNodePools/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandGkemulticloudAwsNodePoolAutoscaling(o interface{}) *gkemulticloud.AwsNodePoolAutoscaling {
	if o == nil {
		return gkemulticloud.EmptyAwsNodePoolAutoscaling
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAwsNodePoolAutoscaling
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AwsNodePoolAutoscaling{
		MaxNodeCount: dcl.Int64(int64(obj["max_node_count"].(int))),
		MinNodeCount: dcl.Int64(int64(obj["min_node_count"].(int))),
	}
}

func flattenGkemulticloudAwsNodePoolAutoscaling(obj *gkemulticloud.AwsNodePoolAutoscaling) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"max_node_count": obj.MaxNodeCount,
		"min_node_count": obj.MinNodeCount,
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAwsNodePoolConfig(o interface{}) *gkemulticloud.AwsNodePoolConfig {
	if o == nil {
		return gkemulticloud.EmptyAwsNodePoolConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAwsNodePoolConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AwsNodePoolConfig{
		IamInstanceProfile: dcl.String(obj["iam_instance_profile"].(string)),
		InstanceType:       dcl.String(obj["instance_type"].(string)),
		Labels:             checkStringMap(obj["labels"]),
		RootVolume:         expandGkemulticloudAwsNodePoolConfigRootVolume(obj["root_volume"]),
		SecurityGroupIds:   expandStringArray(obj["security_group_ids"]),
		SshConfig:          expandGkemulticloudAwsNodePoolConfigSshConfig(obj["ssh_config"]),
		Tags:               checkStringMap(obj["tags"]),
		Taints:             expandGkemulticloudAwsNodePoolConfigTaintsArray(obj["taints"]),
	}
}

func flattenGkemulticloudAwsNodePoolConfig(obj *gkemulticloud.AwsNodePoolConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"iam_instance_profile": obj.IamInstanceProfile,
		"instance_type":        obj.InstanceType,
		"labels":               obj.Labels,
		"root_volume":          flattenGkemulticloudAwsNodePoolConfigRootVolume(obj.RootVolume),
		"security_group_ids":   obj.SecurityGroupIds,
		"ssh_config":           flattenGkemulticloudAwsNodePoolConfigSshConfig(obj.SshConfig),
		"tags":                 obj.Tags,
		"taints":               flattenGkemulticloudAwsNodePoolConfigTaintsArray(obj.Taints),
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAwsNodePoolConfigRootVolume(o interface{}) *gkemulticloud.AwsNodePoolConfigRootVolume {
	if o == nil {
		return gkemulticloud.EmptyAwsNodePoolConfigRootVolume
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAwsNodePoolConfigRootVolume
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AwsNodePoolConfigRootVolume{
		Iops:       dcl.Int64(int64(obj["iops"].(int))),
		KmsKeyArn:  dcl.String(obj["kms_key_arn"].(string)),
		SizeGib:    dcl.Int64(int64(obj["size_gib"].(int))),
		VolumeType: gkemulticloud.AwsNodePoolConfigRootVolumeVolumeTypeEnumRef(obj["volume_type"].(string)),
	}
}

func flattenGkemulticloudAwsNodePoolConfigRootVolume(obj *gkemulticloud.AwsNodePoolConfigRootVolume) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"iops":        obj.Iops,
		"kms_key_arn": obj.KmsKeyArn,
		"size_gib":    obj.SizeGib,
		"volume_type": obj.VolumeType,
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAwsNodePoolConfigSshConfig(o interface{}) *gkemulticloud.AwsNodePoolConfigSshConfig {
	if o == nil {
		return gkemulticloud.EmptyAwsNodePoolConfigSshConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAwsNodePoolConfigSshConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AwsNodePoolConfigSshConfig{
		Ec2KeyPair: dcl.String(obj["ec2_key_pair"].(string)),
	}
}

func flattenGkemulticloudAwsNodePoolConfigSshConfig(obj *gkemulticloud.AwsNodePoolConfigSshConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"ec2_key_pair": obj.Ec2KeyPair,
	}

	return []interface{}{transformed}

}
func expandGkemulticloudAwsNodePoolConfigTaintsArray(o interface{}) []gkemulticloud.AwsNodePoolConfigTaints {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]gkemulticloud.AwsNodePoolConfigTaints, 0, len(objs))
	for _, item := range objs {
		i := expandGkemulticloudAwsNodePoolConfigTaints(item)
		items = append(items, *i)
	}

	return items
}

func expandGkemulticloudAwsNodePoolConfigTaints(o interface{}) *gkemulticloud.AwsNodePoolConfigTaints {
	if o == nil {
		return gkemulticloud.EmptyAwsNodePoolConfigTaints
	}

	obj := o.(map[string]interface{})
	return &gkemulticloud.AwsNodePoolConfigTaints{
		Effect: gkemulticloud.AwsNodePoolConfigTaintsEffectEnumRef(obj["effect"].(string)),
		Key:    dcl.String(obj["key"].(string)),
		Value:  dcl.String(obj["value"].(string)),
	}
}

func flattenGkemulticloudAwsNodePoolConfigTaintsArray(objs []gkemulticloud.AwsNodePoolConfigTaints) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenGkemulticloudAwsNodePoolConfigTaints(&item)
		items = append(items, i)
	}

	return items
}

func flattenGkemulticloudAwsNodePoolConfigTaints(obj *gkemulticloud.AwsNodePoolConfigTaints) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"effect": obj.Effect,
		"key":    obj.Key,
		"value":  obj.Value,
	}

	return transformed

}

func expandGkemulticloudAwsNodePoolMaxPodsConstraint(o interface{}) *gkemulticloud.AwsNodePoolMaxPodsConstraint {
	if o == nil {
		return gkemulticloud.EmptyAwsNodePoolMaxPodsConstraint
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAwsNodePoolMaxPodsConstraint
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AwsNodePoolMaxPodsConstraint{
		MaxPodsPerNode: dcl.Int64(int64(obj["max_pods_per_node"].(int))),
	}
}

func flattenGkemulticloudAwsNodePoolMaxPodsConstraint(obj *gkemulticloud.AwsNodePoolMaxPodsConstraint) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"max_pods_per_node": obj.MaxPodsPerNode,
	}

	return []interface{}{transformed}

}
