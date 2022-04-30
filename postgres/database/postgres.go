package database

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

func PostgresInstance() *sql.DB {
	password := os.Getenv("PG_PASS")
	user := os.Getenv("PG_USER")
	dbName := os.Getenv("PG_DB_PG")
	host := os.Getenv("PG_HOST")
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", user, dbName, password, host)
	connection, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Unable to establish connection:", err.Error())
	}
	return connection
}

// func DbConect() *sql.DB {
// 	conexao := "user=bob dbname=Store password=admin host=localhost sslmode=disable"
// 	db, err := sql.Open("postgres", conexao)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }
