SET SQL_LOG_BIN=0; 
CREATE USER IF NOT EXISTS rpl_user@'%' IDENTIFIED BY 'password'; 
GRANT REPLICATION SLAVE ON *.* TO rpl_user@'%'; 
GRANT BACKUP_ADMIN ON *.* TO rpl_user@'%'; 
FLUSH PRIVILEGES; 
SET SQL_LOG_BIN=1;
CHANGE MASTER TO MASTER_USER='rpl_user', MASTER_PASSWORD='password' FOR CHANNEL 'group_replication_recovery';
INSTALL PLUGIN group_replication SONAME 'group_replication.so';
{% if item.isPrimary is defined %}
SET GLOBAL group_replication_bootstrap_group=ON;
START GROUP_REPLICATION;
SET GLOBAL group_replication_bootstrap_group=OFF;
{% else %}
SELECT sleep(15);
START GROUP_REPLICATION;
{% endif %}
