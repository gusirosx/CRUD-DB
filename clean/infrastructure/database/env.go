package database

import "os"

// env variables
var connString = "user=%s dbname=%s password=%s host=%s sslmode=%s"
var host = os.Getenv("PG_HOST")
var databaseName = os.Getenv("PG_DB_PG")
var user = os.Getenv("PG_USER")
var password = os.Getenv("PG_PASS")
var ssl = os.Getenv("DB_SSLMODE")
