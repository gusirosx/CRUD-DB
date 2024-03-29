package database

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

var connString = "user=%s dbname=%s password=%s host=%s sslmode=%s"
var host = os.Getenv("PG_HOST")
var databaseName = os.Getenv("PG_DB_PG")
var user = os.Getenv("PG_USER")
var password = os.Getenv("PG_PASS")
var ssl = os.Getenv("DB_SSLMODE")

// Create a custom SqlClient type which wraps the redis.Client connection pool.
type SqlClient struct {
	*sql.DB
}

// create the postgres database connection
func PostgresInstance() *SqlClient {
	connection, err := sql.Open("postgres", fmt.Sprintf(connString, user, databaseName, password, host, ssl))
	if err != nil {
		log.Fatalln("Unable to establish connection:", err.Error())
	}
	log.Println("Connected to postgres database!")
	return &SqlClient{connection}
}
