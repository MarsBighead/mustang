#!/bin/python3

'''
Trees: Is This a Binary Search Tree?

https://www.hackerrank.com/challenges/ctci-is-binary-search-tree/problem

The data value of every node in a node's left subtree is less than the data value of that node.
The data value of every node in a node's right subtree is greater than the data value of that node.
The data value of every node is distinct.
'''
class Node: 
    def __init__(self,data): 
        self.left = None
        self.right = None
        self.data = data 

def inOrderTraverse(root, data):
    if root:
        inOrderTraverse(root.left, data)
        data.append(root.data)
        inOrderTraverse(root.right, data)
    else:
        return

def checkBST(root):
    if root is None:
        return True
    data=[]
    inOrderTraverse(root,data)
    max=data[0]
    for v in data[1:]:
        if v>max:
            max=v
        else:
            print("No")
            return False
        #print(v)
    print("Yes")
    return True
    
# Driver code 
root = Node(4) 
root.left      = Node(2) 
root.right     = Node(6) 
root.left.left  = Node(1) 
root.left.right  = Node(3)
root.right.left  = Node(5) 
root.right.right  = Node(7)
checkBST(root)
