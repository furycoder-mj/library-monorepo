package repository

import (
	"epam.com/web-services/library-management/users-service/models"
	"epam.com/web-services/library-management/users-service/ormHandler"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

var o orm.Ormer

func init() {
	o = *ormHandler.New()
}

// Query all users from `users` table
func GetUsers() ([]*models.User, int64, error) {
	var users []*models.User
	num, err := o.QueryTable("user").All(&users)
	if err != nil {
		return nil, num, err
	}
	return users, num, nil
}

// Query user with specific id from `users` table
func GetUserById(id int64) (*models.User, error) {
	user := models.User{Id: id}
	err := o.Read(&user)
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}

// Delete user with specific id from `users` table
func DeleteUser(id int64) (int64, error) {
	num, err := o.Delete(&models.User{Id: id})
	if err != nil {
		return num, err
	}
	return num, nil
}

// Add user to `users` table
func AddUser(user models.User) (int64, error) {
	id, err := o.Insert(&user)
	if err != nil {
		return id, err
	}
	return id, nil
}

// update user in `users` table
func UpdateUser(user models.User) (int64, error) {
	num, err := o.Update(&user)
	if err != nil {
		return num, err
	}
	return num, nil
}
