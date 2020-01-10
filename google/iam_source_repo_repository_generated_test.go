// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//     ***     DIFF TEST DIFF TEST    ***
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

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccSourceRepoRepositoryIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/viewer",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSourceRepoRepositoryIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_sourcerepo_repository_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/repos/%s roles/viewer", getTestProjectFromEnv(), fmt.Sprintf("my/repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccSourceRepoRepositoryIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_sourcerepo_repository_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/repos/%s roles/viewer", getTestProjectFromEnv(), fmt.Sprintf("my/repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSourceRepoRepositoryIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/viewer",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccSourceRepoRepositoryIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_sourcerepo_repository_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/repos/%s roles/viewer user:admin@hashicorptest.com", getTestProjectFromEnv(), fmt.Sprintf("my/repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSourceRepoRepositoryIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/viewer",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSourceRepoRepositoryIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_sourcerepo_repository_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/repos/%s", getTestProjectFromEnv(), fmt.Sprintf("my/repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSourceRepoRepositoryIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_sourcerepo_repository_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/repos/%s", getTestProjectFromEnv(), fmt.Sprintf("my/repository%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSourceRepoRepositoryIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sourcerepo_repository" "my-repo" {
  name = "my/repository%{random_suffix}"
}

resource "google_sourcerepo_repository_iam_member" "foo" {
  project = "${google_sourcerepo_repository.my-repo.project}"
  repository = "${google_sourcerepo_repository.my-repo.name}"
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccSourceRepoRepositoryIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sourcerepo_repository" "my-repo" {
  name = "my/repository%{random_suffix}"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_sourcerepo_repository_iam_policy" "foo" {
  project = "${google_sourcerepo_repository.my-repo.project}"
  repository = "${google_sourcerepo_repository.my-repo.name}"
  policy_data = "${data.google_iam_policy.foo.policy_data}"
}
`, context)
}

func testAccSourceRepoRepositoryIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sourcerepo_repository" "my-repo" {
  name = "my/repository%{random_suffix}"
}

data "google_iam_policy" "foo" {
}

resource "google_sourcerepo_repository_iam_policy" "foo" {
  project = "${google_sourcerepo_repository.my-repo.project}"
  repository = "${google_sourcerepo_repository.my-repo.name}"
  policy_data = "${data.google_iam_policy.foo.policy_data}"
}
`, context)
}

func testAccSourceRepoRepositoryIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sourcerepo_repository" "my-repo" {
  name = "my/repository%{random_suffix}"
}

resource "google_sourcerepo_repository_iam_binding" "foo" {
  project = "${google_sourcerepo_repository.my-repo.project}"
  repository = "${google_sourcerepo_repository.my-repo.name}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccSourceRepoRepositoryIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sourcerepo_repository" "my-repo" {
  name = "my/repository%{random_suffix}"
}

resource "google_sourcerepo_repository_iam_binding" "foo" {
  project = "${google_sourcerepo_repository.my-repo.project}"
  repository = "${google_sourcerepo_repository.my-repo.name}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
