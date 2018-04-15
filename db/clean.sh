#! /bin/bash

mysql -u streaking -pstreaking < migrate.sql
mysql -u streaking -pstreaking < seed.sql