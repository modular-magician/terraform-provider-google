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

package corebilling_test

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

func TestAccCoreBillingProjectInfo_billingProjectInfoBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckCoreBillingProjectInfoDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCoreBillingProjectInfo_billingProjectInfoBasicExample(context),
			},
		},
	})
}

func testAccCoreBillingProjectInfo_billingProjectInfoBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
  deletion_policy = "NONE"
  lifecycle {
    ignore_changes = [billing_account]
  }
}

resource "google_billing_project_info" "default" {
  project         = google_project.project.project_id
  billing_account = "%{billing_account}"
}
`, context)
}

func testAccCheckCoreBillingProjectInfoDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_billing_project_info" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			// After deleting a project, you can still query its billing account
			// (it will be empty). We change the destroy check to ensure the
			// project has no billing account linked after destroying the
			// google_billing_project_info resource

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{CoreBillingBasePath}}projects/{{project}}/billingInfo")
			if err != nil {
				return err
			}

			res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err != nil {
				return nil
			}

			if ba := res["billingAccountName"]; ba == "" {
				return nil
			}

			return fmt.Errorf("Billing account still linked at %s", url)
		}

		return nil
	}
}
