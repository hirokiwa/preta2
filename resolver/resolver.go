package resolver

import (
	"hackz.com/m/v2/client"
	"hackz.com/m/v2/graph/gormmodel"
	"hackz.com/m/v2/graph/model"
	"fmt"
)

func Findfollowee(userid string)([]*model.User, error){
	db, err := client.GetDatabase() 
	if err != nil {
		panic(err)
	}
	var follow []*gormmodel.Follow = []*gormmodel.Follow{}
	var user []*model.User = []*model.User{}
	if err := db.Model(&follow).Select("*").Joins("inner join `users` on follows.follower = users.userid").Where("follows.followee = ?", userid).Scan(&user).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	return user, err
}

func Findfollower(userid string)([]*model.User, error){
	db, err := client.GetDatabase() 
	if err != nil {
		panic(err)
	}
	var follow []*gormmodel.Follow = []*gormmodel.Follow{}
	var user []*model.User = []*model.User{}
	if err := db.Model(&follow).Select("*").Joins("inner join `users` on follows.followee = users.userid").Where("follows.follower = ?", userid).Scan(&user).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	return user, err
}



func FindUser(userid string)(*model.User, error){
	db, err := client.GetDatabase()
	if err != nil {
		panic(err)
	}
	var user *model.User = &model.User{}	
	if err := db.Where("userid = ?", userid).First(&user).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	return user, err
}

func FindDiary(userid string)([]*model.Diary,error){
	db, err := client.GetDatabase();
	if err != nil {
		panic(err)
	}
	var Diary []*model.Diary = []*model.Diary{}
	if err := db.Where("userid = ?",userid).Find(&Diary).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	return Diary, err
}

func FindfolloweeDiary(userid string)([]*model.UserDiary,error){
	var UserDiaryList []*model.UserDiary  = []*model.UserDiary{}

	Followee,err := Findfollowee(userid)
	for i := 0;i < len(Followee);i++{
		fmt.Println(Followee[i].Userid)
		Diary,err := FindDiary(Followee[i].Userid)
		if err  != nil{

		}
		User,err :=FindUser(Followee[i].Userid)
		if err  != nil{

		}
		var UserDiary *model.UserDiary = &model.UserDiary{Diary: Diary,User: User}
		UserDiaryList = append(UserDiaryList, UserDiary)
	}

	return UserDiaryList,err
}

func FindfollowerDiary(userid string)([]*model.UserDiary,error){
	var UserDiaryList []*model.UserDiary

	Follower,err := Findfollower(userid)
	for i := 0;i < len(Follower);i++{
		Diary,err := FindDiary(Follower[i].Userid)
		if err  != nil{

		}
		User,err :=FindUser(Follower[i].Userid)
		if err  != nil{

		}
		var UserDiary *model.UserDiary = &model.UserDiary{Diary: Diary,User: User}
		UserDiaryList = append(UserDiaryList,UserDiary)
	}
	return UserDiaryList,err
}



