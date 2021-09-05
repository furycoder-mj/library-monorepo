package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"epam.com/web-services/library-management/library-service/authorization"
	"epam.com/web-services/library-management/library-service/models"
	"epam.com/web-services/library-management/library-service/proxy"
	"epam.com/web-services/library-management/library-service/repository"
)

type BookId int64

// Get Books http FuncHandler
func getBooks(w http.ResponseWriter, r *http.Request) {
	statusCode, booksList, err := proxy.GetBooks()
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
	w.WriteHeader(int(statusCode))
}

// Get Book http FuncHandler
func (b BookId) getBook(w http.ResponseWriter, r *http.Request) {
	statusCode, book, err := proxy.GetBook(int64(b))
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(int(statusCode))
}

// Add Books http FuncHandler
func addBook(w http.ResponseWriter, r *http.Request) {
	_, authErr := authorization.ValidateToken(w, r)
	if authErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Invalid token:", authErr)
		return
	}
	var book models.BookDTO
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	statusCode, body, err := proxy.AddBook(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(body))
	w.WriteHeader(int(statusCode))
}

// Update Book http FuncHandler
func (b BookId) updateBook(w http.ResponseWriter, r *http.Request) {
	var book models.BookDTO
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	statusCode, body, err := proxy.UpdateBook(int64(b), book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(body))
	w.WriteHeader(int(statusCode))
}

// Delete Book http FuncHandler
func (b BookId) deleteBook(w http.ResponseWriter, r *http.Request) {
	_, authErr := authorization.ValidateToken(w, r)
	if authErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Invalid token:", authErr)
		return
	}
	_, err := repository.DeleteByAssignedBook(int64(b))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	statusCode, body, err := proxy.DeleteBook(int64(b))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(body))
	w.WriteHeader(int(statusCode))
}
