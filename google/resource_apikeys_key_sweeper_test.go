package google

import (
	"context"
	"log"
	"testing"

	apikeys "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apikeys"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("ApikeysKey", &resource.Sweeper{
		Name: "ApikeysKey",
		F:    testSweepApikeysKey,
	})
}

func testSweepApikeysKey(region string) error {
	resourceName := "ApikeysKey"
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

	err = config.clientApikeysDCL.DeleteAllKey(context.Background(), d["project"], isDeletableApikeysKey)
	if err != nil {
		return err
	}
	return nil
}

func isDeletableApikeysKey(r *apikeys.Key) bool {
	return isSweepableTestResource(*r.Name)
}
