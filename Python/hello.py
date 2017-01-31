#!/usr/bin/python
print 'Hello world'

numbers = [1,2,3,4,5,6,7,8,9,10]
print numbers[-3:]

#str=','.join(numbers[-3:])
numbers.reverse()
numbers.remove(1)
numbers.append(10)
tmp=[22,3,4]
numbers.extend(tmp)
print numbers
str = "Pi with three decimals"
# tmp = map(str,tmp)
