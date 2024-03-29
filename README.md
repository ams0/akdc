# Retail Edge CLI

![License](https://img.shields.io/badge/license-MIT-green.svg)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](code_of_conduct.md)

> We use multiple GitHub Repos, so you have to use a PAT

- Create a Personal Access Token (PAT) in your GitHub account
  - Grant repo and package access
  - You can use an existing PAT
  - <https://github.com/settings/tokens>

- Create a personal Codespace secret
  - <https://github.com/settings/codespaces>
  - Name: PAT
  - Value: the PAT you just created
  - Grant access to this repo and any other repos you want

## Create a Codespace

- Click on `Code` then click `New Codespace`

Once Codespaces is running:

> Make sure your terminal is running zsh - bash is not supported and will not work

## Build the CLI

```bash

# change to the src directory
cd src/kic

# run make
make build

# make test
# make cover

```

## Boa Files

- We use `boa` to provide extensibility and customization
- The boa files are located at
  - Fleet CLI (flt)
    - bin/.flt
  - Inner-loop CLI (kic)
    - bin/.kic
  - VM CLI (kic)
    - vm/bin/.kic

## Code Coverage

- [code coverage](https://htmlpreview.github.io/?https://github.com/retaildevcrews/akdc/blob/main/src/kic/cover.html)

## Support

This project uses GitHub Issues to track bugs and feature requests. Please search the existing issues before filing new issues to avoid duplicates.  For new issues, file your bug or feature request as a new issue.

## Contributing

This project welcomes contributions and suggestions and has adopted the [Contributor Covenant Code of Conduct](https://www.contributor-covenant.org/version/2/1/code_of_conduct.html).

For more information see [Contributing.md](./.github/CONTRIBUTING.md)

## Trademarks

This project may contain trademarks or logos for projects, products, or services. Any use of third-party trademarks or logos are subject to those third-party's policies.
