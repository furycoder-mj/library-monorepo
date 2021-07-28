package models

import (
	"fmt"
)

type BookDTO struct {
	Id     int64
	Author string
	Title  string
}

func (b BookDTO) String() string {
	return fmt.Sprintf("BookDTO{Id:%v Author: %s Title: %s}", b.Id, b.Author, b.Title)
}
