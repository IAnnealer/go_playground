package controller

import (
	"encoding/json"
	"github.com/iannealer/go_playground/go-rest-api/entity"
	"github.com/iannealer/go_playground/go-rest-api/repository"
	"math/rand"
	"net/http"
)

//var Posts []entity.Post

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error getting the posts"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post entity.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}

	post.ID = rand.Int63()
	repo.Save(&post)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(post)
}
