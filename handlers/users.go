package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	fmt.Println("id:", id)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	fmt.Println("id:", id)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	fmt.Println("id:", id)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	fmt.Println("id:", id)
}
