#!/usr/bin/python

name=''

while not name.strip():
# while not name or name.isspace():
    name =raw_input('Please enter your name: ')
print 'Hello, %s!' % name
