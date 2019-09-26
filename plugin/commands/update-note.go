package commands

import (
	"github.com/urfave/cli"
	"../resources"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
)

type UpdateNote struct {
	ui      terminal.UI
	context plugin.PluginContext
}

func ActionUpdateNote(ui terminal.UI, context plugin.PluginContext) *UpdateNote {
	return &UpdateNote{
		ui:      ui,
		context: context,
	}
}

func (cmd *UpdateNote) Run(c *cli.Context) {
	if c.NArg() != 3 {
		FailWithUsage(c, cmd.ui)
	}
	config := resources.GetConfig(cmd.context)
	var body string
	providerID := c.Args().First()
	noteID := c.Args().Get(1)
	file := c.Args().Get(2)
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
	cmd.ui.Say("Updating note...")
	_, err := config.PutNote(providerID, noteID, body)
	if err != nil {
		FailWithError(err.Error(), cmd.ui)
	}
	cmd.ui.Say("Note updated successfully!")
}
