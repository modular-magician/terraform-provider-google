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

func TestAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupFunctionsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"zip_path":      acctest.CreateZIPArchiveForCloudFunctionSource(t, "./test-fixtures/http_trigger.js"),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionNetworkEndpointGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupFunctionsExample(context),
			},
			{
				ResourceName:            "google_compute_region_network_endpoint_group.function_neg",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "subnetwork"},
			},
		},
	})
}

func testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupFunctionsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
// Cloud Functions Example
resource "google_compute_region_network_endpoint_group" "function_neg" {
  name                  = "tf-test-function-neg%{random_suffix}"
  network_endpoint_type = "SERVERLESS"
  region                = "us-central1"
  cloud_function {
    function = google_cloudfunctions_function.function_neg.name
  }
}

resource "google_cloudfunctions_function" "function_neg" {
  name        = "tf-test-function-neg%{random_suffix}"
  description = "My function"
  runtime     = "nodejs10"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = google_storage_bucket_object.archive.name
  trigger_http          = true
  timeout               = 60
  entry_point           = "helloGET"
}

resource "google_storage_bucket" "bucket" {
  name     = "tf-test-cloudfunctions-function-example-bucket%{random_suffix}"
  location = "US"
}

resource "google_storage_bucket_object" "archive" { 
  name   = "index.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"
}
`, context)
}

func TestAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupCloudrunExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionNetworkEndpointGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupCloudrunExample(context),
			},
			{
				ResourceName:            "google_compute_region_network_endpoint_group.cloudrun_neg",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "subnetwork"},
			},
		},
	})
}

func testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupCloudrunExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
// Cloud Run Example
resource "google_compute_region_network_endpoint_group" "cloudrun_neg" {
  name                  = "tf-test-cloudrun-neg%{random_suffix}"
  network_endpoint_type = "SERVERLESS"
  region                = "us-central1"
  cloud_run {
    service = google_cloud_run_service.cloudrun_neg.name
  }
}

resource "google_cloud_run_service" "cloudrun_neg" {
  name     = "tf-test-cloudrun-neg%{random_suffix}"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
`, context)
}

func TestAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupAppengineExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"archive": {},
		},
		CheckDestroy: testAccCheckComputeRegionNetworkEndpointGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupAppengineExample(context),
			},
			{
				ResourceName:            "google_compute_region_network_endpoint_group.appengine_neg",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "subnetwork"},
			},
		},
	})
}

func testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupAppengineExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
// App Engine Example
resource "google_compute_region_network_endpoint_group" "appengine_neg" {
  name                  = "tf-test-appengine-neg%{random_suffix}"
  network_endpoint_type = "SERVERLESS"
  region                = "us-central1"
  app_engine {
    service = google_app_engine_flexible_app_version.appengine_neg.service
    version = google_app_engine_flexible_app_version.appengine_neg.version_id
  }
}

resource "google_app_engine_flexible_app_version" "appengine_neg" {
  version_id = "v1"
  service    = "appengine-network-endpoint-group"
  runtime    = "nodejs"

  entrypoint {
    shell = "node ./app.js"
  }

  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.appengine_neg.name}/${google_storage_bucket_object.appengine_neg.name}"
    }
  }

  liveness_check {
    path = "/"
  }

  readiness_check {
    path = "/"
  }

  env_variables = {
    port = "8080"
  }

  handlers {
    url_regex        = ".*\\/my-path\\/*"
    security_level   = "SECURE_ALWAYS"
    login            = "LOGIN_REQUIRED"
    auth_fail_action = "AUTH_FAIL_ACTION_REDIRECT"

    static_files {
      path = "my-other-path"
      upload_path_regex = ".*\\/my-path\\/*"
    }
  }

  automatic_scaling {
    cool_down_period = "120s"
    cpu_utilization {
      target_utilization = 0.5
    }
  }

  delete_service_on_destroy = true
}

resource "google_storage_bucket" "appengine_neg" {
  name     = "tf-test-appengine-neg%{random_suffix}"
  location = "US"
}

data "archive_file" "app" {
  type        = "zip"
  source_dir = "./test-fixtures/hello-world-node-standard"
  output_path = "./test-fixtures/hello-world-node-standard.zip"
}

resource "google_storage_bucket_object" "appengine_neg" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.appengine_neg.name
  source = data.archive_file.app.output_path
}
`, context)
}

func TestAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupAppengineEmptyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionNetworkEndpointGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupAppengineEmptyExample(context),
			},
			{
				ResourceName:            "google_compute_region_network_endpoint_group.appengine_neg",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "subnetwork"},
			},
		},
	})
}

func testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupAppengineEmptyExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
// App Engine Example
resource "google_compute_region_network_endpoint_group" "appengine_neg" {
  name                  = "tf-test-appengine-neg%{random_suffix}"
  network_endpoint_type = "SERVERLESS"
  region                = "us-central1"
  app_engine {
  }
}
`, context)
}

func TestAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupPscExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionNetworkEndpointGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupPscExample(context),
			},
			{
				ResourceName:            "google_compute_region_network_endpoint_group.psc_neg",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "subnetwork"},
			},
		},
	})
}

func testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupPscExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_network_endpoint_group" "psc_neg" {
  name                  = "tf-test-psc-neg%{random_suffix}"
  region                = "asia-northeast3"

  network_endpoint_type = "PRIVATE_SERVICE_CONNECT"
  psc_target_service    = "asia-northeast3-cloudkms.googleapis.com"
}
`, context)
}

func TestAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupPscServiceAttachmentExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionNetworkEndpointGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupPscServiceAttachmentExample(context),
			},
			{
				ResourceName:            "google_compute_region_network_endpoint_group.psc_neg_service_attachment",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "subnetwork"},
			},
		},
	})
}

func testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupPscServiceAttachmentExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_network" "default" {
  name = "tf-test-psc-network%{random_suffix}"
}

resource "google_compute_subnetwork" "default" {
  name          = "tf-test-psc-subnetwork%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "europe-west4"
  network       = google_compute_network.default.id
}

resource "google_compute_subnetwork" "psc_subnetwork" {
  name          = "tf-test-psc-subnetwork-nat%{random_suffix}"
  ip_cidr_range = "10.1.0.0/16"
  region        = "europe-west4"
  purpose       = "PRIVATE_SERVICE_CONNECT"
  network       = google_compute_network.default.id
}

resource "google_compute_health_check" "default" {
  name = "tf-test-psc-healthcheck%{random_suffix}"

  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}
resource "google_compute_region_backend_service" "default" {
  name   = "tf-test-psc-backend%{random_suffix}"
  region = "europe-west4"

  health_checks = [google_compute_health_check.default.id]
}

resource "google_compute_forwarding_rule" "default" {
  name   = "tf-test-psc-forwarding-rule%{random_suffix}"
  region = "europe-west4"

  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.default.id
  all_ports             = true
  network               = google_compute_network.default.name
  subnetwork            = google_compute_subnetwork.default.name
}

resource "google_compute_service_attachment" "default" {
  name        = "tf-test-psc-service-attachment%{random_suffix}"
  region      = "europe-west4"
  description = "A service attachment configured with Terraform"

  enable_proxy_protocol = false
  connection_preference = "ACCEPT_AUTOMATIC"
  nat_subnets           = [google_compute_subnetwork.psc_subnetwork.self_link]
  target_service        = google_compute_forwarding_rule.default.self_link
}

resource "google_compute_region_network_endpoint_group" "psc_neg_service_attachment" {
  name                  = "tf-test-psc-neg%{random_suffix}"
  region                = "europe-west4"

  network_endpoint_type = "PRIVATE_SERVICE_CONNECT"
  psc_target_service    = google_compute_service_attachment.default.self_link

  network               = google_compute_network.default.self_link
  subnetwork            = google_compute_subnetwork.default.self_link
}
`, context)
}

func TestAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupInternetIpPortExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionNetworkEndpointGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupInternetIpPortExample(context),
			},
			{
				ResourceName:            "google_compute_region_network_endpoint_group.region_network_endpoint_group_internet_ip_port",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "subnetwork"},
			},
		},
	})
}

func testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupInternetIpPortExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_network_endpoint_group" "region_network_endpoint_group_internet_ip_port" {
  name                  = "tf-test-ip-port-neg%{random_suffix}"
  region                = "us-central1"
  network               = google_compute_network.default.id

  network_endpoint_type = "INTERNET_IP_PORT"
}

resource "google_compute_network" "default" {
  name                    = "network%{random_suffix}"
}
`, context)
}

func TestAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupInternetFqdnPortExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeRegionNetworkEndpointGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupInternetFqdnPortExample(context),
			},
			{
				ResourceName:            "google_compute_region_network_endpoint_group.region_network_endpoint_group_internet_fqdn_port",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"network", "region", "subnetwork"},
			},
		},
	})
}

func testAccComputeRegionNetworkEndpointGroup_regionNetworkEndpointGroupInternetFqdnPortExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_region_network_endpoint_group" "region_network_endpoint_group_internet_fqdn_port" {
  name                  = "tf-test-ip-port-neg%{random_suffix}"
  region                = "us-central1"
  network               = google_compute_network.default.id

  network_endpoint_type = "INTERNET_FQDN_PORT"
}

resource "google_compute_network" "default" {
  name                    = "network%{random_suffix}"
}
`, context)
}

func testAccCheckComputeRegionNetworkEndpointGroupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_region_network_endpoint_group" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}")
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
				return fmt.Errorf("ComputeRegionNetworkEndpointGroup still exists at %s", url)
			}
		}

		return nil
	}
}
