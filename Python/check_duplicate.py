#!/usr/bin/python

#1-5, Time efficiency O(n), Space efficency O(1)
numbers=[1,3,2,2,5]
n=5
def checker(numbers, n):
    print numbers,n
    for i in range(0, n):
        if numbers[i] == i+1:
            pass
        else:
            temp = numbers[i];
            numbers[i] = numbers[temp-1];
            numbers[temp-1] = temp;
    print numbers,n
    for i in range(0, n):
        if numbers[i] != i+1:
           print "Duplicated"
           return True
    print "Non-duplicated"
    return False
    
isDuplicate=checker(numbers,n)
isDuplicate=checker([2,1,3,5,4],n)
