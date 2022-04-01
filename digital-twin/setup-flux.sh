#!/bin/bash

if [ ! -f "$HOME/.ssh/akdc.pat" ]
then
  echo "Please export AKDC_PAT to $HOME/.ssh/akdc.pat"
  exit 1
fi

# create variables from command line
Cluster=${1}
Branch=${2}

if [ -z "$Cluster" ] || [ -z "$Branch" ]
then
  echo "Usage: $0 ClusterName branch"
  exit 1
fi

# bootstrap flux
flux bootstrap git \
--url https://github.com/retaildevcrews/edge-gitops \
--branch "$Branch" \
--password "$(cat "$HOME"/.ssh/akdc.pat)" \
--token-auth true \
--path "./deploy/dev/$Cluster"

# add the GitOps repo
flux create source git gitops \
--url https://github.com/retaildevcrews/edge-gitops \
--branch "$Branch" \
--password "$(cat "$HOME"/.ssh/akdc.pat)"

# add the apps kustomization
flux create kustomization apps \
--source GitRepository/gitops \
--path "./deploy/apps/$Cluster" \
--prune true \
--interval 1m

flux reconcile source git gitops

kubectl get pods -A
