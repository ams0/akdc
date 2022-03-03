# Kubernetes in Codespaces

> Outer-loop Kubernetes Developer Experience using GitHub Codespaces and k3d

## Start in outer-loop directory

- aioa-beta is a sample A/B deployment for `AI Order Accuracy`

```bash

cd outer-loop/aioa-beta

```

## Build and push docker image

- Normally, this would happen in ci-cd

```bash

make docker

```

## Verify GitOps targets

- GitOps is configured to deploy to redmond-101 and redmond-102

```bash

kic targets list

```

## Deploy with GitOps

- Normally, this would happen in ci-cd

```bash

make gitops

```

- You can remove the `version` file as it's only used to force GitOps to run
