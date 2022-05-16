// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

package kic

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
	boa.AddScriptCommand(ClusterCmd, "rebuild", "Create and bootstrap a local k3d cluster and deploy the apps", kicClusterRebuildScript())
	boa.AddScriptCommand(ClusterCmd, "create", "Create a new local k3d cluster", kicClusterCreateScript())
	boa.AddScriptCommand(ClusterCmd, "delete", "Delete the k3d cluster", kicClusterDeleteScript())
	boa.AddScriptCommand(ClusterCmd, "jumpbox", "Deploy a 'jumpbox' to the local k3d cluster", kicClusterJumpboxScript())
	boa.AddScriptCommand(ClusterCmd, "clean", "Remove the apps from the local k3d cluster", kicClusterCleanScript())
	boa.AddScriptCommand(ClusterCmd, "deploy", "Deploy the apps to the local k3d cluster", kicClusterDeployScript())
}

func kicClusterCreateScript() string {

	return `
#!/bin/bash

# validate directories
if [ ! -f "$REPO_BASE/.devcontainer/k3d.yaml" ]; then echo "File not found - \"$REPO_BASE\"/.devcontainer/k3d.yaml"; exit 1; fi

kic cluster delete

echo ""
echo "Creating cluster ..."

k3d cluster create \
    --registry-use k3d-registry.localhost:5500 \
    --k3s-server-arg '--no-deploy=traefik' \
    --config "$REPO_BASE/.devcontainer/k3d.yaml"

`
}

func kicClusterDeleteScript() string {

	return "k3d cluster delete"
}

func kicClusterJumpboxScript() string {

	return `
#!/bin/bash

kubectl delete pod jumpbox --ignore-not-found=true
kubectl run jumpbox --image=ghcr.io/cse-labs/jumpbox --restart=Always

`
}

func kicClusterCleanScript() string {

	return `
#!/bin/bash

kubectl delete ns imdb --ignore-not-found=true
kubectl delete pod jumpbox --ignore-not-found=true

`
}

func kicClusterDeployScript() string {

	return `
#!/bin/bash

# validate directories
if [ ! -d ./deploy ]; then echo "./deploy directory not found. Please cd to an appropriate directory"; exit 1; fi
if [ ! -d ./deploy/webv ]; then echo "./deploy/webv directory not found. Please cd to an appropriate directory"; exit 1; fi
if [ ! -d ./deploy/imdb ]; then echo "./deploy/imdb directory not found. Please cd to an appropriate directory"; exit 1; fi
if [ ! -d ./deploy/bootstrap ]; then echo "./deploy/bootstrap directory not found. Please cd to an appropriate directory"; exit 1; fi

# create namespace
kubectl apply -f deploy

# deploy imdb reference app
kubectl apply -f deploy/imdb

# deploy heartbeat, prometheus, grafana, and fluent bit
kubectl apply -f deploy/bootstrap
kubectl apply -R -f deploy/bootstrap

"$(dirname "${BASH_SOURCE[0]}")/jumpbox"

# deploy WebV after the app starts
kubectl wait pod -l app=imdb -n imdb --for condition=ready --timeout=60s

kubectl apply -f deploy/webv

`
}

func kicClusterRebuildScript() string {

	return `
#!/bin/bash

kic cluster create
kic cluster deploy

`
}
