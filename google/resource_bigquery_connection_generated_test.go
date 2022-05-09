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
)

func TestAccBigqueryConnectionConnection_bigqueryConnectionBasicExample(t *testing.T) {
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckBigqueryConnectionConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryConnectionConnection_bigqueryConnectionBasicExample(context),
			},
			{
				ResourceName:            "google_bigquery_connection.connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "cloud_sql.0.credential"},
			},
		},
	})
}

func testAccBigqueryConnectionConnection_bigqueryConnectionBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "instance" {
    name             = "tf-test-my-database-instance%{random_suffix}"
    database_version = "POSTGRES_11"
    region           = "us-central1"
    settings {
		tier = "db-f1-micro"
	}

    deletion_protection  = "%{deletion_protection}"
}

resource "google_sql_database" "db" {
    instance = google_sql_database_instance.instance.name
    name     = "db"
}

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    name = "user%{random_suffix}"
    instance = google_sql_database_instance.instance.name
    password = random_password.pwd.result
}

resource "google_bigquery_connection" "connection" {
    friendly_name = "👋"
    description   = "a riveting description"
    cloud_sql {
        instance_id = google_sql_database_instance.instance.connection_name
        database    = google_sql_database.db.name
        type        = "POSTGRES"
        credential {
          username = google_sql_user.user.name
          password = google_sql_user.user.password
        }
    }
}
`, context)
}

func TestAccBigqueryConnectionConnection_bigqueryConnectionFullExample(t *testing.T) {
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckBigqueryConnectionConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryConnectionConnection_bigqueryConnectionFullExample(context),
			},
			{
				ResourceName:            "google_bigquery_connection.connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "cloud_sql.0.credential"},
			},
		},
	})
}

func testAccBigqueryConnectionConnection_bigqueryConnectionFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "instance" {
    name             = "tf-test-my-database-instance%{random_suffix}"
    database_version = "POSTGRES_11"
    region           = "us-central1"
    settings {
		tier = "db-f1-micro"
	}

    deletion_protection  = "%{deletion_protection}"
}

resource "google_sql_database" "db" {
    instance = google_sql_database_instance.instance.name
    name     = "db"
}

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    name = "user%{random_suffix}"
    instance = google_sql_database_instance.instance.name
    password = random_password.pwd.result
}

resource "google_bigquery_connection" "connection" {
    connection_id = "tf-test-my-connection%{random_suffix}"
    location      = "US"
    friendly_name = "👋"
    description   = "a riveting description"
    cloud_sql {
        instance_id = google_sql_database_instance.instance.connection_name
        database    = google_sql_database.db.name
        type        = "POSTGRES"
        credential {
          username = google_sql_user.user.name
          password = google_sql_user.user.password
        }
    }
}
`, context)
}

func TestAccBigqueryConnectionConnection_bigqueryConnectionCloudResourceExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckBigqueryConnectionConnectionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryConnectionConnection_bigqueryConnectionCloudResourceExample(context),
			},
			{
				ResourceName:            "google_bigquery_connection.connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccBigqueryConnectionConnection_bigqueryConnectionCloudResourceExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_connection" "connection" {
   connection_id = "tf-test-my-connection%{random_suffix}"
   location      = "US"
   friendly_name = "👋"
   description   = "a riveting description"
   cloud_resource {}
}
`, context)
}

func testAccCheckBigqueryConnectionConnectionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_bigquery_connection" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{BigqueryConnectionBasePath}}projects/{{project}}/locations/{{location}}/connections/{{connection_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("BigqueryConnectionConnection still exists at %s", url)
			}
		}

		return nil
	}
}
