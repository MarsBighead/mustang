#!/bin/bash

list=$(ls)
echo ${list[*]}
nums=(1 2 3 4)
echo Length of \$nums is ${#nums[*]}
echo "\${nums[2]} is ${nums[2]}"
nums[7]=18
nums[6]=17
echo Length of \$nums is ${#nums[*]}
echo ${!nums[*]}
echo Length of \$nums is ${!nums[*]}
for i in ${!nums[*]}
do
   echo index = $i, value = ${nums[$i]}
done

