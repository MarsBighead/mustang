#!/usr/bin/python
#iterator protocal

class TestIterator:
    value = 0
    def next(self):
        self.value += 1
        if self.value>10: raise StopIteration
        return self.value
    def __iter__(self):
        return self

ti = TestIterator()
print list(ti)

