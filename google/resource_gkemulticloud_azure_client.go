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

func resourceGkemulticloudAzureClient() *schema.Resource {
	return &schema.Resource{
		Create: resourceGkemulticloudAzureClientCreate,
		Read:   resourceGkemulticloudAzureClientRead,
		Delete: resourceGkemulticloudAzureClientDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGkemulticloudAzureClientImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
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

			"tenant_id": {
				Type:        schema.TypeString,
				Required:    true,
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

			"certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},
		},
	}
}

func resourceGkemulticloudAzureClientCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkemulticloud.AzureClient{
		ApplicationId: dcl.String(d.Get("application_id").(string)),
		Location:      dcl.String(d.Get("location").(string)),
		Name:          dcl.String(d.Get("name").(string)),
		TenantId:      dcl.String(d.Get("tenant_id").(string)),
		Project:       dcl.String(project),
	}

	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/azureClients/{{name}}")
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
	res, err := client.ApplyAzureClient(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating AzureClient: %s", err)
	}

	log.Printf("[DEBUG] Finished creating AzureClient %q: %#v", d.Id(), res)

	return resourceGkemulticloudAzureClientRead(d, meta)
}

func resourceGkemulticloudAzureClientRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkemulticloud.AzureClient{
		ApplicationId: dcl.String(d.Get("application_id").(string)),
		Location:      dcl.String(d.Get("location").(string)),
		Name:          dcl.String(d.Get("name").(string)),
		TenantId:      dcl.String(d.Get("tenant_id").(string)),
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
	res, err := client.GetAzureClient(context.Background(), obj)
	if err != nil {
		// Resource not found
		d.SetId("")
		return err
	}

	if err = d.Set("application_id", res.ApplicationId); err != nil {
		return fmt.Errorf("error setting application_id in state: %s", err)
	}
	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("tenant_id", res.TenantId); err != nil {
		return fmt.Errorf("error setting tenant_id in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("certificate", res.Certificate); err != nil {
		return fmt.Errorf("error setting certificate in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("uid", res.Uid); err != nil {
		return fmt.Errorf("error setting uid in state: %s", err)
	}

	return nil
}

func resourceGkemulticloudAzureClientDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkemulticloud.AzureClient{
		ApplicationId: dcl.String(d.Get("application_id").(string)),
		Location:      dcl.String(d.Get("location").(string)),
		Name:          dcl.String(d.Get("name").(string)),
		TenantId:      dcl.String(d.Get("tenant_id").(string)),
		Project:       dcl.String(project),
	}

	log.Printf("[DEBUG] Deleting AzureClient %q", d.Id())
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
	if err := client.DeleteAzureClient(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting AzureClient: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting AzureClient %q", d.Id())
	return nil
}

func resourceGkemulticloudAzureClientImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/azureClients/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/azureClients/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}
