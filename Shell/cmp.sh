#!/bin/bash
a=a1
num=3
for value in a1 a1 a1
do
    echo $value
	if [ $a == $value ];then
	    let num=num-1
	fi
done
echo $num