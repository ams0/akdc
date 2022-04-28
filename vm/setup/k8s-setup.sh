#!/bin/bash

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

echo "$(date +'%Y-%m-%d %H:%M:%S')  k8s-setup start" >> "/home/${AKDC_ME}/status"

sudo systemctl start snapd.socket
sudo snap install microk8s --classic

# sleep to avoid timing issues
sleep 5

# setup kubectl
sudo microk8s kubectl config view --raw > "$HOME/.kube/config"
sudo chown -f -R akdc:akdc "$HOME"

kubectl wait node --all  --for condition=ready --timeout 30s
sleep 5
kubectl wait pod -l k8s-app=calico-node -n kube-system --for condition=ready --timeout 30s

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing DNS" >> "/home/${AKDC_ME}/status"
sudo microk8s enable dns ingress

pip=$(ip -4 a show eth0 | grep inet | sed "s/inet//g" | sed "s/ //g" | cut -d / -f 1 | grep 10.0)

echo "$(date +'%Y-%m-%d %H:%M:%S')  installing load balancer" >> "/home/${AKDC_ME}/status"
echo "$(date +'%Y-%m-%d %H:%M:%S')  IP: $pip" >> "/home/${AKDC_ME}/status"

microk8s enable metallb:"$pip-$pip"

echo "$(date +'%Y-%m-%d %H:%M:%S')  k8s-setup complete" >> "/home/${AKDC_ME}/status"
