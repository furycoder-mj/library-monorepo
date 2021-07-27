package models

import (
	"fmt"

	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id        int64
	Name      string
	BirthDate time.Time
}

func init() {
	orm.RegisterModel(new(User))
}

func (b User) String() string {
	return fmt.Sprintf("Book{Id:%v Author: %s Title: %v}", b.Id, b.Name, b.BirthDate)
}
