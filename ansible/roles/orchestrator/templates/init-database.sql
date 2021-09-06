CREATE USER 'orchestrator'@'%' IDENTIFIED BY 'orchestrator';
GRANT SUPER, PROCESS, REPLICATION SLAVE, RELOAD ON *.* TO 'orchestrator'@'%';
GRANT SELECT ON mysql.slave_master_info TO 'orchestrator'@'%';
CREATE DATABASE IF NOT EXISTS orchestrator;
GRANT ALL ON orchestrator.* TO 'orchestrator'@'%';
use orchestrator;
CREATE TABLE IF NOT EXISTS orchestrator.cluster(
  anchor TINYINT NOT NULL,
  cluster_name VARCHAR(128) CHARSET ascii NOT NULL DEFAULT ' ',
  cluster_domain VARCHAR(128) CHARSET ascii NOT NULL DEFAULT ' ',
  repl_user VARCHAR(128) CHARSET ascii NOT NULL DEFAULT ' ',
  repl_pass VARCHAR(128) CHARSET ascii NOT NULL DEFAULT ' ',
  PRIMARY KEY(anchor)
) ENGINE=InnoDB;
CREATE TABLE IF NOT EXISTS orchestrator.datacenter(
  anchor TINYINT NOT NULL,
  datacenter_name VARCHAR(128) CHARSET ascii NOT NULL DEFAULT ' ',
  PRIMARY KEY(anchor)
) ENGINE=InnoDB;
CREATE TABLE IF NOT EXISTS promotionrule(
  anchor TINYINT NOT NULL,
  promotion_rule enum('must','prefer','neutral','prefer_not','must_not') NOT NULL DEFAULT 'neutral',
  PRIMARY KEY(anchor)
) ENGINE=InnoDB;

