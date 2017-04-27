#!/bin/bash

days=10
cmd="find /data/assassin -mtime +$days -type f -name *.json"
num=`$cmd | wc -l`
# log_time=`date "+%Y-%m-%d %H:%M:%S"`
if [ $num -gt 0 ]; then
    echo "$(date +%Y-%m-%d) $(date +%H:%M:%S)   Move $num files to archive directory built $days days ago."
    `find /data/assassin/  -mtime +$days -type f -name *.json -exec mv {} /data/archive/assassin/ \;`
    if [ $? -eq 0 ]; then
       echo "$(date +%Y-%m-%d) $(date +%H:%M:%S)        Move $num files to archive directory successfully!"
       cd /data/archive/assassin/
       `tar zcf $(date +%Y-%m-%d).assassin.tgz *.json`
       if [ $? -eq 0 ]; then
           echo "$(date +%Y-%m-%d) $(date +%H:%M:%S)    Archive $num files successfully!"
           rm *.json
       else
           echo "$(date +%Y-%m-%d) $(date +%H:%M:%S)    Archive $num files failed!"
       fi
    else
       echo "$(date +%Y-%m-%d) $(date +%H:%M:%S)        Move $num files to archive directory failed!"
    fi
else
    echo "$(date +%Y-%m-%d) $(date +%H:%M:%S)   There are no files to remove!"
fi
