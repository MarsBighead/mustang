CREATE USER 'repl'@'%' IDENTIFIED BY '!togerme';
GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%';
FLUSH PRIVILEGES;

INSTALL PLUGIN rpl_semi_sync_master SONAME 'semisync_master.so';
INSTALL PLUGIN rpl_semi_sync_slave SONAME 'semisync_slave.so';
SET GLOBAL rpl_semi_sync_master_enabled = 1;
SET GLOBAL rpl_semi_sync_slave_enabled=1;
SET GLOBAL rpl_semi_sync_master_timeout=5000000;
SET GLOBAL rpl_semi_sync_master_wait_for_slave_count=0;
SET GLOBAL rpl_semi_sync_master_wait_point=AFTER_SYNC;
