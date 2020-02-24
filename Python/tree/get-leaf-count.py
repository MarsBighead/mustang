#!/bin/python3

'''
Python program to count leaf nodes in Binary Tree 

https://www.geeksforgeeks.org/write-a-c-program-to-get-count-of-leaf-nodes-in-a-binary-tree/

'''
# A Binary tree node 
class Node: 
      
    # Constructor to create a new node 
    def __init__(self, data): 
        self.data = data  
        self.left = None
        self.right = None
  
# Function to get the count of leaf nodes in binary tree 
def getLeafCount(node): 
    if node is None: 
        return 0 
    if(node.left is None and node.right is None): 
        return 1 
    else: 
        return getLeafCount(node.left) + getLeafCount(node.right) 
  
if __name__ == '__main__':
    # Driver program to test above function 
    root = Node(1) 
    root.left = Node(2) 
    root.right = Node(3) 
    root.left.left = Node(4) 
    root.left.right = Node(5) 
    
    print ("Leaf count of the tree is %d" %(getLeafCount(root)))
    print("{}".format(root.left.left.data))
    #This code is contributed by Nikhil Kumar Singh(nickzuck_007) 