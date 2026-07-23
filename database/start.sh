#!/bin/bash

set -e

sed -i "s/^#listen_addresses =.*/listen_addresses = '*'/" /etc/postgresql/*/main/postgresql.conf
sed -i "s/^listen_addresses =.*/listen_addresses = '*'/" /etc/postgresql/*/main/postgresql.conf

echo "host all all 0.0.0.0/0 md5" >> /etc/postgresql/*/main/pg_hba.conf

service postgresql start

sleep 5

su postgres -c "psql -c \"ALTER USER $DB_USER WITH PASSWORD '$DB_PASSWORD';\""

# --- First Database (Sightings) ---
su postgres -c "psql -tc \"SELECT 1 FROM pg_database WHERE datname='$DB_NAME'\" | grep -q 1 || psql -c \"CREATE DATABASE $DB_NAME\""
su postgres -c "psql $DB_NAME < /app/init.sql"

# --- Second Database (Users) ---
su postgres -c "psql -tc \"SELECT 1 FROM pg_database WHERE datname='$USERS_DB_NAME'\" | grep -q 1 || psql -c \"CREATE DATABASE $USERS_DB_NAME\""
su postgres -c "psql $USERS_DB_NAME < /app/users_init.sql"

echo "PostgreSQL is ready with both databases."

tail -f /dev/null