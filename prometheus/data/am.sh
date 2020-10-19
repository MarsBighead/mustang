#!/bin/bash
docker rm -vf alertmanager
CURRENT_DIR=$(cd $(dirname $0);pwd)
echo $CURRENT_DIR
#exit 1
docker run -d \
  -p 9093:9093 \
  --name alertmanager \
  --restart=always \
  -v $CURRENT_DIR/alertmanager:/etc/alertmanager \
  -v $CURRENT_DIR/alertmanager/storage:/alertmanager \
  prom/alertmanager