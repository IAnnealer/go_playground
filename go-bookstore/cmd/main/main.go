package main

import (
	"github.com/gorilla/mux"
	"github.com/iannealer/go_playground/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":/3000", r))
}
