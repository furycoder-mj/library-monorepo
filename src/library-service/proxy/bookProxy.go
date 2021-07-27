package proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"epam.com/web-services/library-management/library-service/client"

	"epam.com/web-services/library-management/library-service/config"
	"epam.com/web-services/library-management/library-service/models"
)

// Get All books from `books` table using `books-service`
func GetBooks() (uint, []models.BookDTO, error) {
	response, err := client.GetRequest(config.Config.BooksURI, "")
	if err != nil {
		return uint(response.StatusCode), nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return uint(response.StatusCode), nil, err
	}
	books := make([]models.BookDTO, 0)
	unmarshallErr := json.Unmarshal(body, &books)
	if unmarshallErr != nil {
		return uint(response.StatusCode), nil, unmarshallErr
	}
	return uint(response.StatusCode), books, nil
}

// Get book with specific id from `books` table using `books-service`
func GetBook(id int64) (uint, models.BookDTO, error) {
	response, err := client.GetRequest(config.Config.BooksURI, fmt.Sprint(id))
	if err != nil {
		return uint(response.StatusCode), models.BookDTO{}, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return uint(response.StatusCode), models.BookDTO{}, err
	}
	book := models.BookDTO{}
	unmarshallErr := json.Unmarshal(body, &book)
	if unmarshallErr != nil {
		return uint(response.StatusCode), models.BookDTO{}, unmarshallErr
	}
	return uint(response.StatusCode), book, nil
}

// Delete book with specific id from `books` table using `books-service`
func DeleteBook(id int64) (uint, string, error) {
	response, err := client.DeleteRequest(config.Config.BooksURI, fmt.Sprint(id))
	if err != nil {
		return uint(response.StatusCode), "0", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return uint(response.StatusCode), "0", err
	}
	return uint(response.StatusCode), string(body), nil
}

// Add book to `books` table using `books-service`
func AddBook(book models.BookDTO) (uint, string, error) {
	serializedBook, marshallErr := json.Marshal(book)
	if marshallErr != nil {
		return uint(http.StatusInternalServerError), "0", marshallErr
	}
	response, err := client.PostRequest(config.Config.BooksURI, "", bytes.NewReader(serializedBook))
	if err != nil {
		return uint(http.StatusInternalServerError), "0", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return uint(response.StatusCode), "0", err
	}
	return uint(response.StatusCode), string(body), nil
}

// Update book with specific id from `books` table using `books-service`
func UpdateBook(id int64, book models.BookDTO) (uint, string, error) {
	serializedBook, marshallErr := json.Marshal(book)
	if marshallErr != nil {
		return uint(http.StatusInternalServerError), "0", marshallErr
	}
	response, err := client.PutRequest(config.Config.BooksURI, fmt.Sprint(id), bytes.NewReader(serializedBook))
	if err != nil {
		return uint(response.StatusCode), "0", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return uint(response.StatusCode), "0", err
	}
	return uint(response.StatusCode), string(body), nil
}
