package models

import (
	"github.com/iannealer/go_playground/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectDB()
	DB = config.GetDB()
	DB.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	//DB.NewRecord(b)
	DB.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	DB.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := DB.Where("ID = ?", Id).Find(book)
	return &book, db
}

func DeleteBook(ID int64) Book {
	var book Book
	DB.Where("ID = ?", ID).Delete(&book)
	return book
}
