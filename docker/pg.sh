#!/bin/bash

CONTAINER=toger
PG_CID=$(docker create -p 5432:5432 --rm  --name $CONTAINER \
                       -e POSTGRES_USER=postgres \
                       -e POSTGRES_PASSWORD=postgres \
                       postgres:latest)
docker start $PG_CID