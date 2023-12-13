package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	// "github.com/justinas/alice"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret"))
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", registerHandler).Methods("GET")


	// r.Handle("/admin", alice.New(usernameMiddleware, roleMiddleware("admin")).ThenFunc(adminHandler))


	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
