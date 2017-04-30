#!/usr/bin/python


import sys


n = int(raw_input().strip())

for i in range (1,11):
    print "%d x %d = %d"%(n, i, n*i)
