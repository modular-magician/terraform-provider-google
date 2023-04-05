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

func TestAccNetworkServicesEdgeCacheOrigin_networkServicesEdgeCacheOriginBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesEdgeCacheOriginDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesEdgeCacheOrigin_networkServicesEdgeCacheOriginBasicExample(context),
			},
			{
				ResourceName:            "google_network_services_edge_cache_origin.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "timeout"},
			},
		},
	})
}

func testAccNetworkServicesEdgeCacheOrigin_networkServicesEdgeCacheOriginBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_network_services_edge_cache_origin" "default" {
  name                 = "tf-test-my-origin%{random_suffix}"
  origin_address       = "gs://media-edge-default"
  description          = "The default bucket for media edge test"
}
`, context)
}

func TestAccNetworkServicesEdgeCacheOrigin_networkServicesEdgeCacheOriginAdvancedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesEdgeCacheOriginDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesEdgeCacheOrigin_networkServicesEdgeCacheOriginAdvancedExample(context),
			},
			{
				ResourceName:            "google_network_services_edge_cache_origin.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "timeout"},
			},
		},
	})
}

func testAccNetworkServicesEdgeCacheOrigin_networkServicesEdgeCacheOriginAdvancedExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_network_services_edge_cache_origin" "fallback" {
  name                 = "tf-test-my-fallback%{random_suffix}"
  origin_address       = "fallback.example.com"
  description          = "The default bucket for media edge test"
  max_attempts         = 3
  protocol = "HTTP"
  port = 80

  retry_conditions = [
    "CONNECT_FAILURE",
    "NOT_FOUND",
    "HTTP_5XX",
    "FORBIDDEN",
  ]
  timeout {
    connect_timeout = "10s"
    max_attempts_timeout = "20s"
    response_timeout = "60s"
    read_timeout = "5s"
  }
  origin_override_action {
    url_rewrite {
      host_rewrite = "example.com"
    }
    header_action {
      request_headers_to_add {
        header_name = "x-header"
	header_value = "value"
	replace = true
      }
    }
  }
  origin_redirect {
    redirect_conditions = [
      "MOVED_PERMANENTLY",
      "FOUND",
      "SEE_OTHER",
      "TEMPORARY_REDIRECT",
      "PERMANENT_REDIRECT",
    ]
  }
}

resource "google_network_services_edge_cache_origin" "default" {
  name                 = "tf-test-my-origin%{random_suffix}"
  origin_address       = "gs://media-edge-default"
  failover_origin      = google_network_services_edge_cache_origin.fallback.id
  description          = "The default bucket for media edge test"
  max_attempts         = 2
  labels = {
    a = "b"
  }

  timeout {
    connect_timeout = "10s"
  }
}
`, context)
}

func TestAccNetworkServicesEdgeCacheOrigin_networkServicesEdgeCacheOriginV4authExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkServicesEdgeCacheOriginDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesEdgeCacheOrigin_networkServicesEdgeCacheOriginV4authExample(context),
			},
			{
				ResourceName:            "google_network_services_edge_cache_origin.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "timeout"},
			},
		},
	})
}

func testAccNetworkServicesEdgeCacheOrigin_networkServicesEdgeCacheOriginV4authExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "tf-test-secret-name%{random_suffix}"

  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret-version-basic" {
  secret = google_secret_manager_secret.secret-basic.id

  secret_data = "secret-data"
}

resource "google_network_services_edge_cache_origin" "default" {
  name           = "tf-test-my-origin%{random_suffix}"
  origin_address = "gs://media-edge-default"
  description    = "The default bucket for V4 authentication"
  aws_v4_authentication {
    access_key_id             = "ACCESSKEYID"
    secret_access_key_version = google_secret_manager_secret_version.secret-version-basic.id
    origin_region             = "auto"
  }
}
`, context)
}

func testAccCheckNetworkServicesEdgeCacheOriginDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_services_edge_cache_origin" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/edgeCacheOrigins/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("NetworkServicesEdgeCacheOrigin still exists at %s", url)
			}
		}

		return nil
	}
}
