#!/bin/bash

cd ..

./create-cluster.sh west ca east 101 westus &
./create-cluster.sh west ca east 102 westus &
./create-cluster.sh west ca east 103 westus &
./create-cluster.sh west ca east 104 westus &
./create-cluster.sh west ca east 105 westus &

./create-cluster.sh west ca la 101 westus &
./create-cluster.sh west ca la 102 westus &
./create-cluster.sh west ca la 103 westus &
./create-cluster.sh west ca la 104 westus &
./create-cluster.sh west ca la 105 westus &

./create-cluster.sh west ca north 101 westus &
./create-cluster.sh west ca north 102 westus &
./create-cluster.sh west ca north 103 westus &
./create-cluster.sh west ca north 104 westus &
./create-cluster.sh west ca north 105 westus &

./create-cluster.sh west ca sac 101 westus &
./create-cluster.sh west ca sac 102 westus &
./create-cluster.sh west ca sac 103 westus &
./create-cluster.sh west ca sac 104 westus &
./create-cluster.sh west ca sac 105 westus &

./create-cluster.sh west ca sfo 101 westus &
./create-cluster.sh west ca sfo 102 westus &
./create-cluster.sh west ca sfo 103 westus &
./create-cluster.sh west ca sfo 104 westus &
./create-cluster.sh west ca sfo 105 westus &

./create-cluster.sh west ca sd 101 westus &
./create-cluster.sh west ca sd 102 westus &
./create-cluster.sh west ca sd 103 westus &
./create-cluster.sh west ca sd 104 westus &
./create-cluster.sh west ca sd 105 westus &

./create-cluster.sh west ca south 101 westus &
./create-cluster.sh west ca south 102 westus &
./create-cluster.sh west ca south 103 westus &
./create-cluster.sh west ca south 104 westus &
./create-cluster.sh west ca south 105 westus &

./create-cluster.sh west ca south 101 westus &
./create-cluster.sh west ca south 102 westus &
./create-cluster.sh west ca south 103 westus &
./create-cluster.sh west ca south 104 westus &
./create-cluster.sh west ca south 105 westus &