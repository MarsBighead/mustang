CREATE USER 'semaphore'@'%' IDENTIFIED BY 'semaphore';
CREATE DATABASE IF NOT EXISTS semaphore;
GRANT ALL ON semaphore.* TO 'semaphore'@'%';
GRANT SELECT ON *.* TO 'semaphore'@'%';
FLUSH PRIVILEGES;
