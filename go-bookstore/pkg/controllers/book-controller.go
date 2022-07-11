package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"start/go-bookstore/pkg/model"
	"start/go-bookstore/pkg/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var CONTENT_TYPE = "content-type"
var APPLICATION_JSON = "application/json"

var NewBook model.Book

func finishResponse(w http.ResponseWriter, responseBody []byte) {
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := model.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	finishResponse(w, res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := model.GetBookById(id)
	res, _ := json.Marshal(bookDetails)
	finishResponse(w, res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &model.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	finishResponse(w, res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := model.DeleteBook(id)
	res, _ := json.Marshal(book)
	finishResponse(w, res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &model.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	booksDetails, db := model.GetBookById(id)
	if updateBook.Name != "" {
		booksDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		booksDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		booksDetails.Publication = updateBook.Publication
	}
	db.Save(&booksDetails)
	res, _ := json.Marshal(booksDetails)
	finishResponse(w, res)
}
