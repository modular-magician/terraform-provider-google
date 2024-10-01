// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package fwprovider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

// TestAccFwProvider_terraform_attribution_label_addition_strategy is a series of acc tests asserting how the plugin-framework provider handles terraform_attribution_label_addition_strategy arguments
// It is plugin-framework specific because the HCL used provisions plugin-framework-implemented resources
// It is a counterpart to TestAccSdkProvider_terraform_attribution_label_addition_strategy
func TestAccFwProvider_terraform_attribution_label_addition_strategy(t *testing.T) {
	testCases := map[string]func(t *testing.T){
		// Configuring the provider using inputs
		"config sets terraform_attribution_label_addition_strategy values":                                                                                 testAccFwProvider_terraform_attribution_label_addition_strategy_configUsed,
		"when terraform_attribution_label_addition_strategy is unset in the config, the default value'CREATION_ONLY' is NOT set in the provider meta data": testAccFwProvider_terraform_attribution_label_addition_strategy_defaultValue,
	}

	for name, tc := range testCases {
		// shadow the tc variable into scope so that when
		// the loop continues, if t.Run hasn't executed tc(t)
		// yet, we don't have a race condition
		// see https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		tc := tc
		t.Run(name, func(t *testing.T) {
			tc(t)
		})
	}
}

func testAccFwProvider_terraform_attribution_label_addition_strategy_configUsed(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API

	context1 := map[string]interface{}{
		"terraform_attribution_label_addition_strategy": "CREATION_ONLY",
	}
	context2 := map[string]interface{}{
		"terraform_attribution_label_addition_strategy": "PROACTIVE",
	}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFwProvider_terraform_attribution_label_addition_strategy_inProviderBlock(context1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "terraform_attribution_label_addition_strategy", "CREATION_ONLY"),
				),
			},
			{
				Config: testAccFwProvider_terraform_attribution_label_addition_strategy_inProviderBlock(context2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.google_provider_config_plugin_framework.default", "terraform_attribution_label_addition_strategy", "PROACTIVE"),
				),
			},
		},
	})
}

func testAccFwProvider_terraform_attribution_label_addition_strategy_defaultValue(t *testing.T) {
	acctest.SkipIfVcr(t) // Test doesn't interact with API

	context := map[string]interface{}{}

	acctest.VcrTest(t, resource.TestCase{
		// No PreCheck for checking ENVs
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFwProvider_terraform_attribution_label_addition_strategy_inEnvsOnly(context),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("data.google_provider_config_plugin_framework.default", "terraform_attribution_label_addition_strategy"),
				),
			},
		},
	})
}

// testAccFwProvider_terraform_attribution_label_addition_strategy_inProviderBlock allows setting the terraform_attribution_label_addition_strategy argument in a provider block.
func testAccFwProvider_terraform_attribution_label_addition_strategy_inProviderBlock(context map[string]interface{}) string {
	return acctest.Nprintf(`
provider "google" {
	terraform_attribution_label_addition_strategy = "%{terraform_attribution_label_addition_strategy}"
}

data "google_provider_config_plugin_framework" "default" {}
`, context)
}

// testAccFwProvider_terraform_attribution_label_addition_strategy_inEnvsOnly allows testing when the terraform_attribution_label_addition_strategy argument
// is only supplied via ENVs
func testAccFwProvider_terraform_attribution_label_addition_strategy_inEnvsOnly(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_provider_config_plugin_framework" "default" {}
`, context)
}
