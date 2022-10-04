package user

import (
	"hackz.com/m/v2/domain/repository"
	"hackz.com/m/v2/graph/model"
)

type getFolloweeDiaryUseCaseImpl struct{
	Userid string
	UserRepository repository.UserRepository
	DiaryRepository repository.DiaryRepository
}

type getFolloweeDiaryUseCase interface {
	GetFolloweeDiary()([]*model.UserDiary,error)
}

func NewGetFolloweeUseCaseImpl(Userid string,userrepository repository.UserRepository, diaryRepository repository.DiaryRepository) getFolloweeDiaryUseCase{
	return getFolloweeDiaryUseCaseImpl{
		Userid: Userid,
		UserRepository: userrepository,
		DiaryRepository: diaryRepository,
	}
}

func (impl getFolloweeDiaryUseCaseImpl) GetFolloweeDiary() ([]*model.UserDiary,error){
	FolloweeDiary,err := impl.UserRepository.FindFolloweeDiary(impl.Userid)
	return FolloweeDiary,err

}