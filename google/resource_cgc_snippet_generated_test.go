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

func TestAccCGCSnippet_flaskGoogleCloudQuickstartExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_flaskGoogleCloudQuickstartExample(context),
			},
			{
				ResourceName:            "google_compute_instance.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata", "metadata_startup_script"},
			},
		},
	})
}

func testAccCGCSnippet_flaskGoogleCloudQuickstartExample(context map[string]interface{}) string {
	return Nprintf(`
# Create a single Compute Engine instance
resource "google_compute_instance" "default" {
  name         = "tf-test-flask-vm%{random_suffix}"
  machine_type = "f1-micro"
  zone         = "us-west1-a"
  tags         = ["ssh"]

  metadata = {
    enable-oslogin = "TRUE"
  }
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  # Install Flask
  metadata_startup_script = "sudo apt-get update; sudo apt-get install -yq build-essential python-pip rsync; pip install flask"

  network_interface {
    network = "default"

    access_config {
      # Include this section to give the VM an external IP address
    }
  }
}

resource "google_compute_firewall" "ssh" {
  name = "tf-test-allow-ssh%{random_suffix}"
  allow {
    ports    = ["22"]
    protocol = "tcp"
  }
  direction     = "INGRESS"
  network       = "default"
  priority      = 1000
  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["ssh"]
}


# [START vpc_flask_quickstart_5000_fw]
resource "google_compute_firewall" "flask" {
  name    = "tf-test-flask-app-firewall%{random_suffix}"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["5000"]
  }
  source_ranges = ["0.0.0.0/0"]
}
# [END vpc_flask_quickstart_5000_fw]

# Create new multi-region storage bucket in the US
# with versioning enabled

resource "google_storage_bucket" "default" {
  name          = "tf-test-bucket-tfstate%{random_suffix}"
  force_destroy = false
  location      = "US"
  storage_class = "STANDARD"
  versioning {
    enabled = true
  }
}
`, context)
}

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

func TestAccCGCSnippet_sqlSqlserverInstanceCloneExample(t *testing.T) {
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
				Config: testAccCGCSnippet_sqlSqlserverInstanceCloneExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.clone",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection", "root_password", "clone"},
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverInstanceCloneExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "source" {
  name             = "tf-test-sqlserver-instance-source-name%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2017_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  settings {
    tier = "db-custom-2-7680"
  }
  deletion_protection =  "%{deletion_protection}"
}

