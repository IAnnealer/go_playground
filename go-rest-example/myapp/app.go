package myapp

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// NewHandler make a new app handler
func NewHandler() http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Go")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get UserInfo by /user/{id}")
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["id"])
}
