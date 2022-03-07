# App Setup

> From the `deploy` directory in this repo

```bash

# create namespace
kubectl apply -f deploy

```

> From the `aioa` directory in this repo

```bash

# set temporary Log Analytics secrets
kubectl create secret generic log-secrets \
  --from-literal=WorkspaceId=dev \
  --from-literal=SharedKey=dev

# display the secrets (base 64 encoded)
kubectl get secret log-secrets -o jsonpath='{.data}'

# deploy AI Order Accuracy app
kubectl apply -f aioa.yaml

# check pods until running
kubectl get pods -n ai-order-accuracy

# check local logs
kubectl logs ai-order-accuracy -n ai-order-accuracy

# check the version and genres endpoints
http $PIP:30080/version
http $PIP:30080/api/genres

# check logs
kubectl logs ai-order-accuracy -n ai-order-accuracy

# delete AI Order Accuracy
kubectl delete -f aioa.yaml

# check pods
kubectl get pods -n ai-order-accuracy

# Result - No resources found in default namespace.

```
