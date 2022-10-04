package dto

import "time"



type UserandDiary struct {
	Userid   string `sql:"userid"`
	Name string `sql:"name"`
	Diaryid       int `gorm:"AUTO_INCREMENT"`
	Word     string `sql:"Word"`
	Imageurl string `sql:"imageurl"`
	Englishword string `sql:"Englishword"`
	Happy     string `json:"Happy"`
	Angry     string  `json:"Angry"`
	Surprise  string  `json:"Surprise"`
	Sad       string  `json:"Sad"`
	Fear      string  `json:"Fear"`
	CreatedAt time.Time `gorm:"type:datetime(6)"`
	UpdatedAt time.Time `gorm:"type:datetime(6)"`
}