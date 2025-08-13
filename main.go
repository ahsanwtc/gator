package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ahsanwtc/gator/internal/config"
	"github.com/ahsanwtc/gator/internal/database"
	_ "github.com/lib/pq"
)

func main () {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := sql.Open("postgres", cfg.DB_URL)
	if err != nil {
		fmt.Println("could not connect to the database")
		os.Exit(1)
	}

	state := State{
		config: &cfg,
		db: database.New(db),
	}

	commands := Commands{
		commands: make(map[string]func(*State, Command) error),
	}

	// commands
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	
	args := os.Args

	if len(args) < 2 {
		fmt.Println("too few arguments, user `help`")
		os.Exit(1)
	}
	
	command := args[1]
	parameters := args[2:]
	
	var commandError error
	switch command {
		case "login":
			commandError = commands.run(&state, Command{
				name: "login",
				parameters: parameters,
			})
		case "register":
			commandError = commands.run(&state, Command{
				name: "register",
				parameters: parameters,
			})
		case "reset":
			commandError = commands.run(&state, Command{
				name: "reset",
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