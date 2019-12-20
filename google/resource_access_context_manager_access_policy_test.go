package google

import (
	"context"
	"fmt"
	"log"
	neturl "net/url"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func init() {
	resource.AddTestSweepers("gcp_access_context_manager_policy", &resource.Sweeper{
		Name: "gcp_access_context_manager_policy",
		F:    testSweepAccessContextManagerPolicies,
	})
}

func testSweepAccessContextManagerPolicies(region string) error {
	config, err := sharedConfigForRegion(region)
	if err != nil {
		log.Fatalf("error getting shared config for region %q: %s", region, err)
	}

	err := config.LoadAndValidate(context.Background())
	if err != nil {
		log.Fatalf("error loading and validating shared config for region %q: %s", region, err)
	}

	log.Printf("[DEBUG] Listing Access Policies for org %q", testOrg)

	parent := neturl.QueryEscape(fmt.Sprintf("organizations/%s", testOrg))
	listUrl := fmt.Sprintf("%saccessPolicies?parent=%s", config.AccessContextManagerBasePath, parent)

	resp, err := sendRequest(config, "GET", "", listUrl, nil)
	if err != nil && !isGoogleApiErrorWithCode(err, 404) {
		return fmt.Errorf("unable to list AccessPolicies for organization %q: %v", testOrg, err)
	}
	var policies []interface{}
	if resp != nil {
		if v, ok := resp["accessPolicies"]; ok {
			policies = v.([]interface{})
		}
	}

	if len(policies) == 0 {
		log.Printf("[DEBUG] no access policies found, exiting sweeper")
		return nil
	}
	if len(policies) > 1 {
		return fmt.Errorf("unexpected - more than one access policies found, change the tests")
	}

	policy := policies[0].(map[string]interface{})
	log.Printf("[DEBUG] Deleting test Access Policies %q", policy["name"])

	policyUrl := config.AccessContextManagerBasePath + policy["name"].(string)
	if _, err := sendRequest(config, "DELETE", "", policyUrl, nil); err != nil && !isGoogleApiErrorWithCode(err, 404) {
		return fmt.Errorf("unable to delete access policy %q", policy["name"].(string))
	}

	return nil
}

// Since each test here is acting on the same organization and only one AccessPolicy
// can exist, they need to be ran serially
func TestAccAccessContextManager(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		"access_policy":            testAccAccessContextManagerAccessPolicy_basicTest,
		"service_perimeter":        testAccAccessContextManagerServicePerimeter_basicTest,
		"service_perimeter_update": testAccAccessContextManagerServicePerimeter_updateTest,
		"access_level":             testAccAccessContextManagerAccessLevel_basicTest,
		"access_level_full":        testAccAccessContextManagerAccessLevel_fullTest,
	}

	for name, tc := range testCases {
		// shadow the tc variable into scope so that when
		// the loop continues, if t.Run hasn't executed tc(t)
		// yet, we don't have a race condition
		// see https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tc := tc
		t.Run(name, func(t *testing.T) {
			tc(t)
		})
	}
}

func testAccAccessContextManagerAccessPolicy_basicTest(t *testing.T) {
	org := getTestOrgFromEnv(t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccessContextManagerAccessPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessContextManagerAccessPolicy_basic(org, "my policy"),
			},
			{
				ResourceName:      "google_access_context_manager_access_policy.test-access",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAccessContextManagerAccessPolicy_basic(org, "my new policy"),
			},
			{
				ResourceName:      "google_access_context_manager_access_policy.test-access",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckAccessContextManagerAccessPolicyDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "google_access_context_manager_access_policy" {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{AccessContextManagerBasePath}}accessPolicies/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("AccessPolicy still exists at %s", url)
		}
	}

	return nil
}

func testAccAccessContextManagerAccessPolicy_basic(org, title string) string {
	return fmt.Sprintf(`
resource "google_access_context_manager_access_policy" "test-access" {
  parent = "organizations/%s"
  title  = "%s"
}
`, org, title)
}
