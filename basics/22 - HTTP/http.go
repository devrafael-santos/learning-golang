package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(templates.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := user{"Rafael", "rafael@gmail.com"}

		templates.ExecuteTemplate(w, "home.html", u)
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Users page"))
	})

	fmt.Println("Port: 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
