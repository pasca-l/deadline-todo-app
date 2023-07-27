package controllers

import (
	"net/http"
	"log"
	"html/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/templates/home.html")
	if err != nil {
		log.Fatalln(err)
	}

	t.Execute(w, "data")
}
