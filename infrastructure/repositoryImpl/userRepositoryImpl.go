package repositoryImpl

import (
	"fmt"

	"hackz.com/m/v2/domain/repository"
	"hackz.com/m/v2/graph/model"
	"hackz.com/m/v2/infrastructure"
	"hackz.com/m/v2/infrastructure/dto"
)

type UserRepositoryImpl struct {}


func NewUserRepositoryImpl() repository.UserRepository{
	return &UserRepositoryImpl{}
}

func (repositoryImpl *UserRepositoryImpl) Findfollowee(followeeid string)([]*model.User,error){
	db := infrastructure.GetDB()
	var err error
	var follow []*dto.Follow = []*dto.Follow{}
	var user []*model.User = []*model.User{}
	if err := db.Model(&follow).Select("*").Joins("inner join `users` on follows.follower = users.userid").Where("follows.followee = ?", userid).Scan(&user).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	return user, err
	
}

func (repositoryImpl *UserRepositoryImpl) Findfollower(followerid string)([]*model.User,error){
	db := infrastructure.GetDB()
	var err error
	var follow []*dto.Follow = []*dto.Follow{}
	var user []*model.User = []*model.User{}
	if err := db.Model(&follow).Select("*").Joins("inner join `users` on follows.followee = users.userid").Where("follows.follower = ?", userid).Scan(&user).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	return user, err
	
}
func (repositoryImpl *UserRepositoryImpl) FindUser(followerid string)(*model.User,error){
		db := infrastructure.GetDB()
		var err error
		var user *model.User = &model.User{}	
		if err := db.Where("userid = ?", followerid).First(&user).Error; err != nil {
			//エラーハンドリング
			fmt.Printf("db select Error!!!! err:%v\n", err)
		}
		return user, err
	
	
}
func (repositoryImpl *UserRepositoryImpl) CreateUser(input model.NewUser)(*model.User,error){
	db := infrastructure.GetDB()
	var err error
	if err := db.Create(&gormmodel.User{Userid: input.Userid, Name: input.Name}).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("user create Error!!!! err:%v\n", err)
		return &model.User{
			Userid: input.Userid,
			Name:   input.Name,
		}, err
	}

	return &model.User{
		Userid: input.Userid,
		Name:   input.Name,
	}, err
}
