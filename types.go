package main

import (
	"github.com/ahsanwtc/gator/internal/config"
	"github.com/ahsanwtc/gator/internal/database"
	"github.com/ahsanwtc/gator/internal/services"
)

type State struct {
	config *config.Config
	db  *database.Queries
	userService *services.UserService
}

type Command struct {
	name string
	parameters []string
}

type CommandHandler func(s *State, cmd Command) error

type CommandHandlerWithUser func(s *State, cmd Command, user database.User) error

type Commands struct {
	commands map[string]CommandHandler
}
