package commands

import (
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
	"github.ibm.com/gaurgosw/ibmcloud_sa_cli/plugin/resources"
)

type DeleteChannel struct {
	ui      terminal.UI
	context plugin.PluginContext
}

func ActionDeleteChannel(ui terminal.UI, context plugin.PluginContext) *DeleteChannel {
	return &DeleteChannel{
		ui:      ui,
		context: context,
	}
}

func (cmd *DeleteChannel) Run(c *cli.Context) {
	if c.NArg() != 1 {
		FailWithUsage(c, cmd.ui)
	}
	config := resources.GetConfig(cmd.context)
	channelID := c.Args().First()
	cmd.ui.Say("Deleting channel...")
	_, err := config.DeleteChannel(channelID)
	if err != nil {
		FailWithError(err.Error(), cmd.ui)
	}
	cmd.ui.Ok()
	cmd.ui.Say("Channel deleted successfully!")
}
