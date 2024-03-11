package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, _:= sql.Open("sqlite3", "./ponderada.db")
	defer db.Close() // Defer Closing the database

	// Criando a tabla
	sqlStmt := `
  CREATE TABLE IF NOT EXISTS sensores
  (id INTEGER PRIMARY KEY, nomeSensor TEXT, valor INTEGER)
  `
	// Preparando o sqlstatement de forma segura
	command, err := db.Prepare(sqlStmt) 
	if err != nil {
		log.Fatal(err.Error())
	}

	// Executando o comando sql
	command.Exec()

	// Criando uma função para inserir usuários
	insertUser := func(db *sql.DB, nomeSensor string, valor int) {
		stmt := `INSERT INTO sensores(nomeSensor, valor) VALUES (?, ?)`
		statement, err := db.Prepare(stmt)
		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = statement.Exec(nomeSensor, valor)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	insertUser(db, "NH3_ppm", 360)
	insertUser(db, "CO_ppm", 680)

	displayUsers(db)
}

func displayUsers(db *sql.DB) {
	row, err := db.Query("SELECT * FROM sensores ORDER BY nomeSensor")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var name string
		var valor string
		row.Scan(&id, &name, &valor)
		log.Println("Sensores: ", id, " ", name, " ", valor )
	}
}