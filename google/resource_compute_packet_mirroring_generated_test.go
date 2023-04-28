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

	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccComputePacketMirroring_computePacketMirroringFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputePacketMirroringDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputePacketMirroring_computePacketMirroringFullExample(context),
			},
			{
				ResourceName:      "google_compute_packet_mirroring.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputePacketMirroring_computePacketMirroringFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "mirror" {
  name = "tf-test-my-instance%{random_suffix}"
  machine_type = "e2-medium"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = google_compute_network.default.id
    access_config {
    }
  }
}

resource "google_compute_network" "default" {
  name = "tf-test-my-network%{random_suffix}"
}

resource "google_compute_subnetwork" "default" {
  name = "tf-test-my-subnetwork%{random_suffix}"
  network       = google_compute_network.default.id
  ip_cidr_range = "10.2.0.0/16"

}

resource "google_compute_region_backend_service" "default" {
  name = "tf-test-my-service%{random_suffix}"
  health_checks = [google_compute_health_check.default.id]
}

resource "google_compute_health_check" "default" {
  name = "tf-test-my-healthcheck%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_forwarding_rule" "default" {
  depends_on = [google_compute_subnetwork.default]
  name       = "tf-test-my-ilb%{random_suffix}"

  is_mirroring_collector = true
  ip_protocol            = "TCP"
  load_balancing_scheme  = "INTERNAL"
  backend_service        = google_compute_region_backend_service.default.id
  all_ports              = true
  network                = google_compute_network.default.id
  subnetwork             = google_compute_subnetwork.default.id
  network_tier           = "PREMIUM"
}

resource "google_compute_packet_mirroring" "foobar" {
  name = "tf-test-my-mirroring%{random_suffix}"
  description = "bar"
  network {
    url = google_compute_network.default.id
  }
  collector_ilb {
    url = google_compute_forwarding_rule.default.id
  }
  mirrored_resources {
    tags = ["foo"]
    instances {
      url = google_compute_instance.mirror.id
    }
  }
  filter {
    ip_protocols = ["tcp"]
    cidr_ranges = ["0.0.0.0/0"]
    direction = "BOTH"
  }
}
`, context)
}

func testAccCheckComputePacketMirroringDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_packet_mirroring" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputePacketMirroring still exists at %s", url)
			}
		}

		return nil
	}
}
