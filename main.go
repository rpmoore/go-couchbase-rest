package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage")
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Users Page")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Fprintf(w, "User page for user (%s)", id)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/users", UsersHandler)
	r.HandleFunc("/users/{id}", UserHandler)

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
