package main

import (
	"fmt"
	"log"

	"github.com/agustfricke/go-htmx-crud/database"
	"github.com/agustfricke/go-htmx-crud/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.ConnectDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/public", "public")

	e.GET("/", handlers.GetTasks)
	e.POST("/add", handlers.CreateTask)
	e.DELETE("/delete", handlers.DeleteTask)
	e.GET("/edit/form", handlers.FormEditTask)
	e.PUT("/put", handlers.EditTask)

	fmt.Println("Running on port 8000")
	log.Fatal(e.Start(":8000"))
}
