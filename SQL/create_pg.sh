#!/bin/bash
POSTGRES_PORT=5433
CONTAINER=pg11
# hbu is short of Heibei University
POSTGRE_USER=hbu
POSRGRES_PASSWORD=marsbighead
IMG=postgres:11
PG_CID=$(docker create -p $POSTGRES_PORT:5432  --name $CONTAINER \
                       -e POSTGRES_USER=$POSTGRE_USER \
                       -e POSTGRES_PASSWORD=$POSRGRES_PASSWORD \
                       $IMG)
docker start $PG_CID