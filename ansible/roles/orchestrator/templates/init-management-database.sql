CREATE DATABASE IF NOT EXISTS orchestrator;
CREATE USER 'orchestrator'@'%' IDENTIFIED BY 'orchestrator';
GRANT ALL ON orchestrator.* TO 'orchestrator'@'%';
FLUSH PRIVILEGES;

