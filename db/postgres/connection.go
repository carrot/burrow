package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var database *sql.DB

func Open() {
	// Pulling environment vars
	databaseUrl := os.Getenv("POSTGRES_DATABASE_URL")

	// Opening + storing the connection
	database, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Pinging the database
	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	database.Close()
}

func Get() *sql.DB {
	return database
}
