CREATE DATABASE IF NOT EXISTS orchestrator;
CREATE USER IF NOT EXISTS 'orchestrator'@'%' IDENTIFIED BY 'orchestrator';
GRANT ALL ON orchestrator.* TO 'orchestrator'@'%';
FLUSH PRIVILEGES;

