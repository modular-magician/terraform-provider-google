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

import "github.com/hashicorp/terraform/helper/schema"

var AppEngineDefaultBasePath = "https://appengine.googleapis.com/v1/"
var AppEngineCustomEndpointEntryKey = "app_engine_custom_endpoint"
var AppEngineCustomEndpointEntry = &schema.Schema{
	Type:         schema.TypeString,
	Optional:     true,
	ValidateFunc: validateCustomEndpoint,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_APP_ENGINE_CUSTOM_ENDPOINT",
	}, AppEngineDefaultBasePath),
}

var GeneratedAppEngineResourcesMap = map[string]*schema.Resource{
	"google_app_engine_firewall_rule": resourceAppEngineFirewallRule(),
}
