// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
	"kic/cfmt"
)

// checkCmd adds check subcommands
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check status on the local k3d cluster",
	Long:  `check status on the local k3d cluster`,
}

// add sub commands
func init() {
	checkCmd.AddCommand(checkAllCmd)
	checkCmd.AddCommand(checkGrafanaCmd)
	checkCmd.AddCommand(checkNgsaCmd)
	checkCmd.AddCommand(checkPrometheusCmd)
	checkCmd.AddCommand(checkWebVCmd)
}

// define the commands so we can call them from other commands
var checkGrafanaCmd = addCheckCommand("grafana", "Check Grafana status on the local k3d cluster", 32000, "/healthz")
var checkNgsaCmd = addCheckCommand("ngsa", "Check NGSA status on the local k3d cluster", 30080, "/version")
var checkPrometheusCmd = addCheckCommand("prometheus", "Check Prometheus status on the local k3d cluster", 30000, "/")
var checkWebVCmd = addCheckCommand("webv", "Check WebV status on the local k3d cluster", 30088, "/version")

// checkAllCmd checks all of the end points
var checkAllCmd = &cobra.Command{
	Use:   "all",
	Short: "check status on the local k3d cluster",
	Long:  `check status on the local k3d cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		cfmt.Info("Checking cluster status")
		checkNgsaCmd.Run(cmd, args)
		checkWebVCmd.Run(cmd, args)
		checkGrafanaCmd.Run(cmd, args)
		checkPrometheusCmd.Run(cmd, args)
	},
}
