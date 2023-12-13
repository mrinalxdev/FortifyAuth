package main

import (
	"fmt"
	"html/template"
	"net/http"
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
	username := r.FormValue("username")
	password := r.FormValue("password")

	if !isValidUsername(username) || !isValidPassword(password) {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	if _, exists := users[username]; exists{
		http.Error(w, "Username already taken", http.StatusConflict)
		return
	}

	roles := []string{"user"}

	if preProcessFunc := getPreProcessFunc(username, password); preProcessFunc != nil {
		preProcessFunc()
	}

	users[username] = User{Username: username, Password: password, Roles: roles}

}

func isValidUsername (username string) bool {
	return len(username) >= 3
}

func isValidPassword(password string) bool {
	return len(password) >= 6
}

func getPreProcessFunc(username, password string) func() {
	return func() {
		fmt.Printf("Pre-processing for user registration: %s\n", username)
	}
}

func getPostRegistrationFunc(username string) func() {
	return func() { 
		fmt.Printf("Post-registration logic for user: %s\n", username) 
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "login.html", nil)
}

func loginSubmitHandler(w http.ResponseWriter, r *http.Request){
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, found := users[username]
	if !found || user.Password != password {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	session, _ := store.Get(r, "session-name")
	session.Values["username"] = username
	session.Save(r, w)

	http.Redirect(w, r, "/user", http.StatusSeeOther)
}

func logoutHandler(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session-time")

	username := getUsernameFromContext(r.Context())
	if preLogoutFunc := getPreLogoutFunc(username); preLogoutFunc != nil {
		preLogoutFunc()
	}

	delete(session.Values, "username")
	session.Save(r, w)

	if postLogoutFunc := getPostLogoutFunc(username); postLogoutFunc != nil {
		postLogoutFunc()
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}


func getPostLogoutFunc(username string) func() {
	if username == "admin" {
		return func() {
			fmt.Printf("Executing admin post-logout logic")
		}
	}

	return nil
}


func getPreLogoutFunc(username string) func() {
	if username == "admin" {
		return func() {
			fmt.Printf("Executing admin post-logout logic")
		}
	}

	return nil
}

func adminHandler(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "admin.html", nil)
}

func userhandler(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "user.html", nil)
}