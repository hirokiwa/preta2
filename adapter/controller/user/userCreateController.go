package user

import (
	"context"

	"hackz.com/m/v2/graph/model"
	"hackz.com/m/v2/infrastructure/repositoryImpl"
	"hackz.com/m/v2/usecase/user"
)

type UserCreateController struct{}

func (ctrl UserCreateController) Create(ctx context.Context, input *model.NewUser)(*model.User,error) {
	userRepository := repositoryImpl.NewUserRepositoryImpl()

	result,err := user.NewCreateUserUseCaseImpl(input,userRepository).Create()

	return result,err
}