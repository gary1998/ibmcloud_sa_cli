package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/trace"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
	"../models"
)

func FailWithUsage(context *cli.Context, ui terminal.UI) {
	ui.Say(terminal.CrashedColor("FAILED"))
	ui.Say("Incorrect Usage.\n")
	cli.ShowCommandHelp(context, context.Command.Name)
	os.Exit(0)
}

func FailWithError(message string, ui terminal.UI) {
	trace.Logger.Println("Failing with error: ", message)
	ui.Failed(message)
	os.Exit(0)
}

func GetFlag(flag string, context *cli.Context) bool {
	return context.Bool(flag)
}

func GetRequiredString(flagName string, defaultVal interface{}, context *cli.Context, ui terminal.UI) string {
	val := context.String(flagName)
	if val != "" {
		return val
	} else if defaultVal != nil {
		return defaultVal.(string)
	} else {
		FailWithError(fmt.Sprintf("Parameter --%s is required.", flagName), ui)
		return ""
	}
}

func GetRequiredInt(flagName string, defaultVal interface{}, context *cli.Context, ui terminal.UI) int {
	val := context.Int(flagName)
	if val != 0 {
		return val
	} else if defaultVal != nil {
		return defaultVal.(int)
	} else {
		FailWithError(fmt.Sprintf("Parameter --%s is required.", flagName), ui)
		return 0
	}
}

func AssertConfigured(context plugin.PluginContext, ui terminal.UI) bool {
	isConfigured := true
	if isConfigured {
		return true
	}
	FailWithError("The client is not yet configured. Run `ibmcloud login`.", ui)
	return false
}

func GetCLIApp(context plugin.PluginContext, ui terminal.UI) *cli.App {
	app := cli.NewApp()
	app.Name = "ibmcloud sa"
	app.Commands = make([]cli.Command, len(models.CLICommands))

	actionGetGraph := func(c *cli.Context) {
		ActionGetGraph(ui, context).Run(c)
	}
	actionGetProviders := func(c *cli.Context) {
		ActionGetProviders(ui, context).Run(c)
	}
	actionCreateNote := func(c *cli.Context) {
		ActionCreateNote(ui, context).Run(c)
	}
	actionUpdateNote := func(c *cli.Context) {
		ActionUpdateNote(ui, context).Run(c)
	}
	actionDeleteNote := func(c *cli.Context) {
		ActionDeleteNote(ui, context).Run(c)
	}
	actionGetNote := func(c *cli.Context) {
		ActionGetNote(ui, context).Run(c)
	}
	actionGetNotes := func(c *cli.Context) {
		ActionGetNotes(ui, context).Run(c)
	}
	actionCreateChannel := func(c *cli.Context) {
		ActionCreateChannel(ui, context).Run(c)
	}
	actionUpdateChannel := func(c *cli.Context) {
		ActionUpdateChannel(ui, context).Run(c)
	}
	actionDeleteChannel := func(c *cli.Context) {
		ActionDeleteChannel(ui, context).Run(c)
	}
	actionDeleteChannels := func(c *cli.Context) {
		ActionDeleteChannels(ui, context).Run(c)
	}
	actionTestChannel := func(c *cli.Context) {
		ActionTestChannel(ui, context).Run(c)
	}
	actionGetKey := func(c *cli.Context) {
		ActionGetKey(ui, context).Run(c)
	}
	actionGetChannel := func(c *cli.Context) {
		ActionGetChannel(ui, context).Run(c)
	}
	actionGetChannels := func(c *cli.Context) {
		ActionGetChannels(ui, context).Run(c)
	}

	actions := map[string]func(c *cli.Context){
		models.CmdGetGraph:       actionGetGraph,
		models.CmdGetProviders:   actionGetProviders,
		models.CmdCreateNote:     actionCreateNote,
		models.CmdUpdateNote:     actionUpdateNote,
		models.CmdGetNotes:       actionGetNotes,
		models.CmdGetNote:        actionGetNote,
		models.CmdDeleteNote:     actionDeleteNote,
		models.CmdCreateChannel:  actionCreateChannel,
		models.CmdUpdateChannel:  actionUpdateChannel,
		models.CmdGetChannels:    actionGetChannels,
		models.CmdGetChannel:     actionGetChannel,
		models.CmdTestChannel:    actionTestChannel,
		models.CmdGetKey:         actionGetKey,
		models.CmdDeleteChannel:  actionDeleteChannel,
		models.CmdDeleteChannels: actionDeleteChannels,
	}

	for index, command := range models.CLICommands {
		app.Commands[index] = cli.Command{
			Name:        command.Name,
			Description: command.Description,
			Usage:       command.Usage,
			Flags:       command.Flags,
			Action:      actions[command.Name],
		}
	}
	return app
}

func GetCommandContext(pluginContext plugin.PluginContext, name string, args []string) *cli.Context {

	app := GetCLIApp(pluginContext, terminal.NewStdUI())
	set := flag.NewFlagSet(app.Name, flag.ContinueOnError)
	for _, f := range app.Command(name).Flags {
		f.Apply(set)
	}
	set.Parse(args)
	cli := cli.NewContext(app, set, nil)
	cli.Command.Name = name
	return cli
}
