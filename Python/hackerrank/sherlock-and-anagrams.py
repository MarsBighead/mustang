#!/bin/python3

import math
import os
import random
import re
import sys

'''
Sherlock and Anagrams

https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem
'''
# Complete the sherlockAndAnagrams function below.
def sherlockAndAnagrams(s):
    length=len(s)
    d,count={},0
    #n=i+1 is the length of sub string, it's really a great method to get all substrings
    for i in range(length):
        for j in range(length-i):
            sub=''.join(sorted(s[j:j+i+1]))
            print (sub," ", end="")
            if d.get(sub)==None:
                d[sub]=0
            else:
                d[sub]+=1
        print()
    for key in d:
        count+=d[key]*(d[key]+1)//2
    return count

if __name__ == '__main__':
    q = 2
    s = "cdcd"
    result = sherlockAndAnagrams(s)
    print("result=",result)

