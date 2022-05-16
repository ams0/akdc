// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

// special version of kic for the kubecon hands-on labs

package kubekic

import (
	"kic/boa"
	"kic/cmd/test"

	"github.com/spf13/cobra"
)

var KicCmd = &cobra.Command{
	Use:   "kic",
	Short: "Kubernetes in Codespaces CLI",
	Long:  "Kubernetes in Codespaces CLI\n\n  A CLI for automating many Kubernetes inner-loop tasks",
}

func init() {

	KicCmd.AddCommand(BuildCmd)
	KicCmd.AddCommand(CheckCmd)
	KicCmd.AddCommand(ClusterCmd)
	KicCmd.AddCommand(test.TestCmd)
	//KicCmd.AddCommand(targets.TargetsCmd)

	boa.AddScriptCommand(KicCmd, "pods", "Get all pods on the local dev cluster", kicPodsScript())
	boa.AddScriptCommand(KicCmd, "svc", "Get all services on the local dev cluster", kicSvcScript())
	boa.AddScriptCommand(KicCmd, "events", "Get all Kubernetes events on the local dev cluster", kicEventsScript())
	// boa.AddScriptCommand(KicCmd, "env", "List the environment variables", kicEnvScript())
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

func kicEventsScript() string {

	return "kubectl get events --all-namespaces --sort-by='.metadata.creationTimestamp'"
}

func kicEnvScript() string {

	return "env | grep AKDC | sort"
}
