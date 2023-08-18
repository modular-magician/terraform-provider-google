// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package google

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/services/secretmanager"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSecretManagerSecret_import(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerSecretDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerSecret_basic(context),
			},
			{
				ResourceName:            "google_secret_manager_secret.secret-basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl"},
			},
		},
	})
}

func TestAccSecretManagerSecret_cmek(t *testing.T) {
	t.Parallel()

	kmscentral := acctest.BootstrapKMSKeyInLocation(t, "us-central1")
	kmseast := acctest.BootstrapKMSKeyInLocation(t, "us-east1")
	context1 := map[string]interface{}{
		"pid":                  envvar.GetTestProjectFromEnv(),
		"random_suffix":        acctest.RandString(t, 10),
		"kms_key_name_central": kmscentral.CryptoKey.Name,
		"kms_key_name_east":    kmseast.CryptoKey.Name,
	}
	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerSecretDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretMangerSecret_cmek(context1),
			},
			{
				ResourceName:            "google_secret_manager_secret.secret-basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl"},
			},
		},
	})
}

func TestAccSecretManagerSecret_annotationsUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerSecretDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretManagerSecret_annotationsBasic(context),
			},
			{
				ResourceName:            "google_secret_manager_secret.secret-with-annotations",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl"},
			},
			{
				Config: testAccSecretManagerSecret_annotationsUpdate(context),
			},
			{
				ResourceName:            "google_secret_manager_secret.secret-with-annotations",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl"},
			},
			{
				Config: testAccSecretManagerSecret_annotationsBasic(context),
			},
			{
				ResourceName:            "google_secret_manager_secret.secret-with-annotations",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl"},
			},
		},
	})
}

func TestAccSecretManagerSecret_automaticCmekUpdate(t *testing.T) {
	t.Parallel()

	suffix := acctest.RandString(t, 10)
	key1 := acctest.BootstrapKMSKeyWithPurposeInLocationAndName(t, "ENCRYPT_DECRYPT", "global", "tf-secret-manager-automatic-key1")
	key2 := acctest.BootstrapKMSKeyWithPurposeInLocationAndName(t, "ENCRYPT_DECRYPT", "global", "tf-secret-manager-automatic-key2")
	context := map[string]interface{}{
		"pid":            envvar.GetTestProjectFromEnv(),
		"random_suffix":  suffix,
		"kms_key_name_1": key1.CryptoKey.Name,
		"kms_key_name_2": key2.CryptoKey.Name,
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckSecretManagerSecretDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSecretMangerSecret_automaticCmekBasic(context),
			},
			{
				ResourceName:            "google_secret_manager_secret.secret-basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl"},
			},
			{
				Config: testAccSecretMangerSecret_automaticCmekUpdate(context),
			},
			{
				ResourceName:            "google_secret_manager_secret.secret-basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl"},
			},
			{
				Config: testAccSecretMangerSecret_automaticCmekUpdate2(context),
			},
			{
				ResourceName:            "google_secret_manager_secret.secret-basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl"},
			},
			{
				Config: testAccSecretMangerSecret_automaticCmekBasic(context),
			},
			{
				ResourceName:            "google_secret_manager_secret.secret-basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl"},
			},
		},
	})
}

func testAccSecretManagerSecret_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "tf-test-secret-%{random_suffix}"
  
  labels = {
    label = "my-label"
  }

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
      replicas {
        location = "us-east1"
      }
    }
  }

  ttl = "3600s"

}
`, context)
}

func testAccSecretMangerSecret_cmek(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
  project_id = "%{pid}"
}
resource "google_project_iam_member" "kms-secret-binding" {
  project = data.google_project.project.project_id
  role    = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member  = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-secretmanager.iam.gserviceaccount.com"
}
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "tf-test-secret-%{random_suffix}"
  
  labels = {
    label = "my-label"
  }
  replication {
    user_managed {
      replicas {
		location = "us-central1"
		customer_managed_encryption {
			kms_key_name = "%{kms_key_name_central}"
		}
	  }
	replicas {
		location = "us-east1"
		customer_managed_encryption {
			kms_key_name = "%{kms_key_name_east}"
		}
      }
	  
    }
  }
  project   = google_project_iam_member.kms-secret-binding.project
}
`, context)
}

func testAccSecretManagerSecret_annotationsBasic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_secret" "secret-with-annotations" {
  secret_id = "tf-test-secret-%{random_suffix}"

  labels = {
    label = "my-label"
  }

  annotations = {
    key1 = "someval"
    key2 = "someval2"
    key3 = "someval3"
    key4 = "someval4"
    key5 = "someval5"
  }

  replication {
    automatic {}
  }
}
`, context)
}

func testAccSecretManagerSecret_annotationsUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_secret_manager_secret" "secret-with-annotations" {
  secret_id = "tf-test-secret-%{random_suffix}"

  labels = {
    label = "my-label"
  }

  annotations = {
    key1 = "someval"
    key2update = "someval2"
    key3 = "someval3update"
    key4update = "someval4update"
  }

  replication {
    automatic {}
  }
}
`, context)
}

