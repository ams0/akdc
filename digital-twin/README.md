# K8s with Codespaces

> Digital Twin Kubernetes Developer Experience using GitHub Codespaces

- Scenario
  - There is a broken app (badapp) in central-tx-austin-103
  - Validate the app is broken in the store
  - Create a Digital Twin of the store in GitHub Codespaces
  - Debug, fix and test the issue in the Digital Twin
  - Fix and Deploy the issue with GitOps
  - Validate Digital Twin and store are working

> Not part of the demo - reset your change!

## Start in digital-twin directory

```bash

cd digital-twin

```

## Create the local k3d cluster

> kic is the Kubernetes in Codespaces CLI

```bash

# create the cluster
kic create

# deploy jumpbox
kic jumpbox

```

## Add GitOps pointing to the Store Cluster config

> This allows you to locally debug store config issues

```bash

# setup flux
./setup-flux central-tx-austin-103

# check pods for "badapp"
k get po -A

# disconnect flux so GitOps doesn't change your config
flux uninstall -s

```

## Check the store cluster for the problem

> akdc is the CLI for working with simulated store clusters

```bash

# check heartbeat
akdc check heartbeat

# get the pods
akdc exec "k get po -A"

```

## Fix the app

```bash

# change to the GitOps directory that Flux uses for this store
cd /workspaces/edge-gitops/apps/central-tx-austin-103/badapp/dev/badapp
git pull

# edit the file and change "-url" to "-u" under "args"
code badapp.yaml

# check the change
git diff

# apply the change to the Digital Twin
k apply -f badapp.yaml

# check the pods
k get po -A

# use the jumpbox to execute the http command against the endpoint
kje http http://badapp.badapp.svc.cluster.local:8080/badapp/17

# undo your edit
git checkout .
git status

```

## Make the change and deploy with GitOps

```bash

# edit the app
cd /workspaces/edge-apps
git pull

# make the same change to the file
code apps/badapp/autogitops/dev/badapp.yaml

# push the change
akdc targets push

# change to the digital twin directory
cd /workspaces/akdc/digital-twin

# setup flux
./setup-flux.sh central-tx-austin-103

# wait for flux to sync
### you may have to repeat these two steps a few times while CI-CD runs and Flux syncs
kic sync
k get po -A

# use the jumpbox to execute the http command against the endpoint
kje http http://badapp.badapp.svc.cluster.local:8080/badapp/17

```

## Validate store cluster

```bash

# validate pods
### you may have to repeat these two steps a few times while CI-CD runs and Flux syncs
akdc sync
akdc exec "k get po -A"

# test the endpoint
http https://central-tx-austin-103.cseretail.com/badapp/17

```

## Undo your change

> Don't forget to undo your change!
>
> Seriously, please undo your change :)

```bash

# edit the app
cd /workspaces/edge-apps
git pull

# change "-u" back to "-url"
code apps/badapp/autogitops/dev/badapp.yaml

# push the change
akdc targets push

# wait for flux to sync
cd /workspaces/akdc/digital-twin

# pod should be broken again
### you may have to repeat these commands
kic sync
k get po -A

# store should be broken again
### you may have to repeat these commands
akdc sync
akdc exec "k get po -A"

```
