package main

import (
	"fmt"
	"os"

	"github.com/ahsanwtc/gator/internal/config"
)

func main () {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	state := State{
		config: &cfg,
	}

	commands := Commands{
		commands: make(map[string]func(*State, Command) error),
	}

	commands.register("login", handlerLogin)
	args := os.Args

	if len(args) < 2 {
		fmt.Println("too few arguments, user `help`")
		os.Exit(1)
	}

	parameters := args[1:]
	
	var commandError error
	switch parameters[0] {
		case "login":
			commandError = commands.run(&state, Command{
				name: "login",
				parameters: parameters,
			})
		default:
			fmt.Printf("`%s`: unknown command\n", parameters[0])
	}

	if commandError != nil {
		fmt.Println("Error: ", commandError)
		os.Exit(1)
	}
}