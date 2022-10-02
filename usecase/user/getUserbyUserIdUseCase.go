package user

import (
	"hackz.com/m/v2/domain/repository"
	"hackz.com/m/v2/graph/model"
)

type getUserByUserIdUseCaseImpl struct {
	Userid string
	UserRepository repository.UserRepository
}

type getUserByUserIdUseCase interface {
	GetUser() (*model.User,error)
}

func NewGetUserByUserIdImpl(Userid string, userRepository repository.UserRepository) getUserByUserIdUseCase {
	return getUserByUserIdUseCaseImpl{
		Userid: Userid,
		UserRepository: userRepository,
	}
}

func (impl getUserByUserIdUseCaseImpl) GetUser() (*model.User,error) {
	user,err := impl.UserRepository.FindUser(impl.Userid)
	return user,err
}