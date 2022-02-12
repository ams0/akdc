#!/bin/bash

# this runs at Codespace creation - not part of pre-build

echo "$(date)    post-create start" >> ~/status

# pull repos
git -C /workspaces/ngsa pull
git -C /workspaces/webvalidate pull
git -C /workspaces/ngsa-app pull

# save ssl certs
mkdir -p $HOME/.ssh
echo "$INGRESS_KEY" | base64 -d > $HOME/.ssh/zone.key
echo "$INGRESS_CERT" | base64 -d > $HOME/.ssh/zone.crt

# add shared ssh key
echo "$ID_RSA" | base64 -d > $HOME/.ssh/id_rsa
echo "$ID_RSA_PUB" | base64 -d > $HOME/.ssh/id_rsa.pub

# set file mode
chmod 600 $HOME/.ssh/id*
chmod 600 $HOME/.ssh/zone.*

echo "$(date)    post-create complete" >> ~/status
