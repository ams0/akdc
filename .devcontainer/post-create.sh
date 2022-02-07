#!/bin/bash

echo "post-create start" >> ~/status

# clone repos
git clone https://github.com/retaildevcrews/edge-ngsa /workspaces/ngsa

# copy grafana.db to /grafana
sudo cp inner-loop/grafanadata/grafana.db /grafana
sudo chown -R 472:0 /grafana

# create local registry
docker network create k3d
k3d registry create registry.localhost --port 5500
docker network connect k3d k3d-registry.localhost

# add scripts to path
echo 'export PATH=$PATH:/workspaces/akdc/src/scripts' >> $HOME/.zshrc

# add akdc completions
cp src/scripts/_akdc $HOME/.oh-my-zsh/completions
unfunction _akdc && autoload -Uz _akdc

echo "post-create complete" >> ~/status
