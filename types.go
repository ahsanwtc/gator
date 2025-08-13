package main

import (
	"github.com/ahsanwtc/gator/internal/config"
	"github.com/ahsanwtc/gator/internal/database"
)

type State struct {
	config *config.Config
	db  *database.Queries
}

type Command struct {
	name string
	parameters []string
}

type Commands struct {
	commands map[string]func(*State, Command)error
}