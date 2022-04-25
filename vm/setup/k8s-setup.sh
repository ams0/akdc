#!/bin/bash

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

echo "$(date +'%Y-%m-%d %H:%M:%S')  k8s-setup start" >> "/home/${AKDC_ME}/status"

sudo systemctl start snapd.socket
sudo snap install microk8s --classic
sudo systemctl stop snapd.socket

# sleep to avoid timing issues
sleep 5
kubectl wait node --all  --for condition=ready --timeout 30s
sleep 5
kubectl wait pod -l k8s-app=kube-dns -n kube-system --for condition=ready --timeout 30s

# Install istio resources on cluster
echo "$(date +'%Y-%m-%d %H:%M:%S')  installing istio resources" >> "/home/${AKDC_ME}/status"
microk8s.enable istio

echo "$(date +'%Y-%m-%d %H:%M:%S')  k8s-setup complete" >> "/home/${AKDC_ME}/status"
