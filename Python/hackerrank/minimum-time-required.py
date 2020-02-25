#!/bin/python3

import math
import os
import random
import re
import sys

'''
Minimum Time Required

https://www.hackerrank.com/challenges/minimum-time-required/problem
'''
# Complete the minTime function below.
def minTime(machines, goal):
    minday = math.ceil(goal / len(machines)) * min(machines)
    maxday = math.ceil(goal / len(machines) * max(machines))
    while minday < maxday:
        day = (maxday + minday) // 2
        if sum(day // m for m in machines) >= goal:
            maxday = day
        else:
            minday = day + 1
    return minday

if __name__ == '__main__':

    nGoal = "3 12".split()

    n = int(nGoal[0])

    goal = int(nGoal[1])

    machines = list(map(int, "4 5 6".split()))

    ans = minTime(machines, goal)
    print("min time=", ans)
  