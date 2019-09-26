package commands

import (
	"strconv"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
	"../resources"
)

type GetProviders struct {
	ui      terminal.UI
	context plugin.PluginContext
}

func ActionGetProviders(ui terminal.UI, context plugin.PluginContext) *GetProviders {
	return &GetProviders{
		ui:      ui,
		context: context,
	}
}

func (cmd *GetProviders) Run(c *cli.Context) {
	if c.NArg() != 0 {
		FailWithUsage(c, cmd.ui)
	}
	config := resources.GetConfig(cmd.context)
	providers, err := config.GetProviders()
	if err != nil {
		FailWithError(err.Error(), cmd.ui)
	}
	providersCount := len(providers.Providers)
	cmd.ui.Say("Found %v provider(s).", terminal.EntityNameColor(strconv.Itoa(providersCount)))
	output := GetFlag("json", c)
	if output {
		json, err := resources.PrintJSON(providers)
		if err != nil {
			FailWithError(err.Error(), cmd.ui)
		}
		cmd.ui.Say(json)
	} else {
		cmd.ui.Say("")
		tbl := cmd.ui.Table([]string{"ID", "Name"})
		for _, entry := range providers.Providers {
			tbl.Add(entry.Id, entry.Name)
		}
		tbl.Print()
		cmd.ui.Say("")
		cmd.ui.Say("Add " + terminal.CommandColor("-j") + " or " + terminal.CommandColor("--json") + " in the command to get full JSON payload.")
	}
}
