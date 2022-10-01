package diary

import (
	"context"

	"hackz.com/m/v2/graph/model"
	"hackz.com/m/v2/infrastructure/queryImpl/diary"
)

type DairyController struct{}

func (ctrl DairyController) Show(ctx context.Context)([]*model.Diary){
	result := diary.NewdiaryQueryImpl().FindAll()
	return result
}