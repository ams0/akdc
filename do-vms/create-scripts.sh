#!/bin/bash

if [ -z $AKDC_PAT ]; then
  echo "Please export AKDC_PAT env variables"
  exit 1
fi

if [ $# -lt 4 ]; then
  echo "Usage: $0 Region State City Store-Number [AzureRegion:centralus]"
  exit 1
fi

Region=${1}
State=${2}
City=${3}
Number=${4}
Location=${5:-centralus}
District=${Region}-${State}-${City}
Store=${District}-${Number}

# create the install script from the template

# replace the host, pat, district and region
rm -f cluster-$Store.sh
sed "s/{{pat}}/$AKDC_PAT/g" ./akdc-do.templ | \
    sed "s/{{cluster}}/$Store/g" | \
    sed "s~{{repo}}~retaildevcrews/red-gitops~g" \
    > cluster-$Store.sh
