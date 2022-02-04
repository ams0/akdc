#!/bin/bash

# check to see if the kustomize-controller is running in each server

for line in $(cat ips | cut -f2);
do
  ssh -o "StrictHostKeyChecking=no" akdc@$line "hostname && kubectl get pods -n flux-system | grep kustomize-controller"
done

