#!/bin/bash

#Extract the largest number in column 3, whole line
cat data.txt |awk -F  "," '{print $1"\t"$3;}' | sort -k2,2nr | awk '!a[$3]++'
#a	5
#Extract the largest number in column 3, special value
cat data.txt |awk -F  "," '{print $1"\t"$3;}' | sort -k2,2nr | awk '!a[$3]++' | awk '{print $2}'
#5

#sort -t $',' : split text content with ','
#uniq -w :  compare no more than N characters in lines
#           1 charater from the line head
sort -t $',' -k 1 -u data.txt | awk -F "," '{print $1,$2}' | uniq -w 1
#1 2
#2 3
#a x
#           3 charater from the line head
sort -t $',' -k 1 -u data.txt | awk -F "," '{print $1,$2}' | uniq -w 3
#1 2
#2 3
#a x
#azx y
sort -t $',' -k 1 -u data.txt | awk -F "," '{print $1,$2}' | uniq -w 6
#1 2
#2 3
#a x
#azx y
#azx z

