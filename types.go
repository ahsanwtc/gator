package main

import "github.com/ahsanwtc/gator/internal/config"

type State struct {
	config *config.Config
}

type Command struct {
	name string
	parameters []string
}

type Commands struct {
	commands map[string]func(*State, Command)error
}