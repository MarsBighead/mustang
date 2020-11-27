set @query_id=1;

SELECT STATE, SUM(DURATION) as Total_R,
ROUND (
    100*SUM(DURATION) /
    (SELECT SUM(DURATION)
    FROM INFORMATION_SCHEMA.PROFILING
    WHERE QUERY_ID=@query_id
    ), 2) as Pct_R,
    COUNT(*) AS Calls,
    SUM(DURATION)/COUNT(*) as "R/Call"
FROM INFORMATION_SCHEMA.PROFILING
WHERE QUERY_ID=@query_id
GROUP BY STATE
ORDER BY Total_R DESC
/*
mysql> show profiles;
+----------+------------+-------------------------------------------+
| Query_ID | Duration   | Query                                     |
+----------+------------+-------------------------------------------+
|        1 | 0.05346550 | select  * from nicer_but_slower_film_list |
+----------+------------+-------------------------------------------+
1 row in set, 1 warning (0.00 sec)

mysql> show profile for query 1;
+--------------------------------+----------+
| Status                         | Duration |
+--------------------------------+----------+
| starting                       | 0.000066 |
| Executing hook on transaction  | 0.000007 |
| starting                       | 0.000009 |
| checking permissions           | 0.000008 |
| Opening tables                 | 0.000124 |
| checking permissions           | 0.000099 |
| checking permissions           | 0.000009 |
| checking permissions           | 0.000005 |
| checking permissions           | 0.000003 |
| checking permissions           | 0.000005 |
| checking permissions           | 0.000194 |
| checking permissions           | 0.000011 |
| checking permissions           | 0.000003 |
| checking permissions           | 0.000003 |
| checking permissions           | 0.000003 |
| checking permissions           | 0.000027 |
| init                           | 0.000010 |
| System lock                    | 0.000013 |
| optimizing                     | 0.000005 |
| optimizing                     | 0.000013 |
| statistics                     | 0.000107 |
| preparing                      | 0.000022 |
| Creating tmp table             | 0.000064 |
| statistics                     | 0.000008 |
| preparing                      | 0.000012 |
| executing                      | 0.051559 |
| end                            | 0.000023 |
| query end                      | 0.000006 |
| waiting for handler commit     | 0.000241 |
| removing tmp table             | 0.000163 |
| waiting for handler commit     | 0.000014 |
| removing tmp table             | 0.000408 |
| waiting for handler commit     | 0.000030 |
| closing tables                 | 0.000058 |
| freeing items                  | 0.000043 |
| removing tmp table             | 0.000015 |
| freeing items                  | 0.000041 |
| cleaning up                    | 0.000037 |
+--------------------------------+----------+
38 rows in set, 1 warning (0.01 sec)

mysql> set @query_id=1;
Query OK, 0 rows affected (0.00 sec)

mysql> 
mysql> SELECT STATE, SUM(DURATION) as Total_R,
    -> ROUND (
    ->     100*SUM(DURATION) /
    ->     (SELECT SUM(DURATION)
    ->     FROM INFORMATION_SCHEMA.PROFILING
    ->     WHERE QUERY_ID=@query_id
    ->     ), 2) as Pct_R,
    ->     COUNT(*) AS Calls,
    ->     SUM(DURATION)/COUNT(*) as "R/Call"
    -> FROM INFORMATION_SCHEMA.PROFILING
    -> WHERE QUERY_ID=@query_id
    -> GROUP BY STATE
    -> ORDER BY Total_R DESC;
+--------------------------------+----------+-------+-------+--------------+
| STATE                          | Total_R  | Pct_R | Calls | R/Call       |
+--------------------------------+----------+-------+-------+--------------+
| executing                      | 0.051559 | 96.43 |     1 | 0.0515590000 |
| removing tmp table             | 0.000586 |  1.10 |     3 | 0.0001953333 |
| checking permissions           | 0.000370 |  0.69 |    12 | 0.0000308333 |
| waiting for handler commit     | 0.000285 |  0.53 |     3 | 0.0000950000 |
| Opening tables                 | 0.000124 |  0.23 |     1 | 0.0001240000 |
| statistics                     | 0.000115 |  0.22 |     2 | 0.0000575000 |
| freeing items                  | 0.000084 |  0.16 |     2 | 0.0000420000 |
| starting                       | 0.000075 |  0.14 |     2 | 0.0000375000 |
| Creating tmp table             | 0.000064 |  0.12 |     1 | 0.0000640000 |
| closing tables                 | 0.000058 |  0.11 |     1 | 0.0000580000 |
| cleaning up                    | 0.000037 |  0.07 |     1 | 0.0000370000 |
| preparing                      | 0.000034 |  0.06 |     2 | 0.0000170000 |
| end                            | 0.000023 |  0.04 |     1 | 0.0000230000 |
| optimizing                     | 0.000018 |  0.03 |     2 | 0.0000090000 |
| System lock                    | 0.000013 |  0.02 |     1 | 0.0000130000 |
| init                           | 0.000010 |  0.02 |     1 | 0.0000100000 |
| Executing hook on transaction  | 0.000007 |  0.01 |     1 | 0.0000070000 |
| query end                      | 0.000006 |  0.01 |     1 | 0.0000060000 |
+--------------------------------+----------+-------+-------+--------------+
18 rows in set, 19 warnings (0.00 sec)
*/
-- p85
show status where variable_name like "Handler%" or variable_name like "Created%"


CREATE TABLE enum_test (
    e enum('fish', 'apple', 'dog') NOT NULL
);

insert into enum_test(e) VALUES('fish'), ('dog'),('apple');