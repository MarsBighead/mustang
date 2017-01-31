#!/usr/bin/python
fibs = [0,1]
num = input('How many fibonacci numbers  do  you want? ')
for i in range(num-2):
    fibs.append(fibs[-2]+fibs[-1])
    # print fibs[i]

print fibs
