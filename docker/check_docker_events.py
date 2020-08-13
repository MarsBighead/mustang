#!/usr/bin/python3

import json
import docker
import sys

cli=docker.DockerClient(base_url='unix://var/run/docker.sock')
events=cli.events(since=sys.argv[1], until=sys.argv[2])
for e in events:
    print(e)

'''
[hbu@localhost docker]$ python3  check_docker_events.py 1595821000 1595821608
b'{"status":"pull","id":"runseb/flask:latest","Type":"image","Action":"pull","Actor":{"ID":"runseb/flask:latest","Attributes":{"name":"runseb/flask"}},"scope":"local","time":1595821525,"timeNano":1595821525326709684}\n'
b'{"status":"create","id":"a384d69f4ff4fc7a202032d461789da982c3d8d30611326f47e49fa38706ab86","from":"runseb/flask","Type":"container","Action":"create","Actor":{"ID":"a384d69f4ff4fc7a202032d461789da982c3d8d30611326f47e49fa38706ab86","Attributes":{"image":"runseb/flask","name":"priceless_napier"}},"scope":"local","time":1595821525,"timeNano":1595821525808970702}\n'
b'{"Type":"network","Action":"connect","Actor":{"ID":"16e6b9cae0928506de33ffc6ebd48e9bd4f3b5df87e274d21094c767465f15c9","Attributes":{"container":"a384d69f4ff4fc7a202032d461789da982c3d8d30611326f47e49fa38706ab86","name":"bridge","type":"bridge"}},"scope":"local","time":1595821526,"timeNano":1595821526414650910}\n'
b'{"status":"start","id":"a384d69f4ff4fc7a202032d461789da982c3d8d30611326f47e49fa38706ab86","from":"runseb/flask","Type":"container","Action":"start","Actor":{"ID":"a384d69f4ff4fc7a202032d461789da982c3d8d30611326f47e49fa38706ab86","Attributes":{"image":"runseb/flask","name":"priceless_napier"}},"scope":"local","time":1595821527,"timeNano":1595821527040308734}\n'
'''


