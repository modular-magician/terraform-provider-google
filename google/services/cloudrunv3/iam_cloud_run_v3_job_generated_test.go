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

package cloudrunv3_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccCloudRunV3JobIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunV3JobIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloud_run_v3_job_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/jobs/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudrun-job%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccCloudRunV3JobIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_cloud_run_v3_job_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/jobs/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudrun-job%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudRunV3JobIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccCloudRunV3JobIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloud_run_v3_job_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/jobs/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudrun-job%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudRunV3JobIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunV3JobIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_cloud_run_v3_job_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_cloud_run_v3_job_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/jobs/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudrun-job%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCloudRunV3JobIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_cloud_run_v3_job_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/jobs/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudrun-job%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudRunV3JobIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_run_v2_job" "default" {
  name     = "tf-test-cloudrun-job%{random_suffix}"
  location = "us-central1"

  template {
    template {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }
}

resource "google_cloud_run_v3_job_iam_member" "foo" {
  project = google_cloud_run_v3_job.default.project
  location = google_cloud_run_v3_job.default.location
  name = google_cloud_run_v3_job.default.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccCloudRunV3JobIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_run_v2_job" "default" {
  name     = "tf-test-cloudrun-job%{random_suffix}"
  location = "us-central1"

  template {
    template {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_cloud_run_v3_job_iam_policy" "foo" {
  project = google_cloud_run_v3_job.default.project
  location = google_cloud_run_v3_job.default.location
  name = google_cloud_run_v3_job.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_cloud_run_v3_job_iam_policy" "foo" {
  project = google_cloud_run_v3_job.default.project
  location = google_cloud_run_v3_job.default.location
  name = google_cloud_run_v3_job.default.name
  depends_on = [
    google_cloud_run_v3_job_iam_policy.foo
  ]
}
`, context)
}

func testAccCloudRunV3JobIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_run_v2_job" "default" {
  name     = "tf-test-cloudrun-job%{random_suffix}"
  location = "us-central1"

  template {
    template {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }
}

data "google_iam_policy" "foo" {
}

resource "google_cloud_run_v3_job_iam_policy" "foo" {
  project = google_cloud_run_v3_job.default.project
  location = google_cloud_run_v3_job.default.location
  name = google_cloud_run_v3_job.default.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccCloudRunV3JobIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_run_v2_job" "default" {
  name     = "tf-test-cloudrun-job%{random_suffix}"
  location = "us-central1"

  template {
    template {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }
}

resource "google_cloud_run_v3_job_iam_binding" "foo" {
  project = google_cloud_run_v3_job.default.project
  location = google_cloud_run_v3_job.default.location
  name = google_cloud_run_v3_job.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccCloudRunV3JobIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_cloud_run_v2_job" "default" {
  name     = "tf-test-cloudrun-job%{random_suffix}"
  location = "us-central1"

  template {
    template {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }
}

resource "google_cloud_run_v3_job_iam_binding" "foo" {
  project = google_cloud_run_v3_job.default.project
  location = google_cloud_run_v3_job.default.location
  name = google_cloud_run_v3_job.default.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
