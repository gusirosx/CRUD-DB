package infrastructure

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/lib/pq"
)

// create the postgres database connection
func PostgresInstance() *sql.DB {
	connection, err := sql.Open("postgres", fmt.Sprintf(connString, user, databaseName, password, host, ssl))
	if err != nil {
		log.Fatalln("Unable to establish connection:", err.Error())
	}
	log.Println("Connected to postgres database!")
	return connection
}

// type PostgresClient struct {
// 	Client *sql.DB
// }

// func NewPostgresClient() *PostgresClient {
// 	return &PostgresClient{Client: PostgresInstance()}
// }

// type SQLRow struct {
// 	Rows *sql.Rows
// }

// func (r SQLRow) Scan(dest ...interface{}) {
// 	r.Rows.Scan(dest...)
// }

// func (r SQLRow) Next() bool {
// 	return r.Rows.Next()
// }
