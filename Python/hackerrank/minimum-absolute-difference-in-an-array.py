#!/bin/python3

import math
import os
import random
import re
import sys

def merge_sort(arr):
    if len(arr) == 1: 
        return arr
    mid = len(arr) // 2
    left = arr[:mid]
    right = arr[mid:]
    return merge(merge_sort(left), merge_sort(right))


def merge(left, right):
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
            j+=1
    #if len(right)==2: print("end i=%d, j=%d" % (i,j))
    result += left[i:]
    result += right[j:]
    return result

# Complete the minimumAbsoluteDifference function below.
def minimumAbsoluteDifference(arr):
    arr=merge_sort(arr)
    n=abs(arr[1]-arr[0])
    for i in range(1,len(arr)):
        m=abs(arr[i]-arr[i-1])
        if m<n:
            n=m
    return n

if __name__ == '__main__':
    n=5
    arr=[1, -3, 71, 68, 17]
    result=minimumAbsoluteDifference(arr)
    print("max=%d" % (result))