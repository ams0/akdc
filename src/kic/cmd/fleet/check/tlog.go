// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package check

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// CheckAioaCmd checks each cluster to see if AI Order Accuracy is installed
var TLogCmd = &cobra.Command{
	Use:   "tlog",
	Short: "Check Virtual TLog status on each cluster",
	Args:  argsFleetCheck,

	RunE: func(cmd *cobra.Command, args []string) error {
		return boa.ExecClusters("./fleet-vm/scripts/check-tlog", grep)
	},
}
