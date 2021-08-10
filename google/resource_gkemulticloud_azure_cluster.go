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

func resourceGkemulticloudAzureCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceGkemulticloudAzureClusterCreate,
		Read:   resourceGkemulticloudAzureClusterRead,
		Delete: resourceGkemulticloudAzureClusterDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGkemulticloudAzureClusterImport,
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
				Elem:        GkemulticloudAzureClusterAuthorizationSchema(),
			},

			"azure_client": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"azure_region": {
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
				Elem:        GkemulticloudAzureClusterControlPlaneSchema(),
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
				Elem:        GkemulticloudAzureClusterNetworkingSchema(),
			},

			"resource_group_id": {
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
				Elem:        GkemulticloudAzureClusterWorkloadIdentityConfigSchema(),
			},
		},
	}
}

func GkemulticloudAzureClusterAuthorizationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_users": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        GkemulticloudAzureClusterAuthorizationAdminUsersSchema(),
			},
		},
	}
}

func GkemulticloudAzureClusterAuthorizationAdminUsersSchema() *schema.Resource {
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

func GkemulticloudAzureClusterControlPlaneSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ssh_config": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAzureClusterControlPlaneSshConfigSchema(),
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

			"database_encryption": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAzureClusterControlPlaneDatabaseEncryptionSchema(),
			},

			"main_volume": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAzureClusterControlPlaneMainVolumeSchema(),
			},

			"root_volume": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkemulticloudAzureClusterControlPlaneRootVolumeSchema(),
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

