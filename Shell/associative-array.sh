#!/bin/bash

declare disk_mem
for i in a b c d e f; do
	disk_mem+=("sd$i")
done

echo Values: ${disk_mem[*]}
echo Index: ${!disk_mem[*]}
#echo Index of \$disk_mem is ${!disk_mem[*]}
for i in ${!disk_mem[*]}
do
   echo index = $i, value = ${disk_mem[$i]}
done
#for i in ${disk_mem[*]}
for i in ${disk_mem[@]}
do
   echo value = $i
done
unset disk_mem
echo Value of variable with UNSET: $disk_mem
