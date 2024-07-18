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

package securesourcemanager_test

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

func TestAccSecureSourceManagerInstance_secureSourceManagerInstanceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecureSourceManagerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecureSourceManagerInstance_secureSourceManagerInstanceBasicExample(context),
			},
			{
				ResourceName:            "google_secure_source_manager_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccSecureSourceManagerInstance_secureSourceManagerInstanceBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secure_source_manager_instance" "default" {
    location = "us-central1"
    instance_id = "tf-test-my-instance%{random_suffix}"
    labels = {
      "foo" = "bar"
    }
}
`, context)
}

func TestAccSecureSourceManagerInstance_secureSourceManagerInstanceCmekExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecureSourceManagerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecureSourceManagerInstance_secureSourceManagerInstanceCmekExample(context),
			},
			{
				ResourceName:            "google_secure_source_manager_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccSecureSourceManagerInstance_secureSourceManagerInstanceCmekExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_kms_key_ring" "key_ring" {
  name     = "tf-test-my-keyring%{random_suffix}"
  location = "us-central1"
}

resource "google_kms_crypto_key" "crypto_key" {
  name     = "tf-test-my-key%{random_suffix}"
  key_ring = google_kms_key_ring.key_ring.id
}

resource "google_kms_crypto_key_iam_member" "crypto_key_binding" {
  crypto_key_id = google_kms_crypto_key.crypto_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"

  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-sourcemanager.iam.gserviceaccount.com"
}

resource "google_secure_source_manager_instance" "default" {
    location = "us-central1"
    instance_id = "tf-test-my-instance%{random_suffix}"
    kms_key = google_kms_crypto_key.crypto_key.id

    depends_on = [
      google_kms_crypto_key_iam_member.crypto_key_binding
    ]
}

data "google_project" "project" {}
`, context)
}

func TestAccSecureSourceManagerInstance_secureSourceManagerInstancePrivateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckSecureSourceManagerInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecureSourceManagerInstance_secureSourceManagerInstancePrivateExample(context),
			},
			{
				ResourceName:            "google_secure_source_manager_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"instance_id", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccSecureSourceManagerInstance_secureSourceManagerInstancePrivateExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "ca_pool" {
  name     = "tf-test-ca-pool%{random_suffix}"
  location = "us-central1"
  tier     = "ENTERPRISE"
  publishing_options {
    publish_ca_cert = true
    publish_crl     = true
  }
}

resource "google_privateca_certificate_authority" "root_ca" {
  pool                     = google_privateca_ca_pool.ca_pool.name
  certificate_authority_id = "tf-test-root-ca%{random_suffix}"
  location                 = "us-central1"
  config {
    subject_config {
      subject {
        organization = "google"
        common_name = "my-certificate-authority"
      }
    }
    x509_config {
      ca_options {
        is_ca = true
      }
      key_usage {
        base_key_usage {
          cert_sign = true
          crl_sign = true
        }
        extended_key_usage {
          server_auth = true
        }
      }
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }

  // Disable deletion protections for easier test cleanup purposes
  deletion_protection = false
  ignore_active_certificates_on_deletion = true
  skip_grace_period = true
}

resource "google_privateca_ca_pool_iam_binding" "ca_pool_binding" {
  ca_pool = google_privateca_ca_pool.ca_pool.id
  role = "roles/privateca.certificateRequester"

  members = [
    "serviceAccount:service-${data.google_project.project.number}@gcp-sa-sourcemanager.iam.gserviceaccount.com"
  ]
}

resource "google_secure_source_manager_instance" "default" {
  instance_id = "tf-test-my-instance%{random_suffix}"
  location = "us-central1"
  private_config {
    is_private = true
    ca_pool = google_privateca_ca_pool.ca_pool.id
  }
  depends_on = [
    google_privateca_certificate_authority.root_ca,
    time_sleep.wait_60_seconds
  ]
}

# ca pool IAM permissions can take time to propagate
resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_privateca_ca_pool_iam_binding.ca_pool_binding]

  create_duration = "60s"
}

data "google_project" "project" {}
`, context)
}

func testAccCheckSecureSourceManagerInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_secure_source_manager_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{SecureSourceManagerBasePath}}projects/{{project}}/locations/{{location}}/instances/{{instance_id}}")
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
				return fmt.Errorf("SecureSourceManagerInstance still exists at %s", url)
			}
		}

		return nil
	}
}
