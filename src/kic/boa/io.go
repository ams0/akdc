// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package boa

import (
	"fmt"
	"kic/boa/cfmt"
	"log"
	"os"
	"path"
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
func ReadHostIPs(grep string) ([]string, error) {
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
		cfmt.ErrorE("fleet file not found")
	}

	if grep != "" {
		err := CheckForBadChars(grep, "grep")
		if err != nil {
			cfmt.ErrorE(err)
			return nil, err
		}

		command += " | grep " + grep
	}

	res, err := ShellExecOut(command, false)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(res, "\n")

	var ips []string = nil

	for _, line := range lines {
		ip := strings.Split(line, "\t")

		if len(ip) > 1 {
			ips = append(ips, line)
		}
	}

	return ips, nil
}

// read a completion file
func ReadCompletionFile(fileName string) ([]string, error) {
	path := path.Join(GetBoaPath(), fileName)

	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	return ReadLinesFromFile(path), nil
}

// get the path to the executable's directory
func GetBinDir() string {
	// read from env var
	ex := os.Getenv("KIC_PATH")
	if ex != "" {
		return ex
	}

	ex, _ = os.Getwd()

	// return the working directory on tests
	if strings.HasPrefix(ex, "/tmp/") || strings.HasPrefix(GetBinName(), "__debug") {
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
	// read from env var
	ex := os.Getenv("KIC_NAME")
	if ex != "" {
		return ex
	}

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
	appConfig := "." + app

	// check current directory first
	if local, err := os.Getwd(); err == nil {
		local = filepath.Join(local, appConfig)

		if _, err := os.Stat(local); err == nil {
			return local
		}
	}

	// running in debugger
	if strings.HasPrefix(app, "__debug") {
		// assume package name == source directory
		app = filepath.Base(boaPath)
		appConfig = "." + app

		if _, err := os.Stat(appConfig); err != nil {
			// walk the path to find the first bin dir
			tpath := filepath.Dir(boaPath)
			_, err := os.Stat(filepath.Join(tpath, "bin", appConfig))

			for err != nil && tpath != "/" {
				tpath = filepath.Dir(tpath)
				_, err = os.Stat(filepath.Join(tpath, "bin", appConfig))
			}

			if tpath != "/" {
				boaPath = filepath.Join(tpath, "bin")
			}
		}
	}

	// complete the path
	return filepath.Join(boaPath, appConfig)
}

// get the path to the boa commands
func GetBoaCommandPath() string {
	return filepath.Join(GetBoaPath(), "commands")
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
