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

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccBigqueryConnectionConnectionIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryConnectionConnectionIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_connection_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/connections/%s roles/viewer", envvar.GetTestProjectFromEnv(), "US", fmt.Sprintf("tf-test-my-connection%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccBigqueryConnectionConnectionIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_connection_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/connections/%s roles/viewer", envvar.GetTestProjectFromEnv(), "US", fmt.Sprintf("tf-test-my-connection%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigqueryConnectionConnectionIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccBigqueryConnectionConnectionIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_connection_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/connections/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), "US", fmt.Sprintf("tf-test-my-connection%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigqueryConnectionConnectionIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryConnectionConnectionIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_bigquery_connection_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_bigquery_connection_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/connections/%s", envvar.GetTestProjectFromEnv(), "US", fmt.Sprintf("tf-test-my-connection%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBigqueryConnectionConnectionIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_bigquery_connection_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/connections/%s", envvar.GetTestProjectFromEnv(), "US", fmt.Sprintf("tf-test-my-connection%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBigqueryConnectionConnectionIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_connection" "connection" {
   connection_id = "tf-test-my-connection%{random_suffix}"
   location      = "US"
   friendly_name = "👋"
   description   = "a riveting description"
   cloud_resource {}
}

resource "google_bigquery_connection_iam_member" "foo" {
  project = google_bigquery_connection.connection.project
  location = google_bigquery_connection.connection.location
  connection_id = google_bigquery_connection.connection.connection_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccBigqueryConnectionConnectionIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_connection" "connection" {
   connection_id = "tf-test-my-connection%{random_suffix}"
   location      = "US"
   friendly_name = "👋"
   description   = "a riveting description"
   cloud_resource {}
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_bigquery_connection_iam_policy" "foo" {
  project = google_bigquery_connection.connection.project
  location = google_bigquery_connection.connection.location
  connection_id = google_bigquery_connection.connection.connection_id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_bigquery_connection_iam_policy" "foo" {
  project = google_bigquery_connection.connection.project
  location = google_bigquery_connection.connection.location
  connection_id = google_bigquery_connection.connection.connection_id
  depends_on = [
    google_bigquery_connection_iam_policy.foo
  ]
}
`, context)
}

func testAccBigqueryConnectionConnectionIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_connection" "connection" {
   connection_id = "tf-test-my-connection%{random_suffix}"
   location      = "US"
   friendly_name = "👋"
   description   = "a riveting description"
   cloud_resource {}
}

data "google_iam_policy" "foo" {
}

resource "google_bigquery_connection_iam_policy" "foo" {
  project = google_bigquery_connection.connection.project
  location = google_bigquery_connection.connection.location
  connection_id = google_bigquery_connection.connection.connection_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccBigqueryConnectionConnectionIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_connection" "connection" {
   connection_id = "tf-test-my-connection%{random_suffix}"
   location      = "US"
   friendly_name = "👋"
   description   = "a riveting description"
   cloud_resource {}
}

resource "google_bigquery_connection_iam_binding" "foo" {
  project = google_bigquery_connection.connection.project
  location = google_bigquery_connection.connection.location
  connection_id = google_bigquery_connection.connection.connection_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccBigqueryConnectionConnectionIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_bigquery_connection" "connection" {
   connection_id = "tf-test-my-connection%{random_suffix}"
   location      = "US"
   friendly_name = "👋"
   description   = "a riveting description"
   cloud_resource {}
}

resource "google_bigquery_connection_iam_binding" "foo" {
  project = google_bigquery_connection.connection.project
  location = google_bigquery_connection.connection.location
  connection_id = google_bigquery_connection.connection.connection_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
