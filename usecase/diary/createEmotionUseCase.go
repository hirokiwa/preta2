package diary

import (
	"hackz.com/m/v2/domain/repository"
	"hackz.com/m/v2/graph/model"
)

type CreateEmotionUseCaseImpl struct {
	NewEmotion *model.NewEmotion
	DiaryRepository repository.DiaryRepository
}


type CreateEmotionUseCase interface {
	Create()(*model.Emotion,error)
}

func NewCreateEmotionUseCaseImpl(input *model.NewEmotion,diaryrepository repository.DiaryRepository) CreateEmotionUseCase {
	return CreateEmotionUseCaseImpl{
		NewEmotion: input,
		DiaryRepository: diaryrepository,
	}
}

func (impl CreateEmotionUseCaseImpl) Create() (*model.Emotion,error) {
	emotion,err := impl.DiaryRepository.CreateEmotion(*impl.NewEmotion)
	return  emotion,err
}