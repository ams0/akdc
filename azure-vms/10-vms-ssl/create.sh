#!/bin/bash

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

create-cluster central ks kc 104 -l centralus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &
create-cluster central ks kc 105 -l centralus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &

create-cluster central tx aus 104 -l centralus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &
create-cluster central tx aus 105 -l centralus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &

create-cluster east ga atl 104 -l eastus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &
create-cluster east ga atl 105 -l eastus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &

create-cluster west ca sand 104 -l westus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &
create-cluster west ca sand 105 -l westus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &

create-cluster west wa sea 104 -l westus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &
create-cluster west wa sea 105 -l westus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &
