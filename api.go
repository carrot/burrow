package main

import (
	db "github.com/carrot/burrow/db/postgres"
	"github.com/carrot/burrow/environment"
	"github.com/carrot/burrow/middleware"
	"github.com/carrot/burrow/response"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	echo_middleware "github.com/labstack/echo/middleware"
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

	e := echo.New()

	// -----------
	// Middleware
	// -----------

	e.Use(echo_middleware.Logger())
	e.Use(middleware.Recover())

	// -------------------
	// HTTP Error Handler
	// -------------------

	e.SetHTTPErrorHandler(func(err error, context echo.Context) {
		httpError, ok := err.(*echo.HTTPError)
		if ok {
			response := response.New(context)
			response.SetResponse(httpError.Code, nil)
			response.Render()
		}
	})

	// -------
	// Routes
	// -------

	prepareRoutes(e)

	// ----
	// Run
	// ----

	port := environment.GetEnvVar(environment.PORT)
	std := standard.New(":" + port)
	std.SetHandler(e)
	log.Println("Server starting on :" + port)
	graceful.ListenAndServe(std.Server, 5*time.Second) // Graceful shutdown
}
