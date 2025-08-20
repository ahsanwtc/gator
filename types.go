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

type Commands struct {
	commands map[string]func(*State, Command)error
}