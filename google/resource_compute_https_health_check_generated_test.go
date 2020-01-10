// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//     ***     DIFF TEST DIFF TEST    ***
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
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccComputeHttpsHealthCheck_httpsHealthCheckBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeHttpsHealthCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHttpsHealthCheck_httpsHealthCheckBasicExample(context),
			},
			{
				ResourceName:      "google_compute_https_health_check.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHttpsHealthCheck_httpsHealthCheckBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_https_health_check" "default" {
  name         = "authentication-health-check%{random_suffix}"
  request_path = "/health_check"

  timeout_sec        = 1
  check_interval_sec = 1
}
`, context)
}

func testAccCheckComputeHttpsHealthCheckDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_https_health_check" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/httpsHealthChecks/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeHttpsHealthCheck still exists at %s", url)
		}
	}

	return nil
}
