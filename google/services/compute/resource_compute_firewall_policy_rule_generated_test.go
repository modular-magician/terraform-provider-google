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

package compute_test

import (
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccComputeFirewallPolicyRule_firewallPolicyRuleExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"service_account": envvar.GetTestServiceAccountFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeFirewallPolicyRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeFirewallPolicyRule_firewallPolicyRuleExample(context),
			},
			{
				ResourceName:            "google_compute_firewall_policy_rule.policy_rule",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"firewall_policy"},
			},
		},
	})
}

func testAccComputeFirewallPolicyRule_firewallPolicyRuleExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_address_group" "basic_global_networksecurity_address_group" {
  name        = "address%{random_suffix}"
  parent      = "organizations/%{org_id}"
  description = "Sample global networksecurity_address_group"
  location    = "global"
  items       = ["208.80.154.224/32"]
  type        = "IPV4"
  capacity    = 100
}

resource "google_folder" "folder" {
  display_name        = "folder%{random_suffix}"
  parent              = "organizations/%{org_id}"
  deletion_protection = false
}

resource "google_compute_firewall_policy" "default" {
  parent      = google_folder.folder.id
  short_name  = "policy%{random_suffix}"
  description = "Resource created for Terraform acceptance testing"
}

resource "google_compute_firewall_policy_rule" "policy_rule" {
  firewall_policy = google_compute_firewall_policy.default.id
  description     = "Resource created for Terraform acceptance testing"
  priority        = 9000
  enable_logging  = true
  action          = "allow"
  direction       = "EGRESS"
  disabled        = false

  match {
    layer4_configs {
      ip_protocol = "tcp"
      ports = [8080]
    }
    layer4_configs {
      ip_protocol = "udp"
      ports = [22]
    }
    dest_ip_ranges = ["11.100.0.1/32"]
    dest_fqdns = []
    dest_region_codes = ["US"]
    dest_threat_intelligences = ["iplist-known-malicious-ips"]
    src_address_groups = []
    dest_address_groups = [google_network_security_address_group.basic_global_networksecurity_address_group.id]
  }
  target_service_accounts = ["%{service_account}"]
}
`, context)
}

func testAccCheckComputeFirewallPolicyRuleDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_firewall_policy_rule" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			log.Printf("[DEBUG] Ignoring destroy during test")
		}

		return nil
	}
}
