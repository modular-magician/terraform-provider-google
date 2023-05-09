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

func TestAccComputeBackendBucket_backendBucketBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendBucketDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendBucket_backendBucketBasicExample(context),
			},
			{
				ResourceName:      "google_compute_backend_bucket.image_backend",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendBucket_backendBucketBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_bucket" "image_backend" {
  name        = "tf-test-image-backend-bucket%{random_suffix}"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.image_bucket.name
  enable_cdn  = true
}

resource "google_storage_bucket" "image_bucket" {
  name     = "tf-test-image-store-bucket%{random_suffix}"
  location = "EU"
}
`, context)
}

func TestAccComputeBackendBucket_backendBucketFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendBucketDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendBucket_backendBucketFullExample(context),
			},
			{
				ResourceName:      "google_compute_backend_bucket.image_backend_full",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendBucket_backendBucketFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_bucket" "image_backend_full" {
  name        = "tf-test-image-backend-bucket-full%{random_suffix}"
  description = "Contains beautiful beta mages"
  bucket_name = google_storage_bucket.image_backend_full.name
  enable_cdn  = true
  cdn_policy {
    cache_mode = "CACHE_ALL_STATIC"
    default_ttl = 3600
    client_ttl  = 7200
    max_ttl     = 10800
    negative_caching = true
  }
  custom_response_headers = [
    "X-Client-Geo-Location:{client_region},{client_city}",
    "X-Tested-By:Magic-Modules"
  ]
}

resource "google_storage_bucket" "image_backend_full" {
  name     = "tf-test-image-store-bucket-full%{random_suffix}"
  location = "EU"
}
`, context)
}

func TestAccComputeBackendBucket_backendBucketSecurityPolicyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendBucketDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendBucket_backendBucketSecurityPolicyExample(context),
			},
			{
				ResourceName:      "google_compute_backend_bucket.image_backend",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendBucket_backendBucketSecurityPolicyExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_bucket" "image_backend" {
  name        = "tf-test-image-backend-bucket%{random_suffix}"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.image_backend.name
  enable_cdn  = true
  edge_security_policy = google_compute_security_policy.policy.id
}

resource "google_storage_bucket" "image_backend" {
  name     = "tf-test-image-store-bucket%{random_suffix}"
  location = "EU"
}

resource "google_compute_security_policy" "policy" {
  name        = "tf-test-image-store-bucket%{random_suffix}"
  description = "basic security policy"
	type = "CLOUD_ARMOR_EDGE"
}
`, context)
}

func TestAccComputeBackendBucket_backendBucketQueryStringWhitelistExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendBucketDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendBucket_backendBucketQueryStringWhitelistExample(context),
			},
			{
				ResourceName:      "google_compute_backend_bucket.image_backend",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendBucket_backendBucketQueryStringWhitelistExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_bucket" "image_backend" {
  name        = "tf-test-image-backend-bucket%{random_suffix}"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.image_bucket.name
  enable_cdn  = true
  cdn_policy {
    cache_key_policy {
        query_string_whitelist = ["image-version"]
    }
  }
}

resource "google_storage_bucket" "image_bucket" {
  name     = "tf-test-image-backend-bucket%{random_suffix}"
  location = "EU"
}
`, context)
}

func TestAccComputeBackendBucket_backendBucketIncludeHttpHeadersExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendBucketDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendBucket_backendBucketIncludeHttpHeadersExample(context),
			},
			{
				ResourceName:      "google_compute_backend_bucket.image_backend",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendBucket_backendBucketIncludeHttpHeadersExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_bucket" "image_backend" {
  name        = "tf-test-image-backend-bucket%{random_suffix}"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.image_bucket.name
  enable_cdn  = true
  cdn_policy {
    cache_key_policy {
        include_http_headers = ["X-My-Header-Field"]
    }
  }
}

resource "google_storage_bucket" "image_bucket" {
  name     = "tf-test-image-backend-bucket%{random_suffix}"
  location = "EU"
}
`, context)
}

func TestAccComputeBackendBucket_externalCdnLbWithBackendBucketExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendBucketDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendBucket_externalCdnLbWithBackendBucketExample(context),
			},
			{
				ResourceName:      "google_compute_backend_bucket.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendBucket_externalCdnLbWithBackendBucketExample(context map[string]interface{}) string {
	return Nprintf(`
# CDN load balancer with Cloud bucket as backend

