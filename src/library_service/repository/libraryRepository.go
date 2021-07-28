package repository

import (
	"fmt"

	"epam.com/web-services/library-management/library-service/models"
	"epam.com/web-services/library-management/library-service/ormHandler"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

var o orm.Ormer

func init() {
	o = *ormHandler.New()
}

// Add new entry in the `library` table
func AddLibraryEntry(library models.Library) (int64, error) {
	id, err := o.Insert(&library)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Get Assigned Books for a specific user from `library` table
func GetAssignedBooksForUser(userId int64) ([]models.Library, error) {
	var library []models.Library
	qs := o.QueryTable(&models.Library{})
	cond := orm.NewCondition().And("User", userId)
	_, err := qs.SetCond(cond).All(&library)
	if err != nil {
		return nil, err
	}
	fmt.Println(library)
	return library, nil
}

// Find if a book exist in `library` table
func BookExistInLibrary(bookId int64) bool {
	return o.QueryTable(&models.Library{}).Filter("AssignedBook", bookId).Exist()
}

// Delete entry in `library` table for user-book pair
func DeleteUserAndAssignedBookEntry(userId int64, bookId int64) (int64, error) {
	qs := o.QueryTable(&models.Library{})
	cond := orm.NewCondition().And("User", userId).And("AssignedBook", bookId)
	num, err := qs.SetCond(cond).Delete()
	if err != nil {
		return 0, err
	}
	return num, nil
}

// Delete entry in `library` table for specific user
func DeleteByUser(userId int64) (int64, error) {
	num, err := o.QueryTable(&models.Library{}).Filter("User", userId).Delete()
	if err != nil {
		return 0, err
	}
	return num, nil
}

// Delete entry in `library` table for specific book
func DeleteByAssignedBook(bookId int64) (int64, error) {
	num, err := o.QueryTable(&models.Library{}).Filter("AssignedBook", bookId).Delete()
	if err != nil {
		return 0, err
	}
	return num, nil
}
