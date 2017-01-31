#!/usr/bin/python

def fibs(num):
    'calculates given numbers Fibonacci'
    result = [0,1]
    for i in range(num-2):
        result.append(result[-2]+result[-1])
    return result

num = input('How many fibonacci numbers  do  you want? ')
print fibs(num)