# Cloud Storage bucket
resource "google_storage_bucket" "default" {
  name                        = "tf-test-my-bucket%{random_suffix}"
  location                    = "us-east1"
  uniform_bucket_level_access = true
  storage_class               = "STANDARD"
  // delete bucket and contents on destroy.
  force_destroy = true
  // Assign specialty files
  website {
    main_page_suffix = "index.html"
    not_found_page   = "404.html"
  }
}


# make bucket public
resource "google_storage_bucket_iam_member" "default" {
  bucket = google_storage_bucket.default.name
  role   = "roles/storage.objectViewer"
  member = "allUsers"
}

resource "google_storage_bucket_object" "index_page" {
  name    = "tf-test-index-page%{random_suffix}"
  bucket  = google_storage_bucket.default.name
  content = <<-EOT
    <html><body>
    <h1>Congratulations on setting up Google Cloud CDN with Storage backend!</h1>
    </body></html>
  EOT
}

resource "google_storage_bucket_object" "error_page" {
  name    = "tf-test-404-page%{random_suffix}"
  bucket  = google_storage_bucket.default.name
  content = <<-EOT
    <html><body>
    <h1>404 Error: Object you are looking for is no longer available!</h1>
    </body></html>
  EOT
}

# image object for testing, try to access http://<your_lb_ip_address>/test.jpg
resource "google_storage_bucket_object" "test_image" {
  name = "tf-test-test-object%{random_suffix}"
  # Uncomment and add valid path to an object.
  #  source       = "/path/to/an/object"
  #  content_type = "image/jpeg"

  # Delete after uncommenting above source and content_type attributes
  content      = "Data as string to be uploaded"
  content_type = "text/plain"

  bucket = google_storage_bucket.default.name
}

# reserve IP address
resource "google_compute_global_address" "default" {
  name = "tf-test-example-ip%{random_suffix}"
}

# forwarding rule
resource "google_compute_global_forwarding_rule" "default" {
  name                  = "tf-test-http-lb-forwarding-rule%{random_suffix}"
  ip_protocol           = "TCP"
  load_balancing_scheme = "EXTERNAL"
  port_range            = "80"
  target                = google_compute_target_http_proxy.default.id
  ip_address            = google_compute_global_address.default.id
}

# http proxy
resource "google_compute_target_http_proxy" "default" {
  name    = "tf-test-http-lb-proxy%{random_suffix}"
  url_map = google_compute_url_map.default.id
}

# url map
resource "google_compute_url_map" "default" {
  name            = "tf-test-http-lb%{random_suffix}"
  default_service = google_compute_backend_bucket.default.id
}

# backend bucket with CDN policy with default ttl settings
resource "google_compute_backend_bucket" "default" {
  name        = "tf-test-cat-backend-bucket%{random_suffix}"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.default.name
  enable_cdn  = true
  cdn_policy {
    cache_mode        = "CACHE_ALL_STATIC"
    client_ttl        = 3600
    default_ttl       = 3600
    max_ttl           = 86400
    negative_caching  = true
    serve_while_stale = 86400
  }
}
`, context)
}

func TestAccComputeBackendBucket_backendBucketBypassCacheExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendBucketDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendBucket_backendBucketBypassCacheExample(context),
			},
			{
				ResourceName:      "google_compute_backend_bucket.image_backend",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendBucket_backendBucketBypassCacheExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_bucket" "image_backend" {
  name        = "tf-test-image-backend-bucket%{random_suffix}"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.image_bucket.name
  enable_cdn  = true
  cdn_policy {
    bypass_cache_on_request_headers {
      header_name = "test"
    }
  }
}

resource "google_storage_bucket" "image_bucket" {
  name     = "tf-test-image-store-bucket%{random_suffix}"
  location = "EU"
}
`, context)
}

func TestAccComputeBackendBucket_backendBucketCoalescingExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeBackendBucketDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeBackendBucket_backendBucketCoalescingExample(context),
			},
			{
				ResourceName:      "google_compute_backend_bucket.image_backend",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeBackendBucket_backendBucketCoalescingExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_backend_bucket" "image_backend" {
  name        = "tf-test-image-backend-bucket%{random_suffix}"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.image_bucket.name
  enable_cdn  = true
  cdn_policy {
    request_coalescing = true
  }
}

resource "google_storage_bucket" "image_bucket" {
  name     = "tf-test-image-store-bucket%{random_suffix}"
  location = "EU"
}
`, context)
}

func testAccCheckComputeBackendBucketDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_backend_bucket" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/backendBuckets/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeBackendBucket still exists at %s", url)
			}
		}

		return nil
	}
}
