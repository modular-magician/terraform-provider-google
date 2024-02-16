// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package functions

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = LocationFromIdFunction{}

func NewLocationFromIdFunction() function.Function {
	return &LocationFromIdFunction{}
}

type LocationFromIdFunction struct{}

func (f LocationFromIdFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "location_from_id"
}

func (f LocationFromIdFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Returns the location name within the resource id or self link provided as an argument.",
		Description: "Takes a single string argument, which should be an id or self link of a resource. This function will either return the location name from the input string or raise an error due to no location being present in the string. The function uses the presence of \"locations/{{location}}/\" in the input string to identify the location name, e.g. when the function is passed the id \"projects/my-project/locations/us-central1/jobs/my-cloudrun-job\" as an argument it will return \"us-central1\".",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "id",
				Description: "An id of a resouce, or a self link. For example, both \"projects/my-project/locations/us-central1/jobs/my-cloudrun-job\" and \"https://us-central1-run.googleapis.com/v2/projects/my-project/locations/us-central1/jobs/my-cloudrun-job\" are valid inputs",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f LocationFromIdFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {

	// Load arguments from function call
	var arg0 string
	resp.Diagnostics.Append(req.Arguments.GetArgument(ctx, 0, &arg0)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Prepare how we'll identify location name from input string
	regex := regexp.MustCompile("locations/(?P<LocationName>[^/]+)/") // Should match the pattern below
	template := "$LocationName"                                       // Should match the submatch identifier in the regex
	pattern := "locations/{location}/"                                // Human-readable pseudo-regex pattern used in errors and warnings

	// Get and return element from input string
	location := GetElementFromId(ctx, arg0, regex, template, pattern, req, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.Result.Set(ctx, location)...)
}
