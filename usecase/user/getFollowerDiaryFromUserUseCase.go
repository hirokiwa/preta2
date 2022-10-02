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
	var UserDiaryList []*model.UserDiary
	Follower ,err := impl.UserRepository.Findfollower(impl.Userid)
	for i := 0;i < len(Follower);i++{
		Diary,err := impl.DiaryRepository.FindDiary(Follower[i].Userid)
		if err  != nil{

		}
		var UserDiary *model.UserDiary = &model.UserDiary{Diary: Diary,User: Follower[i]}
		UserDiaryList = append(UserDiaryList,UserDiary)
	}
	return UserDiaryList,err

	
}