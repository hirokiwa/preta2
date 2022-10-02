package user

import (
	"context"

	"hackz.com/m/v2/graph/model"
	"hackz.com/m/v2/infrastructure/repositoryImpl"
	"hackz.com/m/v2/usecase/user"
)

type UserGetFollowerDiaryController struct{}

func (ctrl UserGetFollowerDiaryController) Get(ctx context.Context,argument *string)([]*model.UserDiary,error) {
	userRepository := repositoryImpl.NewUserRepositoryImpl()
	diaryRepository := repositoryImpl.NewDiaryRepositoryImpl()
	result,err := user.NewGetFollowerDiaryUseCaseImpl(*argument,userRepository,diaryRepository).GetFollowerDiary()
	return result,err
}