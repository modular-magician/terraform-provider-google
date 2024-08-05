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

package cloudfunctions2_test

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

func TestAccCloudfunctions2function_cloudfunctions2BasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"location":      "us-central1",
		"zip_path":      "./test-fixtures/function-source.zip",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudfunctions2functionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudfunctions2function_cloudfunctions2BasicExample(context),
			},
			{
				ResourceName:            "google_cloudfunctions2_function.function",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"build_config.0.source.0.storage_source.0.bucket", "build_config.0.source.0.storage_source.0.object", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccCloudfunctions2function_cloudfunctions2BasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
locals {
  project = "%{project}" # Google Cloud Platform Project ID
}

resource "google_storage_bucket" "bucket" {
  name     = "${local.project}-tf-test-gcf-source%{random_suffix}"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}
 
resource "google_cloudfunctions2_function" "function" {
  name = "tf-test-function-v2%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime = "nodejs16"
    entry_point = "helloHttp"  # Set the entry point 
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }
 
  service_config {
    max_instance_count  = 1
    available_memory    = "256M"
    timeout_seconds     = 60
  }
}
`, context)
}

func TestAccCloudfunctions2function_cloudfunctions2FullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":             envvar.GetTestProjectFromEnv(),
		"location":            "us-central1",
		"primary_resource_id": "terraform-test",
		"zip_path":            "./test-fixtures/function-source-pubsub.zip",
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudfunctions2functionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudfunctions2function_cloudfunctions2FullExample(context),
			},
			{
				ResourceName:            "google_cloudfunctions2_function.function",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"build_config.0.source.0.storage_source.0.bucket", "build_config.0.source.0.storage_source.0.object", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccCloudfunctions2function_cloudfunctions2FullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
locals {
  project = "%{project}" # Google Cloud Platform Project ID
}

resource "google_service_account" "account" {
  account_id = "tf-test-gcf-sa%{random_suffix}"
  display_name = "Test Service Account"
}

resource "google_pubsub_topic" "topic" {
  name = "tf-test-functions2-topic%{random_suffix}"
}

resource "google_storage_bucket" "bucket" {
  name     = "${local.project}-tf-test-gcf-source%{random_suffix}"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}
 
resource "google_cloudfunctions2_function" "function" {
  name = "tf-test-gcf-function%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime = "nodejs16"
    entry_point = "helloPubSub"  # Set the entry point 
    environment_variables = {
        BUILD_CONFIG_TEST = "build_test"
    }
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }
 
  service_config {
    max_instance_count  = 3
    min_instance_count = 1
    available_memory    = "4Gi"
    timeout_seconds     = 60
    max_instance_request_concurrency = 80
    available_cpu = "4"
    environment_variables = {
        SERVICE_CONFIG_TEST = "config_test"
        SERVICE_CONFIG_DIFF_TEST = google_service_account.account.email
    }
    ingress_settings = "ALLOW_INTERNAL_ONLY"
    all_traffic_on_latest_revision = true
    service_account_email = google_service_account.account.email
  }

  event_trigger {
    trigger_region = "us-central1"
    event_type = "google.cloud.pubsub.topic.v1.messagePublished"
    pubsub_topic = google_pubsub_topic.topic.id
    retry_policy = "RETRY_POLICY_RETRY"
  }
}
`, context)
}

func TestAccCloudfunctions2function_cloudfunctions2BasicGcsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":             envvar.GetTestProjectFromEnv(),
		"policyChanged":       acctest.BootstrapPSARole(t, "service-", "gcp-sa-pubsub", "roles/cloudkms.cryptoKeyEncrypterDecrypter"),
		"primary_resource_id": "terraform-test",
		"zip_path":            "./test-fixtures/function-source-eventarc-gcs.zip",
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudfunctions2functionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudfunctions2function_cloudfunctions2BasicGcsExample(context),
			},
			{
				ResourceName:            "google_cloudfunctions2_function.function",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"build_config.0.source.0.storage_source.0.bucket", "build_config.0.source.0.storage_source.0.object", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccCloudfunctions2function_cloudfunctions2BasicGcsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_storage_bucket" "source-bucket" {
  name     = "tf-test-gcf-source-bucket%{random_suffix}"
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.source-bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}

