#!/bin/bash
seq 5 | awk 'BEGIN{sum=0;print "Summation:"}{print $1"+"; sum+=$1} END{print "=="; print sum}'

VAR=10000
echo | awk -v VARIABLE=$VAR '{print VARIABLE }'

seq 5 | awk 'BEGIN { getline; print "Read ahead first line", $0 } {print $0}'

awk 'BEGIN{ FS=","}{ print "data",$1,$2,$3,$4 }' data.txt

for i in {2..40};
do
    echo $i
done
