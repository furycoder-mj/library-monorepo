package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"epam.com/web-services/library-management/books-service/models"
	"epam.com/web-services/library-management/books-service/repository"
)

var booksPath = os.Getenv("BasePath")

func handleBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		booksList, _, err := repository.GetBooks()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(booksList)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusFound)
	case http.MethodPost:
		var book models.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		addedBookId, err := repository.AddBook(book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte(fmt.Sprint(addedBookId)))
		w.WriteHeader(http.StatusCreated)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", booksPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bookID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		book, err := repository.GetBookById(int64(bookID))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusFound)
	case http.MethodPut:
		var book models.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if book.Id != int64(bookID) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		updateCount, err := repository.UpdateBook(book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte(fmt.Sprint(updateCount)))
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		delCount, err := repository.DeleteBook(int64(bookID))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprint(delCount)))
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func badRequest(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// SetupRoutes :
func SetupRoutes() {
	booksHandler := http.HandlerFunc(handleBooks)
	bookHandler := http.HandlerFunc(handleBook)
	http.Handle(booksPath, booksHandler)
	http.Handle(fmt.Sprintf("%s/", booksPath), bookHandler)
}
