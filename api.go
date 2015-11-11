package main

import (
	db "github.com/carrot/go-base-api/db/postgres"
	"github.com/carrot/go-base-api/request"
	"github.com/joho/godotenv"
	"github.com/tylerb/graceful"
	"log"
	"os"
	"time"
)

func main() {
	// ----------------------
	// Environment Variables
	// ----------------------

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// ---------
	// Database
	// ---------

	db.Open()
	defer db.Close()

	// -----
	// Echo
	// -----

	e := request.BuildEcho()

	// ----
	// Run
	// ----

	port := os.Getenv("PORT")
	log.Println("Server started on :" + port)
	graceful.ListenAndServe(e.Server(":" + port), 5 * time.Second) // Graceful shutdown
}
