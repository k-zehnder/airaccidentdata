#!/bin/sh

echo "Running CSV downloader..."
go run ./cmd/csvdownloader/main.go

echo "Running CSV-to-MySQL processor..."
go run ./cmd/csvtomysql/main.go
