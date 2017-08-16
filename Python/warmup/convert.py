#!/usr/bin/python
import csv

input_file='test.csv'
STANDARD_FIELDS = [#'Name',  'Status', 'Host','
				  'Memory Size (MB)', 'Reservation (MB)','Uptime',
				  'State','UUID',
				  #'IP Address', 'DNS Name', 'MoRef',
			     ]	
COMPUTED_FIELDS = ['UptimeDay','MemoryMB', 'ResMemMB']

print STANDARD_FIELDS
print COMPUTED_FIELDS
preserved_fields=STANDARD_FIELDS[:]
print type(preserved_fields)
print preserved_fields
header_indices = dict()
print type(header_indices)

with open(input_file, 'rU') as fh:
    csv_reader = csv.reader(fh,dialect='excel')
    #header=csv_reader.shift()
    for row in csv_reader:
        #print row
        for index, value in enumerate(row):
            print value
            #pass