resource "google_storage_bucket" "trigger-bucket" {
  name     = "tf-test-gcf-trigger-bucket%{random_suffix}"
  location = "us-central1" # The trigger must be in the same location as the bucket
  uniform_bucket_level_access = true
}

data "google_storage_project_service_account" "gcs_account" {
}

# To use GCS CloudEvent triggers, the GCS service account requires the Pub/Sub Publisher(roles/pubsub.publisher) IAM role in the specified project.
# (See https://cloud.google.com/eventarc/docs/run/quickstart-storage#before-you-begin)
resource "google_project_iam_member" "gcs-pubsub-publishing" {
  project = "%{project}"
  role    = "roles/pubsub.publisher"
  member  = "serviceAccount:${data.google_storage_project_service_account.gcs_account.email_address}"
}

resource "google_service_account" "account" {
  account_id   = "tf-test-gcf-sa%{random_suffix}"
  display_name = "Test Service Account - used for both the cloud function and eventarc trigger in the test"
}

# Permissions on the service account used by the function and Eventarc trigger
resource "google_project_iam_member" "invoking" {
  project = "%{project}"
  role    = "roles/run.invoker"
  member  = "serviceAccount:${google_service_account.account.email}"
  depends_on = [google_project_iam_member.gcs-pubsub-publishing]
}

resource "google_project_iam_member" "event-receiving" {
  project = "%{project}"
  role    = "roles/eventarc.eventReceiver"
  member  = "serviceAccount:${google_service_account.account.email}"
  depends_on = [google_project_iam_member.invoking]
}

resource "google_project_iam_member" "artifactregistry-reader" {
  project = "%{project}"
  role     = "roles/artifactregistry.reader"
  member   = "serviceAccount:${google_service_account.account.email}"
  depends_on = [google_project_iam_member.event-receiving]
}

