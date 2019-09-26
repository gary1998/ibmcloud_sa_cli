package models

import (
	"reflect"
	"sort"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/urfave/cli"
)

type NameSorter []cli.Command

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

const (
	Namespace = "sa"

	CmdGetGraph = "get-graph"

	CmdGetProviders = "get-providers"

	CmdCreateNote = "create-note"

	CmdUpdateNote = "update-note"

	CmdGetNotes = "get-notes"

	CmdGetNote = "get-note"

	CmdDeleteNote = "delete-note"

	CmdCreateChannel = "create-channel"

	CmdUpdateChannel = "update-channel"

	CmdGetChannels = "get-channels"

	CmdGetChannel = "get-channel"

	CmdTestChannel = "test-channel"

	CmdGetKey = "get-key"

	CmdDeleteChannel = "delete-channel"

	CmdDeleteChannels = "delete-channels"
)

var (
	HelpTemplate = "NAME:" + `
		{{.Name}} - {{.Description}}{{with .ShortName}}
		` + "ALIAS:" + `
		   {{.}}{{end}}

		` + "USAGE:" + `
		   {{.Usage}}
		{{with .Flags}}
		` + "OPTIONS:" + `
		{{range .}}   {{.}}
		{{end}}{{end}}
`

	CLICommands = []cli.Command{
		cli.Command{
			Category:    Namespace,
			Name:        CmdGetGraph,
			Description: "Get notes and occurrences using GraphQL queries.",
			Usage:       "ibmcloud sa get-graph <QUERY_FILE>",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "query-type, q",
					Usage:    "[Optional] Type of Query, either 'json' or 'graphql'.",
					Value:    "json",
					Required: false,
				},
			},
			ArgsUsage: "<QUERY_FILE>",
		},
		cli.Command{
			Category:    Namespace,
			Name:        CmdGetProviders,
			Description: "Returns the list of all providers using IBM Cloud Findings API.",
			Usage:       "ibmcloud sa get-providers",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:     "json, j",
					Usage:    "[Optional] Activate if JSON response is required.",
					Required: false,
				},
			},
		},
		cli.Command{
			Category:    Namespace,
			Name:        CmdCreateNote,
			Description: "Create a new note using IBM Cloud Security Advisor Findings API.",
			Usage:       "ibmcloud sa create-note <PROVIDER_ID> <BODY_FILE>",
			Flags:       []cli.Flag{},
			ArgsUsage:   "<PROVIDER_ID> <BODY_FILE>",
		},
		cli.Command{
			Category:    Namespace,
			Name:        CmdUpdateNote,
			Description: "Updates the details of a specified note using IBM Cloud Security Advisor Findings API.",
			Usage:       "ibmcloud sa create-note <PROVIDER_ID> <NOTE_ID> <BODY_FILE>",
			Flags:       []cli.Flag{},
			ArgsUsage:   "<PROVIDER_ID> <NOTE_ID> <BODY_FILE>",
		},
		cli.Command{
			Category:    Namespace,
			Name:        CmdDeleteNote,
			Description: "Delete a note using IBM Cloud Security Advisor Findings API.",
			Usage:       "ibmcloud sa delete-note <PROVIDER_ID> <NOTE_ID>",
			Flags:       []cli.Flag{},
			ArgsUsage:   "<PROVIDER_ID> <NOTE_ID>",
		},
		{
			Category:    Namespace,
			Name:        CmdGetNotes,
			Description: "Return a list of all notes using IBM Cloud Security Advisor Findings API.",
			Usage:       "ibmcloud sa get-notes <PROVIDER_ID>",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:     "json, j",
					Usage:    "[Optional] Activate if JSON response is required.",
					Required: false,
				},
			},
			ArgsUsage: "<PROVIDER_ID>",
		},
		{
			Category:    Namespace,
			Name:        CmdGetNote,
			Description: "Return the details of a specified note using IBM Cloud Security Advisor Findings API.",
			Usage:       "ibmcloud sa get-note <PROVIDER_ID> <RESOURCE_ID>",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "resource-type, r",
					Usage:    "[Optional] Type of resource whose ID is supplied, either 'occurrence' or 'note'.",
					Value:    "note",
					Required: false,
				},
				cli.BoolFlag{
					Name:     "json, j",
					Usage:    "[Optional] Activate if JSON response is required.",
					Required: false,
				},
			},
			ArgsUsage: "<PROVIDER_ID> <RESOURCE_ID>",
		},
		{
			Category:    Namespace,
			Name:        CmdCreateChannel,
			Description: "Create a new notification channel using IBM Cloud Security Advisor Notifications API.",
			Usage:       "ibmcloud sa create-channel <BODY_FILE>",
			Flags:       []cli.Flag{},
			ArgsUsage:   "<BODY_FILE>",
		},
		{
			Category:    Namespace,
			Name:        CmdUpdateChannel,
			Description: "Update the details of specific notification channel using IBM Cloud Security Advisor Notifications API.",
			Usage:       "ibmcloud sa update-channel <CHANNEL_ID> <BODY_FILE>",
			Flags:       []cli.Flag{},
			ArgsUsage:   "<CHANNEL_ID> <BODY_FILE>",
		},
		{
			Category:    Namespace,
			Name:        CmdGetChannel,
			Description: "Return the details of a specified notification channel using IBM Cloud Security Advisor Notifications API.",
			Usage:       "ibmcloud sa get-channel <CHANNEL_ID>",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:     "json, j",
					Usage:    "[Optional] Activate if JSON response is required.",
					Required: false,
				},
			},
			ArgsUsage: "<CHANNEL_ID>",
		},
		{
			Category:    Namespace,
			Name:        CmdGetChannels,
			Description: "Return the list of all notification channels using IBM Cloud Security Advisor Notifications API.",
			Usage:       "ibmcloud sa get-channels",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:     "json, j",
					Usage:    "[Optional] Activate if JSON response is required.",
					Required: false,
				},
			},
		},
		{
			Category:    Namespace,
			Name:        CmdDeleteChannel,
			Description: "Delete a specific notification channel using IBM Cloud Security Advisor Notifications API.",
			Usage:       "ibmcloud sa delete-channel <CHANNEL_ID>",
			Flags:       []cli.Flag{},
			ArgsUsage:   "<CHANNEL_ID>",
		},
		{
			Category:    Namespace,
			Name:        CmdDeleteChannels,
			Description: "Delete multiple notification channels (bulk) using IBM Cloud Security Advisor Notifications API.",
			Usage:       "ibmcloud sa delete-channels <BODY_FILE>",
			Flags:       []cli.Flag{},
			ArgsUsage:   "<BODY_FILE>",
		},
		{
			Category:    Namespace,
			Name:        CmdTestChannel,
			Description: "Test a specific notification channel using IBM Cloud Security Advisor Notifications API.",
			Usage:       "ibmcloud sa test-channel <CHANNEL_ID>",
			Flags:       []cli.Flag{},
			ArgsUsage:   "<CHANNEL_ID>",
		},
		{
			Category:    Namespace,
			Name:        CmdGetKey,
			Description: "Returns public key using IBM Cloud Security Advisor Notifications API.",
			Usage:       "ibmcloud sa get-key",
			Flags:       []cli.Flag{},
		},
	}

	Namespaces = []plugin.Namespace{
		plugin.Namespace{
			Name:        Namespace,
			Description: "Interact with IBM Cloud Security Advisor.",
		},
	}
)

func GetPluginCommandDefinition() []plugin.Command {

	sort.Sort(NameSorter(CLICommands))

	pluginCommands := make([]plugin.Command, len(CLICommands))

	for index, cliCmd := range CLICommands {
		pFlags := make([]plugin.Flag, len(cliCmd.Flags))
		for i, flag := range cliCmd.Flags {
			value := reflect.ValueOf(flag)
			flagType := reflect.TypeOf(flag)
			pFlags[i] = plugin.Flag{
				Name:        value.FieldByName("Name").String(),
				Description: value.FieldByName("Usage").String(),
				HasValue:    flagType.String() != "cli.BoolFlag",
			}
		}
		pluginCommands[index] = plugin.Command{
			Namespace:   cliCmd.Category,
			Name:        cliCmd.Name,
			Description: cliCmd.Description,
			Usage:       cliCmd.Usage,
			Flags:       pFlags,
		}
	}
	return pluginCommands
}
