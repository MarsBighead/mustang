#!/bin/bash

OUTPUT_DIR=$(cd $(dirname $0);pwd)
awk 'BEGIN{printf "序号\t名字\t课程\t分数\n"} {print}' $OUTPUT_DIR/marks.txt
while read app; do
    echo -e "upgrade $app \n\n"
    # app_timeout=${TIMEOUT_CONFIG[$app]:-2m}
    # read data from file with while loop    
done < "$OUTPUT_DIR/marks.txt"
