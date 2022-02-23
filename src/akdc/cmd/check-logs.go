// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

// checkLogsCmd represents the dapr command
var checkLogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "check the cloudinit logs on the VMs",
	Long:  `check the cloudinit logs on the VMs`,
	Run: func(cmd *cobra.Command, args []string) {
		execClusters("'echo \"$(hostname)\ncloud-init log\n$(tail /var/log/cloud-init-output.log)\n\nstatus log\n$(tail status)\"'", grep)
	},
}
