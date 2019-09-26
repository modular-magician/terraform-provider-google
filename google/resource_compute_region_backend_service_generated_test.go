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
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccComputeRegionBackendService_regionBackendServiceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeRegionBackendServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_regionBackendServiceBasicExample(context),
			},
			{
				ResourceName:      "google_compute_region_backend_service.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionBackendService_regionBackendServiceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_backend_service" "default" {
  name                            = "region-backend-service%{random_suffix}"
  region                          = "us-central1"
  health_checks                   = ["${google_compute_health_check.default.self_link}"]
  connection_draining_timeout_sec = 10
  session_affinity                = "CLIENT_IP"
}

resource "google_compute_health_check" "default" {
  name               = "health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}
`, context)
}

func testAccCheckComputeRegionBackendServiceDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_region_backend_service" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeRegionBackendService still exists at %s", url)
		}
	}

	return nil
}