resource "google_cloudfunctions2_function" "function" {
  depends_on = [
    google_project_iam_member.event-receiving,
    google_project_iam_member.artifactregistry-reader,
  ]
  name = "tf-test-gcf-function%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime     = "nodejs12"
    entry_point = "entryPoint" # Set the entry point in the code
    environment_variables = {
      BUILD_CONFIG_TEST = "build_test"
    }
    source {
      storage_source {
        bucket = google_storage_bucket.source-bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }
 
  service_config {
    max_instance_count  = 3
    min_instance_count = 1
    available_memory    = "256M"
    timeout_seconds     = 60
    environment_variables = {
        SERVICE_CONFIG_TEST = "config_test"
    }
    ingress_settings = "ALLOW_INTERNAL_ONLY"
    all_traffic_on_latest_revision = true
    service_account_email = google_service_account.account.email
  }

  event_trigger {
    event_type = "google.cloud.storage.object.v1.finalized"
    retry_policy = "RETRY_POLICY_RETRY"
    service_account_email = google_service_account.account.email
    event_filters {
      attribute = "bucket"
      value = google_storage_bucket.trigger-bucket.name
    }
  }
}
`, context)
}

func TestAccCloudfunctions2function_cloudfunctions2BasicAuditlogsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":             envvar.GetTestProjectFromEnv(),
		"policyChanged":       acctest.BootstrapPSARole(t, "service-", "gcp-sa-pubsub", "roles/cloudkms.cryptoKeyEncrypterDecrypter"),
		"primary_resource_id": "terraform-test",
		"zip_path":            "./test-fixtures/function-source-eventarc-gcs.zip",
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudfunctions2functionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudfunctions2function_cloudfunctions2BasicAuditlogsExample(context),
			},
			{
				ResourceName:            "google_cloudfunctions2_function.function",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"build_config.0.source.0.storage_source.0.bucket", "build_config.0.source.0.storage_source.0.object", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccCloudfunctions2function_cloudfunctions2BasicAuditlogsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
# This example follows the examples shown in this Google Cloud Community blog post
# https://medium.com/google-cloud/applying-a-path-pattern-when-filtering-in-eventarc-f06b937b4c34
# and the docs:
# https://cloud.google.com/eventarc/docs/path-patterns

resource "google_storage_bucket" "source-bucket" {
  name     = "tf-test-gcf-source-bucket%{random_suffix}"
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.source-bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}

resource "google_service_account" "account" {
  account_id   = "tf-test-gcf-sa%{random_suffix}"
  display_name = "Test Service Account - used for both the cloud function and eventarc trigger in the test"
}

# Note: The right way of listening for Cloud Storage events is to use a Cloud Storage trigger.
# Here we use Audit Logs to monitor the bucket so path patterns can be used in the example of
# google_cloudfunctions2_function below (Audit Log events have path pattern support)
resource "google_storage_bucket" "audit-log-bucket" {
  name     = "tf-test-gcf-auditlog-bucket%{random_suffix}"
  location = "us-central1"  # The trigger must be in the same location as the bucket
  uniform_bucket_level_access = true
}

# Permissions on the service account used by the function and Eventarc trigger
resource "google_project_iam_member" "invoking" {
  project = "%{project}"
  role    = "roles/run.invoker"
  member  = "serviceAccount:${google_service_account.account.email}"
}

resource "google_project_iam_member" "event-receiving" {
  project = "%{project}"
  role    = "roles/eventarc.eventReceiver"
  member  = "serviceAccount:${google_service_account.account.email}"
  depends_on = [google_project_iam_member.invoking]
}

resource "google_project_iam_member" "artifactregistry-reader" {
  project = "%{project}"
  role     = "roles/artifactregistry.reader"
  member   = "serviceAccount:${google_service_account.account.email}"
  depends_on = [google_project_iam_member.event-receiving]
}

resource "google_cloudfunctions2_function" "function" {
  depends_on = [
    google_project_iam_member.event-receiving,
    google_project_iam_member.artifactregistry-reader,
  ]
  name = "tf-test-gcf-function%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime     = "nodejs12"
    entry_point = "entryPoint" # Set the entry point in the code
    environment_variables = {
      BUILD_CONFIG_TEST = "build_test"
    }
    source {
      storage_source {
        bucket = google_storage_bucket.source-bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }
 
  service_config {
    max_instance_count  = 3
    min_instance_count = 1
    available_memory    = "256M"
    timeout_seconds     = 60
    environment_variables = {
        SERVICE_CONFIG_TEST = "config_test"
    }
    ingress_settings = "ALLOW_INTERNAL_ONLY"
    all_traffic_on_latest_revision = true
    service_account_email = google_service_account.account.email
  }

  event_trigger {
    trigger_region = "us-central1" # The trigger must be in the same location as the bucket
    event_type = "google.cloud.audit.log.v1.written"
    retry_policy = "RETRY_POLICY_RETRY"
    service_account_email = google_service_account.account.email
    event_filters {
      attribute = "serviceName"
      value = "storage.googleapis.com"
    }
    event_filters {
      attribute = "methodName"
      value = "storage.objects.create"
    }
    event_filters {
      attribute = "resourceName"
      value = "/projects/_/buckets/${google_storage_bucket.audit-log-bucket.name}/objects/*.txt" # Path pattern selects all .txt files in the bucket
      operator = "match-path-pattern" # This allows path patterns to be used in the value field
    }
  }
}
`, context)
}

