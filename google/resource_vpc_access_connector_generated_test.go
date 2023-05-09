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

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccVPCAccessConnector_vpcAccessConnectorExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVPCAccessConnectorDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCAccessConnector_vpcAccessConnectorExample(context),
			},
			{
				ResourceName:            "google_vpc_access_connector.connector",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"self_link", "region"},
			},
		},
	})
}

func testAccVPCAccessConnector_vpcAccessConnectorExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vpc_access_connector" "connector" {
  name          = "tf-test-vpc-con%{random_suffix}"
  ip_cidr_range = "10.8.0.0/28"
  network       = "default"
}
`, context)
}

func TestAccVPCAccessConnector_vpcAccessConnectorSharedVpcExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVPCAccessConnectorDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCAccessConnector_vpcAccessConnectorSharedVpcExample(context),
			},
			{
				ResourceName:            "google_vpc_access_connector.connector",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"self_link", "region"},
			},
		},
	})
}

func testAccVPCAccessConnector_vpcAccessConnectorSharedVpcExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_vpc_access_connector" "connector" {
  name          = "tf-test-vpc-con%{random_suffix}"
  subnet {
    name = google_compute_subnetwork.custom_test.name
  }
  machine_type = "e2-standard-4"
}

resource "google_compute_subnetwork" "custom_test" {
  name          = "tf-test-vpc-con%{random_suffix}"
  ip_cidr_range = "10.2.0.0/28"
  region        = "us-central1"
  network       = google_compute_network.custom_test.id
}

resource "google_compute_network" "custom_test" {
  name                    = "tf-test-vpc-con%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func testAccCheckVPCAccessConnectorDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_vpc_access_connector" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{VPCAccessBasePath}}projects/{{project}}/locations/{{region}}/connectors/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("VPCAccessConnector still exists at %s", url)
			}
		}

		return nil
	}
}
