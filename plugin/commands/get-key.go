package commands

import (
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
	"github.ibm.com/gaurgosw/ibmcloud_sa_cli/plugin/resources"
)

type GetKey struct {
	ui      terminal.UI
	context plugin.PluginContext
}

func ActionGetKey(ui terminal.UI, context plugin.PluginContext) *GetKey {
	return &GetKey{
		ui:      ui,
		context: context,
	}
}

func (cmd *GetKey) Run(c *cli.Context) {
	if c.NArg() != 0 {
		FailWithUsage(c, cmd.ui)
	}
	config := resources.GetConfig(cmd.context)
	key, err := config.GetKey()
	if err != nil {
		FailWithError(err.Error(), cmd.ui)
	}
	cmd.ui.Ok()
	cmd.ui.Say(key.Key)
}
