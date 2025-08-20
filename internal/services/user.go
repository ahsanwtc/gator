package services

import (
	"context"
	"fmt"

	"github.com/ahsanwtc/gator/internal/database"
	"github.com/google/uuid"
)

type UserService struct {
	db *database.Queries
}

func NewUserService(db *database.Queries) *UserService  {
	return &UserService{
		db: db,
	}
}

func (us *UserService) FetchUserById(id uuid.UUID) (database.User, error) {
	var user database.User
	user, err := us.db.GetUserById(context.Background(), id)
	if err != nil {
		return user, fmt.Errorf("error fetching user from the database: %s", err)
	}

	return user, nil
}

func (us *UserService) FetchUserByName(name string) (database.User, error) {
	var user database.User
	user, err := us.db.GetUser(context.Background(), name)
	if err != nil {
		return user, fmt.Errorf("error fetching user from the database: %s", err)
	}

	return user, nil
}