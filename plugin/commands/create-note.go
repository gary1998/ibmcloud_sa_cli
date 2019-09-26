package commands

import (
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
	"github.ibm.com/gaurgosw/ibmcloud_sa_cli/plugin/resources"
)

type CreateNote struct {
	ui      terminal.UI
	context plugin.PluginContext
}

func ActionCreateNote(ui terminal.UI, context plugin.PluginContext) *CreateNote {
	return &CreateNote{
		ui:      ui,
		context: context,
	}
}

func (cmd *CreateNote) Run(c *cli.Context) {
	if c.NArg() != 2 {
		FailWithUsage(c, cmd.ui)
	}
	config := resources.GetConfig(cmd.context)
	var body string
	providerID := c.Args().First()
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
	cmd.ui.Say("Creating new note...")
	_, err := config.PostNote(providerID, body)
	if err != nil {
		FailWithError(err.Error(), cmd.ui)
	}
	cmd.ui.Say("Note created successfully!")
}
