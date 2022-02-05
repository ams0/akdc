# AKDC

- Create a k3d cluster in an Azure VM

## Setup

> From GitHub Codespaces

- CD to the vms directory

```bash

cd vms

```

- Add to your .zshrc

```bash

# Flux needs a PAT for the repo
export AKDC_PAT=YourPAT

```

- Check env variables

```bash

# check the value
echo $AKDC_PAT

```

- Login to Azure

```bash

az login --use-device-code

```

> The vms/scripts directory has scripts for creating and deleting groups of clusters

- Create the k3d cluster

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
ss austin-105

# check the VM setup status
# wait for "complete"
cat status

# force Flux to sync
sync

# check for flux pods
kubectl get pods -n flux-system

# reinstall flux on the VM if required
./flux-reset.sh

# exit the VM ssh shell
exit

```
