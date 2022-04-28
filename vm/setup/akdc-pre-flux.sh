#!/bin/bash

# this runs before flux-setup.sh

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-flux start" >> "/home/${AKDC_ME}/status"

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

# create the tls secret
kubectl create secret tls ssl-cert -n ingress --key="$HOME/.ssh/certs.key" --cert="$HOME/.ssh/certs.pem"

# create admin service account
kubectl create serviceaccount admin-user
kubectl create clusterrolebinding admin-user-binding --clusterrole cluster-admin --serviceaccount default:admin-user

if [ -f /home/akdc/.ssh/fluent-bit.key ]
then
    kubectl create ns fluent-bit
    kubectl create secret generic fluent-bit-secrets -n fluent-bit --from-file "/home/${AKDC_ME}/.ssh/fluent-bit.key"
fi

if [ -f /home/akdc/.ssh/prometheus.key ]
then
    kubectl create ns prometheus
    kubectl create secret -n prometheus generic prom-secrets --from-file "/home/${AKDC_ME}/.ssh/prometheus.key"
fi

if [ -d ./bootstrap ]
then
    kubectl apply -f ./bootstrap
    kubectl apply -R -f ./bootstrap
fi

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre-flux complete" >> "/home/${AKDC_ME}/status"
