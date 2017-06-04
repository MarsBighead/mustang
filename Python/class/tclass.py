#!/usr/bin/python
class Class:
    def method(self):
        print 'I have a self'
        self.name = 'Tom'
def function():
    print 'I don\'t...'

instance = Class()
instance.method()
print instance.name
instance.method= function
instance.method()

