#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the alternatingCharacters function below.
def alternatingCharacters(s):
    cursor=s[0]
    #print(cursor)
    cnt=0
    for char in s[1:]:
        if char==cursor:
            cnt+=1
        else:
            cursor=char
    return cnt


if __name__ == '__main__':

    q = 10
    s="AAAA"

    result = alternatingCharacters(s)
