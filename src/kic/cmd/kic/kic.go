// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package kic

import (
	"kic/boa"
	"kic/cmd/targets"
	"kic/cmd/test"

	"github.com/spf13/cobra"
)

var KicCmd = &cobra.Command{
	Use:   "kic",
	Short: "Kubernetes in Codespaces CLI",
	Long:  "Kubernetes in Codespaces CLI\n\n  A CLI for automating many Kubernetes inner-loop tasks",
}

func AddCommands() *cobra.Command {

	if len(KicCmd.Commands()) == 0 {
		KicCmd.AddCommand(BuildCmd)
		KicCmd.AddCommand(CheckCmd)
		KicCmd.AddCommand(ClusterCmd)
		KicCmd.AddCommand(targets.TargetsCmd)
		KicCmd.AddCommand(test.TestCmd)

		boa.AddScriptCommand(KicCmd, "env", "List the environment variables", kicEnvScript())
		boa.AddScriptCommand(KicCmd, "events", "Get all Kubernetes events on the local dev cluster", kicEventsScript())
		boa.AddScriptCommand(KicCmd, "pods", "Get all pods on the local dev cluster", kicPodsScript())
		boa.AddScriptCommand(KicCmd, "svc", "Get all services on the local dev cluster", kicSvcScript())
	}

	return KicCmd
}

func kicEnvScript() string {

	return "env | grep AKDC | sort"
}

func kicEventsScript() string {

	return "kubectl get events --all-namespaces --sort-by='.metadata.creationTimestamp'"
}

func kicPodsScript() string {

	return `
#!/bin/bash

hdrsort()
{
    read -r
    printf "%s\\n" "$REPLY"
    sort
}

kubectl get pods -A | hdrsort

`
}

func kicSvcScript() string {

	return `
#!/bin/bash

hdrsort()
{
    read -r
    printf "%s\\n" "$REPLY"
    sort
}

kubectl get svc -A | hdrsort

`
}
