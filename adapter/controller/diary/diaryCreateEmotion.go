package diary

import (
	"context"

	"hackz.com/m/v2/graph/model"
	"hackz.com/m/v2/infrastructure/repositoryImpl"
	"hackz.com/m/v2/usecase/diary"
)

type DiaryCreateEmotionController struct{}

func (ctrl DiaryCreateEmotionController) Create(ctx context.Context,input *model.NewEmotion)(*model.Emotion,error) {
	diaryRepository := repositoryImpl.NewDiaryRepositoryImpl()
	result,err := diary.NewCreateEmotionUseCaseImpl(input,diaryRepository).Create()
	return result,err
} 