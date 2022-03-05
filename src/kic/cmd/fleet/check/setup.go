// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package check

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// SetupCmd checks each cluster for the setup status
var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Check setup status on each cluster",

	Args: func(cmd *cobra.Command, args []string) error {
		// this will exit with an error
		boa.ReadHostIPs("")
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		// don't use a command on the VM as it's not available until late in setup
		boa.ExecClusters("tail -n1 status", grep)
	},
}
