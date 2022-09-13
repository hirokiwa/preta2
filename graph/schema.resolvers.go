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
)

// CreateDiary is the resolver for the createDiary field.
func (r *mutationResolver) CreateDiary(ctx context.Context, input model.NewDiary) (*model.Diary, error) {
	db, err := client.GetDatabase()
	if err != nil {
		panic(err)
	}
	if err := db.Create(&gormmodel.Diary{Userid: input.UserID, Word: input.Word, Imageurl: input.Imageurl}).Error;
	err != nil{
		//エラーハンドリング
		fmt.Printf("diary create Error!!!! err:%v\n", err)
	}

	return &model.Diary{
		Userid:   input.UserID,
		Word:     input.Word,
		Imageurl: input.Imageurl,
	}, err
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	db, err := client.GetDatabase()
	if err != nil {
		panic(err)
	}
	if err := db.Create(&gormmodel.User{ID: input.UserID, Name: input.Name}).Error
		err != nil{
			//エラーハンドリング
			fmt.Printf("user create Error!!!! err:%v\n", err)
			return &model.User{
				ID:   input.UserID,
				Name: input.Name,
			}, err
		}

	return &model.User{
		ID:   input.UserID,
		Name: input.Name,
	}, err
}

// Diary is the resolver for the diary field.
func (r *queryResolver) Diary(ctx context.Context, argument *string) ([]*model.Diary, error) {
	db, err := client.GetDatabase()
	if err != nil {
		panic(err)
	}

	var query = argument
	var diary []*model.Diary
	if argument != nil {
		if err := db.First(&diary, "userid = ?", query).Error;
		err != nil{
			//エラーハンドリング
			fmt.Printf("db select Error!!!! err:%v\n", err)
		}
		return diary, err
	} else {
		if err := db.First(&diary).Error;
		err != nil{
			//エラーハンドリング
			fmt.Printf("db select Error!!!! err:%v\n", err)
		}
		return diary, err
	}
}

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context, argument *string) ([]*model.User, error) {
	db, err := client.GetDatabase()
	if err != nil {
		panic(err)
	}
	var query = argument
	var user []*model.User
	if argument != nil {
		if err := db.First(&user, "id = ?", query).Error;
		err != nil{
			//エラーハンドリング
			fmt.Printf("db select Error!!!! err:%v\n", err)
		}
		return user, err
	} else {
		if err := db.First(&user).Error;
		err != nil{
			//エラーハンドリング
			fmt.Printf("db select Error!!!! err:%v\n", err)
		}
		return user, err
	}

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
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Diary, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewDiary) (*model.Diary, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}
