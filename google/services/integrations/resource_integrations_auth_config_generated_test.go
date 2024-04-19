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

package integrations_test

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

func TestAccIntegrationsAuthConfig_integrationsAuthConfigAdvanceExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIntegrationsAuthConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationsAuthConfig_integrationsAuthConfigAdvanceExample(context),
			},
			{
				ResourceName:            "google_integrations_auth_config.advance_example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_certificate", "location"},
			},
		},
	})
}

func testAccIntegrationsAuthConfig_integrationsAuthConfigAdvanceExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_integrations_client" "client" {
  location = "asia-east2"
  provision_gmek = true
}

resource "google_integrations_auth_config" "advance_example" {
    location = "asia-east2"
    display_name = "tf-test-test-authconfig%{random_suffix}"
    description = "Test auth config created via terraform"
    visibility = "CLIENT_VISIBLE"
    expiry_notification_duration = ["3.500s"]
    override_valid_time = "2014-10-02T15:01:23Z"
    decrypted_credential {
        credential_type = "USERNAME_AND_PASSWORD"
        username_and_password {
            username = "test-username"
            password = "test-password"
        }
    }
    depends_on = [google_integrations_client.client]
}
`, context)
}

func TestAccIntegrationsAuthConfig_integrationsAuthConfigUsernameAndPasswordExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIntegrationsAuthConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationsAuthConfig_integrationsAuthConfigUsernameAndPasswordExample(context),
			},
			{
				ResourceName:            "google_integrations_auth_config.username_and_password_example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_certificate", "location"},
			},
		},
	})
}

func testAccIntegrationsAuthConfig_integrationsAuthConfigUsernameAndPasswordExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_integrations_client" "client" {
  location = "northamerica-northeast2"
  provision_gmek = true
}

resource "google_integrations_auth_config" "username_and_password_example" {
    location = "northamerica-northeast2"
    display_name = "tf-test-test-authconfig-username-and-password%{random_suffix}"
    description = "Test auth config created via terraform"
    decrypted_credential {
        credential_type = "USERNAME_AND_PASSWORD"
        username_and_password {
            username = "test-username"
            password = "test-password"
        }
    }
    depends_on = [google_integrations_client.client]
}
`, context)
}

func TestAccIntegrationsAuthConfig_integrationsAuthConfigOauth2AuthorizationCodeExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIntegrationsAuthConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationsAuthConfig_integrationsAuthConfigOauth2AuthorizationCodeExample(context),
			},
			{
				ResourceName:            "google_integrations_auth_config.oauth2_authotization_code_example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_certificate", "location"},
			},
		},
	})
}

func testAccIntegrationsAuthConfig_integrationsAuthConfigOauth2AuthorizationCodeExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_integrations_client" "client" {
  location = "asia-east1"
  provision_gmek = true
}

resource "google_integrations_auth_config" "oauth2_authotization_code_example" {
    location = "asia-east1"
    display_name = "tf-test-test-authconfig-oauth2-authorization-code%{random_suffix}"
    description = "Test auth config created via terraform"
    decrypted_credential {
        credential_type = "OAUTH2_AUTHORIZATION_CODE"
        oauth2_authorization_code {
            client_id = "Kf7utRvgr95oGO5YMmhFOLo8"
            client_secret = "D-XXFDDMLrg2deDgczzHTBwC3p16wRK1rdKuuoFdWqO0wliJ"
            scope = "photo offline_access"
            auth_endpoint = "https://authorization-server.com/authorize"
            token_endpoint = "https://authorization-server.com/token"
        }
    }
    depends_on = [google_integrations_client.client]
}
`, context)
}

func TestAccIntegrationsAuthConfig_integrationsAuthConfigOauth2ClientCredentialsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIntegrationsAuthConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationsAuthConfig_integrationsAuthConfigOauth2ClientCredentialsExample(context),
			},
			{
				ResourceName:            "google_integrations_auth_config.oauth2_client_credentials_example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_certificate", "location"},
			},
		},
	})
}

func testAccIntegrationsAuthConfig_integrationsAuthConfigOauth2ClientCredentialsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_integrations_client" "client" {
  location = "southamerica-east1"
  provision_gmek = true
}

resource "google_integrations_auth_config" "oauth2_client_credentials_example" {
    location = "southamerica-east1"
    display_name = "tf-test-test-authconfig-oauth2-client-credentials%{random_suffix}"
    description = "Test auth config created via terraform"
    decrypted_credential {
        credential_type = "OAUTH2_CLIENT_CREDENTIALS"
        oauth2_client_credentials {
            client_id = "demo-backend-client"
            client_secret = "MJlO3binatD9jk1"
            scope = "read"
            token_endpoint = "https://login-demo.curity.io/oauth/v2/oauth-token"
            request_type = "ENCODED_HEADER"
            token_params {
                entries {
                    key {
                        literal_value {
                            string_value = "string-key"
                        }
                    }
                    value {
                        literal_value {
                            string_value = "string-value"
                        }
                    }
                }
            }
        }
    }
    depends_on = [google_integrations_client.client]
}
`, context)
}

