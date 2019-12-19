// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func init() {
	resource.AddTestSweepers("ComputeVpnGateway", &resource.Sweeper{
		Name: "ComputeVpnGateway",
		F:    testSweepComputeVpnGateway,
	})
}

func testSweepComputeVpnGateway(region string) error {
	resource_name := "ComputeVpnGateway"
	log.Printf("[SWEEP_LOG] Sweeping %s", resource_name)

	config, err := sharedConfigForRegion(region)
	if err != nil {
		log.Fatalf("[SWEEP_LOG] error getting shared config for region: %s", err)
	}

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		log.Fatalf("[SWEEP_LOG] error loading: %s", err)
	}

	list_template := strings.Split("https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/targetVpnGateways", "?")[0]

	d := &ResourceDataMock{
		FieldsInSchema: map[string]interface{}{
			"project":  config.Project,
			"region":   region,
			"location": region,
			"zone":     "-",
		},
	}

	list_url, err := replaceVars(d, config, list_template)

	if err != nil {
		log.Printf("[SWEEP_LOG] error preparing sweeper list url: %s", err)
		return nil
	}

	if strings.Count(list_url, "//") > 1 {
		log.Printf("[SWEEP_LOG] Invalid list url for %s sweeper: %s", resource_name, list_url)
		return nil
	}

	res, err := sendRequest(config, "GET", config.Project, list_url, nil)
	if err != nil {
		log.Printf("[SWEEP_LOG] Unable to list %s: %s", resource_name, err)
		return nil
	}

	resource_list, ok := res["items"]
	if !ok {
		log.Printf("[SWEEP_LOG] Nothing found in response.")
		return nil
	}

	rl := resource_list.([]interface{})

	log.Printf("[SWEEP_LOG] Found %d items in %s list response.", len(rl), resource_name)
	non_prefix_count := 0
	for _, ri := range rl {
		r := ri.(map[string]interface{})
		if r["name"] == nil {
			log.Printf("[SWEEP_LOG] %s resource name was nil", resource_name)
			return nil
		}

		name_segs := strings.Split(r["name"].(string), "/")
		name := name_segs[len(name_segs)-1]

		// Only sweep resources with the test prefix
		if strings.HasPrefix(name, "tf-test") {
			delete_template := "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}"

			delete_url, err := replaceVars(d, config, delete_template)
			if err != nil {
				log.Printf("[SWEEP_LOG] error preparing delete url: %s", err)
				return nil
			}
			delete_url = delete_url + name

			_, err = sendRequest(config, "DELETE", config.Project, delete_url, nil)
			if err != nil {
				log.Printf("[SWEEP_LOG] Error deleting for url %s : %s", delete_url, err)
			} else {
				log.Printf("[SWEEP_LOG] Deleted %s resource: %s", resource_name, name)
			}
		} else {
			non_prefix_count++
		}
	}

	if non_prefix_count > 0 {
		log.Printf("[SWEEP_LOG] %d items without tf_test prefix remain.", non_prefix_count)
	}

	return nil
}
