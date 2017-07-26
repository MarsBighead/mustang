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
sort -t $',' -k 1 -u data.txt | awk -F "," 'BEGINE{print $1,$2} END{print "lines: "NR}' | uniq -w 6
#1 2
#2 3
#a x
#azx y
#azx z

#uniq by column 1, with out considering character width
sort -t $"," -k1 data.txt | awk -F "," '{if (!keys[$1]) print $0; keys[$1] = 1;}'

paste  col1 col2 -d ","
#[hbu@hbu Shell]$ paste  col1 col2 -d ","
#duan,24
#annan,31
#nil,78
#onecloud,23
paste  col1 col2 -d " "
#[hbu@hbu Shell]$ paste  col1 col2 -d " "
#duan 24
#annan 31
#nil 78
#onecloud 23
#Default connect with "\t"

cut -f2,3 -d " " data.sort.txt 
#[hbu@hbu Shell]$ cut -f2,3 -d " " data.sort.txt 
#mac 2000
#2 winxp
#bsd 1000
#linux 1000

paste  col1 col2 -d "\t"
#[hbu@hbu Shell]$ paste  col1 col2 -d "\t"
#duan	24
#annan	31
#nil	78
#onecloud	23

