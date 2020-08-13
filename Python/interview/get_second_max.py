#!/usr/bin/python3


__all__ = ['hello']
'''
Alibaba simple algorithm to find the second maximum number from a int list
'''
def get_second_max(arr):
    if len(arr)<=1:
        return None
    elif len(arr)==2:
        return min(arr)
    coordinate=[arr[0], arr[0]]
    for a in arr[1:]:
        if a>coordinate[1]:
            coordinate[0]=coordinate[1]
            coordinate[1]=a
        else:
            if a>coordinate[0]:
                coordinate[0]=a
            else:
                pass
    return coordinate[0]
'''
print(get_second_max([]))
print(get_second_max([4]))
print(get_second_max([4,2]))
print(get_second_max([1,4,2,6,3,1,2]))
'''
def hello():
    print("Hello world")