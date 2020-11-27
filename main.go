package main

import (
	"goblog/data"
	"goblog/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	defer data.CloseClientDB()

	// Routing
	router := mux.NewRouter().StrictSlash(true)
	v1Router := router.PathPrefix("/api/v1").Subrouter()
	v1Router.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	v1Router.HandleFunc("/posts", handlers.CreateNewPost).Methods("POST")

	v1Router.HandleFunc("/signin", handlers.Signin).Methods("POST")
	v1Router.HandleFunc("/home", handlers.Homepage).Methods("GET")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
