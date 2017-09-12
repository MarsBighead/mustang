#!/bin/bash

for ((a = 1 ; a <= 3; a ++ ))
do
    echo "Starting loop $a:"
    for ((b = 1; b <= 100; b++))
    do
        if [ $b -gt 4 ]
        then
            if [ $a -gt 1 ]
            then
                break  2
            else
                break
            fi 
        fi 
	    echo "    inside loop: $b"
    done
done

#continue

for (( var1 = 1; var1 <= 15; var1++))
do
    if [ $var1 -gt 5 ] && [ $var1 -le 10 ]; then
        continue
    fi
    echo "Iteration numberL $var1"
done

