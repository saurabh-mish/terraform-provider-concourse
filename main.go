package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/saurabh-mish/terraform-provider-concourse/app"
)

func main() {
	plugin.Serve(&plugin.ServeOpts {
		ProviderFunc: app.Provider,
	})
}