func TestAccIntegrationsAuthConfig_integrationsAuthConfigJwtExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIntegrationsAuthConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationsAuthConfig_integrationsAuthConfigJwtExample(context),
			},
			{
				ResourceName:            "google_integrations_auth_config.jwt_example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_certificate", "location"},
			},
		},
	})
}

func testAccIntegrationsAuthConfig_integrationsAuthConfigJwtExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_integrations_client" "client" {
  location = "us-west4"
  provision_gmek = true
}

resource "google_integrations_auth_config" "jwt_example" {
    location = "us-west4"
    display_name = "tf-test-test-authconfig-jwt%{random_suffix}"
    description = "Test auth config created via terraform"
    decrypted_credential {
        credential_type = "JWT"
        jwt {
            jwt_header = "{\"alg\": \"HS256\", \"typ\": \"JWT\"}"
            jwt_payload = "{\"sub\": \"1234567890\", \"name\": \"John Doe\", \"iat\": 1516239022}"
            secret = "secret"
        }
    }
    depends_on = [google_integrations_client.client]
}
`, context)
}

func TestAccIntegrationsAuthConfig_integrationsAuthConfigAuthTokenExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIntegrationsAuthConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationsAuthConfig_integrationsAuthConfigAuthTokenExample(context),
			},
			{
				ResourceName:            "google_integrations_auth_config.auth_token_example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_certificate", "location"},
			},
		},
	})
}

func testAccIntegrationsAuthConfig_integrationsAuthConfigAuthTokenExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_integrations_client" "client" {
  location = "us-west2"
  provision_gmek = true
}

resource "google_integrations_auth_config" "auth_token_example" {
    location = "us-west2"
    display_name = "tf-test-test-authconfig-auth-token%{random_suffix}"
    description = "Test auth config created via terraform"
    decrypted_credential {
        credential_type = "AUTH_TOKEN"
        auth_token {
            type = "Basic"
            token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
        }
    }
    depends_on = [google_integrations_client.client]
}
`, context)
}

func TestAccIntegrationsAuthConfig_integrationsAuthConfigServiceAccountExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIntegrationsAuthConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationsAuthConfig_integrationsAuthConfigServiceAccountExample(context),
			},
			{
				ResourceName:            "google_integrations_auth_config.service_account_example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_certificate", "location"},
			},
		},
	})
}

func testAccIntegrationsAuthConfig_integrationsAuthConfigServiceAccountExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_integrations_client" "client" {
  location = "northamerica-northeast1"
  provision_gmek = true
}

resource "google_service_account" "service_account" {
  account_id   = "sa%{random_suffix}"
  display_name = "Service Account"
}

