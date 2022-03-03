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

func TestAccHealthcareFhirStore_healthcareFhirStoreBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHealthcareFhirStoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareFhirStore_healthcareFhirStoreBasicExample(context),
			},
			{
				ResourceName:            "google_healthcare_fhir_store.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"self_link", "dataset"},
			},
		},
	})
}

func testAccHealthcareFhirStore_healthcareFhirStoreBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_healthcare_fhir_store" "default" {
  name    = "tf-test-example-fhir-store%{random_suffix}"
  dataset = google_healthcare_dataset.dataset.id
  version = "R4"

  enable_update_create          = false
  disable_referential_integrity = false
  disable_resource_versioning   = false
  enable_history_import         = false

  notification_config {
    pubsub_topic = google_pubsub_topic.topic.id
  }

  validation_config {
    disable_profile_validation        = false
    disable_required_field_validation = false
    disable_reference_type_validation = false
    disable_fhirpath_validation       = false
    enabled_implementation_guides = [
      "https://some.url/to-an-implementation-guide"
    ]
  }

  labels = {
    label1 = "labelvalue1"
  }
}

resource "google_pubsub_topic" "topic" {
  name     = "tf-test-fhir-notifications%{random_suffix}"
}

resource "google_healthcare_dataset" "dataset" {
  name     = "tf-test-example-dataset%{random_suffix}"
  location = "us-central1"
}
`, context)
}

func TestAccHealthcareFhirStore_healthcareFhirStoreStreamingConfigExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHealthcareFhirStoreDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccHealthcareFhirStore_healthcareFhirStoreStreamingConfigExample(context),
			},
			{
				ResourceName:            "google_healthcare_fhir_store.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"self_link", "dataset"},
			},
		},
	})
}

func testAccHealthcareFhirStore_healthcareFhirStoreStreamingConfigExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_healthcare_fhir_store" "default" {
  name    = "tf-test-example-fhir-store%{random_suffix}"
  dataset = google_healthcare_dataset.dataset.id
  version = "R4"

  enable_update_create          = false
  disable_referential_integrity = false
  disable_resource_versioning   = false
  enable_history_import         = false

  labels = {
    label1 = "labelvalue1"
  }

  stream_configs {
    resource_types = ["Observation"]
    bigquery_destination {
      dataset_uri = "bq://${google_bigquery_dataset.bq_dataset.project}.${google_bigquery_dataset.bq_dataset.dataset_id}"
      schema_config {
        recursive_structure_depth = 3
      }
    }
  }
}

resource "google_pubsub_topic" "topic" {
  name     = "tf-test-fhir-notifications%{random_suffix}"
}

resource "google_healthcare_dataset" "dataset" {
  name     = "tf-test-example-dataset%{random_suffix}"
  location = "us-central1"
}

resource "google_bigquery_dataset" "bq_dataset" {
  dataset_id    = "tf_test_bq_example_dataset%{random_suffix}"
  friendly_name = "test"
  description   = "This is a test description"
  location      = "US"
  delete_contents_on_destroy = true
}
`, context)
}

func testAccCheckHealthcareFhirStoreDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_healthcare_fhir_store" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{HealthcareBasePath}}{{dataset}}/fhirStores/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("HealthcareFhirStore still exists at %s", url)
			}
		}

		return nil
	}
}
