# App Setup

> From the `ngsa` directory in this repo

```bash

# set temporary Log Analytics secrets
kubectl create secret generic log-secrets \
  --from-literal=WorkspaceId=dev \
  --from-literal=SharedKey=dev

# display the secrets (base 64 encoded)
kubectl get secret log-secrets -o jsonpath='{.data}'

# deploy ngsa app
kubectl apply -f ngsa.yaml

# check pods until running
kubectl get pods

# check local logs
kubectl logs ngsa

# check the version and genres endpoints
http $PIP:30080/version
http $PIP:30080/api/genres

# check logs
kubectl logs ngsa

# delete ngsa
kubectl delete -f ngsa.yaml

# check pods
kubectl get pods

# Result - No resources found in default namespace.

```
