// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"akdc/cfmt"
	"github.com/spf13/cobra"
)

// checkFluxCmd checks each cluster for flux-check namespace
var targetsPushCmd = &cobra.Command{
	Use:   "push",
	Short: "push the changes to GitHub",
	Long:  `push the changes to GitHub`,
	Args: func(cmd *cobra.Command, args []string) error {
		return checkForConfigFile()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// make sure the repo is up to date
		if shellExecOut("git status -s") == "" {
			cfmt.Info("nothing to push")
		} else {
			cfmt.Info("pulling from GitHub")
			shellExec("git pull")

			cfmt.Info("committing changes to GitHub")
			shellExec("git commit -am 'updated targets with akdc CLI'")

			cfmt.Info("pushing changes to GitHub")
			shellExec("git push")
		}
	},
}
