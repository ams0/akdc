#!/bin/bash

# check to see if the kustomize-controller is running in each server

for line in $(cat ips | sort | cut -f2);
do
  ssh -p 2222 -o "StrictHostKeyChecking=no" akdc@$line 'if [[ $(kubectl get pods -A) == *"ngsa-memory"* ]]; then echo "$(hostname) found"; else echo "$(hostname) not found"; fi'
done
