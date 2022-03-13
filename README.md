# AKDC

![License](https://img.shields.io/badge/license-MIT-green.svg)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](code_of_conduct.md)

- Open this repo in Codespaces

> Best Practice - set AKDC_PAT as a Codespaces secret

- Export a valid GitHub PAT for Flux
  - Flux will use this PAT to connect to the GitHub repo
  - Flux needs write access

```bash

echo "YourValidGitHubPAT" > ~/.ssh/akdc.pat
chmod 600 ~/.ssh/akdc.pat

```

- Validate the PAT

```bash

rm -rf /workspaces/private-test
git clone https://$(cat ~/.ssh/akdc.pat)@github.com/retaildevcrews/private-test /workspaces/private-test

```

## Retail Edge Demo Dashboard

- [dashboard](https://retailedge.grafana.net/d/pQOetffnz)

## CLI Code Coverage

- [code coverage](https://htmlpreview.github.io/?https://github.com/retaildevcrews/akdc/blob/main/src/kic/cover.html)

## Inner-loop Developer Experience

- [Readme](./inner-loop/README.md)

## Outer-loop Developer Experience

- [Readme](./outer-loop/README.md)

## Digital Twin Developer Experience

- [Readme](./digital-twin/README.md)

## Arc Developer Experience

- work in progress

## Azure Stack HCI Developer Experience

- work in progress

## Fleet Developer Experience

- [Readme](./azure-vms/README.md)

## Support

This project uses GitHub Issues to track bugs and feature requests. Please search the existing issues before filing new issues to avoid duplicates.  For new issues, file your bug or feature request as a new issue.

## Contributing

This project welcomes contributions and suggestions and has adopted the [Contributor Covenant Code of Conduct](https://www.contributor-covenant.org/version/2/1/code_of_conduct.html).

For more information see [Contributing.md](./.github/CONTRIBUTING.md)

## Trademarks

This project may contain trademarks or logos for projects, products, or services. Any use of third-party trademarks or logos are subject to those third-party's policies.
