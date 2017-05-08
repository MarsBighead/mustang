#!/usr/bin/python
class Person:
	def __init__(self, firstName, lastName, idNumber):
		self.firstName = firstName
		self.lastName = lastName
		self.idNumber = idNumber
	def printPerson(self):
		print "Name:", self.lastName + ",", self.firstName
		print "ID:", self.idNumber

class Student(Person):
    def __init__(self,firstName, lastName, idNum, scores ):
        Person.__init__(self,firstName, lastName, idNum)
        self.scores=scores
    def calculate(self):
      
        l=len(self.scores)
        sum = 0
        for s in self.scores:
            sum += s
        avg=sum/len(self.scores)
        if avg>=90 and avg <=100 :
            return 'O'
        elif avg<90 and avg >=80:
            return 'E'
        elif avg<80 and avg >=70:
            return 'A'
        elif avg<70 and avg >=55:
            return 'P'
        elif avg<55 and avg >=40:
            return 'D'
        elif avg<40:
            return 'T'




line = raw_input().split()
firstName = line[0]
lastName = line[1]
idNum = line[2]
numScores = int(raw_input()) # not needed for Python
scores = map(int, raw_input().split())
s = Student(firstName, lastName, idNum, scores)
s.printPerson()
print "Grade:", s.calculate()
