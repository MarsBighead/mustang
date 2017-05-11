#!/usr/bin/python

import sys

def number(s):
    try:
        print int(s)
    except ValueError:
        print 'Bad String'

   
#s = raw_input().strip()
 
a="2"
number(a)
a="ab"
number(a)
