#!/bin/bash

### must run as akdc user

# this runs before flux-setup.sh

set -e

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-flux start" >> /home/akdc/status

# change ownership
sudo chown -R "$USER:$USER" /home/akdc
chmod 600 /home/akdc/.ssh/akdc.pat

sudo chown -R "$USER:$USER" /home/akdc

# create the tls secret
# this has to be installed before flux
if [ -f /home/akdc/.ssh/certs.pem ]
then
    kubectl create secret generic ssl-cert -n istio-system --from-file="key=/home/akdc/.ssh/certs.key" --from-file="cert=/home/akdc/.ssh/certs.pem"
fi

# create admin service account
kubectl create serviceaccount admin-user
kubectl create clusterrolebinding admin-user-binding --clusterrole cluster-admin --serviceaccount default:admin-user

if [ -f /home/akdc/.ssh/fluent-bit.key ]
then
    kubectl create ns fluent-bit
    kubectl create secret generic fluent-bit-secrets -n fluent-bit --from-file /home/akdc/.ssh/fluent-bit.key
fi

if [ -f /home/akdc/.ssh/prometheus.key ]
then
    kubectl create ns prometheus
    kubectl create secret -n prometheus generic prom-secrets --from-file /home/akdc/.ssh/prometheus.key
fi

if [ -d ./bootstrap ]
then
    kubectl apply -f ./bootstrap
    kubectl apply -R -f ./bootstrap
fi

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-flux complete" >> /home/akdc/status
