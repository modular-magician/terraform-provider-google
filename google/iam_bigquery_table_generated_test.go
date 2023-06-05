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

	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccBigQueryTableIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/bigquery.dataOwner",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_table_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s roles/bigquery.dataOwner", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccBigQueryTableIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_table_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s roles/bigquery.dataOwner", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTableIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/bigquery.dataOwner",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccBigQueryTableIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_table_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s roles/bigquery.dataOwner user:admin@hashicorptest.com", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTableIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/bigquery.dataOwner",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_table_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBigQueryTableIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_bigquery_table_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTableIamBindingGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/bigquery.dataOwner",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableIamBinding_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_table_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s roles/bigquery.dataOwner %s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTableIamBindingGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/bigquery.dataOwner",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableIamBinding_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_table_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s roles/bigquery.dataOwner", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_bigquery_table_iam_binding.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s roles/bigquery.dataOwner %s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_bigquery_table_iam_binding.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s roles/bigquery.dataOwner %s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTableIamMemberGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/bigquery.dataOwner",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableIamMember_withConditionGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_table_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s roles/bigquery.dataOwner user:admin@hashicorptest.com %s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTableIamMemberGenerated_withAndWithoutCondition(t *testing.T) {
	// Multiple fine-grained resources
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/bigquery.dataOwner",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableIamMember_withAndWithoutConditionGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_table_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s roles/bigquery.dataOwner user:admin@hashicorptest.com", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_bigquery_table_iam_member.foo2",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s roles/bigquery.dataOwner user:admin@hashicorptest.com %s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"]), context["condition_title"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "google_bigquery_table_iam_member.foo3",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s roles/bigquery.dataOwner user:admin@hashicorptest.com %s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"]), context["condition_title_no_desc"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTableIamPolicyGenerated_withCondition(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":           RandString(t, 10),
		"role":                    "roles/bigquery.dataOwner",
		"condition_title":         "expires_after_2019_12_31",
		"condition_expr":          `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
		"condition_desc":          "Expiring at midnight of 2019-12-31",
		"condition_title_no_desc": "expires_after_2019_12_31-no-description",
		"condition_expr_no_desc":  `request.time < timestamp(\"2020-01-01T00:00:00Z\")`,
	}

	// Test should have 2 bindings: one with a description and one without. Any < chars are converted to a unicode character by the API.
	expectedPolicyData := Nprintf(`{"bindings":[{"condition":{"description":"%{condition_desc}","expression":"%{condition_expr}","title":"%{condition_title}"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"},{"condition":{"expression":"%{condition_expr}","title":"%{condition_title}-no-description"},"members":["user:admin@hashicorptest.com"],"role":"%{role}"}]}`, context)
	expectedPolicyData = strings.Replace(expectedPolicyData, "<", "\\u003c", -1)

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableIamPolicy_withConditionGenerated(context),
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO(SarahFrench) - uncomment once https://github.com/GoogleCloudPlatform/magic-modules/pull/6466 merged
					// resource.TestCheckResourceAttr("data.google_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttr("google_bigquery_table_iam_policy.foo", "policy_data", expectedPolicyData),
					resource.TestCheckResourceAttrWith("data.google_iam_policy.foo", "policy_data", checkGoogleIamPolicy),
				),
			},
			{
				ResourceName:      "google_bigquery_table_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s", acctest.GetTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBigQueryTableIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_table" "test" {
  deletion_protection = false
  table_id   = "tf_test_table_id%{random_suffix}"
  dataset_id = google_bigquery_dataset.test.dataset_id
  time_partitioning {
    type = "DAY"
  }
  schema = <<EOH
[
  {
    "name": "city",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "coord",
    "type": "RECORD",
    "fields": [
    {
      "name": "lon",
      "type": "FLOAT"
    },
    {
      "name": "lat",
      "type": "FLOAT"
    }
    ]
  }
    ]
  },
  {
    "name": "country",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "name",
    "type": "STRING"
  }
    ]
  }
]
EOH
}

