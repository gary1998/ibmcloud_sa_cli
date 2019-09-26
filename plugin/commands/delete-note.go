package commands

import (
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
	"github.ibm.com/gaurgosw/ibmcloud_sa_cli/plugin/resources"
)

type DeleteNote struct {
	ui      terminal.UI
	context plugin.PluginContext
}

func ActionDeleteNote(ui terminal.UI, context plugin.PluginContext) *DeleteNote {
	return &DeleteNote{
		ui:      ui,
		context: context,
	}
}

func (cmd *DeleteNote) Run(c *cli.Context) {
	if c.NArg() != 2 {
		FailWithUsage(c, cmd.ui)
	}
	config := resources.GetConfig(cmd.context)
	providerID := c.Args().First()
	noteID := c.Args().Get(1)
	cmd.ui.Say("Deleting note...")
	_, err := config.DeleteNote(providerID, noteID)
	if err != nil {
		FailWithError(err.Error(), cmd.ui)
	}
	cmd.ui.Ok()
	cmd.ui.Say("Note deleted successfully!")
}
