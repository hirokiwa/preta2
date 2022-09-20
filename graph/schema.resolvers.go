package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"hackz.com/m/v2/client"
	"hackz.com/m/v2/graph/generated"
	"hackz.com/m/v2/graph/gormmodel"
	"hackz.com/m/v2/graph/model"
	"hackz.com/m/v2/resolver"
)

// CreateDiary is the resolver for the createDiary field.
func (r *mutationResolver) CreateDiary(ctx context.Context, input model.NewDiary) (*model.Diary, error) {
	db, err := client.GetDatabase()
	if err != nil {
		panic(err)
	}
	if err := db.Create(&gormmodel.Diary{Userid: input.Userid, Word: input.Word, Imageurl: input.Imageurl}).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("diary create Error!!!! err:%v\n", err)
	}

	return &model.Diary{}, err
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	db, err := client.GetDatabase()
	if err != nil {
		panic(err)
	}
	if err := db.Create(&gormmodel.User{Userid: input.Userid, Name: input.Name}).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("user create Error!!!! err:%v\n", err)
		return &model.User{
			Userid: input.Userid,
			Name:   input.Name,
		}, err
	}

	return &model.User{
		Userid: input.Userid,
		Name:   input.Name,
	}, err
}

// CreateFollow is the resolver for the createFollow field.
func (r *mutationResolver) CreateFollow(ctx context.Context, input *model.NewFollow) (*model.User, error) {
	db, err := client.GetDatabase()
	if err != nil {
		panic(err)
	}

	var folllower = &gormmodel.User{}

	if err := db.First(&folllower, "userid=?", input.Follower).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("follower user Notfound:%v\n", err)
		return &model.User{
			Userid: input.Follower,
			Name:   folllower.Name,
		}, err
	}

	var folllowee = &gormmodel.User{}

	if err := db.First(&folllowee, "userid=?", input.Followee).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("followee user Notfound:%v\n", err)
		return &model.User{
			Userid: input.Follower,
			Name:   folllower.Name,
		}, err
	}

	if err := db.Create(&gormmodel.Follow{Followee: input.Followee, Follower: input.Follower}).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("followe err:%v\n", err)
		return &model.User{
			Userid: input.Follower,
			Name:   folllower.Name,
		}, err
	}

	return &model.User{
		Userid: input.Follower,
		Name:   folllower.Name,
	}, err
}

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context, argument string) (*model.Me, error) {
	var User *model.User
	var Diary []*model.Diary
	var Followee []*model.UserDiary
	var Follower []*model.UserDiary

	User, err := resolver.FindUser(argument)
	if err != nil {

	}

	Diary, err = resolver.FindDiary(argument)
	if err != nil {

	}

	Followee, err = resolver.FindfolloweeDiary(argument)
	if err != nil {

	}
	Follower, err = resolver.FindfollowerDiary(argument)
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
	Diaries, err := resolver.FindAllDiary()
	if err != nil {

	}
	return Diaries, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Diary(ctx context.Context, argument *string) ([]*model.Diary, error) {
	db, err := client.GetDatabase()
	if err != nil {
		panic(err)
	}

	var query = argument
	var diary []*model.Diary
	if argument != nil {
		if err := db.Find(&diary, "userid = ?", query).Error; err != nil {
			//エラーハンドリング
			fmt.Printf("db select Error!!!! err:%v\n", err)
		}
		return diary, err
	} else {
		if err := db.Find(&diary).Error; err != nil {
			//エラーハンドリング
			fmt.Printf("db select Error!!!! err:%v\n", err)
		}
		return diary, err
	}
}
func (r *queryResolver) Followee(ctx context.Context, argument string) ([]*model.User, error) {
	panic("")
}
func (r *queryResolver) Follower(ctx context.Context, argument string) ([]*model.User, error) {
	db, err := client.GetDatabase()
	if err != nil {
		panic(err)
	}
	var follow []*gormmodel.Follow
	var user []*model.User
	if err := db.Model(&follow).Select("*").Joins("inner join `users` on follows.followee = users.userid").Where("follows.follower = ?", argument).Scan(&user).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	return user, err
}
