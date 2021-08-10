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
	"log"
	"testing"

	gkemulticloud "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkemulticloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("GkemulticloudAzure_client", &resource.Sweeper{
		Name: "GkemulticloudAzure_client",
		F:    testSweepGkemulticloudAzure_client,
	})
}

func testSweepGkemulticloudAzure_client(region string) error {
	resourceName := "GkemulticloudAzure_client"
	log.Printf("[INFO][SWEEPER_LOG] Starting sweeper for %s", resourceName)

	config, err := sharedConfigForRegion(region)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error getting shared config for region: %s", err)
		return err
	}

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error loading: %s", err)
		return err
	}

	t := &testing.T{}
	billingId := getTestBillingAccountFromEnv(t)

	// Setup variables to be used for Delete arguments.
	d := map[string]string{
		"project":         config.Project,
		"region":          region,
		"location":        region,
		"zone":            "-",
		"billing_account": billingId,
	}

	client := NewDCLGkemulticloudClient(config, config.userAgent, "")
	err = client.DeleteAllAzureClient(context.Background(), d["project"], d["location"], isDeletableGkemulticloudAzure_client)
	if err != nil {
		return err
	}
	return nil
}

func isDeletableGkemulticloudAzure_client(r *gkemulticloud.AzureClient) bool {
	return isSweepableTestResource(*r.Name)
}
