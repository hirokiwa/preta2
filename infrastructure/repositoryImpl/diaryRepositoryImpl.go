package repositoryImpl

import (
	"fmt"
	"strconv"

	"hackz.com/m/v2/domain/repository"
	"hackz.com/m/v2/graph/model"
	"hackz.com/m/v2/infrastructure"
	"hackz.com/m/v2/infrastructure/dto"
)

type DiaryRepositoryImpl struct {}

func NewDiaryRepositoryImpl() repository.DiaryRepository{
	return &DiaryRepositoryImpl{}
}


func (repositoryImpl *DiaryRepositoryImpl) FindDiary(userid string)([]*model.Diary,error){
	db := infrastructure.GetDB()
	var err error

	var Diaries []*model.Diary = []*model.Diary{}
	var GormDiaries []*dto.Diary = []*dto.Diary{}
	var GormDiaryandEmotion []*dto.DiaryandEmotion = []*dto.DiaryandEmotion{}

	if err := db.Model(&GormDiaries).Select("*").Where("diaries.userid=?",userid).Joins("inner join `emotions` on diaries.diaryid = emotions.diaryid").Joins("inner join `users` on diaries.userid = users.userid").Joins("inner join `englishes` on diaries.diaryid = englishes.diaryid").Order("diaries.created_at").Scan(&GormDiaryandEmotion).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	
	for i := 0;i < len(GormDiaryandEmotion);i++ {
		fmt.Printf("(%%#v) %#v\n", GormDiaryandEmotion[i])
		var Emotion *model.Emotion = &model.Emotion{Diaryid:strconv.Itoa(GormDiaryandEmotion[i].Diaryid),Happy: GormDiaryandEmotion[i].Happy,Angry: GormDiaryandEmotion[i].Angry,Surprise:GormDiaryandEmotion[i].Surprise,Sad: GormDiaryandEmotion[i].Sad,Fear: GormDiaryandEmotion[i].Fear}
		var User *model.User = &model.User{Userid: GormDiaryandEmotion[i].Userid,Name: GormDiaryandEmotion[i].Name}
		var Diary *model.Diary = &model.Diary{Diaryid: strconv.Itoa(GormDiaryandEmotion[i].Diaryid),Word:GormDiaryandEmotion[i].Word,Imageurl: GormDiaryandEmotion[i].Imageurl,CreatedAt: GormDiaryandEmotion[i].CreatedAt.String(),UpdatedAt: GormDiaryandEmotion[i].UpdatedAt.String(),User:User,Emotion: Emotion,Englishword: GormDiaryandEmotion[i].Englishword}
		Diaries = append(Diaries, Diary)
	}
	return Diaries, err
}
func (repositoryImpl *DiaryRepositoryImpl) CreateDiary(input model.NewDiary)(*model.Diary,error){
	db := infrastructure.GetDB()
	var err error
	var Diary = &dto.Diary{}
	if err := db.Create(&dto.Diary{Userid: input.Userid, Word: input.Word, Imageurl: input.Imageurl}).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("diary create Error!!!! err:%v\n", err)
	}

	if err := db.Where("imageurl=?", input.Imageurl).First(&Diary).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("diary create Error!!!! err:%v\n", err)
	}

	if err := db.Create(&dto.English{Diaryid: Diary.Diaryid, Englishword: input.Englishword}).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("english create Error!!!! err:%v\n", err)
	}

	fmt.Printf("(%%#v) %#v\n", Diary)

	return &model.Diary{
		Diaryid:   strconv.Itoa(Diary.Diaryid),
		Word:      Diary.Word,
		Imageurl:  Diary.Imageurl,
		CreatedAt: Diary.CreatedAt.String(),
		UpdatedAt: Diary.UpdatedAt.String(),
	}, err
}

func (repositoryImpl *DiaryRepositoryImpl) CreateEmotion(input model.NewEmotion)(*model.Emotion,error) {
	db:= infrastructure.GetDB()
	var err error
	if err := db.Create(&dto.Emotion{Diaryid:input.Diaryid,Happy: input.Happy,Angry: input.Angry,Surprise:input.Surprise,Sad: input.Sad,Fear: input.Fear}).Error; err != nil {
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
