package controller

import (
	"encoding/json"
	"github.com/iannealer/go_playground/go-rest-api/model"
	"net/http"
)

var Posts []model.Post

func init() {
	Posts = []model.Post{model.Post{Id: 1, Title: "Title 1", Text: "Text 1"}}
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := json.Marshal(Posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}

	post.Id = len(Posts) + 1
	posts = append(Posts, post)
	w.WriteHeader(http.StatusOK)

	result, err := json.Marshal(post)
	w.Write(result)
}
