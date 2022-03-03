// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package cmd

import (
	"fmt"
	"io/ioutil"
	"kic/cfmt"
	"kic/utils"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var modPath string
var cmdPath string

// reset cmdRoot to new__root module if set
// otherwise remove any hidden root commands
func setNewRoot() {
	var cmd *cobra.Command

	// create a new root command
	cmd = &cobra.Command{Use: rootCmd.Use, Short: rootCmd.Short, Long: rootCmd.Long}

	// check for new__root
	nr := getCommandByUse(rootCmd, "new__root")

	if nr == nil {
		// add all non-hidden commands to the new root
		for _, c := range rootCmd.Commands() {
			if !c.Hidden {
				cmd.AddCommand(c)
			}
			rootCmd = cmd
		}
	} else {
		// get the new root
		cmd = getCommandByUse(rootCmd, nr.Short)

		if cmd == nil {
			// new__root not found
			cfmt.ExitErrorMessage("New root command not found", nr.Short)
		}

		// create the new rootCmd
		// we have to do this as the parent is set to the existing rootCmd
		nr = &cobra.Command{Use: rootCmd.Use, Short: cmd.Short, Long: cmd.Long}

		for _, c := range cmd.Commands() {
			// add all commands that aren't hidden
			if !c.Hidden {
				nr.AddCommand(c)
			}
		}
		rootCmd = nr
	}
}

// get a command by Use (Name)
func getCommandByUse(cmd *cobra.Command, use string) *cobra.Command {
	if cmd != nil && len(cmd.Commands()) > 0 {
		for _, c := range cmd.Commands() {
			if c.Use == use {
				return c
			}
		}
	}
	return nil
}

// example of traversing the command tree
// prints out a tree
// not used
func traverseCommands(cmd *cobra.Command, indent string) {
	fmt.Println(indent + cmd.Use)
	for _, c := range cmd.Commands() {
		if len(c.Commands()) > 0 {
			traverseCommands(c, indent+"   ")
		} else {
			fmt.Println(indent + "   " + c.Use + "  --  " + c.Short)
		}
	}
}

// load modules from *.boa
// modules are in the .binName directory
//    that is a subdirectory of where the bin file is located
// example: /bin/.kic
func loadModules() {
	modPath = utils.GetBinDir()
	app := utils.GetBinName()

	if strings.HasPrefix(app, "__debug") {
		// running in debugger - assume package name == source directory
		app = filepath.Base(modPath)
	}

	// complete the paths
	modPath += "/." + app + "/"
	cmdPath = modPath + "commands/"

	// load modPath/*.boa
	files, err := ioutil.ReadDir(modPath)
	if err == nil {
		for _, f := range files {
			if !f.IsDir() && strings.HasSuffix(strings.ToLower(f.Name()), ".boa") {
				loadModule(f.Name())
			}
		}
	}
}

// return the stop word for file reads
func getStopWord(line string) string {

	chk := strings.ToLower(line)

	if strings.HasPrefix(chk, "root") {
		return "root"
	}
	if strings.HasPrefix(chk, "module") {
		return "module"
	}
	if strings.HasPrefix(chk, "runcommand") {
		return "runCommand"
	}
	if strings.HasPrefix(chk, "popcommand") {
		return "popCommand"
	}

	return ""
}

// load module(s) from a file
func loadModule(fileName string) {
	var moduleCmd *cobra.Command

	// todo - convert to struct
	var modType string
	var name string
	var short string
	var long string
	var path string
	var parent string
	var hidden bool

	// read file into an array
	lines := utils.ReadLinesFromFile(modPath + fileName)

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
					// popCommand resets the command to the parent or rootCmd
					addModuleOrCommand(fileName, moduleCmd, modType, name, short, long, path, parent, hidden)
					if moduleCmd != nil {
						moduleCmd = moduleCmd.Parent()
						if moduleCmd == rootCmd {
							moduleCmd = nil
						}
					}
				} else {
					// add the module and set the new type based on stopWord
					moduleCmd = addModuleOrCommand(fileName, moduleCmd, modType, name, short, long, path, parent, hidden)
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
					cfmt.Error("unrecognized line: " + line)
				}
			}
		}
	}

	// handle last module at EOF
	if modType != "" {
		addModuleOrCommand(fileName, moduleCmd, modType, name, short, long, path, parent, hidden)
	}
}

