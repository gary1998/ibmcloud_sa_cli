package commands

import (
	"encoding/json"

	"github.com/urfave/cli"
	"github.ibm.com/gaurgosw/ibmcloud_sa_cli/plugin/resources"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
)

type GetGraph struct {
	ui      terminal.UI
	context plugin.PluginContext
}

func ActionGetGraph(ui terminal.UI, context plugin.PluginContext) *GetGraph {
	return &GetGraph{
		ui:      ui,
		context: context,
	}
}

func (cmd *GetGraph) Run(c *cli.Context) {
	if c.NArg() != 1 {
		FailWithUsage(c, cmd.ui)
	}

	var body string

	config := resources.GetConfig(cmd.context)
	file := c.Args().First()
	query_type := GetRequiredString("query-type", "json", c, cmd.ui)

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
	cmd.ui.Say("Querying...")
	resp, err := config.PostGraph(body, query_type)
	if err != nil {
		FailWithError(err.Error(), cmd.ui)
	}
	a, err := json.Marshal(resp)
	if err != nil {
		FailWithError(err.Error(), cmd.ui)
	} else {
		cmd.ui.Say(string(a))
	}
}
