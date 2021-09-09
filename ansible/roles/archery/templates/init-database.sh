#!/bin/bash
CREATE USER IF NOT EXISTS archery@'%' IDENTIFIED BY 'Archery!23';
CREATE DATABASE `archery` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
GRANT ALL ON archery.* to  archery@'%';
FLUSH PRIVILEGES;