func GkemulticloudAzureClusterControlPlaneSshConfigSchema() *schema.Resource {
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

func GkemulticloudAzureClusterControlPlaneDatabaseEncryptionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kms_key_identifier": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"resource_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func GkemulticloudAzureClusterControlPlaneMainVolumeSchema() *schema.Resource {
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

func GkemulticloudAzureClusterControlPlaneRootVolumeSchema() *schema.Resource {
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

func GkemulticloudAzureClusterNetworkingSchema() *schema.Resource {
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

			"virtual_network_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func GkemulticloudAzureClusterWorkloadIdentityConfigSchema() *schema.Resource {
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

func resourceGkemulticloudAzureClusterCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkemulticloud.AzureCluster{
		Authorization:   expandGkemulticloudAzureClusterAuthorization(d.Get("authorization")),
		AzureClient:     dcl.String(d.Get("azure_client").(string)),
		AzureRegion:     dcl.String(d.Get("azure_region").(string)),
		ControlPlane:    expandGkemulticloudAzureClusterControlPlane(d.Get("control_plane")),
		Location:        dcl.String(d.Get("location").(string)),
		Name:            dcl.String(d.Get("name").(string)),
		Networking:      expandGkemulticloudAzureClusterNetworking(d.Get("networking")),
		ResourceGroupId: dcl.String(d.Get("resource_group_id").(string)),
		Annotations:     checkStringMap(d.Get("annotations")),
		Description:     dcl.String(d.Get("description").(string)),
		Project:         dcl.String(project),
	}

	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/azureClusters/{{name}}")
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
	res, err := client.ApplyAzureCluster(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating AzureCluster: %s", err)
	}

	log.Printf("[DEBUG] Finished creating AzureCluster %q: %#v", d.Id(), res)

	return resourceGkemulticloudAzureClusterRead(d, meta)
}

func resourceGkemulticloudAzureClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkemulticloud.AzureCluster{
		Authorization:   expandGkemulticloudAzureClusterAuthorization(d.Get("authorization")),
		AzureClient:     dcl.String(d.Get("azure_client").(string)),
		AzureRegion:     dcl.String(d.Get("azure_region").(string)),
		ControlPlane:    expandGkemulticloudAzureClusterControlPlane(d.Get("control_plane")),
		Location:        dcl.String(d.Get("location").(string)),
		Name:            dcl.String(d.Get("name").(string)),
		Networking:      expandGkemulticloudAzureClusterNetworking(d.Get("networking")),
		ResourceGroupId: dcl.String(d.Get("resource_group_id").(string)),
		Annotations:     checkStringMap(d.Get("annotations")),
		Description:     dcl.String(d.Get("description").(string)),
		Project:         dcl.String(project),
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
	res, err := client.GetAzureCluster(context.Background(), obj)
	if err != nil {
		// Resource not found
		d.SetId("")
		return err
	}

	if err = d.Set("authorization", flattenGkemulticloudAzureClusterAuthorization(res.Authorization)); err != nil {
		return fmt.Errorf("error setting authorization in state: %s", err)
	}
	if err = d.Set("azure_client", res.AzureClient); err != nil {
		return fmt.Errorf("error setting azure_client in state: %s", err)
	}
	if err = d.Set("azure_region", res.AzureRegion); err != nil {
		return fmt.Errorf("error setting azure_region in state: %s", err)
	}
	if err = d.Set("control_plane", flattenGkemulticloudAzureClusterControlPlane(res.ControlPlane)); err != nil {
		return fmt.Errorf("error setting control_plane in state: %s", err)
	}
	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("networking", flattenGkemulticloudAzureClusterNetworking(res.Networking)); err != nil {
		return fmt.Errorf("error setting networking in state: %s", err)
	}
	if err = d.Set("resource_group_id", res.ResourceGroupId); err != nil {
		return fmt.Errorf("error setting resource_group_id in state: %s", err)
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
	if err = d.Set("workload_identity_config", flattenGkemulticloudAzureClusterWorkloadIdentityConfig(res.WorkloadIdentityConfig)); err != nil {
		return fmt.Errorf("error setting workload_identity_config in state: %s", err)
	}

	return nil
}

func resourceGkemulticloudAzureClusterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkemulticloud.AzureCluster{
		Authorization:   expandGkemulticloudAzureClusterAuthorization(d.Get("authorization")),
		AzureClient:     dcl.String(d.Get("azure_client").(string)),
		AzureRegion:     dcl.String(d.Get("azure_region").(string)),
		ControlPlane:    expandGkemulticloudAzureClusterControlPlane(d.Get("control_plane")),
		Location:        dcl.String(d.Get("location").(string)),
		Name:            dcl.String(d.Get("name").(string)),
		Networking:      expandGkemulticloudAzureClusterNetworking(d.Get("networking")),
		ResourceGroupId: dcl.String(d.Get("resource_group_id").(string)),
		Annotations:     checkStringMap(d.Get("annotations")),
		Description:     dcl.String(d.Get("description").(string)),
		Project:         dcl.String(project),
	}

	log.Printf("[DEBUG] Deleting AzureCluster %q", d.Id())
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
	if err := client.DeleteAzureCluster(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting AzureCluster: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting AzureCluster %q", d.Id())
	return nil
}

func resourceGkemulticloudAzureClusterImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/azureClusters/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/azureClusters/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandGkemulticloudAzureClusterAuthorization(o interface{}) *gkemulticloud.AzureClusterAuthorization {
	if o == nil {
		return gkemulticloud.EmptyAzureClusterAuthorization
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAzureClusterAuthorization
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AzureClusterAuthorization{
		AdminUsers: expandGkemulticloudAzureClusterAuthorizationAdminUsersArray(obj["admin_users"]),
	}
}

func flattenGkemulticloudAzureClusterAuthorization(obj *gkemulticloud.AzureClusterAuthorization) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"admin_users": flattenGkemulticloudAzureClusterAuthorizationAdminUsersArray(obj.AdminUsers),
	}

	return []interface{}{transformed}

}
func expandGkemulticloudAzureClusterAuthorizationAdminUsersArray(o interface{}) []gkemulticloud.AzureClusterAuthorizationAdminUsers {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]gkemulticloud.AzureClusterAuthorizationAdminUsers, 0, len(objs))
	for _, item := range objs {
		i := expandGkemulticloudAzureClusterAuthorizationAdminUsers(item)
		items = append(items, *i)
	}

	return items
}

func expandGkemulticloudAzureClusterAuthorizationAdminUsers(o interface{}) *gkemulticloud.AzureClusterAuthorizationAdminUsers {
	if o == nil {
		return gkemulticloud.EmptyAzureClusterAuthorizationAdminUsers
	}

	obj := o.(map[string]interface{})
	return &gkemulticloud.AzureClusterAuthorizationAdminUsers{
		Username: dcl.String(obj["username"].(string)),
	}
}

func flattenGkemulticloudAzureClusterAuthorizationAdminUsersArray(objs []gkemulticloud.AzureClusterAuthorizationAdminUsers) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenGkemulticloudAzureClusterAuthorizationAdminUsers(&item)
		items = append(items, i)
	}

	return items
}

func flattenGkemulticloudAzureClusterAuthorizationAdminUsers(obj *gkemulticloud.AzureClusterAuthorizationAdminUsers) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"username": obj.Username,
	}

	return transformed

}

