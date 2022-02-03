#!/bin/bash

cd ..

./create-cluster.sh east ga atlanta 101 &
./create-cluster.sh east ga atlanta 102 &
./create-cluster.sh east ga atlanta 103 &
./create-cluster.sh east ga athens 101 &
./create-cluster.sh east ga athens 102 &
./create-cluster.sh east ga athens 103 &