func TestAccCloudfunctions2function_cloudfunctions2BasicBuilderExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"location":      "us-central1",
		"zip_path":      "./test-fixtures/function-source.zip",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		CheckDestroy: testAccCheckCloudfunctions2functionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudfunctions2function_cloudfunctions2BasicBuilderExample(context),
			},
			{
				ResourceName:            "google_cloudfunctions2_function.function",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"build_config.0.source.0.storage_source.0.bucket", "build_config.0.source.0.storage_source.0.object", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccCloudfunctions2function_cloudfunctions2BasicBuilderExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
locals {
  project = "%{project}" # Google Cloud Platform Project ID
}

resource "google_service_account" "account" {
  account_id = "tf-test-gcf-sa%{random_suffix}"
  display_name = "Test Service Account"
}

resource "google_project_iam_member" "log_writer" {
  project = google_service_account.account.project
  role    = "roles/logging.logWriter"
  member  = "serviceAccount:${google_service_account.account.email}"
}

resource "google_project_iam_member" "artifact_registry_writer" {
  project = google_service_account.account.project
  role    = "roles/artifactregistry.writer"
  member  = "serviceAccount:${google_service_account.account.email}"
}

resource "google_project_iam_member" "storage_object_admin" {
  project = google_service_account.account.project
  role    = "roles/storage.objectAdmin"
  member  = "serviceAccount:${google_service_account.account.email}"
}

resource "google_storage_bucket" "bucket" {
  name     = "${local.project}-tf-test-gcf-source%{random_suffix}"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}

# builder permissions need to stablize before it can pull the source zip
resource "time_sleep" "wait_60s" {
  create_duration = "60s"

  depends_on = [
    google_project_iam_member.log_writer,
    google_project_iam_member.artifact_registry_writer,
    google_project_iam_member.storage_object_admin,
  ]
}
 
resource "google_cloudfunctions2_function" "function" {
  name = "tf-test-function-v2%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime = "nodejs16"
    entry_point = "helloHttp"  # Set the entry point 
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
    service_account = google_service_account.account.id
  }
 
  service_config {
    max_instance_count  = 1
    available_memory    = "256M"
    timeout_seconds     = 60
  }

  depends_on = [time_sleep.wait_60s]
}
`, context)
}

func TestAccCloudfunctions2function_cloudfunctions2SecretEnvExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"location":      "us-central1",
		"policyChanged": acctest.BootstrapPSARole(t, "service-", "gcp-sa-pubsub", "roles/cloudkms.cryptoKeyEncrypterDecrypter"),
		"zip_path":      "./test-fixtures/function-source.zip",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudfunctions2functionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudfunctions2function_cloudfunctions2SecretEnvExample(context),
			},
			{
				ResourceName:            "google_cloudfunctions2_function.function",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"build_config.0.source.0.storage_source.0.bucket", "build_config.0.source.0.storage_source.0.object", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccCloudfunctions2function_cloudfunctions2SecretEnvExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
locals {
  project = "%{project}" # Google Cloud Platform Project ID
}

resource "google_storage_bucket" "bucket" {
  name     = "${local.project}-tf-test-gcf-source%{random_suffix}"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}
 
resource "google_cloudfunctions2_function" "function" {
  name = "tf-test-function-secret%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime = "nodejs16"
    entry_point = "helloHttp"  # Set the entry point 
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }
 
  service_config {
    max_instance_count  = 1
    available_memory    = "256M"
    timeout_seconds     = 60

    secret_environment_variables {
      key        = "TEST"
      project_id = local.project
      secret     = google_secret_manager_secret.secret.secret_id
      version    = "latest"
    }
  }
  depends_on = [google_secret_manager_secret_version.secret]
}

resource "google_secret_manager_secret" "secret" {
  secret_id = "secret%{random_suffix}"

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
    }
  }  
}

resource "google_secret_manager_secret_version" "secret" {
  secret = google_secret_manager_secret.secret.name

  secret_data = "secret%{random_suffix}"
  enabled = true
}
`, context)
}

