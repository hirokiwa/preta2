package user

import (
	"hackz.com/m/v2/domain/repository"
	"hackz.com/m/v2/graph/model"
)

type CreateFollowUseCaseImpl struct {
	NewFollow *model.NewFollow
	UserRepository repository.UserRepository
}

type CreateFollowUseCase interface{
	Create()(*model.User,error)
}

func NewCreateFollowUseCaseImpl (input *model.NewFollow, userrepository repository.UserRepository) CreateFollowUseCase{
	return CreateFollowUseCaseImpl{
		NewFollow: input,
		UserRepository: userrepository,
	}
}

func (impl CreateFollowUseCaseImpl) Create() (*model.User,error) {
	user,err := impl.UserRepository.CreateFollow(*impl.NewFollow)
	return user,err
}