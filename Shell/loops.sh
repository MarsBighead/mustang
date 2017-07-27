#!/bin/bash

for ip in 192.168.1.{1..255};
do
    (echo $ip )&
done

wait
