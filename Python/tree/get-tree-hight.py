#!/bin/python3

# A Binary tree node 
class Node: 
      
    # Constructor to create a new node 
    def __init__(self, data): 
        self.data = data
        self.hight= 0
        self.left = None
        self.right = None

        #
# A function to do postorder tree traversal 
def printPreorder(root): 
  
    if root: 
        if root.data!=-1: print("hight=",root.hight, root.data) 
        # First recur on left child 
        printPreorder(root.left) 
  
        # the recur on right child 
        printPreorder(root.right) 

def generateTree(indexes):
    count=0
    root=Node(1)
    buf=[root]
    tree=root
    for left,right in indexes:
        while root.data == -1:
            buf.pop(0)
            root=buf[0]

        root.left = Node(left)
        root.left.hight=root.hight+1
        root.right = Node(right)
        root.right.hight=root.hight+1
        buf.append(root.left)
        buf.append(root.right)
        buf.pop(0)
        if len(buf) != 0:
            root = buf[0]
    for node in buf:
        print("data=",node.data)
    return tree
   
def swapNode(root, query, buf):
    
    if  root:
        if root.hight==query-1:
            buf.append(root)
        else:
            swapNode(root.left,query, buf)
            swapNode(root.right,query, buf)


# Complete the swapNodes function below.
#
def swapNodes(indexes, queries):
    #
    # Write your code here.
    #
    
    root=generateTree(indexes)
    print("\nPostorder traversal of binary tree is")
    printPreorder(root)
    buf=[]
    swapNode(root,queries[0],buf)
    for node in buf:
        node.left,node.right=node.right,node.left
    print("\nX binary tree is", len(buf))
    printPreorder(root)
    return "x"

if __name__ == '__main__':

    n = 11

    indexes = [
        [2,3],
        [4,-1],
        [5,-1],
        [6,-1],
        [7,8],
        [-1,9],
        [-1,-1],
        [10,11],
        [-1,-1],
        [-1,-1],
        [-1,-1],
    ]


    queries_count = 2

    queries = [2, 4]


    result = swapNodes(indexes, queries)
    print("x=",result)
    
