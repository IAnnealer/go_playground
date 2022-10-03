package main

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iannealer/go_playground/go-rest-api/controller"
	"google.golang.org/api/option"
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

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("./go-rest-api-e0ab6 ee10ffee1e.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
