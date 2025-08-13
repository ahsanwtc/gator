package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ahsanwtc/gator/internal/database"
	"github.com/google/uuid"
)

func (cmds *Commands) run(state *State, cmd Command) error  {
	command, ok := cmds.commands[cmd.name]
	if !ok {
		return fmt.Errorf("`%s`: command not found", cmd.name)
	}

	return command(state, cmd)
}

func (cmds *Commands) register(name string, f func(state *State, cmd Command) error) error  {
	cmds.commands[name] = f
	return  nil
}

func handlerLogin(s *State, cmd Command) error {
	if cmd.name != "login" {
		return fmt.Errorf("wrong command handler")
	}

	if len(cmd.parameters) != 2 {
		return fmt.Errorf("login expects 2 argument but got %d", len(cmd.parameters))
	}

	username := cmd.parameters[1]
	err := s.config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("%s has been set successfully\n", username)
	return nil
}

func doesUserExists(username string, db *database.Queries) (bool, error) {
	_, err := db.GetUser(context.Background(), username)
	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func handlerRegister(s *State, cmd Command) error {
	if cmd.name != "register" {
		return fmt.Errorf("wrong command handler")
	}

	if len(cmd.parameters) != 2 {
		return fmt.Errorf("login expects 2 argument but got %d", len(cmd.parameters))
	}

	username := cmd.parameters[1]
	_doesUserExists, err := doesUserExists(username, s.db)
	if err != nil {
		fmt.Println("error fetching the user")
		return err
	}

	if _doesUserExists {
		return fmt.Errorf("user already exists")
	}

	_, err = s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		Name: username,
	})
	
	if err != nil {
		fmt.Println("error creating the user")
		return err
	}

	err = s.config.SetUser(username)
	if err != nil {
		fmt.Println("error setting the active user")
		return err
	}

	fmt.Printf("%s has been created successfully\n", username)
	return nil
}