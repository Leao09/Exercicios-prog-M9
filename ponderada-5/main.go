package main

import (
	"database/sql"


	_ "github.com/mattn/go-sqlite3"
	api "pond5/API"
)

func main() {
	database, _ := sql.Open("sqlite3", "DBMT/ponderada.db")
	defer database.Close()

	api.CreateTables(database)

	api.Subscriber(database)
	

	select {}
}