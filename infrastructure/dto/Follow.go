package dto


type Follow struct {
	Followee string `gorm:"primaryKey"`
	Follower string `gorm:"primaryKey"`
}