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

// func FindDiary(userid string)([]*model.Diary,error){
// 	db, err := client.GetDatabase();
// 	if err != nil {
// 		panic(err)
// 	}

// 	// var User *model.User = &model.User{}
// 	// if err := db.Where("userid = ?",userid).First(&User).Error; err != nil {
// 	// 	//エラーハンドリング
// 	// 	fmt.Printf("db select Error!!!! err:%v\n", err)
// 	// }

// 	var Diaries []*gormmodel.Diary = []*gormmodel.Diary{}
// 	var DiaryandUser []*gormmodel.DiaryandUser = []*gormmodel.DiaryandUser{}
// 	if err := db.Model(&Diaries).Select("*").Joins("inner join `users` on diaries.userid = users.userid").Order("diaries.created_at").Scan(&DiaryandUser).Error; err != nil {
// 		//エラーハンドリング
// 		fmt.Printf("db select Error!!!! err:%v\n", err)
// 	}
	
// 	if err := db.Where("userid = ?",userid).Find(&Diaries).Error; err != nil {
// 		//エラーハンドリング
// 		fmt.Printf("db select Error!!!! err:%v\n", err)
// 	}

// 	var DiaryList []*model.Diary = []*model.Diary{}
// 	for i:=0;i<len(Diaries);i++{
// 		// var Emotion *model.Emotion = &model.Emotion{}
// 		Emotion,err := FindEmotion(Diaries[i].Diaryid)
// 		if err != nil{

// 		}
// 		var Diary *model.Diary = &model.Diary{Diaryid: strconv.Itoa(Diaries[i].Diaryid),Word:Diaries[i].Word,Imageurl: Diaries[i].Imageurl,CreatedAt: Diaries[i].CreatedAt.String(),UpdatedAt: Diaries[i].UpdatedAt.String(),Userid: userid,Emotion: Emotion}
// 		DiaryList = append(DiaryList, Diary)
// 	}
// 	return DiaryList, err
// }

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
	var GormDiaryandEmotion []*gormmodel.DiaryandEmotion = []*gormmodel.DiaryandEmotion{}

	if err := db.Model(&GormDiaries).Select("*").Joins("inner join `emotions` on diaries.diaryid = emotions.diaryid").Joins("inner join `users` on diaries.userid = users.userid").Order("diaries.created_at").Scan(&GormDiaryandEmotion).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}

	// if err := db.Order("created_at").Find(&GormDiaries).Error; err != nil {
	// 	//エラーハンドリング
	// 	fmt.Printf("db select Error!!!! err:%v\n", err)
	// }
	
	for i := 0;i < len(GormDiaryandEmotion);i++ {
		// var User *model.User = &model.User{}
		// if err := db.Where("userid = ?",GormDiaries[i].Userid).First(&User).Error; err != nil {
		// 	//エラーハンドリング
		// 	fmt.Printf("db select Error!!!! err:%v\n", err)
		// }
		fmt.Printf("(%%#v) %#v\n", GormDiaryandEmotion[i])
		var Emotion *model.Emotion = &model.Emotion{Diaryid:strconv.Itoa(GormDiaryandEmotion[i].Diaryid),Happy: GormDiaryandEmotion[i].Happy,Angry: GormDiaryandEmotion[i].Angry,Surprise:GormDiaryandEmotion[i].Surprise,Sad: GormDiaryandEmotion[i].Sad,Fear: GormDiaryandEmotion[i].Fear}
		var User *model.User = &model.User{Userid: GormDiaryandEmotion[i].Userid,Name: GormDiaryandEmotion[i].Name}
		var Diary *model.Diary = &model.Diary{Diaryid: strconv.Itoa(GormDiaryandEmotion[i].Diaryid),Word:GormDiaryandEmotion[i].Word,Imageurl: GormDiaryandEmotion[i].Imageurl,CreatedAt: GormDiaryandEmotion[i].CreatedAt.String(),UpdatedAt: GormDiaryandEmotion[i].UpdatedAt.String(),User:User,Emotion: Emotion}
		Diaries = append(Diaries, Diary)
	}
	return Diaries, err
}


