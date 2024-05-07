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

package datalossprevention_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"location":      envvar.GetTestRegionFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigBasicExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent", "location"},
			},
		},
	})
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_discovery_config" "basic" {
	parent = "projects/%{project}/locations/us"
    location = "us"
    status = "RUNNING"

    targets {
        big_query_target {
            filter {
                other_tables {}
            }
        }
    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"]
}

resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "My description"
	display_name = "display_name"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
    }
}
`, context)
}

func TestAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigActionsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigActionsExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.actions",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent", "location"},
			},
		},
	})
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigActionsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_discovery_config" "actions" {
	parent = "projects/%{project}/locations/us"
    location = "us"
    status = "RUNNING"

    targets {
        big_query_target {
            filter {
                other_tables {}
            }
        }
    }
    actions {
        export_data {
            profile_table {
                project_id = "project"
                dataset_id = "dataset"
                table_id = "table"
            }
        }
    }
    actions { 
        pub_sub_notification {
            topic = "projects/%{project}/topics/${google_pubsub_topic.actions.name}"
            event = "NEW_PROFILE"
            pubsub_condition {
                expressions {
                    logical_operator = "OR"
                    conditions {
                        minimum_sensitivity_score = "HIGH"
                    }
                }
            }
            detail_of_message = "TABLE_PROFILE"
        }
    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"] 
}

resource "google_pubsub_topic" "actions" {
    name = "fake-topic"
}

resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "My description"
	display_name = "display_name"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
    }
}
`, context)
}

func TestAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigOrgRunningExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"organization":  envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigOrgRunningExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.org_running",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent", "location"},
			},
		},
	})
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigOrgRunningExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_discovery_config" "org_running" {
	parent = "organizations/%{organization}/locations/us"
    location = "us"

    targets {
        big_query_target {
            filter {
                other_tables {}
            }
        }
    }
    org_config {
        project_id = "%{project}"
        location {
            organization_id = "%{organization}"
        }
    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"] 
    status = "RUNNING"
}

resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "My description"
	display_name = "display_name"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
    }
}
`, context)
}

func TestAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigOrgFolderPausedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"organization":  envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigOrgFolderPausedExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.org_folder_paused",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent", "location"},
			},
		},
	})
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigOrgFolderPausedExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_discovery_config" "org_folder_paused" {
	parent = "organizations/%{organization}/locations/us"
    location = "us"

    targets {
        big_query_target {
            filter {
                other_tables {}
            }
        }
    }
    org_config {
        project_id = "%{project}"
        location {
            folder_id = 123
        }
    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"]
    status = "PAUSED"
}

resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "My description"
	display_name = "display_name"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
    }
}
`, context)
}

func TestAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigConditionsCadenceExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigConditionsCadenceExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.conditions_cadence",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent", "location"},
			},
		},
	})
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigConditionsCadenceExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_discovery_config" "conditions_cadence" {
	parent = "projects/%{project}/locations/us"
    location = "us"
    status = "RUNNING"

    targets {
        big_query_target {
            filter {
                other_tables {}
            }
            conditions {
                type_collection = "BIG_QUERY_COLLECTION_ALL_TYPES"
            }
            cadence {
                schema_modified_cadence {
                    types = ["SCHEMA_NEW_COLUMNS"]
                    frequency = "UPDATE_FREQUENCY_DAILY"
                }
                table_modified_cadence {
                    types = ["TABLE_MODIFIED_TIMESTAMP"]
                    frequency = "UPDATE_FREQUENCY_DAILY"
                }
            }
        }
    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"]
}

resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "My description"
	display_name = "display_name"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
    }
}
`, context)
}

func TestAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigFilterRegexesAndConditionsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigFilterRegexesAndConditionsExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.filter_regexes_and_conditions",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent", "location"},
			},
		},
	})
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigFilterRegexesAndConditionsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_discovery_config" "filter_regexes_and_conditions" {
	parent = "projects/%{project}/locations/us"
    location = "us"
    status = "RUNNING"

    targets {
        big_query_target {
            filter {
                tables {
                    include_regexes {
                        patterns {
                            project_id_regex = ".*"
                            dataset_id_regex = ".*"
                            table_id_regex = ".*"
                        }
                    }
                }
            }
            conditions {
                created_after = "2023-10-02T15:01:23Z"
                types {
                    types = ["BIG_QUERY_TABLE_TYPE_TABLE", "BIG_QUERY_TABLE_TYPE_EXTERNAL_BIG_LAKE"]
                }
                or_conditions {
                    min_row_count = 10
                    min_age = "10800s"
                }
            }
        }
    }
    targets {
        big_query_target {
            filter {
                other_tables {}
            }
        }
    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"] 
}

resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "My description"
	display_name = "display_name"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
    }
}
`, context)
}

func TestAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigCloudSqlExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigCloudSqlExample(context),
			},
			{
				ResourceName:            "google_data_loss_prevention_discovery_config.cloud_sql",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent", "location"},
			},
		},
	})
}

func testAccDataLossPreventionDiscoveryConfig_dlpDiscoveryConfigCloudSqlExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_data_loss_prevention_discovery_config" "cloud_sql" {
	parent = "projects/%{project}/locations/us"
    location = "us"
    status = "RUNNING"

    targets {
        cloud_sql_target {
            filter {
                collection {
                    include_regexes {
                        patterns {
                            project_id_regex = ".*"
                            instance_regex = ".*"
                            database_regex = ".*"
                            database_resource_name_regex = "mytable.*"
                        }
                    }
                }
            }
            conditions {
                database_engines = ["ALL_SUPPORTED_DATABASE_ENGINES"]
                types = ["DATABASE_RESOURCE_TYPE_ALL_SUPPORTED_TYPES"]
            }
            generation_cadence {
                schema_modified_cadence {
                    types = ["NEW_COLUMNS", "REMOVED_COLUMNS"]
                    frequency = "UPDATE_FREQUENCY_DAILY"
                }
                refresh_frequency = "UPDATE_FREQUENCY_MONTHLY"
            }
        }
    }
    targets {
        cloud_sql_target {
            filter {
                collection {
                    include_regexes {
                        patterns {
                            project_id_regex = ".*"
                            instance_regex = ".*"
                            database_regex = "do-not-scan.*"
                            database_resource_name_regex = ".*"
                        }
                    }
                }
            }
            disabled {}
        }
    }
    targets {
        cloud_sql_target {
            filter {
                others {}
            }
            generation_cadence {
                schema_modified_cadence {
                    types = ["NEW_COLUMNS"]
                    frequency = "UPDATE_FREQUENCY_MONTHLY"
                }
                refresh_frequency = "UPDATE_FREQUENCY_MONTHLY"
            }
        }

    }
    inspect_templates = ["projects/%{project}/inspectTemplates/${google_data_loss_prevention_inspect_template.basic.name}"] 
}

resource "google_pubsub_topic" "cloud_sql" {
    name = "fake-topic"
}

resource "google_data_loss_prevention_inspect_template" "basic" {
	parent = "projects/%{project}"
	description = "My description"
	display_name = "display_name"

	inspect_config {
		info_types {
			name = "EMAIL_ADDRESS"
		}
    }
}
`, context)
}

func testAccCheckDataLossPreventionDiscoveryConfigDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_data_loss_prevention_discovery_config" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DataLossPreventionBasePath}}{{parent}}/discoveryConfigs/{{name}}")
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
				return fmt.Errorf("DataLossPreventionDiscoveryConfig still exists at %s", url)
			}
		}

		return nil
	}
}
