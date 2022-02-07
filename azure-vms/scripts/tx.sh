#!/bin/bash

cd ..

./create-cluster.sh central tx austin 101 centralus &
./create-cluster.sh central tx austin 102 centralus &
./create-cluster.sh central tx austin 103 centralus &

./create-cluster.sh central tx dallas 101 centralus &
./create-cluster.sh central tx dallas 102 centralus &
./create-cluster.sh central tx dallas 103 centralus &

./create-cluster.sh central tx houston 101 centralus &
./create-cluster.sh central tx houston 102 centralus &
./create-cluster.sh central tx houston 103 centralus &
