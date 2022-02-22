#!/bin/bash

echo "$(date +'%Y-%m-%d %H:%M:%S')  flux bootstrap start" >> status

# don't run if --debug
if [ "$AKDC_DEBUG" = "true" ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  AKDC_DEBUG is set to true" >> status
else
  if [ -z "$AKDC_CLUSTER" ]
  then
    echo "$(date +'%Y-%m-%d %H:%M:%S')  AKDC_CLUSTER not set" >> status
    echo "$(date +'%Y-%m-%d %H:%M:%S')  flux bootstrap failed" >> status
    echo "AKDC_CLUSTER not set"
    exit 1
  fi

  if [ ! -f .ssh/akdc.pat ]
  then
    echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc.pat not found" >> status
    echo "$(date +'%Y-%m-%d %H:%M:%S')  flux bootstrap failed" >> status
    echo "akdc.pat not found"
    exit 1
  fi

  status_code=1
  retry_count=0

  until [ $status_code == 0 ]; do

    echo "flux retries: $retry_count"
    echo "$(date +'%Y-%m-%d %H:%M:%S')  flux retries: $retry_count" >> status

    if [ $retry_count -gt 0 ]
    then
      sleep 20
    fi

    retry_count=$((retry_count + 1))

    flux bootstrap git \
    --url https://github.com/$AKDC_REPO \
    --password $(cat .ssh/akdc.pat) \
    --token-auth true \
    --path ./bootstrap/$AKDC_CLUSTER

    status_code=$?
  done

  echo "adding flux sources"
  echo "$(date +'%Y-%m-%d %H:%M:%S')  adding flux sources" >> status

  flux create source git gitops \
  --url https://github.com/$AKDC_REPO \
  --branch main \
  --password $(cat .ssh/akdc.pat) \

  flux create kustomization bootstrap \
  --source GitRepository/gitops \
  --path ./bootstrap/$AKDC_CLUSTER \
  --prune true \
  --interval 1m

  flux create kustomization apps \
  --source GitRepository/gitops \
  --path ./apps/$AKDC_CLUSTER \
  --prune true \
  --interval 1m

  flux reconcile source git gitops

  kubectl get pods -A
fi

echo "$(date +'%Y-%m-%d %H:%M:%S')  flux bootstrap complete" >> status
