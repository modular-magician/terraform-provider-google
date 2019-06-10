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

var FilestoreDefaultBasePath = "https://file.googleapis.com/v1/"
var FilestoreCustomEndpointEntryKey = "filestore_custom_endpoint"
var FilestoreCustomEndpointEntry = &schema.Schema{
	Type:         schema.TypeString,
	Optional:     true,
	ValidateFunc: validateCustomEndpoint,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_FILESTORE_CUSTOM_ENDPOINT",
	}, FilestoreDefaultBasePath),
}

var GeneratedFilestoreResourcesMap = map[string]*schema.Resource{
	"google_filestore_instance": resourceFilestoreInstance(),
}
