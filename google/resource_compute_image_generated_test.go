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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccComputeImage_imageBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImage_imageBasicExample(context),
			},
			{
				ResourceName:            "google_compute_image.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"raw_disk", "source_disk", "source_image", "source_snapshot"},
			},
		},
	})
}

func testAccComputeImage_imageBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_image" "example" {
  name = "tf-test-example-image%{random_suffix}"

  raw_disk {
    source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"
  }
}
`, context)
}

func TestAccComputeImage_imageGuestOsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImage_imageGuestOsExample(context),
			},
			{
				ResourceName:            "google_compute_image.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"raw_disk", "source_disk", "source_image", "source_snapshot"},
			},
		},
	})
}

func testAccComputeImage_imageGuestOsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_image" "example" {
  name = "tf-test-example-image%{random_suffix}"

  raw_disk {
    source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"
  }

  guest_os_features {
    type = "SECURE_BOOT"
  }

  guest_os_features {
    type = "MULTI_IP_SUBNET"
  }
}
`, context)
}

func TestAccComputeImage_imageBasicStorageLocationExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeImageDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeImage_imageBasicStorageLocationExample(context),
			},
			{
				ResourceName:            "google_compute_image.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"raw_disk", "source_disk", "source_image", "source_snapshot"},
			},
		},
	})
}

func testAccComputeImage_imageBasicStorageLocationExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_image" "example" {
  name = "tf-test-example-sl-image%{random_suffix}"

  raw_disk {
    source = "https://storage.googleapis.com/bosh-gce-raw-stemcells/bosh-stemcell-97.98-google-kvm-ubuntu-xenial-go_agent-raw-1557960142.tar.gz"
  }
  storage_locations = ["us-central1"]
}
`, context)
}

func testAccCheckComputeImageDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_image" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/images/{{name}}")
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
				return fmt.Errorf("ComputeImage still exists at %s", url)
			}
		}

		return nil
	}
}
