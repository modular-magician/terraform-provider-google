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

func TestAccComputeUrlMap_urlMapBasicExample(t *testing.T) {
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
		CheckDestroy: testAccCheckComputeUrlMapDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeUrlMap_urlMapBasicExample(context),
			},
			{
				ResourceName:            "google_compute_url_map.urlmap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"default_service"},
			},
		},
	})
}

func testAccComputeUrlMap_urlMapBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_url_map" "urlmap" {
  name        = "urlmap%{random_suffix}"
  description = "a description"

  default_service = google_compute_backend_service.home.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "mysite"
  }

  host_rule {
    hosts        = ["myothersite.com"]
    path_matcher = "otherpaths"
  }

  path_matcher {
    name            = "mysite"
    default_service = google_compute_backend_service.home.id

    path_rule {
      paths   = ["/home"]
      service = google_compute_backend_service.home.id
    }

    path_rule {
      paths   = ["/login"]
      service = google_compute_backend_service.login.id
    }

    path_rule {
      paths   = ["/static"]
      service = google_compute_backend_bucket.static.id
    }
  }

  path_matcher {
    name            = "otherpaths"
    default_service = google_compute_backend_service.home.id
  }

  test {
    service = google_compute_backend_service.home.id
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_backend_service" "login" {
  name        = "login%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_backend_service" "home" {
  name        = "home%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}

resource "google_compute_backend_bucket" "static" {
  name        = "tf-test-static-asset-backend-bucket%{random_suffix}"
  bucket_name = google_storage_bucket.static.name
  enable_cdn  = true
}

resource "google_storage_bucket" "static" {
  name     = "tf-test-static-asset-bucket%{random_suffix}"
  location = "US"
}
`, context)
}

func TestAccComputeUrlMap_urlMapTrafficDirectorRouteExample(t *testing.T) {
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
		CheckDestroy: testAccCheckComputeUrlMapDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeUrlMap_urlMapTrafficDirectorRouteExample(context),
			},
			{
				ResourceName:            "google_compute_url_map.urlmap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"default_service"},
			},
		},
	})
}

func testAccComputeUrlMap_urlMapTrafficDirectorRouteExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_url_map" "urlmap" {
  name        = "urlmap%{random_suffix}"
  description = "a description"
  default_service = google_compute_backend_service.home.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = google_compute_backend_service.home.id

    route_rules {
      priority = 1
      header_action {
        request_headers_to_remove = ["RemoveMe2"]
        request_headers_to_add {
          header_name = "AddSomethingElse"
          header_value = "MyOtherValue"
          replace = true
        }
        response_headers_to_remove = ["RemoveMe3"]
        response_headers_to_add {
          header_name = "AddMe"
          header_value = "MyValue"
          replace = false
        }
      }
      match_rules {
        full_path_match = "a full path"
        header_matches {
          header_name = "someheader"
          exact_match = "match this exactly"
          invert_match = true
        }
        ignore_case = true
        metadata_filters {
          filter_match_criteria = "MATCH_ANY"
          filter_labels {
            name = "PLANET"
            value = "MARS"
          }
        }
        query_parameter_matches {
          name = "a query parameter"
          present_match = true
        }
      }
      url_redirect {
        host_redirect = "A host"
        https_redirect = false
        path_redirect = "some/path"
        redirect_response_code = "TEMPORARY_REDIRECT"
        strip_query = true
      }
    }
  }

  test {
    service = google_compute_backend_service.home.id
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_backend_service" "home" {
  name        = "home%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_health_check.default.id]
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeUrlMap_urlMapTrafficDirectorRoutePartialExample(t *testing.T) {
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
		CheckDestroy: testAccCheckComputeUrlMapDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeUrlMap_urlMapTrafficDirectorRoutePartialExample(context),
			},
			{
				ResourceName:            "google_compute_url_map.urlmap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"default_service"},
			},
		},
	})
}

func testAccComputeUrlMap_urlMapTrafficDirectorRoutePartialExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_url_map" "urlmap" {
  name        = "urlmap%{random_suffix}"
  description = "a description"
  default_service = google_compute_backend_service.home.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = google_compute_backend_service.home.id

    route_rules {
      priority = 1
      match_rules {
        prefix_match = "/someprefix"
        header_matches {
          header_name = "someheader"
          exact_match = "match this exactly"
          invert_match = true
        }
      }
      url_redirect {
        path_redirect = "some/path"
        redirect_response_code = "TEMPORARY_REDIRECT"
      }
    }
  }

  test {
    service = google_compute_backend_service.home.id
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_backend_service" "home" {
  name        = "home%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_health_check.default.id]
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeUrlMap_urlMapTrafficDirectorPathExample(t *testing.T) {
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
		CheckDestroy: testAccCheckComputeUrlMapDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeUrlMap_urlMapTrafficDirectorPathExample(context),
			},
			{
				ResourceName:            "google_compute_url_map.urlmap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"default_service"},
			},
		},
	})
}

func testAccComputeUrlMap_urlMapTrafficDirectorPathExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_url_map" "urlmap" {
  name        = "urlmap%{random_suffix}"
  description = "a description"
  default_service = google_compute_backend_service.home.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = google_compute_backend_service.home.id

    path_rule {
      paths   = ["/home"]
      route_action {
        cors_policy {
          allow_credentials = true
          allow_headers = ["Allowed content"]
          allow_methods = ["GET"]
          allow_origin_regexes = ["abc.*"]
          allow_origins = ["Allowed origin"]
          expose_headers = ["Exposed header"]
          max_age = 30
          disabled = false
        }
        fault_injection_policy {
          abort {
            http_status = 234
            percentage = 5.6
          }
          delay {
            fixed_delay {
              seconds = 0
              nanos = 50000
            }
            percentage = 7.8
          }
        }
        request_mirror_policy {
          backend_service = google_compute_backend_service.home.id
        }
        retry_policy {
          num_retries = 4
          per_try_timeout {
            seconds = 30
          }
          retry_conditions = ["5xx", "deadline-exceeded"]
        }
        timeout {
          seconds = 20
          nanos = 750000000
        }
        url_rewrite {
          host_rewrite = "A replacement header"
          path_prefix_rewrite = "A replacement path"
        }
        weighted_backend_services {
          backend_service = google_compute_backend_service.home.id
          weight = 400
          header_action {
            request_headers_to_remove = ["RemoveMe"]
            request_headers_to_add {
              header_name = "AddMe"
              header_value = "MyValue"
              replace = true
            }
            response_headers_to_remove = ["RemoveMe"]
            response_headers_to_add {
              header_name = "AddMe"
              header_value = "MyValue"
              replace = false
            }
          }
        }
      }
    }
  }

  test {
    service = google_compute_backend_service.home.id
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_backend_service" "home" {
  name        = "home%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_health_check.default.id]
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeUrlMap_urlMapTrafficDirectorPathPartialExample(t *testing.T) {
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
		CheckDestroy: testAccCheckComputeUrlMapDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeUrlMap_urlMapTrafficDirectorPathPartialExample(context),
			},
			{
				ResourceName:            "google_compute_url_map.urlmap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"default_service"},
			},
		},
	})
}

func testAccComputeUrlMap_urlMapTrafficDirectorPathPartialExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_url_map" "urlmap" {
  name        = "urlmap%{random_suffix}"
  description = "a description"
  default_service = google_compute_backend_service.home.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = google_compute_backend_service.home.id

    path_rule {
      paths   = ["/home"]
      route_action {
        cors_policy {
          allow_credentials = true
          allow_headers = ["Allowed content"]
          allow_methods = ["GET"]
          allow_origin_regexes = ["abc.*"]
          allow_origins = ["Allowed origin"]
          expose_headers = ["Exposed header"]
          max_age = 30
          disabled = false
        }
        weighted_backend_services {
          backend_service = google_compute_backend_service.home.id
          weight = 400
          header_action {
            request_headers_to_remove = ["RemoveMe"]
            request_headers_to_add {
              header_name = "AddMe"
              header_value = "MyValue"
              replace = true
            }
            response_headers_to_remove = ["RemoveMe"]
            response_headers_to_add {
              header_name = "AddMe"
              header_value = "MyValue"
              replace = false
            }
          }
        }
      }
    }
  }

  test {
    service = google_compute_backend_service.home.id
    host    = "hi.com"
    path    = "/home"
  }
}

resource "google_compute_backend_service" "home" {
  name        = "home%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_health_check.default.id]
  load_balancing_scheme = "INTERNAL_SELF_MANAGED"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func TestAccComputeUrlMap_urlMapHeaderBasedRoutingExample(t *testing.T) {
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
		CheckDestroy: testAccCheckComputeUrlMapDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeUrlMap_urlMapHeaderBasedRoutingExample(context),
			},
			{
				ResourceName:            "google_compute_url_map.urlmap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"default_service"},
			},
		},
	})
}

func testAccComputeUrlMap_urlMapHeaderBasedRoutingExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_url_map" "urlmap" {
  name        = "urlmap%{random_suffix}"
  description = "header-based routing example"
  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts = ["*"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = google_compute_backend_service.default.id

    route_rules {
      priority = 1
      service = google_compute_backend_service.service-a.id
      match_rules {
        prefix_match = "/"
        ignore_case = true
        header_matches {
          header_name = "abtest"
          exact_match = "a"
        }
      }
    }
    route_rules {
      priority = 2
      service = google_compute_backend_service.service-b.id
      match_rules {
        ignore_case = true
        prefix_match = "/"
        header_matches {
          header_name = "abtest"
          exact_match = "b"
        }
      }
    }
  }
}

resource "google_compute_backend_service" "default" {
  name        = "default%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_backend_service" "service-a" {
  name        = "tf-test-service-a%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_backend_service" "service-b" {
  name        = "tf-test-service-b%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, context)
}

func TestAccComputeUrlMap_urlMapParameterBasedRoutingExample(t *testing.T) {
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
		CheckDestroy: testAccCheckComputeUrlMapDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeUrlMap_urlMapParameterBasedRoutingExample(context),
			},
			{
				ResourceName:            "google_compute_url_map.urlmap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"default_service"},
			},
		},
	})
}

func testAccComputeUrlMap_urlMapParameterBasedRoutingExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_url_map" "urlmap" {
  name        = "urlmap%{random_suffix}"
  description = "parameter-based routing example"
  default_service = google_compute_backend_service.default.id

  host_rule {
    hosts = ["*"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name = "allpaths"
    default_service = google_compute_backend_service.default.id

    route_rules {
      priority = 1
      service = google_compute_backend_service.service-a.id
      match_rules {
        prefix_match = "/"
        ignore_case = true
        query_parameter_matches {
          name = "abtest"
          exact_match = "a"
        }
      }
    }
    route_rules {
      priority = 2
      service = google_compute_backend_service.service-b.id
      match_rules {
        ignore_case = true
        prefix_match = "/"
        query_parameter_matches {
          name = "abtest"
          exact_match = "b"
        }
      }
    }
  }
}

resource "google_compute_backend_service" "default" {
  name        = "default%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_backend_service" "service-a" {
  name        = "tf-test-service-a%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_backend_service" "service-b" {
  name        = "tf-test-service-b%{random_suffix}"
  port_name   = "http"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_http_health_check.default.id]
}

resource "google_compute_http_health_check" "default" {
  name               = "tf-test-health-check%{random_suffix}"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
`, context)
}

func testAccCheckComputeUrlMapDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_url_map" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/urlMaps/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeUrlMap still exists at %s", url)
			}
		}

		return nil
	}
}
