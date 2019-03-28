package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-google/google"
)

func main() {
	"github.com/terraform-providers/terraform-provider-google/google"
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: google.Provider})
}