func testAccSecretMangerSecret_automaticCmekBasic(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
  project_id = "%{pid}"
}
resource "google_kms_crypto_key_iam_member" "kms-secret-binding-1" {
  crypto_key_id = "%{kms_key_name_1}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-secretmanager.iam.gserviceaccount.com"
}
resource "google_kms_crypto_key_iam_member" "kms-secret-binding-2" {
  crypto_key_id = "%{kms_key_name_2}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-secretmanager.iam.gserviceaccount.com"
}
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "tf-test-secret-%{random_suffix}"
  
  labels = {
    label = "my-label"
  }
  replication {
    automatic {}
  }
  depends_on = [
    google_kms_crypto_key_iam_member.kms-secret-binding-1,
    google_kms_crypto_key_iam_member.kms-secret-binding-2,
  ]
}
`, context)
}

func testAccSecretMangerSecret_automaticCmekUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
  project_id = "%{pid}"
}
resource "google_kms_crypto_key_iam_member" "kms-secret-binding-1" {
  crypto_key_id = "%{kms_key_name_1}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-secretmanager.iam.gserviceaccount.com"
}
resource "google_kms_crypto_key_iam_member" "kms-secret-binding-2" {
  crypto_key_id = "%{kms_key_name_2}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-secretmanager.iam.gserviceaccount.com"
}
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "tf-test-secret-%{random_suffix}"
  
  labels = {
    label = "my-label"
  }
  replication {
    automatic {
      customer_managed_encryption {
        kms_key_name = "%{kms_key_name_1}"
      }
    }
  }
  depends_on = [
    google_kms_crypto_key_iam_member.kms-secret-binding-1,
    google_kms_crypto_key_iam_member.kms-secret-binding-2,
  ]
}
`, context)
}

func testAccSecretMangerSecret_automaticCmekUpdate2(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
  project_id = "%{pid}"
}
resource "google_kms_crypto_key_iam_member" "kms-secret-binding-1" {
  crypto_key_id = "%{kms_key_name_1}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-secretmanager.iam.gserviceaccount.com"
}
resource "google_kms_crypto_key_iam_member" "kms-secret-binding-2" {
  crypto_key_id = "%{kms_key_name_2}"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-secretmanager.iam.gserviceaccount.com"
}
resource "google_secret_manager_secret" "secret-basic" {
  secret_id = "tf-test-secret-%{random_suffix}"
  
  labels = {
    label = "my-label"
  }
  replication {
    automatic {
      customer_managed_encryption {
        kms_key_name = "%{kms_key_name_2}"
      }
    }
  }
  depends_on = [
    google_kms_crypto_key_iam_member.kms-secret-binding-1,
    google_kms_crypto_key_iam_member.kms-secret-binding-2,
  ]
}
`, context)
}

func testSecretManagerSecretReplicationWithAutomaticV0() map[string]any {
	return map[string]any{
		"replication": []interface{}{
			map[string]interface{}{
				"automatic":    true,
				"user_managed": []interface{}{},
			},
		},
	}
}

func testSecretManagerSecretReplicationWithAutomaticV1() map[string]any {
	return map[string]any{
		"replication": []interface{}{
			map[string]interface{}{
				"automatic":    []interface{}{},
				"user_managed": []interface{}{},
			},
		},
	}
}

func TestSecretManagerSecretReplicationWithAutomaticStateUpgradeV0(t *testing.T) {
	expected := testSecretManagerSecretReplicationWithAutomaticV1()
	actual, err := secretmanager.ResourceSecretManagerSecretUpgradeV0(context.Background(), testSecretManagerSecretReplicationWithAutomaticV0(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
	}
}

func testSecretManagerSecretReplicationWithoutAutomaticV0() map[string]any {
	return map[string]any{
		"replication": []interface{}{
			map[string]interface{}{
				"automatic": false,
				"user_managed": []interface{}{
					map[string]interface{}{
						"location": "us-central1",
					},
					map[string]interface{}{
						"location": "us-east1",
					},
				},
			},
		},
	}
}

func testSecretManagerSecretReplicationWithoutAutomaticV1() map[string]any {
	return map[string]any{
		"replication": []interface{}{
			map[string]interface{}{
				"automatic": []interface{}{},
				"user_managed": []interface{}{
					map[string]interface{}{
						"location": "us-central1",
					},
					map[string]interface{}{
						"location": "us-east1",
					},
				},
			},
		},
	}
}

func TestSecretManagerSecretReplicationWithoutAutomaticStateUpgradeV0(t *testing.T) {
	expected := testSecretManagerSecretReplicationWithoutAutomaticV1()
	actual, err := secretmanager.ResourceSecretManagerSecretUpgradeV0(context.Background(), testSecretManagerSecretReplicationWithoutAutomaticV0(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
	}
}
