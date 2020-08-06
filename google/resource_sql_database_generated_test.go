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

func TestAccSQLDatabase_sqlDatabaseBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSQLDatabaseDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSQLDatabase_sqlDatabaseBasicExample(context),
			},
			{
				ResourceName:      "google_sql_database.database",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSQLDatabase_sqlDatabaseBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database" "database" {
  name     = "tf-test-my-database%{random_suffix}"
  instance = google_sql_database_instance.instance.name
}

resource "google_sql_database_instance" "instance" {
  name   = "tf-test-my-database-instance%{random_suffix}"
  region = "us-central1"
  settings {
    tier = "db-f1-micro"
  }
}
`, context)
}

func testAccCheckSQLDatabaseDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_sql_database" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{SQLBasePath}}projects/{{project}}/instances/{{instance}}/databases/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("SQLDatabase still exists at %s", url)
			}
		}

		return nil
	}
}
