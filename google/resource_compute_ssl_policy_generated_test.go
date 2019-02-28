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

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccComputeSslPolicy_sslPolicyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeSslPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSslPolicy_sslPolicyBasicExample(context),
			},
			{
				ResourceName:      "google_compute_ssl_policy.prod-ssl-policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeSslPolicy_sslPolicyBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_ssl_policy" "prod-ssl-policy" {
  name    = "production-ssl-policy-%{random_suffix}"
  profile = "MODERN"
}

resource "google_compute_ssl_policy" "nonprod-ssl-policy" {
  name            = "nonprod-ssl-policy-%{random_suffix}"
  profile         = "MODERN"
  min_tls_version = "TLS_1_2"
}

resource "google_compute_ssl_policy" "custom-ssl-policy" {
  name            = "custom-ssl-policy-%{random_suffix}"
  min_tls_version = "TLS_1_2"
  profile         = "CUSTOM"
  custom_features = ["TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"]
}
`, context)
}

func testAccCheckComputeSslPolicyDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_ssl_policy" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/sslPolicies/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeSslPolicy still exists at %s", url)
		}
	}

	return nil
}
