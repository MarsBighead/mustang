#!/bin/python3

import math
import os
import random
import re
import sys

from collections import Counter

'''
Reverse Shuffle Merge

https://www.hackerrank.com/challenges/reverse-shuffle-merge/problem
'''

# Complete the reverseShuffleMerge function below.
def reverseShuffleMerge(s):
    #reverse the charactor
    s = s[::-1]

    cnt = Counter(s)
    print(cnt,s)
    need = {k: cnt[k]//2 for k in cnt.keys()}
    left = {k: cnt[k] for k in cnt.keys()}
    res = []
    for c in s:
        left[c] -= 1
        if not need[c]:
            continue
        need[c] -= 1
        while res and res[-1] > c and left[res[-1]] > need[res[-1]]:
            print(res[-1], c, left)
            need[res[-1]] += 1
            res.pop()
        res.append(c)
    print(res)
    return ''.join(res)

if __name__ == '__main__':
    s = "abcdefgabcdefg"
    s = "eggegg"
    result = reverseShuffleMerge(s)
    print(result)