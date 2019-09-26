package commands

import (
	"strconv"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
	"../resources"
)

type GetChannels struct {
	ui      terminal.UI
	context plugin.PluginContext
}

func ActionGetChannels(ui terminal.UI, context plugin.PluginContext) *GetChannels {
	return &GetChannels{
		ui:      ui,
		context: context,
	}
}

func (cmd *GetChannels) Run(c *cli.Context) {
	if c.NArg() != 0 {
		FailWithUsage(c, cmd.ui)
	}
	config := resources.GetConfig(cmd.context)
	channels, err := config.GetChannels()
	if err != nil {
		FailWithError(err.Error(), cmd.ui)
	}
	channelsCount := len(channels.Channels)
	cmd.ui.Say("Found %v channel(s).", terminal.EntityNameColor(strconv.Itoa(channelsCount)))
	output := GetFlag("json", c)
	if output {
		json, err := resources.PrintJSON(channels)
		if err != nil {
			FailWithError(err.Error(), cmd.ui)
		}
		cmd.ui.Say(json)
	} else {
		cmd.ui.Say("")
		tbl := cmd.ui.Table([]string{"ID", "Name", "Endpoint"})
		for _, entry := range channels.Channels {
			tbl.Add(entry.ID, entry.Name, entry.Endpoint)
		}
		tbl.Print()
		cmd.ui.Say("")
		cmd.ui.Say("Add " + terminal.CommandColor("-j") + " or " + terminal.CommandColor("--json") + " in the command to get full JSON payload.")
	}
}
