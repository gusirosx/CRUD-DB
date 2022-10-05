package infrastructure

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresClient struct {
	Client *sql.DB
}

func NewPostgresClient() *PostgresClient {
	return &PostgresClient{Client: PostgresInstance()}
}

// create the postgres database connection
func PostgresInstance() *sql.DB {
	connection, err := sql.Open("postgres", fmt.Sprintf(connString, user, databaseName, password, host, ssl))
	if err != nil {
		log.Fatalln("Unable to establish connection:", err.Error())
	}
	log.Println("Connected to postgres database!")
	return connection
}
