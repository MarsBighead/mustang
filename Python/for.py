#!/usr/bin/python
import logging
logger = logging.getLogger()
logger.setLevel(logging.DEBUG)
formatter = logging.Formatter('%(asctime)s - %(levelname)s - %(message)s')

ch = logging.StreamHandler()
ch.setLevel(logging.DEBUG)
ch.setFormatter(formatter)
logger.addHandler(ch)

words =['this','is','an','ex','parrot']
for word in words:
    logger.debug(word) 
for id  in range (1,13):
    print (id);
