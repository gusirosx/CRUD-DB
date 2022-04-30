package database

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

var connString = "user=%s dbname=%s password=%s host=%s sslmode=disable"
var host = os.Getenv("PG_HOST")
var databaseName = os.Getenv("PG_DB_PG")
var user = os.Getenv("PG_USER")
var password = os.Getenv("PG_PASS")

// create the postgres database connection
func PostgresInstance() *sql.DB {
	connection, err := sql.Open("postgres", fmt.Sprintf(connString, user, databaseName, password, host))
	if err != nil {
		log.Println("Unable to establish connection:", err.Error())
	}
	return connection
}
