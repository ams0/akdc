// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"fmt"
	"kic/cfmt"
	"kic/utils"
	"strings"

	"github.com/spf13/cobra"
)

// sshCmd opens an ssh terminal on a cluster
var SshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Open an SSH shell to the cluster",
	Args: func(cmd *cobra.Command, args []string) error {
		// this will exit with an error
		utils.ReadHostIPs("")
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("akdc ssh requires a server name from the ips file")
		} else {
			hostIPs := utils.ReadHostIPs(args[0])

			ip := ""

			if len(hostIPs) > 0 {
				ip = hostIPs[len(hostIPs)-1]
				cols := strings.Split(ip, "\t")

				if len(cols) > 1 {
					ip = strings.TrimSpace(cols[1])
				}
			}

			if ip != "" {
				utils.ShellExec("ssh -p 2222 -o \"StrictHostKeyChecking=no\" akdc@" + ip)
			} else {
				cfmt.ExitErrorMessage("unable to find host or IP")
			}
		}
	},
}
