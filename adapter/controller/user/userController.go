package user

import (
	"context"

	"hackz.com/m/v2/graph/model"
	"hackz.com/m/v2/infrastructure/repositoryImpl"
	"hackz.com/m/v2/usecase/user"
)

type UserController struct{}

func (ctrl UserController) Show(ctx context.Context, argument *string)(*model.User,error){
	userRepository := repositoryImpl.NewUserRepositoryImpl()

	result,err := user.NewGetUserByUserIdImpl(*argument,userRepository).GetUser()

	return result,err
}