#!/usr/bin/python

# https://www.hackerrank.com/challenges/30-binary-search-trees
# Binary Search Tree(BSTs)
class Node:
    def __init__(self,data):
        self.right=self.left=None
        self.data = data
class Solution:
    def insert(self,root,data):
        if root==None:
            return Node(data)
        else:
            if data<=root.data:
                cur=self.insert(root.left,data)
                root.left=cur
            else:
                cur=self.insert(root.right,data)
                root.right=cur
        return root

    def getHeight(self,root):
        #Write your code here
        if root:     
            if self.getHeight(root.left)> self.getHeight(root.right):
                return self.getHeight(root.left)+1
            else: 
                return self.getHeight(root.right)+1
        else:
            return -1
            
T=int(raw_input())
myTree=Solution()
root=None
for i in range(T):
    data=int(raw_input())
    root=myTree.insert(root,data)
height=myTree.getHeight(root)
print height       
