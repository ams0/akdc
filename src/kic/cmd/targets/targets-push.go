// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package targets

import (
	"kic/cfmt"
	"kic/utils"

	"github.com/spf13/cobra"
)

// checkFluxCmd checks each cluster for flux-check namespace
var PushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push the changes to GitHub",

	Args: func(cmd *cobra.Command, args []string) error {
		return checkForConfigFile()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// make sure the repo is up to date
		if utils.ShellExecOut("git status -s") == "" {
			cfmt.Info("nothing to push")
		} else {
			cfmt.Info("pulling from GitHub")
			utils.ShellExec("git pull")

			cfmt.Info("committing changes to GitHub")
			utils.ShellExec("git commit -am 'updated targets with akdc CLI'")

			cfmt.Info("pushing changes to GitHub")
			utils.ShellExec("git push")
		}
	},
}
