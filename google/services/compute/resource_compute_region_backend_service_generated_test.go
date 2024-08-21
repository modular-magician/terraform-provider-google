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

package compute_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccComputeRegionBackendService_regionBackendServiceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_regionBackendServiceBasicExample(context),
			},
			{
				ResourceName:            "google_compute_region_backend_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"iap.0.oauth2_client_secret", "network", "region"},
			},
		},
	})
}

func testAccComputeRegionBackendService_regionBackendServiceBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_backend_service" "default" {
  name                            = "tf-test-region-service%{random_suffix}"
  region                          = "us-central1"
  health_checks                   = [google_compute_health_check.default.id]
  connection_draining_timeout_sec = 10
  session_affinity                = "CLIENT_IP"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-rbs-health-check%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}
`, context)
}

func TestAccComputeRegionBackendService_regionBackendServiceExternalIapExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_regionBackendServiceExternalIapExample(context),
			},
			{
				ResourceName:            "google_compute_region_backend_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"iap.0.oauth2_client_secret", "network", "region"},
			},
		},
	})
}

func testAccComputeRegionBackendService_regionBackendServiceExternalIapExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_backend_service" "default" {
  name                  = "tf-test-tf-test-region-service-external%{random_suffix}"
  region                = "us-central1"
  protocol              = "HTTP"
  load_balancing_scheme = "EXTERNAL"
  iap {
    oauth2_client_id     = "abc"
    oauth2_client_secret = "xyz"
  }
}
`, context)
}

func TestAccComputeRegionBackendService_regionBackendServiceIlbRoundRobinExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_regionBackendServiceIlbRoundRobinExample(context),
			},
			{
				ResourceName:            "google_compute_region_backend_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"iap.0.oauth2_client_secret", "network", "region"},
			},
		},
	})
}

func testAccComputeRegionBackendService_regionBackendServiceIlbRoundRobinExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_backend_service" "default" {
  region = "us-central1"
  name = "tf-test-region-service%{random_suffix}"
  health_checks = [google_compute_health_check.health_check.id]
  protocol = "HTTP"
  load_balancing_scheme = "INTERNAL_MANAGED"
  locality_lb_policy = "ROUND_ROBIN"
}

resource "google_compute_health_check" "health_check" {
  name               = "tf-test-rbs-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeRegionBackendService_regionBackendServiceExternalWeightedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_regionBackendServiceExternalWeightedExample(context),
			},
			{
				ResourceName:            "google_compute_region_backend_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"iap.0.oauth2_client_secret", "network", "region"},
			},
		},
	})
}

func testAccComputeRegionBackendService_regionBackendServiceExternalWeightedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_backend_service" "default" {
  region                = "us-central1"
  name                  = "tf-test-region-service%{random_suffix}"
  health_checks         = [google_compute_region_health_check.health_check.id]
  protocol              = "TCP"
  load_balancing_scheme = "EXTERNAL"
  locality_lb_policy    = "WEIGHTED_MAGLEV"
}

resource "google_compute_region_health_check" "health_check" {
  name               = "tf-test-rbs-health-check%{random_suffix}"
  region             = "us-central1"

  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeRegionBackendService_regionBackendServiceIlbRingHashExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_regionBackendServiceIlbRingHashExample(context),
			},
			{
				ResourceName:            "google_compute_region_backend_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"iap.0.oauth2_client_secret", "network", "region"},
			},
		},
	})
}

func testAccComputeRegionBackendService_regionBackendServiceIlbRingHashExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_backend_service" "default" {
  region = "us-central1"
  name = "tf-test-region-service%{random_suffix}"
  health_checks = [google_compute_health_check.health_check.id]
  load_balancing_scheme = "INTERNAL_MANAGED"
  locality_lb_policy = "RING_HASH"
  session_affinity = "HTTP_COOKIE"
  protocol = "HTTP"
  circuit_breakers {
    max_connections = 10
  }
  consistent_hash {
    http_cookie {
      ttl {
        seconds = 11
        nanos = 1111
      }
      name = "mycookie"
    }
  }
  outlier_detection {
    consecutive_errors = 2
  }
}

resource "google_compute_health_check" "health_check" {
  name               = "tf-test-rbs-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeRegionBackendService_regionBackendServiceBalancingModeExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionBackendServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionBackendService_regionBackendServiceBalancingModeExample(context),
			},
			{
				ResourceName:            "google_compute_region_backend_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"iap.0.oauth2_client_secret", "network", "region"},
			},
		},
	})
}

func testAccComputeRegionBackendService_regionBackendServiceBalancingModeExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_backend_service" "default" {
  load_balancing_scheme = "INTERNAL_MANAGED"

  backend {
    group          = google_compute_region_instance_group_manager.rigm.instance_group
    balancing_mode = "UTILIZATION"
    capacity_scaler = 1.0
  }

  region      = "us-central1"
  name        = "tf-test-region-service%{random_suffix}"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.id]
}

data "google_compute_image" "debian_image" {
  family   = "debian-11"
  project  = "debian-cloud"
}

resource "google_compute_region_instance_group_manager" "rigm" {
  region   = "us-central1"
  name     = "tf-test-rbs-rigm%{random_suffix}"
  version {
    instance_template = google_compute_instance_template.instance_template.id
    name              = "primary"
  }
  base_instance_name = "internal-glb"
  target_size        = 1
}

resource "google_compute_instance_template" "instance_template" {
  name         = "template-tf-test-region-service%{random_suffix}"
  machine_type = "e2-medium"

  network_interface {
    network    = google_compute_network.default.id
    subnetwork = google_compute_subnetwork.default.id
  }

  disk {
    source_image = data.google_compute_image.debian_image.self_link
    auto_delete  = true
    boot         = true
  }

  tags = ["allow-ssh", "load-balanced-backend"]
}

resource "google_compute_region_health_check" "default" {
  region = "us-central1"
  name   = "tf-test-rbs-health-check%{random_suffix}"
  http_health_check {
    port_specification = "USE_SERVING_PORT"
  }
}

resource "google_compute_network" "default" {
  name                    = "tf-test-rbs-net%{random_suffix}"
  auto_create_subnetworks = false
  routing_mode            = "REGIONAL"
}

resource "google_compute_subnetwork" "default" {
  name          = "tf-test-rbs-net%{random_suffix}-default"
  ip_cidr_range = "10.1.2.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.id
}
`, context)
}

func testAccCheckComputeRegionBackendServiceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_region_backend_service" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/backendServices/{{name}}")
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
				return fmt.Errorf("ComputeRegionBackendService still exists at %s", url)
			}
		}

		return nil
	}
}
