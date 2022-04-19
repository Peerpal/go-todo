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
	"todo/utils"
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

func (r *mutationResolver) MarkTodoAsDone(ctx context.Context, id string) (bool, error) {
	todoRepository := repository.NewTodoRepository(r.client, ctx)

	return todoRepository.UpdateDoneStatus(id, true), nil
}

func (r *mutationResolver) MarkTodoAsUndone(ctx context.Context, id string) (bool, error) {
	todoRepository := repository.NewTodoRepository(r.client, ctx)

	return todoRepository.UpdateDoneStatus(id, false), nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	userRepository := repository.NewUserRepository(r.client, ctx)

	user, _ := userRepository.FindOne(input.Email)

	if user == nil {
		return nil, fmt.Errorf("User not found")
	}

	if user.Password != input.Password {
		return nil, fmt.Errorf("Invalid password")
	}

	// create a new token
	JWT := utils.JWT{
		Secret: "jwtsigingkeyhere",
	}

	// create token
	token, _ := JWT.CreateToken(user.ID)

	return &model.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*prisma.TodoModel, error) {
	todoRepository := repository.NewTodoRepository(r.client, ctx)

	todos, err := todoRepository.FindMany()

	if err != nil {
		log.Fatal("Error Fetching Todos")
		return nil, err
	}

	// convert to a slice of pointers
	todosPointer := make([]*prisma.TodoModel, len(todos))

	for index := range todos {
		todosPointer[index] = &todos[index]

	}

	return todosPointer, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*prisma.TodoModel, error) {
	todoRepository := repository.NewTodoRepository(r.client, ctx)

	todo, err := todoRepository.FindOne(id)

	if err != nil {
		log.Fatal("Error Fetching Todo")
		return nil, err
	}

	return todo, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*prisma.UserModel, error) {
	userRepository := repository.NewUserRepository(r.client, ctx)

	users, err := userRepository.FindMany()

	if err != nil {
		log.Fatal("Error Fetching Users")
		return nil, err
	}

	return users, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id string) (*prisma.UserModel, error) {
	userRepository := repository.NewUserRepository(r.client, ctx)

	user, err := userRepository.FindOne(id)

	if err != nil {
		log.Fatal("Error Fetching Users")
		return nil, err
	}

	return user, nil
}

func (r *queryResolver) UserTodos(ctx context.Context, id string) ([]*prisma.TodoModel, error) {
	todoRepository := repository.NewTodoRepository(r.client, ctx)

	todos, err := todoRepository.TodosByUser(id)

	if err != nil {
		log.Fatalf("Error fetching todos %V: ", err)

		return nil, err
	}

	return todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *prisma.TodoModel) (*prisma.UserModel, error) {
	userRepository := repository.NewUserRepository(r.client, ctx)

	user, err := userRepository.FindOne(obj.UserID)

	if err != nil {
		log.Fatal("Error Fetching Users")
		return nil, err
	}

	return user, nil
}

func (r *userResolver) Todos(ctx context.Context, obj *prisma.UserModel) ([]*prisma.TodoModel, error) {
	todoRepository := repository.NewTodoRepository(r.client, ctx)

	todos, err := todoRepository.TodosByUser(obj.ID)

	if err != nil {
		log.Fatalf("Error fetching todos %V: ", err)

		return nil, err
	}

	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) User(ctx context.Context, id string) (*prisma.UserModel, error) {
	userRepository := repository.NewUserRepository(r.client, ctx)

	user, err := userRepository.FindOne(id)

	if err != nil {
		log.Fatal("Error Fetching Users")
		return nil, err
	}

	return user, nil
}
