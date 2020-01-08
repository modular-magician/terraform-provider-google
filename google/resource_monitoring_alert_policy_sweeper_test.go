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
	resource.AddTestSweepers("MonitoringAlertPolicy", &resource.Sweeper{
		Name: "MonitoringAlertPolicy",
		F:    testSweepMonitoringAlertPolicy,
	})
}

// At the time of writing, the CI only passes us-central1 as the region
func testSweepMonitoringAlertPolicy(region string) error {
	resourceName := "MonitoringAlertPolicy"
	log.Printf("[INFO] Sweeping %s", resourceName)

	config, err := sharedConfigForRegion(region)
	if err != nil {
		log.Printf("[INFO] error getting shared config for region: %s", err)
		return err
	}

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		log.Printf("[INFO] error loading: %s", err)
		return err
	}

	listTemplate := strings.Split("https://monitoring.googleapis.com/v3/projects/{{project}}/alertPolicies", "?")[0]

	d := &ResourceDataMock{
		FieldsInSchema: map[string]interface{}{
			"project":  config.Project,
			"region":   region,
			"location": region,
			"zone":     "-",
		},
	}

	listUrl, err := replaceVars(d, config, listTemplate)
	if err != nil {
		log.Printf("[INFO] error preparing sweeper list url: %s", err)
		return nil
	}

	if strings.Count(listUrl, "//") > 1 {
		log.Printf("[INFO] Invalid list url for %s sweeper: %s", resourceName, listUrl)
		return nil
	}

	res, err := sendRequest(config, "GET", config.Project, listUrl, nil)
	if err != nil {
		log.Printf("[INFO] Unable to list resource %s (url %s ): %s", resourceName, listUrl, err)
		return nil
	}

	resourceList, ok := res["alertPolicies"]
	if !ok {
		log.Printf("[INFO] Nothing found in response.")
		return nil
	}

	rl := resourceList.([]map[string]interface{})

	log.Printf("[INFO] Found %d items in %s list response.", len(rl), resourceName)
	// items who don't match the tf-test prefix
	nonPrefixCount := 0
	for _, obj := range rl {
		if obj["name"] == nil {
			log.Printf("[INFO] %s resource name was nil", resourceName)
			return nil
		}

		nameSegs := strings.Split(obj["name"].(string), "/")
		name := nameSegs[len(nameSegs)-1]

		// Only sweep resources with the test prefix
		if !strings.HasPrefix(name, "tf-test") {
			nonPrefixCount++
			continue
		}
		deleteTemplate := "https://monitoring.googleapis.com/v3/{{name}}"
		deleteUrl, err := replaceVars(d, config, deleteTemplate)
		if err != nil {
			log.Printf("[INFO] error preparing delete url: %s", err)
			return nil
		}
		deleteUrl = deleteUrl + name
		// Don't wait on operations as we may have a lot to delete
		_, err = sendRequest(config, "DELETE", config.Project, deleteUrl, nil)
		if err != nil {
			log.Printf("[INFO] Error deleting for url %s : %s", deleteUrl, err)
		} else {
			log.Printf("[INFO] Sent delete request for %s resource: %s", resourceName, name)
		}
	}

	if nonPrefixCount > 0 {
		log.Printf("[INFO] %d items without tf_test prefix remain.", nonPrefixCount)
	}

	return nil
}
