#!/bin/bash

# change to this directory
cd "$(dirname "${BASH_SOURCE[0]}")" || exit

create-cluster central ks kc 101 -l centralus &
create-cluster central ks kc 102 -l centralus &

create-cluster central tx aus 101 -l centralus &
create-cluster central tx aus 102 -l centralus &

create-cluster east ga atl 101 -l eastus &
create-cluster east ga atl 102 -l eastus &

create-cluster west ca sand 101 -l westus &
create-cluster west ca sand 102 -l westus &

create-cluster west wa sea 101 -l westus &
create-cluster west wa sea 102 -l westus &
