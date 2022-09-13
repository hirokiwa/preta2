package gormmodel

type Diary struct {
	ID       string  `sql:"id"`
	Word     string `sql:"word"`
	Imageurl string  `sql:"imageurl"`
	UserID     string  `sql:"userid"`
}