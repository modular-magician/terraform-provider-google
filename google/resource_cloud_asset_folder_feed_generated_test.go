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

func TestAccCloudAssetFolderFeed_cloudAssetFolderFeedExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudAssetFolderFeedDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudAssetFolderFeed_cloudAssetFolderFeedExample(context),
			},
			{
				ResourceName:            "google_cloud_asset_folder_feed.folder_feed",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"billing_project", "feed_id", "folder"},
			},
		},
	})
}

func testAccCloudAssetFolderFeed_cloudAssetFolderFeedExample(context map[string]interface{}) string {
	return Nprintf(`
# Create a feed that sends notifications about network resource updates under a
# particular folder.
resource "google_cloud_asset_folder_feed" "folder_feed" {
  billing_project  = "%{project}"
  folder           = google_folder.my_folder.folder_id
  feed_id          = "tf-test-network-updates%{random_suffix}"
  content_type     = "RESOURCE"

  asset_types = [
    "compute.googleapis.com/Subnetwork",
    "compute.googleapis.com/Network",
  ]

  feed_output_config {
    pubsub_destination {
      topic = google_pubsub_topic.feed_output.id
    }
  }

  condition {
    expression = <<-EOT
    !temporal_asset.deleted &&
    temporal_asset.prior_asset_state == google.cloud.asset.v1.TemporalAsset.PriorAssetState.DOES_NOT_EXIST
    EOT
    title = "created"
    description = "Send notifications on creation events"
  }
}

# The topic where the resource change notifications will be sent.
resource "google_pubsub_topic" "feed_output" {
  project  = "%{project}"
  name     = "tf-test-network-updates%{random_suffix}"
}

# The folder that will be monitored for resource updates.
resource "google_folder" "my_folder" {
  display_name = "Networking%{random_suffix}"
  parent       = "organizations/%{org_id}"
}

# Find the project number of the project whose identity will be used for sending
# the asset change notifications.
data "google_project" "project" {
  project_id = "%{project}"
}
`, context)
}

func testAccCheckCloudAssetFolderFeedDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloud_asset_folder_feed" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{CloudAssetBasePath}}{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil)
			if err == nil {
				return fmt.Errorf("CloudAssetFolderFeed still exists at %s", url)
			}
		}

		return nil
	}
}
