package main

import (
	"github.com/iannealer/go_playground/go-rest-example/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
