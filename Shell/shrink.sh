#!/bin/bash
 
days=10
cmd="find /data/assassin -mtime +$days -type f -name *.json"
num=`$cmd | wc -l`
# log_time=`date "+%Y-%m-%d %H:%M:%S"`
if [ $num -gt 0 ]; then
    echo "$(date +%Y-%m-%d) $(date +%H:%M:%S)   There are $num files to remove $days days ago."
    `find /data/assassin -mtime +$days -type f -name *.json -exec rm {} \;`
    if [ $? -eq 0 ]; then
       echo "$(date +%Y-%m-%d) $(date +%H:%M:%S)        Remove $num files successfully!"
    else
       echo "$(date +%Y-%m-%d) $(date +%H:%M:%S)        Remove $num files failed!"
    fi
else
    echo "$(date +%Y-%m-%d) $(date +%H:%M:%S)   There are no files to remove!"
fi
