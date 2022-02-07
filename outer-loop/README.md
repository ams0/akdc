# K8s with Codespaces

> Outer-loop Kubernetes Developer Experience using GitHub Codespaces

## Start in outer-loop directory

```bash

cd outer-loop

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

## Add GitOps as a Store Cluster

> This allows you to locally debug store config issues

- setup Flux for GitOps

```bash

./setup-flux Region State City Number

```

- Valid params (case sensitive!)

```text

Example: ./setup-flux.sh central tx dallas 104

Region    State  City       Number
central   tx     dallas     104 or 105
central   tx     houston
central   mo     stlouis
east      ga     athens
east      nc     charlotte
west      ca     la
west      ca     sfo

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
   make reconcile        - force Flux reconcile
   make check            - check the endpoints with curl
   make test             - run a WebValidate test
   make load-test        - run a 60 second WebValidate test
   make jumpbox          - deploy a 'jumpbox' pod

```
