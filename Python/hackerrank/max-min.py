#!/bin/python3

import math
import os
import random
import re
import sys

'''
Max Min

https://www.hackerrank.com/challenges/angry-children/problem
'''
# Complete the maxMin function below.
def maxMin(k, arr):
    arr=sorted(arr)
    if len(arr)<=k:
        return arr[-1]-arr[0]
    unfairness=arr[k-1]-arr[0]
    print(unfairness)
    for i in range(k,len(arr)):
        print(i-k,i)
        if arr[i]-arr[i-k+1]<unfairness:
            unfairness=arr[i]-arr[i-k+1]
            print(unfairness)

    print(arr)
    print("unfairness=",unfairness)
    return unfairness 

if __name__ == '__main__':
    n = 7
    k = 3
    arr = [ 100,
200,
300,
350,
400,
401,
402]

    result = maxMin(k, arr)

    
