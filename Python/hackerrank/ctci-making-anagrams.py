#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the makeAnagram function below.
def makeAnagram(a, b):
    d={}
    s=set()
    for char in a:
        s.add(char)
        if char not in d:
            d[char]=1
        else:
            d[char]+=1
    d1={}
    s1=set()
    for char in b:
        s1.add(char)
        if char not in d1:
            d1[char]=1
        else:
            d1[char]+=1
    for char in s.intersection(s1):
        if d[char]==d1[char]:
            d.pop(char)
            d1.pop(char)
        else:
            n=abs(d[char]-d1[char])
            d.pop(char)
            d1.pop(char)
            d[char]=n
            #print("changed",char,d1[char])
    return sum(d1.values())+sum(d.values())

if __name__ == '__main__':

    a = "cde"

    b = "abc"

    res = makeAnagram(a, b)
    print(res)


