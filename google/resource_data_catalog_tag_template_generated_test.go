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

	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataCatalogTagTemplate_dataCatalogTagTemplateBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"force_delete":  true,
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDataCatalogTagTemplateDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataCatalogTagTemplate_dataCatalogTagTemplateBasicExample(context),
			},
			{
				ResourceName:            "google_data_catalog_tag_template.basic_tag_template",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region", "tag_template_id", "force_delete"},
			},
		},
	})
}

func testAccDataCatalogTagTemplate_dataCatalogTagTemplateBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_catalog_tag_template" "basic_tag_template" {
  tag_template_id = "tf_test_my_template%{random_suffix}"
  region = "us-central1"
  display_name = "Demo Tag Template"

  fields {
    field_id = "source"
    display_name = "Source of data asset"
    type {
      primitive_type = "STRING"
    }
    is_required = true
  }

  fields {
    field_id = "num_rows"
    display_name = "Number of rows in the data asset"
    type {
      primitive_type = "DOUBLE"
    }
  }

  fields {
    field_id = "pii_type"
    display_name = "PII type"
    type {
      enum_type {
        allowed_values {
          display_name = "EMAIL"
        }
        allowed_values {
          display_name = "SOCIAL SECURITY NUMBER"
        }
        allowed_values {
          display_name = "NONE"
        }
      }
    }
  }

  force_delete = "%{force_delete}"
}
`, context)
}

func testAccCheckDataCatalogTagTemplateDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_data_catalog_tag_template" {
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
				return fmt.Errorf("DataCatalogTagTemplate still exists at %s", url)
			}
		}

		return nil
	}
}
