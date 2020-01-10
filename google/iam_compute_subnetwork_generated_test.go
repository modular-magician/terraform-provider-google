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

func TestAccComputeSubnetworkIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/compute.networkUser",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetworkIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s roles/compute.networkUser", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("test-subnetwork%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccComputeSubnetworkIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s roles/compute.networkUser", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("test-subnetwork%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeSubnetworkIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/compute.networkUser",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccComputeSubnetworkIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s roles/compute.networkUser user:admin@hashicorptest.com", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("test-subnetwork%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeSubnetworkIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/compute.networkUser",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetworkIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("test-subnetwork%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeSubnetworkIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_compute_subnetwork_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/subnetworks/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("test-subnetwork%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeSubnetworkIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.self_link
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "test-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork_iam_member" "foo" {
  project = "${google_compute_subnetwork.network-with-private-secondary-ip-ranges.project}"
  region = "${google_compute_subnetwork.network-with-private-secondary-ip-ranges.region}"
  subnetwork = "${google_compute_subnetwork.network-with-private-secondary-ip-ranges.name}"
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccComputeSubnetworkIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.self_link
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "test-network%{random_suffix}"
  auto_create_subnetworks = false
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_compute_subnetwork_iam_policy" "foo" {
  project = "${google_compute_subnetwork.network-with-private-secondary-ip-ranges.project}"
  region = "${google_compute_subnetwork.network-with-private-secondary-ip-ranges.region}"
  subnetwork = "${google_compute_subnetwork.network-with-private-secondary-ip-ranges.name}"
  policy_data = "${data.google_iam_policy.foo.policy_data}"
}
`, context)
}

func testAccComputeSubnetworkIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.self_link
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "test-network%{random_suffix}"
  auto_create_subnetworks = false
}

data "google_iam_policy" "foo" {
}

resource "google_compute_subnetwork_iam_policy" "foo" {
  project = "${google_compute_subnetwork.network-with-private-secondary-ip-ranges.project}"
  region = "${google_compute_subnetwork.network-with-private-secondary-ip-ranges.region}"
  subnetwork = "${google_compute_subnetwork.network-with-private-secondary-ip-ranges.name}"
  policy_data = "${data.google_iam_policy.foo.policy_data}"
}
`, context)
}

func testAccComputeSubnetworkIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.self_link
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "test-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork_iam_binding" "foo" {
  project = "${google_compute_subnetwork.network-with-private-secondary-ip-ranges.project}"
  region = "${google_compute_subnetwork.network-with-private-secondary-ip-ranges.region}"
  subnetwork = "${google_compute_subnetwork.network-with-private-secondary-ip-ranges.name}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccComputeSubnetworkIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "test-subnetwork%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.self_link
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "test-network%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork_iam_binding" "foo" {
  project = "${google_compute_subnetwork.network-with-private-secondary-ip-ranges.project}"
  region = "${google_compute_subnetwork.network-with-private-secondary-ip-ranges.region}"
  subnetwork = "${google_compute_subnetwork.network-with-private-secondary-ip-ranges.name}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
