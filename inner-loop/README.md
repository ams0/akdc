# K8s with Codespaces

> Inner-loop Kubernetes Developer Experience using GitHub Codespaces

## Start in inner-loop directory

```bash

cd inner-loop

```

## Create and bootstrap k3d cluster

```bash

# make all runs
  # make delete
  # make create
  # make bootstrap
  # make jumpbox
make all

```

## Deploy local apps

```bash

make deploy

```

## Add GitOps as a Store Cluster

> This allows you to locally debug store config issues

- setup Flux for GitOps

```bash

./setup-flux Region State City Number

```

- Valid params (case sensitive!)

```text

Example: ./setup-flux central tx austin 104

Region    State  City       Number
central   tx     austin     104 or 105
central   tx     dallas
central   tx     houston
central   mo     kc
central   mo     stlouis
east      ga     athens
east      ga     atlanta
east      nc     charlotte
east      nc     raleigh
west      ca     la
west      ca     sd
west      ca     sfo
west      wa     seattle

```

## Other make commands

- We use make to implement and document various cluster management tasks

```bash

make

```

- Output

```text
Usage:
   make all              - create and bootstrap a cluster and deploy the apps
   make create           - create a k3d cluster
   make delete           - delete the k3d cluster
   make bootstrap        - deploy monitoring and logging
   make flux             - deploy Flux for GitOps
   make deploy           - deploy the apps to the cluster
   make check            - check the endpoints with curl
   make test             - run a WebValidate test
   make load-test        - run a 60 second WebValidate test
   make clean            - delete the apps from the cluster
   make app              - build and deploy a local app docker image
   make webv             - build and deploy a local WebV docker image
   make jumpbox          - deploy a 'jumpbox' pod

```
