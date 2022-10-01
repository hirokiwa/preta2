package diary

import (
	"hackz.com/m/v2/domain/repository"
	"hackz.com/m/v2/graph/model"
)

type getDiaryFromUseCaseImpl struct {
	Userid string
	DiaryRepository repository.DiaryRepository
}

type getDiaryFromUseCase interface {
	Get()([]*model.Diary,error)
}

func NewGetDiaryFromUseCaseImpl(Userid string , diaryrepository repository.DiaryRepository) getDiaryFromUseCase {
	return getDiaryFromUseCaseImpl{
		Userid: Userid,
		DiaryRepository:diaryrepository,
	}
}

func (impl getDiaryFromUseCaseImpl) Get() ([]*model.Diary,error){
	diary,err := impl.DiaryRepository.FindDiary(impl.Userid)
	return diary,err
}