package plugin

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/trace"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/common/rest"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"

	"strings"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.ibm.com/gaurgosw/ibmcloud_sa_cli/plugin/commands"
	"github.ibm.com/gaurgosw/ibmcloud_sa_cli/plugin/models"
	"github.ibm.com/gaurgosw/ibmcloud_sa_cli/plugin/resources"
)

type SAPlugin struct {
	ui      terminal.UI
	context plugin.PluginContext
}

var DynamicBuildNumber string

var FindingsAPIURL = "FINDINGS_API_ENDPOINT"

var NotificationsAPIURL = "NOTIFICATIONS_API_ENDPOINT"

var BaseURL = "secadvisor.cloud.ibm.com"

func (c *SAPlugin) GetMetadata() plugin.PluginMetadata {

	if DynamicBuildNumber == "" {
		DynamicBuildNumber = "10000"
	}
	build, err := strconv.Atoi(DynamicBuildNumber)
	if err != nil {
		panic(err.Error())
	}

	return plugin.PluginMetadata{

		Name: resources.PluginName,

		Version: plugin.VersionType{
			Major: resources.PluginMajorVersion,
			Minor: resources.PluginMinorVersion,
			Build: build,
		},

		MinCliVersion: plugin.VersionType{
			Major: 0,
			Minor: 4,
			Build: 9,
		},

		Commands: models.GetPluginCommandDefinition(),

		Namespaces: models.Namespaces,
	}
}

