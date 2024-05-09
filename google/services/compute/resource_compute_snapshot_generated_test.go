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

package compute_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccComputeSnapshot_snapshotBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeSnapshotDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSnapshot_snapshotBasicExample(context),
			},
			{
				ResourceName:            "google_compute_snapshot.snapshot",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "source_disk", "source_disk_encryption_key", "terraform_labels", "zone"},
			},
		},
	})
}

func testAccComputeSnapshot_snapshotBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_snapshot" "snapshot" {
  name        = "tf-test-my-snapshot%{random_suffix}"
  source_disk = google_compute_disk.persistent.id
  zone        = "us-central1-a"
  labels = {
    my_label = "value"
  }
  storage_locations = ["us-central1"]
}

data "google_compute_image" "debian" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_disk" "persistent" {
  name  = "tf-test-debian-disk%{random_suffix}"
  image = data.google_compute_image.debian.self_link
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-a"
}
`, context)
}

func TestAccComputeSnapshot_snapshotChainnameExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeSnapshotDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSnapshot_snapshotChainnameExample(context),
			},
			{
				ResourceName:            "google_compute_snapshot.snapshot",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "source_disk", "source_disk_encryption_key", "terraform_labels", "zone"},
			},
		},
	})
}

func testAccComputeSnapshot_snapshotChainnameExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_snapshot" "snapshot" {
  name        = "tf-test-my-snapshot%{random_suffix}"
  source_disk = google_compute_disk.persistent.id
  zone        = "us-central1-a"
  chain_name  = "tf-test-snapshot-chain%{random_suffix}"
  labels = {
    my_label = "value"
  }
  storage_locations = ["us-central1"]
}

data "google_compute_image" "debian" {
  family  = "debian-11"
  project = "debian-cloud"
}

resource "google_compute_disk" "persistent" {
  name  = "tf-test-debian-disk%{random_suffix}"
  image = data.google_compute_image.debian.self_link
  size  = 10
  type  = "pd-ssd"
  zone  = "us-central1-a"
}
`, context)
}

func testAccCheckComputeSnapshotDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_snapshot" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/snapshots/{{name}}")
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
				return fmt.Errorf("ComputeSnapshot still exists at %s", url)
			}
		}

		return nil
	}
}
