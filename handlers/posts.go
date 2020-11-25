package handlers

import (
	"encoding/json"
	"net/http"
)

type post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type allPosts []post

var posts = allPosts{
	{
		ID:    "1",
		Title: "First post",
		Body:  "This is the first post",
	},
}

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var newPost post

	json.NewDecoder(r.Body).Decode(&newPost)
	posts = append(posts, newPost)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newPost)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(posts)
}
