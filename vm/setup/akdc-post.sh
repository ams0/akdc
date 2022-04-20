#!/bin/bash

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-post start" >> "/home/${AKDC_ME}/status"

docker pull ghcr.io/cse-labs/webv-red:latest
docker pull ghcr.io/cse-labs/webv-red:beta

docker pull ghcr.io/cse-labs/heartbeat
docker pull ghcr.io/cse-labs/imdb-app

docker network create akdc

if [[ "$(hostname)" != *"monitor"* ]]
then
    docker run -d --restart always --net akdc --name heartbeat -p 30082:8080 ghcr.io/cse-labs/heartbeat
    docker run -d --restart always --net akdc --name imdb      -p 30080:8080 ghcr.io/cse-labs/imdb-app
    docker run -d --restart always --net akdc --name webv      -p 30088:8080 ghcr.io/cse-labs/webv-red:beta -s http://imdb:8080 -f benchmark.json --run-loop --prometheus --sleep 100 --verbose
fi

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-post complete" >> "/home/${AKDC_ME}/status"
