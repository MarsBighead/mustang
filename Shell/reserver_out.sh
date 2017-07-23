#!/bin/bash

#tac reverse output
ls -tl | awk '{if ( $9!="" ) {print $9}}' |tac

#Remove blank rows with awk NF
ls -lt | awk  '{print $9}' | awk NF | tac

#Output with ls parameter -r
ls  -tr | awk '{print $1 }'

