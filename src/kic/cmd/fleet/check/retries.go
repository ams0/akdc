// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package check

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// RetriesCmd checks each server for the number of flux retries during installation
var RetriesCmd = &cobra.Command{
	Use:   "retries",
	Short: "Check number of retries on each cluster",
	Args:  argsFleetCheck,

	RunE: func(cmd *cobra.Command, args []string) error {
		return boa.ExecClusters("./fleet-vm/scripts/check-retries", grep)
	},
}
