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
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccComputeNetworkEndpointGroup_networkEndpointGroupExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeNetworkEndpointGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNetworkEndpointGroup_networkEndpointGroupExample(context),
			},
			{
				ResourceName:      "google_compute_network_endpoint_group.neg",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeNetworkEndpointGroup_networkEndpointGroupExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network_endpoint_group" "neg" {
  name         = "tf-test-my-lb-neg%{random_suffix}"
  network      = google_compute_network.default.self_link
  subnetwork   = google_compute_subnetwork.default.self_link
  default_port = "90"
  zone         = "us-central1-a"
}

resource "google_compute_network" "default" {
  name                    = "tf-test-neg-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "tf-test-neg-subnetwork%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.default.self_link
}
`, context)
}

func testAccCheckComputeNetworkEndpointGroupDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_network_endpoint_group" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeNetworkEndpointGroup still exists at %s", url)
		}
	}

	return nil
}
