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
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccDataLossPreventionJob_dlpRiskJobWithJobIdExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       acctest.GetTestProjectFromEnv(),
		"name":          "tf_test_" + RandString(t, 10),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionJob_dlpRiskJobWithJobIdExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_job.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"job_id", "parent", "state", "end_time"},
			},
		},
	})
}

func testAccDataLossPreventionJob_dlpRiskJobWithJobIdExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_job" "basic" {
  parent = "projects/%{project}"
  job_id = "%{name}"

  risk_job {
    actions {
      job_notification_emails {}
    }
    source_table {
      project_id = "%{project}"
      dataset_id = google_bigquery_dataset.default.dataset_id
      table_id   = google_bigquery_table.default.table_id
    }
    privacy_metric {
      numerical_stats_config {
        field {
          name = "permalink"
        }
      }
    }
  }
}

resource "google_bigquery_dataset" "default" {
  dataset_id                  = "%{name}"
  friendly_name               = "terraform-test"
  description                 = "Description for the dataset created by terraform"
  location                    = "US"
  default_table_expiration_ms = 3600000

  labels = {
    env = "default"
  }
}

resource "google_bigquery_table" "default" {
  dataset_id          = google_bigquery_dataset.default.dataset_id
  table_id            = "%{name}"
  deletion_protection = false

  time_partitioning {
    type = "DAY"
  }

  labels = {
    env = "default"
  }

  schema = <<EOF
   [
    {
     "name": "permalink",
     "type": "NUMERIC",
     "mode": "NULLABLE",
     "description": "The Permalink"
    },
    {
     "name": "state",
     "type": "STRING",
     "mode": "NULLABLE",
     "description": "State where the head office is located"
    }
   ]
  EOF
}
`, context)
}

func TestAccDataLossPreventionJob_dlpRiskJobBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       acctest.GetTestProjectFromEnv(),
		"name":          "tf_test_" + RandString(t, 10),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionJob_dlpRiskJobBasicExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_job.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"job_id", "parent", "state", "end_time"},
			},
		},
	})
}

func testAccDataLossPreventionJob_dlpRiskJobBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_data_loss_prevention_job" "basic" {
  parent = "projects/%{project}"

  risk_job {
    actions {
      job_notification_emails {}
    }
    source_table {
      project_id = "%{project}"
      dataset_id = google_bigquery_dataset.default.dataset_id
      table_id   = google_bigquery_table.default.table_id
    }
    privacy_metric {
      categorical_stats_config {
        field {
          name = "state"
        }
      }
    }
  }
}

resource "google_bigquery_dataset" "default" {
  dataset_id                  = "%{name}"
  friendly_name               = "terraform-test"
  description                 = "Description for the dataset created by terraform"
  location                    = "US"
  default_table_expiration_ms = 3600000

  labels = {
    env = "default"
  }
}

resource "google_bigquery_table" "default" {
  dataset_id          = google_bigquery_dataset.default.dataset_id
  table_id            = "%{name}"
  deletion_protection = false

  time_partitioning {
    type = "DAY"
  }

  labels = {
    env = "default"
  }

  schema = <<EOF
   [
    {
     "name": "permalink",
     "type": "NUMERIC",
     "mode": "NULLABLE",
     "description": "The Permalink"
    },
    {
     "name": "state",
     "type": "STRING",
     "mode": "NULLABLE",
     "description": "State where the head office is located"
    }
   ]
  EOF
}
`, context)
}

func testAccCheckDataLossPreventionJobDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_data_loss_prevention_job" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := acctest.ReplaceVarsForTest(config, rs, "{{DataLossPreventionBasePath}}{{parent}}/dlpJobs/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("DataLossPreventionJob still exists at %s", url)
			}
		}

		return nil
	}
}
