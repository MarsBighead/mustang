#!/bin/python3

import math
import os
import random
import re
import sys

'''
Pairs

https://www.hackerrank.com/challenges/pairs/problem

Constraints

each integer arr[i] will be unique
'''

# Complete the pairs function below.
def pairs(k, arr):
    arr=sorted(arr)
    d={}
    count=0
    for v in arr:
        if d.get(v)!=None:
            count+=1
        d[v+k]=1
        print(v)
    print("count=",count)
    return count

if __name__ == '__main__':

    nk = "5 2".split()

    n = int(nk[0])

    k = int(nk[1])

    arr = list(map(int, "1 5 3 4 2".split()))

    result = pairs(k, arr)
    print (result)

