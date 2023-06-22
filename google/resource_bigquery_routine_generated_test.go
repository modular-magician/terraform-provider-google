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

func TestAccBigQueryRoutine_bigQueryRoutineBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBigQueryRoutineDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryRoutine_bigQueryRoutineBasicExample(context),
			},
			{
				ResourceName:      "google_bigquery_routine.sproc",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBigQueryRoutine_bigQueryRoutineBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
	dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_routine" "sproc" {
  dataset_id = google_bigquery_dataset.test.dataset_id
  routine_id     = "tf_test_routine_id%{random_suffix}"
  routine_type = "PROCEDURE"
  language = "SQL"
  definition_body = "CREATE FUNCTION Add(x FLOAT64, y FLOAT64) RETURNS FLOAT64 AS (x + y);"
}
`, context)
}

func TestAccBigQueryRoutine_bigQueryRoutineJsonExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBigQueryRoutineDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryRoutine_bigQueryRoutineJsonExample(context),
			},
			{
				ResourceName:      "google_bigquery_routine.sproc",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBigQueryRoutine_bigQueryRoutineJsonExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
	dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_routine" "sproc" {
  dataset_id = google_bigquery_dataset.test.dataset_id
  routine_id     = "tf_test_routine_id%{random_suffix}"
  routine_type = "SCALAR_FUNCTION"
  language = "JAVASCRIPT"
  definition_body = "CREATE FUNCTION multiplyInputs return x*y;"
  arguments {
    name = "x"
    data_type = "{\"typeKind\" :  \"FLOAT64\"}"
  } 
  arguments {
    name = "y"
    data_type = "{\"typeKind\" :  \"FLOAT64\"}"
  }
   
  return_type = "{\"typeKind\" :  \"FLOAT64\"}"
}
`, context)
}

func TestAccBigQueryRoutine_bigQueryRoutineTvfExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckBigQueryRoutineDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryRoutine_bigQueryRoutineTvfExample(context),
			},
			{
				ResourceName:      "google_bigquery_routine.sproc",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBigQueryRoutine_bigQueryRoutineTvfExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
	dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_routine" "sproc" {
  dataset_id      = google_bigquery_dataset.test.dataset_id
  routine_id      = "tf_test_routine_id%{random_suffix}"
  routine_type    = "TABLE_VALUED_FUNCTION"
  language        = "SQL"
  definition_body = <<-EOS
    SELECT 1 + value AS value
  EOS
  arguments {
    name          = "value"
    argument_kind = "FIXED_TYPE"
    data_type     = jsonencode({ "typeKind" : "INT64" })
  }
  return_table_type = jsonencode({"columns" : [
    { "name" : "value", "type" : { "typeKind" : "INT64" } },
  ] })
}
`, context)
}

func testAccCheckBigQueryRoutineDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_bigquery_routine" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}/routines/{{routine_id}}")
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
				return fmt.Errorf("BigQueryRoutine still exists at %s", url)
			}
		}

		return nil
	}
}
