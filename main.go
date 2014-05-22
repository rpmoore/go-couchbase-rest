package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/couchbaselabs/go-couchbase"
	"github.com/gorilla/mux"
)

var userBucket *couchbase.Bucket

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage")
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Users Page")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user := User{}

	err := userBucket.Get(id, &user)
	if err != nil {
		if strings.Contains(err.Error(), "KEY_ENOENT") {
			fmt.Printf("The entry %s does not exist\n", id)
			http.NotFound(w, r)
			return
		}
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(user)
}

func main() {

	connection, err := couchbase.Connect("http://localhost:8091")
	if err != nil {
		log.Fatalf("Failed to connect to couchbase (%s)\n", err)
	}

	pool, err := connection.GetPool("default")
	if err != nil {
		log.Fatalf("Failed to get pool from couchbase (%s)\n", err)
	}

	userBucket, err = pool.GetBucket("default")
	if err != nil {
		log.Fatalf("Failed to get bucket from couchbase (%s)\n", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/users", UsersHandler)
	r.HandleFunc("/users/{id}", UserHandler)

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
