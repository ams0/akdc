#!/bin/bash

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-post start" >> "/home/${AKDC_ME}/status"

docker pull ghcr.io/cse-labs/webv-red:latest
docker pull ghcr.io/cse-labs/webv-red:beta

docker pull ghcr.io/cse-labs/heartbeat
docker pull ghcr.io/cse-labs/imdb-app

docker network create akdc
docker run -d --restart always --net akdc --name heartbeat -p 30082:8080 ghcr.io/cse-labs/heartbeat

if [[ "$(hostname)" == *"monitor"* ]]
then
    docker run -d --restart always --name webv-heartbeat -p 30088:8080 ghcr.io/cse-labs/webv-red:beta -s https://central-tx-atx-512.cseretail.com https://east-ga-atl-512.cseretail.com https://west-wa-sea-512.cseretail.com -f heartbeat-benchmark.json --run-loop --prometheus --sleep 5000
else
    docker run -d --restart always --net akdc --name imdb -p 30080:8080 ghcr.io/cse-labs/imdb-app
fi

echo "$(date +'%Y-%m-%d %H:%M:%S')  akdc-post complete" >> "/home/${AKDC_ME}/status"
