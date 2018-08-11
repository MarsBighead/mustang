#!/usr/bin/python
import re
str = "a23b\na34b"

list1=re.findall(r"^a(\d+)b", str)
print list1
list2=re.findall(r"^a(\d+)b", str, re.M)
print list2

