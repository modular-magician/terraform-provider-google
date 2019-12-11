// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Nathan is editing this to generate diffs in lots of files.
//     He won't submit this change.
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

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccBigtableAppProfile_bigtableAppProfileMulticlusterExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigtableAppProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBigtableAppProfile_bigtableAppProfileMulticlusterExample(context),
			},
			{
				ResourceName:            "google_bigtable_app_profile.ap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ignore_warnings"},
			},
		},
	})
}

func testAccBigtableAppProfile_bigtableAppProfileMulticlusterExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigtable_instance" "instance" {
  name = "tf-test-instance-%{random_suffix}"
  cluster {
    cluster_id   = "tf-test-instance-%{random_suffix}"
    zone         = "us-central1-b"
    num_nodes    = 3
    storage_type = "HDD"
  }
}

resource "google_bigtable_app_profile" "ap" {
  instance       = google_bigtable_instance.instance.name
  app_profile_id = "tf-test-profile-%{random_suffix}"

  multi_cluster_routing_use_any = true
  ignore_warnings               = true
}
`, context)
}

func TestAccBigtableAppProfile_bigtableAppProfileSingleclusterExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigtableAppProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBigtableAppProfile_bigtableAppProfileSingleclusterExample(context),
			},
			{
				ResourceName:            "google_bigtable_app_profile.ap",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ignore_warnings"},
			},
		},
	})
}

func testAccBigtableAppProfile_bigtableAppProfileSingleclusterExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigtable_instance" "instance" {
  name = "tf-test-instance-%{random_suffix}"
  cluster {
    cluster_id   = "tf-test-instance-%{random_suffix}"
    zone         = "us-central1-b"
    num_nodes    = 3
    storage_type = "HDD"
  }
}

resource "google_bigtable_app_profile" "ap" {
  instance       = google_bigtable_instance.instance.name
  app_profile_id = "tf-test-profile-%{random_suffix}"

  single_cluster_routing {
    cluster_id                 = "tf-test-instance-%{random_suffix}"
    allow_transactional_writes = true
  }

  ignore_warnings = true
}
`, context)
}

func testAccCheckBigtableAppProfileDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_bigtable_app_profile" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{BigtableBasePath}}projects/{{project}}/instances/{{instance}}/appProfiles/{{app_profile_id}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("BigtableAppProfile still exists at %s", url)
		}
	}

	return nil
}
