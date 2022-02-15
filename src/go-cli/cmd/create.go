// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
)

// createCmd creates a new cluster
// todo - implement
var cluster string
var group string
var location string
var repo string
var zone string
var pem string
var key string
var quiet bool
var ssl bool

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new cluster",
	Long:  `create a new cluster`,

	Args: func(cmd *cobra.Command, args []string) error {
		if ssl && zone == "" {
			return fmt.Errorf("you must specify --zone to use --ssl")
		}

		if ssl && pem == "" {
			return fmt.Errorf("you must specify --pem to use --ssl")
		}

		if ssl && key == "" {
			return fmt.Errorf("you must specify --key to use --ssl")
		}

		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {

		ex, err := os.Executable()

		if err != nil {
			log.Fatal(err)
		}

		// build the command line
		command := path.Join(path.Dir(ex), "../create-cluster/create-cluster") + " --cluster " + cluster

		if quiet {
			command += " --quiet"
		}

		if group != "" {
			command += " --group " + group
		}

		if location != "" {
			command += " --location " + location
		}

		if repo != "" {
			command += " --repo " + repo
		}

		if zone != "" {
			command += " --zone " + zone
		}

		if ssl {
			command += " --ssl"
			if pem != "" {
				command += " --pem " + pem
			}
			if key != "" {
				command += " --key " + key
			}
		}

		// execute the command
		shellExec(command)
	},
}

func init() {
	createCmd.Flags().StringVarP(&cluster, "cluster", "c", "", "Kubernetes cluster name (required)")
	createCmd.MarkFlagRequired("cluster")
	createCmd.Flags().StringVarP(&group, "group", "g", "", "Azure resource group name")
	createCmd.Flags().StringVarP(&location, "location", "l", "centralus", "Azure location")
	createCmd.Flags().StringVarP(&repo, "repo", "r", "retaildevcrews/edge-gitops", "GitOps repo name")
	createCmd.Flags().StringVarP(&zone, "zone", "z", "", "DNS domain name")
	createCmd.Flags().BoolVarP(&ssl, "ssl", "s", false, "Use SSL cert (must specify --zone)")
	createCmd.Flags().StringVarP(&pem, "pem", "p", "~/.ssh/certs.pem", "Path to SSL .pem file")
	createCmd.Flags().StringVarP(&key, "key", "k", "~/.ssh/certs.key", " Path to SSL .key file")
	createCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Quiet mode")
}
