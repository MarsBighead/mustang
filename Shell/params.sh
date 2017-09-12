#!/bin/bash
echo There are $# parameters supplied.

if [ $# -ne 2 ]
then
    echo Usage: params.sh a b
else 
    total=$[ $1 + $2 ]
    echo THe total is $total
fi

