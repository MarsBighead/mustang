#!/bin/python3

import math
import os
import random
import re
import sys

'''
Two Strings

https://www.hackerrank.com/challenges/two-strings/problem
'''

# Complete the twoStrings function below.
def twoStrings(s1, s2):
    ds={}
    for char in s1:
        if ds.get(char)==None:
            ds[char]=1
        else:
            ds[char]+=1
    for char in s2:
        if ds.get(char)==None:
            continue
        else:
            return "YES"
    return "NO"

if __name__ == '__main__':

    s1 = "hello"
    s2 = "world"
    result = twoStrings(s1, s2)
    ss1=set( [v for v in s1 ])
    ss2=set( [v for v in s2 ])
    #print(ss1,ss2)
    if ss1.intersection(ss2):
        print("YES")
    else:
        print("NO")

