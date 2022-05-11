// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package boa

import (
	"fmt"
	"kic/boa/cfmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var grep string

// get a command by Use (Name)
func GetCommandByUse(cmd *cobra.Command, use string) *cobra.Command {
	if cmd != nil && len(cmd.Commands()) > 0 {
		for _, c := range cmd.Commands() {
			if c.Use == use {
				return c
			}
		}
	}
	return nil
}

// add a command that has sub-commands
func CreateCommand(use string, short string, long string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
	}

	return cmd
}

// create a command that runs the bash command across the fleet
func AddFltCommand(use string, short string, long string, command string) *cobra.Command {
	fltCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,

		RunE: func(cmd *cobra.Command, args []string) error {
			// VM cli could be in ./cli (defalut) or ./gitops
			script := os.Getenv("AKDC_VM_REPO")

			if script == "" {
				script = "gitops"
			}

			script = filepath.Join(".", script, "vm", "scripts", command)

			// add the paramaters
			if len(args) > 0 {
				script += " " + strings.Join(args, " ")
			}

			// check-setup can run before the cli is fully setup
			if command == "check-setup" {
				script = "tail -n1 status"
			}

			return ExecClusters(script, grep)
		},
	}

	fltCmd.PersistentFlags().StringVarP(&grep, "grep", "g", "", "grep conditional to filter by host")

	return fltCmd
}

// add a script command to a command in the command tree
func AddScriptCommand(parent *cobra.Command, name string, short string, script string) error {
	if script == "" {
		cfmt.ErrorE("script is required", name, short)
		os.Exit(1)
	}

	// name and short are required
	if name == "" || short == "" {
		return fmt.Errorf("name and short are required")
	}

	// add the command
	parent.AddCommand(addScriptWorker(name, short, script))

	return nil
}

// create a command that runs the script
func addScriptWorker(use string, short string, command string) *cobra.Command {
	runCmd := &cobra.Command{
		Use:   use,
		Short: short,

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				for i := len(args); i > 0; i-- {
					command = strings.ReplaceAll(command, "$"+strconv.Itoa(i), args[i-1])
				}
			}

			return ShellExecE(command)
		},
	}

	return runCmd
}

// run a command on all clusters
func ExecClusters(cmd string, grep string) error {
	hostIPs, err := ReadHostIPs(grep)

	if err != nil {
		return err
	}

	ch := make(chan string)

	for _, hostIP := range hostIPs {
		cols := strings.Split(hostIP, "\t")

		if len(cols) > 1 {
			go ExecCluster(cols[0], cols[1], cmd, ch)
		}
	}

	// todo - add timeout
	for i := 0; i < len(hostIPs); i++ {
		<-ch
	}

	return nil
}

// run a command on one cluster via ssh
func ExecCluster(host string, ip string, cmd string, ch chan string) {
	cmd = fmt.Sprintf("ssh -p 2222 -o \"StrictHostKeyChecking=no\" -o ConnectTimeout=5 akdc@%s %s", ip, cmd)

	res, _ := ShellExecOut(cmd, true)

	if !strings.HasSuffix(res, "\n") {
		res += "\n"
	}

	if strings.HasPrefix(strings.TrimSpace(res), host) {
		fmt.Print(res)
	} else {
		fmt.Printf("%-25s %s", host, res)
	}

	ch <- host
}
