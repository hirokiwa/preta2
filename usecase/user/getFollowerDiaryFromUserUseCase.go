package user

import (
	"hackz.com/m/v2/domain/repository"
	"hackz.com/m/v2/graph/model"
)

type getFollowerDiaryUseCaseImpl struct {
	Userid string
	UserRepository repository.UserRepository
	DiaryRepository repository.DiaryRepository
}

type getFollowerDiaryUseCase interface {
	GetFollowerDiary()([]*model.UserDiary,error)
}

func NewGetFollowerDiaryUseCaseImpl(Userid string,userRepository repository.UserRepository, diaryReoisitory repository.DiaryRepository) getFollowerDiaryUseCase{
	return getFollowerDiaryUseCaseImpl{
		Userid: Userid,
		UserRepository: userRepository,
		DiaryRepository: diaryReoisitory, 
	}
}

func (impl getFollowerDiaryUseCaseImpl) GetFollowerDiary() ([]*model.UserDiary,error){
	FollowerDiary,err := impl.UserRepository.FindFollowerDiary(impl.Userid)
	return FollowerDiary,err
}