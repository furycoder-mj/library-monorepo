package repository

import (
	"epam.com/web-services/library-management/books-service/models"
	"epam.com/web-services/library-management/books-service/ormHandler"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

var o orm.Ormer

func init() {
	o = *ormHandler.New()
}

// get all books in `books` table
func GetBooks() ([]*models.Book, int64, error) {
	var books []*models.Book
	num, err := o.QueryTable("book").All(&books)
	if err != nil {
		return nil, num, err
	}
	return books, num, nil
}

// get book with specific id in `books` table
func GetBookById(id int64) (models.Book, error) {
	book := models.Book{Id: id}
	err := o.Read(&book)
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

// delete book in `books` table
func DeleteBook(id int64) (int64, error) {
	num, err := o.Delete(&models.Book{Id: id})
	if err != nil {
		return num, err
	}
	return num, nil
}

// add book in `books` table
func AddBook(book models.Book) (int64, error) {
	id, err := o.Insert(&book)
	if err != nil {
		return id, err
	}
	return id, nil
}

// update book in `books` table
func UpdateBook(book models.Book) (int64, error) {
	num, err := o.Update(&book)
	if err != nil {
		return num, err
	}
	return num, nil
}
