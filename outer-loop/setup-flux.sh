#!/bin/bash

if [ -z $AKDC_PAT ]
then
  echo "Please export AKDC_PAT=ValidGitOpsPAT"
  exit 1
fi

# create variables from command line
Region=${1:-$REGION}
State=${2:-$STATE}
City=${3:-$CITY}
Number=${4:-$NUMBER}
District=$Region-$State-$City
Store=$District-$Number

if [ -z $Region ] || [ -z $State ] || [ -z $City ] || [ -z $Number ]
then
  echo "Usage: $0 Region State City Number"
  exit 1
fi

# bootstrap flux
flux bootstrap git \
--url=https://github.com/retaildevcrews/edge-gitops \
--password=$AKDC_PAT \
--token-auth=true \
--path=./deploy/$Store

# add the GitOps repo
flux create source git gitops \
--url=https://github.com/retaildevcrews/edge-gitops \
--branch=main \
--password $AKDC_PAT

# add the store kustomization
flux create kustomization store \
--source GitRepository/gitops \
--path=./deploy/$Store \
--prune true \
--interval 1m

# add the district kustomization
flux create kustomization district \
--source GitRepository/gitops \
--path=./deploy/$District \
--prune true \
--interval 1m

# add the region kustomization
flux create kustomization region \
--source GitRepository/gitops \
--path=./deploy/$Region \
--prune true \
--interval 1m

flux reconcile source git gitops

kubectl get pods -A
