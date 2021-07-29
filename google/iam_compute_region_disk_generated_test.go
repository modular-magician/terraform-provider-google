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
)

func TestAccComputeRegionDiskIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionDiskIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_region_disk_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/disks/%s roles/viewer", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-my-region-disk%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccComputeRegionDiskIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_compute_region_disk_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/disks/%s roles/viewer", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-my-region-disk%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionDiskIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccComputeRegionDiskIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_region_disk_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/disks/%s roles/viewer user:admin@hashicorptest.com", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-my-region-disk%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccComputeRegionDiskIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionDiskIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_compute_region_disk_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/disks/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-my-region-disk%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccComputeRegionDiskIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_compute_region_disk_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/regions/%s/disks/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-my-region-disk%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionDiskIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_disk" "regiondisk" {
  name                      = "tf-test-my-region-disk%{random_suffix}"
  snapshot                  = google_compute_snapshot.snapdisk.id
  type                      = "pd-ssd"
  region                    = "us-central1"
  physical_block_size_bytes = 4096

  replica_zones = ["us-central1-a", "us-central1-f"]
}

resource "google_compute_disk" "disk" {
  name  = "tf-test-my-disk%{random_suffix}"
  image = "debian-cloud/debian-9"
  size  = 50
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_snapshot" "snapdisk" {
  name        = "tf-test-my-snapshot%{random_suffix}"
  source_disk = google_compute_disk.disk.name
  zone        = "us-central1-a"
}

resource "google_compute_region_disk_iam_member" "foo" {
  project = google_compute_region_disk.regiondisk.project
  region = google_compute_region_disk.regiondisk.region
  name = google_compute_region_disk.regiondisk.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccComputeRegionDiskIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_disk" "regiondisk" {
  name                      = "tf-test-my-region-disk%{random_suffix}"
  snapshot                  = google_compute_snapshot.snapdisk.id
  type                      = "pd-ssd"
  region                    = "us-central1"
  physical_block_size_bytes = 4096

  replica_zones = ["us-central1-a", "us-central1-f"]
}

resource "google_compute_disk" "disk" {
  name  = "tf-test-my-disk%{random_suffix}"
  image = "debian-cloud/debian-9"
  size  = 50
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_snapshot" "snapdisk" {
  name        = "tf-test-my-snapshot%{random_suffix}"
  source_disk = google_compute_disk.disk.name
  zone        = "us-central1-a"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_compute_region_disk_iam_policy" "foo" {
  project = google_compute_region_disk.regiondisk.project
  region = google_compute_region_disk.regiondisk.region
  name = google_compute_region_disk.regiondisk.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeRegionDiskIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_disk" "regiondisk" {
  name                      = "tf-test-my-region-disk%{random_suffix}"
  snapshot                  = google_compute_snapshot.snapdisk.id
  type                      = "pd-ssd"
  region                    = "us-central1"
  physical_block_size_bytes = 4096

  replica_zones = ["us-central1-a", "us-central1-f"]
}

resource "google_compute_disk" "disk" {
  name  = "tf-test-my-disk%{random_suffix}"
  image = "debian-cloud/debian-9"
  size  = 50
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_snapshot" "snapdisk" {
  name        = "tf-test-my-snapshot%{random_suffix}"
  source_disk = google_compute_disk.disk.name
  zone        = "us-central1-a"
}

data "google_iam_policy" "foo" {
}

resource "google_compute_region_disk_iam_policy" "foo" {
  project = google_compute_region_disk.regiondisk.project
  region = google_compute_region_disk.regiondisk.region
  name = google_compute_region_disk.regiondisk.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccComputeRegionDiskIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_disk" "regiondisk" {
  name                      = "tf-test-my-region-disk%{random_suffix}"
  snapshot                  = google_compute_snapshot.snapdisk.id
  type                      = "pd-ssd"
  region                    = "us-central1"
  physical_block_size_bytes = 4096

  replica_zones = ["us-central1-a", "us-central1-f"]
}

resource "google_compute_disk" "disk" {
  name  = "tf-test-my-disk%{random_suffix}"
  image = "debian-cloud/debian-9"
  size  = 50
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_snapshot" "snapdisk" {
  name        = "tf-test-my-snapshot%{random_suffix}"
  source_disk = google_compute_disk.disk.name
  zone        = "us-central1-a"
}

resource "google_compute_region_disk_iam_binding" "foo" {
  project = google_compute_region_disk.regiondisk.project
  region = google_compute_region_disk.regiondisk.region
  name = google_compute_region_disk.regiondisk.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccComputeRegionDiskIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_disk" "regiondisk" {
  name                      = "tf-test-my-region-disk%{random_suffix}"
  snapshot                  = google_compute_snapshot.snapdisk.id
  type                      = "pd-ssd"
  region                    = "us-central1"
  physical_block_size_bytes = 4096

  replica_zones = ["us-central1-a", "us-central1-f"]
}

resource "google_compute_disk" "disk" {
  name  = "tf-test-my-disk%{random_suffix}"
  image = "debian-cloud/debian-9"
  size  = 50
  type  = "pd-ssd"
  zone  = "us-central1-a"
}

resource "google_compute_snapshot" "snapdisk" {
  name        = "tf-test-my-snapshot%{random_suffix}"
  source_disk = google_compute_disk.disk.name
  zone        = "us-central1-a"
}

resource "google_compute_region_disk_iam_binding" "foo" {
  project = google_compute_region_disk.regiondisk.project
  region = google_compute_region_disk.regiondisk.region
  name = google_compute_region_disk.regiondisk.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
