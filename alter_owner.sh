#!/bin/bash
for tbl in `psql -qAt -c"SELECT 
  c.relname as "Name"
FROM pg_catalog.pg_class c
     LEFT JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
WHERE c.relkind IN ('r','v','m','S','f','')
      AND n.nspname <> 'pg_catalog'
      AND n.nspname <> 'information_schema'
      AND n.nspname !~ '^pg_toast'
  AND pg_catalog.pg_table_is_visible(c.oid)" usgmtr` ;
do
    psql -U postgres -d usgmtr -c"alter table \"$tbl\" owner to usgmtr"
    #echo $tbl
done

