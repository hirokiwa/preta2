package repository

import (
	"hackz.com/m/v2/graph/model"
)

type UserRepository interface {
	Findfollowee(followerid string)([]*model.User,error)
	Findfollower(followeeid string)([]*model.User,error)
	FindUser(userid string)(*model.User)
	CreateUser(input model.NewUser)(*model.User)
}





