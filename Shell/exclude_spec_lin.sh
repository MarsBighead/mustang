#!/bin/bash

#Output from line two
awk 'NR>1'  data.txt 

echo "exclued line azx,z,1,b with regular expression"
awk '!/azx,z,1,b/' data.txt

echo "Output line 2 with awk"
awk 'NR==2' data.txt 
#[hbu@hbu Shell]$ awk 'NR==2' data.txt 
#1,2,3,4
echo "Output line 2 with sed"
sed -n 2p  data.txt 
#[hbu@hbu Shell]$ sed -n 2p  data.txt 
#1,2,3,4

echo "output from line 3 to 4 with cat union tail and head"
cat  data.txt  | tail -n +3 | head -n 2

echo "output from line 3 to 4 with sed"
sed -n "3,4p" data.txt 

echo "output from line 3 to 4 with awk"
awk 'NR==3,NR==4' data.txt 
