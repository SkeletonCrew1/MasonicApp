#!/bin/bash

set -e

sed -i "s/^#listen_addresses =.*/listen_addresses = '*'/" /etc/postgresql/*/main/postgresql.conf
sed -i "s/^listen_addresses =.*/listen_addresses = '*'/" /etc/postgresql/*/main/postgresql.conf

echo "host all all 0.0.0.0/0 md5" >> /etc/postgresql/*/main/pg_hba.conf

service postgresql start

sleep 5

su postgres -c "psql -c \"ALTER USER $DB_USER WITH PASSWORD '$DB_PASSWORD';\""

su postgres -c "psql -tc \"SELECT 1 FROM pg_database WHERE datname='$DB_NAME'\" | grep -q 1 || psql -c \"CREATE DATABASE $DB_NAME\""

su postgres -c "psql $DB_NAME < /app/init.sql"

echo "PostgreSQL is ready."

tail -f /dev/null