package google

import (
	"testing"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccTPUNode_tpuNodeBUpdateTensorFlowVersion(t *testing.T) {
	t.Parallel()

	nodeId := acctest.RandomWithPrefix("tf-test")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTPUNodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTpuNode_tpuNodeTensorFlow(nodeId, 0),
			},
			{
				ResourceName:            "google_tpu_node.tpu",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"zone"},
			},
			{
				Config: testAccTpuNode_tpuNodeTensorFlow(nodeId, 1),
			},
			{
				ResourceName:            "google_tpu_node.tpu",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"zone"},
			},
		},
	})
}

// WARNING: cidr_block must not overlap with other existing TPU blocks
// Make sure if you change this value that it does not overlap with the
// autogenerated examples.
func testAccTpuNode_tpuNodeTensorFlow(nodeId string, versionIdx int) string {
	return fmt.Sprintf(`
data "google_tpu_tensorflow_versions" "available" {
}

resource "google_tpu_node" "tpu" {
  name = "%s"
  zone = "us-central1-b"

  accelerator_type   = "v3-8"
  tensorflow_version = data.google_tpu_tensorflow_versions.available.versions[%d]
  cidr_block         = "10.1.0.0/29"
}
`, nodeId, versionIdx)
}
