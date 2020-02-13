#!/bin/python3

import math
import os
import random
import re
import sys

import numpy as np

'''
Common Child

https://www.hackerrank.com/challenges/common-child/problem
'''
# Complete the commonChild function below.
def commonChild(s1, s2):
    len1 = len(s1)
    len2 = len(s2)
    c =[[ 0 for _ in range(len2+1)] for _ in range(len1+1)]
    for i in range(len1):
        for j in range(len2):
            if s1[i] == s2[j]:
                c[i + 1][j + 1] = c[i][j] + 1
            elif c[i + 1][j] > c[i][j + 1]:
                c[i + 1][j + 1] = c[i + 1][j]
            else:
                c[i + 1][j + 1] = c[i][j + 1]
        #print(c[i])
    return c[len1][len2]


def printLcs(flag, a, i, j):
    if i == 0 or j == 0:
        return
    if flag[i][j] == 'ok':
        printLcs(flag, a, i - 1, j - 1)
        print(a[i - 1], end='')
    elif flag[i][j] == 'left':
        printLcs(flag, a, i, j - 1)
    else:
        printLcs(flag, a, i - 1, j)
    
    

if __name__ == '__main__':

    s1 = "SHINCHAN"

    s2 = "NOHARAAA"
    
    result = commonChild(s1, s2)
    print(result)
