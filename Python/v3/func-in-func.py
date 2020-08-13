#!/usr/bin/python3

def hi(name="Paul"):
    def greet():
        return "now you are in the greet() "+name+" function"
 
    def welcome():
        return "now you are in the welcome() "+name+" function"
 
    if name == "Paul":
        return greet
    else:
        return welcome
 
a = hi()
print(a)
#outputs: <function greet at 0x7f2143c01500>
 
 
print(a())
print("     name=:",hi(name="Paul")())
print("none name=:",hi("Paul")())