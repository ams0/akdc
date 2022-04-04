#!/bin/bash

# this will delete the DNS entries
flt delete central-il-chi-101
flt delete central-il-chi-102
flt delete east-ny-nyc-101
flt delete east-ny-nyc-102
flt delete west-or-pdx-101
flt delete west-or-pdx-102

# delete the RG
az group delete -y --no-wait -g pilot-fleet
