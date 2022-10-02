package user

import (
	"context"

	"hackz.com/m/v2/graph/model"
	"hackz.com/m/v2/infrastructure/repositoryImpl"
	"hackz.com/m/v2/usecase/user"
)

type UserCreateFollowController struct{}

func (ctrl UserCreateFollowController) Create(ctx context.Context, input *model.NewFollow) (*model.User,error) {
	userRepository := repositoryImpl.NewUserRepositoryImpl()
	result,err := user.NewCreateFollowUseCaseImpl(input,userRepository).Create()
	return result,err
}