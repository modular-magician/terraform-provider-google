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

	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDialogflowIntent_dialogflowIntentFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		CheckDestroy: testAccCheckDialogflowIntentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDialogflowIntent_dialogflowIntentFullExample(context),
			},
			{
				ResourceName:      "google_dialogflow_intent.full_intent",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDialogflowIntent_dialogflowIntentFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "agent_project" {
  project_id = "tf-test-dialogflow-%{random_suffix}"
  name = "tf-test-dialogflow-%{random_suffix}"
  org_id = "%{org_id}"
}

resource "google_project_service" "agent_project" {
  project = google_project.agent_project.project_id
  service = "dialogflow.googleapis.com"
  disable_dependent_services = false
}

resource "google_service_account" "dialogflow_service_account" {
  account_id = "tf-test-dialogflow-%{random_suffix}"
}

resource "google_project_iam_member" "agent_create" {
  project = google_project_service.agent_project.project
  role    = "roles/dialogflow.admin"
  member  = "serviceAccount:${google_service_account.dialogflow_service_account.email}"
}

resource "google_dialogflow_agent" "basic_agent" {
  project = google_project.agent_project.project_id
  display_name = "example_agent"
  default_language_code = "en"
  time_zone = "America/New_York"
}

resource "google_dialogflow_intent" "full_intent" {
  project = google_project.agent_project.project_id
  depends_on = [google_dialogflow_agent.basic_agent]
  display_name = "tf-test-full-intent%{random_suffix}"
  webhook_state = "WEBHOOK_STATE_ENABLED"
  priority = 1
  is_fallback = false
  ml_disabled = true
  action = "some_action"
  reset_contexts = true
  input_context_names = ["projects/${google_project.agent_project.project_id}/agent/sessions/-/contexts/some_id"]
  events = ["some_event"]
  default_response_platforms = ["FACEBOOK","SLACK"]
}
`, context)
}

func testAccCheckDialogflowIntentDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dialogflow_intent" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{DialogflowBasePath}}{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("DialogflowIntent still exists at %s", url)
			}
		}

		return nil
	}
}
