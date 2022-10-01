package repository

import (
	"hackz.com/m/v2/graph/model"
)


type DiaryRepository interface{
	FindDiary(serid string)([]*model.Diary,error)
	CreateDiary(input model.NewDiary )(*model.Diary,error)
}