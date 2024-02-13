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

package clouddeploy_test

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

func TestAccClouddeployCustomTargetType_clouddeployCustomTargetTypeBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckClouddeployCustomTargetTypeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccClouddeployCustomTargetType_clouddeployCustomTargetTypeBasicExample(context),
			},
			{
				ResourceName:            "google_clouddeploy_custom_target_type.custom-target-type",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "annotations", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccClouddeployCustomTargetType_clouddeployCustomTargetTypeBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_clouddeploy_custom_target_type" "custom-target-type" {
    location = "us-central1"
    name = "tf-test-my-custom-target-type%{random_suffix}"
    description = "My custom target type"
    annotations = {
      my_first_annotation = "example-annotation-1"
      my_second_annotation = "example-annotation-2"
    }
    labels = {
      my_first_label = "example-label-1"
      my_second_label = "example-label-2"
    }
    custom_actions {
      render_action = "renderAction"
      deploy_action = "deployAction"
    }
}
`, context)
}

func TestAccClouddeployCustomTargetType_clouddeployCustomTargetTypeGitSkaffoldModulesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckClouddeployCustomTargetTypeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccClouddeployCustomTargetType_clouddeployCustomTargetTypeGitSkaffoldModulesExample(context),
			},
			{
				ResourceName:            "google_clouddeploy_custom_target_type.custom-target-type",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "annotations", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccClouddeployCustomTargetType_clouddeployCustomTargetTypeGitSkaffoldModulesExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_clouddeploy_custom_target_type" "custom-target-type" {
    location = "us-central1"
    name = "tf-test-my-custom-target-type%{random_suffix}"
    description = "My custom target type"
    custom_actions {
      render_action = "renderAction"
      deploy_action = "deployAction"
      include_skaffold_modules {
        configs = ["my-config"]
        git {
            repo = "http://github.com/example/example-repo.git"
            path = "configs/skaffold.yaml"
            ref = "main"
        }
      }
    }
}
`, context)
}

func TestAccClouddeployCustomTargetType_clouddeployCustomTargetTypeGcsSkaffoldModulesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckClouddeployCustomTargetTypeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccClouddeployCustomTargetType_clouddeployCustomTargetTypeGcsSkaffoldModulesExample(context),
			},
			{
				ResourceName:            "google_clouddeploy_custom_target_type.custom-target-type",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "annotations", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccClouddeployCustomTargetType_clouddeployCustomTargetTypeGcsSkaffoldModulesExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_clouddeploy_custom_target_type" "custom-target-type" {
    location = "us-central1"
    name = "tf-test-my-custom-target-type%{random_suffix}"
    description = "My custom target type"
    custom_actions {
      render_action = "renderAction"
      deploy_action = "deployAction"
      include_skaffold_modules {
        configs = ["my-config"]
        google_cloud_storage {
            source = "gs://example-bucket/dir/configs/*"
            path = "skaffold.yaml"
        }
      }
    }
}
`, context)
}

func testAccCheckClouddeployCustomTargetTypeDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_clouddeploy_custom_target_type" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ClouddeployBasePath}}projects/{{project}}/locations/{{location}}/customTargetTypes/{{name}}")
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
				return fmt.Errorf("ClouddeployCustomTargetType still exists at %s", url)
			}
		}

		return nil
	}
}
