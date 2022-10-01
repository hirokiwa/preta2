package dto

type User struct {
	Userid   string `sql:"userid"`
	Name string `sql:"name"`
	// Diary Diary `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
