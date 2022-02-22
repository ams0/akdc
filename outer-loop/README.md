# K8s with Codespaces

> Outer-loop Kubernetes Developer Experience using GitHub Codespaces

## Start in outer-loop directory

```bash

cd outer-loop

```

## Create the k3d cluster

> kic is the Kubernetes in Codespaces CLI

```bash

kic create

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
