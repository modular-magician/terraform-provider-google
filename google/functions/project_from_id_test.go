// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package functions_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccProviderFunction_project_from_id(t *testing.T) {
	t.Parallel()

	projectId := envvar.GetTestProjectFromEnv()
	projectIdRegex := regexp.MustCompile(fmt.Sprintf("^%s$", projectId))

	validSelfLink := fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/zones/us-central1-c/instances/my-instance", projectId)
	validId := fmt.Sprintf("projects/%s/zones/us-central1-c/instances/my-instance", projectId)
	repetitiveInput := fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/%s/projects/not-this-1/projects/not-this-2/instances/my-instance", projectId)
	invalidInput := "zones/us-central1-c/instances/my-instance"

	context := map[string]interface{}{
		"function_name": "project_from_id",
		"output_name":   "project_id",
		"resource_name": fmt.Sprintf("tf-test-project-id-func-%s", acctest.RandString(t, 10)),
		"input":         "", // overridden in test cases
	}

	acctest.VcrTest(t, resource.TestCase{
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Given valid resource self_link input, the output value matches the expected value
				Config: testProviderFunction_generic_element_from_id(context, validSelfLink),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchOutput(context["output_name"].(string), projectIdRegex),
				),
			},
			{
				// Given valid resource id input, the output value matches the expected value
				Config: testProviderFunction_generic_element_from_id(context, validId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchOutput(context["output_name"].(string), projectIdRegex),
				),
			},
			{
				// Given repetitive input, the output value is the left-most match in the input
				Config: testProviderFunction_generic_element_from_id(context, repetitiveInput),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchOutput(context["output_name"].(string), projectIdRegex),
				),
			},
			{
				// Given invalid input, an error occurs
				Config:      testProviderFunction_generic_element_from_id(context, invalidInput),
				ExpectError: regexp.MustCompile("Error in function call"), // ExpectError doesn't inspect the specific error messages
			},
			{
				// Can get the project from a resource's id in one step
				// Uses google_pubsub_topic resource's id attribute with format projects/{{project}}/topics/{{name}}
				Config: testProviderFunction_get_project_from_resource_id(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchOutput(context["output_name"].(string), projectIdRegex),
				),
			},
			{
				// Can get the project from a resource's self_link in one step
				// Uses google_compute_subnetwork resource's self_link attribute
				Config: testProviderFunction_get_project_from_resource_self_link(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchOutput(context["output_name"].(string), projectIdRegex),
				),
			},
		},
	})
}

func testProviderFunction_generic_element_from_id(context map[string]interface{}, input string) string {
	context["input"] = input

	return acctest.Nprintf(`
	# terraform block required for provider function to be found
	terraform {
		required_providers {
			google = {
				source = "hashicorp/google"
			}
		}
	}

	output "%{output_name}" {
		value = provider::google::%{function_name}("%{input}")
	}
`, context)
}

func testProviderFunction_get_project_from_resource_id(context map[string]interface{}) string {
	return acctest.Nprintf(`
# terraform block required for provider function to be found
terraform {
	required_providers {
		google = {
			source = "hashicorp/google"
		}
	}
}

resource "google_pubsub_topic" "default" {
  name = "%{resource_name}"
}

output "%{output_name}" {
	value = provider::google::%{function_name}(google_pubsub_topic.default.id)
}
`, context)
}

func testProviderFunction_get_project_from_resource_self_link(context map[string]interface{}) string {
	return acctest.Nprintf(`
# terraform block required for provider function to be found
terraform {
	required_providers {
		google = {
			source = "hashicorp/google"
		}
	}
}

data "google_compute_network" "default" {
  name = "default"
}

resource "google_compute_subnetwork" "default" {
  name          = "%{resource_name}"
  ip_cidr_range = "10.2.0.0/16"
  network        = data.google_compute_network.default.id
}

output "%{output_name}" {
	value = provider::google::%{function_name}(google_compute_subnetwork.default.self_link)
}
`, context)
}