func (c *SAPlugin) Run(context plugin.PluginContext, args []string) {

	defer func() {
		if l, ok := trace.Logger.(trace.Closer); ok {
			l.Close()
		}
	}()

	defer func() {

		if err := recover(); err != nil {

			exitCode, er := strconv.Atoi(fmt.Sprint(err))
			if er == nil {
				os.Exit(exitCode)
			}

			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
	}()

	if os.Getenv("IBMCLOUD_TRACE") != "" {
		trace.Logger = trace.NewLogger(os.Getenv("IBMCLOUD_TRACE"))
	} else {
		trace.Logger = trace.NewLogger(context.Trace())
	}

	if c.ui == nil {
		c.ui = terminal.NewStdUI()
	}

	c.CheckPluginVersion(context, c.ui)

	if !allOK(context, c.ui) {
		os.Exit(0)
	}

	cli.CommandHelpTemplate = models.HelpTemplate
	app := commands.GetCLIApp(context, c.ui)
	app.Run(append([]string{"sa"}, args...))
}

func (c *SAPlugin) CheckPluginVersion(ctx plugin.PluginContext, ui terminal.UI) {

	isNewVersion := false
	currentVersion := c.GetMetadata().Version
	latestVersion := getLatestPluginVersion(ctx)

	if currentVersion.Major < latestVersion.Major {
		isNewVersion = true
	} else if currentVersion.Major == latestVersion.Major {
		if currentVersion.Minor < latestVersion.Minor {
			isNewVersion = true
		} else if currentVersion.Minor == latestVersion.Minor {
			if currentVersion.Build < latestVersion.Build {
				isNewVersion = true
			}
		}
	}

	if isNewVersion {
		updateCmd := "ibmcloud plugin update security-advisor -r Bluemix"
		if strings.Contains(ctx.APIEndpoint(), "api.stage1.ng.bluemix.net") {
			updateCmd = "ibmcloud plugin update security-advisor -r stage"
		}
		ui.Say(fmt.Sprintf("\nPlugin version %s is now available. To update please run: %s\n",
			terminal.EntityNameColor(latestVersion.String()),
			terminal.CommandColor(updateCmd)))
	}
}

func getLatestPluginVersion(ctx plugin.PluginContext) plugin.VersionType {

	cachedLatestValue := resources.GetConfigWithDefault(resources.LatestVersionValueFlag, "", ctx.PluginConfig())
	cachedLatestTime := resources.GetConfigWithDefault(resources.LatestVersionCheckedFlag, "", ctx.PluginConfig())
	if cachedLatestValue != "" && cachedLatestTime != "" {
		trace.Logger.Printf("Located cached latest version: %s @ %s", cachedLatestValue, cachedLatestTime)
		if cachedTime, err := time.Parse(time.RFC3339, cachedLatestTime); err == nil {
			elapsed := time.Since(cachedTime)
			trace.Logger.Printf("Last cache check %s.  Threshold is %s", string(int(elapsed.Seconds())), string(resources.LatestVersionCheckThresh))
			if int(elapsed.Seconds()) <= resources.LatestVersionCheckThresh {
				trace.Logger.Printf("Cache is still valid, returning value: %s", cachedLatestValue)
				if version, err := getVersion(cachedLatestValue); err == nil {
					return version
				}
			}
		}
	}

	var successV []struct {
		Version string    `json:"version"`
		Updated time.Time `json:"updated"`
	}

	var version plugin.VersionType
	url := "https://plugins.ng.bluemix.net/bx/list/security-advisor"
	if strings.Contains(ctx.APIEndpoint(), "api.stage1.ng.bluemix.net") {
		url = strings.Replace(url, "https://plugins.", "https://plugins.stage1.", 1)
	}

	client := rest.NewClient()
	client.HTTPClient.Timeout = 3 * time.Second

	r := rest.GetRequest(url).Method("GET")
	trace.Logger.Printf("calling plugin list endpoint: %s", url)
	if _, err := client.Do(r, &successV, nil); err == nil {
		trace.Logger.Printf("located sa versions: %s", successV)
		if len(successV) > 0 {
			latest := successV[len(successV)-1].Version
			if version, err = getVersion(latest); err == nil {
				resources.SetConfig(resources.LatestVersionValueFlag, latest, ctx.PluginConfig())
				resources.SetConfig(resources.LatestVersionCheckedFlag, time.Now().Format(time.RFC3339), ctx.PluginConfig())
			}
		}
	}

	trace.Logger.Printf("latest version is %s", version)
	return version
}

func getVersion(s string) (plugin.VersionType, error) {
	var version plugin.VersionType
	latestSplit := strings.Split(s, ".")
	if len(latestSplit) == 3 {
		version.Major, _ = strconv.Atoi(latestSplit[0])
		version.Minor, _ = strconv.Atoi(latestSplit[1])
		version.Build, _ = strconv.Atoi(latestSplit[2])
		return version, nil
	}
	return version, errors.New("could not parse version string")
}

func allOK(ctx plugin.PluginContext, ui terminal.UI) bool {
	if CheckLogin(ctx) {
		region := GetRegion(ctx)
		if region == "us-south" || region == "eu-gb" {
			trace.Logger.Println(fmt.Sprintf("Region logged in is %s", region))
			os.Setenv(FindingsAPIURL, "https://"+region+"."+BaseURL+"/findings/v1")
			os.Setenv(NotificationsAPIURL, "https://"+region+"."+BaseURL+"/notifications/v1")
			trace.Logger.Println(fmt.Sprintf("FindingsAPI Endpoint: %s", os.Getenv(FindingsAPIURL)))
			trace.Logger.Println(fmt.Sprintf("NotificationsAPI Endpoint: %s", os.Getenv(NotificationsAPIURL)))
			return true
		} else {
			ui.Say("egion is not supported yet. Try any of %s or %s regions.", terminal.EntityNameColor("us-south"), terminal.EntityNameColor("eu-gb"))
			os.Exit(0)
		}
	} else {
		ui.Failed("Not logged in. Use %s to log in.", terminal.CommandColor(fmt.Sprintf("%s login", ctx.CLIName())))
		os.Exit(0)
	}
	return false
}

func GetRegion(context plugin.PluginContext) string {
	trace.Logger.Println(fmt.Sprintf("Checking region of logged in user"))
	return context.CurrentRegion().Name
}

func CheckLogin(context plugin.PluginContext) bool {
	trace.Logger.Println(fmt.Sprintf("Checking if the user is logged in or not"))
	return context.IsLoggedIn()
}
