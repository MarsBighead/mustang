#!/bin/bash

for (( a=1; a<10; a++ ))
do
    echo "The number is $a"
done > number.txt
echo "The command is finished."

for state in "North Dakota" Connecticut Illions Alabma Tennessee
do
    echo "$state is the next place to go..."
done | sort
echo "This completes our travels"

