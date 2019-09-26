package commands

import (
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
	"../resources"
)

type CreateChannel struct {
	ui      terminal.UI
	context plugin.PluginContext
}

func ActionCreateChannel(ui terminal.UI, context plugin.PluginContext) *CreateChannel {
	return &CreateChannel{
		ui:      ui,
		context: context,
	}
}

func (cmd *CreateChannel) Run(c *cli.Context) {
	if c.NArg() != 1 {
		FailWithUsage(c, cmd.ui)
	}
	config := resources.GetConfig(cmd.context)
	var body string
	file := c.Args().First()
	if file == "" {
		FailWithError("No file specified.", cmd.ui)
	} else {
		cmd.ui.Say(terminal.HeaderColor("READING FILE..."))
		query, err := config.ReadFile(file)
		if err != nil {
			FailWithError(err.Error(), cmd.ui)
			return
		}
		cmd.ui.Ok()
		body = string(query)
	}
	cmd.ui.Say("Creating new channel...")
	_, err := config.PostChannel(body)
	if err != nil {
		FailWithError(err.Error(), cmd.ui)
	}
	cmd.ui.Say("Channel created successfully!")
}
