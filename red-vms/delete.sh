#!/bin/bash

echo "uncomment the commands before running"

### remove the lock from the retail-edge-demo RG

# this will delete the DNS entries
### RG delete will fail and can be ignored
# akdc delete central-tx-austin-101
# akdc delete central-tx-austin-102
# akdc delete central-mo-kc-101
# akdc delete central-mo-kc-102
# akdc delete east-ga-atlanta-101
# akdc delete east-ga-atlanta-102
# akdc delete west-ca-sd-101
# akdc delete west-ca-sd-102
# akdc delete west-wa-seattle-101
# akdc delete west-wa-seattle-102
# akdc delete corp-monitoring-101

# delete the RG
#az group delete -y --no-wait -g retail-edge-demo

# add Digital Ocean clusters
#echo -e "east-nc-raleigh-101\t138.197.79.242" > "$(dirname "${BASH_SOURCE[0]}")/ips"
#echo -e "east-nc-raleigh-102\t209.97.144.124" >> "$(dirname "${BASH_SOURCE[0]}")/ips"
