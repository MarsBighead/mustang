#!/usr/bin/python
#hello.py

#print 'Hello world'

'''
[hbu@hbu packages]$ python
Python 2.7.5 (default, Nov  6 2016, 00:28:07) 
[GCC 4.8.5 20150623 (Red Hat 4.8.5-11)] on linux2
Type "help", "copyright", "credits" or "license" for more information.
>>> import hello
Hello world
Before update with __main__
'''

def hello():
    print "Hello world"
'''
[hbu@hbu packages]$ python
Python 2.7.5 (default, Nov  6 2016, 00:28:07) 
[GCC 4.8.5 20150623 (Red Hat 4.8.5-11)] on linux2
Type "help", "copyright", "credits" or "license" for more information.
>>> import hello
>>> hello.hello()
Hello world
>>> exit()
'''

#Running as a script, instead of a module
if __name__ == '__main__':hello()
