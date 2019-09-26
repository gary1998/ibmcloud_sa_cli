package commands

import (
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
	"../resources"
)

type TestChannel struct {
	ui      terminal.UI
	context plugin.PluginContext
}

func ActionTestChannel(ui terminal.UI, context plugin.PluginContext) *TestChannel {
	return &TestChannel{
		ui:      ui,
		context: context,
	}
}

func (cmd *TestChannel) Run(c *cli.Context) {
	if c.NArg() != 1 {
		FailWithUsage(c, cmd.ui)
	}
	config := resources.GetConfig(cmd.context)
	channelID := c.Args().First()
	resp, err := config.TestChannel(channelID)
	if err != nil {
		FailWithError(err.Error(), cmd.ui)
	}
	cmd.ui.Ok()
	cmd.ui.Say(resp.Message)
}
