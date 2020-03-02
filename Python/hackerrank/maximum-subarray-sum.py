#!/bin/python3

import math
import os
import random
import re
import sys
from bisect import insort, bisect_right

# Complete the maximumSum function below.
def maximumSum(a, m):
    print(a,m)
    # Create prefix tree
    prefix =[]
    p = 0
    for v in a:
        p = (v % m + p) % m
        prefix.append(p) 
    prefix=sorted(prefix, reverse=True)
    maxmodsum = prefix[0]
    pq=[maxmodsum]
    for i in range(1, len(a)):
        # Find cheapest prefix larger than prefix[i]
        left = bisect_right(pq, prefix[i])
        print("left",left,pq)
        if left != len(pq):
            # Update maxmodsum if possible
            modsum = (prefix[i] - pq[left] + m) % m
            maxmodsum = max(maxmodsum, modsum)
            print(pq)

        # add current prefix to heap
        insort(pq, prefix[i])
    print("maximumSum=",maxmodsum)
    return maxmodsum

if __name__ == '__main__':

    q = 1

    for q_itr in range(q):
        nm = "5 7".split()

        n = int(nm[0])

        m = int(nm[1])

        a = list(map(int, "3 3 9 9 5".split()))

        result = maximumSum(a, m)


