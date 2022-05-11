// Copyright © Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
//   See LICENSE in the project root for license information.

// special version of kic for the kubecon hands-on labs

package kubekic

import (
	"kic/boa"

	"github.com/spf13/cobra"
)

// ListCmd lists the clusters in the fleet
var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build local apps",
}

func init() {
	boa.AddScriptCommand(BuildCmd, "imdb", "Build and deploy the IMDb reference app to the local cluster", kicBuildImdbScript())
	boa.AddScriptCommand(BuildCmd, "webv", "Build and deploy WebValidate to the local cluster", kicBuildWebvScript())
}

func kicBuildImdbScript() string {

	return `
#!/bin/bash

# validate directories
if [ ! -d /workspaces/imdb-app ]; then echo "/workspaces/imdb-app directory not found. Please clone the imdb-app repo to /workspaces"; exit 1; fi
if [ ! -d ./deploy ]; then echo "./deploy directory not found. Please cd to an appropriate directory"; exit 1; fi
if [ ! -d ./deploy/webv ]; then echo "./deploy/webv directory not found. Please cd to an appropriate directory"; exit 1; fi
if [ ! -d ./deploy/imdb ]; then echo "./deploy/imdb directory not found. Please cd to an appropriate directory"; exit 1; fi
if [ ! -d ./deploy/imdb-local ]; then echo "./deploy/imdb-local directory not found. Please cd to an appropriate directory"; exit 1; fi

# delete webv and imdb
kubectl delete -f deploy/webv --ignore-not-found=true
kubectl delete -f deploy/imdb --ignore-not-found=true

# build and push the local image
docker build /workspaces/imdb-app -t k3d-registry.localhost:5500/imdb-app:local
docker push k3d-registry.localhost:5500/imdb-app:local

# wait for delete to finish
kubectl wait pod -l app=webv -n imdb --for delete --timeout=30s
kubectl wait pod -l app=imdb -n imdb --for delete --timeout=30s

# deploy local app and re-deploy webv
kubectl apply -f deploy/imdb-local
kubectl wait pod -l app=imdb -n imdb --for condition=ready --timeout=30s
kubectl apply -f deploy/webv
kubectl wait pod -l app=webv -n imdb --for condition=ready --timeout=30s

# show status and curl results
kubectl get po -n imdb
http localhost:30080/version

`
}

func kicBuildWebvScript() string {

	return `
#!/bin/bash

# validate directories
if [ ! -d /workspaces/webvalidate ]; then echo "/workspaces/webvalidate directory not found. Please clone the webvalidate repo to /workspaces"; exit 1; fi
if [ ! -d ./deploy ]; then echo "./deploy directory not found. Please cd to an appropriate directory"; exit 1; fi
if [ ! -d ./deploy/webv ]; then echo "./deploy/webv directory not found. Please cd to an appropriate directory"; exit 1; fi
if [ ! -d ./deploy/webv-local ]; then echo "./deploy/webv-local directory not found. Please cd to an appropriate directory"; exit 1; fi

# delete local deployment
kubectl delete -f deploy/webv --ignore-not-found=true

# build and push the local docker image
docker build /workspaces/webvalidate -t k3d-registry.localhost:5500/webv:local
docker push k3d-registry.localhost:5500/webv:local

# create deployment from webv-local
kubectl wait pod -l app=webv -n imdb --for delete --timeout=30s
kubectl apply -f deploy/webv-local
kubectl wait pod -l app=webv -n imdb --for condition=ready --timeout=30s

# show pods and curl results
kubectl get pods -n imdb
"$(dirname "${BASH_SOURCE[0]}")/../check/webv"

`
}
