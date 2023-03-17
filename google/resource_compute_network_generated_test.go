// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccComputeNetwork_networkBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeNetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNetwork_networkBasicExample(context),
			},
			{
				ResourceName:      "google_compute_network.vpc_network",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeNetwork_networkBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network" "vpc_network" {
  name = "tf-test-vpc-network%{random_suffix}"
}
`, context)
}

func TestAccComputeNetwork_networkCustomMtuExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeNetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNetwork_networkCustomMtuExample(context),
			},
			{
				ResourceName:      "google_compute_network.vpc_network",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeNetwork_networkCustomMtuExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network" "vpc_network" {
  project                 = "%{project}"
  name                    = "tf-test-vpc-network%{random_suffix}"
  auto_create_subnetworks = true
  mtu                     = 1460
}
`, context)
}

func testAccCheckComputeNetworkDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_network" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/networks/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeNetwork still exists at %s", url)
			}
		}

		return nil
	}
}
