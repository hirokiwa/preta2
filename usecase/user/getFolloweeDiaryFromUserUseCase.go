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
	var UserDiaryList []*model.UserDiary  = []*model.UserDiary{}
	Followee,err := impl.UserRepository.Findfollowee(impl.Userid)
	for i := 0;i < len(Followee);i++{
		Diary,err := impl.DiaryRepository.FindDiary(Followee[i].Userid)
		if err  != nil{

		}
		var UserDiary *model.UserDiary = &model.UserDiary{Diary: Diary,User: Followee[i]}
		UserDiaryList = append(UserDiaryList, UserDiary)
	}
	return UserDiaryList,err

}