func TestAccCloudfunctions2function_cloudfunctions2SecretVolumeExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"location":      "us-central1",
		"policyChanged": acctest.BootstrapPSARole(t, "service-", "gcp-sa-pubsub", "roles/cloudkms.cryptoKeyEncrypterDecrypter"),
		"zip_path":      "./test-fixtures/function-source.zip",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudfunctions2functionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudfunctions2function_cloudfunctions2SecretVolumeExample(context),
			},
			{
				ResourceName:            "google_cloudfunctions2_function.function",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"build_config.0.source.0.storage_source.0.bucket", "build_config.0.source.0.storage_source.0.object", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccCloudfunctions2function_cloudfunctions2SecretVolumeExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
locals {
  project = "%{project}" # Google Cloud Platform Project ID
}

resource "google_storage_bucket" "bucket" {
  name     = "${local.project}-tf-test-gcf-source%{random_suffix}"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}
 
resource "google_cloudfunctions2_function" "function" {
  name = "tf-test-function-secret%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime = "nodejs16"
    entry_point = "helloHttp"  # Set the entry point 
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
  }
 
  service_config {
    max_instance_count  = 1
    available_memory    = "256M"
    timeout_seconds     = 60

    secret_volumes {
      mount_path = "/etc/secrets"
      project_id = local.project
      secret     = google_secret_manager_secret.secret.secret_id
    }
  }
  depends_on = [google_secret_manager_secret_version.secret]
}

resource "google_secret_manager_secret" "secret" {
  secret_id = "secret%{random_suffix}"

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
    }
  }  
}

resource "google_secret_manager_secret_version" "secret" {
  secret = google_secret_manager_secret.secret.name

  secret_data = "secret%{random_suffix}"
  enabled = true
}
`, context)
}

func TestAccCloudfunctions2function_cloudfunctions2PrivateWorkerpoolExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"location":      "us-central1",
		"zip_path":      "./test-fixtures/function-source.zip",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCloudfunctions2functionDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudfunctions2function_cloudfunctions2PrivateWorkerpoolExample(context),
			},
			{
				ResourceName:            "google_cloudfunctions2_function.function",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"build_config.0.source.0.storage_source.0.bucket", "build_config.0.source.0.storage_source.0.object", "labels", "location", "terraform_labels"},
			},
		},
	})
}

func testAccCloudfunctions2function_cloudfunctions2PrivateWorkerpoolExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
locals {
  project = "%{project}" # Google Cloud Platform Project ID
}

resource "google_storage_bucket" "bucket" {
  name     = "${local.project}-tf-test-gcf-source%{random_suffix}"  # Every bucket name must be globally unique
  location = "US"
  uniform_bucket_level_access = true
}
 
resource "google_storage_bucket_object" "object" {
  name   = "function-source.zip"
  bucket = google_storage_bucket.bucket.name
  source = "%{zip_path}"  # Add path to the zipped function source code
}

resource "google_cloudbuild_worker_pool" "pool" {
  name = "workerpool%{random_suffix}"
  location = "us-central1"
  worker_config {
    disk_size_gb = 100
    machine_type = "e2-standard-8"
    no_external_ip = false
  }
}
 
resource "google_cloudfunctions2_function" "function" {
  name = "tf-test-function-workerpool%{random_suffix}"
  location = "us-central1"
  description = "a new function"
 
  build_config {
    runtime = "nodejs16"
    entry_point = "helloHttp"  # Set the entry point 
    source {
      storage_source {
        bucket = google_storage_bucket.bucket.name
        object = google_storage_bucket_object.object.name
      }
    }
    worker_pool = google_cloudbuild_worker_pool.pool.id
  }
 
  service_config {
    max_instance_count  = 1
    available_memory    = "256M"
    timeout_seconds     = 60
  }
}
`, context)
}

func testAccCheckCloudfunctions2functionDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloudfunctions2_function" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{Cloudfunctions2BasePath}}projects/{{project}}/locations/{{location}}/functions/{{name}}")
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
				return fmt.Errorf("Cloudfunctions2function still exists at %s", url)
			}
		}

		return nil
	}
}
