#!/usr/bin/python
import pika
import time
import uuid
from threading import Thread

class Productor_Thread(Thread):
    def __init__(self, threadname, starttime):
        Thread.__init__(self, name=threadname)
        self.threadname = threadname
        self.starttime = starttime
    def run(self):
        connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
        channel = connection.channel()

        channel.queue_declare(queue='hello')
        for i in range(10000):
            uuids = str(uuid.uuid4())
            boodystr = "%s::%s" %(self.threadname, uuids)
            channel.basic_publish(exchange='',
                      routing_key='hello',
                      body=boodystr)
        end = time.time()
        print end - self.starttime
        connection.close()

if __name__ == "__main__":
    start = time.time()
    for i in range(50):
        m =Productor_Thread(i, start)
        m.start()
