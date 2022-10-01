package diary

import (
	"context"

	"hackz.com/m/v2/infrastructure/repositoryImpl"
)


type DiaryCreateController struct{}

func (ctrl DiaryCreateController) Create(ctx context.Context) {
	result,err := repositoryImpl.NewUserRepositoryImpl().CreateUser()
}