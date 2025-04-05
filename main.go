package main

import (
	"RestAPI/db"
	"RestAPI/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	db.InitDB()
	e := echo.New()

	e.GET("/messages", handlers.GetHandler)
	e.POST("/messages", handlers.PostHandler)
	e.PATCH("/messages/:id", handlers.PatchHandler)
	e.DELETE("/messages/:id", handlers.DeleteHandler)

	e.Start("Localhost:8080")
}
