package main

import (
	"github.com/genymobile/terraform-provider-genymotion/genymotion"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return genymotion.Provider()
		},
	})
}