func FindEmotion(diaryid int)(*model.Emotion,error){
	db, err := client.GetDatabase();
	if err != nil {
		panic(err)
	}
	print("diaryid",diaryid)
	var Emotion *gormmodel.Emotion = &gormmodel.Emotion{}
	if err := db.Where("diaryid=?",strconv.Itoa(diaryid)).First(&Emotion).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	return &model.Emotion{
		Diaryid:Emotion.Diaryid,
		Happy: Emotion.Happy,
		Angry: Emotion.Angry,
		Surprise:Emotion.Surprise,
		Sad: Emotion.Sad,
		Fear: Emotion.Fear ,
	}, err

}

func CreateNewEmotion(input *model.NewEmotion)(*model.Emotion,error){
	db, err := client.GetDatabase();
	if err != nil {
		panic(err)
	}
	if err := db.Create(&gormmodel.Emotion{Diaryid:input.Diaryid,Happy: input.Happy,Angry: input.Angry,Surprise:input.Surprise,Sad: input.Sad,Fear: input.Fear}).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("followe err:%v\n", err)
	}
	return &model.Emotion{
		Diaryid:input.Diaryid,
		Happy: input.Happy,
		Angry: input.Angry,
		Surprise:input.Surprise,
		Sad: input.Sad,
		Fear: input.Fear ,
	}, err
}



func FindDiary(userid string)([]*model.Diary, error){
	db, err := client.GetDatabase();
	if err != nil {
		panic(err)
	}

	var Diaries []*model.Diary = []*model.Diary{}
	var GormDiaries []*gormmodel.Diary = []*gormmodel.Diary{}
	var GormDiaryandEmotion []*gormmodel.DiaryandEmotion = []*gormmodel.DiaryandEmotion{}

	if err := db.Model(&GormDiaries).Select("*").Where("diaries.userid=?",userid).Joins("inner join `emotions` on diaries.diaryid = emotions.diaryid").Joins("inner join `users` on diaries.userid = users.userid").Order("diaries.created_at").Scan(&GormDiaryandEmotion).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}

	// if err := db.Order("created_at").Find(&GormDiaries).Error; err != nil {
	// 	//エラーハンドリング
	// 	fmt.Printf("db select Error!!!! err:%v\n", err)
	// }
	
	for i := 0;i < len(GormDiaryandEmotion);i++ {
		// var User *model.User = &model.User{}
		// if err := db.Where("userid = ?",GormDiaries[i].Userid).First(&User).Error; err != nil {
		// 	//エラーハンドリング
		// 	fmt.Printf("db select Error!!!! err:%v\n", err)
		// }
		fmt.Printf("(%%#v) %#v\n", GormDiaryandEmotion[i])
		var Emotion *model.Emotion = &model.Emotion{Diaryid:strconv.Itoa(GormDiaryandEmotion[i].Diaryid),Happy: GormDiaryandEmotion[i].Happy,Angry: GormDiaryandEmotion[i].Angry,Surprise:GormDiaryandEmotion[i].Surprise,Sad: GormDiaryandEmotion[i].Sad,Fear: GormDiaryandEmotion[i].Fear}
		var User *model.User = &model.User{Userid: GormDiaryandEmotion[i].Userid,Name: GormDiaryandEmotion[i].Name}
		var Diary *model.Diary = &model.Diary{Diaryid: strconv.Itoa(GormDiaryandEmotion[i].Diaryid),Word:GormDiaryandEmotion[i].Word,Imageurl: GormDiaryandEmotion[i].Imageurl,CreatedAt: GormDiaryandEmotion[i].CreatedAt.String(),UpdatedAt: GormDiaryandEmotion[i].UpdatedAt.String(),User:User,Emotion: Emotion}
		Diaries = append(Diaries, Diary)
	}
	return Diaries, err
}



