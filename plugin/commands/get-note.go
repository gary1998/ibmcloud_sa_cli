package commands

import (
	"encoding/json"
	"strconv"

	"github.com/urfave/cli"
	"github.ibm.com/gaurgosw/ibmcloud_sa_cli/plugin/resources"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.ibm.com/gaurgosw/ibmcloud_sa_cli/plugin/models"
)

type GetNote struct {
	ui      terminal.UI
	context plugin.PluginContext
}

func ActionGetNote(ui terminal.UI, context plugin.PluginContext) *GetNote {
	return &GetNote{
		ui:      ui,
		context: context,
	}
}

func (cmd *GetNote) Run(c *cli.Context) {
	if c.NArg() != 2 {
		FailWithUsage(c, cmd.ui)
	}
	config := resources.GetConfig(cmd.context)
	providerId := c.Args().First()
	resId := c.Args().Get(1)
	resourceType := GetRequiredString("resource-type", "note", c, cmd.ui)
	var note models.ApiNote
	var err error
	if resourceType == "note" {
		note, err = config.GetNoteByNoteID(providerId, resId)
	} else {
		note, err = config.GetNoteByOccId(providerId, resId)
	}
	if err != nil {
		FailWithError(err.Error(), cmd.ui)
	}
	output := GetFlag("json", c)
	if output {
		json, err := resources.PrintJSON(note)
		if err != nil {
			FailWithError(err.Error(), cmd.ui)
		}
		cmd.ui.Say(json)
	} else {
		cmd.ui.Say("")
		table := cmd.ui.Table([]string{"Property", "Value"})
		table.Add("Name: ", note.Name)
		table.Add("Short Description: ", note.ShortDescription)
		table.Add("Long Description: ", note.LongDescription)
		table.Add("Kind: ", note.Kind)
		relatedurl, _ := json.Marshal(note.RelatedUrl)
		table.Add("Related URL: ", string(relatedurl))
		table.Add("Expiration Time: ", note.ExpirationTime.String())
		table.Add("Creation Time: ", note.CreateTime.String())
		table.Add("Updation Time: ", note.UpdateTime.String())
		table.Add("Provider ID: ", note.ProviderId)
		table.Add("ID: ", note.Id)
		table.Add("Shared: ", strconv.FormatBool(note.Shared))
		reportedby, _ := json.Marshal(note.ReportedBy)
		table.Add("Reported By: ", string(reportedby))
		table.Add("Finding Severity: ", note.Finding.Severity)
		nextsteps, _ := json.Marshal(note.Finding.NextSteps)
		table.Add("Finding Next Steps: ", string(nextsteps))
		kpi, _ := json.Marshal(note.Kpi)
		table.Add("KPI: ", string(kpi))
		card, _ := json.Marshal(note.Card)
		table.Add("Card: ", string(card))
		section, _ := json.Marshal(note.Section)
		table.Add("Section: ", string(section))
		table.Print()
		cmd.ui.Say("")
		cmd.ui.Say("Add " + terminal.CommandColor("-j") + " or " + terminal.CommandColor("--json") + " in the command to get full JSON payload.")
	}
}
