package app

import (
	"context"
	"fmt"

	"github.com/iancharters/gostart/internal/db"
	"github.com/iancharters/gostart/internal/db/dao"
)

type App struct {
	Config Config
	db     *db.Client
}

func New(cfg Config, db *db.Client) *App {
	return &App{
		Config: cfg,
		db:     db,
	}
}

func (a *App) GetUser(username string) (*dao.User, error) {
	fmt.Printf("Getting user %s from the database...", username)

	user, err := a.db.GetUserByUsername(context.Background(), username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user '%s': %w", username, err)
	}

	return &user, nil
}

func (a *App) CreateUser(username string, password string) (*dao.User, error) {
	fmt.Printf("Creating user %s...", username)

	dbUser := dao.CreateUserParams{
		Username: username,
		Password: password,
	}

	user, err := a.db.CreateUser(context.Background(), dbUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user '%s': %w", username, err)
	}

	return &user, nil
}
