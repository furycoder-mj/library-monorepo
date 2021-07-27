package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Library struct {
	Id           int64
	User         int64
	AssignedBook int64
}

func init() {
	// register model
	orm.RegisterModel(new(Library))
}
func (l Library) String() string {
	return fmt.Sprintf("Book{Id:%v User: %v Assigned Book: %v}", l.Id, l.User, l.AssignedBook)
}
