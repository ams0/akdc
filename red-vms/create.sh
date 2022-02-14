#!/bin/bash

echo "uncomment the commands before running"

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

# todo - remove once bug is fixed
cp ../src/cli/create-cluster .
cp ../src/cli/version.txt .
cp -R ../src/cli/lib .
sed "s/edge-gitops/red-gitops/g" < ../src/cli/akdc.templ > akdc.templ

# add Digital Ocean clusters
#echo -e "east-nc-raleigh-104\t167.71.181.26" > ips
#echo -e "east-nc-raleigh-105\t165.227.74.132" >> ips

# todo - remove ./ once bug is fixed
#./create-cluster central tx austin 101 -l centralus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &
#./create-cluster central tx austin 102 -l centralus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &
#./create-cluster central mo kc 101 -l centralus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &
#./create-cluster central mo kc 102 -l centralus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &
#./create-cluster east ga atlanta 101 -l eastus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &
#./create-cluster east ga atlanta 102 -l eastus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &
#./create-cluster west ca sd 101 -l westus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &
#./create-cluster west ca sd 102 -l westus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &
#./create-cluster west wa seattle 101 -l westus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &
#./create-cluster west wa seattle 102 -l westus -z cseretail.com --ssl -k ~/.ssh/certs.key -c ~/.ssh/certs.pem &

# todo - remove once bug is fixed
rm -f create-cluster
rm -f version.txt
rm -rf lib
rm -f akdc.templ
