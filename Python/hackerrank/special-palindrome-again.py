#!/bin/python3

import math
import os
import random
import re
import sys
import re
from collections import Counter

'''
Special String Again

https://www.hackerrank.com/challenges/special-palindrome-again/problem
'''
def substrCount2(n, s):
    answer = 0
    prev = s[0]
    repeat = False
    middle = False
    repeat_len = 0  # Number of times a char repeats
    comp_index = 0  # Compare current index to this index
    num_middle = 0  # Number of special string where middle is different

    # Loop only checks for substrings of repeating characters or with middle char 
    # different from the rest.
    for i in range(1,len(s)):
        if s[i] == prev:
            if repeat:
                # Continue to count the length
                repeat_len += 1  
            else:
                repeat = True
                repeat_len = 2
                
            if middle:
                if comp_index >= 0:
                    if s[comp_index] == s[i]:
                        comp_index -= 1
                        num_middle += 1         
            prev = s[i]
        else:
            # Middle case or repeat stopped
            if repeat:
                repeat = False
                answer += int(((repeat_len*(repeat_len+1))/2) - repeat_len)
                repeat_len = 0

            if middle:
                middle = False
                answer += num_middle
                num_middle = 0
                
            if i-2 >= 0:
                if s[i] == s[i-2]:
                    # Middle case
                    middle = True 
                    num_middle = 1
                    comp_index = (i-3)
            prev = s[i]

    # Check if loop ended on special substring.
    if repeat:
        answer += int(((repeat_len*(repeat_len+1))/2) - repeat_len)
    if middle:
        answer += num_middle

    return answer + len(s)

# Complete the substrCount function below.
def substrCount1(n, s):
    s1=Counter()
    s2=Counter()
    count=0
    #print(prev,mid,nex)
    prev=s[0]
    cursor=[]
    last=0
    for i, char in enumerate(s):
        #prev,mid,nex=s[i],s[i+1],s[i+2]
        #print(prev,mid,nex)
        if i>0: prev=s[i-1]
        #print(cursor, char)
        if  len(cursor)>0 and char==cursor[0]:
            s2[cursor[1]]+=1

        if char!=prev:
            #first palindrom
            if len(cursor)==0:
                cursor=[prev, char, s1[prev]]
            else:
            # second one
                print(cursor,char, i)
                if char!=cursor[0]:
                    count+=min(s2[cursor[1]], cursor[2])
                    #x=s2[cursor[1]]
                    #if x>cursor[2]:
                    #    print(s[i-x-cursor[2]-4:i+1],s2[cursor[1]], cursor, s1[prev])
                    if s2[cursor[1]]!=0:last=min(s2[cursor[1]], cursor[2])
                    s2[cursor[1]]=0
                    cursor=[prev, char, s1[prev]]
                else:
                    pass
                
            s1[prev]=0
        else:
            pass
                    
        s1[char]+=1
        count+=s1[char]
    count+=last
    print("count=",count)
    return count

# Complete the substrCount function below.
def substrCount(n, s):
    count=0
    #n=i+1 is the length of sub string, it's really a great method to get all substrings
    for i in range(n):
        for j in range(n-i):
            sub=s[j:j+i+1]
            ns=i+1
            #print("sub=",sub)
            if ns%2!=0 and ns>1:
                mid=ns//2
                sub=sub[:mid]+sub[mid+1:]
            std=sub[0]*len(sub)
            if sub==std:
                count+=1
            
    print("count=%d"%(count))
    return count
if __name__ == '__main__':
    
    with open(os.environ['HOME']+'/input05.txt', 'r', encoding='utf-8') as f:
        line=f.readline()
        #print(len(line))
        cnt = 1
        while line:
            
            content=line.strip()
            if len(content)<10 and len(content)>0:
                print("num=%s"%(content))
            else:
                s=content
            line = f.readline()
            cnt += 1
    
    #print(s)
    #s="asasd"
    s="abcbabababaab"
    n= len(s)
    print("meter nums=",n)
    result = substrCount2(n, s)
    print("reference=",result)
    result = substrCount1(n, s)
    print(s)
    print([i for i,v in enumerate(s)])
