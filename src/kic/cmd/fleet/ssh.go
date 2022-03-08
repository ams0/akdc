// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package fleet

import (
	"kic/boa"
	"kic/boa/cfmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	// sshCmd opens an ssh terminal on a cluster
	SshCmd = &cobra.Command{
		Use:   "ssh",
		Short: "Open an SSH shell to the cluster",
		Args:  cobra.ExactValidArgs(1),

		ValidArgsFunction: validArgsFleetSsh,

		RunE: runFleetSsh,
	}
)

// run kic fleet ssh command
func runFleetSsh(cmd *cobra.Command, args []string) error {
	// get the ip from the ips file
	hostIPs, err := boa.ReadHostIPs(args[0])

	if err != nil {
		return err
	}

	ip := args[0]

	// try to lookup partial DNS name
	if len(hostIPs) > 0 {
		ip = hostIPs[len(hostIPs)-1]
		cols := strings.Split(ip, "\t")

		if len(cols) > 1 {
			ip = strings.TrimSpace(cols[1])
		}
	}

	if ip != "" {
		boa.ShellExecE("ssh -p 2222 -o \"StrictHostKeyChecking=no\" akdc@" + ip)
	} else {
		cfmt.ErrorE("unable to find host or IP")
	}

	return nil
}

// validate kic fleet ssh arg
func validArgsFleetSsh(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	hostIPs, err := boa.ReadHostIPs("")

	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	// only one arg
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	// sugest from the ips or defaultIPs file
	return hostIPs, cobra.ShellCompDirectiveNoFileComp
}
