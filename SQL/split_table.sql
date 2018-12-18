DROP FUNCTION if exists split_bigtable("s" timestamp);
CREATE OR REPLACE FUNCTION split_bigtable("s" timestamp)
  RETURNS void AS
$$
DECLARE
    "e" timestamp :="s" + interval '1 months';
    "sId" integer;
    "eId" integer;
    t_name varchar(40) :='bigtable_'||to_char("e", 'YYYYMMDD');
   
BEGIN
    EXECUTE 'SELECT min(id), max(id) FROM "Collection" WHERE time >= $1 AND time <= $2'
        INTO "sId","eId"
    USING "s", "e";
    raise notice 'start Id %', "sId";
    raise notice 'end Id %', "eId";
    raise notice 't_name %', t_name;
    EXECUTE 'CREATE table '||t_name||' as select * from "bigtable"';
    TRUNCATE "bigtable";
    EXECUTE 'insert into "bigtable" (select * from bigtable_id where "collectionId" >= $1 and "collectionId" <= $2 )'
    USING "sId", "eId";
end;  
$$ LANGUAGE plpgsql;