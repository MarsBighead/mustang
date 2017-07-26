#!/bin/bash
var="This is a line of text"
echo ${var/line/REPLACED/}

string=abcdefghigklmnopqrstuvwxyz
#[hbu@hbu ~]$ string=abcdefghigklmnopqrstuvwxyz
echo ${string:4}
#[hbu@hbu ~]$ echo ${string:4}
#efghigklmnopqrstuvwxyz
echo ${string:4:8}
#[hbu@hbu ~]$ echo ${string:4:8}
#efghigkl
echo ${string:(-1)}
#[hbu@hbu ~]$ echo ${string:(-1)}
#z
echo ${string:(-2)}
#[hbu@hbu ~]$ echo ${string:(-2)}
#yz
echo ${string:(-2):1}
#[hbu@hbu ~]$ echo ${string:(-2):1}
#y
echo ${string:(-3):2}
#[hbu@hbu ~]$ echo ${string:(-3):2}
#xy

