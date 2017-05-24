#!/usr/bin/python
import sys
text = sys.stdin.read()
print 'Text:',text
words = text.split()
print 'Words:',words
wordcount = len(words)
print 'Wordcount:',wordcount

