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

Example: ./create-cluster.sh central tx austin 104

```

```bash

# run create-cluster.sh
./create-cluster.sh Region State City Number [Azure Region: centralus]

```

## Check Status

- Wait for VM and k3d to install
  - This usually takes 10-15 minutes total

```bash

# check all servers in ips file
./flux-check.sh

```

- If a server fails
  - Flux is failing when creating multiple clusters due to a timing issue on the git repo

```bash
# ssh into the VM
# use the partial store name City-Number
# ss is a function injected into #HOME/.zshrc during Codespaces creation

ss austin-105

# check the VM setup status
# wait for "complete"
cat status

# force Flux to sync
sync

# check for flux pods
kubectl get pods -n flux-system

# reinstall flux on the VM if required
./flux-setup.sh

# exit the VM ssh shell
exit

```
