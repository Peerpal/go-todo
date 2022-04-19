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

func (repo *TodoRepository) FindOne(id string) (*prisma.TodoModel, error) {
	todo, err := repo.client.Todo.FindFirst(
		prisma.Todo.ID.Equals(id),
	).Exec(repo.ctx)

	if err != nil {
		log.Fatal("Error Fetching Todo")
		return nil, err
	}

	return todo, nil
}

func (repo *TodoRepository) FindMany() ([]prisma.TodoModel, error) {
	todos, err := repo.client.Todo.FindMany().OrderBy(
		prisma.Todo.CreatedAt.Order(prisma.SortOrderDesc),
	).Exec(repo.ctx)

	if err != nil {
		log.Fatalf("Error fetching todos %V: ", err)
		return nil, err
	}

	return todos, nil
}

func (repo *TodoRepository) UpdateDoneStatus(id string, status bool) bool {
	_, err := repo.client.Todo.FindUnique(
		prisma.Todo.ID.Equals(id),
	).Update(
		prisma.Todo.Done.Set(status),
	).Exec(repo.ctx)

	if err != nil {
		log.Fatalf("error updating todo %V", err)
		return false
	}

	return true
}

func (repo *TodoRepository) TodosByUser(userId string) ([]*prisma.TodoModel, error) {
	results, err := repo.client.Todo.FindMany(
		prisma.Todo.UserID.Equals(userId),
	).Exec(repo.ctx)

	if err != nil {
		return nil, err
	}

	// transform slice
	todos := transformTodoToPointer(results)

	return todos, nil
}

func transformTodoToPointer(results []prisma.TodoModel) []*prisma.TodoModel {
	data := make([]*prisma.TodoModel, len(results))

	for i := range results {
		data[i] = &results[i]
	}

	return data
}
