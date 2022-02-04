#!/bin/bash

# check to see if the kustomize-controller is running in each server

for line in $(cat ips | sort | cut -f2);
do
  ssh -p 2222 -o "StrictHostKeyChecking=no" -o ConnectTimeout=5 akdc@$line 'if [[ $(kubectl get pods -A) == *"kustomize-controller"* ]]; then echo "$(hostname) success"; else echo "$(hostname) failed" && ./flux-setup.sh; fi'
done
