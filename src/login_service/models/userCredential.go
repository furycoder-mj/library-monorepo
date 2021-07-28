package models

import "github.com/astaxie/beego/orm"

type UserCredential struct {
	Id       int64
	UserName string
	Password string
}

func init() {
	orm.RegisterModel(new(UserCredential))
}
