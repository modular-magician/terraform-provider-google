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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIapTunnelInstanceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.tunnelResourceAccessor",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapTunnelInstanceIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_tunnel_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_tunnel/zones/%s/instances/%s roles/iap.tunnelResourceAccessor", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-tunnel%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccIapTunnelInstanceIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_iap_tunnel_instance_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_tunnel/zones/%s/instances/%s roles/iap.tunnelResourceAccessor", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-tunnel%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapTunnelInstanceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.tunnelResourceAccessor",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccIapTunnelInstanceIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_tunnel_instance_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_tunnel/zones/%s/instances/%s roles/iap.tunnelResourceAccessor user:admin@hashicorptest.com", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-tunnel%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapTunnelInstanceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.tunnelResourceAccessor",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapTunnelInstanceIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_tunnel_instance_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_tunnel/zones/%s/instances/%s", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-tunnel%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIapTunnelInstanceIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_iap_tunnel_instance_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_tunnel/zones/%s/instances/%s", getTestProjectFromEnv(), getTestZoneFromEnv(), fmt.Sprintf("tf-test-tunnel%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIapTunnelInstanceIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "tunnelvm" {
  name         = "tf-test-tunnel%{random_suffix}"
  zone         = ""
  machine_type = "n1-standard-1"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_iap_tunnel_instance_iam_member" "foo" {
  project = google_compute_instance.tunnelvm.project
  zone = google_compute_instance.tunnelvm.zone
  instance = google_compute_instance.tunnelvm.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccIapTunnelInstanceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "tunnelvm" {
  name         = "tf-test-tunnel%{random_suffix}"
  zone         = ""
  machine_type = "n1-standard-1"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_iap_tunnel_instance_iam_policy" "foo" {
  project = google_compute_instance.tunnelvm.project
  zone = google_compute_instance.tunnelvm.zone
  instance = google_compute_instance.tunnelvm.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccIapTunnelInstanceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "tunnelvm" {
  name         = "tf-test-tunnel%{random_suffix}"
  zone         = ""
  machine_type = "n1-standard-1"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

data "google_iam_policy" "foo" {
}

resource "google_iap_tunnel_instance_iam_policy" "foo" {
  project = google_compute_instance.tunnelvm.project
  zone = google_compute_instance.tunnelvm.zone
  instance = google_compute_instance.tunnelvm.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccIapTunnelInstanceIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "tunnelvm" {
  name         = "tf-test-tunnel%{random_suffix}"
  zone         = ""
  machine_type = "n1-standard-1"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_iap_tunnel_instance_iam_binding" "foo" {
  project = google_compute_instance.tunnelvm.project
  zone = google_compute_instance.tunnelvm.zone
  instance = google_compute_instance.tunnelvm.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccIapTunnelInstanceIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_instance" "tunnelvm" {
  name         = "tf-test-tunnel%{random_suffix}"
  zone         = ""
  machine_type = "n1-standard-1"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_iap_tunnel_instance_iam_binding" "foo" {
  project = google_compute_instance.tunnelvm.project
  zone = google_compute_instance.tunnelvm.zone
  instance = google_compute_instance.tunnelvm.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
