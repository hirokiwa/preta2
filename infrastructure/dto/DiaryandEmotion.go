package dto


import (
	"time"
)




type DiaryandEmotion struct {
	Diaryid       int `gorm:"AUTO_INCREMENT"`
	Word     string `sql:"Word"`
	Englishword string `sql:"Englishword"`
	Imageurl string `sql:"imageurl"`
	Userid   string `sql:"userid"`
	Happy     string `json:"Happy"`
	Angry     string  `json:"Angry"`
	Surprise  string  `json:"Surprise"`
	Sad       string  `json:"Sad"`
	Fear      string  `json:"Fear"`
	Name string `sql:"name"`
	CreatedAt time.Time `gorm:"type:datetime(6)"`
	UpdatedAt time.Time `gorm:"type:datetime(6)"`
}