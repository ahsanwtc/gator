package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ahsanwtc/gator/internal/database"
	"github.com/ahsanwtc/gator/internal/rss"
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

	if len(cmd.parameters) != 1 {
		return fmt.Errorf("login expects 1 argument but got %d", len(cmd.parameters))
	}

	username := cmd.parameters[0]
	_doesUserExists, err := doesUserExists(username, s.db)
	if err != nil {
		fmt.Println("error fetching the user")
		return err
	}

	if !_doesUserExists {
		return fmt.Errorf("user does not exists")
	}

	err = s.config.SetUser(username)
	if err != nil {
		fmt.Println("error setting the active user")
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

	if len(cmd.parameters) != 1 {
		return fmt.Errorf("register expects 1 argument but got %d", len(cmd.parameters))
	}

	username := cmd.parameters[0]
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

func handlerReset(s *State, cmd Command) error {
	if cmd.name != "reset" {
		return fmt.Errorf("wrong command handler")
	}

	if len(cmd.parameters) != 0 {
		return fmt.Errorf("reset expects 0 argument but got %d", len(cmd.parameters))
	}

	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Reset users successful")
	return  nil
}

func handlerUsers(s *State, cmd Command) error {
	if cmd.name != "users" {
		return fmt.Errorf("wrong command handler")
	}

	if len(cmd.parameters) != 0 {
		return fmt.Errorf("users expects 0 argument but got %d", len(cmd.parameters))
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		text := fmt.Sprintf("* %s", user.Name)
		if s.config.CURRENT_USER == user.Name {
			text = fmt.Sprintf("%s (current)", text)
		}
		fmt.Println(text)
	}

	return  nil
}

func handlerAggregate(s *State, cmd Command) error {
	if cmd.name != "agg" {
		return fmt.Errorf("wrong command handler")
	}

	// if len(cmd.parameters) != 0 {
	// 	return fmt.Errorf("users expects 0 argument but got %d", len(cmd.parameters))
	// }

	rssFeed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(rssFeed)
	return  nil
}