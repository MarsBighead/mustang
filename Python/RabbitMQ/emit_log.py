#!/usr/bin/python
import pika
import sys
connection = pika.BlockingConnection(pika.ConnectionParameters(
      host='localhost'))
channel = connection.channel()

channel.exchange_declare(exchange='logs',
                        type='fanout')

msg = ' '.join(sys.argv[1:]) or "info:Hello World!"

channel.basic_publish(exchange='logs',
                      routing_key='',
                      body=msg)

print " [x] Sent %r" % msg
connection.close()
