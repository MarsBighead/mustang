#!/bin/bash
docker rm -vf prometheus
docker run \
    --name prometheus \
    -p 9090:9090 \
    -v $HOME/mustang/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml \
    prom/prometheus


export DATA_SOURCE_NAME='root:togerme@(locahost:3306)/'