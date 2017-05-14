#!/usr/bin/python



class Calculator:
    def power(self,n,p):
        self.n=n
        self.p=p
        if n>=0 and p>=0:
            return int(self.n**self.p)
        else: 
            raise Exception("n and p should be non-negative")
    

myCalculator=Calculator()
T=int(raw_input())
for i in range(T):
    n,p = map(int, raw_input().split())
    try:
        ans=myCalculator.power(n,p)
        print ans
    except Exception,e:
        print e    
    
