package diary

import (
	"context"

	"hackz.com/m/v2/graph/model"
	"hackz.com/m/v2/infrastructure/repositoryImpl"
	"hackz.com/m/v2/usecase/diary"
)

type DiaryGetIdController struct{}

func (ctrl DiaryGetIdController) Get(ctx context.Context,argumnet *string )([]*model.Diary,error) {
	diaryRepository := repositoryImpl.NewDiaryRepositoryImpl()
	result,err := diary.NewGetDiaryFromUseCaseImpl(*argumnet,diaryRepository).Get()
	return result,err
}