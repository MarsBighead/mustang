#!/bin/bash

pg_conf=postgresql.conf


shared_buf="shared_buffers = 1024MB"
val=$(grep -E -o "^shared_buffers = ([0-9]+)MB" $pg_conf |grep -E  -o "([0-9]+)") 
if [ $? -eq 0 ]
then
    echo "Match, and get MB number "$val
else
    echo "Mismatch"
fi

ip="121MB"
if [[ $shared_buf =~ ([0-9]+)MB ]];then
	echo "match"
	echo ${BASH_REMATCH[1]}
else
	echo "Not match"
fi
