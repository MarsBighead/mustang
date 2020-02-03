#!/bin/python3

import math
import os
import random
import re
import sys

'''
Hash Tables: Ransom Note

https://www.hackerrank.com/challenges/ctci-ransom-note/problem
'''

# Complete the checkMagazine function below.
def checkMagazine(magazine, note):
    dm={}
    for word in magazine:
        if dm.get(word)==None:
            dm[word]=1
        else:
            dm[word]+=1
            
    for word in note:
        if dm.get(word)==None:
            print("No")
            return
        else:
            dm[word]-=1
        if dm[word]<0:
            print("No")
            return
    print("Yes")
    return



if __name__ == '__main__':
    n = 6
    d = 4
    magazine = ["give", "me", "one", "grand", "today", "night"]
    note=["give", "one", "grand", "today"]

    checkMagazine(magazine, note)
