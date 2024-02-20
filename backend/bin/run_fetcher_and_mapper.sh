#!/bin/bash

# Navigate to the backend directory where the .env file is located
cd "$(dirname "$0")"/../backend || exit

# Source .env file for MySQL credentials
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
else
    echo "Error: .env file not found."
    exit 1
fi

# Check for command line option to drop the existing database
if [[ $1 == "--drop-db" ]]; then
    echo "Dropping existing 'planecrashdata' database..."
    mysql -u "$MYSQL_USER" -p"$MYSQL_PASSWORD" -e "DROP DATABASE IF EXISTS planecrashdata;"
fi

# Continue with the script to fetch and map data
echo "Fetching data..."
go run cmd/fetcher/main.go

echo "Mapping data..."
go run cmd/mapper/main.go

echo "Script execution completed."

# Usage example (from the project root directory):
# ./scripts/run_db_setup.sh --drop-db
