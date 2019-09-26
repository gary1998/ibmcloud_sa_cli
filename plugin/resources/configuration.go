package resources

import (
	"fmt"
	"os"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
)

// PluginName ...
const PluginName = "security-advisor"

// PluginMajorVersion ...
const PluginMajorVersion = 0

// PluginMinorVersion ...
const PluginMinorVersion = 1

/* CONFIGURATION
   Default values below are configured for prod-dal10 (YS1), however the
   client will download different set of configuration depending on
   environment. The Armada API server hosts a configuration REST endpoint
   for the client to download environment specific data. The client self
   configures on the first init to the API server.
*/

// LatestVersionValueFlag ...
var LatestVersionValueFlag = "latest-version"

// LatestVersionCheckedFlag ...
var LatestVersionCheckedFlag = "latest-version-checked"

// LatestVersionCheckThresh one day
var LatestVersionCheckThresh = 86400

// GetConfigWithDefault ...
func GetConfigWithDefault(name string, defaultValue string, config plugin.PluginConfig) string {
	envValue := os.Getenv(name)
	if envValue == "" {
		configValue := config.Get(name)
		if configValue == nil {
			return defaultValue
		}
		return fmt.Sprintf("%v", configValue)
	}
	return envValue
}

// SetConfig ...
func SetConfig(name string, value interface{}, config plugin.PluginConfig) bool {
	err := config.Set(name, value)
	os.Setenv(name, fmt.Sprintf("%v", value))
	return err != nil
}
