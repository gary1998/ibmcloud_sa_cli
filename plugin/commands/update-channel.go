package commands

import (
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
	"../resources"
)

type UpdateChannel struct {
	ui      terminal.UI
	context plugin.PluginContext
}

func ActionUpdateChannel(ui terminal.UI, context plugin.PluginContext) *UpdateChannel {
	return &UpdateChannel{
		ui:      ui,
		context: context,
	}
}

func (cmd *UpdateChannel) Run(c *cli.Context) {
	if c.NArg() != 2 {
		FailWithUsage(c, cmd.ui)
	}
	config := resources.GetConfig(cmd.context)
	var body string
	channelID := c.Args().First()
	file := c.Args().Get(1)
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
	cmd.ui.Say("Updating channel...")
	_, err := config.PutChannel(channelID, body)
	if err != nil {
		FailWithError(err.Error(), cmd.ui)
		return
	}
	cmd.ui.Say("Channel updated successfully!")
}
