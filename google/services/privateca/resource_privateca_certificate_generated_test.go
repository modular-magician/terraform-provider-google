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

package privateca_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccPrivatecaCertificate_privatecaCertificateConfigExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPrivatecaCertificateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCertificate_privatecaCertificateConfigExample(context),
			},
			{
				ResourceName:            "google_privateca_certificate.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"pool", "name", "location", "certificate_authority", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccPrivatecaCertificate_privatecaCertificateConfigExample(context map[string]interface{}) string {
	return acctest.Nprintf(`

resource "google_privateca_ca_pool" "default" {
  location = "us-central1"
  name = "tf-test-my-pool%{random_suffix}"
  tier = "ENTERPRISE"
}

resource "google_privateca_certificate_authority" "default" {
  location = "us-central1"
  pool = google_privateca_ca_pool.default.name
  certificate_authority_id = "my-authority"
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
        common_name = "my-certificate-authority"
      }
      subject_alt_name {
        dns_names = ["hashicorp.com"]
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

  // Disable CA deletion related safe checks for easier cleanup.
  deletion_protection                    = false
  skip_grace_period                      = true
  ignore_active_certificates_on_deletion = true
}

resource "google_privateca_certificate" "default" {
  location = "us-central1"
  pool = google_privateca_ca_pool.default.name
  certificate_authority = google_privateca_certificate_authority.default.certificate_authority_id
  lifetime = "86000s"
  name = "tf-test-my-certificate%{random_suffix}"
  config {
    subject_config  {
      subject {
        common_name = "san1.example.com"
        country_code = "us"
        organization = "google"
        organizational_unit = "enterprise"
        locality = "mountain view"
        province = "california"
        street_address = "1600 amphitheatre parkway"
      } 
      subject_alt_name {
        email_addresses = ["email@example.com"]
        ip_addresses = ["127.0.0.1"]
        uris = ["http://www.ietf.org/rfc/rfc3986.txt"]
      }
    }
    x509_config {
      ca_options {
        is_ca = true
      }
      key_usage {
        base_key_usage {
          crl_sign = false
          decipher_only = false
        }
        extended_key_usage {
          server_auth = false
        }
      }
      name_constraints {
        critical                  = true
        permitted_dns_names       = ["*.example.com"]
        excluded_dns_names        = ["*.deny.example.com"]
        permitted_ip_ranges       = ["10.0.0.0/8"]
        excluded_ip_ranges        = ["10.1.1.0/24"]
        permitted_email_addresses = [".example.com"]
        excluded_email_addresses  = [".deny.example.com"]
        permitted_uris            = [".example.com"]
        excluded_uris             = [".deny.example.com"]
      }
    }
    public_key {
      format = "PEM"
      key = filebase64("test-fixtures/rsa_public.pem")
    }
  }
}
`, context)
}

func TestAccPrivatecaCertificate_privatecaCertificateWithTemplateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPrivatecaCertificateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCertificate_privatecaCertificateWithTemplateExample(context),
			},
			{
				ResourceName:            "google_privateca_certificate.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"pool", "name", "location", "certificate_authority", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccPrivatecaCertificate_privatecaCertificateWithTemplateExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  location = "us-central1"
  name = "tf-test-my-pool%{random_suffix}"
  tier = "ENTERPRISE"
}

resource "google_privateca_certificate_template" "default" {
  location    = "us-central1"
  name = "tf-test-my-certificate-template%{random_suffix}"
  description = "An updated sample certificate template"

  identity_constraints {
    allow_subject_alt_names_passthrough = true
    allow_subject_passthrough           = true

    cel_expression {
      description = "Always true"
      expression  = "true"
      location    = "any.file.anywhere"
      title       = "Sample expression"
    }
  }

  passthrough_extensions {
    additional_extensions {
      object_id_path = [1, 6]
    }

    known_extensions = ["EXTENDED_KEY_USAGE"]
  }

  predefined_values {
    additional_extensions {
      object_id {
        object_id_path = [1, 6]
      }

      value    = "c3RyaW5nCg=="
      critical = true
    }

    aia_ocsp_servers = ["string"]

    ca_options {
      is_ca                  = false
      max_issuer_path_length = 6
    }

    key_usage {
      base_key_usage {
        cert_sign          = false
        content_commitment = true
        crl_sign           = false
        data_encipherment  = true
        decipher_only      = true
        digital_signature  = true
        encipher_only      = true
        key_agreement      = true
        key_encipherment   = true
      }

      extended_key_usage {
        client_auth      = true
        code_signing     = true
        email_protection = true
        ocsp_signing     = true
        server_auth      = true
        time_stamping    = true
      }

      unknown_extended_key_usages {
        object_id_path = [1, 6]
      }
    }

    policy_ids {
      object_id_path = [1, 6]
    }
  }
}

