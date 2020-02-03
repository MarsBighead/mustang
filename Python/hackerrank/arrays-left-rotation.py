#!/bin/python3

import math
import os
import random
import re
import sys

'''
Arrays: Left Rotation

https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem

'''

# Complete the rotLeft function below.
def rotLeft(a, d):
    return a[d:]+a[:d]


if __name__ == '__main__':
    n = 5
    d = 4
    a = [1, 2, 3, 4, 5]

    result = rotLeft(a, d)

    print(' '.join(map(str, result)))
    
