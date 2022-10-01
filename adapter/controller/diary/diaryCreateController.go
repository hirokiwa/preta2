package diary

import (
	"context"

	"hackz.com/m/v2/graph/model"
	"hackz.com/m/v2/infrastructure/repositoryImpl"
	"hackz.com/m/v2/usecase/diary"
)


type DiaryCreateController struct{}

func (ctrl DiaryCreateController) Create(ctx context.Context, input *model.NewDiary)(*model.Diary,error) {
	diaryRepository := repositoryImpl.NewDiaryRepositoryImpl()

	result,err := diary.NewCreateDiaryUseCaseImpl(input,diaryRepository).Create()
	return result,err
}