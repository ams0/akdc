# Virtual IOT Edge on Linux

> Deploy an Azure VM with IOT Edge running on Ubuntu

## Add extension and set environment variables

```bash

az extension add --name azure-iot

export AKDC_IOT_HUB_NAME=akdc-iot-hub
export AKDC_IOT_HUB_GROUP=iot-hub

```

## Create the IOT Hub

> This is already done in our subscription

```bash

az group create --location centralus -g $AKDC_IOT_HUB_GROUP
az iot hub create --name $AKDC_IOT_HUB_NAME -g $AKDC_IOT_HUB_GROUP --sku S1

```

## Create VM

```bash

flt create --ssl cseretail.com --debug -c $AKDC_IOT_HOST

```

## Register device with IOT Hub

```bash

### change this to your VM name
export AKDC_IOT_HOST=iot-edge-101

az iot hub device-identity create --device-id $AKDC_IOT_HOST --hub-name $AKDC_IOT_HUB_NAME --edge-enabled
az iot hub device-identity list --hub-name $AKDC_IOT_HUB_NAME
az iot hub device-identity connection-string show --device-id $AKDC_IOT_HOST --hub-name $AKDC_IOT_HUB_NAME --query connectionString -o tsv

```

## SSH into the VM

```bash

# ssh into the vm
flt ssh $AKDC_IOT_HOST

### should be on the VM - check your prompt

# delete k3d
k3d delete cluster

az extension add --name azure-iot

# add the package
wget https://packages.microsoft.com/config/ubuntu/20.04/packages-microsoft-prod.deb -O packages-microsoft-prod.deb
sudo dpkg -i packages-microsoft-prod.deb
rm packages-microsoft-prod.deb

# update
sudo apt update
sudo apt upgrade -y

# install iot edge
sudo apt-get install -y iotedge

# get your connection string
export CONN_STRING=$(az iot hub device-identity connection-string show --device-id $(hostname) --hub-name $AKDC_IOT_HUB_NAME --query connectionString -o tsv)

# verify connection string
echo $CONN_STRING

# update config
sudo sed -i s"/<ADD DEVICE CONNECTION STRING HERE>/$CONN_STRING/g" /etc/iotedge/config.yaml

# check config
sudo cat /etc/iotedge/config.yaml | grep akdc

# restart iotedge
sudo systemctl restart iotedge

# check status
sudo systemctl status iotedge

# useful for debugging any issues
journalctl -u iotedge

sudo iotedge check

sudo iotedge list

```

## Run Heartbeat

```bash

# need to install caddy for ingress
# for now, run on http
docker run -d --restart always --name heartbeat -p 80:8080 ghcr.io/bartr/tinybench -u /heartbeat

```