func expandGkemulticloudAzureClusterControlPlane(o interface{}) *gkemulticloud.AzureClusterControlPlane {
	if o == nil {
		return gkemulticloud.EmptyAzureClusterControlPlane
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAzureClusterControlPlane
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AzureClusterControlPlane{
		SshConfig:          expandGkemulticloudAzureClusterControlPlaneSshConfig(obj["ssh_config"]),
		SubnetId:           dcl.String(obj["subnet_id"].(string)),
		Version:            dcl.String(obj["version"].(string)),
		DatabaseEncryption: expandGkemulticloudAzureClusterControlPlaneDatabaseEncryption(obj["database_encryption"]),
		MainVolume:         expandGkemulticloudAzureClusterControlPlaneMainVolume(obj["main_volume"]),
		RootVolume:         expandGkemulticloudAzureClusterControlPlaneRootVolume(obj["root_volume"]),
		Tags:               checkStringMap(obj["tags"]),
		VmSize:             dcl.String(obj["vm_size"].(string)),
	}
}

func flattenGkemulticloudAzureClusterControlPlane(obj *gkemulticloud.AzureClusterControlPlane) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"ssh_config":          flattenGkemulticloudAzureClusterControlPlaneSshConfig(obj.SshConfig),
		"subnet_id":           obj.SubnetId,
		"version":             obj.Version,
		"database_encryption": flattenGkemulticloudAzureClusterControlPlaneDatabaseEncryption(obj.DatabaseEncryption),
		"main_volume":         flattenGkemulticloudAzureClusterControlPlaneMainVolume(obj.MainVolume),
		"root_volume":         flattenGkemulticloudAzureClusterControlPlaneRootVolume(obj.RootVolume),
		"tags":                obj.Tags,
		"vm_size":             obj.VmSize,
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAzureClusterControlPlaneSshConfig(o interface{}) *gkemulticloud.AzureClusterControlPlaneSshConfig {
	if o == nil {
		return gkemulticloud.EmptyAzureClusterControlPlaneSshConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAzureClusterControlPlaneSshConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AzureClusterControlPlaneSshConfig{
		AuthorizedKey: dcl.String(obj["authorized_key"].(string)),
	}
}

func flattenGkemulticloudAzureClusterControlPlaneSshConfig(obj *gkemulticloud.AzureClusterControlPlaneSshConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"authorized_key": obj.AuthorizedKey,
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAzureClusterControlPlaneDatabaseEncryption(o interface{}) *gkemulticloud.AzureClusterControlPlaneDatabaseEncryption {
	if o == nil {
		return gkemulticloud.EmptyAzureClusterControlPlaneDatabaseEncryption
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAzureClusterControlPlaneDatabaseEncryption
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AzureClusterControlPlaneDatabaseEncryption{
		KmsKeyIdentifier: dcl.String(obj["kms_key_identifier"].(string)),
		ResourceGroupId:  dcl.String(obj["resource_group_id"].(string)),
	}
}

func flattenGkemulticloudAzureClusterControlPlaneDatabaseEncryption(obj *gkemulticloud.AzureClusterControlPlaneDatabaseEncryption) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"kms_key_identifier": obj.KmsKeyIdentifier,
		"resource_group_id":  obj.ResourceGroupId,
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAzureClusterControlPlaneMainVolume(o interface{}) *gkemulticloud.AzureClusterControlPlaneMainVolume {
	if o == nil {
		return gkemulticloud.EmptyAzureClusterControlPlaneMainVolume
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAzureClusterControlPlaneMainVolume
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AzureClusterControlPlaneMainVolume{
		SizeGib: dcl.Int64(int64(obj["size_gib"].(int))),
	}
}

func flattenGkemulticloudAzureClusterControlPlaneMainVolume(obj *gkemulticloud.AzureClusterControlPlaneMainVolume) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"size_gib": obj.SizeGib,
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAzureClusterControlPlaneRootVolume(o interface{}) *gkemulticloud.AzureClusterControlPlaneRootVolume {
	if o == nil {
		return gkemulticloud.EmptyAzureClusterControlPlaneRootVolume
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAzureClusterControlPlaneRootVolume
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AzureClusterControlPlaneRootVolume{
		SizeGib: dcl.Int64(int64(obj["size_gib"].(int))),
	}
}

func flattenGkemulticloudAzureClusterControlPlaneRootVolume(obj *gkemulticloud.AzureClusterControlPlaneRootVolume) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"size_gib": obj.SizeGib,
	}

	return []interface{}{transformed}

}

func expandGkemulticloudAzureClusterNetworking(o interface{}) *gkemulticloud.AzureClusterNetworking {
	if o == nil {
		return gkemulticloud.EmptyAzureClusterNetworking
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkemulticloud.EmptyAzureClusterNetworking
	}
	obj := objArr[0].(map[string]interface{})
	return &gkemulticloud.AzureClusterNetworking{
		PodAddressCidrBlocks:     expandStringArray(obj["pod_address_cidr_blocks"]),
		ServiceAddressCidrBlocks: expandStringArray(obj["service_address_cidr_blocks"]),
		VirtualNetworkId:         dcl.String(obj["virtual_network_id"].(string)),
	}
}

func flattenGkemulticloudAzureClusterNetworking(obj *gkemulticloud.AzureClusterNetworking) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"pod_address_cidr_blocks":     obj.PodAddressCidrBlocks,
		"service_address_cidr_blocks": obj.ServiceAddressCidrBlocks,
		"virtual_network_id":          obj.VirtualNetworkId,
	}

	return []interface{}{transformed}

}

func flattenGkemulticloudAzureClusterWorkloadIdentityConfig(obj *gkemulticloud.AzureClusterWorkloadIdentityConfig) interface{} {
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
