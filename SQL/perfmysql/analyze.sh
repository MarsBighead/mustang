#!/bin/bash

awk '
  BEGIN {
      printf "#ts date time load QPS";
      fmt = " %.4f";
  }
  /^TS/ {
      ts      = substr($2, 1, index($2, ".")-1 );
      load    = NF-2
      diff    = ts - prev_ts;
      prev_ts = ts;
      printf "\n%s %s %s %s", ts, $3, $4, substr($load, 1, length($load)-1);
  }
  /Queries/ {
      printf fmt, ($2-Queries)/diff;
      Queries = $2;
  }
' "$@"
echo ""

:<<eof
gnuplot> plot "QPS-per-5-seconds-status" using 5 w lines title "QPS"
gnuplot> 
# 5 w means use the 5th column, "QPS" is the row number 
eof