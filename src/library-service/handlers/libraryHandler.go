package handlers

import (
	"fmt"
	"net/http"

	"epam.com/web-services/library-management/library-service/models"
	"epam.com/web-services/library-management/library-service/repository"
)

type userBookAssociation struct {
	userId int64
	bookId int64
}

// Issues book to user
func (uba userBookAssociation) IssueBookToUser(w http.ResponseWriter, r *http.Request) {
	recordId, err := repository.AddLibraryEntry(models.Library{User: uba.userId, AssignedBook: uba.bookId})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = w.Write([]byte(fmt.Sprint(recordId)))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Releases book from user
func (uba userBookAssociation) ReleaseBookFromUser(w http.ResponseWriter, r *http.Request) {
	deletedRows, err := repository.DeleteUserAndAssignedBookEntry(uba.userId, uba.bookId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = w.Write([]byte(fmt.Sprint(deletedRows)))
	if err != nil {
		fmt.Println(err)
		return
	}
}
