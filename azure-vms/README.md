# AKDC

- Create a k3d cluster in one or more Azure VM(s)
  - Bootstrap the cluster with Flux for GitOps

## Setup

> From GitHub Codespaces

## Login to Azure

```bash

az login --use-device-code

```

## CD to the azure-vms directory

```bash

cd azure-vms

```

## Create the k3d cluster

> The vms/scripts directory has scripts for creating and deleting groups of clusters

- The VM will use `$HOME/.ssh/id_rsa` for SSH on port 2222
  - An SSH key will be generated if one doesn't exist
  - To reuse an existing SSH key, copy `id_rsa` and `id_rsa.pub` to `$HOME/.ssh`

- Valid params (case sensitive!)

```text

Region    State  City       Number
central   tx     dallas     104 or 105
central   tx     houston
central   mo     stlouis
east      ga     athens
east      nc     charlotte
west      ca     la
west      ca     sfo

Example: ./create-cluster.sh central tx dallas 104

```

```bash

# run create-cluster.sh
./create-cluster.sh Region State City Number [Azure Region: centralus]

```

## Check Setup

- Run until `complete` is the status for each server

```bash

# check all servers in ips file
akdc check setup

```

## Check Flux Setup

```bash

# check all servers in ips file
akdc check flux

```

- If a server fails
  - Flux is failing when creating multiple clusters due to a timing issue on the git repo

```bash

# setup flux if missing
akdc flux setup

# check for flux
akdc check flux

```
