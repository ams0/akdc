#!/bin/bash

cd ..

./create-cluster.sh central mo stlouis 101 centralus &
./create-cluster.sh central mo stlouis 102 centralus &
./create-cluster.sh central mo stlouis 103 centralus &

./create-cluster.sh central mo kc 101 centralus &
./create-cluster.sh central mo kc 102 centralus &
./create-cluster.sh central mo kc 103 centralus &

./create-cluster.sh central tx austin 101 centralus &
./create-cluster.sh central tx austin 102 centralus &
./create-cluster.sh central tx austin 103 centralus &

./create-cluster.sh central tx dallas 101 centralus &
./create-cluster.sh central tx dallas 102 centralus &
./create-cluster.sh central tx dallas 103 centralus &

./create-cluster.sh central tx houston 101 centralus &
./create-cluster.sh central tx houston 102 centralus &
./create-cluster.sh central tx houston 103 centralus &
