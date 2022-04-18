#!/bin/bash

### runs as akdc user

echo "$(date +'%Y-%m-%d %H:%M:%S')  flux bootstrap start" >> "/home/${AKDC_ME}/status"

if [ ! "$(flux --version)" ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  flux not found" >> "/home/${AKDC_ME}/status"
  echo "$(date +'%Y-%m-%d %H:%M:%S')  flux bootstrap failed" >> "/home/${AKDC_ME}/status"
  exit 1
fi

if [ -z "$AKDC_BRANCH" ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  AKDC_BRANCH not set" >> "/home/${AKDC_ME}/status"
  echo "$(date +'%Y-%m-%d %H:%M:%S')  flux bootstrap failed" >> "/home/${AKDC_ME}/status"
  echo "AKDC_BRANCH not set"
  exit 1
fi

if [ -z "$AKDC_CLUSTER" ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  AKDC_CLUSTER not set" >> "/home/${AKDC_ME}/status"
  echo "$(date +'%Y-%m-%d %H:%M:%S')  flux bootstrap failed" >> "/home/${AKDC_ME}/status"
  echo "AKDC_CLUSTER not set"
  exit 1
fi

if [ ! -f /home/akdc/.ssh/akdc.pat ]
then
  echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc.pat not found" >> "/home/${AKDC_ME}/status"
  echo "$(date +'%Y-%m-%d %H:%M:%S')  flux bootstrap failed" >> "/home/${AKDC_ME}/status"
  echo "akdc.pat not found"
  exit 1
fi

status_code=1
retry_count=0

until [ $status_code == 0 ]; do

  echo "flux retries: $retry_count"
  echo "$(date +'%Y-%m-%d %H:%M:%S')  flux retries: $retry_count" >> "/home/${AKDC_ME}/status"

  if [ $retry_count -gt 0 ]
  then
    sleep $((RANDOM % 30+15))
  fi

  retry_count=$((retry_count + 1))

  flux bootstrap git \
  --url "https://github.com/$AKDC_REPO" \
  --branch "$AKDC_BRANCH" \
  --password "$(cat /home/akdc/.ssh/akdc.pat)" \
  --token-auth true \
  --path "./deploy/bootstrap/$AKDC_CLUSTER"

  status_code=$?
done

echo "adding flux sources"
echo "$(date +'%Y-%m-%d %H:%M:%S')  adding flux sources" >> "/home/${AKDC_ME}/status"

flux create source git gitops \
--url "https://github.com/$AKDC_REPO" \
--branch "$AKDC_BRANCH" \
--password "$(cat /home/akdc/.ssh/akdc.pat)" \

flux create kustomization bootstrap \
--source GitRepository/gitops \
--path "./deploy/bootstrap/$AKDC_CLUSTER" \
--prune true \
--interval 1m

flux create kustomization apps \
--source GitRepository/gitops \
--path "./deploy/apps/$AKDC_CLUSTER" \
--prune true \
--interval 1m

flux reconcile source git gitops

kubectl get pods -A

echo "$(date +'%Y-%m-%d %H:%M:%S')  flux bootstrap complete" >> "/home/${AKDC_ME}/status"
