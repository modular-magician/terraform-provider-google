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
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccFilestoreInstance_filestoreInstanceBasicExample(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFilestoreInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFilestoreInstance_filestoreInstanceBasicExample(acctest.RandString(10)),
			},
			{
				ResourceName:      "google_filestore_instance.instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccFilestoreInstance_filestoreInstanceBasicExample(val string) string {
	return fmt.Sprintf(`
resource "google_filestore_instance" "instance" {
  name = "test-instance-%s"
  zone = "us-central1-b"
  file_shares {
    capacity_gb = 2660
    name = "share1"
  }
  networks {
    network = "default"
    modes = ["MODE_IPV4"]
  }
  tier = "PREMIUM"
}
`, val,
	)
}
