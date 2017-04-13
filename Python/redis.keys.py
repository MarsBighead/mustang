#!/usr/bin/python
import redis
import time
## Connect local redis service
client =redis.Redis(host='127.0.0.1',port=6379,db=0)
print "Connection to server successfully!"
dicKeys = client.keys("*")
print dicKeys

### Redis Key command part Start ###

# Set key-vlaue and get key-value
client.set('w3ckey','redis')
val = client.get('w3ckey')
print "Get key-vlaue ", val

# Delete key w3ckey
client.delete('w3ckey')
val = client.get('w3ckey')
print "Get none key-vlaue: ", val

# No dump key-value
client.set('greeting','Hello, dumping world!')
val = client.get('greeting')
print "Get key-vlaue ", val

# Exists check
keyList = ['w3ckey','greeting']
for key in keyList:
#    print "Key name: ",key
    isKey =client.exists(key)
    if isKey :
        print "Have value mapping with key: ", key
    else:
        print "No value mapping with key: ",key
# Set Expire time for greeting
client.expire('greeting',2)
#saveTime = client.pttl('greeting')
#print "Remaining time: ",saveTime
time.sleep(2)
isExpire = client.get('greeting')
if not isExpire:
    print "Key expire "
else: 
    print "Key not expire"

# Key get special pattern, Redis command:keys
#set value
keyDic = {'w3c1':'redis', 'w3c2':'mysql', 'w3c3':'mongodb'}
for key in keyDic.keys():
    client.set(key,keyDic[key])
keyList = client.keys("w3c*")
print "Get keys: ",keyList

# Get random key from DB, Redis command:randomkey
randomKey =client.randomkey()
print "Get random key: ",randomKey

# Get key type, Redis command:type
keyType =client.type(randomKey)
print "Get key type: ",keyType


# Push value to list's head
client.lpush('w3ckey','redis')
client.lpush('w3ckey','mongodb')
client.lpush('w3ckey','mysql')
val = client.lrange('w3ckey',0,3)
print "Get key-vlaue list: ", val

hashVal = client.hgetall('profile')
print hashVal
#Empty db
client.flushdb()


