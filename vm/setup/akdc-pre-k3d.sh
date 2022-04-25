#!/bin/bash

# this runs before k3d-setup.sh

# change to this directory
#cd "$(dirname "${BASH_SOURCE[0]}")" || exit

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-k3d start" >> "/home/${AKDC_ME}/status"

echo "$(date +'%Y-%m-%d %H:%M:%S')  creating registry" >> "/home/${AKDC_ME}/status"
# create local registry
docker network create k3d

# create container registry
k3d registry create registry.localhost --port 5500
docker network connect k3d k3d-registry.localhost

# add the iot extension
az extension add -n azure-iot

# login to Azure
flt az-login

# create shared directories / mounts
sudo mkdir -p /k3d/var/lib/kubelet
sudo mkdir -p /k3d/etc/kubernetes
sudo chown -R "$USER":"$USER" /k3d

if [ "$(grep k3d /etc/fstab)" = "" ]
then
      echo "/k3d/var/lib/kubelet /k3d/var/lib/kubelet none bind,shared" | sudo tee -a /etc/fstab
fi

sudo mount -a

if [ "$(grep AKDC_STORAGE_KEY /etc/bash.bashrc)" = "" ]
then
      {
            echo ""
            echo "export AKDC_SP_ID=$AKDC_SP_ID"
            echo "export AKDC_SP_KEY=$AKDC_SP_KEY"
            echo "export AKDC_TENANT=$AKDC_TENANT"
            echo "export REPO_BASE=$REPO_BASE"
            echo "export AKDC_RESOURCE_GROUP=$AKDC_RESOURCE_GROUP"
            echo "export AKDC_STORAGE_NAME=$AKDC_STORAGE_NAME"
            echo "export AKDC_SUBSCRIPTION=$AKDC_SUBSCRIPTION"
            echo "export AKDC_VOLUME=$AKDC_VOLUME"
            echo "export AKDC_STORAGE_KEY=$AKDC_STORAGE_KEY"
            echo "export AKDC_STORAGE_CONNECTION=$AKDC_STORAGE_CONNECTION"
            echo "source <(kubectl completion bash)"
            echo "source <(flux completion bash)"
      } | sudo tee -a "/etc/bash.bashrc"
fi

# save the iot hub info
echo "IOTHUB_CONNECTION_STRING=$(az iot hub connection-string show --hub-name "$AKDC_RESOURCE_GROUP" -o tsv)" > ~/.ssh/iot.env
echo "IOTEDGE_DEVICE_CONNECTION_STRING=$(az iot hub device-identity connection-string show --hub-name "$AKDC_RESOURCE_GROUP" --device-id "$AKDC_CLUSTER" -o tsv)" >> ~/.ssh/iot.env

# create the azure credentials file
cat << EOF > /k3d/etc/kubernetes/azure.json
{
    "cloud":"AzurePublicCloud",
    "tenantId": "$AKDC_TENANT",
    "aadClientId": "$AKDC_SP_ID",
    "aadClientSecret": "$AKDC_SP_KEY",
    "subscriptionId": "$AKDC_SUBSCRIPTION",
    "resourceGroup": "$AKDC_RESOURCE_GROUP",
    "location": "centralus",
    "cloudProviderBackoff": false,
    "useManagedIdentityExtension": false,
    "useInstanceMetadata": true
}

EOF

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-k3d complete" >> "/home/${AKDC_ME}/status"
