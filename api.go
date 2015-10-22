package main

import (
	"github.com/carrot/go-base-api/controllers"
	"github.com/carrot/go-base-api/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tylerb/graceful"
	"log"
	"time"
)

func main() {
	// ---------
	// Database
	// ---------

	db.Open()
	defer db.Close()

	// ------------
	// Controllers
	// ------------

	topicsController := new(controllers.TopicsController)

	// ----------
	// Framework
	// ----------

	e := echo.New()

	// -----------
	// Middleware
	// -----------

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ----------
	// Endpoints
	// ----------

	e.Get("/topics", topicsController.Index)
	e.Get("/topics/:id", topicsController.Show)
	e.Post("/topics", topicsController.Create)
	e.Put("/topics/:id", topicsController.Update)
	e.Delete("/topics/:id", topicsController.Destroy)

	// ----
	// Run
	// ----

	log.Println("Server started on :5000")
	graceful.ListenAndServe(e.Server(":5000"), 5*time.Second) // Graceful shutdown
}
