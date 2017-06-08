#!/bin/bash

docker run --rm \
   -v "$(pwd)":/home/hbu/mustang/docker/hello \ 
   -w "$(pwd)" golang:1.8.1 \
   go build -v
