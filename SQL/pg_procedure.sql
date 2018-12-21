--
-- CREARE PROCEDURE feature will finally have it in version PostgreSQL 11.
-- https://www.postgresql.org/docs/11/sql-alterprocedure.html
-- https://www.cybertec-postgresql.com/en/tech-preview-postgresql-11-create-procedure/
-- root@c1557999abc3:~# psql -U hbu
-- psql (11.1 (Debian 11.1-1.pgdg90+1))
-- Type "help" for help.
-- 
-- hbu=# \l
--                              List of databases
--    Name    | Owner | Encoding |  Collate   |   Ctype    | Access privileges 
-- -----------+-------+----------+------------+------------+-------------------
--  hbu       | hbu   | UTF8     | en_US.utf8 | en_US.utf8 | 
--  postgres  | hbu   | UTF8     | en_US.utf8 | en_US.utf8 | 
--  template0 | hbu   | UTF8     | en_US.utf8 | en_US.utf8 | =c/hbu           +
--            |       |          |            |            | hbu=CTc/hbu
--  template1 | hbu   | UTF8     | en_US.utf8 | en_US.utf8 | =c/hbu           +
--            |       |          |            |            | hbu=CTc/hbu
-- (4 rows)
-- 
-- hbu=# CREATE PROCEDURE sum2(a integer, b integer)
-- hbu-# LANGUAGE SQL
-- hbu-# AS $$
-- hbu$# select a+b;
-- hbu$# $$;
-- CREATE PROCEDURE
-- hbu=# select (1,2);
--   row  
-- -------
--  (1,2)
-- (1 row)
--

CREATE PROCEDURE sum2(a integer, b integer)
LANGUAGE SQL
AS $$
select a+b;
$$;
