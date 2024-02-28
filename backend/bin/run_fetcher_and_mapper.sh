#!/bin/bash

# Navigate to the backend directory
if [ -d "../" ]; then
    echo "Navigating to backend directory..."
    cd "../" || exit
else
    echo "Error: Backend directory not found."
    exit 1
fi

# Check for command line option to drop the existing database
if [[ $1 == "--drop-db" ]]; then
    echo "Dropping existing 'airaccidentdata' database..."
    mysql -u "$MYSQL_USER" -p"$MYSQL_PASSWORD" -e "DROP DATABASE IF EXISTS airaccidentdata;"
fi

# Continue with the script to fetch and map data
echo "Fetching data..."
go run cmd/fetcher/main.go

echo "Mapping data..."
go run cmd/mapper/main.go

echo "Script execution completed."

# Usage example (from the project root directory):
# ./scripts/run_db_setup.sh --drop-db
