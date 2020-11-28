package handlers

import (
	"encoding/json"
	"fmt"
	"goblog/data"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPosts(response http.ResponseWriter, request *http.Request) {
	posts, err := data.GetAllPosts()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(posts)
}

func GetPost(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	post, err := data.GetPost(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(post)
}

func CreateNewPost(response http.ResponseWriter, request *http.Request) {
	var newPost data.Post

	json.NewDecoder(request.Body).Decode(&newPost)
	_, err := data.InsertPost(newPost)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	response.WriteHeader(http.StatusCreated)
}

func UpdatePost(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	fmt.Println("id:", id)
}

func DeletePost(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	_, err = data.DeletePost(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	response.WriteHeader(http.StatusOK)
}

func GetPostsByUserID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	fmt.Println("id:", id)
}
