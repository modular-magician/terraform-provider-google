// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccComputeSnapshot_snapshotBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeSnapshotDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSnapshot_snapshotBasicExample(context),
			},
			{
				ResourceName:            "google_compute_snapshot.snapshot",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"zone", "source_disk_encryption_key"},
			},
		},
	})
}

func testAccComputeSnapshot_snapshotBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_snapshot" "snapshot" {
	name = "my-snapshot%{random_suffix}"
	source_disk = "${google_compute_disk.persistent.name}"
	zone = "us-central1-a"
	labels = {
		my_label = "value"
	}
}

data "google_compute_image" "debian" {
	family  = "debian-9"
	project = "debian-cloud"
}

resource "google_compute_disk" "persistent" {
	name = "debian-disk%{random_suffix}"
	image = "${data.google_compute_image.debian.self_link}"
	size = 10
	type = "pd-ssd"
	zone = "us-central1-a"
}
`, context)
}

func testAccCheckComputeSnapshotDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_snapshot" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/snapshots/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeSnapshot still exists at %s", url)
		}
	}

	return nil
}
