#!/usr/bin/python
import os

print os.getcwd()
os.chdir(os.getcwd()+'/../')
print os.getcwd()
os.chdir('/home/hbu')
print os.getcwd()
