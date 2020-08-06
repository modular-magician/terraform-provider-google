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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataCatalogEntryGroup_dataCatalogEntryGroupBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDataCatalogEntryGroupDestroyProducer(t),
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
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDataCatalogEntryGroupDestroyProducer(t),
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

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{DataCatalogBasePath}}{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("DataCatalogEntryGroup still exists at %s", url)
			}
		}

		return nil
	}
}
