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

## DNS & SSL

A DNS zone and SSL can be used with the cluster setup to provide an ingress for http and http access to resources in k3d.

In order to use a zone, you must specify --zone with a zone name. It's assume the zone is both a domain and also an Azure Zone name.

for example to use the cseretail.com domain

```bash

# run create-cluster.sh
./create-cluster.sh west ca east 105 --zone cseretail.com 

```

### SSL

For SSL, first locate your ssl cert and convert the pfx to a .crt and .key file. The key file must be a decrypted key file as traefik will not work with an encrypted file. Ensure that both files are on the local disk.

### Options 1) Working with a PFX Cert Directly
Assuming your cert is in a key vault called kv-tld which is in a resource group called tld.

* Go to the keyvault named `kv-tld` in the `tld` resource group.
* Navigate to secrets and specifically the cseretail* secret
* Load the current version of the key and scroll down.
* Click the Download as a certificate button to download a local copy.
* This downloads the cert as a PFX file.
* Move the file out of the repo into ~/.ssh
  * `mkdir -p ~/.ssh/akdc`
  * `cp <pfx file> ~/.ssh/akdc/<pfx file>`
* In order to use this with traefik, we must first convert it to a .crt and .key file.
* Ensure you have openssl installed.
  * `cd ~/.ssh/akdc`
  * `openssl pkcs12 -in <pfx file> -nocerts -out cseretail-encrypted.key`
  * `openssl pkcs12 -in <pfx file>  -clcerts -nokeys -out cseretail.crt`
  * `openssl rsa -in cseretail-encrypted.key  -out cseretail-decrypted.key`

### Option 2) Working with CRT and Key Files

In some cases, the certifact conversation can be completed beforehand and stored in the keyvault as well. The `kv-tld` vault has crt and key secrets that can be used. 

To download the secrets via the az cli

* `mkdir -p ~/.ssh/akdc`
* ` az keyvault secret show --vault-name kv-tld -n cse-retail-key --query "value" -o tsv > ~/.ssh/akdc/cseretail-decrypted.key`
* ` az keyvault secret show --vault-name kv-tld -n cse-retail-crt --query "value" -o tsv > ~/.ssh/akdc/cseretail.crt`

### Create the Cluster with ssl
```bash

# run create-cluster.sh
./create-cluster.sh west ca east 105 --zone cseretail.com --ssl -c ~/.ssh/akdc/cseretail.crt -k ~/.ssh/akdc/cseretail-decrypted.key 
```

