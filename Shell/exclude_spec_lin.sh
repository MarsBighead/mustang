#!/bin/bash

#Output from line two
awk 'NR>1'  data.txt 

echo "exclued line azx,z,1,b with regular expression"
awk '!/azx,z,1,b/' data.txt

