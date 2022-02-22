#!/bin/bash

# this runs before flux-setup.sh
# this does not if akdc create --debug is used

# runs as akdc user
# env variables defined in .bashrc
    # AKDC_CLUSTER
    # AKDC_REPO
    # AKDC_FQDN
    # AKDC_DEBUG

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre start" >> status

# change ownership
sudo chown -R $USER:$USER .
chmod 600 .ssh/akdc.pat

if [ "$AKDC_DEBUG" = "true" ]
then
    # clone the GitOps repo
    git clone https://"$(cat .ssh/akdc.pat)@github.com/$AKDC_REPO"
else
    # create the tls secret
    # this has to be installed before flux
    if [ -f .ssh/certs.pem ]
    then
        kubectl create secret tls ssl-cert --cert .ssh/certs.pem --key .ssh/certs.key
    fi

    # remove the certs
    rm -f .ssh/certs.pem
    rm -f .ssh/certs.key

    # create any bootstrap K8s resources
    if [ -d ./bootstrap ]
    then
        kubectl apply -R -f ./bootstrap
    fi
fi

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-pre complete" >> status
