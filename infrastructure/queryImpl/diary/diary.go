package diary

import (
	"fmt"
	"strconv"

	"hackz.com/m/v2/graph/model"
	"hackz.com/m/v2/infrastructure"
	"hackz.com/m/v2/infrastructure/dto"
	"hackz.com/m/v2/query"
)

type DiaryQueryImpl struct{}

func NewdiaryQueryImpl() query.DiaryQuery{
	return &DiaryQueryImpl{}
}

func (repo DiaryQueryImpl) FindAll() []*model.Diary{
	db := infrastructure.GetDB()



	var Diaries []*model.Diary = []*model.Diary{}
	var GormDiaries []*dto.Diary = []*dto.Diary{}
	var GormDiaryandEmotion []*dto.DiaryandEmotion = []*dto.DiaryandEmotion{}

	if err := db.Model(&GormDiaries).Select("*").Joins("inner join `emotions` on diaries.diaryid = emotions.diaryid").Joins("inner join `users` on diaries.userid = users.userid").Joins("inner join `englishes` on diaries.diaryid = englishes.diaryid").Order("diaries.created_at").Scan(&GormDiaryandEmotion).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	
	for i := 0;i < len(GormDiaryandEmotion);i++ {
		fmt.Printf("(%%#v) %#v\n", GormDiaryandEmotion[i])
		var Emotion *model.Emotion = &model.Emotion{Diaryid:strconv.Itoa(GormDiaryandEmotion[i].Diaryid),Happy: GormDiaryandEmotion[i].Happy,Angry: GormDiaryandEmotion[i].Angry,Surprise:GormDiaryandEmotion[i].Surprise,Sad: GormDiaryandEmotion[i].Sad,Fear: GormDiaryandEmotion[i].Fear}
		var User *model.User = &model.User{Userid: GormDiaryandEmotion[i].Userid,Name: GormDiaryandEmotion[i].Name}
		var Diary *model.Diary = &model.Diary{Diaryid: strconv.Itoa(GormDiaryandEmotion[i].Diaryid),Word:GormDiaryandEmotion[i].Word,Imageurl: GormDiaryandEmotion[i].Imageurl,CreatedAt: GormDiaryandEmotion[i].CreatedAt.String(),UpdatedAt: GormDiaryandEmotion[i].UpdatedAt.String(),User:User,Emotion: Emotion,Englishword:GormDiaryandEmotion[i].Englishword}
		Diaries = append(Diaries, Diary)
	}
	return Diaries
}