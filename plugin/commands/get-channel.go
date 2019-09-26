package commands

import (
	"encoding/json"
	"strconv"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
	"github.ibm.com/gaurgosw/ibmcloud_sa_cli/plugin/resources"
)

type GetChannel struct {
	ui      terminal.UI
	context plugin.PluginContext
}

func ActionGetChannel(ui terminal.UI, context plugin.PluginContext) *GetChannel {
	return &GetChannel{
		ui:      ui,
		context: context,
	}
}

func (cmd *GetChannel) Run(c *cli.Context) {
	if c.NArg() != 1 {
		FailWithUsage(c, cmd.ui)
	}
	config := resources.GetConfig(cmd.context)
	channelID := c.Args().First()
	channels, err := config.GetChannel(channelID)
	if err != nil {
		FailWithError(err.Error(), cmd.ui)
	}
	output := GetFlag("json", c)
	if output {
		json, err := resources.PrintJSON(channels)
		if err != nil {
			FailWithError(err.Error(), cmd.ui)
		}
		cmd.ui.Say(json)
	} else {
		cmd.ui.Say("")
		table := cmd.ui.Table([]string{"Property", "Value"})
		table.Add("Name: ", channels.Channel.Name)
		table.Add("Description: ", channels.Channel.Description)
		table.Add("Type: ", channels.Channel.Type)
		table.Add("ID: ", channels.Channel.ID)
		table.Add("Frequency: ", channels.Channel.Frequency)
		table.Add("Severity (Low) : ", strconv.FormatBool(channels.Channel.Severity.Low))
		table.Add("Severity (Medium) : ", strconv.FormatBool(channels.Channel.Severity.Medium))
		table.Add("Severity (High) : ", strconv.FormatBool(channels.Channel.Severity.High))
		table.Add("Endpoint: ", channels.Channel.Endpoint)
		table.Add("Enabled: ", strconv.FormatBool(channels.Channel.Enabled))
		sources, _ := json.Marshal(channels.Channel.AlertSource)
		table.Add("Alert Sources: ", string(sources))
		table.Print()
		cmd.ui.Say("")
		cmd.ui.Say("Add " + terminal.CommandColor("-j") + " or " + terminal.CommandColor("--json") + " in the command to get full JSON payload.")
	}
}
