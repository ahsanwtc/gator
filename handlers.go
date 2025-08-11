package main

import "fmt"

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