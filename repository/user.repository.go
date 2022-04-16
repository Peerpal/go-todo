package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"todo/graph/model"
	"todo/prisma"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	client *prisma.PrismaClient
	ctx    context.Context
}

func NewUserRepository(client *prisma.PrismaClient, ctx context.Context) *UserRepository {
	return &UserRepository{client: client, ctx: ctx}
}

func (repo *UserRepository) Create(input model.NewUser) (*prisma.UserModel, error) {
	// check if a user exists
	if existingUser, _ := repo.userExists(input.Email); existingUser {
		err := errors.Unwrap(fmt.Errorf("a user with the email: %v already exists", input.Email))
		return nil, err
	}

	passwordHash, _ := hashPassword(input.Password)
	// if user does not exist
	newUser, err := repo.client.User.CreateOne(
		prisma.User.FirstName.Set(input.FirstName),
		prisma.User.LastName.Set(input.LastName),
		prisma.User.Email.Set(input.Email),
		prisma.User.Password.Set(passwordHash),
	).Exec(repo.ctx)

	if err != nil {
		log.Fatalf("Error occurred while creating user, %V", err)
		return nil, err
	}

	return newUser, nil
}

func (repo *UserRepository) FindOne(id string) (*prisma.UserModel, error) {
	user, err := repo.client.User.FindFirst(
		prisma.User.ID.Equals(id),
	).Exec(repo.ctx)

	if err != nil {
		log.Fatalf("Error occurred while creating user, %V", err)
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) userExists(email string) (bool, error) {
	user, _ := repo.client.User.FindFirst(
		prisma.User.Email.Equals(email),
	).Exec(repo.ctx)
	if user != nil {
		return true, nil
	} else {
		return false, nil
	}

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password string, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	return err == nil
}
