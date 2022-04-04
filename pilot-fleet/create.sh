#!/bin/bash

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

# create the pilot clusters
### todo - fix this
### --repo has to run in serial
flt create --gitops --branch pilot -g pilot-fleet --ssl cseretail.com -c central-il-chi-101
flt create --gitops --branch pilot -g pilot-fleet --ssl cseretail.com -c central-il-chi-102
flt create --gitops --branch pilot -g pilot-fleet --ssl cseretail.com -c east-ny-nyc-101
flt create --gitops --branch pilot -g pilot-fleet --ssl cseretail.com -c east-ny-nyc-102
flt create --gitops --branch pilot -g pilot-fleet --ssl cseretail.com -c west-or-pdx-101
flt create --gitops --branch pilot -g pilot-fleet --ssl cseretail.com -c west-or-pdx-102
