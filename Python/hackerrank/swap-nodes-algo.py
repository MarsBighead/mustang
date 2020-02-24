#!/bin/python3

import math
import os
import random
import re
import sys
import copy

class Node:
    def __init__(self, data):
        self.data = data
        self.right = None
        self.left = None


    def printTree(self):
        if self.left != None:
            self.left.printTree()
        if self.data != -1:
            print(self.data, end = ' ')

        if self.right != None:
            self.right.printTree()

    def rs(self):
        if self.left != None:
            self.left.rs()
        if self.data != -1:
            r.append(str(self.data))

        if self.right != None:
            self.right.rs()



def getNodes(level, curr_level, node,l):
    if node != None:
        if curr_level==level and node.data != -1:
            l.append(node)
        else:
            getNodes(level,curr_level+1,node.left,l)
            getNodes(level,curr_level+1,node.right,l)


def maxDepth(node):
    if node is None or node.data == -1:
        return 0
    else:
        lDepth = maxDepth(node.left) 
        rDepth = maxDepth(node.right) 
        if (lDepth > rDepth):
            return lDepth+1
        else: 
            return rDepth+1       



def swap(queries, newtemp):
    level = queries
    while(level < max_Depth):
        l = []
        getNodes(level, 1, newtemp, l)
        for x in l: 
            if x.data != -1:x.left, x.right = x.right, x.left  
        level += level
    return newtemp


n = int(input())

indexes = []

for _ in range(n):
    indexes.append(list(map(int, input().rstrip().split())))

queries_count = int(input())

queries = []

for _ in range(queries_count):
    queries_item = int(input())
    queries.append(queries_item)

que = [Node(1)]
root = que[0]
temp = root


for x in indexes:
    while root.data == -1:
        que.pop(0)
        root=que[0]

    root.left = Node(x[0])
    root.right = Node(x[1])
    que.append(root.left)
    que.append(root.right)
    que.pop(0)
    if len(que) != 0:
        root = que[0]


max_Depth = maxDepth(temp)
newtemp = copy.deepcopy(temp)

for x in queries:
    r = []
    newtemp = swap(x, newtemp)
    newtemp.rs()

    fptr.write(' '.join([''.join(map(str, x)) for x in r]))
    fptr.write('\n')

fptr.close()