resource "google_sql_database_instance" "clone" {
  name             = "tf-test-sqlserver-instance-clone-name%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2017_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  clone {
    source_instance_name = google_sql_database_instance.source.id
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

func TestAccCGCSnippet_sqlSqlserverInstanceAuthorizedNetworkExample(t *testing.T) {
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
				Config: testAccCGCSnippet_sqlSqlserverInstanceAuthorizedNetworkExample(context),
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

func testAccCGCSnippet_sqlSqlserverInstanceAuthorizedNetworkExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "default" {
  name = "tf-test-sqlserver-instance-with-authorized-network%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2017_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  settings {
    tier = "db-custom-2-7680"
    ip_configuration {
      authorized_networks {
        name = "Network Name"
        value = "192.0.2.0/24"
        expiration_time = "3021-11-15T16:19:00.094Z"
      }
    }
  }
  deletion_protection = "%{deletion_protection}"
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
  name             = "tf-test-sqlserver-instance-with-backup-location%{random_suffix}"
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

func TestAccCGCSnippet_sqlSqlserverInstanceBackupRetentionExample(t *testing.T) {
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
				Config: testAccCGCSnippet_sqlSqlserverInstanceBackupRetentionExample(context),
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

func testAccCGCSnippet_sqlSqlserverInstanceBackupRetentionExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "default" {
  name             = "tf-test-sqlserver-instance-backup-retention%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2017_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  settings {
    tier = "db-custom-2-7680"
    backup_configuration {
      enabled                        = true
      backup_retention_settings {
        retained_backups               = 365
        retention_unit                 = "COUNT"
      }
    }
  }
  deletion_protection =  "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_sqlSqlserverInstanceReplicaExample(t *testing.T) {
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
				Config: testAccCGCSnippet_sqlSqlserverInstanceReplicaExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.read_replica",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection", "root_password"},
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverInstanceReplicaExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "primary" {
  name             = "tf-test-sqlserver-primary-instance-name%{random_suffix}"
  region           = "europe-west4"
  database_version = "SQLSERVER_2019_ENTERPRISE"
  root_password    = "INSERT-PASSWORD-HERE"
  settings {
    tier = "db-custom-2-7680"
    backup_configuration {
      enabled = "true"
    }
  }
  deletion_protection = "%{deletion_protection}"
}

resource "google_sql_database_instance" "read_replica" {
  name                 = "tf-test-sqlserver-replica-instance-name%{random_suffix}"
  master_instance_name = google_sql_database_instance.primary.name
  region               = "europe-west4"
  database_version     = "SQLSERVER_2019_ENTERPRISE"
  root_password        = "INSERT-PASSWORD-HERE"
  replica_configuration {
    failover_target = false
  }

  settings {
    tier              = "db-custom-2-7680"
    availability_type = "ZONAL"
    disk_size         = "100"
  }
  deletion_protection = "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_sqlSqlserverInstancePublicIpExample(t *testing.T) {
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
				Config: testAccCGCSnippet_sqlSqlserverInstancePublicIpExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.sqlserver_public_ip_instance_name",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_protection", "root_password"},
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverInstancePublicIpExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "sqlserver_public_ip_instance_name" {
  name                 = "tf-test-sqlserver-public-ip-instance-name%{random_suffix}"
  region               = "europe-west4"
  database_version     = "SQLSERVER_2019_ENTERPRISE"
  root_password        = "INSERT-PASSWORD-HERE"
  settings {
    tier              = "db-custom-2-7680"
    availability_type = "ZONAL"
    ip_configuration {
      # Add optional authorized networks
      # Update to match the customer's networks
      authorized_networks {
        name  = "test-net-3"
        value = "203.0.113.0/24"
      }
      # Enable public IP
      ipv4_enabled = true
    }
  }
  deletion_protection =  "%{deletion_protection}"
}
`, context)
}

func TestAccCGCSnippet_sqlSqlserverInstancePrivateIpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_sqlSqlserverInstancePrivateIpExample(context),
			},
			{
				ResourceName:            "google_sql_database_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"root_password", "deletion_protection"},
			},
		},
	})
}

func testAccCGCSnippet_sqlSqlserverInstancePrivateIpExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_compute_network" "private_network" {
  name                    = "tf-test-private-network%{random_suffix}"
  auto_create_subnetworks = "false"
}

resource "google_compute_global_address" "private_ip_address" {
  name          = "tf-test-private-ip-address%{random_suffix}"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.private_network.id
}

resource "google_service_networking_connection" "private_vpc_connection" {
  network                 = google_compute_network.private_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_address.name]
}

resource "google_sql_database_instance" "instance" {
  name             = "tf-test-private-ip-sql-instance%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2019_STANDARD"
  root_password        = "INSERT-PASSWORD-HERE"

  depends_on = [google_service_networking_connection.private_vpc_connection]

  settings {
    tier = "db-custom-2-7680"
    ip_configuration {
      ipv4_enabled    = "false"
      private_network = google_compute_network.private_network.id
    }
  }
  deletion_protection = "false"
}

`, context)
}

func TestAccCGCSnippet_sqlSqlserverInstanceFlagsExample(t *testing.T) {
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
				Config: testAccCGCSnippet_sqlSqlserverInstanceFlagsExample(context),
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

func testAccCGCSnippet_sqlSqlserverInstanceFlagsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "instance" {
  name             = "tf-test-sqlserver-instance%{random_suffix}"
  region           = "us-central1"
  database_version = "SQLSERVER_2019_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  settings {
    database_flags {
      name  = "1204"
      value = "on"
    }
    database_flags {
      name  = "remote access"
      value = "on"
    }
    database_flags {
      name  = "remote query timeout (s)"
      value = "300"
    }
    tier = "db-custom-2-7680"
  }
  deletion_protection = "%{deletion_protection}"
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
# Create new storage bucket in the US multi-region
# with coldline storage
resource "google_storage_bucket" "static" {
  name          = "tf-test-new-bucket%{random_suffix}"
  location      = "US"
  storage_class = "COLDLINE"

  uniform_bucket_level_access = true
} 

# Upload files
# Discussion about using tf to upload a large number of objects
# https://stackoverflow.com/questions/68455132/terraform-copy-multiple-files-to-bucket-at-the-same-time-bucket-creation

# The text object in Cloud Storage
resource "google_storage_bucket_object" "default" {
  name         = "tf-test-new-object%{random_suffix}"
# Uncomment and add valid path to an object.
#  source       = "/path/to/an/object"
  content      = "Data as string to be uploaded"
  content_type = "text/plain"
  bucket       = google_storage_bucket.static.id
}
`, context)
}
