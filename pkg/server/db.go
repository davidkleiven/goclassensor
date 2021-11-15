package server

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func dbName() string {
	env := os.Getenv("CLAS_SENSOR_ENV")
	workdir := os.Getenv("CLAS_SENSOR_WORKDIR")

	if workdir == "" {
		workdir = "."
	}

	if env != "prod" {
		return workdir + "/" + "clas_sensor_test.db"
	}
	return workdir + "/" + "clas_sensor.db"
}

func connect() *sql.DB {
	log.Printf("Connecting to DB")
	db, err := sql.Open("sqlite3", dbName())

	if err != nil {
		log.Printf("Can not connect to DB: %s\n%v\n", dbName(), err)
		return nil
	}
	log.Printf("Connected to DB")
	return db
}

func initDb() {
	db := connect()
	if db == nil {
		return
	}

	query := "CREATE TABLE IF NOT EXISTS temperatures (Timestamp TEXT, DataSource TEXT, OutdoorTemp REAL, IndoorTemp REAL)"
	_, err := db.Query(query)

	if err != nil {
		log.Printf("Error when executing\n%s\n%v\n", query, err)
	}
	db.Close()
}
