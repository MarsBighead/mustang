#!/bin/bash

read -n1 -p "Do you want to continue [Y/N]" answer
case $answer in
Y|y) echo
     echo "Fine, continue on...";;
N|n) echo
     echo Ok, goodbye
     exit;;
*)   echo 
     echo "Usage: [Y/N]";;
esac
echo "This is the end of the script $0"

