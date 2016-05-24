package main

import (
	"github.com/carrot/burrow/controllers"
	"github.com/labstack/echo"
)

func prepareRoutes(e *echo.Echo) {
	topicsController := new(controllers.TopicsController)
	e.Get("/topics", topicsController.Index)
	e.Get("/topics/:id", topicsController.Show)
	e.Post("/topics", topicsController.Create)
	e.Put("/topics/:id", topicsController.Update)
	e.Delete("/topics/:id", topicsController.Delete)
}
