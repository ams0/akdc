// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sshCmd opens an ssh terminal on a cluster
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "open an SSH shell to the cluster",
	Long:  `open an SSH shell to the cluster`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("akdc ssh requires a server name from the ips file")
		} else {
			shellExec(fmt.Sprintf("ssh -p 2222 -o \"StrictHostKeyChecking=no\" akdc@$(cat ips | grep %s | tail -1 | cut -f 2)", args[0]))
		}
	},
}
