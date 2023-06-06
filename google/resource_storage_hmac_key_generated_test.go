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

package google

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

func TestAccStorageHmacKey_storageHmacKeyExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckStorageHmacKeyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccStorageHmacKey_storageHmacKeyExample(context),
			},
			{
				ResourceName:            "google_storage_hmac_key.key",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"secret"},
			},
		},
	})
}

func testAccStorageHmacKey_storageHmacKeyExample(context map[string]interface{}) string {
	return tpgresource.Nprintf(`
# Create a new service account
resource "google_service_account" "service_account" {
  account_id = "tf-test-my-svc-acc%{random_suffix}"
}

#Create the HMAC key for the associated service account 
resource "google_storage_hmac_key" "key" {
  service_account_email = google_service_account.service_account.email
}
`, context)
}

func testAccCheckStorageHmacKeyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_storage_hmac_key" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{StorageBasePath}}projects/{{project}}/hmacKeys/{{access_id}}")
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

			if v := res["state"]; v == "DELETED" {
				return nil
			}

			return fmt.Errorf("StorageHmacKey still exists at %s", url)
		}

		return nil
	}
}
