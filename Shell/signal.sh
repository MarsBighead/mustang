#!/bin/bash

trap "echo 'Sorry! I have trapped Ctrl-C'" SIGINT SIGTERM
trap "echo 'Bye-bye'" EXIT 
echo "This is a test program for singanls"
count=1
while [ $count -le 10 ]
do
    echo "Loop #$count"
    sleep 2
    ((count = $count+1)) 
    count=$[ $count+1 ]
done
echo This is the end of the test program
trap - EXIT
echo "I justremoved the trap"

