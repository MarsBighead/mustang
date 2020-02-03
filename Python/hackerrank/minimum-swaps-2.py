#!/bin/python3

import math
import os
import random
import re
import sys

'''
Minimum Swaps 2

https://www.hackerrank.com/challenges/minimum-swaps-2/problem
'''

def minimumSwaps(arr):
    count=0
    length=len(arr)
    d={v:i for i, v in enumerate(arr)}
    for i in range(length):
        val=arr[i]
        if val==i+1: 
            continue
        else:
            j=d[i+1]
            arr[i]=arr[j]
            arr[j]=val
            d[i+1]=i
            d[val]=j
            count+=1
    return count


if __name__ == '__main__':
    q=[2, 1, 5, 3, 4]
    #q=[2, 5, 1, 3, 4]
    q=[4, 3, 1, 2]
    q=[2, 3, 4, 1, 5]
    #print(q)
    res=minimumSwaps(q)
    print (res)
    
