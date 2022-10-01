package dto

import "time"



type Emotion struct {
	Diaryid   string  `gorm:"primaryKey"`
	Happy     string `json:"Happy"`
	Angry     string  `json:"Angry"`
	Surprise  string  `json:"Surprise"`
	Sad       string  `json:"Sad"`
	Fear      string  `json:"Fear"`
	CreatedAt time.Time  `gorm:"type:datetime(6)"`
	UpdatedAt time.Time  `gorm:"type:datetime(6)"`
}
