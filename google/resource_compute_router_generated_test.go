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

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccComputeRouter_routerBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeRouterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRouter_routerBasicExample(context),
			},
			{
				ResourceName:      "google_compute_router.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRouter_routerBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_router" "foobar" {
  name    = "my-router-%{random_suffix}"
  network = "${google_compute_network.foobar.name}"
  bgp {
    asn               = 64514
    advertise_mode    = "CUSTOM"
    advertised_groups = ["ALL_SUBNETS"]
    advertised_ip_ranges {
      range = "1.2.3.4"
    }
    advertised_ip_ranges {
      range = "6.7.0.0/16"
    }
  }
}

resource "google_compute_network" "foobar" {
  name = "my-network-%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func testAccCheckComputeRouterDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_router" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/compute/v1/projects/{{project}}/regions/{{region}}/routers/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeRouter still exists at %s", url)
		}
	}

	return nil
}
