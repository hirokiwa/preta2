package gormmodel

import (
	"time"
)


type Diary struct {
	Diaryid       int `gorm:"AUTO_INCREMENT"`
	Word     *string `sql:"Word"`
	Imageurl string `sql:"imageurl"`
	Userid   string `sql:"userid"`
	CreatedAt time.Time `sql:"createdAt" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time `sql:"updatedAT" sql:"DEFAULT:current_timestamp on update current_timestamp"`
}

type User struct {
	Userid   string `sql:"userid"`
	Name string `sql:"name"`
	// Diary Diary `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}



type Follow struct {
	Followee string `gorm:"primaryKey"`
	Follower string `gorm:"primaryKey"`
}
