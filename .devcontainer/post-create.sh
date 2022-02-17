#!/bin/bash

# this runs at Codespace creation - not part of pre-build

echo "$(date)    post-create start" >> "$HOME/status"

# save ssl certs
mkdir -p "$HOME/.ssh"
echo "$INGRESS_KEY" | base64 -d > "$HOME/.ssh/certs.key"
echo "$INGRESS_CERT" | base64 -d > "$HOME/.ssh/certs.pem"

# add shared ssh key
echo "$ID_RSA" | base64 -d > "$HOME"/.ssh/id_rsa
echo "$ID_RSA_PUB" | base64 -d > "$HOME"/.ssh/id_rsa.pub

# set file mode
chmod 600 "$HOME"/.ssh/id*
chmod 600 "$HOME"/.ssh/certs.*

# update oh-my-zsh
git -C "$HOME/.oh-my-zsh" pull

# update repos
git -C ../workspaces/webvalidate pull
git -C ../workspaces/ngsa-app pull
git -C ../workspaces/edge-apps pull
git -C ../workspaces/edge-gitops pull
git -C ../workspaces/red-apps pull
git -C ../workspaces/red-gitops pull

# azure cli bug hot fix
# todo - remove after cli is fixed
sudo apt-get remove -y azure-cli
sudo apt-get install -y azure-cli=2.32.0-1~bullseye

echo "$(date)    post-create complete" >> "$HOME/status"
