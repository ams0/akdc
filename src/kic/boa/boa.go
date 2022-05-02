// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package boa

import (
	"fmt"
	"io/ioutil"
	"kic/boa/cfmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var boaPath string
var boaCommandPath string

var boaRootCmd *cobra.Command

// load commands from *.boa
// commands are in the .binName directory
//    that is a subdirectory of where the bin file is located
// example: /bin/.kic
func LoadCommands(appRootCmd *cobra.Command) {
	// set boaRootCmd to the app root command
	boaRootCmd = appRootCmd

	boaPath = GetBoaPath()
	boaCommandPath = GetBoaCommandPath()

	// load boaPath/*.boa
	files, err := ioutil.ReadDir(boaPath)
	if err == nil {
		for _, f := range files {
			if !f.IsDir() && strings.HasSuffix(strings.ToLower(f.Name()), ".boa") {
				loadCommand(f.Name())
			}
		}
	}
}

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

// reset cmdRoot to new__root command if set
// otherwise remove any hidden root commands
func SetNewRoot() *cobra.Command {
	var cmd *cobra.Command

	// create a new root command
	cmd = &cobra.Command{Use: boaRootCmd.Use, Short: boaRootCmd.Short, Long: boaRootCmd.Long}

	// check for new__root
	nr := GetCommandByUse(boaRootCmd, "new__root")

	if nr == nil {
		// add all non-hidden commands to the new root
		for _, c := range boaRootCmd.Commands() {
			if !c.Hidden {
				cmd.AddCommand(c)
			}
			boaRootCmd = cmd
		}
	} else {
		// get the new root
		cmd = GetCommandByUse(boaRootCmd, nr.Short)

		if cmd == nil {
			// new__root not found
			cfmt.ErrorE("New root command not found", nr.Short)
			os.Exit(1)
		}

		// create the new boaRootCmd
		// we have to do this as the parent is set to the existing boaRootCmd
		nr = &cobra.Command{Use: boaRootCmd.Use, Short: cmd.Short, Long: cmd.Long}

		for _, c := range cmd.Commands() {
			// add all commands that aren't hidden
			if !c.Hidden {
				nr.AddCommand(c)
			}
		}
		boaRootCmd = nr
	}

	return boaRootCmd
}

// return the stop word for file reads
func getStopWord(line string) string {

	chk := strings.ToLower(line)

	if strings.HasPrefix(chk, "root") {
		return "root"
	}
	if strings.HasPrefix(chk, "command") {
		return "command"
	}
	if strings.HasPrefix(chk, "runcommand") {
		return "runCommand"
	}
	if strings.HasPrefix(chk, "fltcommand") {
		return "fltCommand"
	}
	if strings.HasPrefix(chk, "popcommand") {
		return "popCommand"
	}

	return ""
}

// load command(s) from a file
func loadCommand(fileName string) {
	var boaCmd *cobra.Command

	// todo - convert to struct
	var modType string
	var name string
	var short string
	var long string
	var path string
	var parent string
	var hidden bool

	// read file into an array
	lines := ReadLinesFromFile(filepath.Join(boaPath, fileName))

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		line = strings.Replace(line, "\\n", "\n", -1)

		// ignore comments
		if !strings.HasPrefix(line, "#") {
			chk := strings.ToLower(line)

			// check for stop word
			sw := getStopWord(chk)
			if sw != "" {
				if strings.HasPrefix(sw, "popCommand") {
					// popCommand resets the command to the parent or boaRootCmd
					addBoaCommand(fileName, boaCmd, modType, name, short, long, path, parent, hidden)
					if boaCmd != nil {
						boaCmd = boaCmd.Parent()
						if boaCmd == boaRootCmd {
							boaCmd = nil
						}
					}
				} else {
					// add the command and set the new type based on stopWord
					boaCmd = addBoaCommand(fileName, boaCmd, modType, name, short, long, path, parent, hidden)
				}

				// reset params
				modType = sw
				name = ""
				parent = ""
				short = ""
				long = ""
				path = ""
				hidden = false
			} else if strings.HasPrefix(strings.ToLower(line), "name:") {
				name = strings.TrimSpace(line[5:])
			} else if strings.HasPrefix(strings.ToLower(line), "parent:") {
				parent = strings.TrimSpace(line[7:])
			} else if strings.HasPrefix(strings.ToLower(line), "short:") {
				short = strings.TrimSpace(line[6:])
			} else if strings.HasPrefix(strings.ToLower(line), "long:") {
				long = strings.TrimSpace(line[5:])
			} else if strings.HasPrefix(strings.ToLower(line), "path:") {
				path = strings.TrimSpace(line[5:])
			} else if strings.HasPrefix(strings.ToLower(line), "hidden:") {
				hidden = strings.ToLower(strings.TrimSpace(line[7:])) == "true"
			} else {
				if line != "" {
					cfmt.ErrorE("unrecognized line: " + line)
				}
			}
		}
	}

	// handle last command at EOF
	if modType != "" {
		addBoaCommand(fileName, boaCmd, modType, name, short, long, path, parent, hidden)
	}
}