resource "google_integrations_auth_config" "service_account_example" {
    location = "northamerica-northeast1"
    display_name = "tf-test-test-authconfig-service-account%{random_suffix}"
    description = "Test auth config created via terraform"
    decrypted_credential {
        credential_type = "SERVICE_ACCOUNT"
        service_account_credentials {
            service_account = google_service_account.service_account.email
            scope = "https://www.googleapis.com/auth/cloud-platform https://www.googleapis.com/auth/adexchange.buyer https://www.googleapis.com/auth/admob.readonly"
        }
    }
    depends_on = [google_service_account.service_account, google_integrations_client.client]
}
`, context)
}

func TestAccIntegrationsAuthConfig_integrationsAuthConfigOidcTokenExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIntegrationsAuthConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationsAuthConfig_integrationsAuthConfigOidcTokenExample(context),
			},
			{
				ResourceName:            "google_integrations_auth_config.oidc_token_example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_certificate", "location"},
			},
		},
	})
}

func testAccIntegrationsAuthConfig_integrationsAuthConfigOidcTokenExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_integrations_client" "client" {
  location = "us-south1"
  provision_gmek = true
}

resource "google_service_account" "service_account" {
  account_id   = "sa%{random_suffix}"
  display_name = "Service Account"
}

resource "google_integrations_auth_config" "oidc_token_example" {
    location = "us-south1"
    display_name = "tf-test-test-authconfig-oidc-token%{random_suffix}"
    description = "Test auth config created via terraform"
    decrypted_credential {
        credential_type = "OIDC_TOKEN"
        oidc_token {
            service_account_email = google_service_account.service_account.email
            audience = "https://us-south1-project.cloudfunctions.net/functionA 1234987819200.apps.googleusercontent.com"
        }
    }
    depends_on = [google_service_account.service_account, google_integrations_client.client]
}
`, context)
}

func TestAccIntegrationsAuthConfig_integrationsAuthConfigClientCertificateOnlyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckIntegrationsAuthConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationsAuthConfig_integrationsAuthConfigClientCertificateOnlyExample(context),
			},
			{
				ResourceName:            "google_integrations_auth_config.client_certificate_example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_certificate", "location"},
			},
		},
	})
}

func testAccIntegrationsAuthConfig_integrationsAuthConfigClientCertificateOnlyExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_integrations_client" "client" {
  location = "us-west3"
  provision_gmek = true
}

