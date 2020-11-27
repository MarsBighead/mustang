#!/bin/sh
#prepare data
#Inject password with export MYSQL_PWD
sysbench --test=/usr/share/sysbench/oltp_read_write.lua --table_size=100000 --mysql-db=test \
--mysql-user=root prepare

sysbench --test=/usr/share/sysbench/oltp_read_write.lua --report-interval=10 --rand-type=uniform --mysql-user=root  --mysql-db=test \
 --table_size=100000 --threads=2 --time=100 --max-requests=0 run
:<<eof
 950s ] thds: 128 tps: 137.50 qps: 2774.71 (r/w/o: 1951.21/548.40/275.10) lat (ms,95%): 1453.01 err/s: 0.00 reconn/s: 0.00
[ 960s ] thds: 128 tps: 161.89 qps: 3234.73 (r/w/o: 2262.78/648.17/323.78) lat (ms,95%): 1376.60 err/s: 0.00 reconn/s: 0.00
[ 970s ] thds: 128 tps: 164.78 qps: 3259.96 (r/w/o: 2273.20/657.21/329.55) lat (ms,95%): 1258.08 err/s: 0.00 reconn/s: 0.00
[ 980s ] thds: 128 tps: 161.41 qps: 3254.05 (r/w/o: 2283.38/647.95/322.72) lat (ms,95%): 1304.21 err/s: 0.00 reconn/s: 0.00
[ 990s ] thds: 128 tps: 159.81 qps: 3192.78 (r/w/o: 2237.53/635.54/319.72) lat (ms,95%): 1453.01 err/s: 0.00 reconn/s: 0.00
[ 1000s ] thds: 128 tps: 166.29 qps: 3344.35 (r/w/o: 2342.09/669.67/332.58) lat (ms,95%): 1304.21 err/s: 0.00 reconn/s: 0.00
SQL statistics:
    queries performed:
        read:                            1901340
        write:                           543240
        other:                           271620
        total:                           2716200
    transactions:                        135810 (135.71 per sec.)
    queries:                             2716200 (2714.14 per sec.)
    ignored errors:                      0      (0.00 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          1000.7567s
    total number of events:              135810

Latency (ms):
         min:                                   19.02
         avg:                                  942.86
         max:                                 9891.50
         95th percentile:                     1938.16
         sum:                            128049438.17

Threads fairness:
    events (avg/stddev):           1061.0156/9.54
    execution time (avg/stddev):   1000.3862/0.19
eof

sysbench --test=/usr/share/sysbench/oltp_read_write.lua --report-interval=10 --rand-type=uniform --mysql-user=root \
 --mysql-db=test --time=100 --max-requests=0 cleanup