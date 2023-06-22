// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

func TestAccComputeGlobalForwardingRule_globalForwardingRuleHttpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalForwardingRule_globalForwardingRuleHttpExample(context),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "port_range", "target"},
			},
		},
	})
}

func testAccComputeGlobalForwardingRule_globalForwardingRuleHttpExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_global_forwarding_rule" "default" {
  name       = "tf-test-global-rule%{random_suffix}"
  target     = google_compute_target_http_proxy.default.id
  port_range = "80"
}

resource "google_compute_target_http_proxy" "default" {
  name        = "tf-test-target-proxy%{random_suffix}"
  description = "a description"
  url_map     = google_compute_url_map.default.id
}

resource "google_compute_url_map" "default" {
  name            = "url-map-tf-test-target-proxy%{random_suffix}"
  description     = "a description"
  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.id
    }
  }
}

resource "google_compute_backend_service" "default" {
  name        = "backend%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "check-backend%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, context)
}

func TestAccComputeGlobalForwardingRule_globalForwardingRuleExternalManagedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalForwardingRule_globalForwardingRuleExternalManagedExample(context),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "port_range", "target"},
			},
		},
	})
}

func testAccComputeGlobalForwardingRule_globalForwardingRuleExternalManagedExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_global_forwarding_rule" "default" {
  name                  = "tf-test-global-rule%{random_suffix}"
  target                = google_compute_target_http_proxy.default.id
  port_range            = "80"
  load_balancing_scheme = "EXTERNAL_MANAGED"
}

resource "google_compute_target_http_proxy" "default" {
  name        = "tf-test-target-proxy%{random_suffix}"
  description = "a description"
  url_map     = google_compute_url_map.default.id
}

resource "google_compute_url_map" "default" {
  name            = "url-map-tf-test-target-proxy%{random_suffix}"
  description     = "a description"
  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.id
    }
  }
}

resource "google_compute_backend_service" "default" {
  name                  = "backend%{random_suffix}"
  port_name             = "http"
  protocol              = "HTTP"
  timeout_sec           = 10
  load_balancing_scheme = "EXTERNAL_MANAGED"
}
`, context)
}

func TestAccComputeGlobalForwardingRule_globalForwardingRuleHybridExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeGlobalForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeGlobalForwardingRule_globalForwardingRuleHybridExample(context),
			},
			{
				ResourceName:            "google_compute_global_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "port_range", "target"},
			},
		},
	})
}

func testAccComputeGlobalForwardingRule_globalForwardingRuleHybridExample(context map[string]interface{}) string {
	return Nprintf(`
// Roughly mirrors https://cloud.google.com/load-balancing/docs/https/setting-up-ext-https-hybrid
variable "subnetwork_cidr" {
  default = "10.0.0.0/24"
}

resource "google_compute_network" "default" {
  name                    = "tf-test-my-network%{random_suffix}"
}

resource "google_compute_network" "internal" {
  name                    = "tf-test-my-internal-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "internal"{
  name                    = "tf-test-my-subnetwork%{random_suffix}"
  network                 = google_compute_network.internal.id
  ip_cidr_range           = var.subnetwork_cidr
  region                  = "us-central1"
  private_ip_google_access= true
}

// Zonal NEG with GCE_VM_IP_PORT
resource "google_compute_network_endpoint_group" "default" {
  name                  = "tf-test-default-neg%{random_suffix}"
  network               = google_compute_network.default.id
  default_port          = "90"
  zone                  = "us-central1-a"
  network_endpoint_type = "GCE_VM_IP_PORT"
}

// Zonal NEG with GCE_VM_IP
resource "google_compute_network_endpoint_group" "internal" {
  name                  = "tf-test-internal-neg%{random_suffix}"
  network               = google_compute_network.internal.id
  subnetwork            = google_compute_subnetwork.internal.id
  zone                  = "us-central1-a"
  network_endpoint_type = "GCE_VM_IP"
}

// Hybrid connectivity NEG
resource "google_compute_network_endpoint_group" "hybrid" {
  name                  = "tf-test-hybrid-neg%{random_suffix}"
  network               = google_compute_network.default.id
  default_port          = "90"
  zone                  = "us-central1-a"
  network_endpoint_type = "NON_GCP_PRIVATE_IP_PORT"
}

resource "google_compute_network_endpoint" "hybrid-endpoint" {
  network_endpoint_group = google_compute_network_endpoint_group.hybrid.name
  port       = google_compute_network_endpoint_group.hybrid.default_port
  ip_address = "127.0.0.1"
}

// Backend service for Zonal NEG
resource "google_compute_backend_service" "default" {
  name                  = "tf-test-backend-default%{random_suffix}"
  port_name             = "http"
  protocol              = "HTTP"
  timeout_sec           = 10
  backend {
    group = google_compute_network_endpoint_group.default.id
    balancing_mode               = "RATE"
    max_rate_per_endpoint        = 10
  }
  health_checks = [google_compute_health_check.default.id]
}

// Backgend service for Hybrid NEG
resource "google_compute_backend_service" "hybrid" {
  name                  = "tf-test-backend-hybrid%{random_suffix}"
  port_name             = "http"
  protocol              = "HTTP"
  timeout_sec           = 10
  backend {
    group                        = google_compute_network_endpoint_group.hybrid.id
    balancing_mode               = "RATE"
    max_rate_per_endpoint = 10
  }
  health_checks = [google_compute_health_check.default.id]
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  timeout_sec        = 1
  check_interval_sec = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_url_map" "default" {
  name            = "url-map-tf-test-target-proxy%{random_suffix}"
  description     = "a description"
  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_backend_service.default.id
    }

    path_rule {
      paths   = ["/hybrid"]
      service = google_compute_backend_service.hybrid.id
    }
  }
}

resource "google_compute_target_http_proxy" "default" {
  name        = "tf-test-target-proxy%{random_suffix}"
  description = "a description"
  url_map     = google_compute_url_map.default.id
}

resource "google_compute_global_forwarding_rule" "default" {
  name       = "tf-test-global-rule%{random_suffix}"
  target     = google_compute_target_http_proxy.default.id
  port_range = "80"
}
`, context)
}

func testAccCheckComputeGlobalForwardingRuleDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_global_forwarding_rule" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/forwardingRules/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ComputeGlobalForwardingRule still exists at %s", url)
			}
		}

		return nil
	}
}
