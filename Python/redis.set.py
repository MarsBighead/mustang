#!/usr/bin/python
import redis
client =  redis.Redis(host='localhost', port=6379, db=0)
if client.ping():
    print "Connect to server successfully!"
else:
    print "Connect to server failed!"

# Build a set myset (sadd), 
# redis support add multiple, maybe python module limit it to one
client.sadd('myset','redis')
client.sadd('myset','hello')
client.sadd('myset','bar')

# Return all members of the set ``name``,smembers(self, name)
setVal = client.smembers('myset')
print "All members in the set:", setVal

# Get numbers of set member, scard(self,name)
setNum = client.scard('myset')
print "Get numbers of element in the set:",setNum

# Return the difference of sets specified by ``keys``, sdiff(self, keys *args)
client.sadd('myset1','redis')
#client.sadd('myset1','hello')
client.sadd('myset2','bar')
client.sadd('myset1','hi')
setDiff = client.sdiff(['myset','myset1','myset2'])
print "The difference of sets specified by keys:",setDiff

# Store the difference of sets specified by ``keys`` into a new
# set named ``dest``.  Returns the number of keys in the new set.
# sdiffstore(self, dest, keys, *args)
client.sdiffstore('setdiff',['myset','myset1'])
setDiff = client.smembers('setdiff')
print "The difference of sets store in the new set setDiff:", setDiff

# Return the intersection of sets specified by ``keys``, sinter(self, keys, *args)
setInter = client.sinter(['myset','myset1'])
print "The intersection of sets specified by keys:",setInter

# sinterstore(self, dest, keys, *args)
# Store the intersection of sets specified by ``keys`` into a new
# set named ``dest``.  Returns the number of keys in the new set.
client.sinterstore('setinter',['myset','myset1'])
setInter = client.smembers('setinter')
print "The intererence of sets store in the new set setInter:", setInter

# Judge if element belongs to te set(sismember)
judge =['redis','foo']
for ele in judge:
    boonVal = client.sismember('myset',ele)
    if boonVal:
        print "Element",ele," belongs to set myset"
    else:
        print "Element",ele," doesn't belong to set myset"

# smove(self, src, dst, value)
# Move ``value`` from set ``src`` to set ``dst`` atomically
client.sadd('myset','world')
client.smove('myset','myset1','world')
setNew = client.smembers('myset1')
print "The distination set:", setNew

# Remove and return a random member of set ``name``, spop(self, name)
spop = client.spop('myset1')
print "Get return random element:",spop, ", remaining elements:",client.smembers('myset1')

# Return a random member of set ``name``,  srandmember(self, name),
# >2.6 return numbers can be self define
srandom = client.srandmember('myset')
print "Return random element:",srandom 

# Remove ``value`` from set ``name``, srem(self, name, value)
rem = client.srem('myset','hello')
print "Get srem boolean:",rem 

# sunion(self, keys, *args)
# Return the union of sets specifiued by ``keys``
client.sadd('myset0','special')
setUnion = client.sunion(['myset','myset0','myset1'])
print "The union of sets specified by keys:",setUnion

# sunionstore(self, dest, keys, *args)
# Store the union of sets specified by ``keys`` into a new
# set named ``dest``.  Returns the number of keys in the new set.
client.sunionstore('setunion',['myset','myset1'])
setUnion = client.smembers('setunion')
print "The union of sets store in the new set setUnion:", setUnion

# no SSCAN

#Empty db
client.flushdb()
