package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iannealer/go_playground/go-rest-api/controller"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Running...")
	})
	router.HandleFunc("/posts", controller.GetPosts).Methods("GET")
	router.HandleFunc("/posts", controller.CreatePost).Methods("POST")

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
