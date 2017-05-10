#!/usr/bin/python

class Node:
    def __init__(self,data):
        self.data = data
        self.next = None
class Solution:
    def display(self,head):
        current = head
        while current:
            print current.data,
            current = current.next	

    def insert(self,head,data): 
        p = Node(data)
        if head==None:
            head = p
        elif (head.next == None):
            head.next = Node(data)
        else: 
            self.insert(head.next, data)
        return head
            
                 




mylist= Solution()
T=int(input())
head=None
for i in range(T):
    data=int(input())
    head=mylist.insert(head,data)    
mylist.display(head);
