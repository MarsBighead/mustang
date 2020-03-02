#!/bin/python3

import math
import os
import random
import re
import sys

'''
Max Array Sum

https://www.hackerrank.com/challenges/max-array-sum/problem
Given an array of integers, find the subset of non-adjacent elements with the maximum sum.
'''

# Complete the maxSubsetSum function below.
def maxSubsetSum(arr):
    if len(arr)==1:return arr[0]
    if len(arr)==2:return max(arr)
    dp = []
    dp.append(arr[0])
    dp.append(max(arr[:2]))
    #print (arr[2:])
    for a in arr[2:]:
        #print(dp[-2]+a,a,dp[-1])
        dp.append(max([
            dp[-2]+a, 
            a, 
            dp[-1]
        ]))
    return dp[-1]


if __name__ == '__main__':

    n = 5

    arr = list(map(int, "3 5 -7 8 10".split()))
    #print("arr[-1]=",arr[-1])
    print("arr[:2]=",arr[:2])
    res = maxSubsetSum(arr)
    print(res)

