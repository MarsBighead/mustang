#!/bin/bash
var1=100
until [ $var1 -eq 0 ]
do
    echo $var1
    (( var1 = $var1 - 2 ))
    var1=$[ $var1 - 23 ]
done
