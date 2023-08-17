package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func h1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/views/index.html"))
	films := map[string][]Film{
		"Films": {
			{Title: "foo1", Director: "bar1"},
			{Title: "foo2", Director: "bar2"},
			{Title: "foo3", Director: "bar3"},
		},
	}
	tmpl.Execute(w, films)
}

func handleAddFilm(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	tmpl := template.Must(template.ParseFiles("public/views/index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
}

func main() {

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film", handleAddFilm)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
