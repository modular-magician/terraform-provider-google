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

package datalossprevention_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigOrgFolderPausedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"organization":  envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigOrgFolderPausedExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.org_folder_paused",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "parent"},
			},
		},
	})
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigOrgFolderPausedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_discovery_config" "org_folder_paused" {
	parent = "organizations/%{organization}/locations/us"
    location = "us"

    targets {
        big_query_target {
            filter {
                other_tables {}
            }
        }
    }
    org_config {
        project_id = "%{project}"
        location {
            folder_id = 123
        }
    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"]
    status = "PAUSED"
}

resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "My description"
	display_name = "display_name"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
    }
}
`, context)
}

func testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_data_loss_prevention_discovery_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DataLossPreventionBasePath}}{{parent}}/discoveryConfigs/{{name}}")
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
				return fmt.Errorf("DataLossPreventionDiscoveryConfig still exists at %s", url)
			}
		}

		return nil
	}
}
