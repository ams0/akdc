# Kubernetes in Codespaces

> Inner-loop Kubernetes Developer Experience using GitHub Codespaces and k3d

## Start in inner-loop directory

```bash

cd inner-loop

```

## Create and bootstrap k3d cluster

> kic is the Kubernetes in Codespaces CLI

```bash

kic manage cluster all

```

## Other CLI commands

- `kic manage cluster` implements various cluster management tasks

```bash

kic manage cluster

```

- Output

```text

k3d manage cluster commands

Usage:
  kic manage cluster [command]

Available Commands:
  all         Create and bootstrap a local k3d cluster and deploy the apps
  clean       Remove the apps from the local k3d cluster
  create      Create a new local k3d cluster
  delete      Delete the k3d cluster
  deploy      Deploy the apps to the local k3d cluster
  jumpbox     Deploy a 'jumpbox' to the local k3d cluster

Flags:
  -h, --help   help for cluster

Use "kic manage cluster [command] --help" for more information about a command.

```

## Hands-on Lab

- A hands-on lab is available at [cse-labs](https://github.com/cse-labs/kubernetes-in-codespaces)
