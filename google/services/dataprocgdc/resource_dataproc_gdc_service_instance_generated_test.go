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

package dataprocgdc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccDataprocGdcServiceInstance_dataprocgdcServiceinstanceExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       "gdce-cluster-monitoring",
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataprocGdcServiceInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocGdcServiceInstance_dataprocgdcServiceinstanceExample(context),
			},
			{
				ResourceName:            "google_dataproc_gdc_service_instance.service-instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"labels", "location", "service_instance_id", "terraform_labels"},
			},
		},
	})
}

func testAccDataprocGdcServiceInstance_dataprocgdcServiceinstanceExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataproc_gdc_service_instance" "service-instance" {
  service_instance_id = "tf-test-tf-e2e-service-instance%{random_suffix}"
  project         = "%{project}"
  location        = "us-west2"
  gdce_cluster {
      gdce_cluster = "projects/gdce-cluster-monitoring/locations/us-west2/clusters/gdce-prism-prober-ord106"
  }
  display_name = "A service instance for a Terraform create test"
  labels = {
    "test-label": "label-value"
  }
  service_account = "dataprocgdc-cep-workflows@gdce-cluster-monitoring.iam.gserviceaccount.com"
}
`, context)
}

func testAccCheckDataprocGdcServiceInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dataproc_gdc_service_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DataprocGdcBasePath}}projects/{{project}}/locations/{{location}}/serviceInstances/{{service_instance_id}}")
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
				return fmt.Errorf("DataprocGdcServiceInstance still exists at %s", url)
			}
		}

		return nil
	}
}
