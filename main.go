package main

import (
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	sa_plugin "github.ibm.com/gaurgosw/ibmcloud_sa_cli/plugin"
)

var PluginBuildNumber string

func main() {
	sa_plugin.DynamicBuildNumber = PluginBuildNumber
	plugin.Start(new(sa_plugin.SAPlugin))
}
