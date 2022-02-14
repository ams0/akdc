# Create Kubernetes Clusters in Azure VMs

- Create one or more k3d clusters in unique Azure VMs
  - Bootstrap the cluster
    - Traefik for ingress
    - Flux for GitOps
    - DNS (optional)
    - SSL (optional)
    - Dapr
    - Radius

## Setup

> From GitHub Codespaces

## Login to Azure

```bash

az login --use-device-code

```

## cd to this directory

```bash

cd azure-vms

```

## Create the k3d cluster

- The VM will use `$HOME/.ssh/id_rsa` for SSH on port 2222
  - An SSH key will be generated if one doesn't exist
  - To reuse an existing SSH key, copy `id_rsa` and `id_rsa.pub` to `$HOME/.ssh`

> Our Codespaces installation automatically installs a "shared" ssh key

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
create-cluster Region State City Number -l centralus

```

## Check Setup

- Run until `complete` is the status for each server
  - akdc is our CLI
    - Add src/cli to your path
    - Our Codespaces install does this automatically

      ```bash

      # check all servers in ips file
      akdc check setup

      ```

## Check Flux Setup

```bash

# check all servers in ips file
akdc check flux

```

## Setup DNS and SSL

- A DNS zone and SSL cert can be used with the cluster setup to provide an ingress for http and https access to the clusters
- In order to use a zone, you must specify --zone with a zone name
  - `create-cluster` assumes the zone is both a domain and also an Azure Zone name

## Add ssl certs

> These files are unencrypted versions of your ssl cert
>
> Protect them appropriately!

- Make sure your CA certs are bundled in your certs.pem file
  - If you don't bundle your CA certs, you can end up with "works on my machine" issues
- If you're using Codespaces, these files are in `~/.ssh/certs.pem` and `~/.ssh/certs.key`
  - You can skip this step

## Download certs from Azure Key Vault

- If you have your keys in Azure Key Vault, you can download them

  ```bash

  # set these values for your Key Vault
  export AKDC_VAULT_NAME=kv-tld
  export AKDC_VAULT_KEY=cse-retail-key
  export AKDC_VAULT_CERT=cse-retail-crt

  # make sure the directory exists
  mkdir -p ~/.ssh

  # delete existing files
  rm -f ~/.ssh/certs.key
  rm -f ~/.ssh/certs.pem

  # get the files from Key Vault
  az keyvault secret show --vault-name "$AKDC_VAULT_NAME" -n "$AKDC_VAULT_KEY" --query "value" -o tsv > ~/.ssh/certs.key
  az keyvault secret show --vault-name "$AKDC_VAULT_NAME" -n "$AKDC_VAULT_CERT" --query "value" -o tsv > ~/.ssh/certs.pem

  # reduce permissions
  chmod 600 ~/.ssh/certs.*

  ```

## Extract certs from pfx file

- Download the certs pfx file
  - Make sure you know the password if set
- Extract the key into `~/.ssh/certs.key`
  - This is a two-step process
- Extract the ssl cert and the CA certs into `~/.ssh/certs.pem`
  - this is a two-step process

    ```bash

    ### make sure to run each command individually as they have prompts

    # chain your certs - order matters
    openssl pkcs12 -in retail.pfx -clcerts -nokeys -out ~/.ssh/certs.pem

    openssl pkcs12 -in retail.pfx -cacerts -nokeys -chain >> ~/.ssh/certs.pem

    # save your key
    # the pass code doesn't matter as the file is used once and deleted
    openssl pkcs12 -in retail.pfx -nocerts -out encrypted.key

    openssl rsa -in encrypted.key -out ~/.ssh/certs.key
    rm -f encrypted.key

    # reduce permissions
    chmod 600 ~/.ssh/certs.*

    ```

## Upload certs to Azure Key Vault

- If you have an Azure Key Vault, you can add the certs
  - Note: because of the line feeds, you can't easily paste the values from the portal

  ```bash

  # set these values for your Key Vault
  export AKDC_VAULT_NAME=kv-tld
  export AKDC_VAULT_KEY=cse-retail-key
  export AKDC_VAULT_CERT=cse-retail-crt

  # upload the files from Key Vault
  az keyvault secret set --vault-name "$AKDC_VAULT_NAME" -n "$AKDC_VAULT_KEY" -f ~/.ssh/certs.key
  az keyvault secret set --vault-name "$AKDC_VAULT_NAME" -n "$AKDC_VAULT_CERT" -f ~/.ssh/certs.pem

  ```

## Add the certs to GitHub Codepaces secrets

- Create an org level or repo level secret for
  - INGRESS_CERT
  - INGRESS_KEY
- base64 encode the values
  - The files have linefeeds
  - Copy the output and paste in the GitHub secret in your browser

    ```bash

    # base64 encode the values
    cat ~/.ssh/certs.key | base64
    cat ~/.ssh/certs.pem | base64

    ```

## Extracting the cert from Codespaces secrets

- We do this in `.devcontainer/post-create.sh`
  - Note Codespaces secrets are not available in `on-create.sh`

    ```bash

    # save ssl certs
    mkdir -p $HOME/.ssh

    echo "$INGRESS_KEY" | base64 -d > $HOME/.ssh/certs.key
    echo "$INGRESS_CERT" | base64 -d > $HOME/.ssh/certs.pem

    ```

## Create the Cluster with DNS and SSL

  ```bash

  # change to your DNS zone
  export AKDC_ZONE=cseretail.com

  create-cluster west ca east 105 -l westus --zone "$AKDC_ZONE" --ssl -c ~/.ssh/certs.pem -k ~/.ssh/akdc/certs.key 

  ```

## Check SSL

- You have to deploy an app with ssl ingress first
  - We deploy [TinyBench](https://github.com/bartr/tinybench) as part of our cluster bootstrap

    ```bash

    akdc check ssl

    https://west-ca-east-105.${AKDC_ZONE}

    ```

## Delete the cluster

> Make sure to use `akdc delete west-ca-east-105` to delete the RG and DNS record

  ```bash

  akdc delete west-ca-east-105

  # Edit or delete your ips file to remove the IP address

  ```
