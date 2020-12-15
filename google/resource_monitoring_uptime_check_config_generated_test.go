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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccMonitoringUptimeCheckConfig_uptimeCheckConfigHttpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckMonitoringUptimeCheckConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringUptimeCheckConfig_uptimeCheckConfigHttpExample(context),
			},
			{
				ResourceName:      "google_monitoring_uptime_check_config.http",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMonitoringUptimeCheckConfig_uptimeCheckConfigHttpExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_monitoring_uptime_check_config" "http" {
  display_name = "tf-test-http-uptime-check%{random_suffix}"
  timeout      = "60s"

  http_check {
    path = "/some-path"
    port = "8010"
    request_method = "POST"
    content_type = "URL_ENCODED"
    body = "Zm9vJTI1M0RiYXI="
  }

  monitored_resource {
    type = "uptime_url"
    labels = {
      project_id = "%{project_id}"
      host       = "192.168.1.1"
    }
  }

  content_matchers {
    content = "example"
  }
}
`, context)
}

func TestAccMonitoringUptimeCheckConfig_uptimeCheckConfigHttpsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_id":    getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckMonitoringUptimeCheckConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringUptimeCheckConfig_uptimeCheckConfigHttpsExample(context),
			},
			{
				ResourceName:      "google_monitoring_uptime_check_config.https",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMonitoringUptimeCheckConfig_uptimeCheckConfigHttpsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_monitoring_uptime_check_config" "https" {
  display_name = "tf-test-https-uptime-check%{random_suffix}"
  timeout = "60s"

  http_check {
    path = "/some-path"
    port = "443"
    use_ssl = true
    validate_ssl = true
  }

  monitored_resource {
    type = "uptime_url"
    labels = {
      project_id = "%{project_id}"
      host = "192.168.1.1"
    }
  }

  content_matchers {
    content = "example"
  }
}
`, context)
}

func TestAccMonitoringUptimeCheckConfig_uptimeCheckTcpExample(t *testing.T) {
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
		CheckDestroy: testAccCheckMonitoringUptimeCheckConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringUptimeCheckConfig_uptimeCheckTcpExample(context),
			},
			{
				ResourceName:      "google_monitoring_uptime_check_config.tcp_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMonitoringUptimeCheckConfig_uptimeCheckTcpExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_monitoring_uptime_check_config" "tcp_group" {
  display_name = "tf-test-tcp-uptime-check%{random_suffix}"
  timeout      = "60s"

  tcp_check {
    port = 888
  }

  resource_group {
    resource_type = "INSTANCE"
    group_id      = google_monitoring_group.check.name
  }
}

resource "google_monitoring_group" "check" {
  display_name = "tf-test-uptime-check-group%{random_suffix}"
  filter       = "resource.metadata.name=has_substring(\"foo\")"
}
`, context)
}

func testAccCheckMonitoringUptimeCheckConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_monitoring_uptime_check_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{MonitoringBasePath}}v3/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil, isMonitoringConcurrentEditError)
			if err == nil {
				return fmt.Errorf("MonitoringUptimeCheckConfig still exists at %s", url)
			}
		}

		return nil
	}
}
