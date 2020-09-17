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

	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccComputeHealthCheck_healthCheckTcpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckTcpExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.tcp-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckTcpExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_health_check" "tcp-health-check" {
  name = "tf-test-tcp-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  tcp_health_check {
    port = "80"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckTcpFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckTcpFullExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.tcp-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckTcpFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_health_check" "tcp-health-check" {
  name        = "tf-test-tcp-health-check%{random_suffix}"
  description = "Health check via tcp"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  tcp_health_check {
    port_name          = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    request            = "ARE YOU HEALTHY?"
    proxy_header       = "NONE"
    response           = "I AM HEALTHY"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckSslExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckSslExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.ssl-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckSslExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_health_check" "ssl-health-check" {
  name = "tf-test-ssl-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  ssl_health_check {
    port = "443"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckSslFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckSslFullExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.ssl-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckSslFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_health_check" "ssl-health-check" {
  name        = "tf-test-ssl-health-check%{random_suffix}"
  description = "Health check via ssl"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  ssl_health_check {
    port_name          = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    request            = "ARE YOU HEALTHY?"
    proxy_header       = "NONE"
    response           = "I AM HEALTHY"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckHttpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckHttpExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.http-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckHttpExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_health_check" "http-health-check" {
  name = "tf-test-http-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckHttpFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckHttpFullExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.http-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckHttpFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_health_check" "http-health-check" {
  name        = "tf-test-http-health-check%{random_suffix}"
  description = "Health check via http"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  http_health_check {
    port_name          = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    host               = "1.2.3.4"
    request_path       = "/mypath"
    proxy_header       = "NONE"
    response           = "I AM HEALTHY"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckHttpsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckHttpsExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.https-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckHttpsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_health_check" "https-health-check" {
  name = "tf-test-https-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  https_health_check {
    port = "443"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckHttpsFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckHttpsFullExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.https-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckHttpsFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_health_check" "https-health-check" {
  name        = "tf-test-https-health-check%{random_suffix}"
  description = "Health check via https"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  https_health_check {
    port_name          = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    host               = "1.2.3.4"
    request_path       = "/mypath"
    proxy_header       = "NONE"
    response           = "I AM HEALTHY"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckHttp2Example(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckHttp2Example(context),
			},
			{
				ResourceName:      "google_compute_health_check.http2-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckHttp2Example(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_health_check" "http2-health-check" {
  name = "tf-test-http2-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  http2_health_check {
    port = "443"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckHttp2FullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckHttp2FullExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.http2-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckHttp2FullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_health_check" "http2-health-check" {
  name        = "tf-test-http2-health-check%{random_suffix}"
  description = "Health check via http2"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  http2_health_check {
    port_name          = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    host               = "1.2.3.4"
    request_path       = "/mypath"
    proxy_header       = "NONE"
    response           = "I AM HEALTHY"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckGrpcExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckGrpcExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.grpc-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckGrpcExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_health_check" "grpc-health-check" {
  name = "tf-test-grpc-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  grpc_health_check {
    port = "443"
  }
}
`, context)
}

func TestAccComputeHealthCheck_healthCheckGrpcFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckComputeHealthCheckDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeHealthCheck_healthCheckGrpcFullExample(context),
			},
			{
				ResourceName:      "google_compute_health_check.grpc-health-check",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeHealthCheck_healthCheckGrpcFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_health_check" "grpc-health-check" {
  name = "tf-test-grpc-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  grpc_health_check {
    port_name          = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    grpc_service_name  = "testservice"
  }
}
`, context)
}

func testAccCheckComputeHealthCheckDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_health_check" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/healthChecks/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("ComputeHealthCheck still exists at %s", url)
			}
		}

		return nil
	}
}
