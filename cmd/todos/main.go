package main

import (
	"html/template"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/moosashah/go-htmx/pkg/database"
	"github.com/moosashah/go-htmx/pkg/handlers"
)

func main() {
	e := echo.New()
	url := "tmp/todos.db"
	database.InitTodosDB(url)

	tmpl, err := template.ParseGlob("./public/views/todos/*.html")

	if err != nil {
		log.Fatalf("could not initialize templates: %+v", err)
	}

	e.Renderer = handlers.NewTemplateRenderer(tmpl)

	e.GET("/", handlers.Hello)

	port := ":1323"
	e.Logger.Fatal(e.Start(port))
}
