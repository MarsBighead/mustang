#!/bin/python3

import math
import os
import random
import re
import sys

from collections import Counter

'''
Frequency Queries

https://www.hackerrank.com/challenges/frequency-queries/problem
forward index, inverted index
'''

def freqQuery1(queries):
    d={}
    freq={}
    result=[]

    for cmd, val in queries:
        if cmd==1:
            if val not in d: d[val]=1
            else: d[val]+=1
            prev_freq= d[val]-1
            if d[val] not in freq:
                freq[d[val]]=1
            else:
                freq[d[val]]+=1 
            if prev_freq in freq:
                freq[prev_freq]-=1
                
        elif cmd==2:
            if val in d and d[val]>0:
                d[val]-=1
                prev_freq= d[val]+1
                if d[val] not in freq:
                    freq[d[val]]=1
                else:
                    freq[d[val]]+=1 
                if prev_freq in freq:
                    freq[prev_freq]-=1
        elif cmd==3:
            flag=0
            if val in freq and freq[val]>0:
                flag=1
            result.append(flag)
    #print(result)
    return result

# Complete the freqQuery function below.
def freqQuery(queries):
    d=Counter()
    freq=Counter()
    result=[]
    #print(queries)
    for cmd, val in queries:
        if cmd==1:
            d[val]+=1
            prev_freq=d[val]-1
            if freq[prev_freq]>0: freq[prev_freq]-=1
            freq[d[val]]+=1
        elif cmd==2:
            if d[val]>0:
                d[val]-=1
                prev_freq= d[val]+1
                if freq[prev_freq]>0: freq[prev_freq]-=1
                freq[d[val]]+=1 
        elif cmd==3:
            flag=0
            if freq[val]>0:
                flag=1
            result.append(flag)
    #print(result)
    return result

if __name__ == '__main__':

    q = 10

    queries = [
        [1,3],
        [2,3],
        [3,2],
        [1,4],
        [1,5],
        [1,5],
        [1,4],
        [3,2],
        [2,4],
        [3,2]
    ]


    ans = freqQuery(queries)
    print("\n".join(map(str,ans)))

   