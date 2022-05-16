#!/bin/bash

# this runs as part of pre-build

echo "on-create start"
echo "$(date +'%Y-%m-%d %H:%M:%S')    on-create start" >> "$HOME/status"

### change these as needed
export REPO_BASE=$PWD
export AKDC_SSL=cseretail.com
export AKDC_GITOPS=true
export AKDC_DNS_RG=tld

export PATH="$PATH:$PWD/bin"
export GOPATH="$HOME/go"

mkdir -p "$HOME/.ssh"
mkdir -p "$HOME/.oh-my-zsh/completions"

{
    echo "defaultIPs: \$AKDC_REPO/ips"
    echo "reservedClusterPrefixes: corp-monitoring central-mo-kc central-tx-austin east-ga-atlanta east-nc-raleigh west-ca-sd west-wa-redmond west-wa-seattle"
} > "$HOME/.kic"

{
    #shellcheck disable=2016,2028
    echo 'hsort() { read -r; printf "%s\\n" "$REPLY"; sort }'

    # add cli to path
    echo "export PATH=\$PATH:$PWD/bin"
    echo "export GOPATH=\$HOME/go"

    # unset secrets
    # todo - change prefix
    echo "unset AKDC_SSL_KEY"
    echo "unset AKDC_SSL_CERT"
    echo "unset AKDC_ID_RSA"
    echo "unset AKDC_ID_RSA_PUB"
    echo "unset AKDC_LOKI_URL"
    echo "unset AKDC_PROMETHEUS_KEY"
    echo "unset AKDC_EVENT_HUB"
    echo "unset AKDC_SP_ID"
    echo "unset AKDC_SP_KEY"
    echo "unset AKDC_TENANT"

    # create aliases
    echo "alias mk='cd $PWD/src/kic && make build; cd \$OLDPWD'"
    echo "export AKDC_SSL=$AKDC_SSL"
    echo "export AKDC_GITOPS=$AKDC_GITOPS"
    echo "export AKDC_DNS_RG=$AKDC_DNS_RG"
    echo "export AKDC_MI=$AKDC_MI"
    echo "export REPO_BASE=$REPO_BASE"
    echo ""

    echo "if [ \"\$PAT\" != \"\" ]"
    echo "then"
    echo ""
    echo "    export GITHUB_TOKEN=\$PAT"
    echo "    export AKDC_PAT=\$PAT"
    echo "fi"

    echo ""
    echo "compinit"
} >> "$HOME/.zshrc"

# create local registry
docker network create k3d
k3d registry create registry.localhost --port 5500
docker network connect k3d k3d-registry.localhost

# install cobra
go install -v github.com/spf13/cobra/cobra@latest

# install go modules
go install -v golang.org/x/lint/golint@latest
go install -v github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
go install -v github.com/ramya-rao-a/go-outline@latest
go install -v github.com/cweill/gotests/gotests@latest
go install -v github.com/fatih/gomodifytags@latest
go install -v github.com/josharian/impl@latest
go install -v github.com/haya14busa/goplay/cmd/goplay@latest
go install -v github.com/go-delve/delve/cmd/dlv@latest
go install -v honnef.co/go/tools/cmd/staticcheck@latest
go install -v golang.org/x/tools/gopls@latest

# build the cli
cd src/kic || exit
make build
cd "$OLDPWD" || exit

echo "generating completions"
flt completion zsh > "$HOME/.oh-my-zsh/completions/_flt"
kic completion zsh > "$HOME/.oh-my-zsh/completions/_kic"
kivm completion zsh > "$HOME/.oh-my-zsh/completions/_kivm"

# only run apt upgrade on pre-build
if [ "$CODESPACE_NAME" = "null" ]
then
    echo "$(date +'%Y-%m-%d %H:%M:%S')    upgrading" >> "$HOME/status"
    sudo apt-get update
    sudo apt-get upgrade -y
fi

echo "on-create complete"
echo "$(date +'%Y-%m-%d %H:%M:%S')    on-create complete" >> "$HOME/status"
