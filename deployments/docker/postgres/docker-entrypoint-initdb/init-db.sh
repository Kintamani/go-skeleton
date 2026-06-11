#!/bin/bash
set -e

# Cek apakah environment variables sudah di-set
if [ -z "$POSTGRES_USER" ] || [ -z "$POSTGRES_PASSWORD" ] || [ -z "$POSTGRES_DB" ]; then
  echo "Error: Environment variables POSTGRES_USER, POSTGRES_PASSWORD, and POSTGRES_DB must be set."
  exit 1
fi

# Tunggu beberapa detik agar PostgreSQL siap
echo "Waiting for PostgreSQL to be ready..."
sleep 5

# Inisialisasi database dan user dengan environment variables dari Docker
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE USER $POSTGRES_USER WITH PASSWORD '$POSTGRES_PASSWORD';
    CREATE DATABASE $POSTGRES_DB WITH OWNER $POSTGRES_USER ENCODING 'UTF8' LC_COLLATE='en_US.utf8' LC_CTYPE='en_US.utf8';
    ALTER DATABASE $POSTGRES_DB SET timezone TO 'Asia/Jakarta';
    GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO $POSTGRES_USER;
EOSQL

echo "Database and user setup complete."
