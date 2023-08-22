package main

import (
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/michaelkad/power-beta-go-sdk/version"
)

func main() {
	log.Println("IBM Cloud Provider version", version.Version, version.VersionPrerelease, version.GitCommit)
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
