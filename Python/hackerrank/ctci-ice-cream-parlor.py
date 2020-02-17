#!/bin/python3

import math
import os
import random
import re
import sys

'''
Hash Tables: Ice Cream Parlor

https://www.hackerrank.com/challenges/ctci-ice-cream-parlor/problem
'''

# Complete the whatFlavors function below.
def whatFlavors(cost, money):
    search={}
    for i,v  in enumerate(cost):
        if search.get(v)!=None :
            print("money=",money-v, v)
            print(search[v], i+1)
            return
        else:
            search[money-v]=i+1
    return
    #print(search)



if __name__ == '__main__':

    money = 4
    n = 4
    cost = list(map(int, "1 4 5 3 2".split()))
    whatFlavors(cost, money)
