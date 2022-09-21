package gormmodel

import (
	"time"
)


type Diary struct {
	Diaryid       int `gorm:"AUTO_INCREMENT"`
	Word     *string `sql:"Word"`
	Imageurl string `sql:"imageurl"`
	Userid   string `sql:"userid"`
	CreatedAt time.Time `gorm:"type:datetime(6)"`
	UpdatedAt time.Time `gorm:"type:datetime(6)"`
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


type Emotion struct {
	Diaryid   string  `gorm:"primaryKey"`
	Happy     *string `json:"Happy"`
	Angry     string  `json:"Angry"`
	Surprise  string  `json:"Surprise"`
	Sad       string  `json:"Sad"`
	Fear      string  `json:"Fear"`
	CreatedAt time.Time  `gorm:"type:datetime(6)"`
	UpdatedAt time.Time  `gorm:"type:datetime(6)"`
}


type DiaryandEmotion struct {
	Diaryid       int `gorm:"AUTO_INCREMENT"`
	Word     *string `sql:"Word"`
	Imageurl string `sql:"imageurl"`
	Userid   string `sql:"userid"`
	Happy     *string `json:"Happy"`
	Angry     string  `json:"Angry"`
	Surprise  string  `json:"Surprise"`
	Sad       string  `json:"Sad"`
	Fear      string  `json:"Fear"`
	Name string `sql:"name"`
	CreatedAt time.Time `gorm:"type:datetime(6)"`
	UpdatedAt time.Time `gorm:"type:datetime(6)"`
}

type DiaryandUser struct {
	Diaryid       int `gorm:"AUTO_INCREMENT"`
	Word     *string `sql:"Word"`
	Imageurl string `sql:"imageurl"`
	Userid   string `sql:"userid"`
	Name string `sql:"name"`
	CreatedAt time.Time `gorm:"type:datetime(6)"`
	UpdatedAt time.Time `gorm:"type:datetime(6)"`
}