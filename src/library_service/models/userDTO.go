package models

import (
	"fmt"

	"time"
)

type UserDTO struct {
	Id              int64
	Name            string
	BirthDate       time.Time
	AssociatedBooks []int64
}

type UserToken struct {
	Id       int64
	UserName string
}

func (u UserToken) String() string {
	return fmt.Sprintf("Book{Id:%v Author: %s}", u.Id, u.UserName)
}

func (u UserDTO) String() string {
	return fmt.Sprintf("UserDTO {Id:%v Name: %s BirtDate: %v Associated books:%v}", u.Id, u.Name, u.BirthDate, u.AssociatedBooks)
}
