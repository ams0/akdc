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

```bash

# run akdc create
akdc create -c central-tx-dallas-105 -l centralus

```

## Check Setup

- Run until `complete` is the status for each server
  - akdc is our CLI
    - Add $REPO_ROOT/bin to your path
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
  - `akdc` assumes the zone is both a domain and also an Azure Zone name

## Add ssl certs

> These files are unencrypted versions of your ssl cert
>
> Protect them appropriately!

- If you're using Codespaces, these files are in `~/.ssh/certs.pem` and `~/.ssh/certs.key`
  - You can skip this step
- [SSL Cert Setup](./CERTS.md)

## Create the Cluster with DNS and SSL

  ```bash

  # change to your DNS zone
  export AKDC_ZONE=cseretail.com

  akdc create west-ca-east-105 -l westus --zone "$AKDC_ZONE" --ssl

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
