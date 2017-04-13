#!/usr/bin/python
import pika
import sys
connection = pika.BlockingConnection(pika.ConnectionParameters(
      host='localhost'))
channel = connection.channel()

channel.exchange_declare(exchange='direct_logs',
                        type='direct')

severity = sys.argv[1] if len(sys.argv)>1 else 'info'
msg = ' '.join(sys.argv[2:]) or "Hello World!"

channel.basic_publish(exchange='direct_logs',
                      routing_key=severity,
                      body=msg)

print " [x] Sent %r:%r" % (severity,msg)
connection.close()
