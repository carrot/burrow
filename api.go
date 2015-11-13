package main

import (
	db "github.com/carrot/go-base-api/db/postgres"
	"github.com/carrot/go-base-api/environment"
	"github.com/carrot/go-base-api/request"
	"github.com/tylerb/graceful"
	"log"
	"os"
	"time"
)

func main() {
	// ---------------------------
	// Setting Active Environment
	// ---------------------------

	if len(os.Args) > 1 {
		err := environment.Set(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Running requires an environment argument")
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

	port := environment.GetEnvVar(environment.PORT)
	log.Println("Server started on :" + port)
	graceful.ListenAndServe(e.Server(":"+port), 5*time.Second) // Graceful shutdown
}
