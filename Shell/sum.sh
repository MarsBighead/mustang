#!/bin/bash
echo There are $# parameters supplied.
echo "The end from  element 2(include 2): ${@:2}";

if [ $# -ne 2 ]
then
    echo Usage: sum.sh a b
else 
    total=$[ $1 + $2 ]
    echo THe total is $total
fi

echo The last parameter was ${!#}
