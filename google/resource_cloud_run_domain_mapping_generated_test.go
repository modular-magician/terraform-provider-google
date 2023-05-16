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

func TestAccCloudRunDomainMapping_cloudRunDomainMappingBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"namespace":     acctest.GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudRunDomainMappingDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunDomainMapping_cloudRunDomainMappingBasicExample(context),
			},
			{
				ResourceName:            "google_cloud_run_domain_mapping.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunDomainMapping_cloudRunDomainMappingBasicExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_cloud_run_service" "default" {
  name     = "tf-test-cloudrun-srv%{random_suffix}"
  location = "us-central1"

  metadata {
    namespace = "%{namespace}"
  }

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }
}

resource "google_cloud_run_domain_mapping" "default" {
  location = "us-central1"
  name     = "tf-test-domain%{random_suffix}.gcp.tfacc.hashicorptest.com"

  metadata {
    namespace = "%{namespace}"
  }

  spec {
    route_name = google_cloud_run_service.default.name
  }
}
`, context)
}

func testAccCheckCloudRunDomainMappingDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloud_run_domain_mapping" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{CloudRunBasePath}}apis/domains.cloudrun.com/v1/namespaces/{{project}}/domainmappings/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:               config,
				Method:               "GET",
				Project:              billingProject,
				RawURL:               url,
				UserAgent:            config.UserAgent,
				ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsCloudRunCreationConflict},
			})
			if err == nil {
				return fmt.Errorf("CloudRunDomainMapping still exists at %s", url)
			}
		}

		return nil
	}
}
