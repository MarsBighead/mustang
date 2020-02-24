#!/bin/python3

import math
import os
import random
import re
import sys

from bisect import insort, bisect_left

'''
Triple sum

https://www.hackerrank.com/challenges/triple-sum/problem
'''

# Complete the triplets function below.
def triplets1(a, b, c):
    a=sorted(set(a))
    b=sorted(set(b))
    c=sorted(set(c))
    return sum((bisect(a, x) * bisect(c, x) for x in b))

# Complete the triplets function below.
def triplets(a, b, c):
    a=sorted(set(a))
    b=sorted(set(b))
    c=sorted(set(c))
    count=0
    n,m=0,0
    for q in b:
        for p in a[n:]:
            if p<=q: n+=1
            else: break
        for r in c[m:]:
            if r<=q: 
                m+=1
            else: 
                break
        count+=n*m
        #print("q, m, n",q,m,n)
    print("count=",count)
    return count
        


if __name__ == '__main__':
    arra = list(map(int, "1 4 5".rstrip().split()))

    arrb = list(map(int, "2 3 3".rstrip().split()))

    arrc = list(map(int, "1 2 3".rstrip().split()))

    ans = triplets(arra, arrb, arrc)
