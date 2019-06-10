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

var DnsDefaultBasePath = "https://www.googleapis.com/dns/v1/"
var DnsCustomEndpointEntryKey = "dns_custom_endpoint"
var DnsCustomEndpointEntry = &schema.Schema{
	Type:         schema.TypeString,
	Optional:     true,
	ValidateFunc: validateCustomEndpoint,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_DNS_CUSTOM_ENDPOINT",
	}, DnsDefaultBasePath),
}

var GeneratedDnsResourcesMap = map[string]*schema.Resource{
	"google_dns_managed_zone": resourceDnsManagedZone(),
}
