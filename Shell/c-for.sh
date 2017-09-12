#!/bin/bash
for ((a=1, b=10; a<=10; a++,b--))
do
    echo "$a - $b "
done

for ((a = 1 ; a <= 3; a ++ ))
do
    echo "Starting loop $a:"
    for ((b = 1; b <= 3; b++))
    do
	    echo "    inside loop: $b"
    done
done
