#!/bin/bash

INTERVAL=5
PRIFIX=${INTERVAL}-sec-status
RUNFILE=$(cd $(dirname $0); pwd)"/running.sh"
echo $RUNFILE
#Inject password with export MYSQL_PWD
mysql -h localhost -u root -e 'SHOW GLOBAL VARIABLES'>> mysql-varibles


while test -e $RUNFILE; do
    file=$(date +%F_%I)
    sleep=$(date +%s.%N | awk "{print $INTERVAL-(\$1 % $INTERVAL)}")
    #echo "$INTERVAL $sleep"
    sleep $sleep
    ts=$(date +"TS %s.%N %F %T")
    loadavg=$(uptime)
    echo "$ts $loadavg">>$PRIFIX-${file}-status
    mysql  -h localhost -u root -e 'SHOW GLOBAL Status'>>$PRIFIX-${file}-status
    #mysql  -h localhost -u root -e 'SHOW ENGINE INNODB Status\G'
    #mysql  -h localhost -u root -e 'SHOW FULL PROCESSLIST\G'
done
echo Exiting becacuse $RUNFILE does not exists.