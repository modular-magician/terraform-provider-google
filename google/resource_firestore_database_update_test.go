package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFirestoreDatabase_update(t *testing.T) {
	t.Parallel()

	dbName := "db-to-update"

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFirestoreDatabase_concurrencyMode(dbName, "OPTIMISTIC"),
			},
			{
				ResourceName:            "google_firestore_database.foobar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "delete_protection_state"},
			},
			{
				Config: testAccFirestoreDatabase_concurrencyMode(dbName, "PESSIMISTIC"),
			},
			{
				ResourceName:            "google_firestore_database.foobar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"etag", "delete_protection_state"},
			},
		},
	})
}

func testAccFirestoreDatabase_concurrencyMode(dbName, concurrencyMode string) string {
	return fmt.Sprintf(`
resource "google_firestore_database" "foobar" {
  name             = "%s"
  type             = "DATASTORE_MODE"
  location_id      = "nam5"
  concurrency_mode = "%s"
}
`, dbName, concurrencyMode)
}
