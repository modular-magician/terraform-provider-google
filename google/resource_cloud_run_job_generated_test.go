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

func TestAccCloudRunJob_cloudRunJobBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunJob_cloudRunJobBasicExample(context),
			},
			{
				ResourceName:            "google_cloud_run_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunJob_cloudRunJobBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_job" "default" {
  name     = ""
  location = "us-central1"
  provider = google-beta

  metadata {
    annotations = {
      "run.googleapis.com/launch-stage" = "BETA"
      generated-by = "magic-modules"
    }
  }

  template {
    spec {
      template {
        spec {
          containers {
            image = "us-docker.pkg.dev/cloudrun/container/hello"
          }
        }
      }
    }
  }
}
`, context)
}

func TestAccCloudRunJob_cloudRunJobSqlExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunJob_cloudRunJobSqlExample(context),
			},
			{
				ResourceName:            "google_cloud_run_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "autogenerate_revision_name"},
			},
		},
	})
}

func testAccCloudRunJob_cloudRunJobSqlExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_job" "default" {
  name     = ""
  location = "us-central1"
  provider = google-beta

  metadata {
    annotations = {
      "run.googleapis.com/launch-stage" = "BETA"
      generated-by = "magic-modules"
    }
  }

  template {
    spec {
      template {
        spec {
          containers {
            image = "us-docker.pkg.dev/cloudrun/container/hello"
          }
        }
      }
    }
    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale"      = "1000"
        "run.googleapis.com/cloudsql-instances" = google_sql_database_instance.instance.connection_name
        "run.googleapis.com/client-name"        = "terraform"
      }
    }
  }

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }

    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale"      = "1000"
        "run.googleapis.com/cloudsql-instances" = google_sql_database_instance.instance.connection_name
        "run.googleapis.com/client-name"        = "terraform"
      }
    }
  }
  autogenerate_revision_name = true
}

resource "google_sql_database_instance" "instance" {
  name             = "tf-test-cloudrun-sql%{random_suffix}"
  region           = "us-east1"
  database_version = "MYSQL_5_7"
  settings {
    tier = "db-f1-micro"
  }

  deletion_protection  = "%{deletion_protection}"
}
`, context)
}

func TestAccCloudRunJob_cloudRunJobNoauthExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunJob_cloudRunJobNoauthExample(context),
			},
			{
				ResourceName:            "google_cloud_run_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location"},
			},
		},
	})
}

func testAccCloudRunJob_cloudRunJobNoauthExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_job" "default" {
  name     = ""
  location = "us-central1"
  provider = google-beta

  metadata {
    annotations = {
      "run.googleapis.com/launch-stage" = "BETA"
      generated-by = "magic-modules"
    }
  }
  template {
    spec {
      template {
        spec {
          containers {
            image = "us-docker.pkg.dev/cloudrun/container/hello"
          }
        }
      }
    }
  }
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_job_iam_policy" "noauth" {
  provider = google-beta
  location    = google_cloud_run_job.default.location
  project     = google_cloud_run_job.default.project
  job     = google_cloud_run_job.default.name

  policy_data = data.google_iam_policy.noauth.policy_data
}
`, context)
}

func TestAccCloudRunJob_cloudRunJobMultipleEnvironmentVariablesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunJob_cloudRunJobMultipleEnvironmentVariablesExample(context),
			},
			{
				ResourceName:            "google_cloud_run_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "autogenerate_revision_name"},
			},
		},
	})
}

func testAccCloudRunJob_cloudRunJobMultipleEnvironmentVariablesExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloud_run_job" "default" {
  name     = ""
  location = "us-central1"
  provider = google-beta

  metadata {
    annotations = {
      "run.googleapis.com/launch-stage" = "BETA"
      generated-by = "magic-modules"
    }
  }
  template {
    spec {
      template {
        spec {
          containers {
            image = "us-docker.pkg.dev/cloudrun/container/hello"
            env {
              name = "SOURCE"
              value = "remote"
            }
            env {
              name = "TARGET"
              value = "home"
            }
          }
        }
      }
    }
  }

  autogenerate_revision_name = true

  lifecycle {
    ignore_changes = [
        metadata.0.annotations,
    ]
  }
}
`, context)
}

