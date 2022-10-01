package dto

type English struct{
	Diaryid int `gorm:"AUTO_INCREMENT"`
	Englishword string `sql:"Englishword"`
}
