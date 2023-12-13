package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret"))
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", registerHandler).Methods("GET")
	r.HandleFunc("/register", registerSubmitHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("GET")
	r.HandleFunc("/login", loginSubmitHandler).Methods("POST")
	r.HandleFunc("/logut", logoutHandler).Methods("GET")


	// r.Handle("/admin", alice.New(usernameMiddleware, roleMiddleware("admin")).ThenFunc(adminHandler))
	r.Handle("/admin", alice.New(usernameMiddleware, roleMiddleware("admin")).ThenFunc(adminHandler))
	r.Handle("/user", alice.New(usernameMiddleware, roleMiddleware("user")).ThenFunc(userhandler))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		renderTemplate(w, "index.html", nil)
	}).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
