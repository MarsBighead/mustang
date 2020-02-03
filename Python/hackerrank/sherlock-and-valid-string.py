#!/bin/python3

import math
import os
import random
import re
import sys
'''
Sherlock and the Valid String

https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem
'''
# Complete the isValid function below.
def isValid(s):
    d={}
    for char in s:
        if d.get(char) is None:
            d[char]=1
        else:
            d[char]+=1
    same={}
    for k in d.keys():
        v=d[k]
        if same.get(v) is None:
            same[v]=[k]
        else:
            same[v].append(k)
    cnt=len(same.keys())
    print("count=%d" %(cnt))
    if cnt==1:
        return "YES"
    else:
        if cnt==2:
            k1,k2=sorted(same.keys())
            print(k1,k2)
            if k1 == 1 and len(same[k1])==1: return "YES"
            if k2-k1==1 and len(same[k2])==1: return "YES"
            return "NO"
        else:
            return "NO"
                     
if __name__ == '__main__':
    s = "ibfdgaeadiaefgbhbdghhhbgdfgeiccbiehhfcggchgghadhdhagfbahhddgghbdehidbibaeaagaeeigffcebfbaieggabcfbiiedcabfihchdfabifahcbhagccbdfifhghcadfiadeeaheeddddiecaicbgigccageicehfdhdgafaddhffadigfhhcaedcedecafeacbdacgfgfeeibgaiffdehigebhhehiaahfidibccdcdagifgaihacihadecgifihbebffebdfbchbgigeccahgihbcbcaggebaaafgfedbfgagfediddghdgbgehhhifhgcedechahidcbchebheihaadbbbiaiccededchdagfhccfdefigfibifabeiaccghcegfbcghaefifbachebaacbhbfgfddeceababbacgffbagidebeadfihaefefegbghgddbbgddeehgfbhafbccidebgehifafgbghafacgfdccgifdcbbbidfifhdaibgigebigaedeaaiadegfefbhacgddhchgcbgcaeaieiegiffchbgbebgbehbbfcebciiagacaiechdigbgbghefcahgbhfibhedaeeiffebdiabcifgccdefabccdghehfibfiifdaicfedagahhdcbhbicdgibgcedieihcichadgchgbdcdagaihebbabhibcihicadgadfcihdheefbhffiageddhgahaidfdhhdbgciiaciegchiiebfbcbhaeagccfhbfhaddagnfieihghfbaggiffbbfbecgaiiidccdceadbbdfgigibgcgchafccdchgifdeieicbaididhfcfdedbhaadedfageigfdehgcdaecaebebebfcieaecfagfdieaefdiedbcadchabhebgehiidfcgahcdhcdhgchhiiheffiifeegcfdgbdeffhgeghdfhbfbifgidcafbfcd" 
    #s = "aabbcccddd"
    result = isValid(s)
    print (result)
