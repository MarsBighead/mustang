#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the flippingBits function below.
def flippingBits(n):
    flipping=2**32-1
    return n^flipping

if __name__ == '__main__':
    n = 0
    n = 802743475
    n = 35601423
    result = flippingBits(n)
    print(result)
