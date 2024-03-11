package api

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTables(db *sql.DB) {
	sensorTableStmt := `
	CREATE TABLE IF NOT EXISTS sensor (id INTEGER PRIMARY KEY, sensor TEXT, NH3_ppm INTEGER, CO_ppm INTEGER, NO2_ppm INTEGER)`
	// userTableStmt := `CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, name TEXT, password TEXT)`

	command, err := db.Prepare(sensorTableStmt)
	if err != nil {
		log.Fatal(err.Error())
	}
	command.Exec()
}

func Select(db *sql.DB) {
	row, err := db.Query("SELECT * FROM sensor ORDER BY timestamp")
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer row.Close()
	for row.Next() {
		var id int
		var sensor string
		var NH3_ppm int
		var CO_ppm int
		var NO2_ppm int
		var timestamp time.Time
		row.Scan(&id, &sensor, &NH3_ppm, &CO_ppm, &NO2_ppm, &timestamp)
		log.Println("Sensor data: %v - %v - %v - %v - %v - %v", id, sensor, NH3_ppm, CO_ppm, NO2_ppm, timestamp)
	}
}
