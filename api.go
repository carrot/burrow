package main

import (
	"github.com/carrot/go-base-api/controllers"
	db "github.com/carrot/go-base-api/db/postgres"
	"github.com/carrot/go-base-api/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	echo_middleware "github.com/labstack/echo/middleware"
	"github.com/tylerb/graceful"
	"log"
	"os"
	"time"
)

func buildEcho() (e *echo.Echo) {
	// ----------
	// Framework
	// ----------

	e = echo.New()

	// -----------
	// Middleware
	// -----------

	e.Use(echo_middleware.Logger())
	e.Use(middleware.Recover())

	// ------------
	// Controllers
	// ------------

	topicsController := new(controllers.TopicsController)

	// ----------
	// Endpoints
	// ----------

	e.Get("/topics", topicsController.Index)
	e.Get("/topics/:id", topicsController.Show)
	e.Post("/topics", topicsController.Create)
	e.Put("/topics/:id", topicsController.Update)
	e.Delete("/topics/:id", topicsController.Delete)

	return
}

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

	e := buildEcho()

	// ----
	// Run
	// ----

	port := os.Getenv("PORT")
	log.Println("Server started on :" + port)
	graceful.ListenAndServe(e.Server(":"+port), 5*time.Second) // Graceful shutdown
}
