// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package dialogflowcx_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccDialogflowCXIntent_dialogflowcxIntentFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDialogflowCXIntentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDialogflowCXIntent_dialogflowcxIntentFullExample(context),
			},
			{
				ResourceName:            "google_dialogflow_cx_intent.basic_intent",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccDialogflowCXIntent_dialogflowcxIntentFullExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dialogflow_cx_agent" "agent" {
  display_name = "tf-test-dialogflowcx-agent%{random_suffix}"
  location = "global"
  default_language_code = "en"
  supported_language_codes = ["fr","de","es"]
  time_zone = "America/New_York"
  description = "Example description."
  avatar_uri = "https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png"
  enable_stackdriver_logging = true
  enable_spell_correction    = true
  speech_to_text_settings {
    enable_speech_adaptation = true
  }
}


resource "google_dialogflow_cx_intent" "basic_intent" {
  parent       = google_dialogflow_cx_agent.agent.id
  display_name = "Example"
  priority     = 1
  description  = "Intent example"
  training_phrases {
     parts {
         text = "training"
     }

     parts {
         text = "phrase"
     }

     parts {
         text = "example"
     }

     repeat_count = 1
  }

  parameters {
    id          = "param1"
    entity_type = "projects/-/locations/-/agents/-/entityTypes/sys.date"
  }

  labels  = {
      label1 = "value1",
      label2 = "value2"
   } 
} 
`, context)
}

func TestAccDialogflowCXIntent_dialogflowcxIntentDefaultNegativeIntentExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDialogflowCXIntentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDialogflowCXIntent_dialogflowcxIntentDefaultNegativeIntentExample(context),
			},
			{
				ResourceName:            "google_dialogflow_cx_intent.default_negative_intent",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccDialogflowCXIntent_dialogflowcxIntentDefaultNegativeIntentExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dialogflow_cx_agent" "agent" {
  display_name          = "tf-test-dialogflowcx-agent%{random_suffix}"
  location              = "global"
  default_language_code = "en"
  time_zone             = "America/New_York"
}


resource "google_dialogflow_cx_intent" "default_negative_intent" {
  parent                     = google_dialogflow_cx_agent.agent.id
  is_default_negative_intent = true
  display_name               = "Default Negative Intent"
  priority                   = 1
  is_fallback                = true
  training_phrases {
     parts {
         text = "Never match this phrase"
     }
     repeat_count = 1
  }
}
`, context)
}

func TestAccDialogflowCXIntent_dialogflowcxIntentDefaultWelcomeIntentExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDialogflowCXIntentDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDialogflowCXIntent_dialogflowcxIntentDefaultWelcomeIntentExample(context),
			},
			{
				ResourceName:            "google_dialogflow_cx_intent.default_welcome_intent",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccDialogflowCXIntent_dialogflowcxIntentDefaultWelcomeIntentExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dialogflow_cx_agent" "agent" {
  display_name          = "tf-test-dialogflowcx-agent%{random_suffix}"
  location              = "global"
  default_language_code = "en"
  time_zone             = "America/New_York"
}


resource "google_dialogflow_cx_intent" "default_welcome_intent" {
  parent                    = google_dialogflow_cx_agent.agent.id
  is_default_welcome_intent = true
  display_name              = "Default Welcome Intent"
  priority                  = 1
  training_phrases {
     parts {
         text = "Hello"
     }
     repeat_count = 1
  }
}
`, context)
}

func testAccCheckDialogflowCXIntentDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dialogflow_cx_intent" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{DialogflowCXBasePath}}{{parent}}/intents/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("DialogflowCXIntent still exists at %s", url)
			}
		}

		return nil
	}
}
