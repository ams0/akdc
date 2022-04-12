// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"github.com/spf13/cobra"
)

var (
	// option variables
	grep string

	// FleetCmd adds the kic fleet command tree
	FleetCmd = &cobra.Command{
		Use:   "fleet",
		Short: "Commands for fleet of clusters",
	}
)

func init() {
	FleetCmd.AddCommand(CreateCmd)
	FleetCmd.AddCommand(DeleteCmd)
	FleetCmd.AddCommand(ExecCmd)
	FleetCmd.AddCommand(ListCmd)
	FleetCmd.AddCommand(SshCmd)
}
