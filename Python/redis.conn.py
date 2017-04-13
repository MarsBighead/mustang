#!/usr/bin/python
import redis
client =  redis.Redis(host='localhost', port=6379, db=1)
if client.ping():
    print "Connect to server successfully!"
else:
    print "Connect to server failed!"
