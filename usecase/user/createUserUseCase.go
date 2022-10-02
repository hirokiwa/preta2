package user

import (
	"hackz.com/m/v2/domain/repository"
	"hackz.com/m/v2/graph/model"
)


type CreateUserUseCaseImpl struct {
	NewUser *model.NewUser
	UserRepository repository.UserRepository
}

type CreateUserUseCase interface{
	Create()(*model.User,error)
}

func NewCreateUserUseCaseImpl (input *model.NewUser, userrepository repository.UserRepository) CreateUserUseCase{
	return CreateUserUseCaseImpl{
		NewUser: input,
		UserRepository: userrepository,
	}
}

func (impl CreateUserUseCaseImpl) Create() (*model.User,error){
	user,err := impl.UserRepository.CreateUser(*impl.NewUser)
	return user,err
}