resource "google_bigquery_table_iam_member" "foo" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccBigQueryTableIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_table" "test" {
  deletion_protection = false
  table_id   = "tf_test_table_id%{random_suffix}"
  dataset_id = google_bigquery_dataset.test.dataset_id
  time_partitioning {
    type = "DAY"
  }
  schema = <<EOH
[
  {
    "name": "city",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "coord",
    "type": "RECORD",
    "fields": [
    {
      "name": "lon",
      "type": "FLOAT"
    },
    {
      "name": "lat",
      "type": "FLOAT"
    }
    ]
  }
    ]
  },
  {
    "name": "country",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "name",
    "type": "STRING"
  }
    ]
  }
]
EOH
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_bigquery_table_iam_policy" "foo" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccBigQueryTableIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_table" "test" {
  deletion_protection = false
  table_id   = "tf_test_table_id%{random_suffix}"
  dataset_id = google_bigquery_dataset.test.dataset_id
  time_partitioning {
    type = "DAY"
  }
  schema = <<EOH
[
  {
    "name": "city",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "coord",
    "type": "RECORD",
    "fields": [
    {
      "name": "lon",
      "type": "FLOAT"
    },
    {
      "name": "lat",
      "type": "FLOAT"
    }
    ]
  }
    ]
  },
  {
    "name": "country",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "name",
    "type": "STRING"
  }
    ]
  }
]
EOH
}

data "google_iam_policy" "foo" {
}

resource "google_bigquery_table_iam_policy" "foo" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccBigQueryTableIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_table" "test" {
  deletion_protection = false
  table_id   = "tf_test_table_id%{random_suffix}"
  dataset_id = google_bigquery_dataset.test.dataset_id
  time_partitioning {
    type = "DAY"
  }
  schema = <<EOH
[
  {
    "name": "city",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "coord",
    "type": "RECORD",
    "fields": [
    {
      "name": "lon",
      "type": "FLOAT"
    },
    {
      "name": "lat",
      "type": "FLOAT"
    }
    ]
  }
    ]
  },
  {
    "name": "country",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "name",
    "type": "STRING"
  }
    ]
  }
]
EOH
}

resource "google_bigquery_table_iam_binding" "foo" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccBigQueryTableIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_table" "test" {
  deletion_protection = false
  table_id   = "tf_test_table_id%{random_suffix}"
  dataset_id = google_bigquery_dataset.test.dataset_id
  time_partitioning {
    type = "DAY"
  }
  schema = <<EOH
[
  {
    "name": "city",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "coord",
    "type": "RECORD",
    "fields": [
    {
      "name": "lon",
      "type": "FLOAT"
    },
    {
      "name": "lat",
      "type": "FLOAT"
    }
    ]
  }
    ]
  },
  {
    "name": "country",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "name",
    "type": "STRING"
  }
    ]
  }
]
EOH
}

resource "google_bigquery_table_iam_binding" "foo" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}

func testAccBigQueryTableIamBinding_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_table" "test" {
  deletion_protection = false
  table_id   = "tf_test_table_id%{random_suffix}"
  dataset_id = google_bigquery_dataset.test.dataset_id
  time_partitioning {
    type = "DAY"
  }
  schema = <<EOH
[
  {
    "name": "city",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "coord",
    "type": "RECORD",
    "fields": [
    {
      "name": "lon",
      "type": "FLOAT"
    },
    {
      "name": "lat",
      "type": "FLOAT"
    }
    ]
  }
    ]
  },
  {
    "name": "country",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "name",
    "type": "STRING"
  }
    ]
  }
]
EOH
}

resource "google_bigquery_table_iam_binding" "foo" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccBigQueryTableIamBinding_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_table" "test" {
  deletion_protection = false
  table_id   = "tf_test_table_id%{random_suffix}"
  dataset_id = google_bigquery_dataset.test.dataset_id
  time_partitioning {
    type = "DAY"
  }
  schema = <<EOH
