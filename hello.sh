#!/bin/bash

read inputString # get a line of input from stdin and save it to our variable

# Your first line of output goes here
echo 'Hello, World.'

# Write the second line of output
echo $inputString


ssh  ubuntu@localhost /bin/bash -c 'echo "">/home/ubuntu/.bash_history'
