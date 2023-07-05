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

func TestAccNotebooksInstanceIamBindingGenerated(t *testing.T) {
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
				Config: testAccNotebooksInstanceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_notebooks_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s roles/viewer", envvar.GetTestProjectFromEnv(), "us-west1-a", fmt.Sprintf("tf-test-notebooks-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccNotebooksInstanceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_notebooks_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s roles/viewer", envvar.GetTestProjectFromEnv(), "us-west1-a", fmt.Sprintf("tf-test-notebooks-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNotebooksInstanceIamMemberGenerated(t *testing.T) {
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
				Config: testAccNotebooksInstanceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_notebooks_instance_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), "us-west1-a", fmt.Sprintf("tf-test-notebooks-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNotebooksInstanceIamPolicyGenerated(t *testing.T) {
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
				Config: testAccNotebooksInstanceIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_notebooks_instance_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_notebooks_instance_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s", envvar.GetTestProjectFromEnv(), "us-west1-a", fmt.Sprintf("tf-test-notebooks-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNotebooksInstanceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_notebooks_instance_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/instances/%s", envvar.GetTestProjectFromEnv(), "us-west1-a", fmt.Sprintf("tf-test-notebooks-instance%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNotebooksInstanceIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_notebooks_instance" "instance" {
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-west1-a"
  machine_type = "e2-medium"
  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }
}

resource "google_notebooks_instance_iam_member" "foo" {
  project = google_notebooks_instance.instance.project
  location = google_notebooks_instance.instance.location
  instance_name = google_notebooks_instance.instance.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccNotebooksInstanceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_notebooks_instance" "instance" {
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-west1-a"
  machine_type = "e2-medium"
  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_notebooks_instance_iam_policy" "foo" {
  project = google_notebooks_instance.instance.project
  location = google_notebooks_instance.instance.location
  instance_name = google_notebooks_instance.instance.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_notebooks_instance_iam_policy" "foo" {
  project = google_notebooks_instance.instance.project
  location = google_notebooks_instance.instance.location
  instance_name = google_notebooks_instance.instance.name
  depends_on = [
    google_notebooks_instance_iam_policy.foo
  ]
}
`, context)
}

func testAccNotebooksInstanceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_notebooks_instance" "instance" {
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-west1-a"
  machine_type = "e2-medium"
  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }
}

data "google_iam_policy" "foo" {
}

resource "google_notebooks_instance_iam_policy" "foo" {
  project = google_notebooks_instance.instance.project
  location = google_notebooks_instance.instance.location
  instance_name = google_notebooks_instance.instance.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccNotebooksInstanceIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_notebooks_instance" "instance" {
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-west1-a"
  machine_type = "e2-medium"
  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }
}

resource "google_notebooks_instance_iam_binding" "foo" {
  project = google_notebooks_instance.instance.project
  location = google_notebooks_instance.instance.location
  instance_name = google_notebooks_instance.instance.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccNotebooksInstanceIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_notebooks_instance" "instance" {
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-west1-a"
  machine_type = "e2-medium"
  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }
}

resource "google_notebooks_instance_iam_binding" "foo" {
  project = google_notebooks_instance.instance.project
  location = google_notebooks_instance.instance.location
  instance_name = google_notebooks_instance.instance.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