[
  {
    "name": "city",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "coord",
    "type": "RECORD",
    "fields": [
    {
      "name": "lon",
      "type": "FLOAT"
    },
    {
      "name": "lat",
      "type": "FLOAT"
    }
    ]
  }
    ]
  },
  {
    "name": "country",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "name",
    "type": "STRING"
  }
    ]
  }
]
EOH
}

resource "google_bigquery_table_iam_binding" "foo" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}

resource "google_bigquery_table_iam_binding" "foo2" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_bigquery_table_iam_binding" "foo3" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccBigQueryTableIamMember_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_table" "test" {
  deletion_protection = false
  table_id   = "tf_test_table_id%{random_suffix}"
  dataset_id = google_bigquery_dataset.test.dataset_id
  time_partitioning {
    type = "DAY"
  }
  schema = <<EOH
[
  {
    "name": "city",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "coord",
    "type": "RECORD",
    "fields": [
    {
      "name": "lon",
      "type": "FLOAT"
    },
    {
      "name": "lat",
      "type": "FLOAT"
    }
    ]
  }
    ]
  },
  {
    "name": "country",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "name",
    "type": "STRING"
  }
    ]
  }
]
EOH
}

resource "google_bigquery_table_iam_member" "foo" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}
`, context)
}

func testAccBigQueryTableIamMember_withAndWithoutConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_table" "test" {
  deletion_protection = false
  table_id   = "tf_test_table_id%{random_suffix}"
  dataset_id = google_bigquery_dataset.test.dataset_id
  time_partitioning {
    type = "DAY"
  }
  schema = <<EOH
[
  {
    "name": "city",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "coord",
    "type": "RECORD",
    "fields": [
    {
      "name": "lon",
      "type": "FLOAT"
    },
    {
      "name": "lat",
      "type": "FLOAT"
    }
    ]
  }
    ]
  },
  {
    "name": "country",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "name",
    "type": "STRING"
  }
    ]
  }
]
EOH
}

resource "google_bigquery_table_iam_member" "foo" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}

resource "google_bigquery_table_iam_member" "foo2" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    title       = "%{condition_title}"
    description = "%{condition_desc}"
    expression  = "%{condition_expr}"
  }
}

resource "google_bigquery_table_iam_member" "foo3" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
  condition {
    # Check that lack of description doesn't cause any issues
    # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
    title       = "%{condition_title_no_desc}"
    expression  = "%{condition_expr_no_desc}"
  }
}
`, context)
}

func testAccBigQueryTableIamPolicy_withConditionGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
  dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_table" "test" {
  deletion_protection = false
  table_id   = "tf_test_table_id%{random_suffix}"
  dataset_id = google_bigquery_dataset.test.dataset_id
  time_partitioning {
    type = "DAY"
  }
  schema = <<EOH
[
  {
    "name": "city",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "coord",
    "type": "RECORD",
    "fields": [
    {
      "name": "lon",
      "type": "FLOAT"
    },
    {
      "name": "lat",
      "type": "FLOAT"
    }
    ]
  }
    ]
  },
  {
    "name": "country",
    "type": "RECORD",
    "fields": [
  {
    "name": "id",
    "type": "INTEGER"
  },
  {
    "name": "name",
    "type": "STRING"
  }
    ]
  }
]
EOH
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      # Check that lack of description doesn't cause any issues
      # Relates to issue : https://github.com/hashicorp/terraform-provider-google/issues/8701
      title       = "%{condition_title_no_desc}"
      expression  = "%{condition_expr_no_desc}"
    }
  }
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
    condition {
      title       = "%{condition_title}"
      description = "%{condition_desc}"
      expression  = "%{condition_expr}"
    }
  }
}

resource "google_bigquery_table_iam_policy" "foo" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}
