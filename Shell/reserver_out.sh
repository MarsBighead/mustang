#!/bin/bash
ls -tl | awk '{if ( $9!="" ) {print $9}}' |tac


ls  -tr | awk '{print $1 }'

