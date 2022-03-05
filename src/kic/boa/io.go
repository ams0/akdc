// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package boa

import (
	"fmt"
	"kic/boa/cfmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// read key from ~/.kic
func ReadConfigValue(key string) string {
	key = strings.TrimSpace(key)

	if !strings.HasSuffix(key, ":") {
		key += ":"
	}

	config := os.Getenv("HOME") + "/.kic"

	if _, err := os.Stat(config); err == nil {
		lines := ReadLinesFromFile(config)

		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, key) {
				line = strings.TrimSpace(strings.Replace(line, key, "", 1))
				return line
			}
		}
	}

	return ""
}

// read the ips file
func ReadHostIPs(grep string) []string {
	command := ""

	if _, err := os.Stat("./ips"); err != nil {
		file := ReadConfigValue("defaultIPs:")
		if file != "" {
			command = "cat " + file + " | sort"
		}
	} else {
		command = "cat ips | sort"
	}

	if command == "" {
		cfmt.ExitErrorMessage("fleet file not found")
	}

	if grep != "" {
		err := CheckForBadChars(grep, "grep")
		if err != nil {
			cfmt.ExitErrorMessage(err)
		}

		command += " | grep " + grep
	}

	res, _ := ShellExecOut(command)

	lines := strings.Split(res, "\n")

	var ips []string = nil

	for _, line := range lines {
		ip := strings.Split(line, "\t")

		if len(ip) > 1 {
			ips = append(ips, line)
		}
	}

	return ips
}

// get the path to the executable's directory
func GetBinDir() string {

	ex, _ := os.Getwd()

	// return the working directory on tests
	if strings.HasPrefix(ex, "/tmp/") {
		return ex
	}

	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

	// get the parent of bin
	return filepath.Dir(ex)
}

// get the file name from the executing directory
func GetBinName() string {
	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

	// get the parent of bin
	return filepath.Base(ex)
}

// get the path to the commands (i.e. /bin/kic/.kic)
func GetBoaPath() string {
	boaPath := GetBinDir()
	app := GetBinName()

	if strings.HasPrefix(app, "__debug") {
		// running in debugger - assume package name == source directory
		app = filepath.Base(boaPath)
	}

	// complete the paths
	return boaPath + "/." + app + "/"

}

// get the path to the boa commands
func GetBoaCommandPath() string {
	return GetBoaPath() + "commands/"
}

// get the path to the repo base
func GetRepoBase() string {
	base := os.Getenv("REPO_BASE")

	if base == "" {
		ex, err := os.Executable()

		if err != nil {
			log.Fatal(err)
		}

		base = filepath.Dir(ex)
		base = filepath.Dir(base)

		if strings.HasSuffix(base, "src") {
			base = filepath.Dir(base)
		}
	}

	return base
}

// read a text file from the boa directory
// i.e. /bin/kic/.kic
func ReadTextFileFromBoaDir(name string) string {
	path := filepath.Join(GetBoaPath(), name)
	return ReadTextFile(path)
}

// read a file and return the text
func ReadTextFile(path string) string {
	content, err := os.ReadFile(path)

	if err != nil {
		return ""
	}

	return string(content)
}

// read lines from a text file
func ReadLinesFromFile(path string) []string {
	return strings.Split(ReadTextFile(path), "\n")
}

// check for dangerous characters sent to bash
func CheckForBadChars(source string, param string) error {

	if source != "" {
		badChars := "|&;<>"

		for _, ch := range badChars {
			if strings.Contains(source, string(ch)) {
				return fmt.Errorf("invalid character in parameter %s", param)
			}
		}
	}

	return nil
}
