// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package kivm

import (
	"kic/boa"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// ListCmd lists the clusters in the fleet
var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check status on the local cluster",
}

func init() {
	boa.AddScriptCommand(CheckCmd, "flux", "Check Flux status on the local cluster", kivmCheckFluxScript())
	boa.AddScriptCommand(CheckCmd, "heartbeat", "Check heartbeat status on the local cluster", kivmCheckHeartbeatScript())
	boa.AddScriptCommand(CheckCmd, "logs", "Check Logs on the local cluster", kivmCheckLogsScript())
	boa.AddScriptCommand(CheckCmd, "imdb", "Check IMDb app status on the local cluster", kivmCheckImdbScript())
	boa.AddScriptCommand(CheckCmd, "setup", "Check setup progress on the local cluster", kivmCheckSetupScript())
}

func kivmCheckFluxScript() string {
	return "flux get kustomization"
}

func kivmCheckLogsScript() string {
	return "cat /var/log/cloud-init-output.log"
}

func kivmCheckSetupScript() string {
	return "cat ~/status"
}

func kivmCheckHeartbeatScript() string {
	return getUrl("/heartbeat/17")
}

func kivmCheckImdbScript() string {
	return getUrl("/version")
}

func getUrl(path string) string {
	ssl := os.Getenv("AKDC_SSL")
	fqdn := os.Getenv("AKDC_FQDN")

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	cmd := fqdn + path

	if ssl == "" {
		cmd = "http http://" + cmd
	} else {
		cmd = "http https://" + cmd
	}

	return cmd
}
