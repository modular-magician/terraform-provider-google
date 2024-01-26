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

package netapp_test

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

func TestAccNetappVolume_volumeBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"network_name":  acctest.BootstrapSharedServiceNetworkingConnection(t, "gcnv-network-config-1", acctest.ServiceNetworkWithParentService("netapp.servicenetworking.goog")),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetappVolumeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetappVolume_volumeBasicExample(context),
			},
			{
				ResourceName:            "google_netapp_volume.test_volume",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "name", "deletion_policy", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetappVolume_volumeBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_netapp_storage_pool" "default" {
  name = "tf-test-test-pool%{random_suffix}"
  location = "us-west2"
  service_level = "PREMIUM"
  capacity_gib = "2048"
  network = data.google_compute_network.default.id
}

resource "google_netapp_volume" "test_volume" {
  location = "us-west2"
  name = "tf-test-test-volume%{random_suffix}"
  capacity_gib = "100"
  share_name = "tf-test-test-volume%{random_suffix}"
  storage_pool = google_netapp_storage_pool.default.name
  protocols = ["NFSV3"]
  deletion_policy = "DEFAULT"
}

data "google_compute_network" "default" {
  name = "%{network_name}"
}
`, context)
}

func testAccCheckNetappVolumeDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_netapp_volume" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetappBasePath}}projects/{{project}}/locations/{{location}}/volumes/{{name}}")
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
				return fmt.Errorf("NetappVolume still exists at %s", url)
			}
		}

		return nil
	}
}
