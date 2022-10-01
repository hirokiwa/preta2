package diary

import (
	"hackz.com/m/v2/domain/repository"
	"hackz.com/m/v2/graph/model"
)

type CreateDiaryUseCaseImpl struct {
	NewDiary *model.NewDiary
	DiaryRepository repository.DiaryRepository
}

type CreateDiaryUseCase interface{
	Create()(*model.Diary,error)
}

func NewCreateDiaryUseCaseImpl(input *model.NewDiary, diaryrepository repository.DiaryRepository) CreateDiaryUseCase {
	return CreateDiaryUseCaseImpl{
		NewDiary: input,
		DiaryRepository: diaryrepository,
	}
}

func (impl CreateDiaryUseCaseImpl) Create() (*model.Diary,error) {
	diary, err := impl.DiaryRepository.CreateDiary(*impl.NewDiary)
	return diary, err
}