// add a command or command to cobra
func addBoaCommand(fileName string, modCmd *cobra.Command, modType string, name string, short string, long string, path string, parent string, hidden bool) *cobra.Command {
	// ignore if modType not set
	if modType == "" || modType == "popCommand" {
		return modCmd
	}

	// handle different command and command types
	if modType == "root" {
		return setRootValues(name, short, long)
	} else if modType == "command" {
		return addParentCommand(modCmd, name, short, long, parent, hidden)
	} else if modType == "runCommand" || modType == "fltCommand" {
		return addCommand(modCmd, modType, name, short, long, path, hidden)
	}

	// bad input file
	cfmt.ErrorE("unrecognized Command in file:", fileName, modType, name, short, long, path)
	os.Exit(1)
	return nil
}

// add a command to the command tree
func addParentCommand(modCmd *cobra.Command, name string, short string, long string, parent string, hidden bool) *cobra.Command {
	if !hidden {
		// name and short are required
		if err := checkNameAndShort(name, short); err != nil {
			return modCmd
		}
	} else {
		short = "hidden"
		long = ""
	}

	// create the new command
	boaCmd := createCommand(name, short, long)
	boaCmd.Hidden = hidden

	if parent != "" {
		// set the parent if specified
		if strings.ToLower(parent) == "boaRootCmd" {
			modCmd = nil
		} else {
			modCmd = GetCommandByUse(boaRootCmd, parent)
			if modCmd == nil {
				cfmt.ErrorE("Parent command not found", parent)
				os.Exit(1)
			}
		}
	}

	if modCmd != nil {
		// check for dupes
		if GetCommandByUse(modCmd, name) != nil {
			if hidden {
				GetCommandByUse(modCmd, name).Hidden = true
			} else {
				cfmt.ErrorE("Command already exists", modCmd.Use, name)
				os.Exit(1)
			}
		} else {
			modCmd.AddCommand(boaCmd)
		}
	} else {
		// check for dupes
		if GetCommandByUse(boaRootCmd, name) != nil {
			if hidden {
				GetCommandByUse(boaRootCmd, name).Hidden = true
			} else {
				cfmt.ErrorE("Command already exists", boaRootCmd.Use, name)
				os.Exit(1)
			}
		} else {
			boaRootCmd.AddCommand(boaCmd)
		}
	}

	// set the new command to the parent
	return boaCmd
}

