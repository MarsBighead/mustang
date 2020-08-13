#!/bin/bash

docker run -d --name webserver -p 80:80 nginx
docker run -d --name logspout -v /var/run.docker.socket:/tmp/docker.socket \
    gliderlabs/logspout syslog://127.0.0.1:5000