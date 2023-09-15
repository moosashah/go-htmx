package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/moosashah/go-htmx/pkg/database"
)

func Hello(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.Render(404, "index.html", "no name")
	}
	return c.Render(http.StatusOK, "index.html", name)
}

func CreateTodo(c echo.Context) error {
	todo := database.Todo{
		Content:   c.FormValue("content"),
		Completed: false,
		Id:        -1,
	}

	idStr := c.Param("id")
	if idStr != "" {
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "")
		}
		todo.Id = idInt
	}
	errors, err := todo.Save()
	if err != nil {
		return err
	}

	return errors
}
