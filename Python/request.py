#!/usr/bin/python
from urllib2 import Request,urlopen
import json

def request_test(url):
    req = Request(  
        url= url
    ) 
    resp = urlopen(req)
    data = resp.read()
    #print data
    d= "\"",data,"\""
    #print type(data)
    #print d 
    jdata= json.loads(data)
    #jdata= json.loads(data)
    #print type(jdata)
    s = json.dumps(jdata,indent=4)
    print type(s)
    print s 
    for value in jdata.values():
        #print value 
        pass 

    for key in jdata.keys():
        pass 
        #print key,':',jdata[key]

    for key, value in jdata.items():
        pass 
        #print key,':',value 
    #print jdata['rating']['max']

if __name__ == '__main__':
    url='https://api.douban.com/v2/book/2129650' 
    request_test(url)
  
