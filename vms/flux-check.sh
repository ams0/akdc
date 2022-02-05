#!/bin/bash

# check to see if the kustomize-controller is running in each server

for line in $(cat ips | sort | cut -f2);
do
  ssh -p 2222 -o "StrictHostKeyChecking=no" -o ConnectTimeout=5 akdc@$line 'if [[ $(kubectl get ns) == *"flux-check"* ]]; then echo "$(hostname) success"; else echo "$(hostname) failed"; fi'
done
