# AKDC

![License](https://img.shields.io/badge/license-MIT-green.svg)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](code_of_conduct.md)

- Open this repo in Codespaces

> Best Practice - set AKDC_PAT as a Codespaces secret

- Export a valid GitHub PAT for Flux
  - Flux will use this PAT to connect to the GitHub repo
  - Flux needs write access

```bash

export AKDC_PAT=YourValidGitHubPAT

```

- Validate the PAT

```bash

  git clone https://${AKDC_PAT}@github.com/retaildevcrews/edge-gitops /workspaces/gitops
  ls -alF /workspaces/gitops/deploy
  
```

## Inner-loop Kubernetes Developer Experience

- [Readme](./inner-loop/README.md)

## Outer-loop Kubernetes Developer Experience

- [Readme](./outer-loop/README.md)

## Create a k3d cluster in an Azure VM(s)

- [Readme](./azure-vms/README.md)

## Support

This project uses GitHub Issues to track bugs and feature requests. Please search the existing issues before filing new issues to avoid duplicates.  For new issues, file your bug or feature request as a new issue.

## Contributing

This project welcomes contributions and suggestions and has adopted the [Contributor Covenant Code of Conduct](https://www.contributor-covenant.org/version/2/1/code_of_conduct.html).

For more information see [Contributing.md](./.github/CONTRIBUTING.md)

## Trademarks

This project may contain trademarks or logos for projects, products, or services. Any use of third-party trademarks or logos are subject to those third-party's policies.
