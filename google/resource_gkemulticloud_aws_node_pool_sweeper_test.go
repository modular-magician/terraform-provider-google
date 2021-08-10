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
	resource.AddTestSweepers("GkemulticloudAws_node_pool", &resource.Sweeper{
		Name: "GkemulticloudAws_node_pool",
		F:    testSweepGkemulticloudAws_node_pool,
	})
}

func testSweepGkemulticloudAws_node_pool(region string) error {
	resourceName := "GkemulticloudAws_node_pool"
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
	err = client.DeleteAllAwsNodePool(context.Background(), d["project"], d["location"], d["awscluster"], isDeletableGkemulticloudAws_node_pool)
	if err != nil {
		return err
	}
	return nil
}

func isDeletableGkemulticloudAws_node_pool(r *gkemulticloud.AwsNodePool) bool {
	return isSweepableTestResource(*r.Name)
}
