#!/usr/bin/env bash

if [ "$AKDC_DEBUG" != "true" ]
then
  if [ "$AKDC_ARC_ENABLED" = "true" ]; then
    # install az cli
    echo "$(date +'%Y-%m-%d %H:%M:%S')   installing az cli" >> status
    sudo apt-get install -y azure-cli

    #shellcheck disable=2129
    echo "$(date +'%Y-%m-%d %H:%M:%S')   az login" >> status
    az login --identity >> status 2>&1

    # add azure arc dependencies
    echo "$(date +'%Y-%m-%d %H:%M:%S')   install azure arc dependencies" >> status
    az extension add --name connectedk8s
    az provider register --namespace Microsoft.Kubernetes
    az provider register --namespace Microsoft.KubernetesConfiguration
    az provider register --namespace Microsoft.ExtendedLocation

    # connect k3d to azure arc
    echo "$(date +'%Y-%m-%d %H:%M:%S')   connect k3d cluster to azure via azure arc" >> status
    az connectedk8s connect --name "$AKDC_CLUSTER" --resource-group "$AKDC_RESOURCE_GROUP"
  fi
fi
