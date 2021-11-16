package server

import (
	"database/sql"
	"flag"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func dbName() string {
	env := os.Getenv("CLAS_SENSOR_ENV")
	workdir := os.Getenv("CLAS_SENSOR_WORKDIR")

	// Check if we are running in unittest
	if flag.Lookup("test.v") != nil {
		return "file::memory:?cache=shared"
	}

	if workdir == "" {
		workdir = "."
	}

	if env != "prod" {
		return workdir + "/" + "clas_sensor_test.db"
	}
	return workdir + "/" + "clas_sensor.db"
}

func Connect() *sql.DB {
	name := dbName()
	log.Printf("Connecting to DB %s\n", name)
	db, err := sql.Open("sqlite3", name)

	if err != nil {
		log.Printf("Can not connect to DB: %s\n%v\n", dbName(), err)
		return nil
	}
	return db
}

func InitDb(db *sql.DB) {
	query := "CREATE TABLE IF NOT EXISTS temperatures (ID INTEGER PRIMARY KEY NOT NULL, Timestamp TEXT, OutdoorTemp REAL, IndoorTemp REAL)"
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Error when executing\n%s\n%v\n", query, err)
	}

}

func insert(db *sql.DB, td TemperatureData) {
	query := "INSERT INTO temperatures (Timestamp, OutdoorTemp, IndoorTemp) VALUES (?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("Error during prepare query\n%s\n", query)
		return
	}

	_, err = stmt.Exec(td.Timestamp, td.OutdoorTemp, td.IndoorTemp)

	if err != nil {
		log.Printf("Could not executre statememt %s\n", query)
	}
}

func query(db *sql.DB, query string) *sql.Rows {
	stmt, _ := db.Prepare(query)
	rows, _ := stmt.Query()
	return rows
}
