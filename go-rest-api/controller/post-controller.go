package controller

import (
	"encoding/json"
	"github.com/iannealer/go_playground/go-rest-api/entity"
	"github.com/iannealer/go_playground/go-rest-api/errors"
	"github.com/iannealer/go_playground/go-rest-api/service"
	"net/http"
)

//var Posts []entity.Post

var (
	postService service.PostService = service.NewPostService()
)

type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	CreatePost(w http.ResponseWriter, r *http.Request)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{"Error getting the posts"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post entity.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{"Error unmarshalling the request"})
		return
	}

	validateErr := postService.Validate(&post)
	if validateErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{validateErr.Error()})
		return
	}

	//postService.Create(&post)
	result, createErr := postService.Create(&post)
	if createErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{createErr.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
