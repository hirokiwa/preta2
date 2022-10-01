package dto

import "time"

type Diary struct {
	Diaryid       int `gorm:"AUTO_INCREMENT"`
	Word     string `sql:"Word"`
	Imageurl string `sql:"imageurl"`
	Userid   string `sql:"userid"`
	CreatedAt time.Time `gorm:"type:datetime(6)"`
	UpdatedAt time.Time `gorm:"type:datetime(6)"`
}