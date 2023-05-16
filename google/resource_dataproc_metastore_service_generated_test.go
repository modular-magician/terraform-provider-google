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

func TestAccDataprocMetastoreService_dataprocMetastoreServiceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocMetastoreServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocMetastoreService_dataprocMetastoreServiceBasicExample(context),
			},
			{
				ResourceName:            "google_dataproc_metastore_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_id", "location"},
			},
		},
	})
}

func testAccDataprocMetastoreService_dataprocMetastoreServiceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_metastore_service" "default" {
  service_id = "tf-test-metastore-srv%{random_suffix}"
  location   = "us-central1"
  port       = 9080
  tier       = "DEVELOPER"

  maintenance_window {
    hour_of_day = 2
    day_of_week = "SUNDAY"
  }

  hive_metastore_config {
    version = "2.3.6"
  }
}
`, context)
}

func TestAccDataprocMetastoreService_dataprocMetastoreServiceCmekTestExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocMetastoreServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocMetastoreService_dataprocMetastoreServiceCmekTestExample(context),
			},
			{
				ResourceName:            "google_dataproc_metastore_service.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_id", "location"},
			},
		},
	})
}

func testAccDataprocMetastoreService_dataprocMetastoreServiceCmekTestExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_project" "project" {}

data "google_storage_project_service_account" "gcs_account" {}


resource "google_dataproc_metastore_service" "default" {
  service_id = "tf-test-example-service%{random_suffix}"
  location   = "us-central1"

  encryption_config {
    kms_key = google_kms_crypto_key.crypto_key.id
  }

  hive_metastore_config {
    version = "3.1.2"
  }

  depends_on = [google_kms_crypto_key_iam_binding.crypto_key_binding]
}

resource "google_kms_crypto_key" "crypto_key" {
  name     = "tf-test-example-key%{random_suffix}"
  key_ring = google_kms_key_ring.key_ring.id

  purpose  = "ENCRYPT_DECRYPT"
}

resource "google_kms_key_ring" "key_ring" {
  name     = "tf-test-example-keyring%{random_suffix}"
  location = "us-central1"
}

resource "google_kms_crypto_key_iam_binding" "crypto_key_binding" {
  crypto_key_id = google_kms_crypto_key.crypto_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"

  members = [
    "serviceAccount:service-${data.google_project.project.number}@gcp-sa-metastore.iam.gserviceaccount.com",
    "serviceAccount:${data.google_storage_project_service_account.gcs_account.email_address}"
  ]
}
`, context)
}

func TestAccDataprocMetastoreService_dataprocMetastoreServiceTelemetryExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocMetastoreServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocMetastoreService_dataprocMetastoreServiceTelemetryExample(context),
			},
			{
				ResourceName:            "google_dataproc_metastore_service.telemetry",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_id", "location"},
			},
		},
	})
}

func testAccDataprocMetastoreService_dataprocMetastoreServiceTelemetryExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_metastore_service" "telemetry" {
  service_id = "telemetry%{random_suffix}"
  location   = "us-central1"
  port       = 9080
  tier       = "DEVELOPER"

  hive_metastore_config {
    version = "3.1.2"
  }

  telemetry_config {
    log_format = "LEGACY"
  }
}
`, context)
}

func testAccCheckDataprocMetastoreServiceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dataproc_metastore_service" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DataprocMetastoreBasePath}}projects/{{project}}/locations/{{location}}/services/{{service_id}}")
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
				return fmt.Errorf("DataprocMetastoreService still exists at %s", url)
			}
		}

		return nil
	}
}
