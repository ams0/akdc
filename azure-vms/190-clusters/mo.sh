#!/bin/bash

cd ..

./create-cluster.sh central mo columbia 101 centralus &
./create-cluster.sh central mo columbia 102 centralus &
./create-cluster.sh central mo columbia 103 centralus &
./create-cluster.sh central mo columbia 104 centralus &
./create-cluster.sh central mo columbia 105 centralus &

./create-cluster.sh central mo east 101 centralus &
./create-cluster.sh central mo east 102 centralus &
./create-cluster.sh central mo east 103 centralus &
./create-cluster.sh central mo east 104 centralus &
./create-cluster.sh central mo east 105 centralus &

./create-cluster.sh central mo kc 101 centralus &
./create-cluster.sh central mo kc 102 centralus &
./create-cluster.sh central mo kc 103 centralus &
./create-cluster.sh central mo kc 104 centralus &
./create-cluster.sh central mo kc 105 centralus &

./create-cluster.sh central mo north 101 centralus &
./create-cluster.sh central mo north 102 centralus &
./create-cluster.sh central mo north 103 centralus &
./create-cluster.sh central mo north 104 centralus &
./create-cluster.sh central mo north 105 centralus &

./create-cluster.sh central mo south 101 centralus &
./create-cluster.sh central mo south 102 centralus &
./create-cluster.sh central mo south 103 centralus &
./create-cluster.sh central mo south 104 centralus &
./create-cluster.sh central mo south 105 centralus &

./create-cluster.sh central mo stlouis 101 centralus &
./create-cluster.sh central mo stlouis 102 centralus &
./create-cluster.sh central mo stlouis 103 centralus &
./create-cluster.sh central mo stlouis 104 centralus &
./create-cluster.sh central mo stlouis 105 centralus &

./create-cluster.sh central mo west 101 centralus &
./create-cluster.sh central mo west 102 centralus &
./create-cluster.sh central mo west 103 centralus &
./create-cluster.sh central mo west 104 centralus &
./create-cluster.sh central mo west 105 centralus &
