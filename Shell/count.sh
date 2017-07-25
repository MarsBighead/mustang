#!/bin/bash

filename="inter.sh"
egrep -o "\b[[:alpha:]]+\b"  $filename


echo "###Words count###"
egrep -o "\b[[:alpha:]]+\b"  $filename | \
awk '{count[$0]++}
END{ printf("%-14s%s\n","Word","Count");
for(ind in count) { printf("%-14s%d\n",ind,count[ind]);}
}'
