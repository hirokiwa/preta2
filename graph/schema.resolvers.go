package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	diaryCtrl "hackz.com/m/v2/adapter/controller/diary"
	"hackz.com/m/v2/adapter/controller/user"
	"hackz.com/m/v2/graph/generated"
	"hackz.com/m/v2/graph/model"
	diaryImpl "hackz.com/m/v2/infrastructure/queryImpl/diary"
)

// CreateDiary is the resolver for the createDiary field.
func (r *mutationResolver) CreateDiary(ctx context.Context, input model.NewDiary) (*model.Diary, error) {
	diaryCtrl := diaryCtrl.DiaryCreateController{}
	Diary, err := diaryCtrl.Create(ctx, &input)
	return Diary, err
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	userCtrl := user.UserCreateController{}
	User, err := userCtrl.Create(ctx, input)
	return User, err
}

// CreateFollow is the resolver for the createFollow field.
func (r *mutationResolver) CreateFollow(ctx context.Context, input *model.NewFollow) (*model.User, error) {
	userCtrl := user.UserCreateFollowController{}
	User, err := userCtrl.Create(ctx, input)
	return User, err
}

// CreateEmotion is the resolver for the createEmotion field.
func (r *mutationResolver) CreateEmotion(ctx context.Context, input *model.NewEmotion) (*model.Emotion, error) {
	diaryCtrl := diaryCtrl.DiaryCreateEmotionController{}
	Emotion, err := diaryCtrl.Create(ctx, input)
	return Emotion, err
}

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context, argument string) (*model.Me, error) {
	userCtrl := user.UserController{}
	User, err := userCtrl.Show(ctx, &argument)
	if err != nil {
	}

	diaryCtrl := diaryCtrl.DiaryGetIdController{}
	Diary, err := diaryCtrl.Get(ctx, &argument)
	if err != nil {
	}

	userFolloweeCtrl := user.UserGetFolloweeDiaryController{}
	Followee, err := userFolloweeCtrl.Get(ctx, &argument)
	if err != nil {
	}

	userFollowerCtrl := user.UserGetFollowerDiaryController{}
	Follower, err := userFollowerCtrl.Get(ctx, &argument)
	if err != nil {
	}
	return &model.Me{
		User:     User,
		Diary:    Diary,
		Followee: Followee,
		Follower: Follower,
	}, err
}

// AllDiary is the resolver for the AllDiary field.
func (r *queryResolver) AllDiary(ctx context.Context) ([]*model.Diary, error) {
	// var Diary []*model.Diary
	var err error
	diaryCtrl := diaryImpl.DiaryQueryImpl{}
	Diary := diaryCtrl.FindAll()
	if err != nil {
	}
	return Diary, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
