// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	// "fmt"
	"io/ioutil"
	// "kic/cfmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command and adds commands, options, and flags
var rootCmd = &cobra.Command{
	Use:   "kic",
	Short: "Kubernetes in Codespaces CLI",
}

// initialize the root command
func init() {
	// this can't be a module [easily]
	// todo - this is different between kic and kic-vm
	rootCmd.AddCommand(testCmd)

	// todo - check for duplicate root commands
	// todo - override?
	loadModules()
}

// todo - design file format / metadata
// todo - move the module commands to ../commands (makes chaining easier)
// todo - for simple commands, should we leave them in the .mod file? - can't chain easily

func loadModules() {
	// modules are in the kic bin/.kic/modules directory
	path := getParentDir() + "/.kic/modules/"
	var moduleCmd *cobra.Command

	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, f := range files {
			moduleCmd = nil

			p := path + "/" + f.Name()

			txt := readTextFile(p)

			lines := strings.Split(txt, "\n")

			for i := 0; i < len(lines); i++ {
				line := lines[i]

				// ignore comments
				if !strings.HasPrefix(line, "#") {
					if strings.HasPrefix(line, "kicModule") {
						if i+1 < len(lines) {
							name := strings.TrimSpace(lines[i+1])
							i++

							if i+1 < len(lines) {
								description := strings.TrimSpace(lines[i+1])
								i++
								moduleCmd = addParentCommand(name, description)
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
										moduleCmd.AddCommand(addRunCommand(name, description, command))
									} else {
										rootCmd.AddCommand(addRunCommand(name, description, command))
									}
								}
							}
						}
					} else if strings.HasPrefix(line, "kicRunCommand") {
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
										moduleCmd.AddCommand(addExecCommand(name, description, command))
									} else {
										rootCmd.AddCommand(addExecCommand(name, description, command))
									}
								}
							}
						}
					} else {
						// fmt.Println(line)
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

// add a command that has sub-commands
func addParentCommand(name string, description string) *cobra.Command {
	parentCmd := &cobra.Command{
		Use:   name,
		Short: description,
	}

	return parentCmd
}

// create a command that runs the bash command(s)
func addRunCommand(name string, description string, command string) *cobra.Command {
	runCmd := &cobra.Command{
		Use:   name,
		Short: description,

		Run: func(cmd *cobra.Command, args []string) {
			shellExec(command)
		},
	}

	return runCmd
}

// create a command that executes a bash script from the bin/.kic/commands directory
func addExecCommand(name string, description string, command string) *cobra.Command {
	execCmd := &cobra.Command{
		Use:   name,
		Short: description,

		Run: func(cmd *cobra.Command, args []string) {
			execCommand(command)
		},
	}

	return execCmd
}
