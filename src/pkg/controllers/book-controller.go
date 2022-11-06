package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AdilBikeev/simple-go-mysql/pkg/models"
	"github.com/AdilBikeev/simple-go-mysql/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

// @Summary     Возвращает все книги из БД
// @Description get books
// @ID get-books
// @Tags Book
// @Produce     json
// @Success     200 {array} models.Book
// @Failure     400 {object} utils.HTTPError
// @Failure     404 {object} utils.HTTPError
// @Failure     500 {object} utils.HTTPError
// @Router      /book [get]
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// @Summary     Возвращает книгу по Id
// @Description get book
// @ID get-book-by-id
// @Tags Book
// @Produce     json
// @Param       bookId    path      int     true  "Book ID"
// @Success     200 {object} models.Book
// @Failure     400 {object} utils.HTTPError
// @Failure     404 {object} utils.HTTPError
// @Failure     500 {object} utils.HTTPError
// @Router      /book/{bookId} [get]
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// @Summary     Создает новую книгу в БД
// @Description Create book
// @ID post-book
// @Tags Book
// @Accept      json
// @Produce     json
// @Param       message  body      models.Book  true  "Book Info"
// @Success     200 {object} models.Book
// @Failure     400 {object} utils.HTTPError
// @Failure     404 {object} utils.HTTPError
// @Failure     500 {object} utils.HTTPError
// @Router      /book [post]
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// @Summary     Создает новую книгу в БД
// @Description Create book
// @ID delete-book-by-id
// @Tags Book
// @Accept      json
// @Produce     json
// @Param       bookId    path      int     true  "Book ID"
// @Success     200 {object} models.Book
// @Failure     400 {object} utils.HTTPError
// @Failure     404 {object} utils.HTTPError
// @Failure     500 {object} utils.HTTPError
// @Router      /book/{bookId} [delete]
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deleteBook := models.DeleteBook(ID)
	res, _ := json.Marshal(deleteBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// @Summary     Обновляем информацию о книге
// @Description Update book
// @ID udapte-book
// @Tags Book
// @Accept      json
// @Produce     json
// @Param       message  body      models.Book  true  "Book Info"
// @Success     200 {object} models.Book
// @Failure     400 {object} utils.HTTPError
// @Failure     404 {object} utils.HTTPError
// @Failure     500 {object} utils.HTTPError
// @Router      /book [put]
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Print("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
