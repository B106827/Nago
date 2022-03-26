#!/bin/sh
# docker-compose mysql initial script

# 環境変数で DB_NAME が指定されればそちらを優先

DB_NAME=${DB_NAME:-nago}

# create database
mysql -uroot -proot < /tmp/ddl/create_database.ddl
# create tables
mysql -uroot -proot $DB_NAME < /tmp/ddl/create_tables.ddl
# insert init data
mysql -uroot -proot $DB_NAME < /tmp/ddl/insert_init_data.sql
