package gormmodel

import (
	"time"
)

type autoincrement struct {
	ID string
}

type Diary struct {
	ID       int `gorm:"AUTO_INCREMENT"`
	Word     *string `sql:"Word"`
	Imageurl string `sql:"imageurl"`
	Userid   string `sql:"userid"`
	CreatedAt time.Time `sql:"createdAt" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time `sql:"updatedAT" sql:"DEFAULT:current_timestamp on update current_timestamp"`
}

type User struct {
	ID   string `sql:"id"`
	Name string `sql:"name"`
}
