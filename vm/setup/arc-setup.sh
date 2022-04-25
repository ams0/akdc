#!/usr/bin/env bash

set -e

echo "$(date +'%Y-%m-%d %H:%M:%S')  arc-setup start" >> "/home/${AKDC_ME}/status"

if [ "$AKDC_ARC_ENABLED" = "true" ]; then
  # add azure arc dependencies
  echo "$(date +'%Y-%m-%d %H:%M:%S')   install azure arc dependencies" >> "/home/${AKDC_ME}/status"
  az extension add --name connectedk8s
  az provider register --namespace Microsoft.Kubernetes
  az provider register --namespace Microsoft.KubernetesConfiguration
  az provider register --namespace Microsoft.ExtendedLocation

  # connect k8s to azure arc
  echo "$(date +'%Y-%m-%d %H:%M:%S')   connect k8s cluster to azure via azure arc" >> "/home/${AKDC_ME}/status"
  az connectedk8s connect --name "$AKDC_CLUSTER" --resource-group "$AKDC_RESOURCE_GROUP"
fi

echo "$(date +'%Y-%m-%d %H:%M:%S')  arc-setup complete" >> "/home/${AKDC_ME}/status"
