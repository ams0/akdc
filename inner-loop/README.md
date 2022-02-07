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
  # make deploy
  # make jumpbox
make all

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
   make deploy           - deploy the apps to the cluster
   make check            - check the endpoints with curl
   make test             - run a WebValidate test
   make load-test        - run a 60 second WebValidate test
   make clean            - delete the apps from the cluster
   make app              - build and deploy a local app docker image
   make webv             - build and deploy a local WebV docker image
   make jumpbox          - deploy a 'jumpbox' pod

```
