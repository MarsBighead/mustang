#!/bin/python3

import math
import os
import random
import re
import sys

'''
Greedy Florist

https://www.hackerrank.com/challenges/greedy-florist/problem
'''

# Complete the getMinimumCost function below.
def getMinimumCost(k, c):
    c=sorted(c, reverse=True)
    print(c)
    minimumCost=0
    multiply=0
    for i in range(len(c)):
        if i%k==0:
            multiply+=1
        minimumCost+=multiply*c[i]
    return minimumCost

if __name__ == '__main__':

    nk = "5 3".split()

    n = int(nk[0])

    k = int(nk[1])

    c = list(map(int, "1 3 5 7 9".rstrip().split()))

    minimumCost = getMinimumCost(k, c)

    print(minimumCost)
