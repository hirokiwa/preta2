package resolver

import (
	"fmt"
	"strconv"
	"hackz.com/m/v2/client"
	"hackz.com/m/v2/graph/gormmodel"
	"hackz.com/m/v2/graph/model"
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

	var User *model.User = &model.User{}
	if err := db.Where("userid = ?",userid).First(&User).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}

	var Diaries []*gormmodel.Diary = []*gormmodel.Diary{}
	if err := db.Where("userid = ?",userid).Find(&Diaries).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}

	var DiaryList []*model.Diary = []*model.Diary{}
	for i:=0;i<len(Diaries);i++{
		var Diary *model.Diary = &model.Diary{Diaryid: strconv.Itoa(Diaries[i].Diaryid),Word:Diaries[i].Word,Imageurl: Diaries[i].Imageurl,CreatedAt: Diaries[i].CreatedAt.String(),UpdatedAt: Diaries[i].UpdatedAt.String(),User:  User}
		DiaryList = append(DiaryList, Diary)
	}
	return DiaryList, err
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

func FindAllDiary()([]*model.Diary, error){
	db, err := client.GetDatabase();
	if err != nil {
		panic(err)
	}

	var Diaries []*model.Diary = []*model.Diary{}
	var GormDiaries []*gormmodel.Diary = []*gormmodel.Diary{}

	if err := db.Order("created_at").Find(&GormDiaries).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	for i := 0;i < len(GormDiaries);i++ {
		var User *model.User = &model.User{}
		if err := db.Where("userid = ?",GormDiaries[i].Userid).First(&User).Error; err != nil {
			//エラーハンドリング
			fmt.Printf("db select Error!!!! err:%v\n", err)
		}
		var Diary *model.Diary = &model.Diary{Diaryid: strconv.Itoa(GormDiaries[i].Diaryid),Word:GormDiaries[i].Word,Imageurl: GormDiaries[i].Imageurl,CreatedAt: GormDiaries[i].CreatedAt.String(),UpdatedAt: GormDiaries[i].UpdatedAt.String(),User:  User}
		Diaries = append(Diaries, Diary)
	}
	return Diaries, err
}


