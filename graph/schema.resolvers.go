package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"todo/graph/generated"
	"todo/graph/model"
	"todo/prisma"
	"todo/repository"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*prisma.TodoModel, error) {
	todoRepository := repository.NewTodoRepository(r.client, ctx)

	todo, err := todoRepository.Create(input)

	if err != nil {
		log.Fatal("Error creating Todo")
		return nil, err
	}

	return todo, nil

}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*prisma.UserModel, error) {
	userRepository := repository.NewUserRepository(r.client, ctx)

	user, err := userRepository.Create(input)

	if err != nil {
		log.Fatal("Error creating User")
		return nil, err
	}

	return user, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*prisma.TodoModel, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todo(ctx context.Context) (*prisma.TodoModel, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*prisma.UserModel, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*prisma.UserModel, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
