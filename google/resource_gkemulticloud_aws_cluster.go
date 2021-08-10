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

func resourceGkemulticloudAwsCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceGkemulticloudAwsClusterCreate,
		Read:   resourceGkemulticloudAwsClusterRead,
		Delete: resourceGkemulticloudAwsClusterDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGkemulticloudAwsClusterImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"authorization": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAwsClusterAuthorizationSchema(),
			},

			"aws_region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"control_plane": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAwsClusterControlPlaneSchema(),
			},

			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"networking": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAwsClusterNetworkingSchema(),
			},

			"annotations": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"description": {
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

			"endpoint": {
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

			"workload_identity_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: ``,
				Elem:        GkemulticloudAwsClusterWorkloadIdentityConfigSchema(),
			},
		},
	}
}

func GkemulticloudAwsClusterAuthorizationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_users": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        GkemulticloudAwsClusterAuthorizationAdminUsersSchema(),
			},
		},
	}
}

func GkemulticloudAwsClusterAuthorizationAdminUsersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func GkemulticloudAwsClusterControlPlaneSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aws_services_authentication": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAwsClusterControlPlaneAwsServicesAuthenticationSchema(),
			},

			"database_encryption": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAwsClusterControlPlaneDatabaseEncryptionSchema(),
			},

			"iam_instance_profile": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"subnet_ids": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"version": {
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

			"main_volume": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAwsClusterControlPlaneMainVolumeSchema(),
			},

			"root_volume": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAwsClusterControlPlaneRootVolumeSchema(),
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
				Elem:        GkemulticloudAwsClusterControlPlaneSshConfigSchema(),
			},

			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func GkemulticloudAwsClusterControlPlaneAwsServicesAuthenticationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"role_arn": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"role_session_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func GkemulticloudAwsClusterControlPlaneDatabaseEncryptionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kms_key_arn": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func GkemulticloudAwsClusterControlPlaneMainVolumeSchema() *schema.Resource {
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

func GkemulticloudAwsClusterControlPlaneRootVolumeSchema() *schema.Resource {
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

func GkemulticloudAwsClusterControlPlaneSshConfigSchema() *schema.Resource {
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

func GkemulticloudAwsClusterNetworkingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pod_address_cidr_blocks": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"service_address_cidr_blocks": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"service_load_balancer_subnet_ids": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func GkemulticloudAwsClusterWorkloadIdentityConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"identity_provider": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"issuer_uri": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"workload_pool": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},
		},
	}
}

func resourceGkemulticloudAwsClusterCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkemulticloud.AwsCluster{
		Authorization: expandGkemulticloudAwsClusterAuthorization(d.Get("authorization")),
		AwsRegion:     dcl.String(d.Get("aws_region").(string)),
		ControlPlane:  expandGkemulticloudAwsClusterControlPlane(d.Get("control_plane")),
		Location:      dcl.String(d.Get("location").(string)),
		Name:          dcl.String(d.Get("name").(string)),
		Networking:    expandGkemulticloudAwsClusterNetworking(d.Get("networking")),
		Annotations:   checkStringMap(d.Get("annotations")),
		Description:   dcl.String(d.Get("description").(string)),
		Project:       dcl.String(project),
	}

	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/awsClusters/{{name}}")
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
	res, err := client.ApplyAwsCluster(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating AwsCluster: %s", err)
	}

	log.Printf("[DEBUG] Finished creating AwsCluster %q: %#v", d.Id(), res)

	return resourceGkemulticloudAwsClusterRead(d, meta)
}

func resourceGkemulticloudAwsClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkemulticloud.AwsCluster{
		Authorization: expandGkemulticloudAwsClusterAuthorization(d.Get("authorization")),
		AwsRegion:     dcl.String(d.Get("aws_region").(string)),
		ControlPlane:  expandGkemulticloudAwsClusterControlPlane(d.Get("control_plane")),
		Location:      dcl.String(d.Get("location").(string)),
		Name:          dcl.String(d.Get("name").(string)),
		Networking:    expandGkemulticloudAwsClusterNetworking(d.Get("networking")),
		Annotations:   checkStringMap(d.Get("annotations")),
		Description:   dcl.String(d.Get("description").(string)),
		Project:       dcl.String(project),
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
	res, err := client.GetAwsCluster(context.Background(), obj)
	if err != nil {
		// Resource not found
		d.SetId("")
		return err
	}

	if err = d.Set("authorization", flattenGkemulticloudAwsClusterAuthorization(res.Authorization)); err != nil {
		return fmt.Errorf("error setting authorization in state: %s", err)
	}
	if err = d.Set("aws_region", res.AwsRegion); err != nil {
		return fmt.Errorf("error setting aws_region in state: %s", err)
	}
	if err = d.Set("control_plane", flattenGkemulticloudAwsClusterControlPlane(res.ControlPlane)); err != nil {
		return fmt.Errorf("error setting control_plane in state: %s", err)
	}
	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("networking", flattenGkemulticloudAwsClusterNetworking(res.Networking)); err != nil {
		return fmt.Errorf("error setting networking in state: %s", err)
	}
	if err = d.Set("annotations", res.Annotations); err != nil {
		return fmt.Errorf("error setting annotations in state: %s", err)
	}
	if err = d.Set("description", res.Description); err != nil {
		return fmt.Errorf("error setting description in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("endpoint", res.Endpoint); err != nil {
		return fmt.Errorf("error setting endpoint in state: %s", err)
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
	if err = d.Set("workload_identity_config", flattenGkemulticloudAwsClusterWorkloadIdentityConfig(res.WorkloadIdentityConfig)); err != nil {
		return fmt.Errorf("error setting workload_identity_config in state: %s", err)
	}

	return nil
}

func resourceGkemulticloudAwsClusterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkemulticloud.AwsCluster{
		Authorization: expandGkemulticloudAwsClusterAuthorization(d.Get("authorization")),
		AwsRegion:     dcl.String(d.Get("aws_region").(string)),
		ControlPlane:  expandGkemulticloudAwsClusterControlPlane(d.Get("control_plane")),
		Location:      dcl.String(d.Get("location").(string)),
		Name:          dcl.String(d.Get("name").(string)),
		Networking:    expandGkemulticloudAwsClusterNetworking(d.Get("networking")),
		Annotations:   checkStringMap(d.Get("annotations")),
		Description:   dcl.String(d.Get("description").(string)),
		Project:       dcl.String(project),
	}

	log.Printf("[DEBUG] Deleting AwsCluster %q", d.Id())
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
	if err := client.DeleteAwsCluster(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting AwsCluster: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting AwsCluster %q", d.Id())
	return nil
}

func resourceGkemulticloudAwsClusterImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/awsClusters/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/awsClusters/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandGkemulticloudAwsClusterAuthorization(o interface{}) *gkemulticloud.AwsClusterAuthorization {
	if o == nil {
		return gkemulticloud.EmptyAwsClusterAuthorization
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAwsClusterAuthorization
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AwsClusterAuthorization{
		AdminUsers: expandGkemulticloudAwsClusterAuthorizationAdminUsersArray(obj["admin_users"]),
	}
}

func flattenGkemulticloudAwsClusterAuthorization(obj *gkemulticloud.AwsClusterAuthorization) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"admin_users": flattenGkemulticloudAwsClusterAuthorizationAdminUsersArray(obj.AdminUsers),
	}

	return []interface{}{transformed}

}
func expandGkemulticloudAwsClusterAuthorizationAdminUsersArray(o interface{}) []gkemulticloud.AwsClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]gkemulticloud.AwsClusterAuthorizationAdminUsers, 0, len(objs))
	for _, item := range objs {
		i := expandGkemulticloudAwsClusterAuthorizationAdminUsers(item)
		items = append(items, *i)
	}

	return items
}

func expandGkemulticloudAwsClusterAuthorizationAdminUsers(o interface{}) *gkemulticloud.AwsClusterAuthorizationAdminUsers {
	if o == nil {
		return gkemulticloud.EmptyAwsClusterAuthorizationAdminUsers
	}

	obj := o.(map[string]interface{})
	return &gkemulticloud.AwsClusterAuthorizationAdminUsers{
		Username: dcl.String(obj["username"].(string)),
	}
}

