#!/bin/python3

import math
import os
import random
import re
import sys
'''
Recursive Digit Sum


https://www.hackerrank.com/challenges/recursive-digit-sum/problem
'''
# Complete the superDigit function below.
def superDigit(n, k):
    if int(n)<10:
        return int(n)
    else:
        v=0
        for dig in n:
            
            v+=int(dig)
        if k>1:
            v=v*k
            print("v=",v,n)
            return superDigit(str(v), 1)
        else:        
            print("k=1, v=",v,n)
            return superDigit(str(v), 1)



if __name__ == '__main__':

    nk = "148 3".split()

    n = nk[0]

    k = int(nk[1])

    result = superDigit(n, k)
    print(result)

    