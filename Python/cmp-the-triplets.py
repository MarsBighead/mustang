#!/bin/python

import math
import os
import random
import re
import sys

# Complete the solve function below.
def solve(a, b):
    m=0
    n=0
    for i in range(3):
        print "a=%d, b=%d"%(a[i],b[i])
        if a[i]>b[i]:
            n=n+1
        elif a[i]<b[i]:
            m+=1
    return m,n

if __name__ == '__main__':
    #fptr = open(os.environ['OUTPUT_PATH'], 'w')

    #a = map(int, raw_input().rstrip().split())

    #b = map(int, raw_input().rstrip().split())
    a = [6, 8, 12]
    b = [7, 9, 15]
    result = solve(a, b)
    print result
    #fptr.write(' '.join(map(str, result)))
   # fptr.write('\n')

    #fptr.close()
