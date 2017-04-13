#!/usr/bin/python
import redis
import time
## Connect local redis service
client =redis.Redis(host='127.0.0.1',port=6379,db=0)
print "Connection to server successfully!"
dicKeys = client.keys("*")
print dicKeys

### Redis hash command part Start ###
# hset: Set key to value with hash name,hset(self, name, key, value)
# hget: Return the value of ``key`` within the hash ``name``, hget(self, name, key)
client.hset('myhash','field1',"foo")
hashVal = client.hget('myhash','field1')
print "Get hash value:",hashVal

# Get none-value
hashVal = client.hget('myhash','field2')
print "None hash value:",hashVal

# hexists: Returns a boolean indicating if ``key`` exists within hash ``name``
keyList= ['field1','field2']
for key in keyList:
    hexists = client.hexists('myhash',key)
    if hexists :
        print "Exist in redis-hash key:",key
    else:
        print "Not exist in redis-hash key:",key

# hgetall: Return a Python dict of the hash's name/value pairs
client.hset('myhash','field2',"bar")
valDict = client.hgetall('myhash')
print "Get python-dict from redis-hash",valDict

# hincrby: Increment the value of ``key`` in hash ``name`` by ``amount``
# default increment is 1,
client.hset('myhash','field',20)
client.hincrby('myhash','field')
print "Get incrby value(Default):",client.hget('myhash','field')
client.hincrby('myhash','field',2)
print "Get incrby value(step: 2):",client.hget('myhash','field')
client.hincrby('myhash','field',-3)
print "Get incrby value(step: -3):",client.hget('myhash','field')

# no method hincrbyfloat

#hkeys: Return the list of keys within hash ``name``
kL = client.hkeys('myhash')
print "Get redis-hash key list",kL

#hlen: Return the number of elements in hash ``name``
lenHash =client.hlen('myhash')
print "All hash length:",lenHash

#hmget: Returns a list of values ordered identically to ``keys``
#hmget(self, name, keys), keys should be python list data structure
val =client.hmget('myhash',['field','field1','field2','field3','fieldx'])
print "Get all redis-hash value list:",val

#hmset:  Sets each key in the ``mapping`` dict to its corresponding value in the hash ``name``
hmDict={'field':'foo','field1':'bar'}
hmKeys=hmDict.keys()
client.hmset('hash',hmDict)
val = client.hmget('hash',hmKeys)
print "Get hmset value:",val

#hdel: Delete ``key`` from hash ``name``
client.hdel('hash','field')
print "Get delete result:",client.hget('hash','field')

#hvals:  Return the list of values within hash ``name``
val = client.hvals('myhash')
print "Get redis-hash values with HVALS",val

#hsetnx: Set ``key`` to ``value`` within hash ``name`` if ``key`` does not exist.
#      Returns 1 if HSETNX created a field, otherwise 0.
r=client.hsetnx('myhash','field',2)
print "Check hsetnx execute result:",r," Value:",client.hget('myhash','field')
r=client.hsetnx('myhash','field10',20)
print "Check hsetnx execute result:",r,"Value",client.hget('myhash','field10')

hashVal = client.hgetall('profile')
print hashVal
#Empty db
client.flushdb()
