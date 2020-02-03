#!/bin/python3

import math
import os
import random
import re
import sys

'''
New Year Chaos

https://www.hackerrank.com/challenges/new-year-chaos/problem
'''

def merge_sort(cnt, arr):
    if len(arr) == 1: 
        return cnt, arr
    mid = len(arr) // 2
    left = arr[:mid]
    right = arr[mid:]
    cnt,left=merge_sort(cnt, left)
    #print ("left num: ", cnt, left)
    cnt,right=merge_sort(cnt, right)
    #print ("right num: ", cnt, right)
    return merge(left, right, cnt)


def merge(left, right, cnt):
    result = []
    lnum=len(left)
    rnum=len(right)
    i,j=0,0
    while lnum>i and rnum>j:
        #if len(left)==2: print("left=%d, right=%d" % (left[i],right[j]))
        if left[i] <= right[j]:
            result.append(left[i])
            i+=1
        else:
            result.append(right[j])
            cnt+=lnum-i
            j+=1
    #if len(right)==2: print("end i=%d, j=%d" % (i,j))
    result += left[i:]
    result += right[j:]
    return cnt, result
    
# Complete the minimumBribes function below.
def minimumBribes(q):
    bribe=0
    num=len(q)
    if len(q) == 1:
        print(bribe)
        return 
    for idx in range(num):
        val=q[idx]
        if val-idx-1>2:
            print("Too chaotic")
            return
    bribe, result=merge_sort(bribe, q)       #bribe+=skip
    print (result)
    print(bribe)
    return


if __name__ == '__main__':
    q=[2, 1, 5, 3, 4]
    #q=[2, 5, 1, 3, 4]
    q=[1, 2, 5, 3, 7, 8, 6, 4]
    minimumBribes(q)
