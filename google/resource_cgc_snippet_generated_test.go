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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCGCSnippet_sqlDatabaseInstanceSqlserverExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlDatabaseInstanceSqlserverExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection", "root_password"},
			},
		},
	})
}

func testAccCGCSnippet_sqlDatabaseInstanceSqlserverExample(context map[string]interface{}) string {
	return Nprintf(`
# [START cloud_sql_sqlserver_instance_80_db_n1_s2]
resource "google_sql_database_instance" "instance" {
  name             = "tf-test-sqlserver-instance%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2017_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  settings {
    tier = "db-custom-2-7680"
  }
  deletion_protection =  "%{deletion_protection}"
}
# [END cloud_sql_sqlserver_instance_80_db_n1_s2]
`, context)
}

func TestAccCGCSnippet_sqlDatabaseInstancePostgresExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlDatabaseInstancePostgresExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection"},
			},
		},
	})
}

func testAccCGCSnippet_sqlDatabaseInstancePostgresExample(context map[string]interface{}) string {
	return Nprintf(`
# [START cloud_sql_postgres_instance_80_db_n1_s2]
resource "google_sql_database_instance" "instance" {
  name             = "tf-test-postgres-instance%{random_suffix}"
  region           = "us-central1"
  database_version = "POSTGRES_12"
  settings {
    tier = "db-custom-2-7680"
  }
  deletion_protection =  "%{deletion_protection}"
}
# [END cloud_sql_postgres_instance_80_db_n1_s2]
`, context)
}

func TestAccCGCSnippet_sqlDatabaseInstanceMySqlExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlDatabaseInstanceMySqlExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection"},
			},
		},
	})
}

func testAccCGCSnippet_sqlDatabaseInstanceMySqlExample(context map[string]interface{}) string {
	return Nprintf(`
# [START cloud_sql_mysql_instance_80_db_n1_s2]
resource "google_sql_database_instance" "instance" {
  name             = "tf-test-mysql-instance%{random_suffix}"
  region           = "us-central1"
  database_version = "MYSQL_8_0"
  settings {
    tier = "db-n1-standard-2"
  }
  deletion_protection =  "%{deletion_protection}"
}
# [END cloud_sql_mysql_instance_80_db_n1_s2]
`, context)
}

func TestAccCGCSnippet_sqlMysqlInstanceBackupExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlMysqlInstanceBackupExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection"},
			},
		},
	})
}

func testAccCGCSnippet_sqlMysqlInstanceBackupExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "instance" {
  name             = "tf-test-mysql-instance-backup%{random_suffix}"
  region           = "asia-northeast1"
  database_version = "MYSQL_5_7"
  settings {
    tier = "db-f1-micro"
    backup_configuration {
      enabled                        = true
      binary_log_enabled             = true
      start_time                     = "20:55"
    }
  }
  deletion_protection =  "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_sqlPostgresInstanceBackupExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlPostgresInstanceBackupExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection"},
			},
		},
	})
}

func testAccCGCSnippet_sqlPostgresInstanceBackupExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "instance" {
  name             = "tf-test-postgres-instance-backup%{random_suffix}"
  region           = "us-central1"
  database_version = "POSTGRES_12"
  settings {
    tier = "db-custom-2-7680"
    backup_configuration {
      enabled                        = true
      start_time                     = "20:55"
    }
  }
  deletion_protection =  "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_sqlSqlserverInstanceBackupExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlSqlserverInstanceBackupExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection", "root_password"},
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverInstanceBackupExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "default" {
  name             = "tf-test-sqlserver-instance-backup%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2017_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  settings {
    tier = "db-custom-2-7680"
    backup_configuration {
      enabled                        = true
      start_time                     = "20:55"
    }
  }
  deletion_protection =  "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_sqlSqlserverInstanceBackupLocationExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlSqlserverInstanceBackupLocationExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"root_password", "deletion_protection"},
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverInstanceBackupLocationExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "default" {
  name             = ""
  region           = "us-central1"
  database_version = "SQLSERVER_2017_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  settings {
    tier = "db-custom-2-7680"
    backup_configuration {
      enabled                        = true
      location                       = "us-central1"
    }
  }
  deletion_protection =  "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_storageNewBucketExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_storageNewBucketExample(context),
			},
			{
				ResourceName:      "google_storage_bucket.static",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCGCSnippet_storageNewBucketExample(context map[string]interface{}) string {
	return Nprintf(`

# Create new storage bucket in the US region
# with coldline storage
resource "google_storage_bucket" "static" {
  name          = "tf-test-new-bucket%{random_suffix}"
  location      = "US"
  storage_class = "COLDLINE"

  uniform_bucket_level_access = true
}

`, context)
}
