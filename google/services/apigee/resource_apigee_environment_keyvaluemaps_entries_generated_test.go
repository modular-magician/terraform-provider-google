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

package apigee_test

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

func TestAccApigeeEnvironmentKeyvaluemapsEntries_apigeeEnvironmentKeyvaluemapsEntriesTestExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckApigeeEnvironmentKeyvaluemapsEntriesDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccApigeeEnvironmentKeyvaluemapsEntries_apigeeEnvironmentKeyvaluemapsEntriesTestExample(context),
			},
			{
				ResourceName:            "google_apigee_environment_keyvaluemaps_entries.apigee_environment_keyvaluemaps_entries",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"env_keyvaluemap_id"},
			},
		},
	})
}

func testAccApigeeEnvironmentKeyvaluemapsEntries_apigeeEnvironmentKeyvaluemapsEntriesTestExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "project" {
  project_id      = "tf-test-%{random_suffix}"
  name            = "tf-test-%{random_suffix}"
  org_id          = "%{org_id}"
  billing_account = "%{billing_account}"
}

resource "google_project_service" "apigee" {
  project = google_project.project.project_id
  service = "apigee.googleapis.com"
}

resource "google_project_service" "compute" {
  project = google_project.project.project_id
  service = "compute.googleapis.com"
  depends_on = [ google_project_service.servicenetworking ]
}

resource "google_project_service" "servicenetworking" {
  project = google_project.project.project_id
  service = "servicenetworking.googleapis.com"
  depends_on = [ google_project_service.apigee ]
}

resource "google_compute_network" "apigee_network" {
  name       = "apigee-network"
  project    = google_project.project.project_id
  depends_on = [ google_project_service.compute ]
}

resource "google_compute_global_address" "apigee_range" {
  name          = "apigee-range"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.apigee_network.id
  project       = google_project.project.project_id
}

resource "google_service_networking_connection" "apigee_vpc_connection" {
  network                 = google_compute_network.apigee_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.apigee_range.name]
}

resource "google_apigee_organization" "apigee_org" {
  analytics_region   = "us-central1"
  project_id         = google_project.project.project_id
  authorized_network = google_compute_network.apigee_network.id
  billing_type       = "EVALUATION"
  runtime_type       = "CLOUD"
  properties { 
    property { 
      name = "features.hybrid.enabled" 
      value = "true" 
    } 
    property { 
      name = "features.mart.connect.enabled" 
      value = "true" 
    } 
  } 
    
  depends_on = [
    google_service_networking_connection.apigee_vpc_connection,
    google_project_service.apigee
  ]
}

resource "google_apigee_environment" "apigee_environment" {
  org_id       = google_apigee_organization.apigee_org.id
  name         = "tf-test-env%{random_suffix}"
  description  = "Apigee Environment"
  display_name = "Apigee Environment"
}

resource "google_apigee_instance" "apigee_instance" {
  name     = "tf-test%{random_suffix}"
  location = "us-central1"
  org_id   = google_apigee_organization.apigee_org.id
}

resource "google_apigee_instance_attachment" "apigee_instance_attachment" {
  instance_id  = google_apigee_instance.apigee_instance.id
  environment  = google_apigee_environment.apigee_environment.name
}

resource "google_apigee_environment_keyvaluemaps" "apigee_environment_keyvaluemaps" {
  env_id    = google_apigee_environment.apigee_environment.id
  name      = "tf-test-env-kvms%{random_suffix}"
  depends_on = [
    google_apigee_organization.apigee_org,
    google_apigee_environment.apigee_environment,
    google_apigee_instance.apigee_instance,
    google_apigee_instance_attachment.apigee_instance_attachment
  ]
}

resource "google_apigee_environment_keyvaluemaps_entries" "apigee_environment_keyvaluemaps_entries" {
  
  env_keyvaluemap_id = google_apigee_environment_keyvaluemaps.apigee_environment_keyvaluemaps.id
  name           = "testName"
  value          = "testValue"
  depends_on = [
    google_apigee_organization.apigee_org,
    google_apigee_environment.apigee_environment,
    google_apigee_instance.apigee_instance,
    google_apigee_instance_attachment.apigee_instance_attachment,
    google_apigee_environment_keyvaluemaps.apigee_environment_keyvaluemaps
  ]
}
`, context)
}

func testAccCheckApigeeEnvironmentKeyvaluemapsEntriesDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_apigee_environment_keyvaluemaps_entries" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ApigeeBasePath}}{{env_keyvaluemap_id}}/entries/{{name}}")
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
				return fmt.Errorf("ApigeeEnvironmentKeyvaluemapsEntries still exists at %s", url)
			}
		}

		return nil
	}
}
