package models

import (
	"crudAPI/database"
	"database/sql"
	"log"
)

// Create an unexported global variable to hold the database connection pool.
var db *sql.DB = database.PostgresInstance()

// Create users table if not exists
func init() {
	query := "create table if not exists users3Te(id varchar primary key, name varchar, gender varchar, age integer);"
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalln("Table creation failed: ", err.Error())
	}
}
