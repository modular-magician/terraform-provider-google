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

func TestAccCertificateManagerCertificate_certificateManagerGoogleManagedCertificateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCertificateManagerCertificateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateManagerCertificate_certificateManagerGoogleManagedCertificateExample(context),
			},
			{
				ResourceName:            "google_certificate_manager_certificate.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"self_managed", "name"},
			},
		},
	})
}

func testAccCertificateManagerCertificate_certificateManagerGoogleManagedCertificateExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_certificate_manager_certificate" "default" {
  name        = "tf-test-dns-cert%{random_suffix}"
  description = "The default cert"
  scope       = "EDGE_CACHE"
  managed {
    domains = [
      google_certificate_manager_dns_authorization.instance.domain,
      google_certificate_manager_dns_authorization.instance2.domain,
      ]
    dns_authorizations = [
      google_certificate_manager_dns_authorization.instance.id,
      google_certificate_manager_dns_authorization.instance2.id,
      ]
  }
}


resource "google_certificate_manager_dns_authorization" "instance" {
  name        = "tf-test-dns-auth%{random_suffix}"
  description = "The default dnss"
  domain      = "subdomain%{random_suffix}.hashicorptest.com"
}

resource "google_certificate_manager_dns_authorization" "instance2" {
  name        = "tf-test-dns-auth2%{random_suffix}"
  description = "The default dnss"
  domain      = "subdomain2%{random_suffix}.hashicorptest.com"
}
`, context)
}

func TestAccCertificateManagerCertificate_certificateManagerSelfManagedCertificateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCertificateManagerCertificateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCertificateManagerCertificate_certificateManagerSelfManagedCertificateExample(context),
			},
			{
				ResourceName:            "google_certificate_manager_certificate.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"self_managed", "name"},
			},
		},
	})
}

func testAccCertificateManagerCertificate_certificateManagerSelfManagedCertificateExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_certificate_manager_certificate" "default" {
  name        = "tf-test-self-managed-cert%{random_suffix}"
  description = "The default cert"
  scope       = "EDGE_CACHE"
  self_managed {
    pem_certificate = file("test-fixtures/certificatemanager/cert.pem")
    pem_private_key = file("test-fixtures/certificatemanager/private-key.pem")
  }
}
`, context)
}

func testAccCheckCertificateManagerCertificateDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_certificate_manager_certificate" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{CertificateManagerBasePath}}projects/{{project}}/locations/global/certificates/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("CertificateManagerCertificate still exists at %s", url)
			}
		}

		return nil
	}
}
