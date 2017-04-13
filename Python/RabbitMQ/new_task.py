#!/usr/bin/python
import pika
import sys 

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()

channel.queue_declare(queue='task', durable=True)

msg = ' '.join(sys.argv[1:]) or "Hello World!"
channel.basic_publish(exchange='',
                      routing_key='task',
                      body= msg,
                      properties=pika.BasicProperties(
                          delivery_mode=2,  #make message persistent
                      ))

print " [x] Sent %r" % msg
connection.close()
