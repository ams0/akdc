#!/bin/bash

# this runs as part of pre-build

echo "on-create start"
echo "$(date +'%Y-%m-%d %H:%M:%S')    on-create start" >> "$HOME/status"

export REPO_BASE=$PWD
export AKDC_REPO=retaildevcrews/edge-gitops

export PATH="$PATH:$REPO_BASE"
export GOPATH="$HOME/go"

mkdir -p "$HOME/.ssh"
mkdir -p "$HOME/.oh-my-zsh/completions"

{
    echo "defaultIPs: /workspaces/edge-gitops/ips"
    echo "reservedClusterPrefixes: corp-monitoring central-mo-kc central-tx-austin east-ga-atlanta east-nc-raleigh west-ca-sd west-wa-redmond west-wa-seattle"
} > "$HOME/.kic"

{
    #shellcheck disable=2016,2028
    echo 'hsort() { read -r; printf "%s\\n" "$REPLY"; sort }'

    # add cli to path
    echo "export PATH=\$PATH:$REPO_BASE/bin"
    echo "export GOPATH=\$HOME/go"

    # create aliases
    echo "alias mk='cd $REPO_BASE/src/kic && make build; cd \$OLDPWD'"

    echo "export REPO_BASE=$PWD"
    echo "export AKDC_REPO=retaildevcrews/edge-gitops"
    echo "export AKDC_SSL=cseretail.com"
    echo "export AKDC_GITOPS=true"
    echo "compinit"

    echo "if [ \"\$PAT\" != \"\" ]"
    echo "then"
    echo "    export GITHUB_TOKEN=\$PAT"
    echo "    export AKDC_PAT=\$PAT"
    echo "fi"

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

# clone repos
cd ..
git clone https://github.com/microsoft/webvalidate
git clone https://github.com/cse-labs/imdb-app
git clone https://github.com/cse-labs/kubernetes-in-codespaces inner-loop
git clone https://github.com/retaildevcrews/edge-gitops
git clone https://github.com/retaildevcrews/red-gitops
git clone https://github.com/retaildevcrews/vtlog
cd "$REPO_BASE" || exit

echo "generating kic completion"
kic completion zsh > "$HOME/.oh-my-zsh/completions/_kic"
flt completion zsh > "$HOME/.oh-my-zsh/completions/_flt"

echo "on-create complete"
echo "$(date +'%Y-%m-%d %H:%M:%S')    on-create complete" >> "$HOME/status"