// add a module or command to cobra
func addModuleOrCommand(fileName string, modCmd *cobra.Command, modType string, name string, short string, long string, path string, parent string, hidden bool) *cobra.Command {
	// ignore if modType not set
	if modType == "" || modType == "popCommand" {
		return modCmd
	}

	// handle different module and command types
	if modType == "root" {
		return setRootValues(name, short, long)
	} else if modType == "module" {
		return addModule(modCmd, name, short, long, parent, hidden)
	} else if modType == "runCommand" {
		return addCommand(modCmd, name, short, long, path, hidden)
	}

	// bad input file
	cfmt.ExitErrorMessage("unrecognized Command in file:", fileName, modType, name, short, long, path)
	return nil
}

// add a module to the command tree
func addModule(modCmd *cobra.Command, name string, short string, long string, parent string, hidden bool) *cobra.Command {
	if !hidden {
		// name and short are required
		checkNameAndShort(name, short)
	} else {
		short = "hidden"
		long = ""
	}

	// create the new module
	moduleCmd := createCommand(name, short, long)
	moduleCmd.Hidden = hidden

	if parent != "" {
		// set the parent if specified
		if strings.ToLower(parent) == "rootcmd" {
			modCmd = nil
		} else {
			modCmd = getCommandByUse(rootCmd, parent)
			if modCmd == nil {
				cfmt.ExitErrorMessage("Parent command not found", parent)
			}
		}
	}

	if modCmd != nil {
		// check for dupes
		if getCommandByUse(modCmd, name) != nil {
			if hidden {
				getCommandByUse(modCmd, name).Hidden = true
			} else {
				cfmt.ExitErrorMessage("Command already exists", modCmd.Use, name)
			}
		} else {
			modCmd.AddCommand(moduleCmd)
		}
	} else {
		// check for dupes
		if getCommandByUse(rootCmd, name) != nil {
			if hidden {
				getCommandByUse(rootCmd, name).Hidden = true
			} else {
				cfmt.ExitErrorMessage("Command already exists", rootCmd.Use, name)
			}
		} else {
			rootCmd.AddCommand(moduleCmd)
		}
	}

	// set the new module to the parent
	return moduleCmd
}

// add a command to a module in the command tree
func addCommand(modCmd *cobra.Command, name string, short string, long string, path string, hidden bool) *cobra.Command {
	if path == "" {
		cfmt.ExitErrorMessage("path is required", name, short, long, path)
	}

	// read the values from the command metadata if necessary
	name = readFromCommandFile(path, "name", name)
	short = readFromCommandFile(path, "short", short)
	long = readFromCommandFile(path, "long", long)

	if !hidden {
		// name and short are required
		checkNameAndShort(name, short)
	} else {
		short = "hidden"
		long = ""
	}

	// use rootCmd if modCmd is nil
	aCmd := modCmd
	if aCmd == nil {
		aCmd = rootCmd
	}

	// check for dupes
	if getCommandByUse(aCmd, name) != nil {
		cfmt.ExitErrorMessage("Command already exists", aCmd.Use, name)
	}

	// add the command
	cmd := addRunCommand(name, short, long, path)
	cmd.Hidden = hidden
	aCmd.AddCommand(cmd)

	// parent doesn't change
	return modCmd
}

// set the root command values
func setRootValues(name string, short string, long string) *cobra.Command {
	// name and short are required
	checkNameAndShort(name, short)

	rootCmd.Use = name
	rootCmd.Short = short

	// this will default to short if not set
	if long != "" {
		rootCmd.Long = long
	}

	// reset parent
	return nil
}

// this will exit if name or short are invalid
func checkNameAndShort(name string, short string) {
	// name and short are required
	if name == "" || short == "" {
		if name == "" {
			cfmt.Error("name: is required")
		}
		if short == "" {
			cfmt.Error("short: is required")
		}

		os.Exit(1)
	}
}

// read the metadata from the command
func readFromCommandFile(path string, key string, value string) string {
	// don't read if already set
	if value != "" {
		return value
	}

	key = strings.TrimSpace(strings.ToLower(key)) + ":"
	p := cmdPath + path

	// read the file into an array
	txt := utils.ReadTextFile(p)
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
				utils.ShellExecArgs(cmdPath+command, args)
			} else {
				utils.ShellExec(cmdPath + command)
			}
		},
	}

	return runCmd
}
