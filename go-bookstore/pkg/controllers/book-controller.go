package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iannealer/go_playground/go-bookstore/pkg/models"
	"github.com/iannealer/go_playground/go-bookstore/pkg/utils"
	"net/http"
	"strconv"
)

var Book models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}

	utils.ParseBody(r, book)

	createdBook := book.CreateBook()

	res, _ := json.Marshal(createdBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()

	res, err := json.Marshal(books)

	if err != nil {
		fmt.Println("Error while marshaling")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	id, err := strconv.Atoi(bookId)

	if err != nil {
		fmt.Println("Error while parsing")
		return
	}

	book, _ := models.GetBookById(int64(id))

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.Atoi(bookId)

	if err != nil {
		fmt.Println("Error while parsing")
		return
	}

	book, _ := models.GetBookById(int64(id))

	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	models.DB.Save(&book)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.Atoi(bookId)

	if err != nil {
		fmt.Println("Error while parsing")
		return
	}

	book := models.DeleteBook(int64(id))

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
