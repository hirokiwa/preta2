package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"hackz.com/m/v2/client"
	"hackz.com/m/v2/graph/generated"
	"hackz.com/m/v2/graph/model"
)

// CreateDiary is the resolver for the createDiary field.
func (r *mutationResolver) CreateDiary(ctx context.Context, input model.NewDiary) (*model.Diary, error) {
	panic(fmt.Errorf("not implemented: CreateDiary - createDiary"))
}

// Diary is the resolver for the diary field.
func (r *queryResolver) Diary(ctx context.Context) ([]*model.Diary, error) {
	db, err := client.GetDatabase()
	if err != nil {
		panic(err)
	}

	var diary []*model.Diary
	db.First(&diary) // find product with integer primary ke
	return diary, err
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
