package handlers

import (
	"encoding/json"
	"goblog/data"
	"net/http"
)

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var newPost data.Post

	json.NewDecoder(r.Body).Decode(&newPost)
	data.InsertPost(newPost)
	w.WriteHeader(http.StatusCreated)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	data.GetAllPosts()
}
