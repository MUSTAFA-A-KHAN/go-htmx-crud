package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/agustfricke/go-htmx-crud/database"
	"github.com/agustfricke/go-htmx-crud/models"
	"github.com/labstack/echo/v4"
)

func GetTasks(c echo.Context) error {
	fmt.Println("tests")
	db := database.DB
	var tasks []models.Task

	if err := db.Find(&tasks).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Error getting tasks from database")
	}

	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Render error")
	}

	return c.Render(http.StatusOK, tmpl.Name(), tasks)
}

func CreateTask(c echo.Context) error {
	time.Sleep(2 * time.Second)

	name := c.FormValue("name")
	if name == "" {
		return c.String(http.StatusBadRequest, "Can't create task without a name")
	}

	db := database.DB
	task := models.Task{Name: name}

	if err := db.Create(&task).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Error creating task in database")
	}

	tmpl, err := template.ParseFiles("templates/item.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Render error")
	}

	return c.Render(http.StatusOK, tmpl.Name(), task)
}

func DeleteTask(c echo.Context) error {
	time.Sleep(2 * time.Second)

	ID := c.QueryParam("ID")
	if ID == "" {
		return c.String(http.StatusBadRequest, "ID not found")
	}

	db := database.DB
	var task models.Task

	if err := db.First(&task, ID).Error; err != nil {
		return c.String(http.StatusNotFound, "Task not found")
	}

	if err := db.Delete(&task).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Error deleting task from database")
	}

	return c.NoContent(http.StatusOK)
}

func FormEditTask(c echo.Context) error {
	name := c.QueryParam("name")
	ID := c.QueryParam("ID")
	if ID == "" || name == "" {
		return c.String(http.StatusBadRequest, "ID or Name not found")
	}

	data := struct{ ID, Name string }{ID: ID, Name: name}

	tmpl, err := template.ParseFiles("templates/edit.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Render error")
	}

	return c.Render(http.StatusOK, tmpl.Name(), data)
}

func EditTask(c echo.Context) error {
	time.Sleep(2 * time.Second)

	name := c.FormValue("name")
	ID := c.QueryParam("ID")
	if ID == "" || name == "" {
		return c.String(http.StatusBadRequest, "ID or Name not found")
	}

	db := database.DB
	var task models.Task

	if err := db.First(&task, ID).Error; err != nil {
		return c.String(http.StatusNotFound, "Task not found")
	}

	task.Name = name
	if err := db.Save(&task).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Error saving task in database")
	}

	tmpl, err := template.ParseFiles("templates/item.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Render error")
	}

	return c.Render(http.StatusOK, tmpl.Name(), task)
}
