#!/bin/bash

cd ..

./create-cluster.sh central mo stlouis 101 centralus &
./create-cluster.sh central mo stlouis 102 centralus &
./create-cluster.sh central mo stlouis 103 centralus &

./create-cluster.sh central mo kc 101 centralus &
./create-cluster.sh central mo kc 102 centralus &
./create-cluster.sh central mo kc 103 centralus &