resource "google_privateca_certificate_authority" "default" {
  location = "us-central1"
  pool = google_privateca_ca_pool.default.name
  certificate_authority_id = "my-authority"
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
        common_name = "my-certificate-authority"
      }
      subject_alt_name {
        dns_names = ["hashicorp.com"]
      }
    }
    x509_config {
      ca_options {
        # is_ca *MUST* be true for certificate authorities
        is_ca = true
      }
      key_usage {
        base_key_usage {
          # cert_sign and crl_sign *MUST* be true for certificate authorities
          cert_sign = true
          crl_sign = true
        }
        extended_key_usage {
          server_auth = false
        }
      }
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }

  // Disable CA deletion related safe checks for easier cleanup.
  deletion_protection                    = false
  skip_grace_period                      = true
  ignore_active_certificates_on_deletion = true
}


resource "google_privateca_certificate" "default" {
  location = "us-central1"
  pool = google_privateca_ca_pool.default.name
  certificate_authority = google_privateca_certificate_authority.default.certificate_authority_id
  name = "tf-test-my-certificate%{random_suffix}"
  lifetime = "860s"
  pem_csr = file("test-fixtures/rsa_csr.pem")
  certificate_template = google_privateca_certificate_template.default.id
}
`, context)
}

func TestAccPrivatecaCertificate_privatecaCertificateCsrExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPrivatecaCertificateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCertificate_privatecaCertificateCsrExample(context),
			},
			{
				ResourceName:            "google_privateca_certificate.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"pool", "name", "location", "certificate_authority", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccPrivatecaCertificate_privatecaCertificateCsrExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  location = "us-central1"
  name = "tf-test-my-pool%{random_suffix}"
  tier = "ENTERPRISE"
}

resource "google_privateca_certificate_authority" "default" {
  location = "us-central1"
  pool = google_privateca_ca_pool.default.name
  certificate_authority_id = "my-authority"
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
        common_name = "my-certificate-authority"
      }
      subject_alt_name {
        dns_names = ["hashicorp.com"]
      }
    }
    x509_config {
      ca_options {
        # is_ca *MUST* be true for certificate authorities
        is_ca = true
      }
      key_usage {
        base_key_usage {
          # cert_sign and crl_sign *MUST* be true for certificate authorities
          cert_sign = true
          crl_sign = true
        }
        extended_key_usage {
          server_auth = false
        }
      }
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }

  // Disable CA deletion related safe checks for easier cleanup.
  deletion_protection                    = false
  skip_grace_period                      = true
  ignore_active_certificates_on_deletion = true
}


resource "google_privateca_certificate" "default" {
  location = "us-central1"
  pool = google_privateca_ca_pool.default.name
  certificate_authority = google_privateca_certificate_authority.default.certificate_authority_id
  name = "tf-test-my-certificate%{random_suffix}"
  lifetime = "860s"
  pem_csr = file("test-fixtures/rsa_csr.pem")
}
`, context)
}

func TestAccPrivatecaCertificate_privatecaCertificateNoAuthorityExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPrivatecaCertificateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCertificate_privatecaCertificateNoAuthorityExample(context),
			},
			{
				ResourceName:            "google_privateca_certificate.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"pool", "name", "location", "certificate_authority", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccPrivatecaCertificate_privatecaCertificateNoAuthorityExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_ca_pool" "default" {
  location = "us-central1"
  name = "tf-test-my-pool%{random_suffix}"
  tier = "ENTERPRISE"
}

resource "google_privateca_certificate_authority" "default" {
  location = "us-central1"
  pool = google_privateca_ca_pool.default.name
  certificate_authority_id = "my-authority"
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
        common_name = "my-certificate-authority"
      }
      subject_alt_name {
        dns_names = ["hashicorp.com"]
      }
    }
    x509_config {
      ca_options {
        is_ca = true
      }
      key_usage {
        base_key_usage {
          digital_signature = true
          cert_sign = true
          crl_sign = true
        }
        extended_key_usage {
          server_auth = true
        }
      }
    }
  }
  lifetime = "86400s"
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }

  // Disable CA deletion related safe checks for easier cleanup.
  deletion_protection                    = false
  skip_grace_period                      = true
  ignore_active_certificates_on_deletion = true
}


resource "google_privateca_certificate" "default" {
  location = "us-central1"
  pool = google_privateca_ca_pool.default.name
  name = "tf-test-my-certificate%{random_suffix}"
  lifetime = "860s"
  config {
    subject_config  {
      subject {
        common_name = "san1.example.com"
        country_code = "us"
        organization = "google"
        organizational_unit = "enterprise"
        locality = "mountain view"
        province = "california"
        street_address = "1600 amphitheatre parkway"
        postal_code = "94109"
      }
    }
    x509_config {
      ca_options {
        is_ca = false
      }
      key_usage {
        base_key_usage {
          crl_sign = true
        }
        extended_key_usage {
          server_auth = true
        }
      }
    }
    public_key {
      format = "PEM"
      key = filebase64("test-fixtures/rsa_public.pem")
    }
  }
  // Certificates require an authority to exist in the pool, though they don't
  // need to be explicitly connected to it
  depends_on = [google_privateca_certificate_authority.default]
}
`, context)
}

func testAccCheckPrivatecaCertificateDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_privateca_certificate" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{PrivatecaBasePath}}projects/{{project}}/locations/{{location}}/caPools/{{pool}}/certificates/{{name}}")
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
				return fmt.Errorf("PrivatecaCertificate still exists at %s", url)
			}
		}

		return nil
	}
}
