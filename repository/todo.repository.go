package repository

import (
	"context"
	"log"
	"todo/graph/model"
	"todo/prisma"
)

type TodoRepository struct {
	client *prisma.PrismaClient
	ctx    context.Context
}

func NewTodoRepository(client *prisma.PrismaClient, ctx context.Context) *TodoRepository {
	return &TodoRepository{client: client, ctx: ctx}
}

func (repo *TodoRepository) Create(input model.NewTodo) (*prisma.TodoModel, error) {
	todo, err := repo.client.Todo.CreateOne(
		prisma.Todo.Text.Set(input.Text),
		prisma.Todo.User.Link(
			prisma.User.ID.Equals(input.UserID),
		),
	).Exec(repo.ctx)

	if err != nil {
		log.Fatalf("error creating new todo entry error is: %V", err)

		return nil, err
	}

	return todo, nil
}
