#!/bin/bash
echo -n "Enter your name: "
read name
echo "Hello $name, welcome to my program!"
read -p "Enter your name: " first last
echo "Checking data fro $last, $first..."
read
#$REPLY is env variable for read
echo "REPLY $REPLY"