func TestAccCloudRunJob_cloudRunJobSecretEnvironmentVariablesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunJob_cloudRunJobSecretEnvironmentVariablesExample(context),
			},
			{
				ResourceName:            "google_cloud_run_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "autogenerate_revision_name"},
			},
		},
	})
}

func testAccCloudRunJob_cloudRunJobSecretEnvironmentVariablesExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_project" "project" {
}

resource "google_secret_manager_secret" "secret" {
  secret_id = "secret%{random_suffix}"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret-version-data" {
  secret = google_secret_manager_secret.secret.name
  secret_data = "secret-data"
}

resource "google_secret_manager_secret_iam_member" "secret-access" {
  secret_id = google_secret_manager_secret.secret.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret]
}

resource "google_cloud_run_job" "default" {
  name     = ""
  location = "us-central1"
  provider = google-beta

  metadata {
    annotations = {
      "run.googleapis.com/launch-stage" = "BETA"
      generated-by = "magic-modules"
    }
  }
  template {
    spec {
      template {
        spec {
          containers {
            image = "us-docker.pkg.dev/cloudrun/container/hello"
            env {
              name = "SECRET_ENV_VAR"
              value_from {
                secret_key_ref {
                  name = google_secret_manager_secret.secret.secret_id
                  key = "1"
                }
              }
            }
          }
        }
      }
    }
  }

  autogenerate_revision_name = true

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }

  depends_on = [google_secret_manager_secret_version.secret-version-data]
}
`, context)
}

func TestAccCloudRunJob_cloudRunJobSecretVolumesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudRunJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudRunJob_cloudRunJobSecretVolumesExample(context),
			},
			{
				ResourceName:            "google_cloud_run_job.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "autogenerate_revision_name"},
			},
		},
	})
}

func testAccCloudRunJob_cloudRunJobSecretVolumesExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_project" "project" {
}

resource "google_secret_manager_secret" "secret" {
  secret_id = "secret%{random_suffix}"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret-version-data" {
  secret = google_secret_manager_secret.secret.name
  secret_data = "secret-data"
}

resource "google_secret_manager_secret_iam_member" "secret-access" {
  secret_id = google_secret_manager_secret.secret.id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
  depends_on = [google_secret_manager_secret.secret]
}

resource "google_cloud_run_job" "default" {
  name     = ""
  location = "us-central1"
  provider = google-beta

  metadata {
    annotations = {
      "run.googleapis.com/launch-stage" = "BETA"
      generated-by = "magic-modules"
    }
  }

  template {
    spec {
      template {
        spec {
          containers {
            image = "us-docker.pkg.dev/cloudrun/container/hello"
            volume_mounts {
              name = "a-volume"
              mount_path = "/secrets"
            }
          }
          volumes {
            name = "a-volume"
            secret {
              secret_name = google_secret_manager_secret.secret.secret_id
              default_mode = 292 # 0444
              items {
                key = "1"
                path = "my-secret"
                mode = 256 # 0400
              }
            }
          }
        }
      }
    }
  }

  autogenerate_revision_name = true

  lifecycle {
    ignore_changes = [
      metadata.0.annotations,
    ]
  }

  depends_on = [google_secret_manager_secret_version.secret-version-data]
}
`, context)
}

func testAccCheckCloudRunJobDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloud_run_job" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{CloudRunBasePath}}apis/run.googleapis.com/v1/namespaces/{{project}}/jobs/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil, isCloudRunCreationConflict)
			if err == nil {
				return fmt.Errorf("CloudRunJob still exists at %s", url)
			}
		}

		return nil
	}
}
