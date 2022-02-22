# Kubernetes in Codespaces

> Inner-loop Kubernetes Developer Experience using GitHub Codespaces and k3d

## Start in inner-loop directory

```bash

cd inner-loop

```

## Create and bootstrap k3d cluster

> kic is the Kubernetes in Codespaces CLI

```bash

kic all

```

## Other CLI commands

- `kic` implements various cluster management tasks

```bash

kic

```

- Output

```text
Kubernetes in Codespaces CLI

Usage:
  kic [command]

Available Commands:
  all         create and bootstrap a local k3d cluster and deploy the apps
  app         build and deploy a local app docker image
  check       check status on the local k3d cluster
  clean       remove the apps from the local k3d cluster
  create      create a new local k3d cluster
  delete      delete the local k3d cluster
  deploy      deploy the apps to the local k3d cluster
  jumpbox     deploy a 'jumpbox' to the local k3d cluster
  test        run cluster tests
  webv        build and deploy a local WebV docker image

Flags:
  -h, --help   help for kic

Use "kic [command] --help" for more information about a command.

```
