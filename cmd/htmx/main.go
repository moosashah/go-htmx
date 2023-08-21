package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/moosashah/go-htmx/pkg/database"
)

type Film struct {
	Title    string
	Director string
}

func h1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/views/index.html"))

	rows, _ := database.Db.Query("SELECT * FROM movies")
	defer rows.Close()

	data := []Film{}
	for rows.Next() {
		i := Film{}
		rows.Scan(&i.Title, &i.Director)
		data = append(data, i)
	}

	tmpl.Execute(w, data)
}

func handleAddFilm(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	database.Db.Exec("INSERT INTO movies VALUES(?,?);", title, director)

	tmpl := template.Must(template.ParseFiles("public/views/index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
}

func main() {
	url := "movies.db"
	database.InitDB(url)

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film", handleAddFilm)

	port := ":8000"

	fmt.Printf("Running server on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
