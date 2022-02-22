// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"
	"io/ioutil"
	"kic/cfmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command and adds commands, options, and flags
var rootCmd = &cobra.Command{
	Use:   "kic",
	Short: "Kubernetes in Codespaces CLI",
	Long:  `Kubernetes in Codespaces CLI`,
}

// initialize the root command
func init() {
	rootCmd.AddCommand(allCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(cleanCmd)
	rootCmd.AddCommand(jumpboxCmd)
	rootCmd.AddCommand(appCmd)
	rootCmd.AddCommand(webvCmd)

	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(testCmd)

	loadModules()
}

// todo - design file format / metadata
func loadModules() {
	// modules are in the kic bin/.kic/modules directory
	path := getParentDir() + "/.kic/modules/"

	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, f := range files {
			var moduleCmd *cobra.Command

			p := path + "/" + f.Name() + "/" + f.Name()

			txt := readTextFile(p)

			lines := strings.Split(txt, "\n")

			for i := 0; i < len(lines); i++ {
				line := lines[i]

				if !strings.HasPrefix(line, "#") {
					if strings.HasPrefix(line, "kicModule") {
						if i+1 < len(lines) {
							name := strings.TrimSpace(lines[i+1])
							i++

							if i+1 < len(lines) {
								description := strings.TrimSpace(lines[i+1])
								i++
								moduleCmd = addParentCommand(name, "[module] "+description)
								rootCmd.AddCommand(moduleCmd)
							}
						}
					} else if strings.HasPrefix(line, "kicCommand") {
						if i+1 < len(lines) {
							name := strings.TrimSpace(lines[i+1])
							i++

							if i+1 < len(lines) {
								description := strings.TrimSpace(lines[i+1])
								i++
								if i+1 < len(lines) {
									command := strings.TrimSpace(lines[i+1])
									i++
									if moduleCmd != nil {
										subCmd := addRunCommand(name, description, command)
										moduleCmd.AddCommand(subCmd)
									}
								}
							}
						}
					} else {
						fmt.Println(line)
					}
				}
			}
		}
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// execute the root command
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// define the commands (so we can call them as sub commands)
// these commands execute the bash script in bin/.kic/commands/name
var allCmd = addExecCommand("all", "Create and bootstrap a local k3d cluster and deploy the apps")
var createCmd = addExecCommand("create", "Create a new local k3d cluster")
var deleteCmd = addExecCommand("delete", "Delete the local k3d cluster (if exists)")
var deployCmd = addExecCommand("deploy", "Deploy the apps to the local k3d cluster")
var cleanCmd = addExecCommand("clean", "Remove the apps from the local k3d cluster")
var jumpboxCmd = addExecCommand("jumpbox", "Deploy a 'jumpbox' to the local k3d cluster")
var appCmd = addExecCommand("app", "Build and deploy a local NGSA docker image")
var webvCmd = addExecCommand("webv", "Build and deploy a local WebV docker image")

func addParentCommand(name string, description string) *cobra.Command {
	parentCmd := &cobra.Command{
		Use:   name,
		Short: description,
		Long:  description,
	}

	return parentCmd
}

// create a command that runs the bash command(s)
func addRunCommand(name string, description string, command string) *cobra.Command {
	runCmd := &cobra.Command{
		Use:   name,
		Short: description,
		Long:  description,

		Run: func(cmd *cobra.Command, args []string) {
			cfmt.Info(description)
			shellExec(command)
		},
	}

	return runCmd
}
