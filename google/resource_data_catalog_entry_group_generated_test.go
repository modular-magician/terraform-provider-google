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

func TestAccDataCatalogEntryGroup_dataCatalogEntryGroupBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataCatalogEntryGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataCatalogEntryGroup_dataCatalogEntryGroupBasicExample(context),
			},
			{
				ResourceName:            "google_data_catalog_entry_group.basic_entry_group",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "entry_group_id"},
			},
		},
	})
}

func testAccDataCatalogEntryGroup_dataCatalogEntryGroupBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_entry_group" "basic_entry_group" {
  entry_group_id = "tf_test_my_group%{random_suffix}"
}
`, context)
}

func TestAccDataCatalogEntryGroup_dataCatalogEntryGroupFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataCatalogEntryGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataCatalogEntryGroup_dataCatalogEntryGroupFullExample(context),
			},
			{
				ResourceName:            "google_data_catalog_entry_group.basic_entry_group",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "entry_group_id"},
			},
		},
	})
}

func testAccDataCatalogEntryGroup_dataCatalogEntryGroupFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_entry_group" "basic_entry_group" {
  entry_group_id = "tf_test_my_group%{random_suffix}"

  display_name = "terraform entry group"
  description = "entry group created by Terraform"
}
`, context)
}

func testAccCheckDataCatalogEntryGroupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_data_catalog_entry_group" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DataCatalogBasePath}}{{name}}")
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
				return fmt.Errorf("DataCatalogEntryGroup still exists at %s", url)
			}
		}

		return nil
	}
}
