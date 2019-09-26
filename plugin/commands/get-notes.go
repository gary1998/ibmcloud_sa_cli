package commands

import (
	"strconv"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
	"../resources"
)

type GetNotes struct {
	ui      terminal.UI
	context plugin.PluginContext
}

func ActionGetNotes(ui terminal.UI, context plugin.PluginContext) *GetNotes {
	return &GetNotes{
		ui:      ui,
		context: context,
	}
}

func (cmd *GetNotes) Run(c *cli.Context) {
	if c.NArg() != 1 {
		FailWithUsage(c, cmd.ui)
	}
	config := resources.GetConfig(cmd.context)
	providerId := c.Args().First()
	notes, err := config.GetNotes(providerId)
	if err != nil {
		FailWithError(err.Error(), cmd.ui)
	}
	notesCount := len(notes.Notes)
	cmd.ui.Say("Found %v note(s).", terminal.EntityNameColor(strconv.Itoa(notesCount)))
	output := GetFlag("json", c)
	if output {
		json, err := resources.PrintJSON(notes)
		if err != nil {
			FailWithError(err.Error(), cmd.ui)
		}
		cmd.ui.Say(json)
	} else {
		cmd.ui.Say("")
		tbl := cmd.ui.Table([]string{"ID", "Name"})
		for _, entry := range notes.Notes {
			tbl.Add(entry.Id, entry.Name)
		}
		tbl.Print()
		cmd.ui.Say("")
		cmd.ui.Say(terminal.EntityNameColor("Next Page Token: ") + notes.NextPageToken)
		cmd.ui.Say("")
		cmd.ui.Say("Add " + terminal.CommandColor("-j") + " or " + terminal.CommandColor("--json") + " in the command to get full JSON payload.")
	}
}
