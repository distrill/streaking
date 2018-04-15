#! /bin/bash

mysql -u root -p < init.sql
mysql -u streaking -pstreaking < migrate.sql
mysql -u streaking -pstreaking < seed.sql