package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// env variables
var connString = "user=%s dbname=%s password=%s host=%s sslmode=%s"
var host = os.Getenv("PG_HOST")
var databaseName = os.Getenv("PG_DB_PG")
var user = os.Getenv("PG_USER")
var password = os.Getenv("PG_PASS")
var ssl = os.Getenv("DB_SSLMODE")

// create the postgres database connection
func PostgresInstance() *sql.DB {
	connection, err := sql.Open("postgres", fmt.Sprintf(connString, user, databaseName, password, host, ssl))
	if err != nil {
		log.Fatalln("Unable to establish connection:", err.Error())
	}
	log.Println("Connected to postgres database!")
	return connection
}
