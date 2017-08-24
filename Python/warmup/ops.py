#!/usr/bin/python
import csv
import glob
import sys

input_filename=sys.argv[1]
print "Input filename",input_filename

file_handler = open(input_filename, 'rU')
csv_dict = csv.DictReader(file_handler)
list_data = []
print type(csv_dict)
for row in csv_dict:
    print type(row)
    data = [row['Memory Size (MB)'],row['Reservation (MB)']]
    print data
