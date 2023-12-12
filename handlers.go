package main

import (
	// "fmt"
	"net/http"
	"html/template"
	// "context"

	// "github.com/gorilla/mux"
	// "github.com/justinas/alice"
)

var templates *template.Template

func init(){
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}){
	err := templates.ExecuteTemplate(w, tmpl, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "register.html", nil)
}

func registerSubmitHandler(w http.ResponseWriter, r *http.Request){
	
}