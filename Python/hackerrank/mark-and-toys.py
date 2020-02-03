#!/bin/python3

import math
import os
import random
import re
import sys

'''
Mark and Toys
https://www.hackerrank.com/challenges/mark-and-toys/problem

'''

# Complete the maximumToys function below.
def maximumToys(prices, k):
    count=0
    total=0
    for price in sorted(prices):
        total+=price
        count+=1
        if total>k:
            count-=1
        else:
            pass
    return count

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    nk = input().split()

    n = int(nk[0])

    k = int(nk[1])

    prices = list(map(int, input().rstrip().split()))

    result = maximumToys(prices, k)

    fptr.write(str(result) + '\n')

    fptr.close()
