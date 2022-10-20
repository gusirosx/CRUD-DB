package config

import (
	"crudAPI/database"
	"database/sql"
	"log"
)

// Create an unexported global variable to hold the database connection pool.
var db *sql.DB = database.PostgresInstance()

// Create users table if not exists
func Init() {
	// Drop table users if exists
	var drop = "DROP TABLE IF EXISTS users"
	queryexec(drop)
	// Create table users if exists
	var create = "create table if not exists users(id varchar primary key, name varchar, gender varchar, age integer)"
	queryexec(create)
	// insert into table users
	var insert = "insert into users(id, name, gender, age) values " +
		"('86ea9d8d-43e1-41e5-8640-c477dbe79ca0','Gustavo R','Male',28)," +
		"('552b62fa-e677-44ce-a5ed-f8f3bcc065d6','Cleide G','Female',58)," +
		"('d88b8eef-f8a5-41f1-b554-6c44fbae2762','Jurandir F','Male',35)," +
		"('0cd5286f-a016-463e-b773-8b18ff43d8f9','Clobison H','Male',24);"
	queryexec(insert)
}

// execute a query statement
func queryexec(query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
