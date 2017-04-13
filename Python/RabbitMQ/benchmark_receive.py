#!/usr/bin/python
import pika
import time
from threading import Thread

class Consumer_Thread(Thread):
    def __init__(self, threadname, starttime):
        Thread.__init__(self, name=threadname)
        self.threadname = threadname
        self.starttime = starttime
    def run(self):
        connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
        channel = connection.channel()

        channel.queue_declare(queue='hello')
        def callback(ch, method, properties, body):
            print " [x] Received %r" % body
        channel.basic_consume(callback,
                      queue ='hello',
                      no_ack = True)
        end = time.time()
        print end - self.starttime
        channel.start_consuming()
        connection.close()

if __name__ == "__main__":
    start = time.time()
    for i in range(10):
        m = Consumer_Thread(i, start)
        m.start()
