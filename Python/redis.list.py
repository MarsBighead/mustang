#!/usr/bin/python
import redis
## Connect local redis service
client =redis.Redis(host='127.0.0.1',port=6379,db=0)
print "Connection to server successfully! \n"
## Stored data in list (lpush)
listDict = {'tutorial-list':['Redis','Mangodb','MySQL','PostgreSQL']}
for key in listDict.keys():
    print "Key ",key
    valueList=listDict[key]
    for val in valueList:
        client.lpush(key,val)

#client.lpush("tutorial-list","Redis")
#client.lpush("tutorial-list","Mangodb")
#client.lpush("tutorial-list","MySQL")

# Get data from Redis (lrange) 
arrayList = client.lrange("tutorial-list",0,5)
print 'Stored srting in redis: ',
print arrayList
# Get element from list with index(lindex)
for i in xrange(4):
    ele = client.lindex("tutorial-list",i)
    print "The element with index number",i,"is",ele
# Get list length (llen)
l = client.llen("tutorial-list")
print "The list tutorial-list length is",l

# Get head element from tutorial-list(blpop)
for i in xrange(2):
    ele = client.blpop("tutorial-list",10)
    print "The",(i+1),"element from left ",ele


# Get tail element from tutorial-list(brpop)
for i in xrange(2):
    ele = client.brpop("tutorial-list",10)
    print "The",(i+1),"element from right ",ele

hashVal = client.hgetall('profile')
print hashVal

#Empty db
client.flushdb()
