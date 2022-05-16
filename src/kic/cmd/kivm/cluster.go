// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package kivm

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// ListCmd lists the clusters in the fleet
var ClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "k3d cluster commands",
}

func init() {
	boa.AddScriptCommand(ClusterCmd, "reset", "Create and bootstrap a local k3d cluster and deploy the apps", kivmClusterResetScript())
	boa.AddScriptCommand(ClusterCmd, "delete", "Delete the k3d cluster", kivmClusterDeleteScript())
	boa.AddScriptCommand(ClusterCmd, "jumpbox", "Deploy a 'jumpbox' to the local k3d cluster", kivmClusterJumpboxScript())
}

func kivmClusterDeleteScript() string {

	return "k3d cluster delete"
}

func kivmClusterJumpboxScript() string {

	return `
#!/bin/bash

kubectl delete pod jumpbox --ignore-not-found=true
kubectl run jumpbox --image=ghcr.io/cse-labs/jumpbox --restart=Always

`
}

func kivmClusterResetScript() string {

	return `
#!/bin/bash

kivm cluster delete

echo ""
echo "Creating cluster ..."

### todo - check az login

cd $HOME/gitops/vm/setup || exit

./akdc-pre-k8s.sh
./k8s-setup.sh
./akdc-pre-flux.sh
./flux-setup.sh
./akdc-pre-arc.sh
./arc-setup.sh
./akdc-post.sh

`
}
