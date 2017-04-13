#!/usr/bin/python
import re


def calculate(num1,num2,op):
    if op == "+":
        return str(num1+num2)
    elif op == "-":
        return str(num1-num2)
    elif op == "*":
        return str(num1*num2)
    elif op == "/" and num2>0:
        return str(num1/num2)

def evalRPN(List):
    if len(List) == 3 and List[0].isdigit() and List[1].isdigit() and re.match('[\+\-\*\/]',List[2]): 
        return calculate(int(List[0]),int(List[1]),List[2])
    else:
       preList=[]
       for i in range(len(List)-3):
           if List[i].isdigit() and List[i+1].isdigit() and re.match('[\+\-\*\/]',List[i+2]): 
               preList.append(calculate(int(List[i]),int(List[i+1]),List[i+2]))
               preList.extend(List[i+3:])
               return evalRPN(preList)
           else:
               preList.append(List[i])



result = evalRPN(["1","2","+"])
print "Simple result: ",result
testList =[
   ["5","1","2","+","4","*","+","3","-"],
   ["4","13","5","/","+"],
   ["2","1","+","3","*"],
]

for testL in testList:
    tResult = int(evalRPN(testL))
    print "RPN result:\t",tResult
