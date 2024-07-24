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

package dataplex_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccDataplexTask_dataplexTaskBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataplexTaskDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexTask_dataplexTaskBasicExample(context),
			},
			{
				ResourceName:            "google_dataplex_task.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "lake", "location", "task_id", "terraform_labels"},
			},
		},
	})
}

func testAccDataplexTask_dataplexTaskBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {

}

resource "google_dataplex_lake" "example" {
  name         = "tf-test-lake%{random_suffix}"
  location     = "us-central1"
  project = "%{project_name}"
}


resource "google_dataplex_task" "example" {

    task_id      = "tf-test-task%{random_suffix}"
    location     = "us-central1"
    lake         = google_dataplex_lake.example.name
    
    description = "Test Task Basic"
    
    display_name = "task-basic"

    labels = { "count": "3" }

    trigger_spec  {
        type = "RECURRING"
        disabled = false
        max_retries = 3
        start_time = "2023-10-02T15:01:23Z"
        schedule = "1 * * * *"
    }
    
    execution_spec {
        service_account = "${data.google_project.project.number}-compute@developer.gserviceaccount.com"
        project = "%{project_name}"
        max_job_execution_lifetime = "100s"
        kms_key = "234jn2kjn42k3n423"
    }
    
    spark {
        python_script_file = "gs://dataproc-examples/pyspark/hello-world/hello-world.py"

    }
    
    project = "%{project_name}"
    
}
`, context)
}

func TestAccDataplexTask_dataplexTaskSparkExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataplexTaskDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexTask_dataplexTaskSparkExample(context),
			},
			{
				ResourceName:            "google_dataplex_task.example_spark",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "lake", "location", "task_id", "terraform_labels"},
			},
		},
	})
}

func testAccDataplexTask_dataplexTaskSparkExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
# VPC network
resource "google_compute_network" "default" {
    name                    = "tf-test-workstation-cluster%{random_suffix}"
    auto_create_subnetworks = true
}

data "google_project" "project" {

}

resource "google_dataplex_lake" "example_spark" {
  name         = "tf-test-lake%{random_suffix}"
  location     = "us-central1"
  project = "%{project_name}"
}


resource "google_dataplex_task" "example_spark" {

    task_id      = "tf-test-task%{random_suffix}"
    location     = "us-central1"
    lake         = google_dataplex_lake.example_spark.name
    trigger_spec  {
        type = "ON_DEMAND"
    }
    
    description = "task-spark-terraform"

    execution_spec {
        service_account = "${data.google_project.project.number}-compute@developer.gserviceaccount.com"
        args = {
            TASK_ARGS  = "--output_location,gs://spark-job/task-result, --output_format, json"
        }

    }
    
    spark {
        infrastructure_spec  {
            batch {
                executors_count = 2
                max_executors_count = 100
            }
            container_image {
                image = "test-image"
                java_jars = ["test-java-jars.jar"]
                python_packages = ["gs://bucket-name/my/path/to/lib.tar.gz"]
                properties = { "name": "wrench", "mass": "1.3kg", "count": "3" }
            }
            vpc_network  {
                    network_tags = ["test-network-tag"]
                    sub_network = google_compute_network.default.id
                }
        }
        file_uris = ["gs://terrafrom-test/test.csv"]
        archive_uris = ["gs://terraform-test/test.csv"]
        sql_script = "show databases"
    }
    
    project = "%{project_name}"
    
}
`, context)
}

func TestAccDataplexTask_dataplexTaskNotebookExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataplexTaskDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexTask_dataplexTaskNotebookExample(context),
			},
			{
				ResourceName:            "google_dataplex_task.example_notebook",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "lake", "location", "task_id", "terraform_labels"},
			},
		},
	})
}

func testAccDataplexTask_dataplexTaskNotebookExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
# VPC network
resource "google_compute_network" "default" {
    name                    = "tf-test-workstation-cluster%{random_suffix}"
    auto_create_subnetworks = true
}


data "google_project" "project" {

}

resource "google_dataplex_lake" "example_notebook" {
  name         = "tf-test-lake%{random_suffix}"
  location     = "us-central1"
  project = "%{project_name}"
}


resource "google_dataplex_task" "example_notebook" {

    task_id      = "tf-test-task%{random_suffix}"
    location     = "us-central1"
    lake         = google_dataplex_lake.example_notebook.name
    trigger_spec  {
        type = "RECURRING"
        schedule = "1 * * * *"
    }
    
    execution_spec {
        service_account = "${data.google_project.project.number}-compute@developer.gserviceaccount.com"
        args = {
            TASK_ARGS  = "--output_location,gs://spark-job-jars-anrajitha/task-result, --output_format, json"
        }
    }
    notebook {
        notebook = "gs://terraform-test/test-notebook.ipynb"
        infrastructure_spec  {
            batch {
                executors_count = 2
                max_executors_count = 100
            }
            container_image {
                image = "test-image"
                java_jars = ["test-java-jars.jar"]
                python_packages = ["gs://bucket-name/my/path/to/lib.tar.gz"]
                properties = { "name": "wrench", "mass": "1.3kg", "count": "3" }
            }
            vpc_network  {
                    network_tags = ["test-network-tag"]
                    network = google_compute_network.default.id
                }
        }
        file_uris = ["gs://terraform-test/test.csv"]
        archive_uris = ["gs://terraform-test/test.csv"]

    }
    project = "%{project_name}"
    
    
}
`, context)
}

func testAccCheckDataplexTaskDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dataplex_task" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DataplexBasePath}}projects/{{project}}/locations/{{location}}/lakes/{{lake}}/tasks/{{task_id}}")
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
				return fmt.Errorf("DataplexTask still exists at %s", url)
			}
		}

		return nil
	}
}
