#!/bin/python3

import math
import os
import random
import re
import sys
from bisect import insort, bisect_left

'''
Direct solution
num=30171, days= 10000
38

real    0m34.166s
user    0m34.013s
sys     0m0.079s

https://www.hackerrank.com/challenges/fraudulent-activity-notifications/problem
'''
# Complete the activityNotifications function below.
def activityNotifications(expenditure, d):
    num=len(expenditure)-d
    print  (num)
    result,midx1,midx2=0,0,0
    if d%2==1:
        midx1=midx2=int(d/2)
    else:
        midx1,midx2=int(d/2)-1,int(d/2)
    ls=sorted(expenditure[0:d])
    for i in range(num):
        val=expenditure[i+d]
        if val>=ls[midx1]+ls[midx2]:
            result+=1
        else:
            pass
        idx=bisect_left(ls,expenditure[i])
        del ls[idx]
        insort(ls, val)
    return result


if __name__ == '__main__':
    d = 10
    expenditure=[0,82,180,55,168,41,179,59,139,71,6,12,135,139,73,157,4,7]
    result = activityNotifications(expenditure, d)
    print (result)