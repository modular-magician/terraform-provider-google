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

package dns_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccDNSPolicy_dnsPolicyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDNSPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSPolicy_dnsPolicyBasicExample(context),
			},
			{
				ResourceName:      "google_dns_policy.example-policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDNSPolicy_dnsPolicyBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dns_policy" "example-policy" {
  name                      = "tf-test-example-policy%{random_suffix}"
  enable_inbound_forwarding = true

  enable_logging = true

  alternative_name_server_config {
    target_name_servers {
      ipv4_address    = "172.16.1.10"
      forwarding_path = "private"
    }
    target_name_servers {
      ipv4_address = "172.16.1.20"
    }
  }

  networks {
    network_url = google_compute_network.network-1.id
  }
  networks {
    network_url = google_compute_network.network-2.id
  }
}

resource "google_compute_network" "network-1" {
  name                    = "tf-test-network-1%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_network" "network-2" {
  name                    = "tf-test-network-2%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccDNSPolicy_dnsPolicyMultiprojectExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDNSPolicyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDNSPolicy_dnsPolicyMultiprojectExample(context),
			},
			{
				ResourceName:      "google_dns_policy.example-policy-multiproject",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDNSPolicy_dnsPolicyMultiprojectExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dns_policy" "example-policy-multiproject" {
  name                      = "tf-test-example-policy-multiproject%{random_suffix}"
  enable_inbound_forwarding = true

  enable_logging = true

  networks {
    network_url = google_compute_network.network_1_project_1.id
  }
  networks {
    network_url = google_compute_network.network_2_project_1.id
  }
  networks {
    network_url = google_compute_network.network_1_project_2.id
  }
  networks {
    network_url = google_compute_network.network_2_project_2.id
  }
  
  depends_on = [
    google_project_service.compute_project_1,
    google_project_service.dns_project_1,
    google_project_service.compute_project_2,
    google_project_service.dns_project_2,
  ]
}

resource "google_project" "project_1" {
  name            = "tf-test-project-1%{random_suffix}"
  project_id      = "tf-test-project-1%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
  deletion_policy = "DELETE"
}

resource "google_project" "project_2" {
  name            = "tf-test-project-2%{random_suffix}"
  project_id      = "tf-test-project-2%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
  deletion_policy = "DELETE"
}

resource "google_compute_network" "network_1_project_1" {
  name                    = "tf-test-network-1%{random_suffix}"
  project                 = google_project.project_1.project_id
  auto_create_subnetworks = false
  depends_on              = [ 
    google_project_service.compute_project_1,
    google_project_service.dns_project_1,
  ]
}

resource "google_compute_network" "network_2_project_1" {
  name                    = "tf-test-network-2%{random_suffix}"
  project                 = google_project.project_1.project_id
  auto_create_subnetworks = false
  depends_on              = [ 
    google_project_service.compute_project_1,
    google_project_service.dns_project_1,
  ]
}

resource "google_compute_network" "network_1_project_2" {
  name                    = "tf-test-network-1%{random_suffix}"
  project                 = google_project.project_2.project_id
  auto_create_subnetworks = false
  depends_on              = [ 
    google_project_service.compute_project_2,
    google_project_service.dns_project_2,
  ]
}

resource "google_compute_network" "network_2_project_2" {
  name                    = "tf-test-network-2%{random_suffix}"
  project                 = google_project.project_2.project_id
  auto_create_subnetworks = false
  depends_on              = [ 
    google_project_service.compute_project_2,
    google_project_service.dns_project_2,
  ]
}

resource "google_project_service" "compute_project_1" {
  project    = google_project.project_1.project_id
  service    = "compute.googleapis.com"
  depends_on = [
    google_project.project_1,
  ]
}

resource "google_project_service" "compute_project_2" {
  project    = google_project.project_2.project_id
  service    = "compute.googleapis.com"
  depends_on = [
    google_project_service.dns_project_1
  ]
}

resource "google_project_service" "dns_project_1" {
  project    = google_project.project_1.project_id
  service    = "dns.googleapis.com"
  depends_on = [
    google_project_service.compute_project_1
  ]
}

resource "google_project_service" "dns_project_2" {
  project    = google_project.project_2.project_id
  service    = "dns.googleapis.com"
  depends_on = [
    google_project_service.compute_project_2,
  ]
}
`, context)
}

func testAccCheckDNSPolicyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dns_policy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DNSBasePath}}projects/{{project}}/policies/{{name}}")
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
				return fmt.Errorf("DNSPolicy still exists at %s", url)
			}
		}

		return nil
	}
}
