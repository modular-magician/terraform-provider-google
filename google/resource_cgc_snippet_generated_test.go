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
      image = "debian-cloud/debian-11"
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

func TestAccCGCSnippet_spotInstanceBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_spotInstanceBasicExample(context),
			},
			{
				ResourceName:      "google_compute_instance.spot_vm_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCGCSnippet_spotInstanceBasicExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_compute_instance" "spot_vm_instance" {
  name         = "tf-test-spot-instance-name%{random_suffix}"
  machine_type = "f1-micro"
  zone         = "us-central1-c"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }
  
  scheduling {
      preemptible = true
      automatic_restart = false
      provisioning_model = "SPOT"
      instance_termination_action = "STOP"
  }

  network_interface {
    # A default network is created for all GCP projects
    network = "default"
    access_config {
    }
  }
}

`, context)
}

func TestAccCGCSnippet_instanceCustomHostnameExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_instanceCustomHostnameExample(context),
			},
			{
				ResourceName:      "google_compute_instance.custom_hostname_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCGCSnippet_instanceCustomHostnameExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_compute_instance" "custom_hostname_instance" {
  name         = "tf-test-custom-hostname-instance-name%{random_suffix}"
  machine_type = "f1-micro"
  zone = "us-central1-c"

  # Set a custom hostname below 
  hostname = "hashicorptest.com"
  
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }
  network_interface {
    # A default network is created for all GCP projects
    network = "default"
    access_config {
    }
  }
}

`, context)
}

func TestAccCGCSnippet_computeReservationExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       getTestProjectFromEnv(),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_computeReservationExample(context),
			},
		},
	})
}

func testAccCGCSnippet_computeReservationExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_compute_reservation" "gce_reservation_local" {
  name = "tf-test-gce-reservation-local%{random_suffix}"
  zone = "us-central1-c"
  project = "%{project}"

  share_settings {
    share_type = "LOCAL"
  }

  specific_reservation {
    count = 1
    instance_properties {
      machine_type     = "n2-standard-2"
    }
  }
}

`, context)
}

func TestAccCGCSnippet_instanceVirtualDisplayEnabledExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_instanceVirtualDisplayEnabledExample(context),
			},
			{
				ResourceName:      "google_compute_instance.instance_virtual_display",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCGCSnippet_instanceVirtualDisplayEnabledExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_compute_instance" "instance_virtual_display" {
  name         = "tf-test-instance-virtual-display%{random_suffix}"
  machine_type = "f1-micro"
  zone = "us-central1-c"
  
  # Set the below to true to enable virtual display
  enable_display = true
  
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }
  network_interface {
    # A default network is created for all GCP projects
    network = "default"
    access_config {
    }
  }
}

`, context)
}

func TestAccCGCSnippet_osLoginExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_osLoginExample(context),
			},
			{
				ResourceName:      "google_compute_instance.oslogin_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCGCSnippet_osLoginExample(context map[string]interface{}) string {
	return Nprintf(`
# [START compute_terraform-oslogin-example]

resource "google_project_service" "project" {
  service            = "oslogin.googleapis.com"
  disable_on_destroy = false
}

resource "google_compute_project_metadata" "default" {
  metadata = {
    enable-oslogin = "TRUE"
  }
}

resource "google_compute_instance" "oslogin_instance" {
  name         = "tf-test-oslogin-instance-name%{random_suffix}"
  machine_type = "f1-micro"
  zone         = "us-central1-c"
  metadata = {
    enable-oslogin: "TRUE"
  }
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }
  network_interface {
    # A default network is created for all GCP projects
    network = "default"
    access_config {
    }
  }
}


data "google_project" "project" {
}
resource "google_project_iam_member" "os-login-admin-users" {
  project  = data.google_project.project.project_id
  role = "roles/compute.osAdminLogin"
  member   = "serviceAccount:service-${data.google_project.project.number}@compute-system.iam.gserviceaccount.com"
}

# [END compute_terraform-oslogin-example]
`, context)
}

func TestAccCGCSnippet_sqlDatabaseInstanceSqlserverExample(t *testing.T) {
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"deletion_protection": false,
		"random_suffix":       randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
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
  database_version = "SQLSERVER_2019_STANDARD"
  root_password = "INSERT-PASSWORD-HERE"
  settings {
    tier = "db-custom-2-7680"
  }
  deletion_protection =  "%{deletion_protection}"
}
# [END cloud_sql_sqlserver_instance_80_db_n1_s2]

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    name = "user"
    instance = google_sql_database_instance.instance.name
    password = random_password.pwd.result
}
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
  database_version = "SQLSERVER_2019_STANDARD"
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
  database_version = "SQLSERVER_2019_STANDARD"
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
  database_version = "SQLSERVER_2019_STANDARD"
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
  database_version = "SQLSERVER_2019_STANDARD"
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

# Get object metadata
data "google_storage_bucket_object" "default" {
  name         = google_storage_bucket_object.default.name
  bucket       = google_storage_bucket.static.id
}

output "object_metadata" {
  value        = data.google_storage_bucket_object.default
}

# Get bucket metadata
data "google_storage_bucket" "default" {
  name         = google_storage_bucket.static.id
}

output "bucket_metadata" {
  value        = data.google_storage_bucket.default
}
`, context)
}

func TestAccCGCSnippet_storageStaticWebsiteExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCGCSnippet_storageStaticWebsiteExample(context),
			},
			{
				ResourceName:      "google_storage_bucket.static_website",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCGCSnippet_storageStaticWebsiteExample(context map[string]interface{}) string {
	return Nprintf(`
# Create new storage bucket in the US multi-region
# with coldline storage and settings for main_page_suffix and not_found_page
resource "google_storage_bucket" "static_website" {
    name          = "tf-test-static-website-bucket%{random_suffix}"
    location      = "US"
    storage_class = "COLDLINE"
    website {
        main_page_suffix = "index.html%{random_suffix}"
        not_found_page = "index.html%{random_suffix}"
    }
}

# Make bucket public by granting allUsers READER access
resource "google_storage_bucket_access_control" "public_rule" {
  bucket = google_storage_bucket.static_website.id
  role   = "READER"
  entity = "allUsers"
}

# Upload a simple index.html page to the bucket
resource "google_storage_bucket_object" "indexpage" {
  name         = "index.html%{random_suffix}"
  content      = "<html><body>Hello World!</body></html>"
  content_type = "text/html"
  bucket       = google_storage_bucket.static_website.id
}

# Upload a simple 404 / error page to the bucket
resource "google_storage_bucket_object" "errorpage" {
  name         = "404.html%{random_suffix}"
  content      = "<html><body>404!</body></html>"
  content_type = "text/html"
  bucket       = google_storage_bucket.static_website.id
}
`, context)
}