func flattenGkemulticloudAwsClusterAuthorizationAdminUsersArray(objs []gkemulticloud.AwsClusterAuthorizationAdminUsers) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenGkemulticloudAwsClusterAuthorizationAdminUsers(&item)
		items = append(items, i)
	}

	return items
}

func flattenGkemulticloudAwsClusterAuthorizationAdminUsers(obj *gkemulticloud.AwsClusterAuthorizationAdminUsers) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"username": obj.Username,
	}

	return transformed

}

func expandGkemulticloudAwsClusterControlPlane(o interface{}) *gkemulticloud.AwsClusterControlPlane {
	if o == nil {
		return gkemulticloud.EmptyAwsClusterControlPlane
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAwsClusterControlPlane
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AwsClusterControlPlane{
		AwsServicesAuthentication: expandGkemulticloudAwsClusterControlPlaneAwsServicesAuthentication(obj["aws_services_authentication"]),
		DatabaseEncryption:        expandGkemulticloudAwsClusterControlPlaneDatabaseEncryption(obj["database_encryption"]),
		IamInstanceProfile:        dcl.String(obj["iam_instance_profile"].(string)),
		SubnetIds:                 expandStringArray(obj["subnet_ids"]),
		Version:                   dcl.String(obj["version"].(string)),
		InstanceType:              dcl.String(obj["instance_type"].(string)),
		MainVolume:                expandGkemulticloudAwsClusterControlPlaneMainVolume(obj["main_volume"]),
		RootVolume:                expandGkemulticloudAwsClusterControlPlaneRootVolume(obj["root_volume"]),
		SecurityGroupIds:          expandStringArray(obj["security_group_ids"]),
		SshConfig:                 expandGkemulticloudAwsClusterControlPlaneSshConfig(obj["ssh_config"]),
		Tags:                      checkStringMap(obj["tags"]),
	}
}

func flattenGkemulticloudAwsClusterControlPlane(obj *gkemulticloud.AwsClusterControlPlane) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"aws_services_authentication": flattenGkemulticloudAwsClusterControlPlaneAwsServicesAuthentication(obj.AwsServicesAuthentication),
		"database_encryption":         flattenGkemulticloudAwsClusterControlPlaneDatabaseEncryption(obj.DatabaseEncryption),
		"iam_instance_profile":        obj.IamInstanceProfile,
		"subnet_ids":                  obj.SubnetIds,
		"version":                     obj.Version,
		"instance_type":               obj.InstanceType,
		"main_volume":                 flattenGkemulticloudAwsClusterControlPlaneMainVolume(obj.MainVolume),
		"root_volume":                 flattenGkemulticloudAwsClusterControlPlaneRootVolume(obj.RootVolume),
		"security_group_ids":          obj.SecurityGroupIds,
		"ssh_config":                  flattenGkemulticloudAwsClusterControlPlaneSshConfig(obj.SshConfig),
		"tags":                        obj.Tags,
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAwsClusterControlPlaneAwsServicesAuthentication(o interface{}) *gkemulticloud.AwsClusterControlPlaneAwsServicesAuthentication {
	if o == nil {
		return gkemulticloud.EmptyAwsClusterControlPlaneAwsServicesAuthentication
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAwsClusterControlPlaneAwsServicesAuthentication
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AwsClusterControlPlaneAwsServicesAuthentication{
		RoleArn:         dcl.String(obj["role_arn"].(string)),
		RoleSessionName: dcl.String(obj["role_session_name"].(string)),
	}
}

func flattenGkemulticloudAwsClusterControlPlaneAwsServicesAuthentication(obj *gkemulticloud.AwsClusterControlPlaneAwsServicesAuthentication) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"role_arn":          obj.RoleArn,
		"role_session_name": obj.RoleSessionName,
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAwsClusterControlPlaneDatabaseEncryption(o interface{}) *gkemulticloud.AwsClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return gkemulticloud.EmptyAwsClusterControlPlaneDatabaseEncryption
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAwsClusterControlPlaneDatabaseEncryption
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AwsClusterControlPlaneDatabaseEncryption{
		KmsKeyArn: dcl.String(obj["kms_key_arn"].(string)),
	}
}

func flattenGkemulticloudAwsClusterControlPlaneDatabaseEncryption(obj *gkemulticloud.AwsClusterControlPlaneDatabaseEncryption) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"kms_key_arn": obj.KmsKeyArn,
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAwsClusterControlPlaneMainVolume(o interface{}) *gkemulticloud.AwsClusterControlPlaneMainVolume {
	if o == nil {
		return gkemulticloud.EmptyAwsClusterControlPlaneMainVolume
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAwsClusterControlPlaneMainVolume
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AwsClusterControlPlaneMainVolume{
		Iops:       dcl.Int64(int64(obj["iops"].(int))),
		KmsKeyArn:  dcl.String(obj["kms_key_arn"].(string)),
		SizeGib:    dcl.Int64(int64(obj["size_gib"].(int))),
		VolumeType: gkemulticloud.AwsClusterControlPlaneMainVolumeVolumeTypeEnumRef(obj["volume_type"].(string)),
	}
}

func flattenGkemulticloudAwsClusterControlPlaneMainVolume(obj *gkemulticloud.AwsClusterControlPlaneMainVolume) interface{} {
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

func expandGkemulticloudAwsClusterControlPlaneRootVolume(o interface{}) *gkemulticloud.AwsClusterControlPlaneRootVolume {
	if o == nil {
		return gkemulticloud.EmptyAwsClusterControlPlaneRootVolume
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAwsClusterControlPlaneRootVolume
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AwsClusterControlPlaneRootVolume{
		Iops:       dcl.Int64(int64(obj["iops"].(int))),
		KmsKeyArn:  dcl.String(obj["kms_key_arn"].(string)),
		SizeGib:    dcl.Int64(int64(obj["size_gib"].(int))),
		VolumeType: gkemulticloud.AwsClusterControlPlaneRootVolumeVolumeTypeEnumRef(obj["volume_type"].(string)),
	}
}

func flattenGkemulticloudAwsClusterControlPlaneRootVolume(obj *gkemulticloud.AwsClusterControlPlaneRootVolume) interface{} {
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

func expandGkemulticloudAwsClusterControlPlaneSshConfig(o interface{}) *gkemulticloud.AwsClusterControlPlaneSshConfig {
	if o == nil {
		return gkemulticloud.EmptyAwsClusterControlPlaneSshConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAwsClusterControlPlaneSshConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AwsClusterControlPlaneSshConfig{
		Ec2KeyPair: dcl.String(obj["ec2_key_pair"].(string)),
	}
}

func flattenGkemulticloudAwsClusterControlPlaneSshConfig(obj *gkemulticloud.AwsClusterControlPlaneSshConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"ec2_key_pair": obj.Ec2KeyPair,
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAwsClusterNetworking(o interface{}) *gkemulticloud.AwsClusterNetworking {
	if o == nil {
		return gkemulticloud.EmptyAwsClusterNetworking
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAwsClusterNetworking
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AwsClusterNetworking{
		PodAddressCidrBlocks:         expandStringArray(obj["pod_address_cidr_blocks"]),
		ServiceAddressCidrBlocks:     expandStringArray(obj["service_address_cidr_blocks"]),
		ServiceLoadBalancerSubnetIds: expandStringArray(obj["service_load_balancer_subnet_ids"]),
		VPCId:                        dcl.String(obj["vpc_id"].(string)),
	}
}

func flattenGkemulticloudAwsClusterNetworking(obj *gkemulticloud.AwsClusterNetworking) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"pod_address_cidr_blocks":          obj.PodAddressCidrBlocks,
		"service_address_cidr_blocks":      obj.ServiceAddressCidrBlocks,
		"service_load_balancer_subnet_ids": obj.ServiceLoadBalancerSubnetIds,
		"vpc_id":                           obj.VPCId,
	}

	return []interface{}{transformed}

}

func flattenGkemulticloudAwsClusterWorkloadIdentityConfig(obj *gkemulticloud.AwsClusterWorkloadIdentityConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"identity_provider": obj.IdentityProvider,
		"issuer_uri":        obj.IssuerUri,
		"workload_pool":     obj.WorkloadPool,
	}

	return []interface{}{transformed}

}
