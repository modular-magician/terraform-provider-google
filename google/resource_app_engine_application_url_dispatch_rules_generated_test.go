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
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAppEngineApplicationUrlDispatchRules_appEngineApplicationUrlDispatchRulesBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckAppEngineApplicationUrlDispatchRulesDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAppEngineApplicationUrlDispatchRules_appEngineApplicationUrlDispatchRulesBasicExample(context),
			},
			{
				ResourceName:      "google_app_engine_application_url_dispatch_rules.web_service",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAppEngineApplicationUrlDispatchRules_appEngineApplicationUrlDispatchRulesBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_app_engine_application_url_dispatch_rules" "web_service" {
  dispatch_rules {
    domain  = "*"
    path    = "/*"
    service = "default"
  }

  dispatch_rules {
    domain  = "*"
    path    = "/admin/*"
    service = google_app_engine_standard_app_version.admin_v3.service
  }
}

resource "google_app_engine_standard_app_version" "admin_v3" {
  version_id = "v3"
  service    = "admin"
  runtime    = "nodejs10"

  entrypoint {
    shell = "node ./app.js"
  }

  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.name}"
    }
  }

  env_variables = {
    port = "8080"
  }

  noop_on_destroy = true
}

resource "google_storage_bucket" "bucket" {
  name     = "tf-test-appengine-test-bucket%{random_suffix}"
  location = "US"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/appengine/hello-world.zip"
}
`, context)
}

func testAccCheckAppEngineApplicationUrlDispatchRulesDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_app_engine_application_url_dispatch_rules" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			log.Printf("[DEBUG] Ignoring destroy during test")
		}

		return nil
	}
}
