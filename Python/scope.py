#!/usr/bin/python

class Difference:
    def __init__(self, a):
        self.__elements = a
    def computeDifference(self):
        a = self.__elements
        max = a[0]
        min = a[0]
        for i in range(0, len(a)):
           if max<a[i]:
              max=a[i]
           else:
              pass
           if min>a[i]:
              min=a[i]
           else:
              pass
        print 'max = %d, min = %d'%(max,min)
        self.maximumDifference= max-min

_ = raw_input()
a = [int(e) for e in raw_input().split(' ')]

d = Difference(a)
d.computeDifference()
# d.maximunDifference is a variable
print d.maximumDifference