// add a command to a command in the command tree
func addCommand(modCmd *cobra.Command, modType string, name string, short string, long string, path string, hidden bool) *cobra.Command {
	if path == "" {
		cfmt.ErrorE("path is required", name, short, long, path)
		os.Exit(1)
	}

	// read the values from the command metadata if necessary
	name = readFromCommandFile(path, "name", name)
	short = readFromCommandFile(path, "short", short)
	long = readFromCommandFile(path, "long", long)

	if !hidden {
		// name and short are required
		if err := checkNameAndShort(name, short); err != nil {
			return modCmd
		}
	} else {
		short = "hidden"
		long = ""
	}

	// use boaRootCmd if modCmd is nil
	aCmd := modCmd
	if aCmd == nil {
		aCmd = boaRootCmd
	}

	// check for dupes
	if GetCommandByUse(aCmd, name) != nil {
		cfmt.ErrorE("Command already exists", aCmd.Use, name)
		os.Exit(1)
	}

	// add the command
	var cmd *cobra.Command

	if modType == "runCommand" {
		cmd = addRunCommand(name, short, long, path)
	} else if modType == "fltCommand" {
		cmd = addFltCommand(name, short, long, path)
	}

	cmd.Hidden = hidden
	aCmd.AddCommand(cmd)

	// parent doesn't change
	return modCmd
}

// set the root command values
func setRootValues(name string, short string, long string) *cobra.Command {
	// name and short are required
	if err := checkNameAndShort(name, short); err == nil {
		boaRootCmd.Use = name
		boaRootCmd.Short = short

		// this will default to short if not set
		if long != "" {
			boaRootCmd.Long = long
		}
	}

	// reset parent
	return nil
}

// this will exit if name or short are invalid
func checkNameAndShort(name string, short string) error {
	// name and short are required
	if name == "" || short == "" {
		if name == "" {
			return fmt.Errorf("name: is required")
		}
		if short == "" {
			return fmt.Errorf("short: is required")
		}
	}

	return nil
}

// read the metadata from the command
func readFromCommandFile(path string, key string, value string) string {
	// don't read if already set
	if value != "" {
		return value
	}

	key = strings.TrimSpace(strings.ToLower(key)) + ":"
	p := filepath.Join(boaCommandPath, path)

	// read the file into an array
	txt := ReadTextFile(p)
	lines := strings.Split(txt, "\n")

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		if strings.HasPrefix(line, "#") {
			// check all comments for metadata
			line = strings.TrimSpace(strings.TrimLeft(line, "#"))

			if strings.HasPrefix(strings.ToLower(line), key) {
				// extract the metadata
				line = strings.TrimSpace(line[len(key):])
				return line
			}
		}
	}

	return ""
}

// add a command that has sub-commands
func createCommand(use string, short string, long string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
	}

	return cmd
}

// create a command that runs the bash command
func addRunCommand(use string, short string, long string, command string) *cobra.Command {
	runCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				ShellExecArgsE(filepath.Join(boaCommandPath, command), args)
			} else {
				ShellExecE(filepath.Join(boaCommandPath, command))
			}
		},
	}

	return runCmd
}

var grep string

// create a command that runs the bash command across the fleet
func addFltCommand(use string, short string, long string, command string) *cobra.Command {
	fltCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,

		RunE: func(cmd *cobra.Command, args []string) error {
			// VM cli could be in ./cli (defalut) or ./gitops
			script := os.Getenv("AKDC_VM_REPO")

			if script == "" {
				script = "cli"
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

// create a command that runs the bash command
func AddRunCommand(use string, short string, long string, command string) *cobra.Command {
	cmdPath := GetBoaCommandPath()

	runCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				ShellExecArgsE(cmdPath+command, args)
			} else {
				ShellExecE(filepath.Join(cmdPath, command))
			}
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

	ShellExecE(cmd)

	ch <- host
}

// execute a command in bin/.kic/commands
func ExecCommandE(cmd string) error {
	path := GetBoaCommandPath() + cmd

	// execute the file with "bash -c" if it exists
	_, err := os.Stat(path)

	if err == nil {
		cfmt.Info("Running command: " + cmd)

		err = ShellExecE(fmt.Sprintf("%s %s", path, os.Args))
	}

	return err
}
