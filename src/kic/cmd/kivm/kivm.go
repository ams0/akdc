// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package kivm

import (
	"kic/boa"
	"kic/cmd/test"

	"github.com/spf13/cobra"
)

var KivmCmd = &cobra.Command{
	Use:   "kivm",
	Short: "Kubernetes in VMs CLI",
	Long:  "Kubernetes in VMs CLI\n\n  A CLI for automating many Kubernetes fleet cluster tasks",
}

func LoadCommands(parent *cobra.Command) *cobra.Command {
	boa.AddCommandIfNotExist(parent, CheckCmd)
	boa.AddCommandIfNotExist(parent, ClusterCmd)
	boa.AddCommandIfNotExist(parent, test.TestCmd)

	boa.AddCommandIfNotExist(parent, boa.CreateScriptCommand("az-login", "Log in to Azure using Managed Identity", kivmAzLoginScript()))
	boa.AddCommandIfNotExist(parent, boa.CreateScriptCommand("env", "List the environment variables", kivmEnvScript()))
	boa.AddCommandIfNotExist(parent, boa.CreateScriptCommand("events", "Get all Kubernetes events on the local dev cluster", kivmEventsScript()))
	boa.AddCommandIfNotExist(parent, boa.CreateScriptCommand("pods", "Get all pods on the local dev cluster", kivmPodsScript()))
	boa.AddCommandIfNotExist(parent, boa.CreateScriptCommand("pull", "Pull latest git repos", kivmPullScript()))
	boa.AddCommandIfNotExist(parent, boa.CreateScriptCommand("svc", "Get all services on the local dev cluster", kivmSvcScript()))
	boa.AddCommandIfNotExist(parent, boa.CreateScriptCommand("sync", "Force Flux to sync (reconcile) to the local cluster", kivmSyncScript()))

	return parent
}

func kivmPodsScript() string {

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

func kivmAzLoginScript() string {
	return "az login --identity -o table"
}

func kivmEnvScript() string {

	return "env | grep AKDC | sort"
}

func kivmEventsScript() string {
	return "kubectl get events --all-namespaces --sort-by='.metadata.creationTimestamp'"
}

func kivmPullScript() string {
	return "git -C $HOME/gitops pull"
}

func kivmSvcScript() string {

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

func kivmSyncScript() string {
	return "flux reconcile source git gitops && kubectl get pods -A"
}