resource "google_integrations_auth_config" "client_certificate_example" {
    location = "us-west3"
    display_name = "tf-test-test-authconfig-client-certificate%{random_suffix}"
    description = "Test auth config created via terraform"
    decrypted_credential {
        credential_type = "CLIENT_CERTIFICATE_ONLY"
    }
    client_certificate {
        ssl_certificate = <<EOT
-----BEGIN CERTIFICATE-----
MIICTTCCAbagAwIBAgIJAPT0tSKNxan/MA0GCSqGSIb3DQEBCwUAMCoxFzAVBgNV
BAoTDkdvb2dsZSBURVNUSU5HMQ8wDQYDVQQDEwZ0ZXN0Q0EwHhcNMTUwMTAxMDAw
MDAwWhcNMjUwMTAxMDAwMDAwWjAuMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElORzET
MBEGA1UEAwwKam9lQGJhbmFuYTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA
vDYFgMgxi5W488d9J7UpCInl0NXmZQpJDEHE4hvkaRlH7pnC71H0DLt0/3zATRP1
JzY2+eqBmbGl4/sgZKYv8UrLnNyQNUTsNx1iZAfPUflf5FwgVsai8BM0pUciq1NB
xD429VFcrGZNucvFLh72RuRFIKH8WUpiK/iZNFkWhZ0CAwEAAaN3MHUwDgYDVR0P
AQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMB
Af8EAjAAMBkGA1UdDgQSBBCVgnFBCWgL/iwCqnGrhTPQMBsGA1UdIwQUMBKAEKey
Um2o4k2WiEVA0ldQvNYwDQYJKoZIhvcNAQELBQADgYEAYK986R4E3L1v+Q6esBtW
JrUwA9UmJRSQr0N5w3o9XzarU37/bkjOP0Fw0k/A6Vv1n3vlciYfBFaBIam1qRHr
5dMsYf4CZS6w50r7hyzqyrwDoyNxkLnd2PdcHT/sym1QmflsjEs7pejtnohO6N2H
wQW6M0H7Zt8claGRla4fKkg=
-----END CERTIFICATE-----
EOT
        encrypted_private_key = <<EOT
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCA/Oj2HXqs5fTk
j/8DrlOQtLG3K9RMsYHvnwICLxkGqVcTfut58hDFLbQM8C3C0ENAKitNJplCJmYG
8VpgZzgq8VxaGnlP/sXUFLMGksd5sATn0sY3SkPndTKk/dqqA4MIh/dYfh19ynEN
hB9Ll/h54Yic2je2Qaxe/uMMu8RODTz3oCn7FcoYpPvfygfU0ntn4IcqH/hts5DG
s+3otJk4entRZglQDxR+sWOsbLtJIQZDP8rH3jDVdl5l3wspgtMTY8b5T5+pLm0p
/OzCmxT0dq/O6BhpxI1xf/zcdRZeWk5DTJxTi5AgPquTlAG/B6A3HkqBJ14hT/Rk
iv7Ma3DLAgMBAAECggEABATkf9VfpiAT9zYdouk50bBpckvymQTyQLD8SlBaX+KY
kgv/pHSXK4Pm4iensrQerFLgfqPA3U+FiqjW5Mv7c1VRK6HJbuVkpdzoXLI9IQsL
vsBY7//9Ajk5P7NokjdB6JPdU/2dHROuQVa59cxPtzpHo0htnPlDOKXfFZZuoZ17
Nr8WQHrHy8P8ABM1tLOzvU9Nlh7TcjQvev+HxkLek4qzYyJ/Ac7XOjg/XKUm1tZk
O3BHr8YLabwyjO7l1t+2b14rUTL/8pfUZnAkEi3FAlPxm3ilftmX65zliC9G4ghk
dr5PByT3DqnuIIglua9bISv1H34ogecd+9a6EU7RxQKBgQC2RPKLounXZo8vYiU4
sFTEvjbs+u9Ypk4OrNLnb8KdacLBUaJGnf++xbBoKpwFCBJfy//fvuQfusYF9Gyn
GxL43tw94C/H5upQYnDsmnQak6TbOu3mA24OGK7Rcq6NEHgeCY4HomutnSiPTZJq
8jlpqgqh1itETe5avgkMNq3zBwKBgQC1KlztGzvbB+rUDc6Kfvk5pUbCSFKMMMa2
NWNXeD6i2iA56zEYSbTjKQ3u9pjUV8LNqAdUFxmbdPxZjheNK2dEm68SVRXPKOeB
EmQT+t/EyW9LqBEA2oZt3h2hXtK8ppJjQm4XUCDs1NphP87eNzx5FLzJWjG8VqDq
jOvApNqPHQKBgDQqlZSbgvvwUYjJOUf5R7mri0LWKwyfRHX0xsQQe43cCC6WM7Cs
Zdbu86dMkqzp+4BJfalHFDl0llp782D8Ybiy6CwZbvNyxptNIW7GYfZ9TVCllBMh
5izIqbgub4DWNtq591l+Bf2BnmstU3uiagYw8awSBP4eo9p6y1IgkDafAoGBAJbi
lIiqEP0IqA06/pWc0Qew3rD7OT0ndqjU6Es2i7xovURf3QDkinJThBZNbdYUzdsp
IgloP9yY33/a90SNLLIYlARJtyNVZxK59X4qiOpF9prlfFvgpOumfbkj15JljTB8
aGKkSvfVA5jRYwLysDwMCHwO0bOR1u3itos5AgsFAoGAKEGms1kuQ5/HyFgSmg9G
wBUzu+5Y08/A37rvyXsR6GjmlZJvULEopJNUNCOOpITNQikXK63sIFry7/59eGv5
UwKadZbfwbVF5ipu59UxfVE3lipf/mYePDqMkHVWv/8p+OnnJt9uKnyW8VSOu5uk
82QF30zbIWDTUjrcugVAs+E=
-----END PRIVATE KEY-----     
EOT
        passphrase = ""
    }
    depends_on = [google_integrations_client.client]
}
`, context)
}

func testAccCheckIntegrationsAuthConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_integrations_auth_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{IntegrationsBasePath}}{{name}}")
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
				return fmt.Errorf("IntegrationsAuthConfig still exists at %s", url)
			}
		}

		return nil
	}
}
