package main 

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/justinas/alice"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret"))

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", registerHandler).Methods("GET")


	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
