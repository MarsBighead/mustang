set sql_log_bin=0;
create user rpl_user@'%' identified WITH 'mysql_native_password' BY '!2Toger';
grant replication slave on *.* to rpl_user@'%' ;
GRANT BACKUP_ADMIN ON *.* TO rpl_user@'%';
flush privileges;

--For second nodes
GRANT BACKUP_ADMIN ON *.* TO rpl_user@'%';
FLUSH PRIVILEGES;
--
set sql_log_bin=1;
--install plugin group_replication soname 'group_replication.so';

-- correct master user to the right one
change master to master_user='rpl_user',master_password='!2Toger' 
for channel 'group_replication_recovery';

-- start with password for 8.0.21
SET GLOBAL group_replication_bootstrap_group=ON;
START GROUP_REPLICATION USER='rpl_user', PASSWORD='!2Toger';
SET GLOBAL group_replication_bootstrap_group=OFF;
