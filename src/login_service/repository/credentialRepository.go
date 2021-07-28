package repository

import (
	"epam.com/web-services/library-management/login-service/models"
	"epam.com/web-services/library-management/login-service/ormHandler"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var o orm.Ormer

func init() {
	o = *(ormHandler.New())
}

func GetCredential(userName string) (*models.UserCredential, error) {
	var userCredential models.UserCredential
	qs := o.QueryTable("user_credential")
	err := qs.Filter("user_name", userName).One(&userCredential)
	if err != nil {
		return nil, err
	}
	// fmt.Println(userCredential)
	return &userCredential, nil
}

// Add user to `users` table
func AddUserCredential(userCredential models.UserCredential) (int64, error) {
	id, err := o.Insert(&userCredential)
	if err != nil {
		return id, err
	}
	return id, nil
}

// // Delete user with specific id from `users` table
// func DeleteUserCredential(userName string) (int64, error) {
// 	num, err := o.Delete(&models.UserCredential{UserName: userName})
// 	if err != nil {
// 		return num, err
// 	}
// 	return num, nil
// }

// // update user in `users` table
// func UpdateUserCredential(user models.User) (int64, error) {
// 	num, err := o.Update(&user)
// 	if err != nil {
// 		return num, err
// 	}
// 	return num, nil
// }
