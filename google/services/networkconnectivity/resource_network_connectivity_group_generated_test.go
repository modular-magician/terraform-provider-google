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

package networkconnectivity_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccNetworkConnectivityGroup_networkConnectivityGroupBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkConnectivityGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkConnectivityGroup_networkConnectivityGroupBasicExample(context),
			},
			{
				ResourceName:            "google_network_connectivity_group.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"hub", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkConnectivityGroup_networkConnectivityGroupBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_connectivity_hub" "basic_hub"  {
 name        = "hub1%{random_suffix}"
 description = "A sample hub"
 labels = {
    label-one = "value-one"
  }
}

resource "google_network_connectivity_group" "primary"  {
 hub         = google_network_connectivity_hub.basic_hub.id
 name        = "default"
 auto_accept {
    auto_accept_projects = [
      "foo%{random_suffix}", 
      "bar%{random_suffix}", 
    ]
  }
}
`, context)
}

func testAccCheckNetworkConnectivityGroupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_connectivity_group" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkConnectivityBasePath}}projects/{{project}}/locations/global/hubs/{{hub}}/groups/{{name}}")
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
				return fmt.Errorf("NetworkConnectivityGroup still exists at %s", url)
			}
		}

		return nil
	}
}
