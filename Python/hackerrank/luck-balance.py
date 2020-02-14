#!/bin/python3

import math
import os
import random
import re
import sys

'''
Luck Balance

https://www.hackerrank.com/challenges/luck-balance/problem
'''

# Complete the luckBalance function below.
def luckBalance(k, contests):
    arr=[] #length=k sorted list
    lb=0
    for l,t in contests:
        if t==1:
            arr.append(l)
        else:
            lb+=l
    arr=sorted(arr)
    #print(arr)
    if len(arr)<=k:
        for l in arr:
            lb+=l
    else:
        length=len(arr)
        for i in range(length):
            if i<length-k:
                lb-=arr[i]
            else:
                lb+=arr[i]

    return lb
            


if __name__ == '__main__':


    n = 6

    k = 3

    contests = [
        [5, 1],
        [2, 1],
        [1, 1],
        [8, 1],
        [10, 0],
        [5, 0]
    ]

    result = luckBalance(k, contests)
    print(result)
