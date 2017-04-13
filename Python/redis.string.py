#!/usr/bin/python
import redis
import time

## Connect local redis service
client =redis.Redis(host='127.0.0.1',port=6379)
print "Connection to server successfully!"

client.set("tutorial-name","Redis tutorial")
print "Stored srting in redis: ",
print client.get("tutorial-name") 

# Get Key's children string(unable to achieve command GETRANGE)
key = "key1"
string ="This my test key"
client.set(key,string)
rangeString=client.substr(key,0,5)
print "The child string for key ",key," child vlaue list",rangeString

# Binding a new value to the key, return the old value(GETSET)
oldVal=client.getset("tutorial-name","Redis")
print "Get old value: ",oldVal
print "New value:",client.get("tutorial-name")

# getbit method not exists in the module(GETBIT) 
#client.set("tutorial-name","Redis tutorial")
#print client.getbit("tutorial-name",2)

# setex(self, name, value, time) command SETEX
# Set the value of key ``name`` to ``value`` 
# that expires in ``time`` seconds 
# ttl(self) command TTL
client.setex('mykey','mysql',10)
time.sleep(2)
timeLeft=client.ttl('mykey')
print "Get left live time",timeLeft
print "Get the value",client.get('mykey')

#setnx Set the value of key ``name`` to ``value`` if key doesn't exist. (SETNX)
client.setnx('mykey','PostgreSQL')
client.setnx('mykey1','MongoDB')
print "Get existed key's new value",client.get('mykey')
print "Get new built value",client.get('mykey1')

# no SETRANGE
# no STRLEN
# mset, Sets each key in the ``mapping`` dict to its corresponding value
client.mset({'key1':'Redis','key2':'MySQl'})
print "Get key1 mapping value",client.get('key1')
print "Get key2 mapping value",client.get('key2')


# msetnx: Sets each key in the ``mapping`` dict to its corresponding value if
#         none of the keys are already set
client.msetnx({'mykeyI':'Redis','mykeyII':'MySQl'})
print "Get mykeyI mapping value",client.get('mykeyI')
print "Get mykeyII mapping value",client.get('mykeyII')

# no PSETEX 
# incr method achieved both INCR and INCRBY command line in redis
# incr, iIncrements the value of ``key`` by ``amount``.
# If no key exists, the value will be initialized as ``amount``
# if not set the increment number, the default is 1. 
# We can set amount an negative value get DECR/DECRBY method.
client.set('page_num',20)
client.incr('page_num',-2)
print "The increased value",client.get('page_num')

# append(self,key,value), 
#       Appends the string ``value`` to the value at ``key``. If ``key``
#       doesn't already exist, create it with a value of ``value``.
#       Returns the new length of the value at ``key``.
client.append('mykey',' DB2')
print "Get appended string",client.get('mykey')

hashVal = client.hgetall('profile')
print hashVal
#Empty db
client.flushdb()

