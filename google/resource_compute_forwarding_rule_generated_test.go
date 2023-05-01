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

func TestAccComputeForwardingRule_forwardingRuleGlobalInternallbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleGlobalInternallbExample(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "region"},
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleGlobalInternallbExample(context map[string]interface{}) string {
	return Nprintf(`
// Forwarding rule for Internal Load Balancing
resource "google_compute_forwarding_rule" "default" {
  name                  = "tf-test-website-forwarding-rule%{random_suffix}"
  region                = "us-central1"
  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.backend.id
  all_ports             = true
  allow_global_access   = true
  network               = google_compute_network.default.name
  subnetwork            = google_compute_subnetwork.default.name
}
resource "google_compute_region_backend_service" "backend" {
  name                  = "tf-test-website-backend%{random_suffix}"
  region                = "us-central1"
  health_checks         = [google_compute_health_check.hc.id]
}
resource "google_compute_health_check" "hc" {
  name               = "check-tf-test-website-backend%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}
resource "google_compute_network" "default" {
  name = "tf-test-website-net%{random_suffix}"
  auto_create_subnetworks = false
}
resource "google_compute_subnetwork" "default" {
  name          = "tf-test-website-net%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.default.id
}
`, context)
}

func TestAccComputeForwardingRule_forwardingRuleBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleBasicExample(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "region", "port_range", "target"},
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_forwarding_rule" "default" {
  name       = "tf-test-website-forwarding-rule%{random_suffix}"
  target     = google_compute_target_pool.default.id
  port_range = "80"
}

resource "google_compute_target_pool" "default" {
  name = "tf-test-website-target-pool%{random_suffix}"
}
`, context)
}

func TestAccComputeForwardingRule_forwardingRuleInternallbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleInternallbExample(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "region", "port_range", "target"},
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleInternallbExample(context map[string]interface{}) string {
	return Nprintf(`
// Forwarding rule for Internal Load Balancing
resource "google_compute_forwarding_rule" "default" {
  name   = "tf-test-website-forwarding-rule%{random_suffix}"
  region = "us-central1"

  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.backend.id
  all_ports             = true
  network               = google_compute_network.default.name
  subnetwork            = google_compute_subnetwork.default.name
}

resource "google_compute_region_backend_service" "backend" {
  name          = "tf-test-website-backend%{random_suffix}"
  region        = "us-central1"
  health_checks = [google_compute_health_check.hc.id]
}

resource "google_compute_health_check" "hc" {
  name               = "check-tf-test-website-backend%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_network" "default" {
  name                    = "tf-test-website-net%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "tf-test-website-net%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.default.id
}
`, context)
}

func TestAccComputeForwardingRule_forwardingRuleRegionalSteeringExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleRegionalSteeringExample(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.steering",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "region"},
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleRegionalSteeringExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_forwarding_rule" "steering" {
  name = "tf-test-steering-rule%{random_suffix}"
  region = "us-central1"
  ip_address = google_compute_address.basic.self_link
  backend_service = google_compute_region_backend_service.external.self_link
  load_balancing_scheme = "EXTERNAL"
  source_ip_ranges = ["34.121.88.0/24", "35.187.239.137"]
  depends_on = [google_compute_forwarding_rule.external]
}

resource "google_compute_address" "basic" {
  name = "tf-test-website-ip%{random_suffix}"
  region = "us-central1"
}

resource "google_compute_region_backend_service" "external" {
  name = "tf-test-service-backend%{random_suffix}"
  region = "us-central1"
  load_balancing_scheme = "EXTERNAL"
}

resource "google_compute_forwarding_rule" "external" {
  name = "tf-test-external-forwarding-rule%{random_suffix}"
  region = "us-central1"
  ip_address = google_compute_address.basic.self_link
  backend_service = google_compute_region_backend_service.external.self_link
  load_balancing_scheme = "EXTERNAL"
}
`, context)
}

func testAccCheckComputeForwardingRuleDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_forwarding_rule" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeForwardingRule still exists at %s", url)
			}
		}

		return nil
	}
}
